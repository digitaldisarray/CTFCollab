<script lang="ts">
    import { Button } from "$lib/components/ui/button";
</script>

<style>
    .CTF {
        color: #dc4405;
    }
    h1 {
        font-size: 3.5rem;
        margin-left: 0px;
        padding-top: 30px;
    }
    h2 {
       font-size: 60px; 
    }
    h3 {
        margin-left: 450px;
        font-size: 2.5rem;
    }
    .logo {
        padding-left: 0px;
    }
    .eventRoomContent{
        display: flex;
        margin-top: 80px;
    }
    .textEditor {
        border: solid black;
        margin-left: 100px;
        padding-bottom: 720px;
        padding-right: 640px;
    }
</style>

<div class="logo">
    <Button
        href="/"
        variant="ghost"
    >
    <h1><span class="CTF">CTF</span><span class="Collab">Collab</span></h1>
    </Button> 
</div>

<div class="eventRoomContent">
    <h2 class="textEditor">Editor</h2>   
    <h3>Users</h3>
</div>