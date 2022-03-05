package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func config() *Config {
	linksRaw := os.Getenv("URLS")
	token := os.Getenv("TOKEN")
	chatId := os.Getenv("CHAT_ID")
	delayRaw := os.Getenv("DELAY")

	if linksRaw == "" {
		log.Fatalln("Не задан список ссылок")
	}

	if token == "" {
		log.Fatalln("Не задан токен telegram")
	}

	if chatId == "" {
		log.Fatalln("Не задан список чатов для отправки сообщений")
	}

	if delayRaw == "" {
		delayRaw = "10"
	}
	delay, _ := strconv.Atoi(delayRaw)

	var links = strings.Split(linksRaw, ",")
	return &Config{
		links:  links,
		token:  token,
		chatId: chatId,
		delay:  delay,
	}
}

type Config struct {
	links  []string
	token  string
	chatId string
	delay  int
}
