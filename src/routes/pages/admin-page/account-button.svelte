<script lang="ts">
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { goto } from "$app/navigation";
    let { cookie } = $props();
    const logout = async(e: Event) => {
        if(confirm("Are you sure you want to logout?")){
            let resp = await fetch("http://localhost:1337/users/logout", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
            })

            if(resp.ok){
                goto("/", {invalidateAll: true})
            }
        }
    }
</script>
<div class="float-right" >

    <DropdownMenu.Root>
        <DropdownMenu.Trigger>
    
            <Button
             variant="ghost"
             size="icon"
             class="size-16 p-0 bg-white text-black"
            >
                Account
            </Button>
       
     
        </DropdownMenu.Trigger>
        <DropdownMenu.Content>
            <DropdownMenu.Item onclick={() => console.log("change password")}>Change Password</DropdownMenu.Item>
            <DropdownMenu.Item onclick={(e: Event) => logout(e)}>Logout</DropdownMenu.Item>
        </DropdownMenu.Content>
    
       
    </DropdownMenu.Root>
</div>