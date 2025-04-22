package notificationService

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"time"
)

type NotificationService struct {
	AccessToken string
}

func NewNotificationService() *NotificationService {
	accessToken, err := getAccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Println("access token:", accessToken)
	return &NotificationService{
		AccessToken: accessToken,
	}
}

func getAccessToken() (string, error) {
	data, err := ioutil.ReadFile("./monitor-614c1-firebase-adminsdk-fbsvc-3921307ae5.json")
	if err != nil {
		return "", err
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		return "", err
	}

	token, err := conf.TokenSource(context.Background()).Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (noti *NotificationService) SendMessage(deviceToken, title, messages string) error {
	url := "https://fcm.googleapis.com/v1/projects/monitor-614c1/messages:send"

	message := map[string]interface{}{
		"message": map[string]interface{}{
			"token": deviceToken,
			"notification": map[string]string{
				"title": title,
				"body":  messages,
			},
		},
	}

	jsonBody, _ := json.Marshal(message)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+noti.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("FCM Error: %s", body)
	}

	fmt.Println("Push sent!")
	return nil
}
