"use client";

import CurrencyTicker from "@/components/currenyTicker";
import ExchangeCard from "@/components/exchangeCard";
import { ThemeSwitch } from "@/components/themeSwitch";
import { ApiResponse, CryptoPrices, CurrencyPrices } from "@/types";
import { formatDate, formatPrice } from "@/utils/format";
import { Accordion, AccordionItem } from "@heroui/accordion";
import { Button } from "@heroui/button";
import { Card, CardBody } from "@heroui/card";
import { Chip } from "@heroui/chip";
import { Spinner } from "@heroui/spinner";
import { addToast } from "@heroui/toast";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { ArrowDownIcon, ArrowUpIcon, RefreshCw } from "lucide-react";

export default function CryptoPage() {
  const queryClient = useQueryClient();

  const {
    data: cryptoData,
    isLoading: isCryptoLoading,
    error,
    dataUpdatedAt,
  } = useQuery<ApiResponse<CryptoPrices> | null>({
    queryKey: ["crypto"],
    queryFn: async () => {
      try {
        const res = await fetch("http://localhost:7777/api/crypto/prices");

        if (!res.ok) {
          throw new Error(
            `HTTP error! status: ${res.status} - ${res.statusText}`,
          );
        }

        const data = await res.json();

        if (!data || !data.success) {
          throw new Error(data?.message || "Invalid response from server");
        }

        return data;
      } catch (error) {
        console.error("Error fetching crypto prices:", error);
        throw error;
      }
    },
    staleTime: Infinity,
    refetchOnWindowFocus: false,
    retry: 3,
  });

  const {
    data: currencyData,
    isLoading: isCurrencyLoading,
    error: currencyError,
  } = useQuery<ApiResponse<CurrencyPrices> | null>({
    queryKey: ["currency"],
    queryFn: async () => {
      try {
        const res = await fetch("http://localhost:7777/api/currency/prices");

        if (!res.ok) {
          throw new Error(
            `HTTP error! status: ${res.status} - ${res.statusText}`,
          );
        }

        const data = await res.json();

        if (!data || !data.success) {
          throw new Error(data?.message || "Invalid response from server");
        }

        return data;
      } catch (error) {
        console.error("Error fetching currency prices:", error);
        throw error;
      }
    },
    staleTime: Infinity,
    refetchOnWindowFocus: false,
    retry: 3,
  });

  const { mutate: refetchCryptoPrices, isPending: isRefetchPending } =
    useMutation({
      mutationFn: async () => {
        const cryptoRes = await fetch(
          "http://localhost:7777/api/crypto/refetch",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
          },
        );

        if (cryptoRes.ok) {
          await queryClient.invalidateQueries({
            queryKey: ["crypto"],
          });
        } else {
          throw new Error("Refetch failed");
        }

        const currencyRes = await fetch(
          "http://localhost:7777/api/currency/refetch",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
          },
        );

        if (currencyRes.ok) {
          await queryClient.invalidateQueries({
            queryKey: ["currency"],
          });
        } else {
          throw new Error("Refetch failed");
        }
      },
      onSuccess: (data) => {
        addToast({
          title: "Fiyatlar Güncellendi",
          description:
            "Kripto ve para birimleri fiyatları başarıyla güncellendi.",
          color: "success",
        });
      },
      onError: (error) => {
        addToast({
          title: "Hata",
          description: "Beklenmedik bir hata oluştu.",
          color: "danger",
        });
      },
    });

  if (isCryptoLoading || isCurrencyLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-background">
        <div className="text-center">
          <Spinner size="lg" color="primary" />
          <p className="mt-4 text-muted-foreground">Veriler yükleniyor...</p>
        </div>
      </div>
    );
  }

  if (
    error ||
    currencyError ||
    !cryptoData?.success ||
    !currencyData?.success
  ) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-background p-4">
        <Card className="max-w-md w-full">
          <CardBody className="text-center py-8">
            <p className="text-danger font-semibold text-lg">
              {error?.message || "Veri yüklenemedi"}
            </p>
          </CardBody>
        </Card>
      </div>
    );
  }

  return (
    <>
      {currencyData?.success && currencyData.data.currencies && (
        <div className="sticky top-0 z-50 shadow-sm mb-8">
          <CurrencyTicker currencies={currencyData.data.currencies} />
        </div>
      )}
      <div className=" mx-auto space-y-6">
        <div className="text-center space-y-2">
          <h1 className="text-4xl font-bold text-foreground">
            Kripto Para Fiyatları
          </h1>
          <p className="text-muted-foreground">
            Anlık kripto para piyasa verileri
          </p>
        </div>

        <div className="flex items-center justify-end gap-4 pt-2 mb-2">
          <ThemeSwitch />

          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <span className="font-medium">Son Güncelleme:</span>
            <span className="text-foreground font-semibold">
              {formatDate(dataUpdatedAt)}
            </span>
          </div>
          <Button
            color="primary"
            size="sm"
            variant="light"
            startContent={
              <RefreshCw
                className={`size-4 ${isRefetchPending ? "animate-spin" : ""}`}
              />
            }
            onPress={() => refetchCryptoPrices()}
            isDisabled={isRefetchPending}
          >
            {isRefetchPending ? "Yenileniyor..." : "Yeniden Yükle"}
          </Button>
        </div>
      </div>

      <Accordion variant="splitted" selectionMode="multiple" className="px-0">
        {(cryptoData?.data.prices
          ? Object.entries(cryptoData?.data.prices)
          : []
        ).map(([coinPair, coinData]) => {
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
                    {firstExchange.change_percent?.toFixed(2) ?? "N/A"}%
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
                    <ExchangeCard exchange={exchange} baseCoin={baseCoin} />
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
