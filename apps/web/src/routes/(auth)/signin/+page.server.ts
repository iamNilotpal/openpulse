import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { signinSchema } from './schema';

export const load = async () => {
  const form = await superValidate(zod(signinSchema), { strict: true });
  return { form };
};
