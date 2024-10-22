<script lang="ts">
  import {
    IconBrandGithubFilled,
    IconBrandGoogleFilled,
    IconLockFilled,
  } from '@tabler/icons-svelte';
  import { superForm, type SuperValidated } from 'sveltekit-superforms';
  import { zodClient } from 'sveltekit-superforms/adapters';

  import Separator from '@/components/separator.svelte';
  import { Button } from '@/components/ui/button';
  import * as Form from '@/components/ui/form';
  import { Input } from '@/components/ui/input';
  import { type SignInFormSchema, signinSchema } from './schema';

  export let data: SuperValidated<SignInFormSchema>;

  const form = superForm(data, {
    dataType: 'json',
    autoFocusOnError: true,
    validators: zodClient(signinSchema),
  });

  const { form: formData, enhance } = form;
</script>

<div class="space-x-3 flex">
  <Button variant="outline" title="Sign in with Google" class="py-5 px-6 bg-background-500">
    <IconBrandGoogleFilled class="h-5 w-5" />
  </Button>
  <Button variant="outline" title="Sign in with GitHub" class="py-5 px-6 bg-background-500">
    <IconBrandGithubFilled class="h-5 w-5" />
  </Button>
</div>

<div class="mx-auto w-[40%]">
  <Separator gradient>
    <p slot="label" class="px-2">or</p>
  </Separator>
</div>

<div class="flex flex-col gap-3">
  <form method="POST" use:enhance>
    <Form.Field {form} name="email">
      <Form.Control let:attrs>
        <Input
          type="email"
          class="w-96"
          placeholder="Email address"
          {...attrs}
          bind:value={$formData.email} />
      </Form.Control>
      <Form.Description />
      <Form.FieldErrors style="margin-bottom: 5px;" />
    </Form.Field>

    <Form.Field {form} name="password">
      <Form.Control let:attrs>
        <Input
          class="w-96"
          type="password"
          placeholder="Password"
          {...attrs}
          bind:value={$formData.password} />
      </Form.Control>
      <Form.Description />
      <Form.FieldErrors />
    </Form.Field>

    <Button type="submit" class="mx-auto w-full mt-5 flex items-center gap-1">
      <IconLockFilled class="h-4 w-4" />
      <span>Sign In</span>
    </Button>
  </form>
</div>

<Button variant="link" href="/signup" class="text-foreground-200 hover:no-underline">
  Don't have an account? Sign Up
</Button>
