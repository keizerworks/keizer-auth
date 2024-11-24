import { z } from "zod";

import apiClient from "~/axios";
import { verifyOtpSchema } from "~/schema/auth";
import { UserInterface } from "~/schema/user";

export type VerifyOtpInterface = z.infer<typeof verifyOtpSchema>;

export async function verifyOtpMutationFn(values: VerifyOtpInterface) {
  return apiClient
    .post<UserInterface>("auth/verify-otp", values)
    .then((r) => r.data);
}
