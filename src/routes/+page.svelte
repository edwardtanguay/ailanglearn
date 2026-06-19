<script lang="ts">
	import { 
		Search, 
		Cpu, 
		Code2, 
		Terminal, 
		Database, 
		Sparkles,
		CheckCircle,
		ArrowUpRight
	} from '@lucide/svelte';

	// Svelte 5 reactive states
	let searchQuery = $state('');
	let selectedTag = $state('All');

	const tags = ['All', 'LLM', 'Prompting', 'RAG', 'Agents'];

	// Real-world notes data
	const notes = [
		{
			id: 'llm-decoding',
			title: 'Understanding LLM Decoding Parameters',
			description: 'Deep dive into token selection parameters: Temperature, Top-P, and Top-K.',
			category: 'LLM',
			updatedAt: '2026-06-12',
			content: [
				{ label: 'Temperature', detail: 'Controls randomness. Lower values (e.g., 0.2) force deterministic, precise output; higher values (e.g., 0.8) increase creativity.' },
				{ label: 'Top-P (Nucleus)', detail: 'Filters tokens to the smallest set whose cumulative probability exceeds P (e.g., 0.90). Helps prune low-probability tokens.' },
				{ label: 'Top-K', detail: 'Restricts selection to the K most probable tokens (e.g., K=40) before performing sampling.' }
			],
			code: `// standard SDK parameters setup
const response = await ai.generateText({
  model: 'gemini-1.5-flash',
  prompt: 'Explain quantum superposition.',
  config: {
    temperature: 0.3,
    topP: 0.95,
    topK: 40
  }
});`
		},
		{
			id: 'prompt-react',
			title: 'The ReAct (Reason + Action) Pattern',
			description: 'How to structure agentic loops that combine reasoning traces with tool execution.',
			category: 'Prompting',
			updatedAt: '2026-06-15',
			content: [
				{ label: 'Thought', detail: 'The model analyzes the user query and outlines what steps or logic it must perform next.' },
				{ label: 'Action', detail: 'The model outputs a tool call (e.g., search_web, calc) with specific parameters.' },
				{ label: 'Observation', detail: 'The environment executes the tool and passes the raw result back to the model context.' }
			],
			code: `Thought: I need to find the current weather in Berlin.
Action: get_weather(location: "Berlin")
Observation: { temp: "22°C", condition: "Sunny" }
Thought: I can now answer the user with the temperature.`
		},
		{
			id: 'rag-chunking',
			title: 'Chunking Strategies for Vector DBs',
			description: 'Optimizing text segmentation to improve semantic search and relevance.',
			category: 'RAG',
			updatedAt: '2026-06-18',
			content: [
				{ label: 'Fixed-size Chunking', detail: 'Simple division by character/token length. Fast but can cut off sentences and split key information.' },
				{ label: 'Semantic Chunking', detail: 'Uses embedding similarity or markdown syntax (headers, bullet points) to group sentences logically.' },
				{ label: 'Overlapping', detail: 'Adding a 10-20% overlap (e.g. 500 token chunks with 100 token overlap) ensures context isn\'t lost at boundaries.' }
			],
			code: `// Overlapping chunk example
const textSplitter = new RecursiveCharacterTextSplitter({
  chunkSize: 500,
  chunkOverlap: 100,
  separators: ["\\n\\n", "\\n", " ", ""]
});`
		},
		{
			id: 'agentic-loops',
			title: 'Autonomous Multi-Agent Systems',
			description: 'Orchestrating specialized nodes working together to complete complex tasks.',
			category: 'Agents',
			updatedAt: '2026-06-19',
			content: [
				{ label: 'Coordinator Agent', detail: 'Breaks down the main task and routes sub-tasks to specialized worker nodes.' },
				{ label: 'Researcher Agent', detail: 'Performs web search and documentation parsing, reporting findings back.' },
				{ label: 'Writer Agent', detail: 'Synthesizes research summaries and drafts the final formatted document.' }
			],
			code: `// Multi-agent routing pseudocode
const result = await coordinator.delegate({
  task: "Research and write report on Next.js 15",
  workers: [researcher, writer]
});`
		}
	];

	// Filtered notes computed reactive block
	let filteredNotes = $derived(
		notes.filter(note => {
			const matchesSearch = note.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
				note.description.toLowerCase().includes(searchQuery.toLowerCase()) ||
				note.category.toLowerCase().includes(searchQuery.toLowerCase());
			
			const matchesTag = selectedTag === 'All' || note.category === selectedTag;

			return matchesSearch && matchesTag;
		})
	);
</script>

<!-- SEO Title & Meta tags in layout / pages -->
<svelte:head>
	<title>Learn AI - Notes & Reference | Fast Follow Learn</title>
	<meta name="description" content="Explore advanced notes on LLM parameters, prompt engineering patterns, chunking strategies for RAG, and autonomous agent loops." />
</svelte:head>

<div class="flex-1 p-6 md:p-12 max-w-5xl mx-auto w-full relative z-10">
	
	<!-- Hero Section -->
	<header class="mb-10 text-left space-y-4">
		<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-indigo-500/10 border border-indigo-500/25 text-indigo-400 text-xs font-semibold uppercase tracking-wider">
			<Sparkles class="w-3.5 h-3.5" />
			<span>Stay on the Front</span>
		</div>
		<h1 class="text-3xl md:text-5xl font-black tracking-tight leading-tight">
			<span class="bg-linear-to-r from-indigo-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">My Notes</span>
			<span class="text-slate-200">on AI</span>
		</h1>
	</header>

		<div class="space-y-6">@@@NOTES</div>

	<!-- Quick Springboard back to Discuss -->
	<footer class="mt-12 p-6 rounded-2xl bg-gradient-to-r from-indigo-900/10 via-purple-900/10 to-pink-900/10 border border-indigo-500/10 flex flex-col md:flex-row items-center justify-between gap-4">
		<div class="space-y-1 text-center md:text-left">
			<h3 class="text-sm font-bold text-slate-200">Want to discuss these concepts live?</h3>
			<p class="text-xs text-slate-400">Join our Discord community or comment on our Fast Follow Forum.</p>
		</div>
		<a href="/discuss" class="inline-flex items-center gap-1.5 px-4 py-2 bg-indigo-600 hover:bg-indigo-500 text-xs font-bold text-white rounded-xl transition-all hover:scale-105 active:scale-95 shadow-[0_0_15px_rgba(99,102,241,0.2)] select-none">
			<span>Go to Discuss AI</span>
			<ArrowUpRight class="w-3.5 h-3.5" />
		</a>
	</footer>

</div>
