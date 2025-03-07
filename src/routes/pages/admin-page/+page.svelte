<script lang="ts">
  import * as Popover from "$lib/components/ui/popover/index.js";
  import { Button } from "$lib/components/ui/button/index.js"; // Add Button component
  import type { PageData } from "./$types.js";
  import SettingsForm from "./settings-form.svelte";
  import DataTable from "./data-table.svelte";
  import { columns, type Challenge, type CTF } from "./columns.js";
  import { ctfData, formatData } from "./columns.js";
  import { getLocalTimeZone, today } from "@internationalized/date";
  import { Calendar } from "$lib/components/ui/calendar/index.js";
  import { goto } from "$app/navigation";
  import { onMount } from 'svelte';

  
  let { data: pageData }: { data: PageData } = $props();

  let value = [today(getLocalTimeZone())];

  let conformData: Challenge[] = $state([]);
  let loading = $state(true);



  const getUsersCTFs = async () => {
    const token = localStorage.getItem("jwtToken");
    if(!token){
      console.error("No token found");
      goto('/pages/signin');
      return;
    }
    
    try {
      const response = await fetch('http://localhost:1337/ctfs/joined', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });
      if (response.ok) {
        let resp = await response.json();
        
        if (Array.isArray(resp) && resp.length > 0) {
          ctfData.set(formatData(resp));
          loading = false;
        } else {
          loading = false;
          conformData = [];
        }

      } else {
        conformData = [];
      }
    }catch (error) {
      console.error("Error occured", error);
    }
  }


  onMount(() => {
    getUsersCTFs();
  })

</script>



<div class="page-container">
  <div class="container">
    <!-- Header -->
    <header>
      <h1>CTF Events</h1>
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
</div>


<style>
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
    color: #333;
    margin-bottom: 10px;
  }


  header p {
    font-size: 1.2rem;
    color: #666;
  }


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
    width: 100%; /* Make button same width as Calendar */
    max-width: 400px; /* Match Calendar's max-width */
  }


  .custom-button:hover {
    background-color: darkorange;
  }
</style>