

<script>
    import { onMount } from 'svelte'; // Import the onMount lifecycle function

    let currentGifIndex = 0; // This will be our reactive variable

    onMount(() => { // This code will run after the component is added to the DOM
        const gifs = document.querySelectorAll('.slider-gif');
        const nextButton = document.getElementById('nextGif');

        nextButton.addEventListener('click', () => {
            gifs[currentGifIndex].style.display = 'none';
            currentGifIndex = (currentGifIndex + 1) % gifs.length; // Loop back to the start if we go past the last gif
            gifs[currentGifIndex].style.display = 'block';
        });
    });
</script>


<style>

    @import url('https://fonts.googleapis.com/css2?family=Inconsolata:wght@300&display=swap');

    .features-section {
        display: flex;
        justify-content: space-between; 
        background-color: #6C91BF;
        padding-top: 15px;  
        padding-bottom: 15px;  
        padding-left: 50px;  
        padding-right: 50px;  
        font-family: 'Inconsolata', monospace;
        color: white;
    }

    .features-text {
        flex: 6; /* This means the text section will take up 3 parts of the total space */
        padding: 20px;
        font-size: 72px;
        display: flex;
        justify-content: center; /* Centers content horizontally within the container */
        align-items: center;     /* Centers content vertically within the container */
    }

    .features-text p {
        text-align: justify; /* This makes the text stretch to fill the full width, giving a more block-like appearance */
        line-height: 1.5; /* Adjust for readability */
        margin: 0 auto; /* This will center the limited-width paragraph within its container */
    }

    .gif-slider {
        flex: 4; /* This means the gif slider section will also take up 50% of the space of the parent container */
        display: flex;
        flex-direction: column; /* This will stack the images vertically */
        gap: 20px; /* This adds space between the gifs. Adjust as necessary */
        padding: 20px; /* Add more padding if necessary */
        align-items: center; /* This centers the gif images and button horizontally */
        justify-content: center; /* This centers the gif images and button vertically */
    }

    .gif-slider img {
        width: 60%; /* Adjust to desired gif size; smaller values make the gif smaller */
        border-radius: 10px; /* Optional: to make the corners of the gifs rounded */
    }

    /* ... Previous CSS ... */

    .slider-gif {
        display: none;
    }

    .gif-slider img:first-child {
        display: block; 
    }

    #nextGif {
        margin-top: 20px; 
        cursor: pointer;
        background-color: white;
        color: #6C91BF;
        border: none;
        padding: 10px 20px;
        border-radius: 10px;
        font-family: 'Inconsolata', monospace;
    }


    
</style>




<div class="features-section">
    <div class="features-text">
        <p>Select your favorite gameplay and voice (or bring your own!) to create the most unique tiktok readings<p/>
    </div>
    
    <div class="gif-slider">
        <img class="slider-gif" src="/assets/example_1.gif" alt="IMG_1"/>
        <img class="slider-gif" src="/assets/example_2.gif" alt="IMG_2"/>
        <img class="slider-gif" src="/assets/example_1.gif" alt="IMG_3"/>
        <img class="slider-gif" src="/assets/example_1.gif" alt="IMG_4"/>
        <!-- Add more gifs as necessary -->
        <button id="nextGif">Next</button>
    </div>
</div>


