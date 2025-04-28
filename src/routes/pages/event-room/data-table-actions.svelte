<script lang="ts">
    import Ellipsis from "lucide-svelte/icons/ellipsis";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { page } from "$app/stores";
    import { challenges } from './columns';

    const ctfPhrase = new URLSearchParams($page.url.search).get("code");
    let { id, description }: { id: string; description: string } = $props();
    let showDetails = $state(false)

    const deleteChal = async () => {
        // A confirmation prompt to help prevent accidental deletions
        const confirmed = confirm('Are you sure you want to delete this Challenge? This action is irreversible.');
        if (!confirmed) return;

        try {
            const response = await fetch(`http://localhost:1337/ctfs/${ctfPhrase}/challenges/${id}`, {
                method: 'DELETE',
                credentials: 'include',
            });

            if (response.ok) {
                challenges.update((current) => current.filter((ch) => ch.id !== id));
                alert('Challenge deleted successfully');

            } else {
                console.error('Delete request failed:', response);
                alert('Failed to delete the Challenge.');
            }
        } catch (error) {
            console.error('Error deleting Challenge:', error);
            alert('An error occurred while deleting the Challenge.');
        }
    };

    const viewDetails = async () => {
        navigator.clipboard.writeText(description)
        if(description === "")
            alert("No description for this event")
        else{
            showDetails = !showDetails
            console.log(showDetails)
            
        }
    }
    const closeDetails = async (e) => {
        e.stopPropagation()
        showDetails = false
    }
   </script>

<style>
    .detailsBox{
        z-index: 99;
        position: absolute;
        border: solid 1px;
        right: 10px;
        padding: 3px;
        border-radius: 4px;
    }
    .close{
        position: relative;
        right: 0;
        top: 0;
    }
</style>
    
   <DropdownMenu.Root>
    <DropdownMenu.Trigger>
     {#snippet child({ props })}
      <Button
       {...props}
       variant="ghost"
       size="icon"
       class="relative size-8 p-0"
      >
       <span class="sr-only">Open menu</span>
       <Ellipsis />
      </Button>
     {/snippet}
    </DropdownMenu.Trigger>
    <DropdownMenu.Content>
     <DropdownMenu.Group>
      <DropdownMenu.GroupHeading>Actions</DropdownMenu.GroupHeading>
      <DropdownMenu.Item onclick={() => navigator.clipboard.writeText(id)}>
       Copy Challenge ID
      </DropdownMenu.Item>
     </DropdownMenu.Group>
     <DropdownMenu.Separator />
     <DropdownMenu.Item>View Active Members</DropdownMenu.Item>
     <DropdownMenu.Item onclick={viewDetails}>View Challenge details</DropdownMenu.Item>
     <DropdownMenu.Item onclick={deleteChal}>
        Delete Challenge
    </DropdownMenu.Item>
    </DropdownMenu.Content>
   </DropdownMenu.Root>

   {#if showDetails}
    <div class="detailsBox">
        <h2>
            {description} <button class="close" onclick={closeDetails}>×</button>
        </h2>
    </div>
    {/if}