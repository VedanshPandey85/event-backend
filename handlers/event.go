package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"weather-event-planner/models"
	"weather-event-planner/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	collection := utils.GetCollection("events")
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	event.ID = primitive.NewObjectID().Hex()
	_, err := collection.InsertOne(context.TODO(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(event)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	collection := utils.GetCollection("events")
	var events []models.Event
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var event models.Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	json.NewEncoder(w).Encode(events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	collection := utils.GetCollection("events")
	params := mux.Vars(r)
	id := params["id"]
	var event models.Event
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(event)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	collection := utils.GetCollection("events")
	params := mux.Vars(r)
	id := params["id"]

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure the ID from the URL is used, not the one in the request body
	event.ID = id

	// Update the event in the database
	result, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": id}, event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	collection := utils.GetCollection("events")
	params := mux.Vars(r)
	id := params["id"]
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
