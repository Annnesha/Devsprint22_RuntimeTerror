package main

import(
  "fmt"  
  "main/Database"
)

func main(){
  database.Connect()
  fmt.Println("Successfully Connected")
}


