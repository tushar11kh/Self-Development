# Day 2

## 🔐 Docker Image Hardening: Minimizing with Distroless

### 🔧 What is Docker Image Hardening?

**Docker image hardening** is the practice of reducing the attack surface of container images by:

* Removing unnecessary files, libraries, and tools
* Reducing image size
* Following security principles like least privilege

### 🎯 Why Minimize Docker Images?

* Smaller **attack surface** (fewer tools = fewer potential exploits)
* Faster **pull times** (smaller images)
* Easier **auditing and vulnerability scanning**
* Better **performance and reliability**



## 📦 What are Distroless Images?

**Distroless images**, created by Google, are minimal Docker images that **don’t include a package manager or shell** like bash.

> They only include the **application and its runtime** dependencies—nothing more.

### ✅ Key Properties:

| Feature               | Traditional Image      | Distroless       |
| --------------------- | ---------------------- | ---------------- |
| Shell (e.g., bash)    | ✅ Present              | ❌ Not present    |
| Package manager (apt) | ✅ Present              | ❌ Not present    |
| Debug tools           | ✅ Yes (curl, ps, etc.) | ❌ No tools       |
| Attack surface        | ⚠️ Large               | ✅ Minimal        |
| Size                  | ⚠️ Large (\~100MB+)    | ✅ Small (5–20MB) |



## 🧪 Example Comparison

Let’s compare a Node.js app using:

* Ubuntu-based image
* Alpine
* Distroless

### 🏗️ Dockerfile with Ubuntu:

```Dockerfile
FROM node:18
WORKDIR /app
COPY . .
RUN npm install
CMD ["node", "index.js"]
```

Size: \~130MB+



### 🪶 Dockerfile with Alpine:

```Dockerfile
FROM node:18-alpine
WORKDIR /app
COPY . .
RUN npm install
CMD ["node", "index.js"]
```

Size: \~80MB

Better, but still includes tools like `sh`, `apk`, `wget`, etc.



### 🔒 Dockerfile with Distroless:

```Dockerfile
# Build stage
FROM node:18 AS builder
WORKDIR /app
COPY . .
RUN npm install --omit=dev

# Final image
FROM gcr.io/distroless/nodejs18
COPY --from=builder /app /app
WORKDIR /app
CMD ["index.js"]
```

✅ No shell, no package manager
✅ Only includes runtime and your app
✅ Size: \~20MB



## 🛡️ Why is Distroless More Secure?

| Security Feature     | Benefit                                    |
| -------------------- | ------------------------------------------ |
| No shell (`/bin/sh`) | Attackers can’t execute arbitrary commands |
| No package manager   | Can’t install malware or backdoors         |
| Read-only by default | Reduces runtime tampering                  |
| Minimal libraries    | Fewer CVEs (Common Vulnerabilities)        |



## 🔍 Downsides of Distroless

1. ❌ **No shell access** for debugging

   * Solution: Use ephemeral debug images like `busybox` or `alpine` mounted to the same volume

2. ❌ **Harder to inspect manually**

   * Solution: Use multi-stage builds and CI-based security scanning (e.g., `Trivy`, `grype`, `Dockle`)

3. 🧰 **Only supports certain runtimes**

   * Distroless images exist for: `node`, `go`, `java`, `python`, `cc`, `static`, etc.



## 🧰 Tools for Image Hardening Beyond Distroless

| Tool     | Purpose                                       |
| -------- | --------------------------------------------- |
| Trivy    | Vulnerability scanning                        |
| Grype    | Vulnerability and license scanning            |
| Dive     | Explore image layers                          |
| Dockle   | Linter for best practices and hardening       |
| BuildKit | Efficient builds with caching and secrets     |
| Slim.ai  | Automatic image minimization and optimization |



## 📚 Best Practices

* ✅ Use **multi-stage builds**: keep dev tools in the build stage only
* ✅ Prefer **distroless** or **alpine** for base images
* ✅ Avoid using `latest` tag — pin exact image versions
* ✅ Use `USER` in Dockerfile to drop root privileges
* ✅ Scan images regularly with tools like `Trivy`
* ✅ Use **read-only volumes** and **no new privileges**


## 📌 Summary

| Principle               | How Distroless Helps                             |
| ----------------------- | ------------------------------------------------ |
| Minimize attack surface | No shell, no package manager                     |
| Reduce size             | Only runtime + app                               |
| Secure by default       | No debugging tools, locked-down                  |
| Production readiness    | Used at Google and in real-world CI/CD pipelines |
