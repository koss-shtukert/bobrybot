package pkg

const (
	START_COMMAND                   = "start"
	ELECTRICITY_STATUS_COMMAND      = "Електрика є?"
	NP_STATUS_COMMAND               = "НП працює?"
	POST_STATUS_COMMAND             = "Пошта працює?"
	MAIN_MENU_COMMAND               = "Головне меню"
	SUBSCRIBE_ELECTRICITY_COMMAND   = "Підписатися на оновлення статусу?"
	UNSUBSCRIBE_ELECTRICITY_COMMAND = "Відписатися від оновлення статусу?"
)

func GetElectricityStatus() string {
	return electricityStatus
}

func GetNPStatus() string {
	return IN_DEVELOPMENT_MESSAGE
}

func GetPostStatus() string {
	return IN_DEVELOPMENT_MESSAGE
}

func SubscribeElectricityStatus(chatId int64) string {
	AddElectricitySubscription(chatId, true)

	return ELECTRICITY_SUBSCRIPTION_ON_MESSAGE
}

func UnsubscribeElectricityStatus(chatId int64) string {
	RemoveElectricitySubscription(chatId, false)

	return ELECTRICITY_SUBSCRIPTION_OFF_MESSAGE
}
