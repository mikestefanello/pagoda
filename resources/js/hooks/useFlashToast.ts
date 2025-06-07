import { FlashMessages } from "@/types/global";
import { useEffect } from "react";
import { toast } from "sonner";

export function useFlashToasts(flash?: FlashMessages) {
  useEffect(() => {
    if (!flash) return;

    Object.entries(flash).forEach(([type, messages]) => {
      messages.forEach((message) => {
        switch (type) {
          case "success":
            toast.success(message);
            break;
          case "info":
            toast.info(message);
            break;
          case "warning":
            toast.warning(message);
            break;
          case "danger":
          case "error":
            toast.error(message);
            break;
          default:
            toast(message);
        }
      });
    });
  }, [flash]);
}
