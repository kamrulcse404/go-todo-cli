# Todo CLI

A simple and beginner-friendly Todo CLI application built with **Go**. This project demonstrates how to build a command-line application using Go while following a clean project structure and modular design.

---

## Features

- ➕ Add a new todo
-  List all todos
-  Edit an existing todo
-  Mark a todo as completed
-  Delete a todo
-  Search a todo by ID
-  List all pending todos
-  Display todo statistics
-  Persistent storage using JSON
-  Input validation
-  Error handling
-  Modular project structure

---

## Project Structure

```
todo-cli/
│── commands/
│── models/
│── storage/
│── util/
│── public/
│── main.go
│── go.mod
│── README.md
```

---

## Concepts Covered

This project covers many fundamental Go concepts, including:

### Go Basics

- Variables
- Constants
- Functions
- Packages
- Imports
- Control Statements
- Loops
- Switch Statements

### Data Structures

- Structs
- Slices
- JSON Encoding/Decoding

### Error Handling

- Returning errors
- Custom errors
- Error propagation

### File Handling

- Reading JSON files
- Writing JSON files
- Persistent storage

### Project Organization

- Modular packages
- Separation of concerns
- Helper functions
- Clean code practices

### Standard Library

- `fmt`
- `os`
- `encoding/json`
- `errors`
- `strings`

---

## Installation

### Requirements

- Go **1.24+** (or the version specified in `go.mod`)

Check your Go version:

```bash
go version
```

Clone the repository:

```bash
git clone https://github.com/kamrulcse404/todo-cli.git
```

Go to the project directory:

```bash
cd todo-cli
```

Run the application:

```bash
go run .
```


---

##  Usage

### Add Todo

```bash
todo add "Learn Golang"
```

---

### List Todos

```bash
todo list
```

---

### Complete Todo

```bash
todo complete 1
```

---

### Edit Todo

```bash
todo edit 1 "Learn Go Deeply"
```

---

### Delete Todo

```bash
todo delete 1
```

---

### Search Todo

```bash
todo search 1
```

---

### List Pending Todos

```bash
todo pending
```

---

### Todo Statistics

```bash
todo stats
```

---

## Data Storage

Todos are stored in a JSON file.

Example:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "completed": false
  }
]
```

---

## Technologies Used

- Go
- JSON
- Standard Library

---

## Learning Goals

This project was built to practice:

- Building CLI applications
- File handling
- JSON manipulation
- Structs and slices
- Package organization
- Error handling
- Clean code
- Go project structure