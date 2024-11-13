import { createFileRoute, Link } from "@tanstack/react-router";

import { VerifyOtpForm } from "~/components/sign-up/verify-otp";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";

export const Route = createFileRoute("/_auth/verify-otp/$id")({
  component: RouteComponent,
});

function RouteComponent() {
  const params = Route.useParams();

  return (
    <Card className="mx-auto shadow-none border-0 w-full sm:w-[350px]">
      <CardHeader className="text-center">
        <CardTitle>Verify Your Account</CardTitle>
        <CardDescription>
          Enter the OTP sent to your email to verify your account
        </CardDescription>
      </CardHeader>

      <CardContent>
        <VerifyOtpForm id={params.id} />
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
