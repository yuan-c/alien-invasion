package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
)

var dirs = []string{"north", "south", "west", "east"}  // Directions
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

// Function to detect if there are aliens in the same city, if so, they fight and destroy the city
func alienFight(alienLocation map[int]string, cityConnects map[string]map[string]string) {
	cityContains := make(map[string][]int)  // Which aliens contained in each city

	// Put aliens in the city they stay
	for k,v := range alienLocation {
		cityContains[v] = append(cityContains[v], k)
	}

	var shouldDelete []string   // Cities should be destroyed after the aliens fight

	for k,v := range cityContains {
		// If any city contains more than one aliens, they fight
		if len(v) > 1 {
			if len(v) == 2 {    // Only two alines in the city
				fmt.Printf("%v has been destroyed by alien-%v and alien-%v\n", k, v[0], v[1])
			} else {            // More than two aliens in the city
				fmt.Printf("%v has been destroyed by ", k)
				for i := 0; i < len(v)-1 ; i++ {
					fmt.Printf("alien-%v, ", v[i])
				}
				fmt.Printf("and alien-%v\n", v[len(v)-1])
			}
			// Aliens died after the fight
			for _, alien := range v {
				delete(alienLocation, alien)
			}

			// Since the city should be destroyed, other cities cannot go to that city anymore
			for dir,city := range cityConnects[k] {
				if city != k {
					cityConnects[city][reverseDir[dir]] = city
					cityConnects[k][dir] = k
				}
			}

			shouldDelete = append(shouldDelete, k)    // Add the city to should delete
		}
	}

	// Delete those cities which should be delete from the city connection
	for _, city := range shouldDelete {
		delete(cityConnects, city)
	}
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

	// For each city, add the connection between itself and its neighbors
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

	var cities []string  // All the cities in the map

	for city := range cityConnects {
        cities = append(cities, city)
    }

	//Get the aline numbers
	var alienNums int
	fmt.Printf("How many aliens do you want to put in the world? ")
	fmt.Scanln(&alienNums)

	var alienLocation = map[int]string{}   // The location for each alien

	// For each alien, put it in a random city
	for n := 0; n < alienNums; n++ {
		rand.Seed(time.Now().UnixNano())
		alienLocation[n] = cities[rand.Intn(len(cities))]
   	}

	var moves int = 10000    // The maximum movements for aliens

	// If aliens did not reach the maximum movements, they keep moving till all aliens died
	for moves > 0 {
		alienFight(alienLocation, cityConnects)   // Aliens fight before each move

		// It should stop if no alive alines
		if (len(alienLocation) == 0) {
			break
		}

		// Each alive alien moves to a random direction
		for k,v := range alienLocation {
			rand.Seed(time.Now().UnixNano())
			alienLocation[k] = cityConnects[v][dirs[rand.Intn(4)]]
		}
		moves--
	}

	if (len(cityConnects) == 0) {           // If all cities has been destroyed
		fmt.Println("All cities are destroyed")
	} else if (len(cityConnects) == 1) {    // If one city left
		fmt.Printf("There is only one city left and it's ")
		for k,_ := range cityConnects {
			fmt.Printf("%v\n", k)
		}
	} else {                                // If more than two cities left
		fmt.Printf("The left cities and their connections are:\n")
		for k,v := range cityConnects {
			fmt.Printf("%v ", k)
			for dir,city := range v {
				if city != k {              // Output the connections of this city
					fmt.Printf("%v: %v ", dir, city)
				}
			}
			fmt.Printf("\n")
		}
	}
}
