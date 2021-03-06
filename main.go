package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tebeka/selenium"
	"os"
	"strings"
	"time"
)

const (
	//Set constants separately chromedriver.exe Address and local call port of
	seleniumPath = `D:\Golang\chromedriver.exe`
	port         = 9515
	log          = "log.file"
)

func main() {

	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	//Call browser urlPrefix: test reference: defaulturlprefix =“ http://127.0.0.1 :4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("https://showroom.hyundai.ru/"); err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	source, err := wd.PageSource()
	if err != nil {
		panic(err)
	}

	if !strings.Contains(source, "На данный момент все автомобили распроданы") {
		//save to file
		f, err := os.Create(log)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = f.WriteString(source)
		if err != nil {
			panic(err)
		}

		//send message
		bot, _ := tgbotapi.NewBotAPI("2031940210:AAHUIaQJndVEtdBonIGelJisaw4g0lL6UhQ")
		bot.Debug = true

		msg := tgbotapi.NewMessage(int64(-1001651654069), "Есть новые автомобили\nhttps://showroom.hyundai.ru/")

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
