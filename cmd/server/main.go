package main

import (
	"fmt"
	"job-scraper/messenger/bot"
	"log"
	"net/http"
	"os"
)

func main() {

	// var bot bot.Bot
	fmt.Println("Server Running on PORT")
	http.HandleFunc("/webhook", bot.Handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
