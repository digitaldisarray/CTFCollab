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

    let { data: pageData }: { data: PageData } = $props();

    let roomcode = '';

    onMount(() => {
      connectWebSocket();
      roomcode = $page.url.searchParams.get('code') || "";
      getCurrentCTF();
      getChallenges();
    });
    
    let manuallyClosed = false;
    onDestroy(() => {
      if (ws) {
        console.log("Cleaning up WebSocket connection.");
        manuallyClosed = true;
        ws.close(1000, "Client component unmounting");
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
    };

    ws.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data);
        console.log("WebSocket message received:", message);

        // Handle different message types from the server
        switch (message.type) {
          case 'chal_added':
             const newChallenge = message.payload as Challenge;
             challenges.update((current) => [...current, newChallenge]);
            break;
          case 'chal_deleted':
            const deletedId = message.payload.id.toString();
            challenges.update((current) => current.filter((ch) => ch.id !== deletedId));
            break;
          case 'chal_updated':
             const updatedChallenge = message.payload as Challenge;
             challenges.update((current) =>
               current.map((challenge) =>
                 challenge.id === updatedChallenge.id ? { ...challenge, ...updatedChallenge } : challenge
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
      if (!manuallyClosed && !event.wasClean) {
        error = "WebSocket connection closed unexpectedly. Trying to reconnect...";
        setTimeout(connectWebSocket, 5000);
      } else {
        error = "WebSocket connection closed.";
      }
    };
  }

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
                        return {
                            name: challenge.challenge_name,
                            active_members: 0, // TODO: need to add a members to backend or have some way to check it
                            status: "pending", // TODO: need to add a status to the backend maybe? or just remove
                            id: challenge.challenge_id.toString(),
                            hedgedoc_url: challenge.hedgedoc_url,
                            description: challenge.challenge_description
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
          <p>Manage your Challenge rooms.</p>
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
        <!-- Header -->
        <header>
          <h1 class="chal-name">Team Members</h1>
          <p>Invite team members to collaborate on challenges.</p>
        </header>
        <Card.Root>
          <Card.Header>
            <Card.Title>Team Members</Card.Title>
            <Card.Description>Invite your team members to collaborate</Card.Description>
          </Card.Header>
          <Card.Content class="grid gap-6">
            <div class="flex items-center justify-between space-x-4">
              <div class="flex items-center space-x-4">
                <Avatar.Root>
                  <Avatar.Image src="/avatars/01.png" alt="Sofia Davis" />
                  <Avatar.Fallback>SD</Avatar.Fallback>
                </Avatar.Root>
                <div>
                  <p class="text-sm font-medium leading-none">Sofia Davis</p>
                  <p class="text-muted-foreground text-sm">m@example.com</p>
                </div>
              </div>
              <Popover.Root>
                <Popover.Trigger>Owner</Popover.Trigger>
                <Popover.Content class="p-0" align="end">
                  <Command.Root>
                    <Command.Input placeholder="Select new role..." />
                    <Command.List>
                      <Command.Empty>No roles found.</Command.Empty>
                      <Command.Group>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Viewer</p>
                          <p class="text-muted-foreground text-sm">
                            Can view and comment.
                          </p>
                        </Command.Item>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Developer</p>
                          <p class="text-muted-foreground text-sm">
                            Can view, comment, and edit.
                          </p>
                        </Command.Item>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Billing</p>
                          <p class="text-muted-foreground text-sm">
                            Can view, comment and manage billing.
                          </p>
                        </Command.Item>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Owner</p>
                          <p class="text-muted-foreground text-sm">
                            Admin-level access to all resources.
                          </p>
                        </Command.Item>
                      </Command.Group>
                    </Command.List>
                  </Command.Root>
                </Popover.Content>
              </Popover.Root>
            </div>
            <div class="flex items-center justify-between space-x-4">
              <div class="flex items-center space-x-4">
                <Avatar.Root>
                  <Avatar.Image src="/avatars/02.png" alt="Jackson Lee" />
                  <Avatar.Fallback>JL</Avatar.Fallback>
                </Avatar.Root>
                <div>
                  <p class="text-sm font-medium leading-none">Jackson Lee</p>
                  <p class="text-muted-foreground text-sm">p@example.com</p>
                </div>
              </div>
              <Popover.Root>
                <Popover.Trigger>Member</Popover.Trigger>
                <Popover.Content class="p-0" align="end">
                  <Command.Root>
                    <Command.Input placeholder="Select new role..." />
                    <Command.List>
                      <Command.Empty>No roles found.</Command.Empty>
                      <Command.Group>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Viewer</p>
                          <p class="text-muted-foreground text-sm">
                            Can view and comment.
                          </p>
                        </Command.Item>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Developer</p>
                          <p class="text-muted-foreground text-sm">
                            Can view, comment, and edit.
                          </p>
                        </Command.Item>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Billing</p>
                          <p class="text-muted-foreground text-sm">
                            Can view, comment and manage billing.
                          </p>
                        </Command.Item>
                        <Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
                          <p>Owner</p>
                          <p class="text-muted-foreground text-sm">
                            Admin-level access to all resources.
                          </p>
                        </Command.Item>
                      </Command.Group>
                    </Command.List>
                  </Command.Root>
                </Popover.Content>
              </Popover.Root>
            </div>
          </Card.Content>
        </Card.Root>
    
      </div>
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
  </style>
