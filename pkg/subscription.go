package pkg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	SUBSCRIBTION_TYPE_ELECTRICITY = 1
)

var electricitySubscriptions map[int64]bool

func init() {
	electricitySubscriptions = make(map[int64]bool)
}

func IsElectricitySubscribed(chatId int64) bool {
	return electricitySubscriptions[chatId]
}

func AddElectricitySubscription(chatId int64, isSubscribed bool) {
	electricitySubscriptions[chatId] = isSubscribed
}

func RemoveElectricitySubscription(chatId int64, isSubscribed bool) {
	electricitySubscriptions[chatId] = isSubscribed
}

func sendSubscriptionNotice(bot *tgbotapi.BotAPI, subscriptionType int, electricityStatus string) {
	for chatId, isSubscribed := range electricitySubscriptions {
		log.Println(chatId)
		log.Println(isSubscribed)

		if isSubscribed {
			switch subscriptionType {
			case SUBSCRIBTION_TYPE_ELECTRICITY:
				var msg = tgbotapi.NewMessage(chatId, ELECTRICITY_SUBSCRIPTION_STATUS_CHANGED_MESSAGE+electricityStatus)

				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
				}
			}
		}
	}
}
