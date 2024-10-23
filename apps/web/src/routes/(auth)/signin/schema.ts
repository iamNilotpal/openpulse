import { z } from 'zod';

export const signinSchema = z.object({
  email: z
    .string({ message: 'This field is required.', required_error: 'This field is required.' })
    .email('Email must be valid.'),
});

export type SignInFormSchema = z.infer<typeof signinSchema>;
