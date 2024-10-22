import { z } from 'zod';

export const signinSchema = z.object({
  email: z
    .string({ message: 'This field is required.', required_error: 'This field is required.' })
    .email('Email must be valid.'),
  password: z
    .string({ message: 'This field is required.', required_error: 'This field is required.' })
    .min(8, 'Password must be at least 8 characters long.')
    .max(50, 'Password can only be 50 characters long.'),
});

export type SignInFormSchema = z.infer<typeof signinSchema>;
