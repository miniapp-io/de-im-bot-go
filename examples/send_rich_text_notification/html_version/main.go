package main

import (
	"fmt"
	"log"
	"time"

	deimbotapi "github.com/miniapp-io/de-im-bot-go"
)

func main() {
	botToken := "APITOKEN"
	bot, err := deimbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create a beautiful reward withdrawal notification message using HTML formatting
	withdrawalAmount := "$1,250.00"
	rewardType := "Referral Bonus"
	transactionID := "RWD20260205ABC123"
	timestamp := time.Now().Format("January 02, 2006 at 03:04 PM")

	// Build rich text message with HTML formatting
	message := fmt.Sprintf(`
<b>🎉 Reward Withdrawal Successful!</b>
<br>
━━━━━━━━━━━━━━
<br>
<b>🎁 Reward Type:</b> <i>%s</i>
<br>

<b>💰 Amount:</b> <code>%s</code>
<br>
<b>📅 Date &amp; Time:</b>
%s
<br>
<b>🔖 Transaction ID:</b>
<code>%s</code>
<br>
<b>📊 Status:</b> <i>Completed</i>
<br>
━━━━━━━━━━━━━━
<br>
<b>ℹ️ Details:</b>
Your reward withdrawal has been processed successfully! The funds should appear in your account within 1-3 business days.
<br>
<i>Congratulations on your reward! 🎊</i>
<br>
Need help? Contact us at support@example.com
`, rewardType, withdrawalAmount, timestamp, transactionID)

	// Create message config with HTML parse mode
	msg := deimbotapi.NewMessage("CHAT_ID", message)
	msg.ParseMode = deimbotapi.ModeHTML

	// Send the message
	res, err := bot.Request(msg)
	if err != nil {
		log.Fatal("Unable to send rich text notification: ", err)
	}

	log.Printf("✅ Rich text notification sent successfully!\nResult: %s", string(res.Result))
}
