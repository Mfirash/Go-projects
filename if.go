package main
import "fmt"

func main() {
  var i string
  fmt.Scanln(&i)
  fmt.Println("rock paper sissor")
  for {
    if i == "rock" {
      fmt.Println("paper")
      
   } else if i == "paper" {
      fmt.Println("sissor")
      
   } else if i == "sissor" {
      fmt.Println("rock")
      
   } else {
      fmt.Println("what")
      continue
   }
  }
  
}