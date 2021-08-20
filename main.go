package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func main() {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", id)
}
