package stream

import (
	"fmt"
	"io"
	"os"
)

type InputOptions struct {
	Filenames []string
}

func InitializeInputStream(opt InputOptions) (io.ReadCloser, error) {
	if len(opt.Filenames) == 0 {
		return os.Stdin, nil
	}
	if len(opt.Filenames) > 1 {
		return nil, fmt.Errorf("multiple input files are not yet supported")
	}

	filename := opt.Filenames[0]
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %w", filename, err)
	}

	return f, nil

	// TODO: tail the file.
	/*
		if *cat {
			inputStream = f
		} else {
			// Close the handle - we just wanted to verify it was valid
			f.Close()

			cmd := exec.Command("tail", "-n", "-1000", "-F", *filename)
			var err error
			inputStream, err = cmd.StdoutPipe()
			if err != nil {
				errorf("cannot collect tail -F output of file: %v", err)
			}
			err = cmd.Start()
			if err != nil {
				errorf("cannot collect tail -F output of file: %v", err)
			}
		}
	*/
}
