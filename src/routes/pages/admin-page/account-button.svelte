<script lang="ts">
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Form from "$lib/components/ui/form/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import {
      type SuperValidated,
      type Infer,
      superForm,
    } from "sveltekit-superforms";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { goto } from "$app/navigation";
    import { changePasswordForm, type ChangePasswordForm } from "./schema";

    let { data } = $props();
    let changePasswordData: SuperValidated<Infer<ChangePasswordForm>> = data.changePasswordForm
    let isModal = $state(false);
    
    const form = superForm(changePasswordData, {
      validators: zodClient(changePasswordForm),
    });
    const { form: formData } = form;
    
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

    const toggleModal = () =>{
        isModal = !isModal
    }
    const changePassword = async(e: Event) => {
        let postreq = {
            'old_password': $formData.oldPassword,
            'new_password': $formData.newPassword
        }
        let user = data.user.username

        if($formData.oldPassword !== $formData.newPassword && $formData.oldPassword !== "" && $formData.newPassword !== ""){
            let resp = await fetch(`http://localhost:1337/users/${user}/password`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
                body: JSON.stringify(postreq)
            })
            
            if(resp.ok){
                alert("Password changed.")
                toggleModal()
            } else {
                console.error(resp)
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
            <DropdownMenu.Item onclick={() => toggleModal()}>Change Password</DropdownMenu.Item>
            <DropdownMenu.Item onclick={(e: Event) => logout(e)}>Logout</DropdownMenu.Item>
        </DropdownMenu.Content>
       
    </DropdownMenu.Root>

    {#if isModal}
        <div id="change-password-model" class="fixed inset-0 flex flex-col items-center justify-center bg-black bg-opacity-70">
                <Form.Field {form} name="oldPassword" id="old-password-input-form">
                    <Form.Control>
                        <Form.Label>Old Password</Form.Label>
                        <Input bind:value={$formData.oldPassword} type="password"/>
                    </Form.Control>
                    <Form.Description>Enter your old password.</Form.Description>
                    <Form.FieldErrors />
                </Form.Field>
                <Form.Field {form} name="newPassword" id="new-password-input-form">
                    <Form.Control>
                        <Form.Label>New Password</Form.Label>
                        <Input bind:value={$formData.newPassword} type="password"/>
                    </Form.Control>
                    <Form.Description>Enter your new password.</Form.Description>
                    <Form.FieldErrors />
                </Form.Field>
                <div class="flex-row">
                    <Form.Button type="submit" onclick={(e) => changePassword(e)}>Enter</Form.Button>
                    <Form.Button onclick={() => toggleModal() }>Cancel</Form.Button>
                </div>
        </div>
    {/if}
</div>