"use client";

import { Currency } from "@/types";
import { formatPrice } from "@/utils/format";
import { useEffect, useRef } from "react";

interface CurrencyTickerProps {
  currencies: Currency[];
}

export default function CurrencyTicker({ currencies }: CurrencyTickerProps) {
  const scrollRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const scrollContainer = scrollRef.current;
    if (!scrollContainer) return;

    let animationFrameId: number;
    let scrollPosition = 0;
    const scrollSpeed = 0.5; // Pixels per frame

    const animate = () => {
      scrollPosition += scrollSpeed;

      const scrollWidth = scrollContainer.scrollWidth / 2;
      if (scrollPosition >= scrollWidth) {
        scrollPosition = 0;
      }

      scrollContainer.style.transform = `translateX(-${scrollPosition}px)`;
      animationFrameId = requestAnimationFrame(animate);
    };

    animationFrameId = requestAnimationFrame(animate);

    return () => {
      cancelAnimationFrame(animationFrameId);
    };
  }, [currencies]);

  if (!currencies || currencies.length === 0) {
    return null;
  }

  const duplicatedCurrencies = [...currencies, ...currencies];

  return (
    <div className="w-full bg-gradient-to-r from-primary-50 via-primary-100 to-primary-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 border-y border-primary-200 dark:border-gray-700 overflow-hidden py-3">
      <div
        ref={scrollRef}
        className="flex gap-8 whitespace-nowrap"
        style={{ width: "fit-content" }}
      >
        {duplicatedCurrencies.map((currency, index) => (
          <div
            key={`${currency.code}-${index}`}
            className="flex items-center gap-3 px-4"
          >
            <div className="flex items-center gap-2">
              <span className="text-sm font-bold text-primary-700 dark:text-primary-400">
                {currency.code}
              </span>
              <span className="text-xs text-muted-foreground">/</span>
              <span className="text-xs font-medium text-muted-foreground">
                TRY
              </span>
            </div>
            <div className="h-4 w-px bg-primary-300 dark:bg-gray-600" />
            <span className="text-sm font-semibold text-foreground">
              ₺{formatPrice(currency.price)}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}
