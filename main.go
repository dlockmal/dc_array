package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a list of hosts
	hosts := []string{"casper.ash1", "larian.dc2", "apple.dc1", "banana.dc2",
		"cherry.ash1", "date.dfw1", "elderberry.dfw2", "fig.dc1", "grape.dc2", "honeydew.dc2",
		"kiwi.dc2", "lemon.ash1", "mango.dc2", "nectarine.dc1", "orange.dc1", "pear.dc1",
		"quince.dc1", "raspberry.dc1", "strawberry.dc1", "tangerine.ash1", "ugli.dc1",
		"vanilla.dc1", "watermelon.dc2", "ximenia.ash1", "yuzu.dc1", "zucchini.dc1"}

	// Create a map of datacenters
	datacenterMap := make(map[string][]string)

	// Loop through the hosts and add them to the datacenter map
	for _, host := range hosts {
		splitHost := strings.Split(host, ".")
		server, datacenter := splitHost[0], splitHost[1]

		datacenterMap[datacenter] = append(datacenterMap[datacenter], server)
	}

	fmt.Println(datacenterMap)

}
