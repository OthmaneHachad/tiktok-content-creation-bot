import re
import praw
from google.cloud import texttospeech
from pydub import AudioSegment
import subprocess
from random import randint

import get_comments
import create_subt
import edit_video



reddit = praw.Reddit(
    client_id='Zl5cZsjPgTr1GW1U5T64sw',
    client_secret='mnNDGr2voE4jhGQeczLew4hTH1YVLQ',
    user_agent='macos:tiktok-reddit-bot:v1.0 (by /u/Lower-Present2622)'
)

if __name__ == "__main__":
    post_path = 'https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/'

    dialogue_entries = get_comments.get_reddit_comments(post_path)["comments"]
    subtitles = create_subt.create_subtitles_file("subtitles.srt", dialogue_entries)["subtitles_path"]
    
    minute_extract = edit_video.cut_video('minecraft_1.mp4')["output_video_path"]
    gameplay = edit_video.merge_audio_video(minute_extract, "audio_files/full_speech.mp3")["output_video_path"]
    edit_video.burn_subtitles(gameplay, subtitles)

    




