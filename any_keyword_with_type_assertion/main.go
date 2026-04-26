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
package main
import "fmt"

func processData(data any) {
    // අපි අහනවා "මේ data පෙට්ටිය ඇතුළේ තියෙන්නේ string එකක්ද?" කියලා
    strVal, ok := data.(string)
    
    if ok { // ok == true නම් (ඒ කියන්නේ ඒක ඇත්තටම string එකක් නම්)
        fmt.Println("මේක String එකක්! අකුරු ගාණ:", len(strVal)) // දැන් len() වැඩ කරනවා!
        return
    }

    // string එකක් නෙවෙයි නම්, ඊළඟට බලනවා int එකක්ද කියලා
    intVal, ok := data.(int)
    
    if ok { // ok == true නම්
        fmt.Println("මේක අංකයක්! ඒකට 10ක් එකතු කළාම:", intVal + 10) // දැන් ගණන් හදන්න පුළුවන්!
        return
    }

    fmt.Println("මේක මොකක්ද කියලා මම දන්නේ නැහැ!")
}

func main() {
    processData("GoLang") // Output: මේක String එකක්! අකුරු ගාණ: 6
    processData(50)       // Output: මේක අංකයක්! ඒකට 10ක් එකතු කළාම: 60
    processData(true)     // Output: මේක මොකක්ද කියලා මම දන්නේ නැහැ!
}