package models

type Event struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Location    string `json:"location" bson:"location"`
	Date        string `json:"date" bson:"date"` // Format: YYYY-MM-DD
	Time        string `json:"time" bson:"time"` // Format: HH:MM
	Description string `json:"description" bson:"description"`
	Email       string `json:"email" bson:"email"` // User email for alerts
}
