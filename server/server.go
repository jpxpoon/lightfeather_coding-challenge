package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/lightfeather_coding-challenge/server/notifications"
	"github.com/lightfeather_coding-challenge/server/supervisors"
)

var db *supervisors.SupervisorDB

func init() {
	db = supervisors.NewDB()
}

type JsonResponse struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func cors(w http.ResponseWriter, req *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	return
}

func getSupervisors(w http.ResponseWriter, req *http.Request) {
	cors(w, req)

	entries := db.GetAll()
	response := JsonResponse{Type: "success", Data: entries}
	json.NewEncoder(w).Encode(response)
}
func saveNotification(w http.ResponseWriter, req *http.Request) {
	cors(w, req)

	var msg string
	var in *notifications.NotificationStruct
	err := json.NewDecoder(req.Body).Decode(&in)
	if err != nil {
		msg = err.Error()
		return
	}

	msg = validateInput(in)
	if msg == "" {
		// ToDo: implemnet save to database
		fmt.Println("ToDo: Save notification data")
	}

	// response
	response := JsonResponse{Type: "success", Data: in, Message: msg}
	json.NewEncoder(w).Encode(response)
}

func validateInput(in *notifications.NotificationStruct) string {
	if in == nil {
		return ""
	}

	var sb strings.Builder
	if in.FirstName == "" {
		sb.WriteString("FirstName cannot be empty, \n")
	}
	if in.LastName == "" {
		sb.WriteString("LastName cannot be empty, \n")
	}
	if in.SupervisorID == "" {
		sb.WriteString("Supervisor cannot be empty, \n")
	}

	return sb.String()
}
