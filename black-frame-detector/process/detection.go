package process

import (
	"github.com/pkg/errors"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
)

func GetVideoMetadata(filename string) ([]byte, error) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return nil, errors.Wrap(err, "pipe")
	}
	os.Stdout = w

	inputKwargs := ffmpeg.KwArgs{}
	inputKwargs["hide_banner"] = ""

	outputKwargs := ffmpeg.KwArgs{}
	outputKwargs["filter:v"] = "blackframe"
	outputKwargs["f"] = "null"

	cmd := ffmpeg.Input(filename, inputKwargs).
		Output("out.null", outputKwargs).
		WithOutput(w).
		ErrorToStdOut()

	err = cmd.Run()

	if err != nil {
		return nil, errors.Wrap(err, "cmd failed")
	}

	err = w.Close()
	if err != nil {
		return nil, errors.Wrap(err, "close w")
	}

	res, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "reading")
	}

	os.Stdout = oldStdout

	return res, nil
}

func ExtractBlackFramesFromMetadata(meta []byte) ([]byte, error) {
}
