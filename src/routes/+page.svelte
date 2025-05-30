<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input/index.js";
    import { goto } from '$app/navigation';

    export let data: {
        user: User | null;
    };
    
    let roomcode = '';
    let errorMessage = '';
    let showNicknamePopup = false;
    let nickname = '';
    let nicknameError = ''; // <-- for popup errors
  
    const handleSubmit = async (e: Event) => {
        e.preventDefault();
        errorMessage = '';

        // Check if blank room code
        if (!roomcode.trim()) {
            errorMessage = "Please enter a room code.";
            return;
        }

        try {
            const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/exists`);
            
            if (!response.ok) {
                if (response.status === 404) {
                    errorMessage = "No room found with that code.";
                } else {
                    errorMessage = `Something went wrong: ${response.statusText}`;
                }
                return;
            }

            // Parse the response
            const result = await response.json();

            // Check if the server’s response indicates the code is valid
            if (result.exists) {
                // Ask for nickname if code was valid
                showNicknamePopup = true;
            } else {
                errorMessage = "That room code doesn't exist.";
            }
        } catch (err) {
            errorMessage = "Failed to contact server. Please try again.";
            console.error(err);
        }
    }

    const closePopup = () => {
        showNicknamePopup = false;
        nicknameError = '';
        nickname = '';
    };

    const handleNicknameSubmit = async (e: Event) => {
        e.preventDefault();
        nicknameError = '';
        if (!nickname.trim()) {
            nicknameError = "Please enter a nickname.";
            return;
        }

        try {
            const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/join-as-guest`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
                body: JSON.stringify({ nickname }),
            });
            if (response.ok) {
                goto(`/pages/event-room?code=${roomcode}&nickname=${encodeURIComponent(nickname)}`);
            } else {
                const data = await response.json();
                nicknameError = data.error || "Failed to join as guest.";
            }
        } catch (err) {
            nicknameError = "Failed to contact server. Please try again.";
            console.error(err);
        }
    };
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
            <Button href="/pages/admin-page" variant="ghost">Dashboard</Button>
        {:else}
            <Button href="/pages/signin" variant="ghost">Login</Button>
        {/if}
        <span class="separator">|</span>
        <Button href="/pages/about" variant="ghost">About</Button>
    </div>

    <header>
    </header>

    <main>
        <h1>Welcome to CTF-Collab!</h1>
        <p>To join an event, enter room code below or login with existing account</p>
        <div class="form-container" on:submit={handleSubmit}>
            {#if showNicknamePopup}
                <div class="popup-overlay">
                    <div class="popup-content">
                        <h2>Enter Your Nickname</h2>
                        <p>Please choose a nickname to join the room</p>
                        <form on:submit|preventDefault={handleNicknameSubmit}>
                            <Input bind:value={nickname} placeholder="Your nickname..." />
                            <div class="popup-buttons">
                                <Button type="submit">Join Room</Button>
                                <button type="button" on:click={closePopup}>
                                    <Button variant="outline">Cancel</Button>
                                </button>
                            </div>
                            {#if nicknameError}
                                <div class="nickname-error">{nicknameError}</div>
                            {/if}
                        </form>
                    </div>
                </div>
            {/if}
            {#if showNicknamePopup==false}
            <form class="flex w-full max-w-sm items-center space-x-2">
                <Input type="room-code" bind:value={roomcode} placeholder="enter room-code here..." />
                <Button type="submit">Submit</Button>
            </form>
            {/if}
            
        </div>
        {#if errorMessage}
            <p2 class="mt-2 text-red-500">{errorMessage}</p2>
        {/if}
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

    main h1 {
        font-size: 2.5rem;
        color: #dc4405;
        margin-bottom: 10px;
    }

    main p {
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

    .popup-overlay {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    }

    .popup-content {
        background: white;
        padding: 2rem;
        border-radius: 8px;
        width: 90%;
        max-width: 400px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    }

    .popup-content h2 {
        color: #dc4405;
        margin-bottom: 1rem;
    }

    .popup-content p {
        margin-bottom: 1.5rem;
        color: #666;
    }

    .popup-buttons {
        display: flex;
        gap: 1rem;
        margin-top: 1.5rem;
        justify-content: flex-end;
    }
    
    .separator {
        color: #dc4405;
    }

    .nickname-error {
        color: red;
        margin-top: 1rem;
        font-size: 0.95rem;
        text-align: left;
    }
</style>
