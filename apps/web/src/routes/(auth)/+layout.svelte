<script>
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { IconLockFilled } from '@tabler/icons-svelte';

  import GridPattern from '@/components/grid-pattern.svelte';
  import ShinnyText from '@/components/shinny-text.svelte';
  import { Button } from '@/components/ui/button';
  import { landingPageNavLinks } from '@/constants/nav-links';
  import { cn } from '@/utils';

  $: isSignup = $page.url.pathname === '/signup';
</script>

<main class="relative h-screen w-full overflow-hidden md:shadow-xl">
  <GridPattern
    width={50}
    height={50}
    strokeDashArray="1 2"
    class={cn(
      '[mask-image:radial-gradient(800px_circle_at_top,white,transparent)]',
      '-top-5 opacity-25',
    )} />

  <header class="mx-auto mt-9 flex w-[60%] items-center justify-between">
    <Button variant="link" href="/" class="hover:no-underline">
      <h1 class="text-xl">
        open<span class="text-purple-400"> pulse</span>
      </h1>
    </Button>
    <nav>
      <ul class="text-nav-link flex gap-8">
        {#each landingPageNavLinks as link (link.label)}
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

  <section class="flex h-[90%] flex-col items-center justify-center space-y-8">
    <img src="/logo.svg" alt="Openpulse logo" class="-mb-2 h-9 w-10 rounded-md bg-foreground p-2" />
    <ShinnyText class="text-2xl">
      {isSignup ? 'Signup to OpenPulse' : 'Signin to OpenPulse'}
    </ShinnyText>
    <slot />
  </section>
</main>
