<script lang="ts">
	type Plain = {
		type: 'plain';
	};

	type Link = {
		type: 'link';
		href: `/${string}`;
	};

	type $$Props = Plain | Link;
</script>

{#if $$props.type === 'link'}
	<a href="/" class="appearance-none">
		<h1 class="text-white font-semibold text-xl">
			open<span class="text-primary-2 text-3xl">pulse</span>
		</h1>
	</a>
{:else}
	<h1 class="text-white font-semibold text-xl">
		open<span class="text-primary-2 text-3xl">pulse</span>
	</h1>
	``
{/if}
