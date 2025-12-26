import argparse, pathlib, sys
from collections import defaultdict
parser = argparse.ArgumentParser(description="Simple text file analyzer. You can specify what you want to analyze with following parameters.")
parser.add_argument("filepath", type=pathlib.Path, help="Path to the file with the text to be analyzed")
parser.add_argument("-l", "--lines", action="store_true", help="Amount of lines")
parser.add_argument("-c", "--chars", action="store_true", help="Amount of characters")
parser.add_argument("-b", "--blanks", action="store_true", help="Amount of blank lines")
parser.add_argument("-f", "--freq", action="store_true", help="Frequency dictionary")
args = parser.parse_args()

filepath: pathlib.Path = args.filepath 

if not filepath.exists():
    print(f"Filepath {filepath} does not exist")
    sys.exit(2)
if not filepath.is_file():
    print(f"{filepath} is not a file")
    sys.exit(2)

lines = 0
chars = 0
blanks = 0
freq = defaultdict(lambda: 0)
with open(filepath, "rt", encoding="utf-8") as f:
    for line in f:
        if args.lines: 
            lines += 1
        if not line.strip() and args.blanks:
            blanks += 1 
        for ch in line.strip():
            if args.chars: chars +=1
            if args.freq:
                freq[ch] += 1

if args.lines: print(f"Amount of lines = {lines}")
if args.chars: print(f"Amount of characters = {chars}")
if args.blanks: print(f"Amount of blank lines = {blanks}")
if args.freq:
    print("Frequency dcitionary:")
    for (ch, n) in freq.items():
        print(f"\t{ch}: {n}")
sys.exit(0)