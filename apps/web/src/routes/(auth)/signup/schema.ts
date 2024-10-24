import { z } from 'zod';

export const signupSchema = z.object({
  firstName: z
    .string({ message: 'This field is required.', required_error: 'This field is required.' })
    .max(50, 'First name can only be 50 characters long.'),
  lastName: z
    .string({ message: 'This field is required.', required_error: 'This field is required.' })
    .max(50, 'Last name can only be 50 characters long.'),
  email: z
    .string({ message: 'This field is required.', required_error: 'This field is required.' })
    .email('Email must be valid.'),
});

export type SignupFormSchema = z.infer<typeof signupSchema>;
