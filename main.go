package main

import (
	"fmt"
	"log"
	"net/http"

	social "github.com/kkdai/line-login-sdk-go"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

// LINE Login related configuration
var channelID, channelSecret string

// LINE MessageAPI related configuration
var serverURL string
var botToken, botSecret string
var socialClient *social.Client

func main() {
	var err error
	serverURL = "https://line-robot-1.dorichor.com"
	channelID = "1657518053"                           // os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELID")
	channelSecret = "77f73a0f78a32b69237b989be0125226" // os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")

	// if bot, err = linebot.New(os.Getenv("LINECORP_PLATFORM_CHATBOT_CHANNELSECRET"), os.Getenv("LINECORP_PLATFORM_CHATBOT_CHANNELTOKEN")); err != nil {
	if bot, err = linebot.New("cec2e43307b02d7d39fd4c48b5d81b19", "ZqxAOZv1Lpi9TjPByENgi12ByPQa8+RbBWZx/uh1AB1JtmOYIE7tyzSEDlIlaD4AqM2iOYH8RfcYSYNhRHFpCfay181J6qG060ReCLaDNXrpXSyO+7qGQRheTtc2SFi5HmVI2LDcVx2xykIk5PH4oAdB04t89/1O/w1cDnyilFU="); err != nil {
		log.Println("Bot:", bot, " err:", err)
		return
	}

	if socialClient, err = social.New(channelID, channelSecret); err != nil {
		log.Println("Social SDK:", socialClient, " err:", err)
		return
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//For LINE login
	http.HandleFunc("/", browse)
	http.HandleFunc("/gotoauthOpenIDpage", gotoauthOpenIDpage)
	http.HandleFunc("/gotoauthpage", gotoauthpage)
	http.HandleFunc("/auth", auth)

	//For linked chatbot
	http.HandleFunc("/callback", callbackHandler)

	//provide by Heroku
	port := "3000"
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Start web service !!\n")
	http.ListenAndServe(addr, nil)
}
