import AuthLayoutTemplate from "@/Layouts/Auth/AuthSimpleLayout";
import { useFlashToasts } from "@/hooks/useFlashToast";
import { Toaster } from "@/components/ui/sonner";
import { FlashMessages } from "@/types/global";
import { usePage } from "@inertiajs/react";

export default function AuthLayout({
  children,
  title,
  description,
  logo,
  ...props
}: {
  children: React.ReactNode;
  title: string;
  description: string;
  logo: string;
}) {
  const { flash } = usePage<{ flash: FlashMessages }>().props;

  useFlashToasts(flash);

  return (
    <AuthLayoutTemplate
      title={title}
      description={description}
      logo={logo}
      {...props}
    >
      <Toaster richColors position="top-center" />
      {children}
    </AuthLayoutTemplate>
  );
}
