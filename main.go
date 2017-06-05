package main

import (
	"fmt"
	"os"

	"github.com/battleroid/bugusbot/bot"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage of %s: [config]\n", os.Args[0])
		os.Exit(1)
	}

	botConfig := bot.LoadYAML(os.Args[1])
	botConfig.Start()
}
