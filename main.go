package main

import (
	//
	//
	//"flag"
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"strings"
	"time"
	//"strings"
	//"time"
	"github.com/go-telegram-bot-api/telegram-bot-api"
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

	//сохранение в файл
	f, err := os.Create(log)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(source)
	if err != nil {
		panic(err)
	}

	newCars := "Есть новые автомобили\nhttps://showroom.hyundai.ru/"
	if strings.Contains(source, "На данный момент все автомобили распроданы") {
		newCars = "Нет новых автомобилей\nhttps://showroom.hyundai.ru/"
	} else {

	}

	bot, _ := tgbotapi.NewBotAPI("2031940210:AAHUIaQJndVEtdBonIGelJisaw4g0lL6UhQ")
	bot.Debug = true

	msg := tgbotapi.NewMessage(int64(-1001651654069), newCars)

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
