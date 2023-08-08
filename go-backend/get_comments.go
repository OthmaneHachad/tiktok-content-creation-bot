package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
	"time"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var voiceMaps = map[string][2]string{
	"English - UK (Male)" : {"en-GB","en-GB-News-J"},
	"English - UK (Female)" : {"en-GB","en-GB-Neural2-F"},
	"English - US (Male)" : {"en-US", "en-US-Wavenet-D"},
	"English - US (Female)" : {"en-US", "en-US-Wavenet-E"},
	"English - INDIA (Male)" : {"en-IN", "en-IN-Wavenet-B"},
	"English - INDIA (Female)" : {"en-IN", "en-IN-Wavenet-D"},
	"English - AUSTRALIA (Male)" : {"en-AU", "en-AU-Neural2-B"},
	"English - AUSTRALIA (Female)" : {"en-AU", "en-AU-Neural2-C"},

}


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

	client, _ := reddit.NewReadonlyClient()
	comments, _, err := client.Post.Get(context.Background(), id)
	if err != nil {
		return "", "", nil, err
	}

	var commentBodies []string
	commentBodies = []string{
		comments.Post.Title,
	}

	var full_text string 
	var comment_text string
	comment_text = comments.Post.Title
	
	for _, comment := range comments.Comments {

		sentences := SplitCommentIntoSentences(comment.Body)

		for _, sentence := range sentences {

			// Trim spaces and check if the sentence is not empty
			trimmedSentence := strings.TrimSpace(sentence)
			if trimmedSentence == "" {
				continue
			}

			if !(strings.Count((comment_text + sentence), " ") > 160) {
				commentBodie2Add := sentence + "."
				commentBodies = append(commentBodies, commentBodie2Add)
				comment_text = comment_text + " " + sentence + "."
			} else {
				break
			}
		}
	}
	

	fmt.Println("starting the parsing here...")

	// ARRANGE THIS, WAS POINTLESS
	full_text, commentBodies = cleanUpData(comment_text, commentBodies)
	
	fmt.Println("\n \n" + full_text)
	fmt.Println("starting the voice synthesize here...")
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
					LanguageCode: voiceMaps[voice][0],
					Name: voiceMaps[voice][1],
			},
			// Select the type of audio file you want returned.
			AudioConfig: &texttospeechpb.AudioConfig{
					AudioEncoding: texttospeechpb.AudioEncoding_MP3,
					SpeakingRate: 1.00, // Default is 1.0
			},
	}

	response, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
			log.Fatal(err)
	}

	return response 


}

func SplitCommentIntoSentences(comment string) ([]string) {
	// here we input in parameter Comment.Body
	comment_sentences := strings.Split(comment, ".")
	return comment_sentences

}

func splitEveryNWords(s []string, n int) []string {
	var chunks []string

	for _, sentence := range s {
		words := strings.Split(sentence, " ")
		var currentChunk []string
		for _, word := range words {
			// Add the word to the current chunk
			currentChunk = append(currentChunk, word)

			// If the word contains a special character or if the chunk size reaches n, append to chunks
			if strings.ContainsAny(word, ":;,.?!") || len(currentChunk) == n {
				chunks = append(chunks, strings.Join(currentChunk, " "))
				currentChunk = nil // Reset the current chunk
			}
		}
		// If there's any remaining words in the current chunk, append it to chunks
		if len(currentChunk) > 0 {
			chunks = append(chunks, strings.Join(currentChunk, " "))
		}
	}

	return chunks
}


func cleanUpData(text string, commentBodies []string) (string, []string) {
    unwantedPairs := map[rune]rune{
        '[': ']',
        '{': '}',
        '(': ')',
        '\\': '\\',
        '*': '*',
        // Add any other pairs of unwanted characters here
    }

    // Clean the main text (string) first
    cleanedText := removeEnclosedText(text, unwantedPairs)

    // Now, clean each individual comment in the slice
    cleanedComments := make([]string, len(commentBodies))
    for i, comment := range commentBodies {
        cleanedComments[i] = removeEnclosedText(comment, unwantedPairs)
    }

    return cleanedText, cleanedComments
}

// Helper function to remove enclosed text
func removeEnclosedText(text string, unwantedPairs map[rune]rune) string {
    for startChar, endChar := range unwantedPairs {
        for {
            startIndex := strings.IndexRune(text, startChar)
            endIndex := strings.IndexRune(text, endChar)
            if startIndex == -1 || endIndex == -1 || endIndex < startIndex {
                break
            }
            text = text[:startIndex] + text[endIndex+1:]
        }
    }
    return text
}


