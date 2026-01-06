

## 🚀 Kurulum ve Çalıştırma

### Ön Gereksinimler

Sisteminizde aşağıdaki yazılımların kurulu olması gerekmektedir:

- [Docker](https://docs.docker.com/get-docker/) (v20.10+)
- [Docker Compose](https://docs.docker.com/compose/install/) (v2.0+)

*Not: Local development için Go 1.21+ ve Node.js 20+ gereklidir.*

### Adım 1: Projeyi İndirin

```bash
git clone https://github.com/AkifhanIlgaz/crypto-platform.git
cd crypto-platform
```

### Adım 2: ENV Dosyasını Oluşturun ve Gerekli Bilgiler Girin

```bash
cp .env.example .env
```

### Adım 3: Backendi Çalıştırın

```bash
docker compose up -d
```

### Adım 4: Frontendi Çalıştırın

```bash
cd frontend
npm install
npm run dev
```

# http://localhost:3000 adresinden erişebilirsiniz
