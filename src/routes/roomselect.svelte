
<script>
   let roomcode = $state('')
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
    <input onchange={(e)=> {roomcode = e.target.value}} type="text" placeholder="Type room key..." >
    <button >Enter Room</button>
    {#if roomcode === null}
        <p>Invalid roomcode</p>
    {:else if roomcode !== '' }
        <p>roomcode: {roomcode}</p>
    {/if}
</div>