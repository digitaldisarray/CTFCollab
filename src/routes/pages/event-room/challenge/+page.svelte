<script lang="ts">
    import { page } from "$app/stores";
    import { onMount } from "svelte";

    let challengeId: string = "";
    let hedgedocUrl: string = "";

    // Extract challenge ID and CTF phrase from query parameters
    $: challengeId = new URLSearchParams($page.url.search).get("challenge") || "";
    const ctfPhrase = new URLSearchParams($page.url.search).get("code");

    const fetchChallenge = async () => {
        if (!ctfPhrase || !challengeId) {
            console.error("Missing challenge or CTF phrase in URL");
            return;
        }

        try {
            const response = await fetch(`http://localhost:1337/ctfs/${encodeURIComponent(ctfPhrase)}/challenge/${challengeId}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: 'include'
            });

            if (response.ok) {
                const challenge = await response.json();
                hedgedocUrl = challenge.hedgedoc_url;
            } else {
                console.error("Failed to fetch challenge details");
            }
        } catch (error) {
            console.error("Error occurred", error);
        }
    };

    onMount(async() => {
        if (challengeId) {
            fetchChallenge();
        }
    });
</script>
<header class="logo-header">
    <div class="absolute left-4 top-4 md:left-8 md:top-7">
      <a href="/" class="logo">
          <span class="ctf">CTF</span>
          <span class="collab">Collab</span>
      </a>
    </div>
</header>

<div class="iframe-container">
    {#if hedgedocUrl}
        <iframe title="Challenge Notes" src={hedgedocUrl} style="width: 100%; border: 1px solid #ccc" class="my-iframe"></iframe>
    {:else}
        <p>Loading challenge...</p>
    {/if}
</div>

<style>
    .logo-header {
    display: flex;
    align-items: center;
    height: 60px;
    padding: 0 1rem;
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
    .iframe-container {
    padding-top: 45px;
    }

    .my-iframe {
    width: 100%;
    height: 87vh;
    }
</style>
