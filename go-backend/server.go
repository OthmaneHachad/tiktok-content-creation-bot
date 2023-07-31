package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ClientVideoCreationInput struct {
	Link string `form:"link" binding:"required"`
	Gameplay string `form:"link" binding:"required"`
	Voice string `form:"link" binding:"required"`
}

// Declare variables, not constants, for values retrieved from environment variables
var (
	accessKeyID     string
	secretAccessKey string
	region          string
	S3Bucket        = "tiktok-processed-videos"
)


func GeneratePresignedURL(objectKey string) (string, error) {
    // Initialize AWS credentials from environment variables.
	// Make sure to set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and AWS_REGION environment variables before running this program.
	// Load .env file
	err := godotenv.Load() // This will load the .env file from the same directory by default.
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve environment variables after loading the .env file
	accessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	region = os.Getenv("AWS_REGION")

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return "", err
	}
	S3_client := s3.New(session) 

    // Set the expiration time for the URL. For this example, the link will be valid for 15 minutes.
    req, _ := S3_client.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(S3Bucket),
        Key:    aws.String(objectKey),
    })
    urlStr, err := req.Presign(15 * time.Minute) // 15 minutes pre-signed url

    if err != nil {
        return "", err
    }

    return urlStr, nil
}

func UploadProcessedVideoS3(video_local_path string, video_uuid string) (string){

	// Initialize AWS credentials from environment variables.
	// Make sure to set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, and AWS_REGION environment variables before running this program.
	// Load .env file
	err := godotenv.Load() // This will load the .env file from the same directory by default.
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve environment variables after loading the .env file
	accessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	region = os.Getenv("AWS_REGION")

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		return ""
	}
	S3_client := s3.New(session) 

	// Open the file for use
	file, err := os.Open(video_local_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	suffixe := fmt.Sprintf("processedVideos-%s", video_uuid)
	objectKey := fmt.Sprintf("processedVideos/%s.mp4", suffixe)
	// Upload the file to S3
	_, err = S3_client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})

	if err != nil {
		fmt.Println("Error uploading to S3:", err)
		return ""
	}

	fmt.Println("Successfully uploaded file to", S3Bucket)

	return objectKey
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

	subreddit_name, postID, err_retrieve := RetrieveSubredditAndPostId(link)
	if err_retrieve != nil {
		log.Fatal(err_retrieve)
	}
	speeche, _, comments, err := GetComments(subreddit_name, postID, voice)
	if err != nil {
		log.Fatal(err)
	}

	parsed_comments := splitEveryNWords(comments, 3)
	subtitles_path, err := createSubtitlesFile("../merging_files/subtitles.srt", parsed_comments)
	if err != nil {
		log.Fatal(err)
	}

	gameplay_v_a, err := CutVideoAddAudio(fmt.Sprintf("../merging_files/%s", gameplay_path), speeche)
	if err != nil {
		log.Fatal(err)
	}

	video_local_path, err := BurnSubtitles(gameplay_v_a, subtitles_path)

	fmt.Println("Now starting the Upload to the s3 bucket")

	// creating a video UUID
	processed_video_uuid := uuid.New()

	objectKey := UploadProcessedVideoS3(video_local_path, processed_video_uuid.String())
	// https://tiktok-processed-videos.s3.eu-west-3.amazonaws.com/processedVideos/526bd3b3-e073-4e50-8e9f-e945e79d6475.mp4

	// here we generate the pre-signed url
	// to be able to embed the video inside the html
	presigned_url, _ := GeneratePresignedURL(objectKey)

	// video_link : fmt.Sprintf("https://%s.s3.%s.amazonaws.com/processedVideos/%s.mp4", S3Bucket, region, processed_video_uuid),

	context.IndentedJSON(http.StatusOK, gin.H{
		"subreddit_post_link" : subreddit_name,
		"TTS_voice" : voice,
		"gameplay_extracted_from" : gameplay_path,
		"video_link": presigned_url,
		"processed_video_uuid": processed_video_uuid,
	})
}

func CreateVideoHandler(context *gin.Context) {
	var input ClientVideoCreationInput

	if err := context.ShouldBindJSON(&input) ; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createTiktokVideo(context, input.Link, input.Voice, fmt.Sprintf("../merging_files/%s.mp4", input.Gameplay))
}


func main() {
	
	//  https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/


	router := gin.Default()

	// Serve the static Svelte files
	router.Static("/build", "../svelte-app/public/build/")
	router.Static("/global.css", "../svelte-app/public/global.css")
	router.Static("/favicon.png", "../svelte-app/public/favicon.png")
	router.Static("/assets", "../svelte-app/public/assets/")

	router.GET("/", func(c *gin.Context) {
		c.File("../svelte-app/public/index.html")
	})

	router.POST("/createVideo", CreateVideoHandler)
	router.Run("localhost:8080")
}



