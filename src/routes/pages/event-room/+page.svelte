<script lang="ts">
    import DataTable from "./data-table.svelte";
    import { columns, type Challenge } from "./columns.js";
    import * as Avatar from "$lib/components/ui/avatar/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Command from "$lib/components/ui/command/index.js";
    import { page } from '$app/stores'; 
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Input } from "$lib/components/ui/input/index.js"; // Add Input component
    import { Button } from "$lib/components/ui/button/index.js"; // Add Button component
    import type { PageData } from "./$types.js";
    import SettingsForm from "./settings-form.svelte";
    import {currentCTF, challenges} from "./columns.js"
    import { type CTF  } from "../admin-page/columns.js" //
    import { onMount, onDestroy } from "svelte";
    import { writable } from "svelte/store";

    export type Participant = {
      id: number;
      ctf_id: number;
      user_id: { Int32: number, Valid: boolean };
      guest_id: { Int32: number, Valid: boolean };
      username: { String: string, Valid: boolean };
      nickname: { String: string, Valid: boolean };
    };

    export const participants = writable<Participant[]>([]);

    let { data: pageData }: { data: PageData } = $props();

    let roomcode = "";
    let participantKey = "";

    const backendUrl = "http://localhost:1337";

    // Increment tab count
    function incrementTabCount() {
      const alreadyIncremented = sessionStorage.getItem(`${participantKey}-tabIncremented`);
      if (alreadyIncremented) {
        console.log("Tab already incremented — skipping");
        return;
      }

      const currentVal = localStorage.getItem(participantKey);
      const count = parseInt(currentVal || '0', 10);
      console.log(`Incrementing tab count for ${participantKey}. Current before: ${currentVal}, Parsed: ${count}`);
      console.trace();
      localStorage.setItem(participantKey, String(count + 1));
      sessionStorage.setItem(`${participantKey}-tabIncremented`, "true");
    }



    // Decrement tab count
    function decrementTabCount() {
      const count = parseInt(localStorage.getItem(participantKey) || '0', 10);
      const newCount = Math.max(count - 1, 0);
      localStorage.setItem(participantKey, String(newCount));
    }

    let initializedForThisInstance = false;
    let tabIncremented = false;


    onMount(() => {
      roomcode = $page.url.searchParams.get('code') || "";
      participantKey = `ctfcollab-${roomcode}-active-tabs`;

      if (!tabIncremented) {
        incrementTabCount();
        tabIncremented = true;

        const currentTotalTabs = parseInt(localStorage.getItem(participantKey) || '0', 10);
        if (currentTotalTabs === 1) {
          console.log("First tab for this room globally, calling add-participant");
          fetch(`${backendUrl}/ctfs/${roomcode}/add-participant`, {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            credentials: "include"
          });
        }
      }
      const currentVal = localStorage.getItem(participantKey);
      const count = parseInt(currentVal || '0', 10);
      console.log(`Incrementing tab count for ${participantKey}. Current before: ${currentVal}, Parsed: ${count}`);


      const handleUnload = () => {
        decrementTabCount();
        console.log("decrementing tab count");

        const count = parseInt(localStorage.getItem(participantKey) || '0', 10);
        if (count <= 0) {
          const url = `${backendUrl}/ctfs/${roomcode}/remove-participant`;
          const blob = new Blob([JSON.stringify({})], { type: 'application/json' });
          navigator.sendBeacon(url, blob);
        }
      };

      connectWebSocket();
      getCurrentCTF();
      getChallenges();
      getParticipants();

      // navigates away from the page or closes the browser
      window.addEventListener("beforeunload", handleUnload);
      return () => {
        window.removeEventListener("beforeunload", handleUnload);
        initializedForThisInstance = false;
      };
      
    });
    
    let manuallyClosed = false;
    onDestroy(() => {
      if (typeof navigator !== 'undefined' && typeof navigator.sendBeacon === 'function') {
        sessionStorage.removeItem(`${participantKey}-tabIncremented`);
        decrementTabCount();
        console.log("decrementing tab count");

        const count = parseInt(localStorage.getItem(participantKey) || '0', 10);
        if (count <= 0) {
          const url = `${backendUrl}/ctfs/${roomcode}/remove-participant`;
          const blob = new Blob([JSON.stringify({})], { type: 'application/json' });
          navigator.sendBeacon(url, blob);
        }
      }

      if (ws) {
        console.log("Cleaning up WebSocket connection.");
        manuallyClosed = true;
        ws.close();
        ws = null;
      }
    });


  let ws: WebSocket | null = null;
  let error = $state<string | null>(null);

  function connectWebSocket() {
    const wsUrl = "ws://localhost:1337/ws"; // if backend runs elsewhere in the future, change

    console.log(`Attempting to connect WebSocket to: ${wsUrl}`);
    ws = new WebSocket(wsUrl);

    ws.onopen = () => {
      console.log("WebSocket connection established");
      error = null;
      if (ws && roomcode) {
        const joinMessage = {
          type: "join_room",
          payload: {
            room_id: roomcode
          }
        };
        ws.send(JSON.stringify(joinMessage));
        console.log(`Sent join_room message for room: ${roomcode}`);
      } else {
        console.error("WebSocket is not open or roomcode is missing, cannot send join_room.")
      }
    };

    ws.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        console.log("WebSocket message received:", message);
        console.log("Received message type:", JSON.stringify(message.type));

        // Handle different message types from the server
        switch (message.type) {
          case 'chal_added':
            const newChallenge = message.payload as Challenge;
            challenges.update((current) => [...current, newChallenge]);
            break;
          case 'chal_deleted':
            const deletedId = message.payload.id.toString();
            challenges.update((current) => current.filter((ch) => ch.id.toString() !== deletedId));
            break;
          case 'chal_flag_updated':
            const { id, flag } = message.payload;
            challenges.update(current =>
                current.map(ch =>
                    ch.id.toString() === id.toString()
                        ? { ...ch, flag, status: flag ? 'complete' : 'pending' }
                        : ch
                )
            );
            break;
          case 'participants_updated':
            console.log("Participants list changed, re-fetching...");
            getParticipants(); // refetch participants when updated
            break;
          // Add cases for challenge updates if needed
          default:
            console.warn("Received unknown WebSocket message type:", message.type);
        }
      } catch (e) {
        console.error("Failed to parse WebSocket message or update store:", e);
      }
    };

    ws.onerror = (event) => {
      console.error("WebSocket error:", event);
      error = "WebSocket connection error. Trying to reconnect...";
      setTimeout(connectWebSocket, 5000);
    };

    ws.onclose = (event) => {
      console.log("WebSocket connection closed:", event.reason, `Code: ${event.code}`);
      ws = null;
      if (!manuallyClosed && !event.wasClean) {
        error = "WebSocket connection closed unexpectedly. Trying to reconnect...";
        setTimeout(connectWebSocket, 5000);
      } else {
        error = "WebSocket connection closed.";
      }
    };
  }

  // Function to fetch participants
  const getParticipants = async () => {
    try {
      const res = await fetch(`${backendUrl}/ctfs/${roomcode}/participants`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include'
      });
      if (res.ok) {
        const data: Participant[] = await res.json();
        participants.set(data || []); // Set to empty array if null/undefined
        console.log("Fetched Participants:", data);
      } else {
        console.error("Failed to fetch participants:", await res.text());
        participants.set([]);
      }
    } catch (err) {
      console.error("Error fetching participants:", err);
      participants.set([]);
    }
  };

    const getCurrentCTF = async () => {
      try {
        const res = await fetch(`http://localhost:1337/ctfs/${roomcode}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
              },
          credentials: 'include'
        });
        if(res.ok){
          let ctfData = await res.json()
          
          currentCTF.set({
            ctf_id: ctfData.id,
            ctf_author_id: ctfData.author_id,
            ctf_description: ctfData.description,
            ctf_name: ctfData.name,
            start_date: ctfData.start_date,
            end_date: ctfData.end_date,
            phrase:ctfData.phrase
          })
        } 
      } catch(error){
        console.error("Error fetching current CTF")
      }
    }

    const getChallenges = async () => {  
        try {
            challenges.set([]);
            const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/challenges`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                  },
              credentials: 'include'
            });

            if (response.ok) {
                let challengeData = await response.json();
                if(Array.isArray(challengeData) && challengeData.length > 0){
                    challenges.set(challengeData.map((challenge) => {
                      const isComplete = !!challenge.flag;
                        return {
                            name: challenge.challenge_name,
                            active_members: 0, // TODO: need to add a members to backend or have some way to check it
                            status: isComplete ? "complete" : "pending", // TODO: need to add a status to the backend maybe? or just remove
                            id: challenge.challenge_id.toString(),
                            hedgedoc_url: challenge.hedgedoc_url,
                            description: challenge.challenge_description,
                            flag: challenge.flag
                        }
                    })
                  )
                }
            } else {
                console.error("Failed to fetch challenges");
            }
        } catch (error) {
            console.error("Error occured", error);
        }
    }
    
    
  </script>
