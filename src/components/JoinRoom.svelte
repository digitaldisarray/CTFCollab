
<script>
    import { goto } from "$app/navigation";
    import { redirect } from "@sveltejs/kit";

    let roomcode = $state('');

    function roomInput(e){
        roomcode = e.target.value;
        goto('/rooms');
    }
</script>
<style>
    div {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        margin: auto;
        border: 1px solid;
        border-radius: 25px; 
        width: 350px;
        height: 150px;
    }

    button {
        margin: 5px;
        width: 150px;
        height: 25px;
    }
    input {
        border-radius: 25px;
        text-align: center;
        height: 30px;
    }
</style>
<div>
    <input onchange={roomInput} type="text" placeholder="Type room key..." >
    <button>Enter Room</button>
    {#if roomcode === null}
        <p>Invalid roomcode</p>
    {:else if roomcode !== '' }
        <p>roomcode: {roomcode}</p>
    {/if}
</div>