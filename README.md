Backend Repository README.md
markdown
Copy
# Weather-Based Event Planner - Backend

This is the backend repository for the Weather-Based Event Planner application. It provides RESTful APIs for event management, weather data fetching, and weather alert notifications. The backend is built using **Golang** with the **gorilla/mux** router and uses **MongoDB** for data storage. It integrates with **OpenWeatherMap API** for weather data and **Mailgun API** for email notifications.

---

## Table of Contents
1. [Features](#features)
2. [Technologies Used](#technologies-used)
3. [Setup Instructions](#setup-instructions)
4. [API Endpoints](#api-endpoints)
5. [Environment Variables](#environment-variables)
6. [Running the Application](#running-the-application)
7. [Third-Party Integrations](#third-party-integrations)
8. [Deployment](#deployment)

---

## Features
- **Event Management**: Create, update, delete, and view events.
- **Weather Fetching**: Fetch weather forecasts for a specific location and date.
- **Weather Alerts**: Notify users via email if the weather forecast changes for their event.

---

## Technologies Used
- **Backend Framework**: Golang
- **Router**: gorilla/mux
- **Database**: MongoDB
- **Third-Party APIs**:
  - OpenWeatherMap API (for weather data)
  - Mailgun API (for email notifications)
- **Cron Job**: For periodic weather updates and alerts.

---

## Setup Instructions

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/weather-event-planner-backend.git
   cd weather-event-planner-backend
Install Dependencies:
Ensure you have Go installed. Then, install the required Go packages:

bash
Copy
go mod tidy
Set Up MongoDB:

Install MongoDB locally or use a cloud service like MongoDB Atlas.

Create a database named weather_event_planner.

Set Up Environment Variables:
Create a .env file in the root directory and add the following variables:

env
Copy
MONGODB_URI=mongodb://localhost:27017
OPENWEATHERMAP_API_KEY=your_openweathermap_api_key
MAILGUN_API_KEY=your_mailgun_api_key
MAILGUN_DOMAIN=your_mailgun_domain
PORT=8080
Run the Application:
Start the backend server:

bash
Copy
go run main.go
The server will start at http://localhost:8080.

API Endpoints
Event Management
Create Event: POST /events

json
Copy
{
  "title": "Beach Party",
  "location": "Miami, FL",
  "date": "2023-12-25",
  "time": "15:00",
  "description": "Annual beach party with friends"
}
Get All Events: GET /events

Get Event by ID: GET /events/{id}

Update Event: PUT /events/{id}

Delete Event: DELETE /events/{id}

Weather Fetching
Get Weather Forecast: GET /weather?location={location}&date={date}

Environment Variables
Variable	Description
MONGODB_URI	MongoDB connection string.
OPENWEATHERMAP_API_KEY	API key for OpenWeatherMap.
MAILGUN_API_KEY	API key for Mailgun.
MAILGUN_DOMAIN	Domain registered with Mailgun.
PORT	Port for the backend server (default: 8080).
Running the Application
Start the backend server:

bash
Copy
go run main.go
The server will be available at http://localhost:8080.

Third-Party Integrations
OpenWeatherMap API: Used to fetch weather forecasts for specific locations and dates.

Mailgun API: Used to send email notifications to users for weather updates.

Deployment
To deploy the backend:

Use a cloud platform like AWS, Railway, or Heroku.

Set the environment variables in the cloud platform's dashboard.

Deploy the application using the platform's CLI or dashboard.

License
This project is licensed under the MIT License. See the LICENSE file for details.

Copy

---

### Key Changes:
1. Replaced **SendGrid** with **Mailgun** in the **Technologies Used** and **Third-Party Integrations** sections.
2. Updated the **Environment Variables** section to include `MAILGUN_API_KEY` and `MAILGUN_DOMAIN`.
3. Removed any references to **SendGrid** and replaced them with **Mailgun**.

---

### Notes for Presentation:
- Highlight the use of **Mailgun** for email notifications and how it integrates seamlessly with the backend.
- Explain the benefits of using **Mailgun**, such as its simplicity and reliability for sending transactional emails.
- If you have any specific features of **Mailgun** that you utilized (e.g., templates, webhooks), be sure to mention them during your presentation.