<header class="logo-header">
  <div class="absolute left-4 top-4 md:left-8 md:top-7">
    <a href="/" class="logo">
        <span class="ctf">CTF</span>
        <span class="collab">Collab</span>
    </a>
  </div>
</header>

  <div class="page-container">
    <div class="container">
      <div class="horizontal-container">
        <!-- Header -->
        <header>
          <h1 class="chal-name">{$currentCTF?.ctf_name}</h1>
          <p>Challenge Editors.</p>
        </header>
  
        <!-- Add Challenge Button -->
        <div class="button-container">
          <Popover.Root>
            <Popover.Trigger>
              <Button class="new-ctf-button">
                Add Challenge
              </Button>
            </Popover.Trigger>
            <Popover.Content><SettingsForm data={pageData} /></Popover.Content>
          </Popover.Root>
        </div>
      </div>
  
      <!-- DataTable -->
      <DataTable data={$challenges} {columns} />
    </div>
    <div class="container">
  <Card.Root>
    <Card.Header>
      <Card.Title>Team Members</Card.Title>
      {#if $participants.length === 0}
        <Card.Description>No members yet. Invite someone!</Card.Description>
      {:else}
        <Card.Description>Currently {$participants.length} member(s) online.</Card.Description>
      {/if}
    </Card.Header>
    <Card.Content class="grid gap-6">
      {#each $participants as participant (participant.id)}
        <div class="flex items-center justify-between space-x-4">
          <div class="flex items-center space-x-4">
            <Avatar.Root>
              <Avatar.Image src="" alt={participant.username.Valid ? participant.username.String : participant.nickname.String} />
              <Avatar.Fallback>
                {(participant.username.Valid
                  ? participant.username.String.substring(0, 2)
                  : participant.nickname.String.substring(0, 2)).toUpperCase()}
              </Avatar.Fallback>
            </Avatar.Root>
            <div>
              <p class="text-sm font-medium leading-none">
                {participant.username.Valid ? participant.username.String : participant.nickname.String}
              </p>
              <p class="text-muted-foreground text-sm">
                {participant.username.Valid ? "User" : "Guest"}
              </p>
            </div>
          </div>
          <span class="text-sm text-muted-foreground">Member</span>
        </div>
      {:else}
          <p class="text-sm text-muted-foreground">Waiting for participants to join...</p>
      {/each}
    </Card.Content>
  </Card.Root>
</div>
  </div>
  <footer>
    <p>&copy; 2025 CTF-Collab. All rights reserved.</p>
  </footer>
  
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

    .chal-name{
        color: #ffffff;
    }

    header h1 {
    color: #ffffff;
  }

    .page-container {
      display: flex;
      justify-content: flex-start; /* Align container to the left */
      width: 100%; /* Ensure the page container takes full width */
    }
  
    .container {
      display: flex;
      flex-direction: column;
      width: 50%; /* Set container width to 50% of the screen */
      padding: 20px;
      box-sizing: border-box; /* Include padding in the width calculation */
    }

    .horizontal-container {
    display: flex; /* Enables Flexbox */
    flex-direction: row; /* Arranges children horizontally (default) */
    align-items: center; /* Vertically centers children */
    justify-content: flex-start; /* Aligns children to the start of the container */
    gap: 16px; /* Adds space between children */
  }
  .button-container {
    margin-left: auto; /* Pushes the button to the far right */
  }
  
    header h1 {
      font-size: 2.5rem;
      color: #666;
      margin-bottom: 10px;
    }
  
    header p {
      font-size: 1.2rem;
      color: #666;
    }

    footer {
        margin-top: 40px;
        font-size: 0.9rem;
        color: #777;
        text-align: center
    }
  </style>
