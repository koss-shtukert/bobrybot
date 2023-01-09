package pkg

import (
	"github.com/koss-shtukert/bobrybot/cmd"
)

var electricityStatus string

func init() {
	electricityStatus = cmd.GetElectricityStatus()
}
