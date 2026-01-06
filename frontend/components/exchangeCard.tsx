import { Exchange } from "@/types";
import { formatPrice, formatVolume } from "@/utils/format";
import { Chip } from "@heroui/chip";
import { ArrowDownIcon, ArrowUpIcon } from "lucide-react";

interface ExchangeCardProps {
  exchange: Exchange;
  baseCoin: string;
}

export default function ExchangeCard({
  exchange,
  baseCoin,
}: ExchangeCardProps) {
  const isPositive = exchange.change_percent >= 0;

  return (
    <div className="flex-1 min-w-[280px] space-y-4">
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
}
