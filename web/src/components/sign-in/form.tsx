"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { GitHubLogoIcon } from "@radix-ui/react-icons";
import { useMutation } from "@tanstack/react-query";
import { Link, useRouter } from "@tanstack/react-router";
import { AxiosError } from "axios";
import * as React from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

import { signInMutationFn } from "~/actions/auth/sign-in";
import { setUser } from "~/global-state/persistant-storage/token";
import { cn } from "~/lib/utils";
import { emailPassSignInSchema } from "~/schema/auth";

import { Button } from "../ui/button";
import { Form, FormField } from "../ui/form";
import { Input } from "../ui/input";
import { PasswordInput } from "../ui/password-input";
import { Separator } from "../ui/separator";

type UserAuthFormProps = React.HTMLAttributes<HTMLDivElement>;
type EmailSignInSchema = z.infer<typeof emailPassSignInSchema>;

export function SignInForm({ className, ...props }: UserAuthFormProps) {
  const router = useRouter();
  const form = useForm<EmailSignInSchema>({
    resolver: zodResolver(emailPassSignInSchema),
  });

  const { mutate, isPending } = useMutation({
    mutationFn: signInMutationFn,
    onSuccess: (res) => {
      console.log(res);
      setUser(res);
      toast.success("Logged in successfully");
      router.navigate({ to: "/" });
    },
    onError: (err) => {
      if (err instanceof AxiosError) {
        if (err.response?.data?.errors) {
          let shouldFocus = true;
          const validationErrors = err.response.data.errors;

          return Object.keys(validationErrors).forEach((field) => {
            const fieldErrors = validationErrors[field];
            const errorMessage = Object.values(fieldErrors).find(
              (message) => message !== "",
            ) as string;

            if (errorMessage) {
              form.setError(
                field as keyof EmailSignInSchema,
                { message: errorMessage },
                { shouldFocus: shouldFocus },
              );
              shouldFocus = false;
            }
          });
        }

        return toast.error(
          err.response?.data?.error || "An unknown error occurred.",
        );
      }

      return toast.error("An unknown error occurred.");
    },
  });

  async function onSubmit(data: EmailSignInSchema) {
    mutate(data);
  }

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-y-4">
          <FormField
            control={form.control}
            name="email"
            label="Email"
            render={({ field }) => <Input {...field} />}
          />

          <FormField
            control={form.control}
            name="password"
            label="Passowrd"
            render={({ field }) => <PasswordInput {...field} />}
          />

          <Button loading={isPending} type="submit" className="mt-4 w-full">
            Sign In
          </Button>
        </form>
      </Form>

      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t" />
        </div>
        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-background px-2 text-muted-foreground">
            Or continue with
          </span>
        </div>
      </div>

      <Button disabled variant="outline" type="button">
        <GitHubLogoIcon className="size-4" /> GitHub
      </Button>

      <small className="text-muted-foreground text-center">
        Don't have an account?{" "}
        <Link
          className="text-primary underline-offset-4 hover:underline"
          to="/sign-up"
        >
          Sign up
        </Link>
      </small>
      <Separator />
    </div>
  );
}
