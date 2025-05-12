<script lang="ts">
  import * as Popover from "$lib/components/ui/popover/index.js";
  import { Button } from "$lib/components/ui/button/index.js"; // Add Button component
  import type { PageData } from "./$types.js";
  import SettingsForm from "./settings-form.svelte";
  import DataTable from "./data-table.svelte";
  import AccountButton from "./account-button.svelte";
  import { columns, type Challenge, type CTF } from "./columns.js";
  import { ctfData, formatData } from "./columns.js";
  import { getLocalTimeZone, today } from "@internationalized/date";
  import { Calendar } from "$lib/components/ui/calendar/index.js";
  import { goto } from "$app/navigation";
  import { onDestroy, onMount } from 'svelte';
  import { changePasswordForm } from "./schema.js";

  
  let { data: pageData }: { data: PageData } = $props();
 
  // svelte-ignore non_reactive_update
  let value = [today(getLocalTimeZone())];

  let conformData: Challenge[] = $state([]);

  let ws: WebSocket | null = null;
  let error = $state<string | null>(null);
  let loading = $state(true);


  function connectWebSocket() {
    const wsUrl = "ws://localhost:1337/ws"; // if backend runs elsewhere in the future, change

    console.log(`Attempting to connect WebSocket to: ${wsUrl}`);
    ws = new WebSocket(wsUrl);

    ws.onopen = () => {
      console.log("WebSocket connection established");
      error = null;
    };

    ws.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        console.log("WebSocket message received:", message);

        // Handle different message types from the server
        switch (message.type) {
          case 'ctf_added':
             const newChallenge = message.payload as Challenge;
             ctfData.update((current) => [...current, newChallenge]);
            break;
          case 'ctf_deleted':
            const deletedId = message.payload.id as string;
            ctfData.update((current) => current.filter((ctf) => ctf.id !== deletedId));
            break;
          case 'ctf_updated':
             const updatedChallenge = message.payload as Challenge;
             ctfData.update((current) =>
               current.map((ctf) =>
                 ctf.id === updatedChallenge.id ? { ...ctf, ...updatedChallenge } : ctf
               )
             );
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
      if (!event.wasClean) {
        error = "WebSocket connection closed unexpectedly. Trying to reconnect...";
        setTimeout(connectWebSocket, 5000);
      } else {
          error = "WebSocket connection closed.";
      }
    };
  }

  // Fetch initial data
  const getUsersCTFs = async () => {
    loading = true;
    error = null;
    try {
      const response = await fetch('http://localhost:1337/ctfs/joined', { 
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include'
      });
      if (response.ok) {
        let resp = await response.json();
        if (Array.isArray(resp)) {
           const formatted = formatData(resp as CTF[]);
           ctfData.set(formatted);
        } else {
          ctfData.set([]);
        }
      } else {
        console.error("Failed to fetch initial CTFs:", response.statusText);
        error = `Failed to load CTFs: ${response.statusText}`;
        ctfData.set([]);
      }
    } catch (err) {
      console.error("Error fetching initial CTFs:", err);
      error = `Error loading CTFs: ${err instanceof Error ? err.message : String(err)}`;
      ctfData.set([]);
    } finally {
       loading = false;
    }
  };

  onMount(() => {
    getUsersCTFs();
    connectWebSocket(); // Establish WebSocket connection
  });

  onDestroy(() => {
    if (ws) {
      console.log("Closing WebSocket connection on component destroy");
      ws.close(1000, "Client component unmounting"); // code 1000 is normal closure
      ws = null;
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

<div class="page-container">
  <div class="container">
    <!-- Header -->
    <header>
      <h1 class="admin-header">Administrator Dashboard</h1>
      <p>Manage your CTF event rooms.</p>
    </header>

    <!-- DataTable -->
    <DataTable data={$ctfData} {columns} />
  </div>
  <div class="container">
    <!-- Header -->
    <header>
      <h1>Calendar</h1>
      <p>View upcoming CTF events.</p>
      
    </header>

    <div class="cal-container">
      <Popover.Root>
        <Popover.Trigger>
          <Button class="new-ctf-button">
            New CTF
          </Button>
        </Popover.Trigger>
        <Popover.Content><SettingsForm data={pageData} /></Popover.Content>
      </Popover.Root>
      <br />
      <div class="calendar-wrapper">
        <Calendar bind:value={value} type="multiple" class="rounded-md border" />
      </div>
    </div>
    
  </div>
  <AccountButton data={pageData}/>

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

    .admin-header {
        color: #ffffff;
    }

  .page-container {
    display: flex;
    justify-content: flex-start;
    width: 100%;
  }
  


  .container {
    display: flex;
    flex-direction: column;
    width: 50%; /* Each container takes 50% of the screen */
    padding: 20px;
    box-sizing: border-box;
  }


  .cal-container {
    display: flex;
    flex-direction: column;
    align-items: flex-start; /* Align items to the left */
    width: 100%; /* Take full width of its container */
    padding: 20px;
    box-sizing: border-box;
  }


  /* Constrain the Calendar and button width */
  .calendar-wrapper {
    width: 100%; /* Adjust as needed */
    max-width: 400px; /* Set a maximum width for the Calendar */
  }


  header h1 {
    font-size: 2.5rem;
    color: #ffffff;
    margin-bottom: 10px;
  }


  header p {
    font-size: 1.2rem;
    color: #666;
  }

/* 
  .custom-button {
    margin-top: 20px;
    padding: 15px 30px;
    background-color: orange;
    color: black;
    font-size: 1.2rem;
    font-weight: bold;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.3s ease;
    width: 100%;
    max-width: 400px;
  }


  .custom-button:hover {
    background-color: darkorange;
  }
   */
</style>