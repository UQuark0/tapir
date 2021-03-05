package main

import (
	"github.com/UQuark0/tapir/telegram"
	"os"
)

func main() {
	tapir, err := telegram.NewTapirBot(os.Getenv("TAPIR_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	tapir.Init()
	err = tapir.Run()
	if err != nil {
		panic(err)
	}
}