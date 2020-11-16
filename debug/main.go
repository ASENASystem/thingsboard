package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"gitlab.com/asenasystem-opensource/go/thingsboard"
)

func main() {

	host := os.Getenv("TB_HOST")
	user := os.Getenv("TB_USER")
	pass := os.Getenv("TB_PASS")

	if host == "" {
		log.Error("TB_HOST missing")
		os.Exit(1)
	}
	if user == "" {
		log.Error("TB_USER missing")
		os.Exit(1)
	}
	if pass == "" {
		log.Error("TB_PASS missing")
		os.Exit(1)
	}

	tb, err := thingsboard.New(host, user, pass)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("%+v\n", tb.User)
	// for key, element := range tb.User.AdditionalInfo.UserPasswordHistory {
	// 	fmt.Println("Key:", key, "=>", "Element:", element)
	// }

	fmt.Printf("\n\nBearer " + tb.Auth.Token + "\n\n")

	fmt.Print("Looking for device ID named: 'test': ")
	device1, _ := tb.GetDeviceByName("test")
	fmt.Printf("%+v\n", device1.ID.ID)

	if d2, err := tb.GetDeviceByID(device1.ID.ID); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n\n%+v\n\n", d2)
	}

	if err := tb.Disconnect(); err != nil {
		fmt.Println(err)
	}

	os.Exit(0)

}
