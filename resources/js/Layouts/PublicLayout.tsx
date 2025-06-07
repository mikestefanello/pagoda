import { Link } from "@inertiajs/react";
import { ReactNode } from "react";

export default function PublicLayout({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen text-foreground flex flex-col">
      <header className="flex items-center justify-end px-6 py-4 border-b border">
        <div className="space-x-4">
          <Link
            href="/user/login"
            className="text-sm hover:text-primary transition-colors"
          >
            Log in
          </Link>
          <Link
            href="/user/register"
            className="text-sm hover:text-primary transition-colors"
          >
            Register
          </Link>
        </div>
      </header>

      <main className="flex-grow flex items-center justify-center flex-col">
        {children}
      </main>
    </div>
  );
}
