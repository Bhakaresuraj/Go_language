# Go Language Learning Repository

A hands-on collection of Go programs covering core language concepts — from basic functions and interfaces to goroutines, channels, HTTP servers, and a real-world DevOps health-checker.

---

## 📁 Repository Structure

```
Go_language/
├── main.go                  # HTTP server with login + JSON marshal/unmarshal
├── function/                # Functions, multiple return values & error handling
├── defer/                   # defer keyword & LIFO execution order
├── interfaces/              # Interface basics with an HTTP health checker
├── interfaces-2/            # Advanced interfaces with a UPI payment system
├── goRoutines/              # Goroutines + WaitGroup for concurrent HTTP checks
├── channels/                # Channels for goroutine communication
├── http-server/             # Basic HTTP server with request inspection
├── marshalling/             # JSON marshalling & unmarshalling with structs
├── package/                 # Standard library packages (os, strings, strconv, color)
└── devops-healthcheck/      # Real-world DevOps service health-checker
    ├── models/              # Service struct & port validation
    └── checker/             # Color-coded status printer
```

---

## 📚 Modules

### 1. `function/` — Functions & Error Handling

Demonstrates how to write Go functions with multiple return values and idiomatic error handling using a port validator.

```go
func checkPort(port int) (string, error) {
    if port <= 0 || port >= 65535 {
        return "", fmt.Errorf("This is invalid port : %d", port)
    }
    return fmt.Sprintf("port %d is valid port.", port), nil
}
```

**Concepts covered:** function declaration, multiple return values, error handling with `if err != nil`.

---

### 2. `defer/` — Defer & LIFO Execution

Shows how `defer` works in Go — deferred calls are executed in **Last-In-First-Out (LIFO)** order after the surrounding function returns.

```go
defer fmt.Println("def 8")
defer fmt.Println("def 7")
// ...
// Output order: def 1 → def 2 → ... → def 8
```

**Concepts covered:** `defer` keyword, LIFO stack behavior.

---

### 3. `interfaces/` — Interface Basics

Defines a `Checker` interface and implements it with `HttpService`. Iterates over a slice of `Checker` values to check service health.

```go
type Checker interface {
    Check()
}

type HttpService struct {
    Name    string
    URL     string
    Healthy bool
}

func (h HttpService) Check() { ... }
```

**Concepts covered:** interface declaration, method receivers, polymorphism with slices.

---

### 4. `interfaces-2/` — Advanced Interfaces with Packages

Demonstrates interface-driven design with a **payment system**. The `PaymentMethod` interface is defined in the `payments` package and implemented by `UPIPayment` in the `upi` package.

```go
// payments/payment.go
type PaymentMethod interface {
    Pay(amount float64) string
}

// upi/upi.go
type UPIPayment struct {
    UpiId string
    App   string
}

func (u UPIPayment) Pay(amount float64) string { ... }
```

Usage in `main.go`:
```go
shubhamUpi := upi.UPIPayment{UpiId: "suraj@oksbi", App: "Gpay"}
msg := Checkout(shubhamUpi, 24.3)
```

**Concepts covered:** interface implementation across packages, modular design, dependency injection via interfaces.

---

### 5. `goRoutines/` — Goroutines & WaitGroup

Launches **9 concurrent HTTP health checks** using goroutines, synchronized with `sync.WaitGroup`. Measures total execution time.

```go
var wg sync.WaitGroup

wg.Add(1)
go checkService("api1", "https://httpbin.org/status/200", &wg)
// ...
wg.Wait()

fmt.Println("total time :", time.Since(startTime))
```

**Concepts covered:** `go` keyword, `sync.WaitGroup`, concurrent HTTP requests, execution time measurement.

---

### 6. `channels/` — Channels

Uses a **buffered channel** to receive a result from a goroutine that performs an HTTP GET request.

```go
ch := make(chan string)

go func() {
    resp, err := http.Get("https://httpbin.org/status/200")
    if err != nil {
        ch <- "❌ failed ....."
    }
    defer resp.Body.Close()
    ch <- fmt.Sprintf("✅status: %d", resp.StatusCode)
}()

result := <-ch
fmt.Println(result)
```

**Concepts covered:** unbuffered channels, goroutine-to-main communication, anonymous functions.

---

### 7. `http-server/` — Basic HTTP Server

