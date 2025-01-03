import { z } from "zod";

import apiClient from "~/axios";
import type { emailPassSignInSchema } from "~/schema/auth";
import { UserInterface } from "~/schema/user";

export const signInMutationFn = async (
  data: z.infer<typeof emailPassSignInSchema>,
) =>
  await apiClient
    .post<UserInterface>("auth/sign-in", data)
    .then((res) => res.data);
