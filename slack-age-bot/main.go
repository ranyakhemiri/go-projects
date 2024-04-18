package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
)

func printCommandEvents(analytisChannel <-chan *slacker.CommandEvent) {
	for event := range analytisChannel {
		fmt.Println("Command: ", event.Command)
		fmt.Println("Parameters: ", event.Parameters)
		fmt.Println("Event: ", event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", godotenv.Load(".env")["SLACK_BOT_TOKEN"])
	os.Setenv("SLACK_APP_TOKEN", godotenv.Load(".env")["SLACK_APP_TOKEN"])
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition){
		Description: "Set your year of birth",
		Example:     "My year of birth is 1990",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("Error: ", err)
				
			}
			age := 2024 - yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		}
	}

	// stop bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
