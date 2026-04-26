package main
import "fmt"

// 1. INTERFACE එක (රැකියා විස්තරය)
// නීතිය: Shape කෙනෙක් වෙන්න නම් අනිවාර්යයෙන්ම "Area()" කියන ක්‍රියාව තියෙන්න ඕනේ.
type Shape interface {
    Area() float64 
}

// 2. STRUCTS (දත්ත ගබඩා කිරීම) - මෙතන implements කියලා මුකුත් නැහැ!
type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

// 3. METHODS (ක්‍රියාවන් ලබා දීම)
// Rectangle එකට Area() හොයන විදිහ කියලා දෙනවා.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Circle එකට Area() හොයන විදිහ කියලා දෙනවා.
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// 4. INTERFACE TYPE එකක් පාවිච්චි කිරීම (Type Safety)
// මේ Function එක ඉල්ලන්නේ Rectangle වත් Circle වත් නෙවෙයි, "Shape" කෙනෙක්වයි!
func PrintArea(s Shape) {
    fmt.Println("Area is:", s.Area())
}



// interface compostion
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

// Interface Composition
// අර පරණ දෙක එකට එකතු කරලා "ReadWriter" කියලා අලුත් නීති පොතක් හැදුවා!
type ReadWriter interface {
    Reader
    Writer
}




func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circ := Circle{Radius: 7}

    // Duck Typing ක්‍රියාත්මක වෙන අවස්ථාව!
    // Go Compiler එක බලනවා "rect ගාව Area() තියෙනවද? ඔව්! එහෙනම් එයා Shape කෙනෙක්."
    PrintArea(rect) // වැඩ කරනවා!
    PrintArea(circ) // වැඩ කරනවා!

	err := withdrawal(10, 5)
	if err != nil {
		fmt.Println(err)
	}
	


}

func withdrawal(amount float64, balance float64) error {
	if amount > balance {
		return InvalidWithdrawalError{
			Amount:      amount,
			Balance:     balance,
			Description: "Insufficient balance",
		}
	}
	return nil
}

// custom errors
type InvalidWithdrawalError struct {
	Amount      float64
	Balance     float64
	Description string
}

func (e InvalidWithdrawalError) Error() string {
	return fmt.Sprintf("Error: %.2f withdrawal failed. Insufficient balance: %.2f", e.Amount, e.Balance)
}

