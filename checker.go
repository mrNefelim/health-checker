package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func check() {

}

func makeRequest(url string) int {
	resp, error := http.Get(url)
	if error != nil {
		return 0
	}

	return resp.StatusCode
}

func checkStatus(link string, chatId string, bot Client, delay int) {
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
