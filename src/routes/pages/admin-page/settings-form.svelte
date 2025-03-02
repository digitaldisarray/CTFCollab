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
  
    // Initialize the date formatter for displaying the selected date
    const df = new DateFormatter("en-US", {
      dateStyle: "long",
    });
  
    // Initialize the value for the calendar
    let value: DateValue | null = $formData.dob
      ? parseDate($formData.dob)
      : null;
  </script>
  
  <form method="POST" use:enhance>
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
    <Form.Field {form} name="dob">
      <Form.Control>
        {#snippet children({ props })}
          <Form.Label>Event Date</Form.Label>
          <Popover.Root>
            <Popover.Trigger
              {...props}
              class={cn(
                buttonVariants({ variant: "outline" }),
                "w-[255px] justify-start pl-4 text-left font-normal",
                !value && "text-muted-foreground"
              )}
            >
              {value
                ? df.format(value.toDate(getLocalTimeZone()))
                : "Pick a date"}
              <CalendarIcon class="ml-auto size-4 opacity-50" />
            </Popover.Trigger>
            <Popover.Content class="w-auto p-2" side="top">
              <Calendar
                type="single"
                value={value as DateValue}
                minValue={today(getLocalTimeZone())}
                maxValue={new CalendarDate(2100, 1, 1)}
                calendarLabel="Event Date"
                onValueChange={(v) => {
                  if (v) {
                    $formData.dob = v.toString();
                    value = v;
                  } else {
                    $formData.dob = "";
                    value = null;
                  }
                }}
              />
            </Popover.Content>
          </Popover.Root>
          <Form.Description>The day your CTF event is scheduled to occur</Form.Description>
          <Form.FieldErrors />
          <input hidden value={$formData.dob} name={props.name} />
        {/snippet}
      </Form.Control>
    </Form.Field>
    <!-- Event Code Field -->
    <Form.Field {form} name="ctfCode">
        <Form.Control>
          {#snippet children({ props })}
            <Form.Label>Access Code</Form.Label>
            <Input {...props} bind:value={$formData.ctfCode} />
          {/snippet}
        </Form.Control>
        <Form.Description>This is your event's access code.</Form.Description>
        <Form.FieldErrors />
      </Form.Field>
  
    <!-- Submit Button -->
    <Form.Button>Submit</Form.Button>
  </form>