<script lang="ts">
    import DataTable from "./data-table.svelte";
    import { columns, type Challenge } from "./columns.js";
    import { data } from "./columns.js";
    import { page } from '$app/stores'; 
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Input } from "$lib/components/ui/input/index.js"; // Add Input component
    import { Button } from "$lib/components/ui/button/index.js"; // Add Button component
    import type { PageData } from "./$types.js";
    import SettingsForm from "./settings-form.svelte";
    let { data: pageData }: { data: PageData } = $props();

    let roomcode = '';
    let challenges: Challenge[] = $state([]);
    const getChallenges = async () => {

        
        const token = localStorage.getItem("jwtToken");
        if(!token){
            console.error("No token found");
            return;
        }
        
        try {
            const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/challenges`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.ok) {
                let challengeData = await response.json();
                if(Array.isArray(challengeData) && challengeData.length > 0){
                    challenges = challengeData.map((challenge) => {
                        return {
                            name: challenge.challenge_name,
                            active_members: 0,
                            status: "pending",
                            id: challenge.phrase
                        }
                    });
                }
            } else {
                console.error("Failed to fetch challenges");
            }
        } catch (error) {
            console.error("Error occured", error);
        }
    }
    $effect(() => {
        roomcode = $page.url.searchParams.get('code') || "";
        getChallenges();
    });
    
    
  </script>
  
  <div class="page-container">
    <div class="container">
      <div class="horizontal-container">
        <!-- Header -->
        <header>
          <h1>CTF Events</h1>
          <p>Manage your CTF event rooms.</p>
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
      <DataTable data={challenges} {columns} />
    </div>
    <div class="container">
        <!-- Header -->
        <header>
          <h1>Team Members</h1>
          <p>Invite team members to collaborate on challenges.</p>
        </header>
    
      </div>
  </div>
  
  <style>
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
      color: #333;
      margin-bottom: 10px;
    }
  
    header p {
      font-size: 1.2rem;
      color: #666;
    }
  </style>
