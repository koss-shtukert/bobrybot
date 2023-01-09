package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	WELLCOME_MESSAGE                                = "Привіт! Що бажаєш дізнатися?"
	UNKNOW_MESSAGE                                  = "Я не знаю цієї команди :("
	IN_DEVELOPMENT_MESSAGE                          = "В розробці"
	ELECTRICITY_SUBSCRIPTION_ON_MESSAGE             = "Ви успішно підписалися на оновлення статусу електроенергії"
	ELECTRICITY_SUBSCRIPTION_OFF_MESSAGE            = "Ви успішно відписалися від оновлення статусу електроенергії"
	ELECTRICITY_SUBSCRIPTION_STATUS_CHANGED_MESSAGE = "Статус електроенергії було змінено на - "
)

func BuildReplyMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ReplyMarkup = GetMainMenuKeyboard()

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case START_COMMAND:
			msg.Text = WELLCOME_MESSAGE
		default:
			msg.Text = UNKNOW_MESSAGE
		}
	} else {
		switch update.Message.Text {
		case MAIN_MENU_COMMAND:
			msg.Text = WELLCOME_MESSAGE
		case ELECTRICITY_STATUS_COMMAND:
			msg.Text = GetElectricityStatus()
			msg.ReplyMarkup = GetElectricityMenuKeyboard(update.Message.Chat.ID)
		case NP_STATUS_COMMAND:
			msg.Text = GetNPStatus()
		case POST_STATUS_COMMAND:
			msg.Text = GetPostStatus()
		case SUBSCRIBE_ELECTRICITY_COMMAND:
			msg.Text = SubscribeElectricityStatus(update.Message.Chat.ID)
			msg.ReplyMarkup = GetElectricityMenuKeyboard(update.Message.Chat.ID)
		case UNSUBSCRIBE_ELECTRICITY_COMMAND:
			msg.Text = UnsubscribeElectricityStatus(update.Message.Chat.ID)
			msg.ReplyMarkup = GetElectricityMenuKeyboard(update.Message.Chat.ID)
		default:
			msg.Text = UNKNOW_MESSAGE
		}
	}

	return msg
}
