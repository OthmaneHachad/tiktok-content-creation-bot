# This file will cut a 1 minute gameplay video 
# and add the audio + subtiles from the 2 other files

import subprocess
from moviepy.editor import VideoFileClip, TextClip, concatenate_videoclips, CompositeVideoClip
from random import randint
from pydub import AudioSegment

def cut_video(video_path):
    # here is the commande template we will use
    #ffmpeg -ss 00:00:00 -i input.mp4 -t 00:01:00 -vf "scale=-1:1080, crop=1080:1920" final_output.mp4
    
    video = VideoFileClip(video_path)
    video_length = int(video.duration)
    video_length_h, remainder = divmod(video_length, 3600)
    video_length_m, video_length_s = divmod(remainder, 60)

    timestamp = randint(5, (video_length-10))
    h, r = divmod(timestamp, 3600)
    m, s = divmod(r, 60)


    COMMAND = f"ffmpeg -ss {h}:{m}:{s} -i {video_path} -t 00:01:00 cropped_gameplay.mp4" # -vf \"scale=1080:1920\"
    try:
        subprocess.run(COMMAND, shell=True, check=True)
        return {
            "output_video_path": "cropped_gameplay.mp4"
        }
    except subprocess.CalledProcessError as e:
        print(f"The subprocess returned a non-zero exit status: {e.returncode}")
    except Exception as e:
        print(f"An error occurred: {str(e)}")






def merge_audio_video(gameplay, speeche):

    try:
        COMMAND = f"ffmpeg -i {gameplay} -i {speeche} -c:v copy -c:a aac gameplay_w_audio_video.mp4"
        subprocess.run(COMMAND, shell=True, check=True)
        return {
            "output_video_path": "gameplay_w_audio_video.mp4"
        }
    except subprocess.CalledProcessError as e:
        print(f"The subprocess returned a non-zero exit status: {e.returncode}")
    except Exception as e:
        print(f"An error occurred: {str(e)}")



def burn_subtitles(video, subtitles_file):
    # "ffmpeg -i input.mp4 -vf ass=subtitles.ass output.mp4"
    COMMAND = f"ffmpeg -i {video} -vf \"subtitles={subtitles_file}:force_style='Alignment=2,MarginV=140'\"  final_tiktok_video.mp4" # -vf \"scale=1080:1920\"
    # -vf \"subtitles={subtitles_file}:force_style='Alignment=2,MarginV=10'\"
    try:
        subprocess.run(COMMAND, shell=True, check=True)
        return {
            "output_video_path": "final_tiktok_video.mp4"
        }
    except subprocess.CalledProcessError as e:
        print(f"The subprocess returned a non-zero exit status: {e.returncode}")
    except Exception as e:
        print(f"An error occurred: {str(e)}")



#minute_extract = cut_video('minecraft_1.mp4')["output_video_path"]
#gameplay = merge_audio_video(minute_extract, "audio_files/full_speech.mp3")["output_video_path"]
#print(gameplay)
#burn_subtitles("gameplay_w_audio_video.mp4", "subtitles.srt")

