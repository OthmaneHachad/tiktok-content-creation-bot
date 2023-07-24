
# tiktok-content-creation-bot

![example](https://github.com/OthmaneHachad/tiktok-content-creation-bot/assets/75754374/3e35efd3-0700-4e96-bac6-ee494ea2564a)

This project automates the creation of reddit post reading with a background gameplay that everybody loves on Tiktok. 

The service will be free of charge and served on a website.

I first decided to do it in Python but I quickly realised that efficiency and performance would become issue.

With that, I thought I would do my backend in GOLANG so I 
could gain some experience with it and consequently I chose GOLANG as the main language of my project. 

Version 1.0.0 will be a server-based implementation, where all the computation
will be done on the server side. The ultimate goal is to reach Version 2.0.0, a WebAssembly-based (WASM) implementation where all of the computation will be done on the client side using wasm.


## GOAL

The ultimate goal here is to be able to create those small reddit post video in just a few clicks. A user can specify the background video and speech voice or none and let the app choose them randomly.

More technically, the goal is to transport all of the computation (which is the costly part of the application) on the client side, this is why reaching Version 2.0.0 is critical for the service to be free of charge.
## Tech Stack

**Client:** Svelte 

**Server:** GO

## Roadmap
 - Making the video editing client side (WASM)
 - Adding the option to enable Hardware Acceleration

