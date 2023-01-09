package pkg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/koss-shtukert/bobrybot/cmd"
	"github.com/robfig/cron/v3"
)

func RunCronJobs(bot *tgbotapi.BotAPI) {
	log.Println("Cron: RunCronJobs")

	cron := cron.New()

	RunElectricityJob(bot, cron)

	cron.Start()
}

func RunElectricityJob(bot *tgbotapi.BotAPI, cron *cron.Cron) {
	log.Println("Cron: RunElectricityJob")

	cron.AddFunc("@every 3m", func() {
		var oldElectricityStatus = electricityStatus

		electricityStatus = cmd.GetElectricityStatus()

		if oldElectricityStatus != electricityStatus {
			sendSubscriptionNotice(bot, SUBSCRIBTION_TYPE_ELECTRICITY, electricityStatus)
		}

		log.Println("Cron: GetElectricityStatus, response: " + electricityStatus)
	})
}
