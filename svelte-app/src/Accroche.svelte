<script>
    import { push } from 'svelte-spa-router';

    function goToCreate(){
        push('/create') ;
    }

</script>


<style>
    @import url('https://fonts.googleapis.com/css2?family=Inconsolata:wght@300;900&display=swap');

    .content-container {
        background-color: #6C91BF;
        padding: 20px;
        /*box-shadow: 0px 0px 15px rgba(0, 0, 0, 0.1); */
        max-width: 600px; 
        margin: 0px auto; 
        text-align: center; /* This will center the button and any inline or inline-block content */
    }

    .content-container h2 {
        color: white;
        font-size: 48px; 
        font-family: 'Inconsolata', monospace;
        font-weight: bold;
        margin-bottom: 20px; 
    }

    .content-container p {
        color: white;
        font-size: 24px;
        font-family: 'Inconsolata', monospace;
        line-height: 1.5; 
    }

    .create-page-button {
        cursor: pointer;
        background-color: #001A23;
        color: #ffffff;                 
        border: none;                    
        border-radius: 25px;             
        width: 125px;                     
        height: 50px;     
        font-size:  12px;               
        display: inline-block;           /* Making it inline-block so it respects the text-align of its container */
        text-align: center;              /* Center the text inside the button */
        line-height: 30px;               /* This helps vertically center the text inside the button */
        transition: 0.3s;                /* For any hover effects you might add later */
        text-decoration: none;           /* Remove underline if ever applied due to being in an <a> tag */
    }

</style>


<div class="content-container">

    <h2>NEED TO FIX THE ROUTING + THE SUBTITLES WRITING (DELETE THIS TAG AFTER READING IT)</h2>
    <h2>Create your Tiktok videos in a few clicks</h2>
    <p>Create 1 minute reading video from your favorite subreddits in a matter of seconds.</p>
    <button class="create-page-button" on:click={goToCreate}>CREATE NOW</button>
</div>