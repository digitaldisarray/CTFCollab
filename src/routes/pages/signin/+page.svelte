<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
    let username = '';
    let password = '';
    let errorMessage = '';

    const login = async(e: Event) => {
        e.preventDefault();
        errorMessage = ''; // Clear previous error
        const loginData = {
            "username": username, 
            "password": password
        };
        try {

            const response = await fetch('http://localhost:1337/users/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                // Tell fetch to include cookies
                credentials: 'include',
                body: JSON.stringify(loginData),
            });
            
            if (response.ok) {
                goto('/pages/admin-page');
            } else {
                errorMessage = 'Invalid username or password.';
            }
        } catch (error) {
            errorMessage = 'An error occurred while trying to log in. Please try again later.';
            console.error("Error occured", error);
        }
    }

</script>

<div class="login-container">
    <header>
    </header>

    <main>
        <h1>Log in as administrator</h1>
        <p>Enter credentials below to access an existing admin account</p>
        
        <!-- Display error message if it exists -->
        {#if errorMessage}
            <p class="error-message">{errorMessage}</p>
        {/if}

        <div class="form-container">
            <form class="flex flex-col w-full max-w-sm items-center space-y-2" on:submit={login}>
                <Input type="username" bind:value={username} placeholder="Username" />
                <Input type="password" bind:value={password} placeholder="Password" />
                <Button type="submit">Log In</Button>
            </form>
        </div>
    </main>

    <footer>
        <p>&copy; 2025 CTF-Collab. All rights reserved.</p>
    </footer>
</div>

<style>
    .login-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh;
        text-align: center;
        padding: 20px;
    }

    main h1 {
        font-size: 2.5rem;
        color: #333;
        margin-bottom: 10px;
    }

    main p {
        font-size: 1.2rem;
        color: #dc4405;
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

    .error-message {
        color: red;
        font-size: 0.9rem;
        margin-top: 10px;
    }
</style>