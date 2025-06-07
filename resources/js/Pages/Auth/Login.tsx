import { useForm } from "@inertiajs/react";
import { FormEvent } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Link } from "@inertiajs/react";
import AuthLayout from "@/layouts/auth-layout";

type LoginProps = {
  flash: FlashMessages;
};

export default function Login() {
  const { data, setData, post, processing, errors } = useForm({
    email: "",
    password: "",
  });

  const submit = (e: FormEvent) => {
    e.preventDefault();
    post("/user/login");
  };

  return (
    <AuthLayout
      title="Log in to your account"
      description="Crafting seamless digital experiences with InertiaJS, React, and Golang"
      logo="🥁 Pagode"
    >
      <Card className="w-full max-w-md border bg-card backdrop-blur-xl shadow-xl">
        <CardContent>
          <form onSubmit={submit} className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                type="email"
                value={data.email}
                onChange={(e) => setData("email", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.email && (
                <p className="text-sm text-destructive">{errors.email}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                value={data.password}
                onChange={(e) => setData("password", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.password && (
                <p className="text-sm text-destructive">{errors.password}</p>
              )}
            </div>

            <Button
              type="submit"
              disabled={processing}
              className="w-full font-semibold text-primary-foreground 
                bg-gradient-to-br from-primary to-primary 
                hover:from-primary/90 hover:to-primary/80 
                shadow-lg shadow-primary/40 border"
            >
              Log in
            </Button>
          </form>

          <div className="mt-6 text-center text-sm text-muted-foreground">
            Don’t have an account?{" "}
            <Link
              href="/user/register"
              className="underline text-foreground hover:text-primary"
            >
              Register
            </Link>
          </div>
        </CardContent>
      </Card>
    </AuthLayout>
  );
}
