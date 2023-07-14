import speech_recognition as sr
from pydub import AudioSegment

# Convert mp3 file to wav                                                       
audio = AudioSegment.from_mp3("post_audio.mp3")
audio.export("post_audio.wav", format="wav")

r = sr.Recognizer()

# Use wav file in speech recognition
with sr.AudioFile('post_audio.wav') as source:
    audio = r.record(source)  # read the entire audio file
    text = r.recognize_google(audio)
    print(text)

# Split the text into chunks of roughly equal length
# Here, I'm just splitting the text into chunks of approximately one sentence each
# For more accurate results, you would want to split based on the actual length of the audio
chunks = text.split('. ')

# Create the SubRip file
subs = pysrt.SubRipFile()
for i, chunk in enumerate(chunks):
    # Create a new SubRip item for each chunk
    item = pysrt.SubRipItem(i, start=((i*5) * 1000), end=((i*5 + 5) * 1000), text=chunk)
    subs.append(item)

# Save the SubRip file
subs.save('subtitles.srt', encoding='utf-8')
