<script lang="ts">
    import * as Popover from "$lib/components/ui/popover/index.js";
    import type { PageData } from "./$types.js";
    import SettingsForm from "./settings-form.svelte";
    export let data: PageData;
    
</script>

<div class="about">
    <h1>New Event Placeholder</h1>
    <p>Event options will go here</p>
    <Popover.Root>
      <Popover.Trigger>Open</Popover.Trigger>
      <Popover.Content><SettingsForm {data} /></Popover.Content>
    </Popover.Root>

    <a href="/">Home</a>
  </div>
  
  <style>
    .about {
      text-align: center;
      display: block;
      margin: 20px auto;
    }
    .about h1 {
      font-size: 2.5rem;
      color: #333;
      margin-bottom: 10px;
    }
  </style>

  