import { z } from "zod";

import apiClient from "~/axios";
import type { emailPassSignUpSchema } from "~/schema/auth";

interface SignUpRes {
  message: string;
}

export const signUpMutationFn = async (
  data: z.infer<typeof emailPassSignUpSchema>,
) =>
  await apiClient.post<SignUpRes>("auth/sign-up", data).then((res) => res.data);
