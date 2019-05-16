package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strings"
)

type Message struct {
	Content  string `json:message`
	Title    string `json:title`
	AppToken string `json:token`
	UserKey  string `json:user`
}

func sendMessage(message Message) {
	path := "https://api.pushover.net/1/messages.json"
	client := &http.Client{}

	data := url.Values{}
	data.Set("user", message.UserKey)
	data.Set("token", message.AppToken)
	data.Set("message", message.Content)
	data.Set("title", message.Title)

	request, _ := http.NewRequest("POST", path, strings.NewReader(data.Encode()))
	response, _ := client.Do(request)

	if response.StatusCode == 200 {
		log.Println("Notification sent")
	} else {
		log.Println("An error ocurred")
	}
}

// these variables will be injected at compile time using -ldflags
var pushover_app_token string
var pushover_user_key string

func main() {
	// in case -ldflags is not used the variables will be loaded from environment
	if pushover_app_token == "" {
		pushover_app_token = os.Getenv("PUSHOVER_APP_TOKEN")
	}
	if pushover_user_key == "" {
		pushover_user_key = os.Getenv("PUSHOVER_USER_KEY")
	}

	hostname, _ := os.Hostname()
	username, _ := user.Current()
	message := Message{
		Content:  fmt.Sprintf("New login from %s", username.Username),
		Title:    hostname,
		AppToken: pushover_app_token,
		UserKey:  pushover_user_key,
	}
	sendMessage(message)
}
