<script lang="ts">
  import { onMount } from 'svelte';
	import '../app.css';
	let { children } = $props();

	let isDarkMode = false;

	onMount(() => {
		const systemTheme = window.matchMedia('(prefers-color-scheme: dark)');
		isDarkMode = systemTheme.matches;

		// Apply dark mode class
		document.documentElement.classList.toggle('dark', isDarkMode);

		// Listen for changes in system setting
		systemTheme.addEventListener('change', (event) => {
			isDarkMode = event.matches;
			document.documentElement.classList.toggle('dark', isDarkMode);
		});
	});
</script>

{@render children()}
