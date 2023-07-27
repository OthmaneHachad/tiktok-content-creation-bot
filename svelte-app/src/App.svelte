<script>
    // ********* IMPORTS *****************
    import Choices from "./Choices.svelte";
    // ***********************************


    
    function eehandleSubmit() {
        // You can make a fetch call here to send the data to your Go backend
        fetch("/createVideo", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                link: subreddit_link,
                voice: TTS_voice
            })
        }).then(response => response.json())
        .then(data => {
            // Handle the response from your Go backend
            console.log(data);
        });
    }
</script>

<Choices/>



<style>
	
</style>