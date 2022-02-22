package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Client struct {
	bot *tgbotapi.BotAPI
}

func main() {
	linksRaw, linksExists := os.LookupEnv("URLS")
	token, tokenExists := os.LookupEnv("TOKEN")
	chatId, chatIdExists := os.LookupEnv("CHAT_ID")
	delayRaw, delayExists := os.LookupEnv("DELAY")

	if !delayExists {
		delayRaw = "10"
	}
	delay, _ := strconv.Atoi(delayRaw)

	if !linksExists {
		log.Fatalln("Не задан список ссылок")
	}

	if !tokenExists {
		log.Fatalln("Не задан токен telegram")
	}

	if !chatIdExists {
		log.Fatalln("Не задан список чатов для отправки сообщений")
	}

	var bot = New(token)
	var links = strings.Split(linksRaw, ",")
	for _, link := range links {
		done := make(chan bool)
		async(link, chatId, *bot, delay)
		done <- true
	}
}

func async(link string, chatId string, bot Client, delay int) {
	var oldStatus = 200
	for {
		var statusCode = makeRequest(strings.TrimSpace(link))
		if statusCode == 200 && oldStatus != 200 {
			sendMessage("Сайт "+link+" снова в работе", chatId, bot)
		} else if statusCode == 0 && oldStatus == 200 {
			sendMessage("Сайт "+link+" отдал невалидный ответ", chatId, bot)
		} else if statusCode != 200 && oldStatus == 200 {
			sendMessage("Сайт "+link+" отдает код "+strconv.Itoa(statusCode), chatId, bot)
		}
		oldStatus = statusCode

		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func sendMessage(text string, chatId string, Bot Client) {
	chatNumId, err := strconv.ParseInt(chatId, 10, 64)

	if err == nil {
		Bot.SendMessage(text, chatNumId)
	}
}

func makeRequest(url string) int {
	resp, error := http.Get(url)
	if error != nil {
		return 0
	}

	return resp.StatusCode
}

func New(apiKey string) *Client {
	bot, error := tgbotapi.NewBotAPI(apiKey)
	if error != nil {
		log.Panic(error)
	}

	return &Client{
		bot: bot,
	}
}

func (c *Client) SendMessage(text string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, error := c.bot.Send(msg)
	return error
}
