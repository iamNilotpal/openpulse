<script>
	import { landingPageNavLinks } from '$lib/constants/nav-links';
	import JoinWaitlistButton from './join-waitlist-button.svelte';
	import Logo from './logo.svelte';
</script>

<header class="flex justify-between items-center w-[70%] mx-auto py-7">
	<Logo type="link" href="/" />
	<nav>
		<ul class="flex gap-8 text-nav-link">
			{#each landingPageNavLinks as link (link.label)}
				<li class="font-medium transition-colors duration-[400ms] ease-in-out hover:text-white">
					<a href={link.path} class="appearance-none">
						{link.label}
					</a>
				</li>
			{/each}
		</ul>
	</nav>
	<JoinWaitlistButton />
</header>
