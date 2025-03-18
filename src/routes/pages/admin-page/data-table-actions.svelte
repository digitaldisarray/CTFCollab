<script lang="ts">
    import Ellipsis from "lucide-svelte/icons/ellipsis";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { ctfData } from './columns';

    let { id }: { id: string } = $props();

    const deleteCTF = async () => {
        // A confirmation prompt to help prevent accidental deletions
        const confirmed = confirm('Are you sure you want to delete this CTF room? This action is irreversible.');
        if (!confirmed) return;

        try {
            const response = await fetch(`http://localhost:1337/ctfs/${id}`, {
                method: 'DELETE',
                credentials: 'include',
            });

            if (response.ok) {
                ctfData.update((ctfs) => ctfs.filter((ctf) => ctf.id !== id));
                alert('CTF room deleted successfully');

            } else {
                console.error('Delete request failed:', response);
                alert('Failed to delete the CTF room.');
            }
        } catch (error) {
            console.error('Error deleting CTF room:', error);
            alert('An error occurred while deleting the CTF room.');
        }
    };
</script>
   
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
       Copy Room Code
      </DropdownMenu.Item>
     </DropdownMenu.Group>
     <DropdownMenu.Separator />
     <DropdownMenu.Item>View Active Members</DropdownMenu.Item>
     <DropdownMenu.Item>View CTF details</DropdownMenu.Item>
     <DropdownMenu.Item onclick={deleteCTF}>
        Delete CTF Room
    </DropdownMenu.Item>
    </DropdownMenu.Content>
   </DropdownMenu.Root>
