<script>
  import Separator from '@/components/separator.svelte';
  import { Button } from '@/components/ui/button';
  import { Input } from '@/components/ui/input';
  import {
    IconBrandGithubFilled,
    IconBrandGoogleFilled,
    IconLockFilled,
  } from '@tabler/icons-svelte';
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

<form class="flex flex-col gap-3">
  <div class="flex gap-3">
    <Input
      min={1}
      required
      type="text"
      aria-required="true"
      placeholder="First name"
      class="w-56" />

    <Input type="text" min={1} required aria-required="true" placeholder="Last name" class="w-56" />
  </div>
  <Input min={1} required type="email" aria-required="true" placeholder="Email address" />
  <Input min={1} required type="password" aria-required="true" placeholder="Password" />
  <Button type="submit" class="mt-3 flex items-center gap-1">
    <IconLockFilled class="h-4 w-4" />
    <span>Sign Up</span>
  </Button>
</form>

<Button variant="link" href="/signin" class="text-foreground-200 hover:no-underline">
  Already have an account? Sign In
</Button>
