<script lang="ts">
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { currentCTF, challenges } from "./columns";
  
    let challengeId: string = "";
    let hedgedocUrl: string = "";
    let roomcode: string = "";
  
    onMount(() => {
      // Get the roomcode / event phrase
      roomcode = $page.url.searchParams.get("code") || "";
  
      // Get the challenge ID
      challengeId = $page.url.searchParams.get("challenge") || "";
  
      // If they exist, fetch the current CTF data & the full challenge list.
      if (roomcode) {
        getCurrentCTF();
        getAllChallenges();
      }
  
      // Once we have the challengeId, fetch the specific challenge data (to get hedgedocUrl).
      if (challengeId && roomcode) {
        fetchChallenge(challengeId, roomcode);
      }
    });
  
    // Fetch the CTF details (event name, etc.)
    const getCurrentCTF = async () => {
      try {
        const res = await fetch(`http://localhost:1337/ctfs/${roomcode}`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
          credentials: "include"
        });
        if (res.ok) {
          const ctfData = await res.json();
          currentCTF.set({
            ctf_id: ctfData.id,
            ctf_author_id: ctfData.author_id,
            ctf_description: ctfData.description,
            ctf_name: ctfData.name,
            start_date: ctfData.start_date,
            end_date: ctfData.end_date,
            phrase: ctfData.phrase
          });
        } else {
          console.error("Failed to fetch the CTF details");
        }
      } catch (error) {
        console.error("Error fetching current CTF", error);
      }
    };
  
    // Fetch all challenges for the event, store them in `challenges`.
    const getAllChallenges = async () => {
      try {
        challenges.set([]);
        const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/challenges`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
          credentials: "include"
        });
        if (response.ok) {
          const challengeData = await response.json();
          if (Array.isArray(challengeData) && challengeData.length > 0) {
            challenges.set(
              challengeData.map((c) => ({
                name: c.challenge_name,
                active_members: 0, // if needed
                status: "pending", // if needed
                id: c.challenge_id.toString(),
                hedgedoc_url: c.hedgedoc_url,
                description: c.challenge_description
              }))
            );
          }
        } else {
          console.error("Failed to fetch challenges");
        }
      } catch (error) {
        console.error("Error fetching challenges", error);
      }
    };
  
    // Fetch the specific challenge and set `hedgedocUrl`.
    const fetchChallenge = async (challengeId: string, phrase: string) => {
      try {
        const response = await fetch(
          `http://localhost:1337/ctfs/${encodeURIComponent(phrase)}/challenge/${challengeId}`,
          {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            credentials: "include"
          }
        );
  
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
  
    // When user selects a different challenge in the dropdown
    function handleChallengeChange(event: Event) {
      const selectedId = (event.currentTarget as HTMLSelectElement).value;
      window.location.href = `/pages/event-room/challenge?code=${roomcode}&challenge=${selectedId}`;
    }
  </script>
  
  <header class="logo-header">
    <div class="absolute left-4 top-4 md:left-8 md:top-7">
      <a href="/" class="logo">
        <span class="ctf">CTF</span>
        <span class="collab">Collab</span>
      </a>
      {#if $currentCTF?.ctf_name}
        <a
          href={`/pages/event-room?code=${roomcode}`}
          class="event-link"
          style="margin-left: 1rem;"
        >
          {$currentCTF.ctf_name}
        </a>
      {/if}
      {#if $challenges && $challenges.length > 0}
        <select
        class="challenge-select"
        bind:value={challengeId}
        on:change={handleChallengeChange}
        >
        {#each $challenges as chal}
            <option value={chal.id}>
            {chal.name}
            </option>
        {/each}
        </select>
      {/if}
    </div>
  </header>
  
  <div class="iframe-container">
    {#if hedgedocUrl}
        <iframe title="Challenge Notes" src={hedgedocUrl} style="width: 100%" class="my-iframe"></iframe>
    {:else}
      <p>Error</p>
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
    .event-link {
      font-weight: bold;
      text-decoration: none;
    }
    .iframe-container {
      padding-top: 25px;
    }
    .my-iframe {
      width: 100%;
      height: 87vh;
    }
    .challenge-select {
    appearance: none;
    background-color: #000000;
    border: 1px solid #000000;
    border-radius: 4px;
    padding: 0.5rem 2rem 0.5rem 0.75rem;
    color: #ffffff;
    font-size: 0.9rem;
    cursor: pointer;
    outline: none;
    position: relative;
    background-image: url("data:image/svg+xml;charset=UTF8,<svg fill='none' height='24' stroke='%23333' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' viewBox='0 0 24 24' width='24' xmlns='http://www.w3.org/2000/svg'><path d='M6 9l6 6 6-6'></path></svg>");
    background-repeat: no-repeat;
    background-position: right 0.75rem center;
    background-size: 1rem 1rem;
    }
    footer {
        margin-top: 40px;
        font-size: 0.9rem;
        color: #777;
    }

    .challenge-select:hover {
    border-color: #000000;
    }

    .challenge-select:focus {
    border-color: #000000;
    }

    .challenge-select:focus-visible {
    outline: none;
    }

  </style>