package services

import (
	"fmt"
	"os"

	"github.com/mailgun/mailgun-go"
)

func SendAlert(email, eventTitle, eventDate, weather string) error {
	mg := mailgun.NewMailgun("sandbox695cdb8d3eab477fbda52f81fb1a8dba.mailgun.org", os.Getenv("MAILGUN_API_KEY"))
	sender := "Weather Event Planner <mailgun@sandbox695cdb8d3eab477fbda52f81fb1a8dba.mailgun.org>"
	subject := "Weather Alert for Your Event"
	body := fmt.Sprintf("<strong>%s</strong>: The weather for your event on %s is %s.", eventTitle, eventDate, weather)
	to := email
	msg := mg.NewMessage(sender, subject, body, to)
	_, id, err := mg.Send(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Email sent with ID: %s\n", id)
	return nil
}
