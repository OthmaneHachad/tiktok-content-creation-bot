import assemblyai as aai
import requests
import json
import time
import os
from dotenv import load_dotenv

load_dotenv()  # This will load environment variables from a .env file located in the same directory as your script


def get_subtitle_file(transcript_id, api_token, file_format):
    if file_format not in ["srt", "vtt"]:
        raise ValueError("Invalid file format. Valid formats are 'srt' and 'vtt'.")


    CHARACTER_LIMIT = 15

    url = f"https://api.assemblyai.com/v2/transcript/{transcript_id}/{file_format}?chars_per_caption={CHARACTER_LIMIT}"
    headers = {"authorization": api_token}

    response = requests.get(url, headers=headers)

    if response.status_code == 200:
        return response.text
    else:
        raise RuntimeError(f"Failed to retrieve {file_format.upper()} file: {response.status_code} {response.reason}")


if __name__ == "__main__":

  api_key = str(os.environ["ASSEMBLYAI_KEY"])
  print(f"HERE IS YOUR API KEY {api_key}")
  aai.settings.api_key = api_key


  base_url = "https://api.assemblyai.com/v2"

  headers = {
      "authorization": api_key
  }


  with open("../audio_files/speech.mp3", "rb") as f:
    response = requests.post(base_url + "/upload",
                            headers=headers,
                            data=f)

  print("HERE IS THE RESPONSE FROM ASSEMBLY \n \n")
  print(response.json())
  print("\n \n")
  upload_url = response.json()["upload_url"]

  data = {
      "audio_url": upload_url
  }

  url = base_url + "/transcript"
  response = requests.post(url, json=data, headers=headers)


  transcript_id = response.json()['id']
  polling_endpoint = f"https://api.assemblyai.com/v2/transcript/{transcript_id}"

  while True:
    transcription_result = requests.get(polling_endpoint, headers=headers).json()

    if transcription_result['status'] == 'completed':
      break

    elif transcription_result['status'] == 'error':
      raise RuntimeError(f"Transcription failed: {transcription_result['error']}")

    else:
      time.sleep(1)


  subtitles = get_subtitle_file(transcript_id, api_key, "srt")
  with open("../merging_files/subtitles.srt", 'w') as file:
    file.write(subtitles)