package main


import (
	"context"
	"fmt"
	"net/url"
	"log"
	"io/ioutil"
	"time"
	"strings"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
        "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)



func RetrieveSubredditAndPostId(link string) (string, string ,error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return "", "", err
	}

	pathParts := strings.Split(parsedURL.Path, "/")
	if len(pathParts) < 5 {
		return "", "", fmt.Errorf("Invalid reddit post link")
	}
	/*  /r/{subreddit_name}/comments/{post_id}/{slug}/
		[0] "" | [1] "r" | [2] "{subreddit_name}" | [3] "comments"
		[4] "{post_id}" | [5] "{slug}" | [6] ""
	*/
	return pathParts[2], pathParts[4], nil
}

func GetComments(subreddit_name string, id string, voice string) (string, string, []string, error){

	client, err := reddit.NewReadonlyClient()
	if err != nil {
		return "", "", nil, err
	}

	comments, _, err := client.Post.Get(context.Background(), id)
	var commentBodies []string
	commentBodies = []string{
		comments.Post.Title,
	}

	var full_text string 
	full_text = comments.Post.Title
	for _, comment := range comments.Comments {
		if !(strings.Count((full_text + comment.Body), " ") > 160) {
			commentBodie2Add := comment.Body
			commentBodies = append(commentBodies, commentBodie2Add)
			full_text = full_text + comment.Body
		} else {
			break
		}
	}

	// synthesize the speech and write it to audio file
	speechStartTime := time.Now()
	speech := SynthesizeSpeech(full_text, voice)
	speechElapsedTime := time.Since(speechStartTime)
	fmt.Printf("The Speech Synthesize took %s", speechElapsedTime)

	// The resp's AudioContent is binary.
	filename := "speech.mp3"
	var file_path string = fmt.Sprintf("../audio_files/%s", filename)
	outcome := ioutil.WriteFile(file_path, speech.AudioContent, 0644)

	if outcome != nil {
			log.Fatal(outcome)
	}
	fmt.Printf("Audio content written to file: %v\n", file_path)

	return file_path, full_text, commentBodies, nil

}

func SynthesizeSpeech(text string, voice string) (*texttospeechpb.SynthesizeSpeechResponse){
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
			log.Fatal(err)
	}
	defer client.Close() // defer means that the action will be done right before return

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
			// Set the text input to be synthesized.
			Input: &texttospeechpb.SynthesisInput{
					InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
			},
			Voice: &texttospeechpb.VoiceSelectionParams{
					LanguageCode: "en-US",
					SsmlGender:   texttospeechpb.SsmlVoiceGender_MALE,
			},
			// Select the type of audio file you want returned.
			AudioConfig: &texttospeechpb.AudioConfig{
					AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			},
	}

	response, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
			log.Fatal(err)
	}

	return response 


}

func splitEveryNWords(s []string, n int) ([]string) {
	var chunks []string
	
	for _, sentence := range s {
		words := strings.Split(sentence, " ")
		var compteur int = 0  
		for _,word := range words {
			fmt.Println(word)
			if compteur+3 >= len(words){
				joined := strings.Join(words[compteur:], " ")
				//fmt.Println(joined)
				chunks = append(chunks, joined)
				break
			} else {
				joined := strings.Join(words[compteur:compteur+3], " ")
				chunks = append(chunks, joined)
				compteur = compteur + 3
			}
		}
	}
	

	return chunks
}


func main() {
	startTime := time.Now()


	subreddit, postID, err_retrieve := RetrieveSubredditAndPostId("https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/")
	if err_retrieve != nil {
		log.Fatal(err_retrieve)
	}
	speeche, _, comments, err := GetComments(subreddit, postID, "en-US")
	if err != nil {
		log.Fatal(err)
	}

	parsed_comments := splitEveryNWords(comments, 3)
	subtitles_path, err := createSubtitlesFile("../merging_files/subtitles.srt", parsed_comments)
	if err != nil {
		log.Fatal(err)
	}

	gameplay_v_a, err := CutVideoAddAudio("../merging_files/minecraft_1.mp4", speeche)
	if err != nil {
		log.Fatal(err)
		fmt.Println("(CutVideoAddAudio)")
	}

	fmt.Println(BurnSubtitles(gameplay_v_a, subtitles_path))
	elapsedTime := time.Since(startTime)
	fmt.Printf("The creation of the Tiktok Video took %s", elapsedTime)
}