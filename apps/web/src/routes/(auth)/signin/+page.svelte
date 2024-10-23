<script lang="ts">
  import {
    IconBrandGithubFilled,
    IconBrandGoogleFilled,
    IconCloverFilled,
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

    <Button type="submit" class="mx-auto mt-4 flex w-full items-center gap-1">
      <IconCloverFilled class="h-4 w-4" />
      <span>Send Magic Link</span>
    </Button>
  </form>
</div>

<Button variant="link" href="/signup" class="text-foreground-200 hover:no-underline">
  Don't have an account? Sign Up
</Button>
