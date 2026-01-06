import { SVGProps } from "react";

export type IconSvgProps = SVGProps<SVGSVGElement> & {
  size?: number;
};

export interface Exchange {
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

export interface CoinData {
  exchanges: Exchange[];
}

export interface PricesData {
  [key: string]: CoinData;
}

export interface ApiResponse {
  success: boolean;
  message: string;
  data: {
    prices: PricesData;
  };
}
