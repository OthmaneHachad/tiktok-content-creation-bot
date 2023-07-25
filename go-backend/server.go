package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type ClientVideoCreationInput struct {
	Link string `form:"link" binding:"required"`
	Voice string `form:"link" binding:"required"`
}



func createTiktokVideo(context *gin.Context, link string, voice string, gameplay_path string) {

	/* 
		Args :
			- string -> gin Context
			- string -> subreddit link provided by the user (mandatory)
			- string -> synthesize voice type provided by the user (optional)
			- string -> The gameplay the user wants the program to extract a minute from (optional) (minecraft, trackmania, subway surfers...)

		Output : 
			- all of the inputs used (because some of them are unknown to the user, since not provided)
	*/

	startTime := time.Now()

	subreddit_name, postID, err_retrieve := RetrieveSubredditAndPostId("https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/")
	if err_retrieve != nil {
		log.Fatal(err_retrieve)
	}
	speeche, _, comments, err := GetComments(subreddit_name, postID, voice)
	if err != nil {
		log.Fatal(err)
	}

	parsed_comments := splitEveryNWords(comments, 3)
	subtitles_path, err := createSubtitlesFile("./merging_files/subtitles.srt", parsed_comments)

	gameplay_v_a, err := CutVideoAddAudio("./merging_files/minecraft_1.mp4", speeche)
	fmt.Println(err)

	fmt.Println(BurnSubtitles(gameplay_v_a, subtitles_path))
	elapsedTime := time.Since(startTime)
	fmt.Printf("The creation of the Tiktok Video took %s", elapsedTime)


	context.IndentedJSON(http.StatusOK, gin.H{
		"subreddit_post_link" : subreddit_name,
		"TTS_voice" : voice,
		"gameplay_extracted_from" : gameplay_path,
	})
}

func CreateVideoHandler(context *gin.Context) {
	var input ClientVideoCreationInput

	if err := context.ShouldBindJSON(&input) ; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createTiktokVideo(context, input.Link, input.Voice, "./merging_files/minecraft_1.mp4")
}


func test() {
	router := gin.Default()

	// Serve the static Svelte files
	router.Static("/static", "./svelte-app/public/build")

	router.GET("/", func(c *gin.Context) {
		c.File("./svelte-app/public/index.html")
	})

	router.POST("/createVideo", CreateVideoHandler)
	router.Run("localhost:8080")
}



