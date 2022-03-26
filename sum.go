package main

import (
    "fmt"
)

// func Err(){
//     return "Sorry Try again"
//     main()
// }
// func Substract (){
//
// }


func WannaPlayAgain(){

    var getAns string

    fmt.Println("Hello , Wanna Play Again with Us (Yes/No)")
    fmt.Scan(&getAns)

    if (getAns == "yes")||(getAns == "y"){
        fmt.Println("Nice to have you Again")
        Sum()
    }else {
        fmt.Println("It seems you dont wanna play \nBye!!")
    }

}

func WannaPlay(){

    var getAns string

    fmt.Println("Hello We are going to make game \n Do you want to get the number \n(yes or no) ")
    fmt.Scan(&getAns)

    if (getAns == "yes")||(getAns == "y"){
        fmt.Println("Nice to have you")
        Sum()
    }else {
        fmt.Println("It seems you dont wanna play")
    }

}

func Sum(){
    var one,two int
    fmt.Println("type two number")
    fmt.Scan(&one,&two)
    result := one + two
    fmt.Println("Your Result is :",result,"total")
    WannaPlayAgain()

}

func main(){
//     var one, two int
    WannaPlay()

}
