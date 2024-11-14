import { createFileRoute, Link } from "@tanstack/react-router";

import { SignInForm } from "~/components/sign-in/form";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";

export const Route = createFileRoute("/_auth/sign-in")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <Card className="mx-auto shadow-none border-0 w-full sm:w-[350px]">
      <CardHeader className="text-center">
        <CardTitle>Sign In to Your Dashboard</CardTitle>
        <CardDescription>
          Please enter your email and password to access your dashboard
        </CardDescription>
      </CardHeader>

      <CardContent>
        <SignInForm />
      </CardContent>

      <CardFooter className="text-center">
        <CardDescription>
          By clicking continue, you agree to our{" "}
          <Link
            to="/"
            className="underline underline-offset-4 hover:text-primary"
          >
            Terms of Service
          </Link>{" "}
          and{" "}
          <Link
            to="/"
            className="underline underline-offset-4 hover:text-primary"
          >
            Privacy Policy
          </Link>
          .
        </CardDescription>
      </CardFooter>
    </Card>
  );
}
