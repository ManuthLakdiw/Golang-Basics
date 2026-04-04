package main

import "fmt"

func main() {
    var num1 int = 10;
    var num2 float64 = 4.3;

    sum := float64(num1) + num2;
    fmt.Println("Sum of num1 and num2 is: ", sum);

    var num3 int = 10;
    var num4 float64 = 4.3;

    diff := float64(num3) - num4;
    fmt.Println("Difference of num3 and num4 is: ", diff);

    var num5 int = 10;
    var num6 float64 = 4.3;

    product := float64(num5) * num6;
    fmt.Println("Product of num5 and num6 is: ", product);

    var num7 int = 10;
    var num8 float64 = 4.3;

    quotient := float64(num7) / num8;
    fmt.Println("Quotient of num7 and num8 is: ", quotient);
}


