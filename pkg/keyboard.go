package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetMainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(ELECTRICITY_STATUS_COMMAND),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(NP_STATUS_COMMAND),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(POST_STATUS_COMMAND),
		),
	)
}

func GetElectricityMenuKeyboard(chatId int64) tgbotapi.ReplyKeyboardMarkup {
	var subscribtionElectricityCommand = SUBSCRIBE_ELECTRICITY_COMMAND

	if IsElectricitySubscribed(chatId) {
		subscribtionElectricityCommand = UNSUBSCRIBE_ELECTRICITY_COMMAND
	}

	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(subscribtionElectricityCommand),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(MAIN_MENU_COMMAND),
		),
	)
}
