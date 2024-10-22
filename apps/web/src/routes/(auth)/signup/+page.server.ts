import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import type { PageServerLoad } from './$types';
import { signupSchema } from './schema';

export const load: PageServerLoad = async () => {
  const form = await superValidate(zod(signupSchema), { strict: true });
  return { form };
};
