<script lang="ts">
  import { cn } from '$lib/utils.js';
  import type { HTMLInputAttributes } from 'svelte/elements';
  import type { InputEvents } from './index.js';

  type $$Props = HTMLInputAttributes;
  type $$Events = InputEvents;

  let className: $$Props['class'] = undefined;
  export let value: $$Props['value'] = undefined;
  export { className as class };

  export let readonly: $$Props['readonly'] = undefined;
</script>

<input
  class={cn(
    'flex h-12 w-full rounded-md border border-white/10 bg-background-500 px-4 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-foreground-300 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-white/10 disabled:cursor-not-allowed disabled:opacity-50',
    className,
  )}
  bind:value
  {readonly}
  on:blur
  on:change
  on:click
  on:focus
  on:focusin
  on:focusout
  on:keydown
  on:keypress
  on:keyup
  on:mouseover
  on:mouseenter
  on:mouseleave
  on:mousemove
  on:paste
  on:input
  on:wheel|passive
  {...$$restProps} />
