import { createFileRoute, Link } from "@tanstack/react-router";

import { SignUpForm } from "~/components/sign-up/form";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";

export const Route = createFileRoute("/_auth/sign-up")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <Card className="mx-auto shadow-none border-0 w-full sm:w-[350px]">
      <CardHeader className="text-center">
        <CardTitle>Create an account</CardTitle>
        <CardDescription>
          Enter your email below to create your account
        </CardDescription>
      </CardHeader>

      <CardContent>
        <SignUpForm />
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
