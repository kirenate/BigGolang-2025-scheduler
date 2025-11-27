package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		panic("Usage: bfd -i {inputPath} -o {outputPath}\n\nThis script only works with mp4 and mov")
	}
	if args[1] != "-i" {
		panic("Usage: bfd -i {inputPath} -o {outputPath}\n\nThis script only works with mp4 and mov")
	}

	inputPath := args[2]
	splittedInput := strings.Split(inputPath, ".")
	if !(len(splittedInput) == 2 && (splittedInput[1] == "mp4" || splittedInput[1] == "mov")) {
		panic("Wrong input path, example: your/input/path/filename.mp4\n\nThis script only works with mp4 and mov")
	}

	outputPath := args[4]
	splittedOutput := strings.Split(outputPath, ".")
	if !(len(splittedOutput) == 2 && (splittedOutput[1] == "mp4" || splittedOutput[1] == "mov")) {
		panic("Wrong output path, example: your/output/path/filename.mp4\n\nThis script only works with mp4 and mov")
	}

	//file := ffmpeg.Input(inputPath)

}

// -i{inputPath}
