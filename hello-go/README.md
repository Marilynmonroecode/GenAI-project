# 👋 Hello VS Code - Simple Go Project

This is a beginner-friendly Go project that prints "Hello, VS Code!" to the terminal and includes a unit test to demonstrate basic testing in Go.

---

## 📦 Project Structure

HELLO-GO/
├── main.go # Main application file
├── main_test.go # Unit test file for SayHello function
└── go.mod # Go module definition
|__ README.md

---

## 🛠️ Requirements

- [Go (Golang)](https://go.dev/doc/install) installed (v1.21+ recommended)
- [Visual Studio Code (VS Code)](https://code.visualstudio.com/)
- Go extension for VS Code (from Microsoft)
- Git (optional, for version control)

---

## 🚀 How to Run the Project

1. **Clone the repository** (or create your own folder):

   ```bash
   git clone https://github.com/your-username/GenAI-project.git
   cd GenAI-project

2. **Initialize a Go module** (if not already initialized)

   ```bash
   go mod init genai-project

3. **Run the program**

   ```bash
   go run main.go

   You should see:

   ```css
   Hello, VS Code!

---

## 🧪 Running Tests

This project includes a unit test for the SayHello() function.

   bash
   go test

You should see output like:

   PASS
   ok      genai-project    0.002s

---

## 🧠 How It Works

✅ main.go

   package main

   import "fmt"

   // SayHello returns a simple greeting message
   func SayHello() string {
       return "Hello, VS Code!"
   }

   func main() {
       fmt.Println(SayHello())
   }

- The SayHello function returns a string and is tested in main_test.go.
- main() calls the function and prints its return value.


✅ main_test.go
   package main

   import "testing"

   func TestSayHello(t *testing.T) {
       got := SayHello()
       want := "Hello, VS Code!"

       if got != want {
           t.Errorf("SayHello() = %q, want %q", got, want)
       }
   }

- This test checks whether the SayHello() function returns the expected greeting.
- If the result doesn't match, the test will fail with a descriptive message.   

---

## 📚 What You’ll Learn
By working through this project, you’ll learn:
- How to write and run a basic Go program
- How to use Go modules (go mod)
- How to write unit tests in Go
- How to structure a Go project
- How to use VS Code effectively for Go development

---
## MIT License

Copyright (c) 2025 Marilyn Nduku
 







