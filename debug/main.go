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

	fmt.Printf("%+v\n", tb.User)

	for key, element := range tb.User.AdditionalInfo.UserPasswordHistory {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	fmt.Println("Bearer " + tb.Auth.Token)

	tb.GetTenantDevices(10, 0)

	err = tb.Disconnect()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(0)

}
