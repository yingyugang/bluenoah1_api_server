package main

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"io/ioutil"
	"log"
	"net/http"
)

func RegisterPushToken(w http.ResponseWriter, r *http.Request) {
	var token =  r.Header.Get("token")
	var uuid =  r.Header.Get("uuid")
	tokenMap[uuid] = token
	returnNull(w)
}

func PushNotification(w http.ResponseWriter, r *http.Request) {
	var key =  r.Header.Get("send_key")
	var msg =  r.Header.Get("send_msg")
	if key == "yyg"{
		cert, err := certificate.FromP12File("APNsDev.p12", "1fd94d19bbe3c")

		ioutil.ReadFile("")
		if err != nil {
			log.Fatal("Cert Error:", err)
		}
		for k := range tokenMap {
			var v = tokenMap[k]
			notification := &apns2.Notification{}
			notification.DeviceToken = v
			notification.Topic = "com.moba"
			notification.Payload = []byte(msg)
			client := apns2.NewClient(cert).Development()//   Production()
			res, err := client.Push(notification)
			if err != nil {
				log.Fatal("Error:", err)
			}
			fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
		}
	}
}