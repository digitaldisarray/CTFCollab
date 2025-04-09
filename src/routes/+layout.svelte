<script lang="ts">
  import { onMount } from 'svelte';
	import '../app.css';
    import { afterNavigate, goto } from '$app/navigation';
	let {  children } = $props();

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

	// force a refresh so cookie data can properly be used in store value. 
	afterNavigate(({from, to}) => {
		if(from?.url.pathname.includes("signin") || from?.url.pathname.includes("logout")){
			goto(to?.url.pathname ? to.url.pathname : "/", {invalidateAll: true})
		}
	})
</script>

{@render children()}
