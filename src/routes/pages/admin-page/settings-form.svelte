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
    import { writable } from "svelte/store";
    import { Button, buttonVariants } from "$lib/components/ui/button/index.js";
    import { Calendar } from "$lib/components/ui/calendar/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import * as Form from "$lib/components/ui/form/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { formSchema, type FormSchema } from "./schema";
    import { ctfData, formatData, type Challenge } from "./columns.js";
    import {
      type SuperValidated,
      type Infer,
      superForm,
    } from "sveltekit-superforms";
    import { z } from "zod";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { id } from "date-fns/locale";
  
    let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();
  
    const form = superForm(data.form, {
      validators: zodClient(formSchema),
    });
    let startOpen = writable(false);
    let endOpen = writable(false);

    const { form: formData, enhance } = form;
  
    // Initialize the date formatter for displaying the selected date
    const df = new DateFormatter("en-US", {
      dateStyle: "long",
    });
    let start_date = $state<DateValue | null>(null);
    let end_date = $state<DateValue | null>(null);
    // Initialize the value for the calendar
    $effect(() => {
      start_date = $formData.start_date ? parseDate($formData.start_date) : null;
      end_date = $formData.end_date ? parseDate($formData.end_date) : null;
    });
    const formatDate = (dateString: string) => {
      if (!dateString) return null;
      return new Date(dateString).toISOString(); 
    };

    const handleSubmit = async (e: Event) => {
      e.preventDefault();
      const token = localStorage.getItem("jwtToken");
        if(!token){
            console.error("No token found");
            return;
      }

      const name = $formData.ctfName;
      const start_date = formatDate($formData.start_date)
      const end_date = formatDate($formData.end_date)
      const description = $formData.ctfDescription
      const jsonBody = JSON.stringify({
          "name": name,
          "start_date": start_date,
          "end_date": end_date,
          "description": description
      })
      
      const response = await fetch('http://localhost:1337/ctfs', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: jsonBody
      });
      if (response.ok) {
        const data = await response.json();
        
        let newCTF: Challenge = {
          id: data,
          members: 0,
          status: "pending",
          date: start_date!!.toString(),
          name: name
        };
        console.log(newCTF)
        ctfData.update((ctf)=>[...ctf, newCTF])
        
      } else {
        console.error("Failed to submit form");
      }

    };

  </script>
  <form method="POST" use:enhance onsubmit={handleSubmit}>
    <!-- Event Name Field -->
    <Form.Field {form} name="ctfName">
      <Form.Control>
        {#snippet children({ props })}
          <Form.Label>Event Name</Form.Label>
          <Input {...props} bind:value={$formData.ctfName} />
        {/snippet}
      </Form.Control>
      <Form.Description>This is your event's public display name.</Form.Description>
      <Form.FieldErrors />
    </Form.Field>
  
    <!-- Event Date Field -->
    <Form.Field {form} name="start_date">
      <Form.Control>
        {#snippet children({ props })}
          <Form.Label>Event Date</Form.Label>
          <Popover.Root bind:open={$startOpen}>
            <Popover.Trigger
              {...props}
              class={cn(
                buttonVariants({ variant: "outline" }),
                "w-[255px] justify-start pl-4 text-left font-normal",
                !start_date && "text-muted-foreground"
              )}
            > 
              {start_date
                ? df.format(start_date.toDate(getLocalTimeZone()))
                : "Pick a date"}
              <CalendarIcon class="ml-auto size-4 opacity-50" />
            </Popover.Trigger>
            <Popover.Content class="w-auto p-2" side="top">
              <Calendar
                type="single"
                value={start_date as DateValue}
                minValue={today(getLocalTimeZone())}
                maxValue={new CalendarDate(2100, 1, 1)}
                calendarLabel="Event Date"
                onValueChange={(v) => {
                  if (v) {
                    $formData.start_date = v.toString();
                    start_date = v;
                    startOpen.set(false); 
                  } else {
                    $formData.start_date = "";
                    start_date = null;
                  }
                }}
              />
            </Popover.Content>
          </Popover.Root>
          <Form.Description>The day your CTF event is scheduled to occur</Form.Description>
          <Form.FieldErrors />
          <input hidden value={$formData.start_date} name={props.name} />
        {/snippet}
      </Form.Control>
    </Form.Field>


    <Form.Field {form} name="end_date">
      <Form.Control>
        {#snippet children({ props })}
          <Form.Label>Event Date</Form.Label>
          <Popover.Root bind:open={$endOpen}>
            <Popover.Trigger
              {...props}
              class={cn(
                buttonVariants({ variant: "outline" }),
                "w-[255px] justify-start pl-4 text-left font-normal",
                !end_date && "text-muted-foreground"
              )}
            > 
              {end_date
                ? df.format(end_date.toDate(getLocalTimeZone()))
                : "Pick a date"}
              <CalendarIcon class="ml-auto size-4 opacity-50" />
            </Popover.Trigger>
            <Popover.Content class="w-auto p-2" side="top">
              <Calendar
                type="single"
                value={end_date as DateValue}
                minValue={today(getLocalTimeZone())}
                maxValue={new CalendarDate(2100, 1, 1)}
                calendarLabel="Event Date"
                onValueChange={(v) => {
                  if (v) {
                    $formData.end_date = v.toString();
                    end_date = v;
                    endOpen.set(false); 
                  } else {
                    $formData.end_date = "";
                    end_date = null;
                  }
                }}
              />
            </Popover.Content>
          </Popover.Root>
          <Form.Description>The day your CTF event is scheduled to occur</Form.Description>
          <Form.FieldErrors />
          <input hidden value={$formData.end_date} name={props.name} />
        {/snippet}
      </Form.Control>
    </Form.Field>
    <!-- Event Code Field -->
    <Form.Field {form} name="ctfDescription">
        <Form.Control>
          {#snippet children({ props })}
            <Form.Label>Description</Form.Label>
            <Input {...props} bind:value={$formData.ctfDescription} />
          {/snippet}
        </Form.Control>
        <Form.Description>(Optional) Add a description.</Form.Description>
        <Form.FieldErrors />
      </Form.Field>
  
    <!-- Submit Button -->
    <Form.Button>Submit</Form.Button>
  </form>