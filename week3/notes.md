# Day 1

## Write docker-compose.yml for NGINX + Redis stack

We want to use **Docker Compose** to launch two services:

* **NGINX** – A high-performance reverse proxy / web server.
* **Redis** – An in-memory key-value store, commonly used as a cache or message broker.

They will be containerized and networked together using Docker Compose.

---

### 🧠 Deep Conceptual Understanding

#### 🧱 Why Docker Compose?

* Docker Compose allows you to **define and run multi-container Docker applications** using a YAML file.
* It handles **networking**, **volume mounts**, and **environment config** between services.

---

### ⚙️ Services Breakdown

#### 1. **NGINX**

* Lightweight web server, reverse proxy, load balancer.
* Can serve static files, or reverse proxy to upstream services (like app servers).

#### 2. **Redis**

* An in-memory database supporting strings, hashes, lists, sets, etc.
* Used for:

  * Caching
  * Pub/Sub messaging
  * Session management

---

### 📄 `docker-compose.yml` File

---

### 🔍 Explained in Detail

#### ✅ `version: '3.9'`

* Specifies the Compose file format.
* `3.9` is stable and widely supported.

#### ✅ `services`

Defines all containers in your stack.

##### nginx:

* `image: nginx:latest` → Pulls the latest official NGINX image.
* `ports: 8080:80` → Redirects traffic from your machine’s port 8080 to container’s port 80 (NGINX default).
* `volumes`:

  * `./nginx.conf:/etc/nginx/nginx.conf:ro` → Mount your custom config.
  * `./html:/usr/share/nginx/html:ro` → Serve static files (like `index.html`).
* `depends_on: redis` → Ensures Redis starts before NGINX.
* `networks: backend` → Both containers share a private Docker network.

###### 🧠 What This Line Means

```yaml
- ./nginx.conf:/etc/nginx/nginx.conf:ro
```

**Split into 3 parts:**

