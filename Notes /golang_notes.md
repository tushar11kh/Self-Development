# Week 1

### Building a Memory Log File in Go

#### Standard Library Packages Used

| Package | Purpose |
| :-- | :-- |
| `bufio` | Buffered I/O (efficient file reading/writing) |
| `fmt` | Formatted Input/Output |
| `os` | OS-level file operations |
| `strconv` | String to number conversions |
| `strings` | String manipulation functions |
| `time` | Time-related utilities |

#### Key Functions and Concepts

- **os.Open**: Opens a file for reading.
- **defer**: Ensures resources (e.g., file handles) are closed after the function exits, even on error.
- **bufio.NewScanner**: Scans input line by line.
    - `sc.Scan()` reads the next line.
    - `sc.Text()` returns the current line’s text.
- **strings.HasPrefix(line, "MemTotal:")**: Checks if a line starts with `"MemTotal:"`.
- **strings.Fields(line)**: Splits a line into whitespace-separated fields.
- **strconv.Atoi(s)**: Converts a string to an integer.
- **bufio.NewWriter**: Buffered writing to a log file for efficiency.
- **time.Now()**: Gets the current time.
- **.Format("02-01-2006 15:04:05")**: Formats date/time as string.


#### Writing to and Flushing the Log

- `fmt.Sprintf()`: Creates formatted strings.
- `fmt.Print()`: Prints to stdout.
- `writer.WriteString()`: Writes data to the buffer.
- `writer.Flush()`: Ensures buffered data is actually written to the file.


#### Example Time Formatting in Go

```go
fmt.Println(time.Now().Format("02-01-2006 15:04:05"))
```

**Example Program:**
[Memory Usage Collector in Go (GitHub)](https://github.com/tushar11kh/Self-Development/blob/main/week1/mem-usage/main.go)

### Building a CLI for CPU and Memory Usage

- Uses Go's `flag` standard library:
    - `flag.Bool/String/Int(...)`: Define command-line flags.
    - `flag.Parse()`: Parse the command-line arguments before using any flags.

**Example Program:**
[CPU \& Memory Usage CLI Tool (GitHub)](https://github.com/tushar11kh/Self-Development/blob/main/week1/flag_cli/main.go)

### Basic HTTP Requests

Functions and methods demonstrated for HTTP GET:

- **http.Get(url)**: Initiates an HTTP GET request.
- **resp.Body**: The body of the HTTP response. Must be closed after reading.
- **io.ReadAll(resp.Body)**: Reads the entire response into a byte slice.
- **string(body)**: Converts raw bytes to a readable string.

**HTTP-enabled CLI Example:**
[Flag-Based CLI with HTTP (GitHub)](https://github.com/tushar11kh/Self-Development/blob/main/week2/flag_cli/main.go)

# Week 2

### Building and Testing a Simple API

#### Key Packages \& Concepts

- **net/http**
    - `http.HandleFunc()`: Registers HTTP handlers.
    - `http.ListenAndServe()`: Starts the server.
    - `http.Error()`: Sends error responses.
    - `r.URL.Query().Get("id")`: Retrieves query parameters.
- **encoding/json**
    - `json.NewEncoder(w).Encode(v)`: Encodes Go struct as JSON and writes to response.
- **log**
    - `log.Fatal()`: Logs fatal errors (also exits the program).


#### Go Types \& Usage

- **http.ResponseWriter**: Interface for HTTP responses.
- **\*http.Request**: Pointer to the incoming HTTP request.
- **map[string]Person**: Stores data indexed by string keys.
- **struct**: Defines custom data types (e.g., `Person` with `Name` and `Age`).


#### Other Notable Elements

- Struct tags: `json:"name"`, `json:"age"`.
- Set HTTP response header: `w.Header().Set("Content-Type", "application/json")`.
- Common status codes:
    - `http.StatusBadRequest`
    - `http.StatusNotFound`
    - `http.StatusInternalServerError`

**API Example:**
[Sample API in Go (GitHub)](https://github.com/tushar11kh/Self-Development/blob/main/week2/testApi/main.go)

# Go Integration With Docker: Building a Go CLI Tool

## 1. Making HTTP Requests to Docker

Go’s `net/http` is used to interact directly with Docker’s REST API:

```go
resp, err := http.Get("http://localhost:2375/containers/json")
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
```

Use when Docker’s API listens over TCP.

## 2. Controlling Docker: Two Approaches

### A. Using `os/exec` to Run Docker CLI

Use Go to execute shell commands—good for scripting, minimal dependencies.

```go
out, err := exec.Command("docker", "ps").Output()
fmt.Println(string(out))
```


### B. Using Docker SDK for Go

For advanced/robust interaction, leverage Docker’s official Go SDK.

**Install:**

```
go get github.com/docker/docker/client
```

**Sample:**

```go
cli, err := client.NewEnvClient()
containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{})
for _, container := range containers {
    fmt.Println(container.ID, container.Image)
}
```

The SDK enables full container and image management.

## 3. Fetching \& Parsing Docker Container Stats

Obtain and decode resource usage stats via SDK:

```go
resp, err := cli.ContainerStats(ctx, containerID, false)
defer resp.Body.Close()
body, _ := ioutil.ReadAll(resp.Body)
var stats types.StatsJSON
json.Unmarshal(body, &stats)
fmt.Println(stats.MemoryStats.Usage)
```

You can also unmarshal into custom structs or `map[string]interface{}` as needed.

## 4. Starting \& Stopping Containers

- **With SDK:**

```go
cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
cli.ContainerStop(ctx, containerID, nil)
```

- **With `os/exec`:**

```go
exec.Command("docker", "start", containerID).Run()
exec.Command("docker", "stop", containerID).Run()
```


## 5. Managing Dependencies With Go Modules

How to use Go Modules for dependency management:

- Initialize:

```
go mod init yourmodule
```

- Add dependencies:

```
go get module@version
```

- Tidy up:

```
go mod tidy
```


## 6. Committing Code to GitHub

**Standard Workflow:**

1. Initialize git and make first commit:

```bash
git init
git add .
git commit -m "Initial commit"
```

2. Add remote and push code:

```bash
git remote add origin https://github.com/youruser/yourrepo.git
git push -u origin main
```

3. For subsequent changes:

```bash
git add <file>
git commit -m "Describe your change"
git push
```


## Summary Table

| Feature | Preferred Go Approach | Key Package |
| :-- | :-- | :-- |
| HTTP (Docker API) | net/http | net/http |
| Shell Docker CLI | os/exec | os/exec |
| Native Docker Ctrl | Docker SDK (Go client) | github.com/docker/docker/client |
| JSON Parsing | encoding/json | encoding/json, api/types |
| Dependency Mgmt | Go Modules | go.mod, go.sum |
| VCS Integration | Git CLI, GitHub Remote | git |

By combining these Go features, you can build robust, maintainable CLI tools that interact efficiently with Docker and other system resources.

