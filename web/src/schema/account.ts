import { z } from "zod";

export const accountSchema = z.object({
  name: z.string(),
  logo: z.string().url().optional(),
  id: z.string(),
  created_at: z.date(),
  updated_at: z.date(),
});

export const createAccountSchema = accountSchema
  .omit({
    updated_at: true,
    created_at: true,
    id: true,
  })
  .extend({
    logo: z.any().optional(),
  });

export type AccountInterface = z.infer<typeof accountSchema>;
export type CreateAccountInterface = z.infer<typeof createAccountSchema>;
