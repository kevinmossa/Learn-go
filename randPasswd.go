package main

import(
  "fmt"
  "math/rand"
  "time"
)

func main(){
  abc := [52]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z", "A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
  rand.Seed(time.Now().UnixNano())
  var password string = ""
  fmt.Println("Generating password!")
  for i:= 0; i <= 7; i++ {
    randomLetter := rand.Intn(52)
    password = password + abc[randomLetter]
  } 
  fmt.Printf("Your password is: %s", password)
}
