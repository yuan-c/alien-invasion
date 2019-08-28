package main

import (
	"testing"
	"reflect"
)

// The function to test if addNeighbor works correctly
func TestAddNeighbor(t *testing.T) {
	cityConnects := map[string]map[string]string{}

	// Add city2 to the north of city1. For city1, only north should be city2, other direction should point to itself
	addNeighbor("city1", "city2", "north", cityConnects)
	expected := map[string]map[string]string{"city1": map[string]string{"east":"city1", "north":"city2", "south":"city1", "west":"city1"},
	"city2": map[string]string{"east":"city2", "north":"city2", "south":"city1", "west":"city2"}}
	if !reflect.DeepEqual(expected, cityConnects) {
		t.Errorf("The city connections are incorrect. Expected %s, but was %s.", expected, cityConnects)
	}

	// Add city3 to the north of city2. For city2, north should be city3, south should be city1, other direction should point to itself
	addNeighbor("city3", "city2", "south", cityConnects)
	expected = map[string]map[string]string{"city1": map[string]string{"east":"city1", "north":"city2", "south":"city1", "west":"city1"},
	"city2": map[string]string{"east":"city2", "north":"city3", "south":"city1", "west":"city2"},
	"city3": map[string]string{"east":"city3", "north":"city3", "south":"city2", "west":"city3"}}
	if !reflect.DeepEqual(expected, cityConnects) {
		t.Errorf("The city connections are incorrect. Expected %v, but was %v.", expected, cityConnects)
	}
}

// The function to test if alienFight works correctly
func TestAlienFight(t *testing.T) {
	// First, we have the initial alien lication and city connections
	alienLocation := map[int]string{0:"city1", 1:"city2", 2:"city1", 3:"city3"}
	cityConnects := map[string]map[string]string{"city1": map[string]string{"east":"city1", "north":"city2", "south":"city1", "west":"city1"},
	"city2": map[string]string{"east":"city2", "north":"city3", "south":"city1", "west":"city2"},
	"city3": map[string]string{"east":"city3", "north":"city3", "south":"city2", "west":"city3"}}

	// Since alien0 and alien2 both in city1, they should fight and destroy city1
	alienFight(alienLocation, cityConnects)

	// Alien0 and alien 2 should die in the fight, only alien1 and alien3 are alive
	expextedAlienLocation := map[int]string{1:"city2", 3:"city3"}
	if !reflect.DeepEqual(expextedAlienLocation, alienLocation) {
		t.Errorf("The aliens' location are incorrect. Expected %v, but was %v.", expextedAlienLocation, alienLocation)
	}

	// The city1 should be destroyed and the south of city2 should change to itself
	expextedCityConnects := map[string]map[string]string{"city2": map[string]string{"east":"city2", "north":"city3", "south":"city2", "west":"city2"},
	"city3": map[string]string{"east":"city3", "north":"city3", "south":"city2", "west":"city3"}}
	if !reflect.DeepEqual(expextedCityConnects, cityConnects) {
		t.Errorf("The city connections are incorrect. Expected %v, but was %v.", expextedCityConnects, cityConnects)
	}

	// Fight again withour move the aliens, since aliens all stay in different cities, no fight occurs
	alienFight(alienLocation, cityConnects)

	// Everything should stay in the same as before
	if !reflect.DeepEqual(expextedAlienLocation, alienLocation) {
		t.Errorf("The aliens' location are incorrect. Expected %v, but was %v.", expextedAlienLocation, alienLocation)
	}
	if !reflect.DeepEqual(expextedCityConnects, cityConnects) {
		t.Errorf("The city connections are incorrect. Expected %v, but was %v.", expextedCityConnects, cityConnects)
	}
}