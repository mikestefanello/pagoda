import { useForm } from "@inertiajs/react";
import { FormEvent } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Link } from "@inertiajs/react";

export default function Register() {
  const { data, setData, post, processing, errors } = useForm({
    name: "",
    email: "",
    password: "",
    password_confirmation: "",
  });

  const submit = (e: FormEvent) => {
    e.preventDefault();
    post("/user/register");
  };

  return (
    <div className="min-h-screen bg-background flex items-center justify-center px-4">
      <Card className="w-full max-w-md border bg-card backdrop-blur-xl shadow-xl">
        <CardHeader className="text-center">
          <CardTitle className="text-3xl font-extrabold tracking-tight text-primary">
            <Link href="/">🥁 Pagode</Link>
          </CardTitle>
          <p className="mt-2 text-sm text-muted-foreground">
            Create your account to join the platform
          </p>
        </CardHeader>

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
              <Label htmlFor="password_confirmation">Confirm Password</Label>
              <Input
                id="password_confirmation"
                type="password"
                value={data.password_confirmation}
                onChange={(e) =>
                  setData("password_confirmation", e.target.value)
                }
                required
                className="bg-muted text-card-foreground placeholder:text-muted-foreground"
              />
              {errors.password_confirmation && (
                <p className="text-sm text-destructive">
                  {errors.password_confirmation}
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
    </div>
  );
}
