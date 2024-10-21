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

<main class="h-screen w-full overflow-hidden md:shadow-xl relative">
  <GridPattern
    width={50}
    height={50}
    strokeDashArray="1 2"
    class={cn(
      '[mask-image:radial-gradient(800px_circle_at_top,white,transparent)]',
      'opacity-25 -top-5',
    )} />

  <header class="flex justify-between items-center w-[60%] mx-auto mt-9">
    <Button variant="link" href="/" class="hover:no-underline">
      <h1 class="text-xl">
        open<span class="text-purple-400"> pulse</span>
      </h1>
    </Button>
    <nav>
      <ul class="flex gap-8 text-nav-link">
        {#each landingPageNavLinks as link (link.label)}
          <li class="hover:text-foreground-300 transition-colors duration-[300ms] ease-in-out">
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

  <section class="flex flex-col space-y-8 justify-center items-center h-[90%]">
    <img src="/logo.svg" alt="Openpulse logo" class="h-9 w-10 bg-foreground p-2 rounded-md -mb-2" />
    <ShinnyText class="text-2xl">
      {isSignup ? 'Signup to OpenPulse' : 'Signin to OpenPulse'}
    </ShinnyText>
    <slot />
  </section>
</main>
