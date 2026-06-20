<script lang="ts">
	import './layout.css';
	import { BRAND_NAME, BRAND_SUFFIX } from '$lib/config';
	import { page } from '$app/state';
	import { 
		Brain, 
		BookOpen, 
		MessageSquare, 
		Info, 
		Menu, 
		X, 
		ExternalLink,
		NotebookPen
	} from '@lucide/svelte';

	let { children } = $props();

	// Menu state for mobile screen drawer
	let isMenuOpen = $state(false);

	// Navigation items
	const navItems = [
		{ href: '/', label: 'Learn AI', icon: Brain },
		{ href: '/learn-languages', label: 'Learn Languages', icon: BookOpen },
		{ href: '/discuss', label: 'Discuss AI', icon: MessageSquare },
		{ href: '/about', label: 'About', icon: Info }
	];

	function closeMenu() {
		isMenuOpen = false;
	}

	function toggleMenu() {
		isMenuOpen = !isMenuOpen;
	}
</script>

<div class="relative min-h-screen flex flex-col md:flex-row overflow-hidden bg-gray-950 text-gray-100 font-sans">
	
	<!-- Ambient Background Glows -->
	<div class="fixed inset-0 pointer-events-none z-0 overflow-hidden select-none">
		<div class="glow-bg absolute -top-[10%] left-[10%] w-[70%] h-[60%] bg-indigo-500/10 animate-pulse-slow"></div>
		<div class="glow-bg absolute bottom-[10%] right-[10%] w-[60%] h-[60%] bg-purple-500/5 animate-pulse-slow" style="animation-delay: 2s;"></div>
		<div class="glow-bg absolute top-[25%] right-[25%] w-[45%] h-[45%] bg-pink-500/5 animate-pulse-slow" style="animation-delay: 4s;"></div>
	</div>

	<!-- Mobile Top Header Bar -->
	<header class="fixed top-0 left-0 right-0 z-50 flex items-center justify-between px-6 py-4 bg-gray-950/80 backdrop-blur-md border-b border-gray-900 md:hidden">
		<div class="flex flex-col gap-1">
			<a href="/" class="flex items-center gap-3" onclick={closeMenu}>
				<!-- Dynamic SVG Logo -->
				<div class="relative w-8 h-8 shrink-0 rounded-lg bg-gradient-to-tr from-indigo-500 via-purple-500 to-pink-500 p-0.5 shadow-md shadow-indigo-500/20">
					<div class="w-full h-full bg-gray-950 rounded-[6px] flex items-center justify-center">
						<NotebookPen class="w-4 h-4 text-indigo-400" />
					</div>
				</div>
				<div class="flex flex-col">
					<span class="text-[10px] text-gray-400 font-mono tracking-widest uppercase">{BRAND_SUFFIX}</span>
				</div>
			</a>
			<div class="text-[11px] font-mono italic text-yellow-400/90 mt-1 pr-12 md:pr-0">
				<span>This site made by <a href="https://tanguay.info" target="_blank" rel="noopener noreferrer" class="underline text-yellow-300 hover:text-yellow-100 transition-colors font-semibold">Edward</a> with SvelteKit 5.0 and Go, see <a href="https://github.com/edwardtanguay/ailanglearn" target="_blank" rel="noopener noreferrer" class="underline text-yellow-300 hover:text-yellow-100 transition-colors font-semibold">the code</a></span>
			</div>
		</div>
		<button 
			class="p-2 rounded-lg border border-gray-800 bg-gray-900/60 hover:bg-gray-800 text-gray-200 hover:text-white transition-all cursor-pointer z-50"
			onclick={toggleMenu}
			aria-label="Toggle Menu"
		>
			{#if isMenuOpen}
				<X class="w-5 h-5" />
			{:else}
				<Menu class="w-5 h-5" />
			{/if}
		</button>
	</header>

	<!-- Sidebar Navigation Drawer -->
	<!-- Mobile: slides out from the right | Desktop: fixed on the left -->
	<aside class="fixed md:static inset-y-0 right-0 md:left-0 z-40 w-64 
		bg-gray-950/95 backdrop-blur-xl border-l border-white/10 md:border-r md:border-l-0 md:border-gray-900
		transition-transform duration-300 ease-out md:translate-x-0 flex flex-col pt-20 md:pt-0
		before:absolute before:inset-0 before:z-[-1] before:bg-gradient-to-b before:from-gray-950 before:via-purple-950/5 before:to-purple-900/10
		{isMenuOpen ? 'translate-x-0' : 'translate-x-full'}"
	>
		<!-- Desktop Header/Branding -->
		<div class="hidden md:flex flex-col gap-2 px-6 py-8 border-b border-gray-900/50">
			<a href="/" class="flex items-center gap-3">
				<!-- Brand Logo -->
				<div class="relative w-10 h-10 shrink-0 rounded-xl bg-gradient-to-tr from-indigo-500 via-purple-500 to-pink-500 p-0.5 shadow-lg shadow-indigo-500/20">
					<div class="w-full h-full bg-gray-950 rounded-[10px] flex items-center justify-center">
						<NotebookPen class="w-5 h-5 text-indigo-400" />
					</div>
				</div>
				<div class="flex flex-col">
					<span class="text-xs text-slate-400 font-mono tracking-widest uppercase">{BRAND_SUFFIX}</span>
				</div>
			</a>
			<div class="text-[11px] font-mono italic text-yellow-400/90 mt-1">
				<span>This site made by <a href="https://tanguay.info" target="_blank" rel="noopener noreferrer" class="underline text-yellow-300 hover:text-yellow-100 transition-colors font-semibold">Edward</a> with SvelteKit 5.0 and Go, see <a href="https://github.com/edwardtanguay/ailanglearn" target="_blank" rel="noopener noreferrer" class="underline text-yellow-300 hover:text-yellow-100 transition-colors font-semibold">the code</a></span>
			</div>
		</div>

		<!-- Nav Links -->
		<nav class="flex-1 px-4 py-8 space-y-1.5 z-10 overflow-y-auto">
			{#each navItems as item}
				{@const IconComponent = item.icon}
				{@const isActive = page.url.pathname === item.href}
				<a 
					href={item.href} 
					onclick={closeMenu}
					class="flex items-center gap-3.5 px-4 py-3 rounded-xl font-medium text-sm transition-all duration-200 border border-transparent 
					{isActive 
						? 'bg-indigo-500/10 text-indigo-300 border-indigo-500/25 shadow-[0_0_12px_rgba(99,102,241,0.08)]' 
						: 'text-slate-400 hover:text-slate-200 hover:bg-slate-900/60 hover:border-gray-900/30'}"
				>
					<IconComponent class="w-5 h-5 transition-transform duration-200 group-hover:scale-110 {isActive ? 'text-indigo-400' : 'text-slate-500'}" />
					<span>{item.label}</span>
				</a>
			{/each}
		</nav>

		<!-- Sidebar Footer / Brand connection -->
		<div class="p-6 border-t border-gray-900/50 text-center z-10">
			<a 
				href="https://tanguay.info" 
				target="_blank" 
				rel="noopener noreferrer" 
				class="inline-flex items-center justify-center gap-1.5 text-[11px] font-medium text-indigo-400/80 hover:text-indigo-300 transition-colors"
			>
				<span>tanguay.info</span>
				<ExternalLink class="w-3 h-3" />
			</a>
			<p class="text-[10px] text-slate-500 mt-1 font-mono">more projects by Edward</p>
		</div>
	</aside>

	<!-- Mobile Overlay -->
	{#if isMenuOpen}
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div 
			class="fixed inset-0 bg-black/60 backdrop-blur-sm z-30 md:hidden"
			onclick={closeMenu}
		></div>
	{/if}

	<!-- Main content area -->
	<main class="flex-1 min-w-0 z-10 pt-24 md:pt-0 overflow-y-auto h-screen flex flex-col">
		<div class="flex-1 flex flex-col">
			{@render children()}
		</div>
	</main>
</div>
