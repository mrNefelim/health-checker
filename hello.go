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

func main() {
	linksRaw, linksExists := os.LookupEnv("URLS")
	token, tokenExists := os.LookupEnv("TOKEN")
	chatId, chatIdExists := os.LookupEnv("CHAT_ID")
	delayRaw, delayExists := os.LookupEnv("DELAY")

	if !delayExists {
		delayRaw = "10"
	}
	delay, _ := strconv.ParseInt(delayRaw, 10, 64)

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

	for {
		for _, link := range strings.Split(linksRaw, ",") {
			var statusCode = makeRequest(strings.TrimSpace(string(link)))
			if statusCode == 0 {
				sendMessage("Сайт "+string(link)+" отдал невалидный ответ", string(chatId), *bot)
			} else if statusCode != 200 {
				sendMessage("Сайт "+string(link)+" отдает код "+strconv.Itoa(statusCode), string(chatId), *bot)
			}
		}
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

type Client struct {
	bot *tgbotapi.BotAPI
}
