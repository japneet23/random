/*******
** Delivery Module help to assign requests
** to delivery representatives
******/

package main

import (
	"fmt"
)

var deliveryLocations map[string]int
var deliveryReps []string
var locationReference map[string][]int
var requestAssignment map[string]string

func main() {
	deliveryLocations = map[string]int{"Whitefiled": 0, "Marathahalli": 1, "Hebbal": 2, "BSK": 3, "HSR": 4}
	deliveryReps = []string{"Anand", "Bala", "Chandra"}
	relativeLocation1 := []int{5, 6, 7, 4, 2}
	relativeLocation2 := []int{2, 5, 3, 7, 8}
	relativeLocation3 := []int{7, 3, 6, 5, 2}
	locationReference := make(map[string][]int)
	locationReference["Anand"] = relativeLocation1
	locationReference["Bala"] = relativeLocation2
	locationReference["Chandra"] = relativeLocation3
	requestAssignment = make(map[string]string)
	// var s string
	var d int

	fmt.Printf("\n deliveryLocations %v", deliveryLocations)
	fmt.Printf("\n deliveryReps %v", deliveryReps)
	fmt.Printf("\n locationReference %v", locationReference)
	requestAssignment = LogRequest("Hebbal", "Bala", requestAssignment)

	//fmt.Printf("\n NearBy Rep for %s  %v X", "Hebbal", getNearbyRep("Hebbal", locationReference))
	//Accept the request
	for {
		//Delivery request
		location := ""
		fmt.Println("\n ***Enter Location for the request:")
		fmt.Scanln(&location)
		fmt.Println("\n ***Request for the location:" + location)
		rep := getNearbyRep(location, locationReference)
		fmt.Printf("\n ***Rep Availabilty: %b %s", isRepAvailable(rep), rep)
		if rep == "" {
			fmt.Printf("\n ### Error No Reps Available")
			break
		}
		requestAssignment = LogRequest(location, rep, requestAssignment)
		fmt.Printf("\n ***requestAssignment %v", requestAssignment)
		//fmt.Printf("\n Do you want to process any more request? (0-exit, 1- continue)")
		d++
		fmt.Printf("\n ### Loop : %d", d)
		if d > 5 {
			break
		}
	}
}

func LogRequest(location string, rep string, requestAssignment map[string]string) map[string]string {
	requestAssignment[rep] = location
	return requestAssignment
}
func getNearbyRep(location string, locationReference map[string][]int) string {
	var nearByRep string
	var minDistance int
	var tempMinDistance int = 0
	fmt.Printf("\n Finding nearbyRep for %s", location)
	fmt.Printf("\n locationReference %v", locationReference)
	for k, v := range locationReference {
		if isRepAvailable(k) {
			// locRefList = v
			index := deliveryLocations[location]
			fmt.Printf("\n location index for %s is %d", location, index)
			minDistance = v[index]
			fmt.Printf("\n distance for %s is %d", k, minDistance)
			if tempMinDistance == 0 || tempMinDistance > minDistance {
				nearByRep = k
			}
		}

	}
	fmt.Printf("\n Rep NearBy %s", nearByRep)
	return nearByRep
}
func isRepAvailable(rep string) bool {
	// if requestAssignment[rep] == "" {
	// 	return true
	// } else {
	// 	return false
	// }
	fmt.Printf("\n ####requestAssignment[rep] %v", requestAssignment[rep])

	return requestAssignment[rep] == ""

}
