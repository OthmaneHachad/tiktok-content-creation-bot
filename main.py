import praw
from google.cloud import texttospeech
import speech_recognition as sr
from pydub import AudioSegment
import get_comments
import create_subt
import re


reddit = praw.Reddit(
    client_id='Zl5cZsjPgTr1GW1U5T64sw',
    client_secret='mnNDGr2voE4jhGQeczLew4hTH1YVLQ',
    user_agent='macos:tiktok-reddit-bot:v1.0 (by /u/Lower-Present2622)'
)

if __name__ == "__main__":
    # 'https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/'
    text_from_reddit = get_comments.get_reddit_comments_text('https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/')
    post_audio = get_comments.comment_to_speach(text_from_reddit)

    #print(text_from_reddit)
    #print(" \n ==================== \n")

    # Convert mp3 file to wav                                                       
    audio = AudioSegment.from_mp3("post_audio.mp3")
    audio.export("post_audio.wav", format="wav")
    r = sr.Recognizer()

    chunks, chunks_length = create_subt.divide_chunks(text_from_reddit)

    #print(chunks)

    create_subt.create_ass_file("subtitles.ass", chunks, chunks_length)