| Part                    | Meaning                                                          |
| ----------------------- | ---------------------------------------------------------------- |
| `./nginx.conf`          | Path on your **host machine** (relative to `docker-compose.yml`) |
| `/etc/nginx/nginx.conf` | Path inside the **NGINX container** (replaces default config)    |
| `:ro`                   | Mount is **read-only** (for safety, NGINX can't modify it)       |

---

##### redis:

* `image: redis:7-alpine` → Alpine-based lightweight Redis.
* `ports: 6379:6379` → Exposes Redis to the host (not always needed unless accessed externally).
* `volumes: redis_data:/data` → Redis persistence.




### 📁 Project Structure

```
nginx-redis-stack/
├── docker-compose.yml
├── nginx.conf
├── html/
│   └── index.html
```

---

### ✍️ Sample `nginx.conf`

This config makes NGINX serve static files and optionally connect to Redis (via a Lua module or upstream app):



---

### 🧪 Running It

```bash
docker compose up -d
```

#### Test:

* Visit `http://localhost:8080` → Should serve `index.html`
* Run `docker exec -it redis_cache redis-cli` → Use the Redis CLI inside container

---

### 🧼 Cleaning Up

```bash
docker compose down -v
```

This removes containers **and** volumes (Redis data).

---

### 🧠 Extra Concepts You Should Understand

| Concept                     | Why It Matters                                                           |
| --------------------------- | ------------------------------------------------------------------------ |
| **Volumes**                 | Keeps Redis data persistent outside container lifecycle                  |
| **Networks**                | Containers can talk to each other using service names (`redis`, `nginx`) |
| **depends\_on**             | Helps order startup but doesn’t wait for Redis to be fully *ready*       |
| **Stateless vs Persistent** | NGINX is stateless, Redis stores data (persistence is optional)          |
| **Configuration Mounts**    | Allows you to control behavior without rebuilding the image              |

---

### 🧭 What You Can Do Next

1. **Add a backend** like a Go or Node.js app that reads/writes to Redis.
2. Use **Lua in NGINX** to interact directly with Redis (advanced).
3. Add **healthchecks** and proper **logging/monitoring**.
4. Setup a **Docker network bridge** for better service isolation.
5. Use **Redis password** and secure the stack.

---

### 🧪 Bonus: How Redis & NGINX Might Interact?

Direct interaction is **not common** unless you:

* Use **OpenResty** (NGINX with Lua scripting)
* Or use NGINX to reverse proxy to an app that uses Redis

So typically the stack looks like this:

```
Client → NGINX (proxy) → App Server (e.g., Go) ↔ Redis
```

---

## Go HTTP Client: Build an HTTP client for a public API (jsonplaceholder.typicode.com)

Perfect. Let’s break this down **from first principles** so you deeply understand how to build a real-world Go HTTP client.

---

## ✅ **Goal**

You will build a Go program that:

* Sends an HTTP **GET** request to a public API
* Receives a **JSON** response
* **Parses** the response into a Go `struct`
* **Prints** the parsed data in a human-readable format

---

## 📍 The API You're Calling

The public API endpoint is:

```
https://jsonplaceholder.typicode.com/posts/1
```

This simulates a RESTful backend. It returns:

```json
{
  "userId": 1,
  "id": 1,
  "title": "sunt aut facere...",
  "body": "quia et suscipit..."
}
```

You need to fetch this and work with it **programmatically**.

---

## 🧠 **Concepts You’ll Learn**

---

### 1. `net/http`

> Go’s standard package for building HTTP clients and servers.

* You’ll use `http.Get()` to make the request.
* This returns an `*http.Response`.

---

### 2. `io.ReadAll`

> Used to read the HTTP response body (which is a stream of bytes).

```go
bodyBytes, err := io.ReadAll(resp.Body)
```

You must **close** the body with `defer resp.Body.Close()` to avoid resource leaks.

---

### 3. `encoding/json`

> Go’s package for encoding/decoding JSON.

* You’ll define a Go `struct` with JSON tags.
* Use `json.Unmarshal()` to decode JSON into that struct.

---

### 4. Struct Mapping (`Unmarshalling`)

You define a struct that looks like the JSON:

```go
type Post struct {
  UserID int    `json:"userId"`
  ID     int    `json:"id"`
  Title  string `json:"title"`
  Body   string `json:"body"`
}
```

The field names in Go **don’t have to match JSON exactly** — the tags (`json:"..."`) handle that.

---

## 🛠️ **Step-by-Step Guide to Build This**

---

### ✅ Step 1: Set Up Your Project

Create a file:

```
go-http-client/
└── main.go
```

---

### ✅ Step 2: Write Full Go Code 

### ✅ Step 3: Run It

```bash
cd go-http-client
go run main.go
```

Expected output:

```
Post Information
----------------
UserID: 1
Post ID: 1
Title: sunt aut facere repellat provident occaecati excepturi optio reprehenderit
Body : quia et suscipit...
```

---

## 🧠 Why This Is Important in DevOps / Backend / Web3

| Skill          | Why it matters                                                          |
| -------------- | ----------------------------------------------------------------------- |
| HTTP Client    | Needed to interact with REST APIs, microservices, backend communication |
| JSON Decoding  | All APIs use JSON; you must decode and work with responses              |
| Error handling | Production-grade Go requires strong, readable error handling            |
| Struct tags    | Crucial for working with APIs, databases (e.g., GORM, DynamoDB, etc.)   |

---

## 🔁 Extension Ideas

If you want to go deeper:

| Task                  | Description                              |
| --------------------- | ---------------------------------------- |
| Use `http.NewRequest` | Set custom headers (e.g., Auth token)    |
| Parse list of posts   | Use `/posts` and unmarshal into `[]Post` |
| Send POST request     | Use `http.Post` and `json.Marshal()`     |
| Add CLI input         | Take post ID as command-line arg         |
| Add unit tests        | Test JSON decoding and error handling    |

---

Would you like to extend this into:

* a CLI that takes post ID as input?
* or a Go program that POSTs data to the API?

Let me know how deep you want to go.



