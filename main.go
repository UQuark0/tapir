package main

import (
	"github.com/UQuark0/tapir/telegram"
	"os"
)

func main() {
	configFile, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()
	tapir, err := telegram.NewTapirBot(os.Getenv("TAPIR_BOT_TOKEN"), configFile)
	if err != nil {
		panic(err)
	}
	tapir.Init()
	err = tapir.Run()
	if err != nil {
		panic(err)
	}
}