/// CLI todo list

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	myMap := map[int]string{}
	taskID := 1

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("This is a CLI TODO LIST")

	for {
		fmt.Println("1. Add a task")
		fmt.Println("2. Display task")
		fmt.Println("3. Deleta a task")
		fmt.Println("4. Exit")

		fmt.Println("enter your choice")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input enter numbers from the given options")
			continue
		}
		switch option {
		case 1:
			fmt.Print("Enter your task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			myMap[taskID] = task
			fmt.Printf("Task added with ID %d\n", taskID)
			taskID++
		case 2:
			if len(myMap) == 0 {
				fmt.Println("no task available")
			} else {
				for key, value := range myMap {
					fmt.Println(key, value)
				}
			}
		case 3:
			fmt.Print("Enter task ID to delete: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)
			id, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Invalid task ID.")
				continue
			}
			if _, exists := myMap[id]; exists {
				delete(myMap, id)
				fmt.Println("Task deleted.")
			} else {
				fmt.Println("Task ID not found.")
			}
		case 4:
			return

		}

	}
}
