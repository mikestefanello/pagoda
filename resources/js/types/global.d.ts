import type { route as routeFn } from "ziggy-js";

declare global {
  const route: typeof routeFn;
}

export type FlashMessages = {
  success?: string[];
  info?: string[];
  warning?: string[];
  danger?: string[];
};
