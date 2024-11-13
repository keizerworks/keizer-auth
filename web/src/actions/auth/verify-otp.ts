import { z } from "zod";

import apiClient from "~/axios";
import { verifyOtpSchema } from "~/schema/auth";

export type VerifyOtpInterface = z.infer<typeof verifyOtpSchema>;

interface VerifyOtpRes {
  message: string;
}

export async function verifyOtpMutationFn(values: VerifyOtpInterface) {
  return apiClient
    .post<VerifyOtpRes>("auth/verify-otp", values)
    .then((r) => r.data);
}
