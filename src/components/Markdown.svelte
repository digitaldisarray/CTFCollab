<script>
    import { document } from "./text.js"
    import { marked } from 'marked'
    import { derived } from "svelte/store";
    let html = derived(document, $document => marked($document))
</script>
<style>
    div {
        border:1px solid black;
        width: 50%;
        height: 80vh;
        margin: 10px;
        padding: 5px;
    }
</style>
<div>
    {#await html}
        <p>waiting</p>
    {:then}
        {@html $html}
    {/await}
</div>