package main

import (
	"bufio"
  "os"
  "fmt"
  "strings"
)

// Define a new Task structure
type Task struct{
  Name string
  Description string
  Status string 
}

// Print Option writes out a separation 
func print_default(s string) {
  fmt.Println("==========================")
  fmt.Println(s)
  fmt.Println("==========================")
}

func create_task() []Task {
  // Init tasks
  var tasks []Task
  var q string

  // Loop until user quits
  for q != "quit" {
    var newTask Task
    scanner := bufio.NewReader(os.Stdin)
    fmt.Println("Create a new task")
    newTask.Name,_ = scanner.ReadString('\n')
    fmt.Println("Enter a description")
    newTask.Description,_ = scanner.ReadString('\n')
    fmt.Println("Enter a status")
    newTask.Status,_ = scanner.ReadString('\n')
    
    //Add to the list of tasks
    tasks = append(tasks, newTask)

    //Ask user if they want to continue
    fmt.Println("Successfully created task! New task? y/n")
    q, _ = scanner.ReadString('\n') 
    q = strings.TrimSpace(q)

    if q == "y" || q == "Y" {
      q = ""
    } else {
      return tasks
    }
  }
  return tasks
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

  var tasks []Task

  // Set option based on user input
  switch option {
  case "a":
    option = "Add a new task"
    tasks = create_task()
  case "d":
    option = "Delete a task"
  case "e":
    option = "List tasks"

  default:
    option = "NoOp"
  }

  if tasks != nil {
    fmt.Println("Here is teh current list of tasks")
    for _,t := range tasks {
      print_default(t.Name)
    }
  }
  return option
}

func main(){
  // Start user input
  option := user_input()
  print_default(option)
}
