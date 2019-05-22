// Simple example for the gosmart libraries.
//
// This is a simple demonstration of how to obtain a token from the smartthings
// API using Oauth2 authorization, and how to request the status of some of your
// sensors (in this case, temperature).
//
// This file is part of gosmart, a set of libraries to communicate with
// the Samsumg SmartThings API using Go (golang).
//
// http://github.com/marcopaganini/gosmart
// (C) 2016 by Marco Paganini <paganini@paganini.net>

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/marcopaganini/gosmart"
)

const (
	tokenFilePrefix = ".example_st_token"
)

var (
	flagClient    = flag.String("client", "", "OAuth Client ID")
	flagSecret    = flag.String("secret", "", "OAuth Secret")
	flagTokenFile = flag.String("tokenfile", "", "Token filename")
	flagDevID     = flag.String("devid", "", "Show information about this particular device ID")
	flagAll       = flag.Bool("all", false, "Show Information about all devices found")
)

var client *http.Client
var endpoint string

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//GetDevices returns a slice of devices as "|" seperated strings
func GetDevices(clientID *string, secretID *string) []string {
	accessAPI(clientID, secretID)

	devices := []string{}
	devs, err := gosmart.GetDevices(client, endpoint)
	check(err)
	for _, d := range devs {
		dInfo, err := gosmart.GetDeviceInfo(client, endpoint, d.ID)
		check(err)
		if dInfo.Attributes["switch"] == "on" {
			devices = append(devices, d.ID+"|"+d.DisplayName+"|on")
		} else {
			devices = append(devices, d.ID+"|"+d.DisplayName+"|")
		}
		fmt.Printf("ID: %s, Name: %q, Display Name: %q\n", d.ID, d.Name, d.DisplayName)
	}
	return devices
}

// SetSwitch sets the value of the switch parameter on the given device
func SetSwitch(clientID *string, secretID *string, deviceID string, switchValue string) bool {
	accessAPI(clientID, secretID)

	status, err := gosmart.SendDeviceCommands(client, endpoint, deviceID, switchValue)
	check(err)
	return status
}

func accessAPI(clientID *string, secretID *string) {
	flagClient = clientID
	flagSecret = secretID

	// No date on log messages
	log.SetFlags(0)

	// If we have a token file from the command line, use that directly.
	// Otherwise, form the name from tokenFilePrefix and the Client ID.
	tfile := *flagTokenFile
	if tfile == "" {
		if *flagClient == "" {
			log.Fatalf("Must specify Client ID or Token File")
		}
		tfile = tokenFilePrefix + "_" + *flagClient + ".json"
	}

	// Create the oauth2.config object and get a token
	config := gosmart.NewOAuthConfig(*flagClient, *flagSecret)
	token, err := gosmart.GetToken(tfile, config)
	flagTokenFile = &tfile
	check(err)

	// Create a client with the token. This client will be used for all ST
	// API operations from here on.
	ctx := context.Background()
	client = config.Client(ctx, token)

	// Retrieve Endpoints URI. All future accesses to the smartthings API
	// for this session should use this URL, followed by the desired URL path.
	endpoint, err = gosmart.GetEndPointsURI(client)
	check(err)
}
