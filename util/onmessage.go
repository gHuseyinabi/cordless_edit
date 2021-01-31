package util

import (
	"strings"
	"time"

	"github.com/Bios-Marcel/discordgo"
)

var (
	commandinfo map[string]string
)

func init() {
	commandinfo = make(map[string]string)
	commandinfo["encsendprefix"] = "/enc "
	commandinfo["advsendprefix"] = "/advenc "
	commandinfo["encrecievefix"] = "ENC"
	commandinfo["advrecievefix"] = "ADVENC"
	commandinfo["legacy_token"] = "golang_malclub_encryption_key111"
	commandinfo["statusfix"] = "/status-set "
}

func OnMessageSend(message string, additional map[string]string, session *discordgo.Session) string {
	if strings.HasPrefix(strings.ToLower(message), commandinfo["encsendprefix"]) {
		return commandinfo["encrecievefix"] + EncryptBase64(
			Encrypt([]byte(message[len(commandinfo["encsendprefix"]):]),
				commandinfo["legacy_token"]))
	} else if strings.HasPrefix(
		strings.ToLower(message), commandinfo["advsendprefix"]) {
		return commandinfo["advrecievefix"] + EncryptBase64(
			Encrypt([]byte(message[len(commandinfo["advsendprefix"]):]),
				additional["userid"]))
	} else if strings.HasPrefix(message, commandinfo["statusfix"]) {
		var customStatus discordgo.CustomStatus
		session.UserUpdateStatusCustom(customStatus)

		go func() {
			var statslist []string
			var stats string = message[len(commandinfo["statusfix"]):]
			statslist = strings.Split(stats, "-")

			var i int = 0
			for range time.Tick(time.Second * 1) {
				if i >= len(statslist) {
					i = 0
				}
				i += 1
				customStatus.Text = statslist[i-1]
				session.UserUpdateStatusCustom(customStatus)
			}
		}()
		return "no_send_message"
	}
	return "nil"
}
