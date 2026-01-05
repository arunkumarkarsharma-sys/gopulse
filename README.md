**  GoPulse — Concurrent URL Status Checker (Go CLI)

GoPulse is a fast and lightweight CLI tool built with Go that checks the status of multiple websites concurrently.  
It reads URLs from a file, sends HTTP requests in parallel, and displays response status and timing.


** Features

-  Concurrent URL checking using **Goroutines**
-  Reads URLs from a text file
-  Shows HTTP status and response time
-  Clean modular project structure
-  Simple and beginner-friendly CLI tool

---

** Project Structure

gopulse/
│
├── cmd/
│ └── main.go # CLI entry point
│
├── checker/
│ └── checker.go # URL checking & concurrency logic
│
├── input/
│ └── reader.go # Reads URLs from file
│
├── urls.txt # Input file with website URLs
├── go.mod # Go module file
└── README.md





** How It Works

1. URLs are read from `urls.txt`
2. Each URL is checked in a separate goroutine
3. Results are sent via channels
4. CLI prints status and response time

