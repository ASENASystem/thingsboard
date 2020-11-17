package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"gitlab.com/asenasystem-opensource/go/thingsboard"
)

type cubeCell struct {
	Battery  string `json:"battery"`
	Distance string `json:"distance"`
}

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

	// fmt.Printf("\n\nBearer " + tb.Auth.Token + "\n\n")

	fmt.Print("Device: 'test' ID: ")
	device1, _ := tb.GetDeviceByName("test")
	fmt.Printf("%+v\n", device1.ID.ID)

	// if d, err := tb.GetDeviceByID(device1.ID.ID); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("\n\n%+v\n\n", d)
	// }

	// fmt.Printf("%+v\n", tb.User)
	// key -> accessToken
	devices := make(map[string]string)
	devices["test"] = ""
	devices["test2"] = ""

	cc := cubeCell{
		Battery:  "4.321",
		Distance: "167",
	}

	for deviceName := range devices {
		token, err := tb.GetDeviceAccessTokenByName(deviceName)
		if err != nil {
			fmt.Println(err)
		}
		devices[deviceName] = token

		fmt.Println(cc)
		err = tb.SaveTelemetry(devices["test"], cc)
		if err != nil {
			fmt.Println(err)
		}

	}

	if err := tb.Disconnect(); err != nil {
		fmt.Println(err)
	}

	os.Exit(0)
}
