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
  import { type SignupFormSchema, signupSchema } from './schema';

  export let data: SuperValidated<SignupFormSchema>;

  const form = superForm(data, {
    dataType: 'json',
    onSubmit: ({}) => {},
    autoFocusOnError: true,
    validators: zodClient(signupSchema),
  });

  const { form: formData, enhance } = form;
</script>

<div class="flex space-x-3">
  <Button variant="outline" title="Sign in with Google" class="bg-background-500 px-6 py-5">
    <IconBrandGoogleFilled class="h-5 w-5" />
  </Button>
  <Button variant="outline" title="Sign in with GitHub" class="bg-background-500 px-6 py-5">
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
    <div class="flex gap-3">
      <Form.Field {form} name="firstName">
        <Form.Control let:attrs>
          <Input
            type="text"
            class="w-56"
            placeholder="First name"
            {...attrs}
            bind:value={$formData.firstName} />
        </Form.Control>
        <Form.Description />
        <Form.FieldErrors style="margin-bottom: 5px;" />
      </Form.Field>

      <Form.Field {form} name="lastName">
        <Form.Control let:attrs>
          <Input
            type="text"
            class="w-56"
            placeholder="Last name"
            {...attrs}
            bind:value={$formData.lastName} />
        </Form.Control>
        <Form.Description />
        <Form.FieldErrors style="margin-bottom: 5px;" />
      </Form.Field>
    </div>

    <Form.Field {form} name="email">
      <Form.Control let:attrs>
        <Input type="email" placeholder="Email address" {...attrs} bind:value={$formData.email} />
      </Form.Control>
      <Form.Description />
      <Form.FieldErrors style="margin-bottom: 5px;" />
    </Form.Field>

    <Form.Field {form} name="password">
      <Form.Control let:attrs>
        <Input type="password" placeholder="Password" {...attrs} bind:value={$formData.password} />
      </Form.Control>
      <Form.Description />
      <Form.FieldErrors />
    </Form.Field>

    <Button type="submit" class="mx-auto mt-5 flex w-full items-center gap-1">
      <IconLockFilled class="h-4 w-4" />
      <span>Sign Up</span>
    </Button>
  </form>
</div>

<Button variant="link" href="/signin" class="text-foreground-200 hover:no-underline">
  Already have an account? Sign In
</Button>
