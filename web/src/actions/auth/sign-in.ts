import { z } from "zod";

import apiClient from "~/axios";
import type { emailPassSignInSchema } from "~/schema/auth";

interface SignInRes {
  message: string;
}

export const signInMutationFn = async (
  data: z.infer<typeof emailPassSignInSchema>,
) =>
  await apiClient.post<SignInRes>("auth/sign-in", data).then((res) => res.data);