A minimal HTTP server that logs full request details — method, path, query, headers, content length, host, and body.

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    fmt.Println("Path:", r.URL.Path)
    body, err := io.ReadAll(r.Body)
    // ...
    fmt.Fprintln(w, "Hello from http-server......")
}

http.HandleFunc("/", homeHandler)
http.ListenAndServe(":8080", nil)
```

**Concepts covered:** `net/http`, `http.HandleFunc`, `http.ResponseWriter`, `http.Request`, reading request body.

---

### 8. `marshalling/` — JSON Marshalling & Unmarshalling

Demonstrates reading a JSON request body and deserializing it into a Go struct using `json.Unmarshal`.

```go
type Student struct {
    Name  string `json:"name"`
    Marks int    `json:"marks"`
    City  string `json:"city"`
}

// Unmarshal request body
json.Unmarshal(body, &stud)
```

**Concepts covered:** struct tags, `encoding/json`, `json.Unmarshal`, HTTP request body parsing.

---

### 9. `package/` — Standard Library Packages

Explores several important standard library and third-party packages:

| Package | Usage |
|---|---|
| `fmt` | Formatted output & string formatting |
| `os` | Reading environment variables (`os.Getenv`) |
| `strings` | String manipulation (`strings.ToUpper`) |
| `strconv` | Type conversion (`strconv.Atoi`) |
| `github.com/fatih/color` | Colored terminal output |

**Concepts covered:** environment variables, type conversions, string manipulation, third-party packages.

---

### 10. `devops-healthcheck/` — Real-World DevOps Health Checker

A production-style application organized into separate packages. Checks the health status of multiple services and prints color-coded output.

**Package structure:**

- **`models/`** — defines the `Service` struct and port validation logic
- **`checker/`** — contains `PrintStatus()` which prints green for healthy, red for unhealthy services

```go
// models/service.go
type Service struct {
    Name    string
    Port    int
    Healthy bool
}

// checker/checker.go
func PrintStatus(s models.Service) {
    if !s.Healthy {
        color.Red(fmt.Sprintf("Name:%s | Port:%d | UnHealthy", s.Name, s.Port))
    } else {
        color.Green(fmt.Sprintf("Name:%s | Port:%d | Healthy", s.Name, s.Port))
    }
}
```

**Sample output:**
```
✅ Name:gateway  | Port:8080 | Healthy
❌ Name:postgres | Port:5432 | UnHealthy
✅ Name:frontend | Port:443  | Healthy
```

**Concepts covered:** multi-package project structure, struct constructors, port validation, color-coded CLI output.

---

### 11. `main.go` — HTTP Login Server with JSON

A complete HTTP server that accepts a JSON login request, validates credentials, and responds with a JSON result.

```go
type UserInfo struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Location string `json:"location"`
}

// POST / with JSON body {"username":"admin","password":"admin"}
// Returns: {"english":"90","maths":"80","total":170}
```

**Concepts covered:** JSON unmarshal from request body, business logic in handlers, JSON marshal for response, HTTP status codes.

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.18+](https://go.dev/dl/) installed

### Run any module

```bash
# Clone the repository
git clone https://github.com/Bhakaresuraj/Go_language.git
cd Go_language

# Run a specific module (example: goRoutines)
cd goRoutines
go run main.go
```

### Run the DevOps health checker

```bash
cd devops-healthcheck
go run main.go
```

### Run the HTTP server

```bash
cd http-server
go run main.go
# Server starts on http://localhost:8080
```

Test it with curl:
```bash
curl -X POST http://localhost:8080/ \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}'
```

---

## 🧠 Concepts Summary

| Concept | Module |
|---|---|
| Functions & multiple return values | `function/` |
| Error handling | `function/`, `main.go` |
| Defer & LIFO | `defer/` |
| Interfaces & polymorphism | `interfaces/`, `interfaces-2/` |
| Goroutines | `goRoutines/` |
| Channels | `channels/` |
| WaitGroup | `goRoutines/` |
| HTTP server | `http-server/`, `main.go` |
| JSON marshal / unmarshal | `marshalling/`, `main.go` |
| Packages & modules | `package/`, `interfaces-2/`, `devops-healthcheck/` |
| Environment variables | `package/` |
| Struct tags | `marshalling/`, `main.go` |
| Multi-package project structure | `devops-healthcheck/`, `interfaces-2/` |

---

## 👤 Author

**Suraj Bhakare**  
GitHub: [@Bhakaresuraj](https://github.com/Bhakaresuraj)
