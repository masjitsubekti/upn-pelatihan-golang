package notification

import (
	"fmt"
	"log"

	"github.com/douglasmakey/go-fcm"
)

func PushNotification(data map[string]interface{}, token []string) (err error) {
	// init client
	// SERVER KEY GSI-ESDM

	client := fcm.NewClient("AAAAeCPPrTQ:APA91bEHMItvOOV7ZKZ5efpb176XoUVZU8AZ7ofgG3M9SXw_DoBkxZgOrMVJpnLzLLpRRYqxqsUe_3tDSdIJzBQStN2M1N6APm4Px9YthtKmRz2Q3E66VP0IR_kcP25Ky6pyKgHb3jEE")
	// // CLIENT TESTING
	// client := fcm.NewClient("AAAATA_--1w:APA91bFNSY0P5Lb3Xd1NBYu3zH2TQFZpEwS84mAnZiPkErYv78L94LFrJcKD8NtqGXu37Uo_caQhqzTcGeSstA7XHIrazlNthGbB3zPIJG8AjnJ7q4PSbgR0nEKACI0gw_l7_fiEDnp4")
	// client := fcm.NewClient("AAAATA_--1w:APA91bFNSY0P5Lb3Xd1NBYu3zH2TQFZpEwS84mAnZiPkErYv78L94LFrJcKD8NtqGXu37Uo_caQhqzTcGeSstA7XHIrazlNthGbB3zPIJG8AjnJ7q4PSbgR0nEKACI0gw_l7_fiEDnp4")
	// CLIENT GSI-AKUNTING
	// client := fcm.NewClient("AAAAD0yBNGE:APA91bGkQtweh-B8541fLt2H3CLHBYd3nYSwR95Op7reY1qocm-MjxcLUvM2WiPZMzTK6twJfQ-xRRVUJVq-sSCveaGWN8Ep8oaFDo7wY-LCwWPf8c-CB7x8B-G6NRqW5sGDWbfConBk")
	// You can use your HTTPClient
	//client.SetHTTPClient(client)

	// You can use PushMultiple or PushSingle
	client.PushMultiple(token, data)
	//client.PushSingle("token 1", data)

	// registrationIds remove and return map of invalid tokens
	badRegistrations := client.CleanRegistrationIds()
	log.Println(badRegistrations)

	status, err := client.Send()
	fmt.Println("ERROR PUSH:", err)
	fmt.Println("STATUS PUSH", status)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return err
}
