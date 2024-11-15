import { z } from "zod";

export const userSchema = z.object({
  last_login: z.string().datetime(),
  created_at: z.string().datetime(),
  updated_at: z.string().datetime(),
  email: z.string().email(),
  first_name: z.string(),
  last_name: z.string().optional(),
  is_verified: z.boolean(),
  is_active: z.boolean(),
});

export type UserInterface = z.infer<typeof userSchema>;
