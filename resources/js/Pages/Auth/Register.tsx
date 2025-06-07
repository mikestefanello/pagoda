import { useForm } from "@inertiajs/react";
import { FormEvent } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Link } from "@inertiajs/react";
import AuthLayout from "@/Layouts/AuthLayout";

export default function Register() {
  const { data, setData, post, processing, errors } = useForm<{
    Name: string;
    Email: string;
    Password: string;
    ConfirmPassword: string;
  }>({
    Name: "",
    Email: "",
    Password: "",
    ConfirmPassword: "",
  });

  const submit = (e: FormEvent) => {
    e.preventDefault();
    post("/user/register");
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
              <Label htmlFor="Name">Name</Label>
              <Input
                id="Name"
                type="text"
                value={data.Name}
                onChange={(e) => setData("Name", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.Name && (
                <p className="text-sm text-destructive">{errors.Name}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="Email">Email</Label>
              <Input
                id="Email"
                type="email"
                value={data.Email}
                onChange={(e) => setData("Email", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.Email && (
                <p className="text-sm text-destructive">{errors.Email}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="Password">Password</Label>
              <Input
                id="Password"
                type="password"
                value={data.Password}
                onChange={(e) => setData("Password", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.Password && (
                <p className="text-sm text-destructive">{errors.Password}</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="ConfirmPassword">Confirm Password</Label>
              <Input
                id="ConfirmPassword"
                type="password"
                value={data.ConfirmPassword}
                onChange={(e) => setData("ConfirmPassword", e.target.value)}
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.ConfirmPassword && (
                <p className="text-sm text-destructive">
                  {errors.ConfirmPassword}
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
