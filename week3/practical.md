You're on **Week 3** of your roadmap, which blends:

* **Docker Compose** (infra + orchestration),
* **Go** (HTTP clients + CLI),
* **GitHub Actions** (CI/CD with security),
* and a **weekend project** that ties it all together.

Below is a **precise breakdown of what to do each day**, which machine to use, and the deep understanding you should aim for in each task.

---

## ✅ OVERALL WEEK GOAL

* Run a **Docker Compose stack** with NGINX + Redis + Go App.
* Build a **Go HTTP client** for public API + container stats.
* Set up a **secure CI/CD pipeline** on GitHub using GitHub Actions + Secrets + Trivy.
* Finish a weekend project called `docker-stack`.

---

## 🗓️ **Day 1: Compose + Go HTTP Client**

**Machine Use:**

* **Ubuntu** → Docker Compose
* **Mac M3** → Go HTTP client

### 1. Docker Compose Basics (2h)

✅ Write a full `docker-compose.yml` for:

* `nginx` (static file server)
* `redis` (data store)

🔎 Understand:

* Services
* Networks
* Port mappings
* Volumes
* `depends_on`

💡 **Already done above**, just build + run it with:

```bash
docker compose up -d
```

---

### 2. Go HTTP Client (2h)

✅ Write a Go program that:

* Makes a GET request to `https://jsonplaceholder.typicode.com/posts/1`
* Prints response in formatted output

🧠 Learn:

* `net/http`
* `io.ReadAll`
* `encoding/json`
* Basic struct mapping

---

### 3. Compose Services + Networks (2h)

✅ Modify your Compose file to:

* Explicitly define a `backend` network
* Add service names and make sure they're reachable internally (e.g., nginx can talk to redis)

🧠 Learn:

* Docker service discovery
* How containers in the same Compose network communicate via DNS

---

### 4. Go API Integration (2h)

✅ Integrate your API logic into your existing **Go CLI** (from Week 2).

* Add a command like `cli fetch-post` to retrieve and print a post

🧠 Learn:

* Command flag handling (if using Cobra or `flag`)
* Modular Go project layout

---

## 🗓️ **Day 2: GitHub Actions CI/CD + Compose Volumes**

**Machine Use:**

* **Ubuntu** → Compose execution
* **Mac M3** → GitHub workflows

### 1. GitHub Actions Setup (2h)

✅ Create `.github/workflows/ci.yml`:

* Use **matrix strategy** to test multiple Go versions
* Steps: checkout, setup Go, run build/test

🧠 Learn:

* Matrix builds
* YAML syntax in Actions

---

### 2. Run Compose Stack Locally (2h)

✅ Run your full stack on Ubuntu:

```bash
docker compose up -d
```

✅ Verify:

* NGINX is up (`curl localhost:8080`)
* Redis responds (`docker exec redis redis-cli ping`)

---

### 3. GitHub Workflow Security (2h)

✅ Add **GitHub Secrets**:

* `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN`
* Reference them in Actions for secure Docker login & push

🧠 Learn:

* How secrets are injected (`${{ secrets.VARIABLE }}`)

---

### 4. Compose Volumes (2h)

✅ Add volume to Redis service:

```yaml
volumes:
  redis_data:
```

And mount it:

```yaml
volumes:
  - redis_data:/data
```

🧠 Learn:

* Persistent volumes vs ephemeral containers

---

## 🗓️ **Day 3: Stats, Linting, Secrets**

**Machine Use: Mac M3**

### 1. Go Container Stats via HTTP (2h)

✅ Extend your Go CLI:

* Query `docker stats` via `http` (e.g., expose container stats via REST API in Go)

🧠 Learn:

* Use Docker client SDK in Go (optional)
* Or write HTTP server that runs `docker stats` and returns JSON

---

### 2. GitHub Actions Linting (2h)

✅ Add:

* [`golangci-lint`](https://golangci-lint.run)
* [`aquasecurity/trivy`](https://github.com/aquasecurity/trivy-action) for image scanning

🧠 Learn:

* Linting & static analysis
* Container vulnerability scanning

---

### 3. Go Error Handling (2h)

✅ Improve your Go client:

* Handle non-200 responses
* Handle bad JSON
* Use `log.Fatal` or `errors.Wrap` for clear debug info

---

### 4. GitHub Secrets for Docker Push (2h)

✅ Confirm DockerHub login is using GitHub Secrets in your workflow:

```yaml
- name: Login to DockerHub
  run: echo "${{ secrets.DOCKERHUB_TOKEN }}" | docker login -u "${{ secrets.DOCKERHUB_USERNAME }}" --password-stdin
```

---

## 🗓️ **Day 4: Compose Go App + Testing**

**Machine Use: Ubuntu for Compose, Mac M3 for Go**

### 1. Add Go App to Compose (2h)

✅ Add a Go API server to your Compose stack:

* Build Go binary
* Create a `Dockerfile`
* Add `go_app` service to Compose file

---

### 2. Parse API Response in CLI (2h)

✅ Parse full JSON object (posts, comments, etc.) and display formatted output

* Use structs with `json` tags

---

### 3. Secure Compose with Env Vars (2h)

✅ Use `.env` file or `env_file:` in Compose

* Store Redis password, app config securely

🧠 Learn:

* Difference between `.env` and `environment:` in Compose

---

### 4. Unit Testing in Go (2h)

✅ Write tests for your HTTP client

* Use `net/http/httptest`
* Test happy & error paths

---

## 🗓️ **Day 5: Git + CI/CD Push + README**

**Machine Use: Mac for Git/Go, Ubuntu for Compose**

### 1. GitHub Actions: Push Docker Image (2h)

✅ Add GitHub Action step:

* Build Go app
* Dockerize it
* Push to DockerHub securely

---

### 2. Git Commit Updated CLI (2h)

✅ Stage, commit, and push changes:

```bash
git add .
git commit -m "Add Go HTTP client and Compose stack"
git push origin main
```

---

### 3. Test Full Stack (2h)

✅ Run your full Compose stack:

```bash
docker compose up --build
```

✅ Test:

* Can Go app connect to Redis?
* Is NGINX accessible?
* Is API client working?

---

### 4. Add README Details (2h)

✅ Include in `README.md`:

* Project overview
* Compose services
* CI/CD pipeline steps
* Security (Secrets, Trivy)
* How to run/test it

---

## 🏁 Weekend Project: `docker-stack`

Your GitHub project should include:

* `docker-compose.yml` with:

  * NGINX
  * Redis
  * Go App (Dockerized)
* `Dockerfile` for Go App
* GitHub Actions CI/CD

  * Linting (golangci-lint)
  * Security (Trivy)
  * Build & Push Docker image
* Go CLI with:

  * HTTP client for API
  * Container stats endpoint (or call)
* ✅ **README with full setup instructions + security setup**

---

## 🔚 Definition of Done

You are done when:

✅ Stack runs with:

* NGINX
* Redis
* Go App

✅ CI/CD:

* Lints Go
* Builds/pushes Docker image
* Scans with Trivy

✅ GitHub repo:

* Code organized
* CI/CD workflows secure
* README explains usage + security

---

If you'd like, I can now:

1. Write your `.github/workflows/ci.yml` file.
2. Build Go HTTP client code.
3. Help create Dockerfile for your Go app.
4. Guide Compose security practices in `.env`.

Let me know which part you want done next.
