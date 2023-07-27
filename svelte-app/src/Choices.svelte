<script>

    export let gameplayOptions = ["minecraft", "trackmania"] ;

    let gameplayType = '';
    let subredditLink = '' ;

    let isProcessing ;
    let processedData ;


    function handleChangeGameplay(event) {
        gameplayType = event.target.value ;
        console.log("gameplay type: ", gameplayType)

    }

    function handleSubmit() {

        isProcessing = true ;

        // You can make a fetch call here to send the data to your Go backend
        fetch("/createVideo", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                link: subredditLink,
                gameplay: gameplayType
            })
        }).then(response => response.json())
        .then(data => {
            // Handle the response from your Go backend
            console.log(data);
            isProcessing = false ;
            processedData = data ;
            /* 
            When you set isProcessing = false; it does not immediately remove the loading message.
            Hence, when your code tries to access videoContainer, it's still not present in the DOM.
            */
        })
    }

    // This reactive statement ($: {}) will trigger when processedData actually changes
    $: if(processedData) {
        // Wait for the DOM to be updated before accessing videoContainer
        setTimeout(() => {
            const videoContainer = document.getElementById('videoContainer');
            videoContainer.innerHTML = `
                <video width="320" height="240" controls>
                    <source src="${processedData.video_link}" type="video/mp4">
                    Your browser does not support the video tag.
                </video> 
                <a href="${processedData.video_link}" download>Download Video</a>
            `;
        });
    }
</script>

<h1>Client Video Input Details</h1>
<div>
    <label for="subredditLink">Link:</label>
    <input id="link" bind:value={subredditLink} type="text" placeholder="Enter subreddit link" />

    <button on:click={handleSubmit}>Submit</button>

    <select bind:value={gameplayType} on:change={handleChangeGameplay}>
        {#each gameplayOptions as option}
            <option value={option}>{option}</option>
        {/each}
    </select>
</div>
{#if isProcessing == true}
    <div>
        <h1>Loading...</h1>
    </div>
{:else}
<div id="videoContainer">

</div>
{/if}


