package notificationService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const fcmURL = "https://fcm.googleapis.com/fcm/send"

type NotificationPayload struct {
	To           string                 `json:"to"`
	Notification map[string]string      `json:"notification"`
	Data         map[string]interface{} `json:"data,omitempty"`
}

func SendNotification(notification NotificationPayload) {
	serverKey := os.Getenv("FCM_SERVER_KEY") // Keep secret in env
	deviceToken := "YOUR_DEVICE_TOKEN_HERE"  // From frontend

	payload := NotificationPayload{
		To: deviceToken,
		Notification: map[string]string{
			"title": "🔥 Go Native FCM",
			"body":  "Sent without third-party libs!",
		},
		Data: map[string]interface{}{
			"customKey": "customValue",
		},
	}

	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "key="+serverKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
