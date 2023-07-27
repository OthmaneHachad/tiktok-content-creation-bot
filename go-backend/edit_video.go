package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
)

func MergeVideoAudio(gameplay_extract string, speeche string) (string, error) {
	// COMMAND = f"ffmpeg -i {gameplay} -i {speeche} -c:v copy -c:a aac gameplay_w_audio_video.mp4"
	var speeche_path string = fmt.Sprintf("../audio_files/%s", speeche)
	fmt.Println(speeche_path)
	
	cmd := exec.Command("ffmpeg", "-y", "-i", gameplay_extract, "-i", speeche_path, "-c:v", "copy", "-c:a", "aac", "../merging_files/gameplay_w_audio_video.mp4")
	
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		// print the stderr or return it in your error message for more details
		fmt.Println("Error:", stderr.String())
		return "Command execution Failed (MergeVideoAudio)", err
	}

	return "../merging_files/gameplay_w_audio_video.mp4", nil
}

func BurnSubtitles(gameplay_v_a string, subtitles_path string) (string, error) {
	// COMMAND = f"ffmpeg -i {video} -vf \"subtitles={subtitles_file}:force_style='Alignment=2,MarginV=140'\"  final_tiktok_video.mp4"
	var subtitle_argument string = fmt.Sprintf("subtitles=%s:force_style='Alignment=2,MarginV=140'", subtitles_path)
	fmt.Println(subtitle_argument)

	cmd := exec.Command("ffmpeg", "-y", "-i", gameplay_v_a, "-vf", subtitle_argument, "../merging_files/final_tiktok_video.mp4")
	
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		// print the stderr or return it in your error message for more details
		fmt.Println("Error:"+ stderr.String())
		return "Command execution Failed (BurnSubtitles)", err
	}

	return "../merging_files/final_tiktok_video.mp4", nil
}


// CreateTikTokVideo creates the desired TikTok video by combining operations.
func CutVideoAddAudio(videoPath string, audioPath string) (string, error) {
	videoLength := 34 * 60 // video is 34 minutes long
	
	randMinute := rand.Intn(videoLength)

	h, r := randMinute/3600, randMinute%3600
	m, s := r/60, r%60
	timestamp := fmt.Sprintf("%d:%d:%d", h, m, s)

	output := "../merging_files/gameplay_w_video_audio.mp4"
	// Combined ffmpeg command
	// ffmpeg -y -ss [TIMESTAMP] -i [video_path] -i [audio_path] -t 00:01:00 -c:v copy -c:a copy [output_path] 630:1120
	cmd := exec.Command("ffmpeg", "-y", "-ss", timestamp, "-i", videoPath, "-i", audioPath, "-t", "00:01:00", "-vf", "scale=630:1120", "-c:a", "copy", output)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", stderr.String())
		return "Command execution failed", err
	}

	return output, nil
}