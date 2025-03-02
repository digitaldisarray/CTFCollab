<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input/index.js";

    import { onMount } from 'svelte';

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

<div class="welcome-container">
    <div class="absolute right-4 top-4 md:right-8 md:top-8">
        <Button
            href="/pages/signin"
            variant="ghost"
        >
            Login
        </Button>
        <span class="separator">|</span>
        <Button
            href="/pages/about"
            variant="ghost"
        >
            About
        </Button>
        <span class="separator">|</span>
        <Button
        href="/pages/event-room"
        variant="ghost"
        >
            Test Event Room
        </Button>
    </div>
    <header>
        <h1>Welcome to CTF-Collab!</h1>
        <p>To join an event, enter room code below or login with existing account</p>
        
        <div class="form-container">
            <form class="flex w-full max-w-sm items-center space-x-2">
                <Input type="room-code" placeholder="enter room-code here..." />
                <Button href="/pages/event-room" type="submit">Submit</Button>
            </form>
        </div>
        
    </header>

    <main>
    </main>

    <footer>
        <p>&copy; 2025 CTF-Collab. All rights reserved.</p>
    </footer>
</div>

<style>
    .welcome-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh;
        text-align: center;
        padding: 20px;
    }

    header h1 {
        font-size: 2.5rem;
        color: #333;
        margin-bottom: 10px;
    }

    header p {
        font-size: 1.2rem;
        color: #dc4405;
    }

    main {
        margin-top: 20px;
    }

    footer {
        margin-top: 40px;
        font-size: 0.9rem;
        color: #777;
    }
    .form-container {
        display: flex;
        justify-content: center; /* Center the form horizontally */
        width: 100%;
        margin-top: 20px;
    }
    .separator{
        color: #dc4405;
    }
</style>