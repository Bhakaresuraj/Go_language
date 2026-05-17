# Go Language Learning Repository

A hands-on collection of Go programs built to learn the language step by step — from basic syntax all the way to real-world DevOps tooling.

---

## 🧭 What Is This Repo?

This repo is a **personal Go learning journal**. Each folder is a self-contained program focused on one concept. The goal is to progress from Go fundamentals to building production-style applications through practical, runnable examples.

---

## 🚀 Getting Started

**Prerequisites:** [Go 1.18+](https://go.dev/dl/) installed.

```bash
git clone https://github.com/Bhakaresuraj/Go_language.git
cd Go_language

# Navigate into any module and run it
cd <module-name>
go run main.go
```

---

## 📈 Learning Path — Basic to Advanced

The modules are best explored in this order:

| # | Module | What You Learn |
|---|---|---|
| 1 | `function/` | Functions, multiple return values, error handling |
| 2 | `defer/` | The `defer` keyword and LIFO execution order |
| 3 | `package/` | Standard library packages (`os`, `strings`, `strconv`) and third-party packages |
| 4 | `http-server/` | Building a basic HTTP server, reading request data |
| 5 | `marshalling/` | JSON marshalling & unmarshalling with structs |
| 6 | `interfaces/` | Interface declaration, method receivers, polymorphism |
| 7 | `interfaces-2/` | Multi-package interface design (UPI payment system) |
| 8 | `goRoutines/` | Goroutines and `sync.WaitGroup` for concurrency |
| 9 | `channels/` | Channels for goroutine communication |
| 10 | `main.go` | Combines HTTP server + JSON + business logic into a login API |
| 11 | `devops-healthcheck/` | Multi-package real-world app — service health checker with color output |

---

## 🛠️ Troubleshooting

### `go: command not found`
Go is not installed or not in your PATH.
```bash
# Verify installation
go version

# If missing, download from https://go.dev/dl/
# Then add to PATH (Linux/macOS):
export PATH=$PATH:/usr/local/go/bin
```

### `no required module provides package ...`
Dependencies are missing. Run this inside the module folder:
```bash
go mod tidy
```

### `cannot find module providing package ...`
You're likely running `go run` from the wrong directory. Make sure you're inside the folder that has a `go.mod` file:
```bash
cd devops-healthcheck   # not from the repo root
go run main.go
```

### `address already in use` (port 8080)
Another process is already using port 8080.
```bash
# Find and kill the process using the port
lsof -i :8080
kill -9 <PID>

# Or just use a different port in the code
http.ListenAndServe(":9090", nil)
```

### `undefined: <PackageName>`
You haven't imported the package, or the import path is wrong. Check your `import` block and confirm the package name matches what's in `go.mod`.

### `json: cannot unmarshal ...`
The JSON keys in your request don't match the struct field names or tags. Double-check your struct tags:
```go
type Student struct {
    Name string `json:"name"`  // must match the key in the JSON body
}
```

### Third-party package not found (e.g. `fatih/color`)
```bash
go get github.com/fatih/color
go mod tidy
```

### Goroutine exits before finishing (`channels/`, `goRoutines/`)
If you launch goroutines but the program exits immediately, you're missing synchronization. Use `sync.WaitGroup` or a channel to wait for results before `main()` returns.

---

## 👤 Author

**Suraj Bhakare** · [@Bhakaresuraj](https://github.com/Bhakaresuraj)
