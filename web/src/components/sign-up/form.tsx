"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { GitHubLogoIcon } from "@radix-ui/react-icons";
import { useMutation } from "@tanstack/react-query";
import { useRouter } from "@tanstack/react-router";
import { AxiosError } from "axios";
import * as React from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

import { signUpMutationFn } from "~/actions/auth/sign-up";
import { cn } from "~/lib/utils";
import { emailPassSignUpSchema } from "~/schema/auth";

import { Button } from "../ui/button";
import { Form, FormField } from "../ui/form";
import { Input } from "../ui/input";
import { PasswordInput } from "../ui/password-input";

type UserAuthFormProps = React.HTMLAttributes<HTMLDivElement>;
type EmailSignUpSchema = z.infer<typeof emailPassSignUpSchema>;

export function SignUpForm({ className, ...props }: UserAuthFormProps) {
  const router = useRouter();

  const form = useForm<EmailSignUpSchema>({
    resolver: zodResolver(emailPassSignUpSchema),
  });

  const { mutate, isPending } = useMutation({
    mutationFn: signUpMutationFn,
    onSuccess: (res) => {
      toast.success(res.message);
      router.navigate({
        to: "/verify-otp/$id",
        params: { id: res.id },
      });
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
                field as keyof EmailSignUpSchema,
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

  async function onSubmit(data: EmailSignUpSchema) {
    mutate(data);
  }

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-y-4">
          <div className="grid grid-cols-2 gap-4">
            <FormField
              control={form.control}
              name="first_name"
              label="First Name"
              render={({ field }) => <Input {...field} />}
            />

            <FormField
              control={form.control}
              name="last_name"
              label="Last Name"
              render={({ field }) => <Input {...field} />}
            />
          </div>

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
            Sign Up
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
        <GitHubLogoIcon className="mr-2 h-4 w-4" /> GitHub
      </Button>
    </div>
  );
}
