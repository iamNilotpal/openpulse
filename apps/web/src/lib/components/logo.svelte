<script lang="ts">
  import { cn } from '@/utils';

  type Plain = {
    type: 'text';
  };

  type Link = {
    type: 'link';
    href: `/${string}`;
  };

  type $$Props = (Plain | Link) & { class?: string };
</script>

{#if $$props.type === 'link'}
  <a href="/" class="appearance-none">
    <h1 class={cn('text-xl font-semibold text-white', $$props.class)}>
      open<span class="text-primary-2 text-3xl">pulse</span>
    </h1>
  </a>
{:else}
  <h1 class={cn('text-xl font-semibold text-white', $$props.class)}>
    open<span class="text-primary-2 text-3xl">pulse</span>
  </h1>
{/if}
