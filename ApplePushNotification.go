package main

import (
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
)

func main() {
	cert, err := certificate.FromP12File("/Users/yingyugang/Documents/Keys/APNsDev.p12", "1fd94d19bbe3c")
	if err != nil {
		log.Fatal("Cert Error:", err)
	}
	notification := &apns2.Notification{}
	notification.DeviceToken = "10c8dc74e78441098c1fd94d19bbe3c9b8c05e41928bd08d6bdfb64f18be3a52"
	notification.Topic = "com.moba"
	notification.Payload = []byte(`{
			"aps" : {
				"alert" : "Hello!"
			},
			"data111":{
				"name" : "It's me"
			}
		}
	`)
	client := apns2.NewClient(cert).Development()//   Production()
	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}
