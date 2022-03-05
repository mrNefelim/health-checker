package main

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Client struct {
	bot *tgbotapi.BotAPI
}

func sendMessage(text string, chatId string, Bot Client) {
	chatNumId, err := strconv.ParseInt(chatId, 10, 64)

	if err == nil {
		Bot.SendMessage(text, chatNumId)
	}
}

func (c *Client) SendMessage(text string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, error := c.bot.Send(msg)
	return error
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

/*


func newBot(apiKey string) *Client {
	bot, error := tgbotapi.NewBotAPI(apiKey)
	if error != nil {
		log.Panic(error)
	}

	return &Client{
		bot: bot,
	}
}

func (c *Client) sendMessage(text string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, error := c.bot.Send(msg)
	return error
}

func sendMessage(text string, chatId string, Bot Client) {
	chatNumId, err := strconv.ParseInt(chatId, 10, 64)

	if err == nil {
		Bot.SendMessage(text, chatNumId)
	}
}
*/
