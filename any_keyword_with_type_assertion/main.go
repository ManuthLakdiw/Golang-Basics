// package main
// import "fmt"

// // මෙතන data කියන variable එකට ඕනෑම දෙයක් භාරගන්න පුළුවන්
// func printSecretBox(data any) {
//     fmt.Println("පෙට්ටිය ඇතුළේ තියෙන්නේ:", data)
// }

// func main() {
//     printSecretBox("Hello") // String එකක් යවනවා
//     printSecretBox(100)     // int එකක් යවනවා
//     printSecretBox(3.14)    // float එකක් යවනවා
// }


// type assertion
// package main
// import "fmt"

// func processData(data any) {
//     // අපි අහනවා "මේ data පෙට්ටිය ඇතුළේ තියෙන්නේ string එකක්ද?" කියලා
//     strVal, ok := data.(string)
    
//     if ok { // ok == true නම් (ඒ කියන්නේ ඒක ඇත්තටම string එකක් නම්)
//         fmt.Println("මේක String එකක්! අකුරු ගාණ:", len(strVal)) // දැන් len() වැඩ කරනවා!
//         return
//     }

//     // string එකක් නෙවෙයි නම්, ඊළඟට බලනවා int එකක්ද කියලා
//     intVal, ok := data.(int)
    
//     if ok { // ok == true නම්
//         fmt.Println("මේක අංකයක්! ඒකට 10ක් එකතු කළාම:", intVal + 10) // දැන් ගණන් හදන්න පුළුවන්!
//         return
//     }

//     fmt.Println("මේක මොකක්ද කියලා මම දන්නේ නැහැ!")
// }

// func main() {
//     processData("GoLang") // Output: මේක String එකක්! අකුරු ගාණ: 6
//     processData(50)       // Output: මේක අංකයක්! ඒකට 10ක් එකතු කළාම: 60
//     processData(true)     // Output: මේක මොකක්ද කියලා මම දන්නේ නැහැ!
// }


//type switch
package main
import "fmt"

func checkTypeWithSwitch(data any) {
    // මෙතන v කියන variable එකට හැම case එකේදීම අදාළ Type එකේ Value එක එනවා
    switch v := data.(type) {
    case string:
        // මේ block එක ඇතුළේ v කියන්නේ අනිවාර්යයෙන්ම string එකක්
        fmt.Printf("String එකක් හම්බුණා: %s (දිග: %d)\n", v, len(v))
    case int:
        // මේ block එක ඇතුළේ v කියන්නේ අනිවාර්යයෙන්ම int එකක්
        fmt.Printf("අංකයක් හම්බුණා. දෙකෙන් වැඩි කළාම: %d\n", v*2)
    case bool:
        fmt.Println("Boolean එකක් හම්බුණා:", v)
    default:
        fmt.Println("දන්නෙ නැති Data Type එකක්!")
    }
}

func main() {
    checkTypeWithSwitch("Go is Awesome")
    checkTypeWithSwitch(25)
    checkTypeWithSwitch(3.14)
}