import { AppContent } from "@/components/AppContent";
import { AppHeader } from "@/components/AppHeader";
import { AppShell } from "@/components/AppShell";
import { type BreadcrumbItem } from "@/types";
import type { PropsWithChildren } from "react";

export default function AppHeaderLayout({
  children,
  breadcrumbs,
}: PropsWithChildren<{ breadcrumbs?: BreadcrumbItem[] }>) {
  return (
    <AppShell>
      <AppHeader breadcrumbs={breadcrumbs} />
      <AppContent>{children}</AppContent>
    </AppShell>
  );
}
