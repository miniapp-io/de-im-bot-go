# Rich Text Notification Example - Reward Withdrawal

This example demonstrates how to send beautiful, formatted reward withdrawal notification messages using the de-im bot API.

## Files

- **`html_version/main.go`** - HTML formatted version (recommended for better styling control)

## Features

Both examples include:
- 🎉 Reward success indicator
- 🎁 Reward type display
- 💰 Formatted amount display
- 📅 Timestamp
- 🔖 Transaction ID in monospace font
- 📊 Status information
- ℹ️ Additional details
- Professional styling with separators
- Celebratory emojis for better visual appeal

## Setup

1. Replace `"APITOKEN"` with your actual bot token
2. Replace `"CHAT_ID"` with the target chat ID

## Run

For HTML version (recommended):
```bash
cd html_version
go run main.go
```


## Formatting Comparison

### HTML Format (`html_version/main.go`)
- Uses `<b>` for bold text
- Uses `<i>` for italic text
- Uses `<code>` for monospace text
- Uses `<br>` for line breaks (insert a newline in the message)
- Better support for special characters
- More flexible styling options
- Special characters need escaping: `&` → `&amp;`, `<` → `&lt;`, `>` → `&gt;`


## Reward Type Examples

You can customize the `rewardType` variable to match different reward scenarios:
- **"Referral Bonus"** - For user referral rewards
- **"Achievement Reward"** - For milestone achievements
- **"Daily Login Bonus"** - For daily check-in rewards
- **"Competition Prize"** - For winning contests
- **"Cashback Reward"** - For transaction cashback
- **"Special Promotion"** - For promotional campaigns

## Customization

You can customize the message by modifying:
- `rewardType` - The type of reward (e.g., "Referral Bonus", "Achievement Reward", "Daily Login Bonus")
- `withdrawalAmount` - The amount to display
- `transactionID` - The transaction identifier
- `timestamp` - The date and time format
- Message text and structure
- Emojis and separators

## Output Example

```
🎉 Reward Withdrawal Successful!

━━━━━━━━━━━━━━━━━━━━
🎁 Reward Type: Referral Bonus

💰 Amount: $1,250.00

📅 Date & Time:
February 05, 2026 at 02:30 PM

🔖 Transaction ID:
RWD20260205ABC123

📊 Status: Completed
━━━━━━━━━━━━━━━━━━━━

ℹ️ Details:
Your reward withdrawal has been processed successfully! The funds should appear in your account within 1-3 business days.

Congratulations on your reward! 🎊

Need help? Contact us at support@example.com
```

## Notes

- Both versions produce visually similar results
- HTML is recommended for production use due to better control over formatting
- Make sure to test the formatting in your actual de-im chat to verify appearance
- You can combine this with inline keyboards for interactive notifications
