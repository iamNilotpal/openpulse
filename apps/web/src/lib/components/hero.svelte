<script>
	import Header from './header.svelte';
</script>

<div
	class="bg-[radial-gradient(ellipse_at_bottom,_var(--tw-gradient-stops))] from-purple-600/20 via-purple-800/20 to-black/20"
>
	<Header />
	<div class="flex justify-center items-center flex-col w-[70%] mx-auto mt-20">
		<h1
			class="text-h1 font-semibold leading-tight bg-gradient-to-bl from-red-300 via-purple-100 to-stone-900 bg-clip-text text-transparent"
		>
			Monitor Uptime.
		</h1>
		<p class="text-caption-2 text-center w-[50%]">
			Get real-time uptime and performance insights with global tracking, proactive monitoring,
			detailed analytics, and instant alerts.
		</p>
		<div class="mt-14 rounded-md">
			<img src="/hero.png" alt="Openpulse hero" class="rounded-md" />
		</div>
	</div>
</div>
