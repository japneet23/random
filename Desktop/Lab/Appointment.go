package main

import "fmt"

func main() {

	//Doctor's data

	doctors := make(map[string]string)
	doctors["9_12mo"] = "A,C,D"
	doctors["9_12tu"] = "B,D,E"
	doctors["9_12we"] = "A,E"
	doctors["9_12th"] = "B,C"
	doctors["9_12fr"] = "C,D,E"
	doctors["3_6mo"] = "B,E"
	doctors["3_6tu"] = "A,C"
	doctors["3_6we"] = "B,C,D"
	doctors["3_6th"] = "A,D,E"
	doctors["3_6fr"] = "A,B"

	mapSpecialist := make(map[string]int)
	mapSpecialist["A"] = 0
	mapSpecialist["B"] = 1
	mapSpecialist["C"] = 2
	mapSpecialist["D"] = 3
	mapSpecialist["E"] = 4

	mapDay := make(map[string]int)
	mapDay["mo"] = 0
	mapDay["tu"] = 1
	mapDay["we"] = 2
	mapDay["th"] = 3
	mapDay["fr"] = 4

	//Patient data
	var patient_name [30]string
	var patient_phone [30]int

	var time_slot [5][5]int
	timesM := [7]string{"9:00", "9:30", "10:00", "10:30", "11:00", "11:30", "12:00"}
	timesE := [7]string{"3:00", "3:30", "4:00", "4:30", "5:00", "5:30", "6:00"}

	i := 0 //counter for name
	j := 0 //counter for number

	patient_name[i] = "Aman"
	time_slot[1][0] += 1
	i++
	patient_name[i] = "Rahul"
	time_slot[4][3] += 1
	i++
	patient_name[i] = "Riya"
	time_slot[2][4] += 1
	i++
	patient_name[i] = "Nisha"
	time_slot[0][1] += 1
	i++
	patient_name[i] = "Roy"
	time_slot[3][2] += 1
	patient_phone[j] = 1234567890
	j++
	patient_phone[j] = 2345678901
	j++
	patient_phone[j] = 3456789012
	j++
	patient_phone[j] = 4567890123
	j++
	patient_phone[j] = 5678901234

	fmt.Println("Enter the Number of Patients you want to register for")
	var x int
	fmt.Scanln(&x)
	for k := 0; k < x; k++ {
		var name string
		var phone int
		fmt.Println("Enter name and number")
		fmt.Scanln(&name)
		fmt.Scanln(&phone)
		flag := 0

		//for phone check
		count := 0
		checknum := phone

		//check for 10 digits number
		for checknum > 0 {
			checknum = checknum / 10
			count += 1
		}
		if count < 10 {
			fmt.Println("Wrong Number")
			break
		}

		//check if number already exists or not
		for l := 0; l < j; l++ {
			if phone == patient_phone[l] {
				fmt.Println("User already exits.Try again")
				flag = 1
				break
			}
		}
		if flag == 0 {

			patient_name[i] = name
			i++

			patient_phone[j] = phone
			j++
		}

		//end of phone check

		//giving time slot to patient
		var day string
		var time string
		fmt.Println("Enter the day [mo,tu,we,th,fr]")

		fmt.Scanln(&day)

		fmt.Println("Enter the time slot(9_12 or 3_6)")

		fmt.Scanln(&time)

		userInput := time + day
		//fmt.Printf("%s",userInput)

		specialists := doctors[userInput]
		fmt.Println("Choose one of the specialist", specialists)
		fmt.Println("Press NO for other specialists")

		var option string
		fmt.Scanln(&option)

		if option != "NO" {

			length := len(userInput)
			day2 := userInput[length-2 : length] //to take the day from userinput
			dayIndex := mapDay[day2]             //to take the value of day
			specialistIndex := mapSpecialist[option]

			if time_slot[dayIndex][specialistIndex] != 6 { //slot allocation
				freq := time_slot[dayIndex][specialistIndex]

				var res string
				if time == "9_12" {
					res = timesM[freq]
				} else {
					res = timesE[freq]
				}

				time_slot[dayIndex][specialistIndex] += 1 //updating frequency

				fmt.Printf("\nYour appointment has been booked for %s at time %s ", day, res)

			} else {
				fmt.Printf("Time slot is full.Try Another Day")
				k = k - 1
			}

		} else {

			length := len(userInput)
			day2 := userInput[length-2 : length] //to take the day from userinput
			dayIndex := mapDay[day2]
			if time == "9_12" {
				time = "3_6"
				userInput := time + day
				specialists := doctors[userInput]
				fmt.Println("Choose one of the specialist", specialists)

				var option string
				fmt.Scanln(&option)

				specialistIndex := mapSpecialist[option]

				if time_slot[dayIndex][specialistIndex] != 6 { //slot allocation
					freq := time_slot[dayIndex][specialistIndex]

					var res string
					res = timesM[freq]

					time_slot[dayIndex][specialistIndex] += 1 //updating frequency

					fmt.Printf("\nYour appointment has been booked for %s at time %s ", day, res)

				} else {
					fmt.Printf("Time slot is full.Try Another Day")
					k = k - 1
				}
			} else {
				time = "9_12"
				userInput := time + day
				specialists := doctors[userInput]
				fmt.Println("Choose one of the specialist", specialists)

				var option string
				fmt.Scanln(&option)

				specialistIndex := mapSpecialist[option]

				if time_slot[dayIndex][specialistIndex] != 6 { //slot allocation
					freq := time_slot[dayIndex][specialistIndex]

					var res string

					res = timesE[freq]

					time_slot[dayIndex][specialistIndex] += 1 //updating frequency

					fmt.Printf("\nYour appointment has been booked for %s at time %s ", day, res)

				} else {
					fmt.Printf("Time slot is full.Try Another Day")
					k = k - 1
				}
			}

		}

	}

}
