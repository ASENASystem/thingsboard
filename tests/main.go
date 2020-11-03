package main

import (
	"fmt"
	"os"

	"gitlab.com/asenasystem-opensource/go/thingsboard"
)

func main() {

	tbHost := os.Getenv("TB_HOST")
	tbUser := os.Getenv("TB_USER")
	tbPass := os.Getenv("TB_PASS")

	tb, err := thingsboard.New(tbHost, tbUser, tbPass)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tb)

}
