package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"os"
)

func main() {
	copyToClipboard := flag.Bool("clip", false, "Copy to clipboard")
	flag.Parse() // add this line

	id, err := uuid.NewUUID()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to generate UUID%s", LineBreak)
	}

	if *copyToClipboard == true {
		err := clipboard.WriteAll(id.String())
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to copy to clipboard%s", LineBreak)
		} else {
			_, _ = fmt.Fprintf(os.Stdout, "Copied to clipboard%s", LineBreak)
		}
		return
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s%s", id, LineBreak)
}
