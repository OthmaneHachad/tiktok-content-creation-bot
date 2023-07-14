import speech_recognition as sr
from pydub import AudioSegment

# Convert mp3 file to wav                                                       
audio = AudioSegment.from_mp3("post_audio.mp3")
audio.export("post_audio.wav", format="wav")

r = sr.Recognizer()

# Split the text into chunks of roughly equal length
# Here, I'm just splitting the text into chunks of approximately one sentence each
# For more accurate results, you would want to split based on the actual length of the audio

chunks = ["They’re quarantined before launch to weed out any infectious diseases and they’re health screened",
            "Of course you can never be 100% sure someone won’t have a brain aneurysm or some other bizarre health emergency",
            "It hasn’t happened yet, but it surely will eventually",
            "I’m sure there are contingency plans on paper for such a thing happening, but space travel is still isolated and dangerous",
            "Just like traveling to the south pole or the bottom of the ocean, there’s an inherent risk involved that the people are accepting when they sign up to ride a gigantic missile at 12,000 mph into a vacuum",
            "You are right about the vetting and medical training",
            "And there is a flight surgeon for every expedition who monitor each crew members health",
            "Usually the flight surgeon is located on the ground but sometimes one of the crew members is a trained medical doctor and can perform the role of flight surgeon from space",
            "There are quite a well stacked medical cabinets on the space station with various drugs and devices that might be needed",
            "You also have to remember that the ISS is a flying research laboratory where they often do biological and medical research"]


def divide_chunks(chunks):
    split_chunks = []

    temp_chunk = []
    for chunk in chunks:
        temp_chunk.append(chunk.split(" "))
    chunks = temp_chunk

    for chunk in chunks:
        chunk_len = len(chunk)//12
        if (chunk_len > 0):
            for i in range(chunk_len):
                split_chunks.append(chunk[(i*12):((i+1)*12)])
        chunk_reste = (len(chunk)%12)*(-1) # negative to get the last *chunk_reste*
        split_chunks.append(chunk[chunk_reste:])

    return [" ".join(split_chunks[i]) for i in range(len(split_chunks)-1)]

print(divide_chunks(chunks=chunks))

def create_ass_file(filename, dialogue_entries):
    with open(filename, 'w') as f:
        f.write("[Script Info]\n")
        f.write("ScriptType: v4.00+\n")
        f.write("WrapStyle: 0\n")
        f.write("ScaledBorderAndShadow: yes\n")
        f.write("YCbCr Matrix: TV.601\n")
        f.write("\n")

        f.write("[V4+ Styles]\n")
        f.write("Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding\n")
        f.write("Style: Default,Arial,20,&H00FFFFFF,&H0000FFFF,&H00000000,&H80000000,0,0,0,0,100,100,0,0,1,1,1,2,10,10,10,0\n")
        f.write("\n")

        f.write("[Events]\n")
        f.write("Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text\n")
        for entry in dialogue_entries:
            f.write("Dialogue: " + entry + "\n")



# note that every 5 seconds, the bot can say 12 to 13 words
# need to split each chunk into 12 words
# time for each subtitle will depend on the number of words of a chunk

dialogue_entries = [
    "0,0:00:00.00,0:00:05.00,Default,,0,0,0,,Hello, world!",
    "0,0:00:06.00,0:00:10.00,Default,,0,0,0,,This is a test."
]


#create_ass_file('test.ass', dialogue_entries)

