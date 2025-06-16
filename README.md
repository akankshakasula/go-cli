CLI To-Do List in Go

This is a simple command-line To-Do List application built using the Go programming language. It allows users to manage their daily tasks directly from the terminal. The application supports adding tasks, displaying the list of tasks, deleting tasks using task IDs, and exiting the program.

Features

Add tasks with automatically generated task IDs

View all current tasks

Delete tasks using their respective task ID

Menu-driven interface that runs in a continuous loop until the user exits

How to Run

Make sure Go is installed on your computer

Save the code in a file named main.go

Open your terminal and navigate to the folder containing main.go

Run the command
go run main.go

How It Works

The program displays a menu with four options

The user is prompted to enter their choice

Based on the selected option, the program either adds a new task, displays all tasks, deletes a task by ID, or exits

Tasks are stored in a map with integer keys for task IDs and string values for task descriptions

Concepts Used

Maps for storing tasks

Loops and switch statements for menu control

bufio for reading input from the user

strconv for converting string input to integers

Error handling for invalid inputs

Example

When the user selects option 1, they can add a task
When the user selects option 2, all current tasks are displayed
When the user selects option 3, they are prompted to enter a task ID to delete
Option 4 exits the application

This project is a basic example of how to use Go for building interactive command-line applications.
