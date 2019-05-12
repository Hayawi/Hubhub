package main

import (
	"fmt"
	"log"
	"net/http"
)

func getDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices := getDevices()
	responseString := "["
	for _, d := range devices {
		responseString += d + ","
	}
	if len(responseString) > 1 {
		responseString = responseString[:len(responseString)-1]
	}
	responseString += "]"
	fmt.Fprintf(w, "%v", responseString)
}

func getDevices() []string {
	client := "dd1a4123-5130-4907-a304-4a19ad2c181a"
	secret := "1770d1b6-3ace-4a8b-9559-4c0abb20863f"
	devices := GetDevices(&client, &secret)
	return devices
}

func main() {
	http.HandleFunc("/getDevices", getDevicesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
