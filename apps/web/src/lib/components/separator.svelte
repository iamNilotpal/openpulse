<script lang="ts">
  import { cn } from '$lib/utils';

  let className: string = '';
  export let gradient: boolean = false;

  export { className as class };
</script>

{#if $$slots.label}
  <div class="flex w-full items-center">
    <div
      class={cn(
        'h-[1px] w-full rounded-full',
        gradient
          ? 'bg-gradient-to-r from-transparent to-zinc-500 dark:from-zinc-800 dark:to-zinc-400'
          : 'bg-zinc-300 dark:bg-zinc-800',
        className,
      )}>
    </div>
    <div class="w-fit text-nowrap text-gray-600 dark:text-gray-300">
      <slot name="label"></slot>
    </div>
    <div
      class={cn(
        'h-[1px] w-full rounded-full',
        gradient
          ? 'bg-gradient-to-r from-zinc-500 to-transparent dark:from-zinc-200 dark:to-zinc-700'
          : 'bg-zinc-300 dark:bg-zinc-800',
        className,
      )}>
    </div>
  </div>
{:else}
  <slot>
    <div
      class={cn(
        'h-[1px] w-full rounded-full',
        gradient
          ? 'bg-gradient-to-r from-transparent via-zinc-500 to-transparent dark:from-zinc-800 dark:via-zinc-200 dark:to-zinc-700'
          : 'bg-zinc-300 dark:bg-zinc-800',
        className,
      )} />
  </slot>
{/if}
