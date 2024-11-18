<script>
    import { document } from "./text.js";
    let width, height = $props();
    function onInput(e){
        if(e.key == 'Tab'){
            e.preventDefault();
        }
        document.update(d => d = e.target.value);
    }
</script>
<style>
    textarea {
        display: block;
        width: 50%;
        height: 100%;
        margin: 10px;
        padding: 5px;
    }
</style>
<textarea
    oninput={onInput}>
</textarea>