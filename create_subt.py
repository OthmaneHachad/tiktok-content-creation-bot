import speech_recognition as sr
from pydub import AudioSegment



# Split the text into chunks of roughly equal length
# Here, I'm just splitting the text into chunks of approximately one sentence each
# For more accurate results, you would want to split based on the actual length of the audio

"""
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
"""

def divide_chunks(chunks):
    SPLITTER = 12 # 2.6 mots / second -> 13 mots / 5 seconds

    split_chunks = []

    temp_chunk = []
    for chunk in chunks:
        temp_chunk.append(chunk.split(" "))
    chunks = temp_chunk

    for chunk in chunks:
        chunk_len = len(chunk)//SPLITTER
        if (chunk_len > 0):
            for i in range(chunk_len):
                split_chunks.append(chunk[(i*SPLITTER):((i+1)*SPLITTER)])
        chunk_reste = (len(chunk)%SPLITTER)*(-1) # negative to get the last *chunk_reste*
        split_chunks.append(chunk[chunk_reste:])

    return [" ".join(split_chunks[i]) for i in range(len(split_chunks)-1)], [len(split_chunks[i]) for i in range(len(split_chunks)-1)]

#print(divide_chunks(chunks=chunks))

def create_ass_file(filename, dialogue_entries, dialogue_length):
    """
    [Script Info]
    Title: Example ASS file
    ScriptType: v4.00+
    WrapStyle: 0
    ScaledBorderAndShadow: yes
    YCbCr Matrix: TV.601
    PlayResX: 384
    PlayResY: 288

    [V4+ Styles]
    Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
    Style: Default,Arial,20,&H00FFFFFF,&H0300FFFF,&H00000000,&H02000000,0,0,0,0,100,100,0,0,1,1,1,2,10,10,10,0

    [Events]
    Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
    Dialogue: 0,0:00:01.00,0:00:05.00,Default,,0,0,0,,First line of subtitle text displayed from 1 to 5 seconds.
    Dialogue: 0,0:00:06.00,0:00:10.00,Default,,0,0,0,,Second line of subtitle text displayed from 6 to 10 seconds.
    """

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
        start_time = 0
        for i, entry in enumerate(dialogue_entries):
            sub_time_length = round((dialogue_length[i]/2.6), 2)
            if (start_time + sub_time_length) >= 10:
                TIMESTAMP = f"0:00:{float(start_time)},0:00:{float(start_time + sub_time_length)}"
            else:
                TIMESTAMP = f"0:00:0{float(start_time)},0:00:0{float(start_time + sub_time_length)}"
            f.write(f"Dialogue: {TIMESTAMP},Default,,0,0,0,,{entry}\n")
            start_time += sub_time_length
    print("succesfully created subtitles !")



# note that every 5 seconds, the bot can say 12 to 13 words
# need to split each chunk into 12 words
# time for each subtitle will depend on the number of words of a chunk
