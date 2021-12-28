package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetPrice(coin string) {
	var url = fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=USD", coin)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}
