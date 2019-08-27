package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

var reverseDir = map[string]string{"north": "south", "south": "north", "east": "west", "west": "east"}  // The reverse direction for each direction

// Function to add two cities to each others connection
func addNeighhbor(city1 string, city2 string, dir string, cityConnects map[string]map[string]string) {
	// If no connection for city1, add the connections for it.
	if (cityConnects[city1] == nil){
		cityConnects[city1] = map[string]string{"north": city1, "south": city1, "east": city1, "west": city1}
	}
	// Add city2 to city1's connection
	cityConnects[city1][dir] = city2
	// If no connection for city1, add the connections for it.
	if (cityConnects[city2] == nil){
		cityConnects[city2] = map[string]string{"north": city2, "south": city2, "east": city2, "west": city2}
	}
	// Add city1 to city2's connection
	cityConnects[city2][reverseDir[dir]] = city1
}

func main() {
	//Get the file path
	var filePath string
	fmt.Printf("What's the path for the map you want to use? ")
	fmt.Scanln(&filePath)

	//Open the file and read it
	fi, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)

	var cityConnects = map[string]map[string]string{}  // The map for the city connections

	lines := strings.Split(string(fd), "\n")
	for _, line := range lines {
		cities := strings.Split(line, " ")

		//Get the city name
		city1 := cities[0]

		// Add connections between each city
		for _, cityDir := range cities[1:] {
			dirAndCity := strings.Split(cityDir, "=")
			dir := dirAndCity[0]
			city2 := dirAndCity[1]
			addNeighhbor(city1, city2, dir, cityConnects)
		}
	}
	fmt.Println(cityConnects)
}
