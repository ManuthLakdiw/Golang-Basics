package main

import (
	"errors"
	"fmt"
)

// =========================================================
// PART 1: සාමාන්‍ය Error Handling (Normal Errors)
// =========================================================

// withdraw: බැංකුවෙන් සල්ලි ගන්න Function එක. 
// සල්ලි මදි වුණොත් මේකෙන් Custom Error එකක් Return කරනවා.
func withdraw(balance int, amount int) (int, error) {
	if amount > balance {
		// සල්ලි මදි නම්: ඉතුරු සල්ලි ගාණයි, Error එකකුයි යවනවා
		return balance, errors.New("ගිණුමේ ප්‍රමාණවත් මුදලක් නොමැත (Insufficient Funds)")
	}
	// සල්ලි තියෙනවා නම්: අලුත් ගාණයි, Error එක nil (හිස්) විදිහටයි යවනවා
	return balance - amount, nil
}

// =========================================================
// PART 2: Exceptions Handling (Panic, Defer, Recover)
// =========================================================

// processCriticalTransaction: අනිවාර්යයෙන්ම වැඩ කරන්න ඕනේ ප්‍රධාන Function එකක්
func processCriticalTransaction() {
	// 1. Defer එක - මේක අනිවාර්යයෙන්ම Function එක ඉවර වෙද්දී (හෝ Crash වෙද්දී) වැඩ කරනවා
	// මේක හරියට Java වල finally block එක සහ catch block එක එකතු කළා වගේ.
	defer func() {
		// 2. Recover එක - පැනපු Panic එක අල්ලගන්නවා
		if r := recover(); r != nil {
			fmt.Println("\n🚨 [RECOVERED] System එක Crash වෙන්න ගියා! හැබැයි අපි බේරගත්තා.")
			fmt.Println("   හේතුව (Panic Message):", r)
		} else {
			fmt.Println("\n✅ [FINALLY] Transaction එක සාර්ථකව අවසන් වුණා.")
		}
	}()

	fmt.Println("\n--- බැංකු සර්වර් එකට සම්බන්ධ වෙමින් පවතී... ---")

	// 3. හිතන්න සර්වර් එක හදිසියේම Down වුණා කියලා. 
	// මේක අපිට සාමාන්‍ය Error එකකින් විසඳන්න බැහැ. ඒ නිසා අපි Panic කරනවා!
	serverIsDown := true 

	if serverIsDown {
		// මේක Run වුණු ගමන් Function එකේ පහළ තියෙන කේත මුකුත් Run වෙන්නේ නැතුව,
		// කෙලින්ම අර උඩ තියෙන 'defer' එක ඇතුළට පනිනවා.
		panic("බැංකුවේ Main Server එක ක්‍රියා විරහිතයි! (Server Disconnected)")
	}

	// Panic එකක් ආපු නිසා මේ පේළිය කවදාවත් Print වෙන්නේ නැහැ.
	fmt.Println("සල්ලි සාර්ථකව මාරු කළා.")
}

// =========================================================
// MAIN FUNCTION
// =========================================================

func main() {
	fmt.Println("=== GO ERROR & EXCEPTION HANDLING DEMO ===")

	// --- 1 වෙනි පරීක්ෂණය: සාමාන්‍ය Errors (if err != nil) ---
	fmt.Println("\n[TEST 1: Normal Errors]")
	myBalance := 1000
	amountToTake := 1500

	fmt.Printf("රු.%d ක ගිණුමෙන් රු.%d ක් ගන්න උත්සාහ කරයි...\n", myBalance, amountToTake)
	
	// Function එක Call කරනවා
	newBalance, err := withdraw(myBalance, amountToTake)

	// Error එකක් ආවද කියලා බලනවා (if err != nil)
	if err != nil {
		fmt.Println("❌ ERROR:", err) // Error එක Print කරනවා
	} else {
		fmt.Println("✅ SUCCESS! අලුත් ශේෂය: රු.", newBalance)
	}

	// --- 2 වෙනි පරීක්ෂණය: සාර්ථකව සල්ලි ගැනීම ---
	fmt.Println("\n[TEST 2: No Errors]")
	newBalance, err = withdraw(myBalance, 200)

	if err != nil {
		fmt.Println("❌ ERROR:", err)
	} else {
		fmt.Println("✅ SUCCESS! අලුත් ශේෂය: රු.", newBalance)
	}

	// --- 3 වෙනි පරීක්ෂණය: Panic සහ Recover ---
	fmt.Println("\n[TEST 3: Panic and Recover]")
	processCriticalTransaction()

	// වැදගත්ම දේ: අර function එක ඇතුළේ Panic එකක් ආවත්, අපි ඒක Recover කරපු නිසා,
	// Program එක නතර වුණේ (Crash වුණේ) නැහැ. මේ අවසාන පේළියත් ලස්සනට වැඩ කරනවා.
	fmt.Println("\n=== PROGRAM එක සාර්ථකව අවසන් විය ===")
}