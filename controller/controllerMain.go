package controller

import (
	"fmt"
	"irptb/handler"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartService() {
	telegramBotApiTocken := handler.GetEnv("TELEGRAM_BOT_TOCKEN")
	bot, err := tgbotapi.NewBotAPI(telegramBotApiTocken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message != nil {
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "پیام شما دریافت شد")
	// 		bot.Send(msg)
	// 	}
	// }
	for update := range updates {
		if update.Message == nil {
			continue
		}

		userInput := update.Message.Text
		var replyText string

		// Validate the input.
		if handler.PackageSourceValidator(userInput) {
			replyText = "✅ کد رهگیری مرسوله دریافت شد و در حال برسی وضعیت میباشد.\n تا لحضاتی دیگر مشخصات بسته برای شما ارسال میشود"
			checkPackageStatus(userInput)
		} else {
			replyText = "❌ کد رهگیری مرسوله اشتباه است یا به صورت نادرستی وارد شده است. \nدقت شود که کد رهگیری مرسوله ۲۴ رقم میباشد\n"
		}

		// Create a message to send back.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
		_, err = bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}
func checkPackageStatus(orderSourceNumber string) {
	fmt.Println("this is my suer input from bot :", orderSourceNumber)
}
