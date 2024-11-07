import { z } from "zod";

export const emailPassSignUpSchema = z.object({
  first_name: z.string().min(3),
  last_name: z.string().optional(),
  email: z.string().email(),
  password: z
    .string()
    .min(8)
    .regex(
      new RegExp(
        /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/,
      ),
      {
        message:
          "Password must contain at least 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character",
      },
    ),
});

export const emailPassSignInSchema = z.object({
  email: z.string().email(),
  password: z
    .string()
    .min(8)
    .regex(
      new RegExp(
        /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/,
      ),
      {
        message:
          "Password must contain at least 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character",
      },
    ),
});

export const emailVerifyCodeSchema = z.object({
  code: z.string().length(6),
  email: z.string().email(),
});
