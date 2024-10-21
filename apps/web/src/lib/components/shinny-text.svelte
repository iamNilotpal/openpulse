<script lang="ts">
	import { cn } from '@/utils';

	let className: string = '';
	export { className as class };
</script>

<h1
	class={cn(
		'text-3xl font-medium inline-flex animate-shine bg-[linear-gradient(110deg,#939393,45%,#1e2631,55%,#939393)] bg-[length:200%_100%] text-transparent bg-clip-text',
		className
	)}
>
	<slot />
</h1>
