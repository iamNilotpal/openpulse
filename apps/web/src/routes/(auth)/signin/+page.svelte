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
  <Input
    min={1}
    required
    type="email"
    aria-required="true"
    placeholder="Email address"
    class="w-80" />
  <Input
    min={1}
    required
    type="password"
    aria-required="true"
    placeholder="Password"
    class="w-80" />
  <Button type="submit" class="mt-3 flex items-center gap-1">
    <IconLockFilled class="h-4 w-4" />
    <span>Sign In</span>
  </Button>
</form>

<Button variant="link" href="/signup" class="text-foreground-200 hover:no-underline">
  Don't have an account? Sign Up
</Button>
