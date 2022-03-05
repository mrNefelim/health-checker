package main

import (
	"sync"
)

func main() {
	var config = config()
	var bot = New(config.token)
	var waitGroup sync.WaitGroup

	for _, link := range config.links {
		waitGroup.Add(1)
		go func(link string, chatId string, bot Client, delay int) {
			defer waitGroup.Done()
			checkStatus(link, chatId, bot, delay)
		}(link, config.chatId, *bot, config.delay)
	}

	waitGroup.Wait()
}
