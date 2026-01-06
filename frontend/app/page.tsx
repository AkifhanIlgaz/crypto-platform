"use client";

import { Accordion, AccordionItem } from "@heroui/accordion";
import { Button } from "@heroui/button";
import { Card, CardBody } from "@heroui/card";
import { Chip } from "@heroui/chip";
import { Spinner } from "@heroui/spinner";

import { ArrowDownIcon, ArrowUpIcon, RefreshCw } from "lucide-react";

import { useEffect, useState } from "react";

interface Exchange {
  exchange: string;
  last_updated_at: string;
  price: number;
  high: number;
  low: number;
  open: number;
  close: number;
  base_volume: number;
  quote_volume: number;
  change: number;
  change_percent: number;
}

interface CoinData {
  exchanges: Exchange[];
}

interface PricesData {
  [key: string]: CoinData;
}

interface ApiResponse {
  success: boolean;
  message: string;
  data: {
    prices: PricesData;
  };
}

export default function CryptoPage() {
  const [cryptoData, setCryptoData] = useState<ApiResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [refetching, setRefetching] = useState(false);
  const [lastUpdateTime, setLastUpdateTime] = useState<string>("");

  const fetchCryptoData = async () => {
    try {
      setLoading(true);
      const response = await fetch("http://localhost:7777/api/crypto/prices");
      const data = await response.json();
      setCryptoData(data);
      setLastUpdateTime(new Date().toLocaleString("tr-TR"));
    } catch (err) {
      setError("Veri yüklenirken bir hata oluştu");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleRefetch = async () => {
    try {
      setRefetching(true);
      const response = await fetch("http://localhost:7777/api/crypto/refetch", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (response.ok) {
        // Refetch başarılı, yeni verileri al
        await fetchCryptoData();
      } else {
        setError("Veriler yenilenirken bir hata oluştu");
      }
    } catch (err) {
      setError("Veriler yenilenirken bir hata oluştu");
      console.error(err);
    } finally {
      setRefetching(false);
    }
  };

  useEffect(() => {
    fetchCryptoData();

    // Her 30 saniyede bir veriyi yenile
    const interval = setInterval(fetchCryptoData, 30000);

    return () => clearInterval(interval);
  }, []);

  const formatPrice = (price: number) => {
    return price.toLocaleString("tr-TR", {
      minimumFractionDigits: 2,
      maximumFractionDigits: price < 1 ? 4 : 2,
    });
  };

  const formatVolume = (volume: number) => {
    if (volume >= 1_000_000_000) {
      return `${(volume / 1_000_000_000).toFixed(2)}B`;
    } else if (volume >= 1_000_000) {
      return `${(volume / 1_000_000).toFixed(2)}M`;
    } else if (volume >= 1_000) {
      return `${(volume / 1_000).toFixed(2)}K`;
    }
    return volume.toFixed(2);
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleTimeString("tr-TR", {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    });
  };

  const renderExchangeCard = (exchange: Exchange, baseCoin: string) => {
    const isPositive = exchange.change_percent >= 0;

    return (
      <div className="flex-1 min-w-[280px] space-y-4">
        {/* Exchange Badge and Time */}
        <div className="flex items-center justify-between mb-3">
          <Chip
            color={exchange.exchange === "Binance" ? "warning" : "success"}
            variant="solid"
            size="lg"
            className="font-bold"
          >
            {exchange.exchange}
          </Chip>
        </div>

        {/* Price and Change */}
        <div className="bg-card rounded-xl p-4 border-2 border-border">
          <p className="text-sm text-muted-foreground mb-1">Güncel Fiyat</p>
          <p className="text-3xl font-bold text-foreground mb-2">
            ${formatPrice(exchange.price)}
          </p>
          <div className="flex items-center gap-2">
            <Chip
              color={isPositive ? "success" : "danger"}
              variant="flat"
              size="sm"
              startContent={
                isPositive ? (
                  <ArrowUpIcon className="size-3" />
                ) : (
                  <ArrowDownIcon className="size-3" />
                )
              }
              className="font-semibold"
            >
              {isPositive ? "+" : ""}
              {exchange.change_percent.toFixed(2)}%
            </Chip>
            <span
              className={`text-sm font-medium ${
                isPositive ? "text-success" : "text-danger"
              }`}
            >
              {isPositive ? "+" : ""}${formatPrice(exchange.change)}
            </span>
          </div>
        </div>

        {/* Stats Grid */}
        <div className="grid grid-cols-2 gap-3">
          <div className="bg-card rounded-lg p-3 border border-border shadow-sm">
            <p className="text-xs text-muted-foreground mb-1">Açılış</p>
            <p className="text-sm font-semibold text-foreground">
              ${formatPrice(exchange.open)}
            </p>
          </div>

          <div className="bg-card rounded-lg p-3 border border-border shadow-sm">
            <p className="text-xs text-muted-foreground mb-1">Kapanış</p>
            <p className="text-sm font-semibold text-foreground">
              ${formatPrice(exchange.close)}
            </p>
          </div>

          <div className="bg-success/10 rounded-lg p-3 border border-success/30 shadow-sm">
            <p className="text-xs text-success mb-1">En Yüksek (24H)</p>
            <p className="text-sm font-semibold text-success">
              ${formatPrice(exchange.high)}
            </p>
          </div>

          <div className="bg-danger/10 rounded-lg p-3 border border-danger/30 shadow-sm">
            <p className="text-xs text-danger mb-1">En Düşük (24H)</p>
            <p className="text-sm font-semibold text-danger">
              ${formatPrice(exchange.low)}
            </p>
          </div>
        </div>

        {/* Volume Info */}
        <div className="grid grid-cols-2 gap-3">
          <div className="bg-primary/10 rounded-lg p-3 border border-primary/30">
            <p className="text-xs text-primary mb-1">Base Hacim ({baseCoin})</p>
            <p className="text-sm font-semibold text-primary">
              {formatVolume(exchange.base_volume)}
            </p>
          </div>

          <div className="bg-accent/10 rounded-lg p-3 border border-accent/30">
            <p className="text-xs text-accent mb-1">Quote Hacim (USDT)</p>
            <p className="text-sm font-semibold text-accent">
              ${formatVolume(exchange.quote_volume)}
            </p>
          </div>
        </div>
      </div>
    );
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-background">
        <div className="text-center">
          <Spinner size="lg" color="primary" />
          <p className="mt-4 text-muted-foreground">
            Kripto verileri yükleniyor...
          </p>
        </div>
      </div>
    );
  }

  if (error || !cryptoData?.success) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-background p-4">
        <Card className="max-w-md w-full">
          <CardBody className="text-center py-8">
            <p className="text-danger font-semibold text-lg">
              {error || "Veri yüklenemedi"}
            </p>
          </CardBody>
        </Card>
      </div>
    );
  }

  return (
    <>
      <div className="max-w-7xl mx-auto space-y-6">
        {/* Header */}
        <div className="text-center space-y-2">
          <h1 className="text-4xl font-bold text-foreground">
            Kripto Para Fiyatları
          </h1>
          <p className="text-muted-foreground">
            Anlık kripto para piyasa verileri
          </p>
        </div>

        <div className="flex items-center justify-center gap-4 pt-2">
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <span className="font-medium">Son Güncelleme:</span>
            <span className="text-foreground font-semibold">
              {lastUpdateTime}
            </span>
          </div>
          <Button
            color="primary"
            size="sm"
            variant="flat"
            startContent={
              <RefreshCw
                className={`size-4 ${refetching ? "animate-spin" : ""}`}
              />
            }
            onPress={handleRefetch}
            isLoading={refetching}
            isDisabled={refetching}
          >
            {refetching ? "Yenileniyor..." : "Yeniden Yükle"}
          </Button>
        </div>
      </div>

      <Accordion variant="splitted" selectionMode="multiple" className="px-0">
        {Object.entries(cryptoData.data.prices).map(([coinPair, coinData]) => {
          const [baseCoin] = coinPair.split("/");
          const firstExchange = coinData.exchanges[0];
          const isPositive = firstExchange.change_percent >= 0;

          return (
            <AccordionItem
              key={coinPair}
              aria-label={coinPair}
              startContent={
                <div className="flex items-center gap-3">
                  <span className="text-xl font-bold text-muted-foreground">
                    {coinPair}
                  </span>
                </div>
              }
              subtitle={
                <div className="flex items-center gap-2 mt-1">
                  <span className="text-lg font-bold text-foreground">
                    ${formatPrice(firstExchange.price)}
                  </span>
                  <Chip
                    color={isPositive ? "success" : "danger"}
                    variant="flat"
                    size="sm"
                    startContent={
                      isPositive ? (
                        <ArrowUpIcon className="size-3" />
                      ) : (
                        <ArrowDownIcon className="size-3" />
                      )
                    }
                  >
                    {isPositive ? "+" : ""}
                    {firstExchange.change_percent.toFixed(2)}%
                  </Chip>
                </div>
              }
              classNames={{
                base: "shadow-md hover:shadow-lg transition-shadow",
                title: "text-lg font-semibold",
                trigger: "py-4 px-6",
                content: "px-6 pb-6",
              }}
            >
              <div className="flex flex-col lg:flex-row gap-6 pt-4">
                {coinData.exchanges.map((exchange, idx) => (
                  <div key={`${exchange.exchange}-${idx}`} className="flex-1">
                    {renderExchangeCard(exchange, baseCoin)}
                  </div>
                ))}
              </div>
            </AccordionItem>
          );
        })}
      </Accordion>
    </>
  );
}
