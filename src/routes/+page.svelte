<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input/index.js";
    import { goto } from '$app/navigation';

    const { data } = $props();
    let roomcode = '';
    let errorMessage = '';
  
    const handleSubmit = async (e: Event) => {
      e.preventDefault();
      errorMessage = '';
  
      if (!roomcode.trim()) {
        // Set an error if it's blank:
        errorMessage = "Please enter a room code.";
        return;
      }
  
      try {
        const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/exists`);
        
        if (!response.ok) {
          // If the server returns 404 Not Found, or something else causes error:
          if (response.status === 404) {
            errorMessage = "No room found with that code.";
          } else {
            errorMessage = `Something went wrong: ${response.statusText}`;
          }
          return;
        }
  
        // If 200 OK, parse response
        const result = await response.json();
  
        // Check if the server’s response indicates the code is valid
        if (result.exists) {
          // Navigate to the event-room if it exists
          goto(`/pages/event-room?code=${roomcode}`);
        } else {
          errorMessage = "That room code doesn't exist.";
        }
      } catch (err) {
        errorMessage = "Failed to contact server. Please try again.";
        console.error(err);
      }
    }

  </script>

<div class="welcome-container">
    <div class="absolute left-4 top-4 md:left-8 md:top-7">
        <a href="/" class="logo">
            <span class="ctf">CTF</span>
            <span class="collab">Collab</span>
        </a>
    </div>
    
    <div class="absolute right-4 top-4 md:right-8 md:top-8">
        {#if data.user}
            <Button href="/pages/admin-page" variant="ghost">Account</Button>
        {:else}
            <Button href="/pages/signin" variant="ghost">Login</Button>
        {/if}
        <span class="separator">|</span>
        <Button href="/pages/about" variant="ghost">About</Button>
        <span class="separator">|</span>
        <Button href="/pages/event-room" variant="ghost">
            Test Event Room
        </Button>
    </div>

    <header>
        <h1>Welcome to CTF-Collab!</h1>
        <p>To join an event, enter room code below or login with existing account</p>
        <div class="form-container" on:submit={handleSubmit}>
            <form class="flex w-full max-w-sm items-center space-x-2">
                <Input type="room-code" bind:value={roomcode} placeholder="enter room-code here..." />
                <Button type="submit">Submit</Button>
            </form>
            
        </div>
        {#if errorMessage}
            <p2 class="mt-2 text-red-500">{errorMessage}</p2>
        {/if}
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

    .logo {
        font-size: 1.5rem;
        font-weight: bold;
    }
    .ctf {
        color: #dc4405;
    }
    .collab {
        color: #666;
        margin-left: 2px;
    }

    header h1 {
        font-size: 2.5rem;
        color: #dc4405;
        margin-bottom: 10px;
    }

    header p {
        font-size: 1.2rem;
        color: #666;
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
        justify-content: center;
        width: 100%;
        margin-top: 20px;
    }
    
    .separator {
        color: #dc4405;
    }
</style>
