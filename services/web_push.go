package services

import (
	"encoding/json"
	"fmt"
	"log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

const (
	subscription    = `{"endpoint":"https://fcm.googleapis.com/fcm/send/eivFwFS3VdE:APA91bHTuJ1q_EtwduZn7brNtYKniWGrmXyEBovTacQCAQi_Jf-94CXv30lKQDDWV7dBamrWaOTj3GE-kEm2oOUJmK4IqgURbV7j23q8ozN8iQoWV-vm21DP3VV45JQzJgDoZYqzHgLn","expirationTime":null,"keys":{"p256dh":"BG3FcqS9QUnwb4Y7YahAnUTMpFwkYhdJgRwKjnO_TexPuAzF7Btbdv2UYVn0hOVRnodm9om2UO3I5x61lLy4kBc","auth":"KaYRivM7P-pXNiYaWiq2KQ"}}`
	vapidPublicKey  = "BL0TX8utP5uYL8Zzqcf6xk7ALK9aTe8RiTqmsjatwbg5tzx651rTXmigD8dLiuv1LO9acQd-83r2M08P5uKHZdQ"
	vapidPrivateKey = "ebEbTNvVl_sWqqLwo2lxnHjGYQp_88YBYt59lmQgq3U"
)

func GenerateVapidKeys() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		// TODO: Handle error
	}
	fmt.Println(privateKey, "private key")
	fmt.Println(publicKey, "public key")
}

func SendPushNotification() {
	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      "example@example.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
