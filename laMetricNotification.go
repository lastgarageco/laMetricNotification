// Package lametricnotification implements simple functions
// to push notifications on your local network to your
// LaMetric device
package lametricnotification

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
)

// NotificationSound holds the sound for the notification
type NotificationSound struct {
	Category string `json:"category"`
	ID       string `json:"id"`
	Repeat   int    `json:"repeat"`
}

// NotificationFrame holds the frame for the notification
type NotificationFrame struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

// NotificationModel holds the frame for the notification
type NotificationModel struct {
	Cycles int                 `json:"cycles"`
	Frames []NotificationFrame `json:"frames"`
	Sound  NotificationSound   `json:"sound"`
}

// Notification Notification object
// This will hold the JSON that will be sent to the device
type Notification struct {
	Priority string            `json:"priority"`
	IconType string            `json:"icon_type"`
	Model    NotificationModel `json:"model"`
}

// SendSimpleNotification sends a simple notification to your laMetric
func SendSimpleNotification(apikey string, ip string, message string) error {
	if apikey == "" || ip == "" {
		return errors.New("Api Key or Device IP can't be empty")
	}

	noti := Notification{
		Priority: "info",
		IconType: "none",
		Model: NotificationModel{Cycles: 1,
			Frames: []NotificationFrame{
				NotificationFrame{
					Icon: "i1233",
					Text: message},
			},
			Sound: NotificationSound{
				Category: "notifications",
				ID:       "letter_email",
				Repeat:   1},
		},
	}

	err := PushNotification(apikey, ip, noti)
	if err != nil {
		return err
	}
	return nil
}

// PushNotification push the Notification message to the device
func PushNotification(apikey string, ip string, notification Notification) error {
	if apikey == "" || ip == "" {
		return errors.New("Api Key or Device IP can't be empty")
	}

	auth := "dev:" + apikey
	passcode := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	laMetricNotificationURL := "http://" + ip + ":8080/api/v2/device/notifications"
	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", laMetricNotificationURL, bytes.NewReader(notificationJSON))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", passcode)
	req.Header.Add("Cache-Control", "no-cache")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
