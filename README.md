# laMetricNotification
a simple package for Go to send notifications via local network to your LaMetreic 
using only the Api Key provided at developer.lametric.com

You will find the key at developer.lametric.com under your user name / devices

The list of notifications, alarms and icons are at developer site.

Usage:

EXAMPLE 1:

package main

import (
	"log"

	"github.com/laMetricNotification"
)

func main() {
	apikey := "YOUR_DEVICE_API_KEY"
	ip := "192.168.1.05"
  myMessage := "Hello World"
  
	if err := lametricnotification.SendSimpleNotification(apikey, ip, myMessage); err != nil {
		log.Println("Error sending Notification:", err)
	}
}
  
EXAMPLE 2:

package main

import (
	"log"

	"github.com/laMetricNotification"
)

func main() {
	var message2 lametricnotification.Notification
	
  message2.Priority = "info"
	message2.IconType = "none"
	message2.Model.Cycles = 1
	message2.Model.Frames = make([]lametricnotification.NotificationFrame, 1)
	message2.Model.Frames[0].Icon = "i120"
	message2.Model.Frames[0].Text = "Hello World"
	message2.Model.Sound.Category = "notifications"
	message2.Model.Sound.ID = "car"
	message2.Model.Sound.Repeat = 2

	apikey := "YOUR_DEVICE_API_KEY"
	ip := "192.168.1.05"

	if err := lametricnotification.PushNotification(apikey, ip, message2); err != nil {
		log.Println("Error sending Notification:", err)
	}
}





