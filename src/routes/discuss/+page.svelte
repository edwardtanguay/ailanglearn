<script lang="ts">
	import {
		MessageSquare,
		ExternalLink,
		Users,
		Flame,
		ArrowRight,
		PlusCircle,
		Sparkles
	} from '@lucide/svelte';

	// Mock forum threads for visual fidelity
	const mockThreads = [
		{
			title: 'Best chunking strategy for long PDF document scanning in RAG?',
			author: 'Klaus_D',
			replies: 18,
			category: 'RAG & Vector DBs',
			hot: true
		},
		{
			title: 'Berlin AI Meetup - July 2026: Demos & Drinks',
			author: 'Edward_T',
			replies: 42,
			category: 'Events & Meetups',
			hot: true
		},
		{
			title: 'Deploying deepseek-r1 locally: quantization tips',
			author: 'Sarah_Codes',
			replies: 9,
			category: 'Local LLMs',
			hot: false
		}
	];
</script>

<svelte:head>
	<title>Discuss AI</title>
	<meta
		name="description"
		content="Connect with the Fast Follow Forum AI Community. Join our Discord or discuss latest AI tools and trends on the forum."
	/>
</svelte:head>

<div class="flex-1 p-6 md:p-12 max-w-5xl mx-auto w-full relative z-10">
	<!-- Header -->
	<header class="mb-10 text-left space-y-4">
		<div
			class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-pink-500/10 border border-pink-500/25 text-pink-400 text-xs font-semibold uppercase tracking-wider"
		>
			<MessageSquare class="w-3.5 h-3.5" />
			<span>Join the Discussion</span>
		</div>
		<h1
			class="text-3xl md:text-5xl font-black tracking-tight leading-tight"
		>
			<span
				class="bg-gradient-to-r from-indigo-400 via-purple-400 to-pink-400 bg-clip-text text-transparent"
				>Discuss AI</span
			>
			<span class="text-slate-200">with us</span>
		</h1>
		<p class="text-slate-400 text-base md:text-lg max-w-2xl font-light">
			Fast Follow Forum is a Discord Server that I created in January 2026 as a
			place to discuss AI. Discord Server which has a nice UX
			for creating various topics, posting videos/screenshots, which
			allows a group to stay up-to-date on various topics. It also allows
			members to create planned or spontaneous meetings via video
			conference and screen sharing.
		</p>
	</header>

	<div class="max-w-2xl">
		<a
			href="https://fastfollow.vercel.app"
			target="_blank"
			rel="noopener noreferrer"
			class="w-full inline-flex items-center justify-center gap-2.5 px-8 py-4 font-bold text-white transition-all duration-300 bg-linear-to-r from-indigo-600 to-purple-600 hover:from-indigo-500 hover:to-purple-500 rounded-xl hover:scale-[1.02] active:scale-[0.98] shadow-[0_0_20px_rgba(99,102,241,0.25)] hover:shadow-[0_0_30px_rgba(99,102,241,0.45)] select-none text-sm uppercase tracking-wider"
		>
			<span>Join Fast Follow Forum</span>
			<ExternalLink class="w-4 h-4" />
		</a>
	</div>
</div>
