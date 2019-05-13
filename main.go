package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func toggleDeviceHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	deviceID := keys.Get("device")
	value := keys.Get("checked")

	fmt.Println("Url Param 'device' is: " + string(deviceID) + " and its value is: " + string(value))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Main.html", nil)
}

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
	http.HandleFunc("/Main/getDevices", getDevicesHandler)
	http.HandleFunc("/Main/toggleDevice/", toggleDeviceHandler)
	http.HandleFunc("/Main", mainHandler)

	http.Handle("/sup/", http.StripPrefix("/sup", http.FileServer(http.Dir("sup/"))))

	tpl = template.Must(template.ParseGlob("*.html"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
