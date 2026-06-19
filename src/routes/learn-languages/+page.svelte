<script lang="ts">
	import { 
		Languages, 
		GraduationCap, 
		HelpCircle, 
		Bookmark, 
		CheckCircle2, 
		Sparkles 
	} from '@lucide/svelte';

	// Svelte 5 state
	let activeLanguage = $state('German');
	const languages = ['German', 'Spanish', 'French'];

	// Real-world structured notes for language study
	const languageData: Record<string, {
		description: string;
		flag: string;
		grammarTipTitle: string;
		grammarTipText: string;
		conjugationTitle: string;
		conjugationVerb: string;
		conjugationRows: { pronoun: string; form: string; translation: string }[];
		vocabulary: { word: string; translation: string; type: string; gender?: 'masculine' | 'feminine' | 'neuter' }[];
	}> = {
		German: {
			description: 'Study notes for German grammar structure, articles, and essential conversational verbs.',
			flag: '🇩🇪',
			grammarTipTitle: 'German Noun Gender Articles',
			grammarTipText: 'All German nouns have genders: Masculine (der), Feminine (die), and Neuter (das). Genders are grammatical, not biological (e.g., "das Mädchen" - the girl is neuter). Tips: Nouns ending in -ung, -heit, -keit, -schaft, or -tät are always Feminine (die). Nouns ending in -chen or -lein are always Neuter (das).',
			conjugationTitle: 'Verb Conjugation: sein (to be)',
			conjugationVerb: 'sein • Irregular Verb • Present Tense',
			conjugationRows: [
				{ pronoun: 'ich', form: 'bin', translation: 'I am' },
				{ pronoun: 'du', form: 'bist', translation: 'you are (informal)' },
				{ pronoun: 'er / sie / es', form: 'ist', translation: 'he / she / it is' },
				{ pronoun: 'wir', form: 'sind', translation: 'we are' },
				{ pronoun: 'ihr', form: 'seid', translation: 'you are (plural)' },
				{ pronoun: 'sie / Sie', form: 'sind', translation: 'they / You (formal) are' }
			],
			vocabulary: [
				{ word: 'der Apfel', translation: 'apple', type: 'Noun', gender: 'masculine' },
				{ word: 'die Zeitung', translation: 'newspaper', type: 'Noun', gender: 'feminine' },
				{ word: 'das Buch', translation: 'book', type: 'Noun', gender: 'neuter' },
				{ word: 'sprechen', translation: 'to speak', type: 'Verb' },
				{ word: 'langsam', translation: 'slowly', type: 'Adverb' },
				{ word: 'wichtig', translation: 'important', type: 'Adjective' }
			]
		},
		Spanish: {
			description: 'Study notes for Spanish sentence structure, verb moods, and essential nouns.',
			flag: '🇪🇸',
			grammarTipTitle: 'Ser vs Estar (The "To Be" Distinction)',
			grammarTipText: 'Spanish uses two verbs for "to be". Use SER for permanent or descriptive attributes (Identity, Occupation, Origin, Time - acronym DOCTOR). Use ESTAR for temporary states or locations (Position, Location, Action, Condition, Emotion - acronym PLACE).',
			conjugationTitle: 'Verb Conjugation: estar (to be - location/state)',
			conjugationVerb: 'estar • Irregular Verb • Present Tense',
			conjugationRows: [
				{ pronoun: 'yo', form: 'estoy', translation: 'I am' },
				{ pronoun: 'tú', form: 'estás', translation: 'you are (informal)' },
				{ pronoun: 'él / ella / Ud.', form: 'está', translation: 'he / she / it is' },
				{ pronoun: 'nosotros', form: 'estamos', translation: 'we are' },
				{ pronoun: 'vosotros', form: 'estáis', translation: 'you are (plural)' },
				{ pronoun: 'ellos / ellas / Uds.', form: 'están', translation: 'they / You are' }
			],
			vocabulary: [
				{ word: 'el coche', translation: 'car', type: 'Noun', gender: 'masculine' },
				{ word: 'la mesa', translation: 'table', type: 'Noun', gender: 'feminine' },
				{ word: 'entender', translation: 'to understand', type: 'Verb' },
				{ word: 'rápidamente', translation: 'quickly', type: 'Adverb' },
				{ word: 'hermoso', translation: 'beautiful', type: 'Adjective' },
				{ word: 'el agua', translation: 'water', type: 'Noun', gender: 'feminine' }
			]
		},
		French: {
			description: 'Study notes for French spelling rules, irregular conjugations, and core phrases.',
			flag: '🇫🇷',
			grammarTipTitle: 'French Adjective Agreement',
			grammarTipText: 'In French, adjectives must agree in gender (masculine/feminine) and number (singular/plural) with the noun they describe. Usually, add "-e" for feminine and "-s" for plural (e.g. "un livre vert" -> "une pomme verte"). Placement: most adjectives go AFTER the noun they describe.',
			conjugationTitle: 'Verb Conjugation: avoir (to have)',
			conjugationVerb: 'avoir • Auxiliary Verb • Present Tense',
			conjugationRows: [
				{ pronoun: 'j\'', form: 'ai', translation: 'I have' },
				{ pronoun: 'tu', form: 'as', translation: 'you have (informal)' },
				{ pronoun: 'il / elle / on', form: 'a', translation: 'he / she / one has' },
				{ pronoun: 'nous', form: 'avons', translation: 'we have' },
				{ pronoun: 'vous', form: 'avez', translation: 'you have (plural/formal)' },
				{ pronoun: 'ils / elles', form: 'ont', translation: 'they have' }
			],
			vocabulary: [
				{ word: 'le stylo', translation: 'pen', type: 'Noun', gender: 'masculine' },
				{ word: 'la maison', translation: 'house', type: 'Noun', gender: 'feminine' },
				{ word: 'comprendre', translation: 'to understand', type: 'Verb' },
				{ word: 'toujours', translation: 'always', type: 'Adverb' },
				{ word: 'facile', translation: 'easy', type: 'Adjective' },
				{ word: 'la fleur', translation: 'flower', type: 'Noun', gender: 'feminine' }
			]
		}
	};

	let selectedData = $derived(languageData[activeLanguage]);
</script>

<svelte:head>
	<title>Learn Languages - Notes & Conjugations | Fast Follow Learn</title>
	<meta name="description" content="Study conjugation tables, essential vocabulary, and core grammar notes for German, Spanish, and French." />
</svelte:head>

<div class="flex-1 p-6 md:p-12 max-w-5xl mx-auto w-full relative z-10">

	<!-- Hero Header -->
	<header class="mb-10 text-left space-y-4">
		<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-purple-500/10 border border-purple-500/25 text-purple-400 text-xs font-semibold uppercase tracking-wider">
			<GraduationCap class="w-3.5 h-3.5" />
			<span>Multilingual</span>
		</div>
		<h1 class="text-3xl md:text-5xl font-black tracking-tight leading-tight">
			<span class="bg-gradient-to-r from-indigo-400 via-purple-400 to-pink-400 bg-clip-text text-transparent">Language</span>
			<span class="text-slate-200">Notes</span>
		</h1>
		<p class="text-slate-400 text-base md:text-lg max-w-2xl font-light">
			I watch videos and read articles in various languages and record useful phrases here that I encounter in order to transfer them from my passive to active vocabulary.
		</p>
	</header>

	<div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">@@@LANGUAGE_NOTES</div>

</div>
