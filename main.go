package main

import (
	"bufio"
  "os"
  "fmt"
  "strings"
)
// Print Option writes out a separation 
func print_default(s string) {
  fmt.Println("==========================")
  fmt.Println(s)
  fmt.Println("==========================")
}

// User input retrieves a user input for the main menu
// A user can add, delete, list, and edit existing tasks.
func user_input() string{
  // Scan in user input for selecting an option 
  scanner := bufio.NewReader(os.Stdin)
  fmt.Println("Select an option:")
  fmt.Print("a: Add new task\n")
  fmt.Print("d: Delete a task\n")
  fmt.Print("e: Edit a task\n")
  fmt.Print("l: List tasks\n")
  option, _ := scanner.ReadString('\n')
  option = strings.TrimSpace(option)

  // Set option based on user input
  switch option {
  case "a":
    option = "Add a new task"
  case "d":
    option = "Delete a task"
  case "e":
    option = "List tasks"

  default:
    option = "NoOp"
  }


  return option
}

func main(){
  // Start user input
  option := user_input()
  print_default(option)
}
