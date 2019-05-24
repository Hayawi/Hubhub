package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
var clientID = "dd1a4123-5130-4907-a304-4a19ad2c181a"
var secretID = "1770d1b6-3ace-4a8b-9559-4c0abb20863f"

var presets = map[string][]string{
	"Fajr": []string{"1a95563f-14d8-4fe9-a1c0-47255d79804f", "1dd43892-e83e-42f2-9c9b-6325afcf5dd0",
		"94b57bdd-df9c-4282-b9ed-bd2648e3012f"},
	"Dhuhr": []string{"1a95563f-14d8-4fe9-a1c0-47255d79804f", "1dd43892-e83e-42f2-9c9b-6325afcf5dd0",
		"94b57bdd-df9c-4282-b9ed-bd2648e3012f", "f07f74a7-a654-41e5-8fa3-7ebbebb93a83", "bc2abaec-42c0-4c4c-8958-8014bac2619d"},
}

func presetTitles(w http.ResponseWriter, r *http.Request) {
	mapOfIDToName(&clientID, &secretID)
	responseString := ""
	for name := range presets {
		responseString += name + " "
	}

	if len(responseString) > 1 {
		responseString = responseString[:len(responseString)-1]
	}

	fmt.Fprintf(w, "%v", responseString)
}

func presetPicked(w http.ResponseWriter, r *http.Request) {
	responseString := []string{}
	keys := r.URL.Query()
	preset := keys.Get("preset")

	status := false

	for id := range IDToName {
		responseString = append(responseString, id+"|off")
		status = SetSwitch(&clientID, &secretID, id, "off")
		if !status {
			fmt.Println("Setting device: " + string(id) + " to value: off failed!")
			tpl.ExecuteTemplate(w, "Main.html", nil)
		}
	}

	for _, id := range presets[preset] {
		responseString = append(responseString, id+"|on")
		status = SetSwitch(&clientID, &secretID, id, "on")
		if !status {
			fmt.Println("Setting device: " + string(id) + " to value: on failed!")
			tpl.ExecuteTemplate(w, "Main.html", nil)
		}
	}

	fmt.Fprintf(w, "%v", responseString)
}

func toggleDeviceHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	deviceID := keys.Get("device")
	value := keys.Get("checked")

	status := false
	if value == "true" {
		status = SetSwitch(&clientID, &secretID, deviceID, "on")
	} else {
		status = SetSwitch(&clientID, &secretID, deviceID, "off")
	}
	if !status {
		fmt.Println("Setting device: " + string(deviceID) + " to value: " + string(value) + " failed!")
		tpl.ExecuteTemplate(w, "Main.html", nil)
	}
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
	devices := GetDevices(&clientID, &secretID)
	return devices
}

func savePreset(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	name := keys.Get("presetName")

	ids := []string{}
	for _, inpIds := range keys["ids"] {
		ids = append(ids, inpIds)
	}
	presets[name] = ids
	fmt.Println(presets)
}

func main() {
	http.HandleFunc("/Main/getDevices", getDevicesHandler)
	http.HandleFunc("/Main/toggleDevice/", toggleDeviceHandler)
	http.HandleFunc("/Main/presets", presetTitles)
	http.HandleFunc("/Main/presetPicked/", presetPicked)
	http.HandleFunc("/Main/savePreset/", savePreset)
	http.HandleFunc("/Main", mainHandler)

	http.Handle("/sup/", http.StripPrefix("/sup", http.FileServer(http.Dir("sup/"))))

	tpl = template.Must(template.ParseGlob("*.html"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
