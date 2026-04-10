package main

import "fmt"

func main() {
	// 1. Map එකක් සාදා Data ඇතුළත් කිරීම (Literal method)
	capitals := map[string]string{
		"Sri Lanka": "Colombo",
		"Japan":     "Tokyo",
		"USA":       "Washington D.C.",
	}

	// 2. make() හරහා Map එකක් සෑදීම සහ Data ඇතුළත් කිරීම
	scores := make(map[string]int)
	scores["Manuth"] = 95
	scores["Kasun"] = 80

	// 3. සම්පූර්ණ Map එකම Print කිරීම
	fmt.Println("All Capitals:", capitals)

	// 4. Key එකක් මඟින් Value එකක් ගැනීම
	myScore := scores["Manuth"]
	fmt.Println("Manuth's Score:", myScore)

	// 5. නැති Key එකක් Access කිරීම සහ (Comma ok) පරීක්ෂාව
	// මෙහි Value data type එක int නිසා, නැතිනම් 0 return වේ.
	score, exists := scores["Nimal"]
	if exists {
		fmt.Println("Nimal's score is:", score)
	} else {
		fmt.Println("Nimal Map එකේ නැහැ! (Zero Value returned:", score, ")")
	}

	// 6. delete() Method එක භාවිතය
	delete(capitals, "USA")
	fmt.Println("After deleting USA:", capitals)

	// 7. Iterating (Key සහ Value වෙන වෙනම ගැනීම)
	fmt.Println("--- Iterating Map ---")
	for country, capital := range capitals {
		fmt.Printf("Country: %s -> Capital: %s\n", country, capital)
	}

	// 8. Values පමණක් Iteration එකෙන් ගැනීම
	fmt.Println("--- Only Values ---")
	for _, capName := range capitals {
		fmt.Println("Capital:", capName)
	}
}