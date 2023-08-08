<script>
    // ********* IMPORTS *****************
    import Choices from "./Choices.svelte";
    import GifSlider from "./GifSlider.svelte";
    import Accroche from "./Accroche.svelte";
    import Create from "./Create.svelte";
    
</script>

<style>
    :global(body) {
        margin: 0;
        background-color: #6C91BF;  /* Use your desired color */
    }

    #landing-page {
        background-color: #6C91BF;
    }
    .top-bar {
        display: flex;
        justify-content: space-between; /* Pushes items to the far edges */
        padding: 10px 20px;             /* Add some padding around */
    }

    .site-name {
        font-size: 20px;                /* Optional: Adjust font size */
        font-weight: bold;              /* Optional: Make it bold */
    }



    .redirect-button {
    padding: 10px 15px;
    cursor: pointer;
    background-color: #001A23;
    color: #ffffff;                  /* Text color set to white for better visibility on dark background */
    border: none;                    /* Remove any default borders */
    border-radius: 25px;              /* This makes it circular */
    width: 75px;                     /* Set a fixed width */
    height: 30px;                    /* Set a fixed height */
    display: flex;                   /* Using flex to center the text/icon inside the button */
    justify-content: center;         /* Horizontally center the text/icon */
    align-items: center;             /* Vertically center the text/icon */
    transition: 0.3s;                /* Smooth transition for hover effects */

    /* Optional: Add hover effect */
    &:hover {
            background-color: #00343f;   /* Slightly darker shade for hover */
    }
}
</style>

<div id="landing-page">
    <div class="top-bar">
        <span class="site-name">Create Your Content</span>
        <a href="https://www.tiktok.com" target="_blank">
            <button class="redirect-button">TikTok</button>
        </a>
    </div>
    
    <Accroche/>
</div>



<GifSlider/>
<Choices/>