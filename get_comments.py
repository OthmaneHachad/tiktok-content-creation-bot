import praw
from google.cloud import texttospeech
import pydub
import os
from pydub import AudioSegment

# export GOOGLE_APPLICATION_CREDENTIALS=/Users/othmane/Development/tiktok-reddit-bot/tiktok-reddit-bot-key.json


reddit = praw.Reddit(
    client_id='Zl5cZsjPgTr1GW1U5T64sw',
    client_secret='mnNDGr2voE4jhGQeczLew4hTH1YVLQ',
    user_agent='macos:tiktok-reddit-bot:v1.0 (by /u/Lower-Present2622)'
)

def get_reddit_comments(post):

    VIDEO_TIME_LIMIT = 55
    submission = reddit.submission(url=post)

    # Replace the "MoreComments" objects with actual comments up to a certain level
    submission.comments.replace_more(limit=0)
    list_comments = [submission.title]
    full_text = submission.title

    for i, top_level_comment in enumerate(submission.comments):
        comment = top_level_comment.body
        temp_full_text = full_text + comment
        if (full_text + comment).count(" ") > 131 : # number of spaces + 1
            break
        else :
            full_text += " " + comment
            list_comments.append(comment)
    
    audio_file = "full_speech.mp3"
    full_speech = comment_to_speach(full_text, audio_file)

    with open(f"audio_files/full_speech.mp3", "wb") as out:
        out.write(full_speech.audio_content)

    audio_length = pydub.AudioSegment.from_file(f"audio_files/full_speech.mp3", format="mp3").duration_seconds

    # split sentences every three word
    final_list_comments = [cut_sentence(comment) for comment in list_comments]


    return {
        "url": post,
        "length" : audio_length,
        "audio_file": audio_file,
        "comments" : final_list_comments
    }







def comment_to_speach(text_to_convert, output_name):

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

    response

    return response

    """# The response's audio_content is binary.
    with open(f"audio_files/{output_name}", "wb") as out:
        out.write(response.audio_content)"""




def cut_sentence(sentence):
    words = sentence.split(" ")
    cut_sentences = []
    for i in range(0, len(words), 3):
        cut_sentence = " ".join(words[i : i + 3])
        cut_sentences.append(cut_sentence)
    return cut_sentences

#print(get_reddit_comments('https://www.reddit.com/r/explainlikeimfive/comments/14wytj0/eli5_how_does_nasa_ensure_that_astronauts_going/'))