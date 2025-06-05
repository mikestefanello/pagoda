import { Link } from "@inertiajs/react";
import { ReactNode } from "react";

export default function PublicLayout({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen bg-[var(--background)] text-[var(--foreground)] flex flex-col">
      <header className="flex items-center justify-between px-6 py-4 border-b border-[var(--border)]">
        <Link href="/" className="text-xl font-bold text-[var(--primary)]">
          🥁 Pagode
        </Link>

        <div className="space-x-4">
          <Link
            href="/user/login"
            className="text-sm hover:text-[var(--primary)] transition-colors"
          >
            Log in
          </Link>
          <Link
            href="/user/register"
            className="text-sm hover:text-[var(--primary)] transition-colors"
          >
            Register
          </Link>
        </div>
      </header>

      <main className="flex-grow flex items-center justify-center">
        {children}
      </main>
    </div>
  );
}
