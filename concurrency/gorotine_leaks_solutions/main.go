//WithCancel
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// // මෙතන අනිවාර්ය Go නීතියක් තියෙනවා: Context එකක් Function එකකට යවද්දී 
// // අනිවාර්යයෙන්ම ඒක පළවෙනි Parameter එක විදිහටම යවන්න ඕනේ. (සම්මතයයි).
// func complexDatabaseTask(ctx context.Context, workerID int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// Remote එකෙන් Cancel කරපු ගමන් මේ Channel එකට Signal එක එනවා!
// 			fmt.Printf("Worker %d: 🛑 Cancel Signal එක ආවා! මම Database වැඩේ නැවැත්තුවා.\n", workerID)
// 			return // Function එකෙන් එළියට පනිනවා (Goroutine එක අවසන් වෙනවා)

// 		default:
// 			// Cancel කරලා නැත්නම්, මේක ඇතුළේ වැඩේ දිගටම කරගෙන යනවා
// 			fmt.Printf("Worker %d: ⏳ Database එකේ Data හොයනවා...\n", workerID)
// 			time.Sleep(500 * time.Millisecond) // වැඩේට යන වෙලාව
// 		}
// 	}
// }

// func main() {
// 	// 1. Cancel කරන්න පුළුවන් Remote එකක් (Context) හදාගන්නවා
// 	ctx, cancel := context.WithCancel(context.Background())

// 	// 2. ඒ Context එක Workers ලාට Pass කරනවා
// 	go complexDatabaseTask(ctx, 1)
// 	go complexDatabaseTask(ctx, 2)

// 	// Main එක තත්පර 2ක් වෙන වැඩක් කරනවා කියලා හිතමු (Users බලන් ඉන්න වෙලාව)
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Main: 🚨 User Browser එක Close කළා! ඔක්කොම වැඩ නවත්තන්න (Cancel)!")
	
// 	// 3. Remote එකේ Cancel බොත්තම ඔබනවා! 
// 	// මේක එබුවාම අර ctx.Done() එකෙන් බලන් ඉන්න ඔක්කොටම Signal එක යනවා.
// 	cancel() 

// 	// Workers ලා නතර වෙනකම් තත්පරයක් විතර බලන් ඉඳලා Program එක ඉවර කරනවා
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main: Program එක සාර්ථකව අවසන් වුණා.")
// }

//WithTimeout

package main

import (
	"context"
	"fmt"
	"time"
)

func fetchUserFromDB(ctx context.Context) {
	// Database එකෙන් Data එන්න තත්පර 5ක් යනවා කියලා හිතමු
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("✅ Database: Data ටික සාර්ථකව ගත්තා.")
	case <-ctx.Done():
		// හැබැයි තත්පර 5 යන්න කලින් Timeout එක පැනලා Cancel වුණොත් මේක වැඩ කරනවා
		fmt.Println("🛑 Database: Request එක Timeout වුණා! Query එක Cancel කළා.")
	}
}

func main() {
	// තත්පර 2කින් ස්වයංක්‍රීයවම Cancel වෙන Context එකක් හදනවා
	// (Database එකට තත්පර 5ක් යන නිසා අනිවාර්යයෙන්ම මේක Timeout වෙන්න ඕනේ)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // මතක ඇතුව හැමවිටම අන්තිමට cancel() කතා කරන්න ඕනේ (Memory leaks නවත්තන්න)

	fmt.Println("Server: Database එකට Request එක යැව්වා (උපරිම තත්පර 2යි දෙන්නේ)...")
	
	go fetchUserFromDB(ctx)

	// Context එක Timeout වෙනකම් (තත්පර 2ක්) බලන් ඉන්නවා
	<-ctx.Done()
	fmt.Println("Server: User ට Error එකක් යැව්වා (504 Gateway Timeout).")
}