package handlers

import (
	"slices"

	"github.com/mymmrac/telego"
	"github.com/the-technat/telegram-minecraft-manager/pkg/config"
)

// AuthMiddleware checks if the user is authorized to run this command
func AuthMiddleware(bot *telego.Bot, update telego.Update, C *config.Config) bool {

	// first of all if the messanger is an admin, everything is allowed
	if isAdmin(C.Admins, update.Message.From.Username) {
		return true
	}

	// when in doubt, the action is not allowed
	return false
}

func isAdmin(admins []string, user string) bool {
	if slices.Contains(admins, user) {
		return true
	}
	return false
}
