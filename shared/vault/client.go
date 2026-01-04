package vault

import (
	"context"
	"fmt"
	"log"
	"time"

	vault "github.com/hashicorp/vault/api"
)

type Client struct {
	client    *vault.Client
	token     string
	renewable bool
	mountPath string
}

type Config struct {
	Address   string `yaml:"address"`
	Token     string `yaml:"token"`
	MountPath string `yaml:"mount_path"`
}

func NewClient(cfg Config) (*Client, error) {
	config := vault.DefaultConfig()
	config.Address = cfg.Address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Vault client: %w", err)
	}

	vaultClient := &Client{
		client:    client,
		renewable: false,
		mountPath: cfg.MountPath,
	}

	if cfg.Token != "" {
		// Development:  Use token directly
		client.SetToken(cfg.Token)
		vaultClient.token = cfg.Token
	} else {
		return nil, fmt.Errorf("no valid authentication method provided")
	}

	return vaultClient, nil
}

func (c *Client) StartTokenRenewal(ctx context.Context) {
	if !c.renewable {
		log.Println("Vault: Token is not renewable, skipping renewal")
		return
	}

	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := c.renewToken(); err != nil {
					log.Printf("Failed to renew Vault token: %v", err)
				}
			case <-ctx.Done():
				log.Println("Vault: Token renewal stopped")
				return
			}
		}
	}()

	log.Println("Vault: Token renewal started")
}

func (c *Client) renewToken() error {
	secret, err := c.client.Auth().Token().RenewSelf(3600)
	if err != nil {
		return err
	}

	log.Printf("Vault: Token renewed, TTL: %d seconds", secret.Auth.LeaseDuration)
	return nil
}

func (c *Client) GetSecret(path string) (map[string]any, error) {
	secret, err := c.client.KVv2(c.mountPath).Get(context.Background(), path)
	if err != nil {
		return nil, fmt.Errorf("unable to get secret from %s/%s: %w", c.mountPath, path, err)
	}
	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("secret not found")
	}
	return secret.Data, nil
}

func (c *Client) SetSecret(path string, data map[string]any) error {
	_, err := c.client.KVv2(c.mountPath).Put(context.Background(), path, data)
	if err != nil {
		return fmt.Errorf("unable to set secret in %s/%s: %w", c.mountPath, path, err)
	}
	return nil
}

// HealthCheck checks Vault connectivity
func (c *Client) HealthCheck() error {
	health, err := c.client.Sys().Health()
	if err != nil {
		return fmt.Errorf("vault health check failed: %w", err)
	}

	if !health.Initialized {
		return fmt.Errorf("vault is not initialized")
	}

	if health.Sealed {
		return fmt.Errorf("vault is sealed")
	}

	return nil
}
