package main

import "fmt"

func main() {
	fmt.Println("Control Flow")

	age := 22

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	// limit කියන variable එක හදන්නෙත්, condition එක check කරන්නෙත් එකම පේළියේ (;) මඟින් වෙන් කර ඇත.
	if limit := 100; limit > 50 {
		fmt.Println("Limit is high:", limit)
	} else {
		fmt.Println("Limit is fine:", limit)
	}

	// fmt.Println(limit) -> මෙහෙම ලිව්වොත් Error එකක් එනවා!
	// මොකද limit කියන variable එක if එකෙන් එළියේ නැහැ.

	///// switch

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Tuesday":
		fmt.Println("Second day") // මේක Print වෙලා ස්වයංක්‍රීයව Switch එකෙන් එළියට යයි.
	default:
		fmt.Println("Some other day")
	}

	switch day {
	case "Saturday", "Sunday": // අගයන් දෙකෙන් එකක් වුණොත් මේක වැඩ කරනවා
		fmt.Println("It's the Weekend!")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a Weekday.")
	}

	score := 85

    switch {
    case score >= 90:
        fmt.Println("Grade A")
    case score >= 80:
        fmt.Println("Grade B") // 85 නිසා මේක වැඩ කරයි.
    case score >= 70:
        fmt.Println("Grade C")
    default:
        fmt.Println("Grade F")
    }

}
