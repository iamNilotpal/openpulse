<script>
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { IconLockFilled } from '@tabler/icons-svelte';

  import { Button } from '@/components/ui/button';
  import { landingPageNavLinks } from '@/constants/nav-links';

  $: isSignup = $page.url.pathname === '/signup';
</script>

<!-- <div
  class="absolute inset-0 -z-10 h-screen w-full items-center px-5 py-24 [background:radial-gradient(125%_125%_at_50%_10%,#000_40%,#63e_100%)]">
</div> -->

<div class="absolute top-0 -z-10 h-screen w-full bg-white">
  <div
    class="absolute bottom-auto left-auto right-0 top-0 h-[500px] w-[500px] -translate-x-[30%] translate-y-[20%] rounded-full bg-[rgba(173,109,244,0.5)] opacity-50 blur-[80px]">
  </div>
</div>

<header class="mx-auto mt-9 flex w-[60%] items-center justify-between">
  <Button variant="link" href="/" class="hover:no-underline">
    <h1 class="text-xl text-background-400">
      open<span class="text-purple-600"> pulse</span>
    </h1>
  </Button>
  <nav>
    <ul class="text-nav-link flex gap-8 text-background-400">
      {#each landingPageNavLinks as link (link.path)}
        <li class="transition-colors duration-[300ms] ease-in-out hover:text-foreground-300">
          <a href={link.path} class="appearance-none">
            {link.label}
          </a>
        </li>
      {/each}
    </ul>
  </nav>
  {#if isSignup}
    <Button class="flex items-center space-x-1" on:click={() => goto('/signin')}>
      <IconLockFilled class="h-4 w-4" />
      <span>Login</span>
    </Button>
  {:else}
    <Button class="flex items-center space-x-1" on:click={() => goto('/signup')}>
      <IconLockFilled class="h-4 w-4" />
      <span>Register</span>
    </Button>
  {/if}
</header>
