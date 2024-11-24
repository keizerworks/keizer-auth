"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { AxiosError } from "axios";
import * as React from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";

import {
  VerifyOtpInterface,
  verifyOtpMutationFn,
} from "~/actions/auth/verify-otp";
import { setUser } from "~/global-state/persistant-storage/token";
import { cn } from "~/lib/utils";
import { verifyOtpSchema } from "~/schema/auth";

import { Button } from "../ui/button";
import { Form, FormField } from "../ui/form";
import { InputOTP, InputOTPGroup, InputOTPSlot } from "../ui/input-otp";

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {
  id: string;
}

export function VerifyOtpForm({ id, className, ...props }: UserAuthFormProps) {
  const form = useForm<VerifyOtpInterface>({
    resolver: zodResolver(verifyOtpSchema),
    defaultValues: {
      id,
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: verifyOtpMutationFn,
    onSuccess: (res) => {
      setUser(res);
      toast.success("OTP verified!");
    },
    onError: (err) => {
      let errMessage = "An unknown error occurred.";
      if (err instanceof AxiosError && err.response?.data?.error)
        errMessage = err.response?.data?.error;
      return toast.error(errMessage);
    },
  });

  async function onSubmit(data: VerifyOtpInterface) {
    mutate(data);
  }

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-y-4">
          <div className="flex flex-col items-center gap-4">
            <FormField
              className="space-y-0 text-center mx-auto"
              control={form.control}
              name="otp"
              description="Enter your OTP"
              render={({ field }) => (
                <InputOTP maxLength={6} {...field}>
                  <InputOTPGroup>
                    <InputOTPSlot index={0} />
                    <InputOTPSlot index={1} />
                    <InputOTPSlot index={2} />
                    <InputOTPSlot index={3} />
                    <InputOTPSlot index={4} />
                    <InputOTPSlot index={5} />
                  </InputOTPGroup>
                </InputOTP>
              )}
            />
          </div>

          <Button loading={isPending} type="submit" className="mt-4 w-full">
            Verify
          </Button>
        </form>
      </Form>
    </div>
  );
}
