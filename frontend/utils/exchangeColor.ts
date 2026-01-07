export function exchangeColor(
  exchange: string,
): "warning" | "success" | "default" | "secondary" {
  switch (exchange) {
    case "Kucoin":
      return "success";
    case "Binance":
      return "warning";
    case "OKX":
      return "secondary";
    default:
      return "default";
  }
}
