# Hands-On: Practice \& Do for Docker Compose, Go HTTP Clients, and GitHub CI/CD

## 1. Docker Compose

### Composition Fundamentals

- Write a minimal `docker-compose.yml` describing two services (e.g., NGINX and Redis).
- Validate your YAML file using `docker compose config` to catch indentation or syntax errors.


### Multi-Service Stack: NGINX and Redis

- Use Compose to spin up the NGINX and Redis services you defined.
- Map NGINX’s port 80 to a host port (e.g., 8080:80) and verify access via browser or curl.
- Use Compose service names (`nginx`, `redis`) to test inter-service communication (e.g., through a custom NGINX config or by inspecting logs).


### Compose Services and Networks

- Add a custom user-defined network to your Compose file.
- Run `docker network inspect` to observe network isolation and service connectivity.
- Attempt to connect services not on the same network and verify they cannot reach each other.


### Persistent Data and Volumes

- Add a named volume for Redis data:

```yaml
volumes:
  redis-data:
services:
  redis:
    ...
    volumes:
      - redis-data:/data
```

- Insert test data into Redis, stop/remove the container, and restart to confirm persistence.


### Compose File Best Practices

- Move environment variables into a `.env` file and reference them in your `docker-compose.yml`.
- Implement secret injection using Compose (or Docker secrets if in Swarm mode) for sensitive config, and verify secrets do not appear in logs or with `docker inspect`.


### Running and Managing Compose Stacks

- Start your stack with `docker compose up -d`, stop it with `docker compose down`, and practice restarting single containers.
- Use `docker compose logs` and `docker compose ps` to monitor your services.
- Remove stopped containers and validate clean-up of volumes/networks as needed.


## 2. Go HTTP Client Development

### Go HTTP Client Fundamentals

- Write a CLI program in Go that fetches data from `https://jsonplaceholder.typicode.com/posts/1` and prints the result.
- Compile and run the program from the command line.


### Request Handling

- Extend your client to send headers (e.g., set a fake `Authorization` header).
- Add context-based timeouts, aborting requests that take too long.


### Response Handling and Parsing

- Parse returned JSON into Go structs and print a user-friendly summary to the terminal.
- Handle error responses: detect HTTP status codes outside the 2xx range and display meaningful error messages.


### CLI Integration

- Refactor your program for flag support (e.g., `-endpoint` to select API resource).
- Modularize code to separate CLI logic from HTTP request/response logic.


### Testing Go HTTP Clients

- Write simple tests for your API functions using table-driven style.
- Use Go’s `httptest` package to mock the API, simulating success and error cases.


## 3. GitHub Actions CI/CD Pipeline

### Workflow Fundamentals

- Create a `.github/workflows/main.yml` that triggers on every push and PR to your repository.
- Define steps to check out the code, install dependencies, and run `go test`.


### Matrix Builds

- Configure a matrix job to run your build/test step on at least two Go versions and two OSes (e.g., `ubuntu-latest` and `windows-latest`).


### Build, Lint, and Security Steps

- Add `golangci-lint` as a step to check for code linting warnings.
- Integrate Trivy to scan your Docker images for vulnerabilities, and address any high/critical issues flagged.


### Secure CI/CD Credentials

- Add a dummy secret (e.g., `DOCKERHUB_PASSWORD`) to your GitHub repo's Secrets.
- Reference this secret in your workflow and use it to log in to Docker Hub during a build step. Ensure the secret isn’t leaked in any logs.


### Workflow Security Best Practices

- Explicitly set minimal permissions for workflow jobs (e.g., `permissions: read-all` or more restrictive as needed).
- Review your workflow for any places secrets might be accidentally exposed (e.g., echoed to output).


## 4. Docker Compose Integration with Go Apps

### Adding Go App to Compose Stack

- Add a new service to your `docker-compose.yml` that builds your Go app from a Dockerfile.
- Ensure all services (Go app, Redis, NGINX) run together, and validate by hitting the app's endpoint.


### Dockerfile Security and Hardening

- Refactor your Dockerfile to use multi-stage builds, resulting in a small, production-ready image.
- Set your app to run as a non-root user and verify by checking container processes.


### Service Integration

- Update your Go app to connect to Redis using the Compose network alias.
- Store and retrieve sample data in Redis from your Go service, verifying cross-service networking.


### Docker Compose Stack Testing

- Use `docker compose logs`, `docker compose ps`, and built-in healthchecks to debug service failures.
- Simulate service failures (e.g., force one to exit) to troubleshoot dependency and restart behavior.


## 5. Robust Go HTTP Client Enhancements

### Error Handling Strategies

- Implement custom error types in your Go client for network errors, HTTP status errors, and timeout errors. Print detailed diagnostics for each case.
- Add a retry mechanism for failed HTTP requests, respecting a maximum retry count and implementing an exponential backoff strategy.


### Fetching Container Stats via HTTP

- Extend your CLI to interact with Docker’s Engine API (e.g., `GET /containers/json` or `/containers/{id}/stats`), either via local Unix socket or configured API endpoint.
- Display core stats (CPU, mem, status) in a clear format.


### Parsing and Validating API Responses

- Add validation layers to ensure data types and values returned from APIs are as expected. Log warnings or errors for unexpected formats or missing fields.
- Include structured logging for all errors and edge cases.


## 6. Git and GitHub Best Practices

### Commit and Branching Workflow

- Make all changes in clearly-named feature branches.
- Commit logically related changes as atomic units with descriptive messages.
- Regularly open pull requests and use review tools to address merge conflicts.


### CI/CD Integration in Git Workflow

- Enable required status checks on your GitHub branch protection rules.
- Only merge pull requests after automated builds/tests pass.


## 7. Documentation and Project Readiness

### Comprehensive README

- Add setup, build, run, and usage sections to your README.
- Include example commands, sample output, CI/CD status badges, and troubleshooting tips.


### Compliance with Definition of Done

- Maintain a project checklist (requirements, security, documentation, coverage).
- Only mark work items complete after they pass this checklist and end-to-end testing.


## 8. Review \& Practice

### End-to-End Testing

- Deploy your full Compose stack, run through every user workflow from initial deployment to shutdown.
- Confirm end-to-end functionality: service communication, secret handling, volume persistence, error reporting, and pipeline security.


### Peer Review (Optional)

- Request reviews from teammates; respond to and incorporate feedback via PR comments.
- Document notes on lessons learned and improvements implemented from review cycles.



