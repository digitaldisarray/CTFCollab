<script lang="ts">
    import { page } from "$app/stores";
    import * as Form from "$lib/components/ui/form/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { formSchema, type FormSchema } from "./schema";
    import {
      type SuperValidated,
      type Infer,
      superForm,
    } from "sveltekit-superforms";
    import { z } from "zod";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { onMount } from "svelte";
    import { type CTF } from "../admin-page/columns.js" 
    import  { challenges, currentCTF, type Challenge } from "./columns.js"
  import { get } from "svelte/store";
    let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();
  
    const form = superForm(data.form, {
      validators: zodClient(formSchema),
    });
    
    
    const { form: formData, enhance } = form;
    
    let roomcode = '';

    onMount(() => {
      roomcode = get(page).url.searchParams.get('code') || "";
    });

    const handleSubmit = async (e: Event) => {
      e.preventDefault();
      const name = $formData.challengeName;
      if (!name) {
        return;
      }
      const description = $formData.challengeDetails;
      const flag = $formData.challengeFlag;

      try {
        const response = await fetch(`http://localhost:1337/ctfs/${roomcode}/challenges`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({
            "name": name,
            "description": description,
            "flag": flag,
          }),
        });
        if (response.ok) {
          const challengeData = await response.json();
          
          let newChallenge: Challenge = {
            id: challengeData.id,
            hedgedoc_url: challengeData.hedgedoc_url,
            status: "pending" as "pending" | "complete",
            active_members: 0,
            name: challengeData.name,
            description: challengeData.description
          };

        } else {
          // Read and log error details
          const errorData = await response.json();
          console.error("Failed to submit form:", errorData);
        }
      } catch (error) {
        console.error("Unable to submit challenge", error);
      }
    };


  </script>
  
  <form method="POST" use:enhance onsubmit={handleSubmit}>
    <!-- Challenge Name Field -->
    <Form.Field {form} name="challengeName">
      <Form.Control>
        {#snippet children({ props })}
          <Form.Label>Challenge Name</Form.Label>
          <Input {...props} bind:value={$formData.challengeName} />
        {/snippet}
      </Form.Control>
      <Form.Description>This sets the Challenge display name.</Form.Description>
      <Form.FieldErrors />
    </Form.Field>
    <!-- Event Code Field -->
    <Form.Field {form} name="challengeDetails">
        <Form.Control>
          {#snippet children({ props })}
            <Form.Label>(Optional) Challenge Details</Form.Label>
            <Input {...props} bind:value={$formData.challengeDetails} />
          {/snippet}
        </Form.Control>
        <Form.Description>Describe the details of this challenge</Form.Description>
        <Form.FieldErrors />
    </Form.Field>
  
    <!-- Submit Button -->
    <Form.Button>Submit</Form.Button>
  </form>