import speech_recognition as sr
from pydub import AudioSegment
import get_comments



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


def create_subtitles_file(filename, dialogue_entries):
    

    with open(filename, 'w') as f:
        start_time = 0
        sub_time_length = 1
        for sentence in dialogue_entries:
            for i, entry in enumerate(sentence):
                TIMESTAMP = f"00:00:{start_time},000 --> 00:00:{start_time + sub_time_length},000"
                f.write(str(i+1) + "\n")
                f.write(TIMESTAMP + "\n")
                f.write(entry + "\n \n")
                start_time += sub_time_length

    return {
        "statusCode" : 200,
        "subtitles_path" : filename
    }