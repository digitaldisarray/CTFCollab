<script lang="ts">
    import CalendarIcon from "lucide-svelte/icons/calendar";
    import {
      CalendarDate,
      DateFormatter,
      type DateValue,
      getLocalTimeZone,
      parseDate,
      today,
    } from "@internationalized/date";
    import { browser } from "$app/environment";
    import { page } from "$app/stores";
    import { cn } from "$lib/utils.js";
  
    import { Button, buttonVariants } from "$lib/components/ui/button/index.js";
    import { Calendar } from "$lib/components/ui/calendar/index.js";
  
    import * as Popover from "$lib/components/ui/popover/index.js";
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
  
    let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();
  
    const form = superForm(data.form, {
      validators: zodClient(formSchema),
    });
  
    const { form: formData, enhance } = form;
  </script>
  
  <form method="POST" use:enhance>
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
            <Form.Label>Challenge Details</Form.Label>
            <Input {...props} bind:value={$formData.challengeDetails} />
          {/snippet}
        </Form.Control>
        <Form.Description>Describe the details of this challenge</Form.Description>
        <Form.FieldErrors />
      </Form.Field>
  
    <!-- Submit Button -->
    <Form.Button>Submit</Form.Button>
  </form>