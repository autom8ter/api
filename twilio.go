package api

import "fmt"

func SearchUSPhoneNumbersURL(account string) string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/AvailablePhoneNumbers/US/Local.json", account)
}

func IncomingPhoneNumbersURL(account string) string {
	return fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/IncomingPhoneNumbers.json", account)
}

func ChatServiceURL() string {
	return "https://chat.twilio.com/v2/Services"
}

