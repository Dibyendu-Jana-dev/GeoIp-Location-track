package main

import (
	"fmt"
	"github.com/oschwald/maxminddb-golang"
	"net"
)

type Location struct {
	City struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Country struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	Postal struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"postal"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
		AccuracyRadius int `maxminddb:"accuracy_radius"`
	} `maxminddb:"location"`
	//ISP string `maxminddb:"isp"`
	//Domain string `maxminddb:"domain"`
	//ConnectionType string `maxminddb:"connection_type"`
}

func main() {
	// Open the GeoIP2 database
	db, err := maxminddb.Open("./GeoLite2-City.mmdb")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// IP address to lookup
	ipAddress := "110.226.50.38"

	// Parse the IP address
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		fmt.Println("Invalid IP address")
		return
	}

	// Perform the lookup
	var location Location
	err = db.Lookup(ip, &location)
	if err != nil {
		fmt.Println("Error looking up IP:", err)
		return
	}

	// Print the retrieved information
	fmt.Printf("IP Address: %s\n", ipAddress)
	fmt.Printf("Location: %s, %s\n", location.City.Names["en"], location.Country.Names["en"])
	fmt.Printf("Postal Code: %s\n", location.Postal.Code)
	fmt.Printf("Latitude / Longitude: %f, %f (Accuracy Radius: %d meters)\n", location.Location.Latitude, location.Location.Longitude, location.Location.AccuracyRadius)
	//fmt.Printf("ISP: %s\n", location.ISP)
	//fmt.Printf("Domain: %s\n", location.Domain)
	//fmt.Printf("Connection Type: %s\n", location.ConnectionType)
}

