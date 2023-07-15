import praw
from google.cloud import texttospeech
import re

reddit = praw.Reddit(
    client_id='Zl5cZsjPgTr1GW1U5T64sw',
    client_secret='mnNDGr2voE4jhGQeczLew4hTH1YVLQ',
    user_agent='macos:tiktok-reddit-bot:v1.0 (by /u/Lower-Present2622)'
)

def get_reddit_comments_text(post):
    submission = reddit.submission(url=post)
    WORD_LIMIT = 150

    # Replace the "MoreComments" objects with actual comments up to a certain level
    submission.comments.replace_more(limit=0)
    full_text = [submission.title]
    compteur = 0

    for top_level_comment in submission.comments:
        comments = top_level_comment.body

        # parsing out typos and bad characters
        comments = re.sub("\n", "", comments)
        comments = re.split(r'\.\s|\.\n|\.', comments)

        for comment in comments:
            compteur += len(comment.split(' '))
            if compteur <= WORD_LIMIT and not(comment in ['', ' ']):
                full_text.append(comment)
            else : break

    return full_text






def comment_to_speach(text_to_convert):

    text_to_convert = " ".join(text_to_convert)

    # Instantiates a client
    client = texttospeech.TextToSpeechClient()

    # Set the text input to be synthesized
    synthesis_input = texttospeech.SynthesisInput(text=text_to_convert)

    # Build the voice request, select the language code and the ssml voice gender
    voice = texttospeech.VoiceSelectionParams(
        language_code="en-US", ssml_gender=texttospeech.SsmlVoiceGender.MALE
    )

    # Select the type of audio file you want returned
    audio_config = texttospeech.AudioConfig(
        audio_encoding=texttospeech.AudioEncoding.MP3
    )

    # Perform the text-to-speech request on the text input with the selected
    # voice parameters and audio file type
    response = client.synthesize_speech(
        input=synthesis_input, voice=voice, audio_config=audio_config
    )

    # The response's audio_content is binary.
    with open("post_audio.mp3", "wb") as out:
        out.write(response.audio_content)
        print('Audio content written to file "post_audio.mp3"')



print(get_reddit_comments_text('https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/'))
