import { useForm } from "@inertiajs/react";
import { FormEvent } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Link } from "@inertiajs/react";
import AuthLayout from "@/Layouts/AuthLayout";

export default function Register() {
  const { data, setData, post, processing, errors } = useForm({
    name: "",
    email: "",
    password: "",
    "password-confirm": "",
  });

  const submit = (e: FormEvent) => {
    e.preventDefault();
    post("/user/register", {
      forceFormData: true,
    });
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
              <Label htmlFor="name">Name</Label>
              <Input
                id="name"
                type="text"
                value={data.name}
                onChange={(e) => setData("name", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.name && (
                <p className="text-sm text-destructive">{errors.name}</p>
              )}
            </div>

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

            <div className="space-y-2">
              <Label htmlFor="password-confirm">Confirm Password</Label>
              <Input
                id="password-confirm"
                type="password"
                value={data["password-confirm"]}
                onChange={(e) => setData("password-confirm", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors["password-confirm"] && (
                <p className="text-sm text-destructive">
                  {errors["password-confirm"]}
                </p>
              )}
            </div>

            <Button
              type="submit"
              disabled={processing}
              className="w-full font-semibold text-primary-foreground bg-primary hover:bg-primary/90 shadow-lg"
            >
              Create Account
            </Button>
          </form>

          <div className="mt-6 text-center text-sm text-muted-foreground">
            Already have an account?{" "}
            <Link
              href="/user/login"
              className="underline text-foreground hover:text-primary"
            >
              Log in
            </Link>
          </div>
        </CardContent>
      </Card>
    </AuthLayout>
  );
}
