<script>

    import {onMount} from "svelte" ;

    export let gameplayOptions = ["minecraft", "trackmania", "subway-surfers"] ;
    export let voiceOptions = ["English - UK (Male)", "English - UK (Female)", "English - US (Male)" , "English - US (Female)" , "English - INDIA (Male)" , "English - INDIA (Female)" , "English - AUSTRALIA (Male)" , "English - AUSTRALIA (Female)" ,]

    let gameplayType = '';
    let voiceType = '' ;
    let subredditLink = '' ;

    let isProcessing ;
    let processedData ;


    function handleChangeGameplay(event) {
        gameplayType = event.target.value ;
        console.log("gameplay type: ", gameplayType)

    }
    function handleChangeVoice(event) {
        voiceType = event.target.value ;
        console.log("voice type: ", gameplayType)

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
                gameplay: gameplayType,
                voice: voiceType
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
                <video width="360" height="640" controls>
                    <source src="${processedData.video_link}" type="video/mp4">
                    Your browser does not support the video tag.
                </video> 
                <a href="${processedData.video_link}" download>Download Video</a>
            `;
        });
    }


    // VIDEO HOVERING CONTROLS MUTE
    onMount(() => {
        // VIDEO HOVERING CONTROLS MUTE
        const divElement = document.getElementById('example-videos');
        const videoElements = divElement.querySelectorAll('video');

        videoElements.forEach(video => {
            video.addEventListener('mouseover', function() {
                video.muted = false;
            });

            video.addEventListener('mouseout', function() {
                video.muted = true;
            });
        });
    });
</script>




<style>
    .styled-div {
        background-color: #6C91BF;
        padding: 20px;
        border-radius: 8px;
        width: 300px; 
        margin: 50px auto; 
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 10px;
    }

    .styled-div input, .styled-div select, .styled-div button {
        width: 100%;
        font-size: 16px;
        color: white; 
        outline: none;
    }

    .styled-div input::placeholder {
        color: rgba(255, 255, 255, 0.7); 
    }

    .styled-div input[type="text"] {
        padding: 10px;
        border-radius: 5px;
        border: none;
        background-color: rgba(255, 255, 255, 0.1);
    }

    .styled-div select {
        padding: 10px;
        border-radius: 5px;
        border: none;
        background-color: rgba(255, 255, 255, 0.1);
        color: white;
        appearance: textfield; /* This removes default OS styling of dropdowns */
    }

    .styled-div button {
        padding: 10px;
        border-radius: 5px;
        border: none;
        background-color: white;
        color: #6C91BF;
        cursor: pointer;
        transition: background-color 0.3s;
    }

    .styled-div button:hover {
        background-color: rgba(255, 255, 255, 0.7);
    }

</style>
<div class="styled-div">
    <input id="link" bind:value={subredditLink} type="text" placeholder="Enter subreddit link" />

    <select bind:value={gameplayType} on:change={handleChangeGameplay}>
        {#each gameplayOptions as option}
            <option value={option}>{option}</option>
        {/each}
    </select>

    <select bind:value={voiceType} on:change={handleChangeVoice}>
        {#each voiceOptions as option}
            <option value={option}>{option}</option>
        {/each}
    </select>

    <button on:click={handleSubmit}>Submit</button>
</div>

{#if isProcessing == true}
    <div>
        <h1>Loading...</h1>
    </div>
{:else}
<div id="videoContainer">

</div>
{/if}


