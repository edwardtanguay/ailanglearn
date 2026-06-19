<script lang="ts">
	import { 
		Info, 
		Code, 
		Layers, 
		BookOpen, 
		Heart, 
		Globe, 
		ExternalLink,
		Sparkles
	} from '@lucide/svelte';

	const techBadges = [
		{ name: 'Svelte 5', desc: 'Reactive Runes & Compile-time performance' },
		{ name: 'SvelteKit 2', desc: 'Routing, layouts, and server-side rendering' },
		{ name: 'Tailwind CSS v4', desc: 'CSS-first theme configuration & modern utility engine' },
		{ name: 'TypeScript', desc: 'Strict compiler checking and design systems' }
	];
</script>

<svelte:head>
	<title>About the Project</title>
	<meta name="description" content="Learn about Fast Follow Learn: the mission, technology stack, and its relation as an extension of Fast Follow Forum." />
</svelte:head>

<div class="flex-1 p-6 md:p-12 max-w-5xl mx-auto w-full relative z-10">

	<!-- Header -->
	<header class="mb-10 text-left space-y-4">
		<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-indigo-500/10 border border-indigo-500/25 text-indigo-400 text-xs font-semibold uppercase tracking-wider">
			<Info class="w-3.5 h-3.5" />
			<span>Learn constantly</span>
		</div>
		<h1 class="text-3xl md:text-5xl font-black tracking-tight leading-tight">
			<span class="bg-linear-to-r from-indigo-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">About</span>
			<span class="text-slate-200">this site</span>
		</h1>
		<p class="text-slate-400 text-base md:text-lg max-w-2xl font-light">My name is <a href="https://tanguay.info" class="text-indigo-400 hover:text-indigo-300 hover:underline transition-colors" target="_blank" rel="noopener noreferrer">Edward Tanguay</a>. I regularly watch videos and read articles about AI in English, German, Italian, Spanish, Dutch and other languages to (1) stay up-to-date on AI topics and (2) improve my language skills.</p>
		<p class="text-slate-400 text-base md:text-lg max-w-2xl font-light">I record all my notes on this site so they are always online and I can easily search and find them.</p>
		<p class="text-slate-400 text-base md:text-lg max-w-2xl font-light">In 2026, I created this site and the <a href="https://www.fastfollowforum.com" class="text-indigo-400 hover:text-indigo-300 hover:underline transition-colors" target="_blank" rel="noopener noreferrer">Fast Follow Forum</a> which is a Discord server to discuss AI topics.</p>
	</header>

</div>
