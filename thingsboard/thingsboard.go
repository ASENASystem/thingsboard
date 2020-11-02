// Package thingsboard ...
// Author: Paweł 'felixd' Wojciechowski - Outsourcing IT Konopnickiej.Com
// (c) 2020 ASENA System
package thingsboard

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// ‘X-Authorization’ to “Bearer $YOUR_JWT_TOKEN”

type Thingsboard struct {
}

func main() {

	// Create a Resty Client
	client := resty.New()

	apiLink := "/api/auth/login"
	tbUser := os.Getenv("TB_USER")
	tbPass := os.Getenv("TB_PASS")
	tbHost := os.Getenv("TB_HOST")

	authSuccess := new(AuthSuccess)

	// POST JSON string
	// No need to set content type, if you have client level setting
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"` + tbUser + `", "password":"` + tbPass + `"}`).
		SetResult(&authSuccess). // or SetResult(&AuthSuccess{}).
		// EnableTrace().
		Post(tbHost + apiLink)

	printOutput(resp, err)

	// // Explore response object
	// fmt.Println("Response Info:")
	// fmt.Println("Error      :", err)
	// fmt.Println("Status Code:", resp.StatusCode())
	// fmt.Println("Status     :", resp.Status())
	// fmt.Println("Proto      :", resp.Proto())
	// fmt.Println("Time       :", resp.Time())
	// fmt.Println("Received At:", resp.ReceivedAt())
	// fmt.Println("Body       :\n", resp)
	// fmt.Println()

	// // Explore trace info
	// fmt.Println("Request Trace Info:")
	// ti := resp.Request.TraceInfo()
	// fmt.Println("DNSLookup    :", ti.DNSLookup)
	// fmt.Println("ConnTime     :", ti.ConnTime)
	// fmt.Println("TCPConnTime  :", ti.TCPConnTime)
	// fmt.Println("TLSHandshake :", ti.TLSHandshake)
	// fmt.Println("ServerTime   :", ti.ServerTime)
	// fmt.Println("ResponseTime :", ti.ResponseTime)
	// fmt.Println("TotalTime    :", ti.TotalTime)
	// fmt.Println("IsConnReused :", ti.IsConnReused)
	// fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	// fmt.Println("ConnIdleTime :", ti.ConnIdleTime)

	fmt.Println(authSuccess.RefreshToken)
	fmt.Println(authSuccess.Token)

}

func printOutput(resp *resty.Response, err error) {
	fmt.Println(resp, err)
}
