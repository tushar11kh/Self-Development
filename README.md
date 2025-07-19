# 13-Week Mixed Roadmap: DevOps, Cloud, Go, and Web3 (8 Hours/Day)

**Goal**: Land a $80K-$120K remote/international job in DevOps or Web3 infrastructure.  
**Hardware**: Ubuntu MacBook (8GB/256GB, EC2-like), M3 MacBook (24GB/512GB).  
**Focus**: Blend Git, Docker, Terraform, Kubernetes, AWS, Go, Ethereum light node, Solana testnet, Polygon RPC, and monitoring weekly to stay engaged, with enhanced Kubernetes (Helm, EKS), Docker security, Terraform (S3 backend), CI/CD (ArgoCD), and system design. No blog reading to ensure relevant learning.  
**Structure**: Each week mixes 2-3 skills with daily topics (Day 1-5, 8 hours), weekend projects, and resources. Each topic has a one-line intro.

---

## Week 1: Linux, Git, and Go Basics
- **Goals**: Master Linux CLI, Git, and Go fundamentals.
- **Day 1** (Ubuntu: Linux, M3: Git):
  - Linux CLI Navigation: Learn `ls`, `cd`, `pwd` to navigate and manage files (2h).
  - File Permissions: Understand `chmod`, `chown` for file access control (2h).
  - Git Setup: Create GitHub repo and push a sample file to start version control (2h).
  - SSH Configuration: Setup passwordless SSH for secure remote access (2h).
- **Day 2** (Ubuntu: Linux, M3: Go):
  - Firewall Setup: Configure `ufw` to secure network ports (2h).
  - Go Installation: Install Go and write a program to print system info (2h).
  - Go Syntax: Learn variables and functions via “A Tour of Go” (go.dev/tour) (2h).
  - System Services: Explore `systemd` for managing background processes (2h).
- **Day 3** (M3):
  - Git Branching: Master `branch` and `pull request` for collaborative workflows (2h).
  - Go Control Flow: Study conditionals and loops in Go (2h).
  - Go Structs: Understand structs for structured data in Go (2h).
  - GitHub Workflow: Practice commits and pull requests on GitHub (2h).
- **Day 4** (Ubuntu: Linux, M3: Go):
  - Bash Scripting: Write a Bash script to monitor CPU usage (2h).
  - Go CLI Development: Build a CLI to log system memory usage (2h).
  - Linux Process Management: Use `ps`, `top` to monitor running processes (2h).
  - Go Error Handling: Implement error handling in Go CLI (2h).
- **Day 5** (Ubuntu: Linux, M3: Git/Go):
  - Git Commits: Commit Go CLI to GitHub with clear messages (2h).
  - Go CLI Flags: Add `--cpu`, `--mem` flags to CLI for flexibility (2h).
  - Systemd Service: Create a systemd service for auto-running CLI (2h).
  - GitHub README: Write a README for the project repo (2h).
- **Weekend Project**: `sys-monitor` – Go CLI + Bash script to log CPU/memory, systemd service, GitHub repo with README.
- **Definition of Done**: CLI runs with flags, systemd service auto-starts, repo has README with setup steps.
- **Resources**: FreeCodeCamp “Linux Basics” (YouTube), GitHub Docs, “A Tour of Go”.
- **Hardware**: Ubuntu MacBook (Linux, systemd), M3 MacBook (Git, Go).
- **Outcome**: Linux CLI proficiency, Git workflow, basic Go CLI.

---

## Week 2: Docker, Go CLI, and AWS Intro
- **Goals**: Learn Docker basics with security, build Go CLI, explore AWS Free Tier.
- **Day 1** (Ubuntu: Docker, M3: Go):
  - Docker Installation: Install Docker to run containerized apps (2h).
  - Docker Multi-Stage Builds: Use multi-stage builds for smaller images (2h).
  - Go HTTP Requests: Add HTTP request functionality to `sys-monitor` CLI (2h).
  - Docker Images: Learn image creation and management with `docker build` (2h).
- **Day 2** (Ubuntu: Docker, M3: AWS):
  - AWS Free Tier: Sign up and explore EC2/S3 in AWS console (2h).
  - Docker Image Hardening: Learn to minimize images with `distroless` (2h).
  - AWS CLI Setup: Install and configure AWS CLI for command-line access (2h).
  - Docker Run: Practice `docker run` with custom options (2h).
- **Day 3** (M3):
  - Go CLI for Docker: Build a CLI to check Docker container status (2h).
  - AWS EC2 Basics: Explore EC2 instance types and security groups (2h).
  - Docker Security Scanning: Scan images with Trivy for vulnerabilities (2h).
  - Go Modules: Use Go modules for dependency management (2h).
- **Day 4** (Ubuntu: Docker, M3: Go):
  - Docker Multi-Container: Run an NGINX container with Docker (2h).
  - Go CLI Enhancements: Add container start/stop functionality to CLI (2h).
  - Docker Secrets: Explore secrets handling in Docker (e.g., environment variables) (2h).
  - Go JSON Parsing: Parse JSON responses in CLI for Docker stats (2h).
- **Day 5** (Ubuntu: Docker, M3: AWS/Go):
  - AWS EC2 Launch: Launch a t3.micro EC2 instance and SSH in (2h).
  - Go CLI Commit: Commit CLI code to GitHub (2h).
  - Docker Hub Push: Push secure Docker image to Docker Hub (2h).
  - AWS Cleanup: Terminate EC2 instance to avoid costs (2h).
- **Weekend Project**: `docker-check` – Go CLI to manage Docker containers, Dockerized with multi-stage build, Trivy-scanned, pushed to GitHub/Docker Hub.
- **Definition of Done**: CLI starts/stops containers, image is secure, repo has README with security details.
- **Resources**: TechWorld with Nana “Docker Tutorial” (YouTube), AWS Docs, Docker Docs, Trivy Docs.
- **Hardware**: Ubuntu MacBook (Docker), M3 MacBook (AWS, Go).
- **Outcome**: Secure Docker basics, Go CLI, AWS console familiarity.

---

## Week 3: Docker Compose, Git CI/CD, and Go HTTP
- **Goals**: Master Docker Compose, setup secure CI/CD, build Go HTTP client.
- **Day 1** (Ubuntu: Compose, M3: Go):
  - Docker Compose Basics: Write `docker-compose.yml` for NGINX + Redis stack (2h).
  - Go HTTP Client: Build an HTTP client for a public API (jsonplaceholder.typicode.com) (2h).
  - Compose Services: Define services and networks in Compose file (2h).
  - Go API Integration: Integrate API calls into existing CLI (2h).
- **Day 2** (M3: Git, Ubuntu: Compose):
  - GitHub Actions Setup: Create a CI/CD pipeline with matrix builds for `docker-check` (2h).
  - Docker Compose Execution: Run NGINX + Redis stack locally (2h).
  - GitHub Workflow Security: Use secrets for secure CI/CD credentials (2h).
  - Compose Volumes: Configure volumes for persistent data (2h).
- **Day 3** (M3):
  - Go Container Stats: Add container stats fetching to `docker-check` via HTTP (2h).
  - GitHub Actions Linting: Add `golangci-lint` and Trivy to CI/CD pipeline (2h).
  - Go Error Handling: Enhance HTTP client with robust error handling (2h).
  - Git Secrets: Setup GitHub Secrets for secure CI/CD credentials (2h).
- **Day 4** (Ubuntu: Compose, M3: Go):
  - Compose Go App: Add Go app to Docker Compose stack (2h).
  - Go Response Parsing: Parse API responses in CLI for structured output (2h).
  - Compose Security: Secure Compose stack with environment variables (2h).
  - Go Testing: Write unit tests for HTTP client functionality (2h).
- **Day 5** (M3: Git/Go, Ubuntu: Compose):
  - GitHub Actions Push: Configure Actions to push secure Docker image (2h).
  - Go Commit: Commit updated CLI to GitHub (2h).
  - Docker Compose Testing: Test full stack functionality locally (2h).
  - GitHub README: Add CI/CD details to repo README (2h).
- **Weekend Project**: `docker-stack` – Compose file + Go HTTP client, secure CI/CD pipeline with Trivy, GitHub repo.
- **Definition of Done**: Stack runs NGINX + Redis + Go app, CI/CD builds/pushes secure image, repo has README with security setup.
- **Resources**: Docker Compose Docs, GitHub Actions Docs, “Go HTTP Clients” (YouTube), Trivy Docs.
- **Hardware**: Ubuntu MacBook (Compose), M3 MacBook (Git, Go).
- **Outcome**: Secure multi-container apps, CI/CD pipeline, Go HTTP skills.

---

## Week 4: AWS EC2, Go AWS SDK, and Linux Networking
- **Goals**: Launch EC2, use Go AWS SDK, configure Linux networking.
- **Day 1** (Ubuntu: Linux, M3: AWS):
  - AWS EC2 Setup: Launch a t3.micro EC2 instance with security groups (2h).
  - Linux Firewall: Configure `ufw` to allow ports 22, 80 (2h).
  - AWS Security Groups: Define inbound/outbound rules for EC2 (2h).
  - Linux File Management: Practice `find`, `grep` for file operations (2h).
- **Day 2** (M3):
  - Go AWS SDK: Use AWS SDK to list EC2 instances in Go (2h).
  - AWS SSH Access: SSH into EC2 instance with key pair (2h).
  - Go SDK Setup: Configure AWS SDK with credentials and region (2h).
  - AWS CLI EC2: Run `aws ec2 describe-instances` to verify setup (2h).
- **Day 3** (Ubuntu: Linux, M3: Go):
  - Linux Security: Install and configure `fail2ban` for intrusion prevention (2h).
  - Go S3 SDK: Build CLI to list S3 buckets using AWS SDK (2h).
  - Linux Networking Tools: Use `ss`, `dig` to inspect network connections (2h).
  - Go Error Handling: Implement error handling for AWS SDK calls (2h).
- **Day 4** (M3):
  - AWS IAM Roles: Create IAM role for EC2-S3 access (2h).
  - Go S3 Integration: Add S3 bucket listing to CLI (2h).
  - AWS CLI S3: Run `aws s3 ls` with IAM role credentials (2h).
  - Go CLI Refinement: Optimize CLI for user-friendly output (2h).
- **Day 5** (Ubuntu: Linux, M3: Go/AWS):
  - Linux Network Diagnostics: Test connectivity with `ping`, `traceroute` (2h).
  - Go Commit: Commit CLI code to GitHub (2h).
  - AWS Cleanup: Terminate EC2 instance to avoid costs (2h).
  - GitHub README: Add AWS CLI usage details to repo README (2h).
- **Weekend Project**: `aws-monitor` – Go CLI to list EC2/S3 resources, GitHub repo with README.
- **Definition of Done**: CLI lists EC2/S3, IAM role works, repo has README with usage.
- **Resources**: AWS EC2 Docs, AWS SDK Go Docs, Ubuntu Server Guide.
- **Hardware**: Ubuntu MacBook (Linux), M3 MacBook (AWS, Go).
- **Outcome**: EC2 management, Go SDK usage, Linux networking.

---

## Week 5: Ethereum Light Node, Docker, and Go
- **Goals**: Run Ethereum light node, secure Dockerized Go app, enhance CLI.
- **Day 1** (Ubuntu: Geth, M3: Go):
  - Ethereum Geth Setup: Install Geth and run light node (`geth --syncmode light`) (2h).
  - Go RPC Query: Write Go code to query Ethereum RPC (`eth_blockNumber`) (2h).
  - Docker Security: Use `distroless` for Go app Dockerfile (2h).
  - Ethereum Sync Modes: Understand light vs. full sync modes (2h).
- **Day 2** (Ubuntu: Docker/Geth, M3: Go):
  - Docker Go App: Write secure Dockerfile with multi-stage build (2h).
  - Ethereum RPC Testing: Test Geth RPC endpoint with `curl` (2h).
  - Docker Trivy Scan: Scan Go app image for vulnerabilities with Trivy (2h).
  - Go RPC Integration: Integrate RPC query into CLI (2h).
- **Day 3** (Ubuntu: Docker, M3: Go):
  - Go Ethereum CLI: Build CLI to log Ethereum block height (2h).
  - Docker Image Build: Build and test secure Docker image (2h).
  - Go Logging: Add file logging to CLI for metrics (2h).
  - Docker Volumes: Configure volumes for persistent CLI data (2h).
- **Day 4** (Ubuntu: Geth, M3: Go):
  - Ethereum Systemd: Setup Geth as a systemd service for auto-start (2h).
  - Go CLI Enhancements: Add flags for customizable RPC queries (2h).
  - Ethereum Node Monitoring: Monitor Geth logs for sync status (2h).
  - Go Error Handling: Handle RPC errors in CLI (2h).
- **Day 5** (Ubuntu: Docker/Geth, M3: Go):
  - Docker Hub Push: Push secure Go app image to Docker Hub (2h).
  - Go Commit: Commit CLI code to GitHub (2h).
  - Ethereum Node Verification: Verify light node sync progress (2h).
  - GitHub README: Add Ethereum CLI setup to repo README (2h).
- **Weekend Project**: `eth-ping` – Secure Dockerized Go CLI to ping Ethereum node, logs to S3, GitHub repo.
- **Definition of Done**: CLI queries block height, Docker image is secure, repo has README with setup and security details.
- **Resources**: Ethereum.org “Running a Node”, Docker Docs, Trivy Docs.
- **Hardware**: Ubuntu MacBook (Geth, Docker), M3 MacBook (Go).
- **Outcome**: Ethereum light node, secure Dockerized Go app.

---

## Week 6: Prometheus, AWS S3, and Polygon RPC
- **Goals**: Setup Prometheus, use S3, explore Polygon RPC.
- **Day 1** (Ubuntu: Prometheus, M3: AWS):
  - Prometheus Installation: Install Prometheus and Node Exporter for metrics (2h).
  - AWS S3 Setup: Create S3 bucket and sync sample logs (2h).
  - Prometheus Scraping: Configure Prometheus to scrape Node Exporter (2h).
  - AWS IAM Basics: Setup IAM user for S3 access (2h).
- **Day 2** (Ubuntu: Prometheus, M3: Go):
  - Go Metrics Tool: Build Go tool to export system metrics to Prometheus (2h).
  - Prometheus Configuration: Setup `prometheus.yml` for custom metrics (2h).
  - Polygon RPC Intro: Query Polygon testnet RPC with `curl` (e.g., block height) (2h).
  - Go Prometheus Client: Use Prometheus Go client library for metrics (2h).
- **Day 3** (M3):
  - AWS S3 IAM Role: Create IAM role for S3 write access (2h).
  - Go S3 Logging: Add S3 logging to Go metrics tool (2h).
  - Polygon RPC Queries: Write Go code to query Polygon RPC endpoint (2h).
  - Go CLI Flags: Add flags for customizable metrics output (2h).
- **Day 4** (Ubuntu: Prometheus, M3: Go):
  - Prometheus PromQL: Query metrics using PromQL basics (2h).
  - Go Custom Metrics: Enhance tool with custom system metrics (2h).
  - Polygon Node Basics: Explore Polygon testnet node setup requirements (2h).
  - Prometheus Alerts: Setup basic alert rules in Prometheus (2h).
- **Day 5** (Ubuntu: Prometheus, M3: Go/AWS):
  - AWS S3 Testing: Verify log sync to S3 bucket (2h).
  - Go Commit: Commit metrics tool to GitHub (2h).
  - Prometheus Dashboard: Build a basic Prometheus dashboard (2h).
  - GitHub README: Add metrics tool setup to repo README (2h).
- **Weekend Project**: `sys-metrics` – Go tool + Prometheus for system metrics, Polygon RPC query, logs in S3, GitHub repo.
- **Definition of Done**: Tool exports metrics, queries Polygon RPC, Prometheus scrapes, S3 stores logs, repo has README with setup.
- **Resources**: Prometheus.io, AWS S3 Docs, Polygon Docs, “Prometheus Go Client” (YouTube).
- **Hardware**: Ubuntu MacBook (Prometheus), M3 MacBook (AWS, Go), AWS Free Tier (Polygon).
- **Outcome**: Monitoring basics, S3 integration, Polygon RPC intro.

---

## Week 7: Grafana, Ethereum Monitoring, and Terraform Intro
- **Goals**: Setup Grafana, monitor Ethereum node, start Terraform, draft blog post.
- **Day 1** (Ubuntu):
  - Grafana Installation: Install Grafana for visualization (2h).
  - Ethereum Log Monitoring: Monitor Geth logs for sync status (2h).
  - Grafana Data Source: Add Prometheus as a Grafana data source (2h).
  - Ethereum RPC Queries: Query Geth RPC for block height (2h).
- **Day 2** (Ubuntu: Grafana, M3: Terraform):
  - Terraform Installation: Install Terraform for infrastructure as code (2h).
  - Grafana Dashboard: Build CPU/memory dashboard in Grafana (2h).
  - Terraform S3: Write Terraform script for S3 bucket (2h).
  - Grafana Panels: Create basic panels for system metrics (2h).
- **Day 3** (Ubuntu: Geth, M3: Go):
  - Ethereum Metrics: Query Geth for metrics (e.g., peer count) (2h).
  - Go Prometheus Exporter: Build Go tool to export Geth metrics (2h).
  - Ethereum Node Optimization: Adjust Geth cache for performance (2h).
  - Go Metrics Integration: Integrate Geth metrics into Prometheus (2h).
- **Day 4** (Ubuntu: Grafana, M3: Terraform):
  - Grafana Ethereum Dashboard: Add Geth metrics to Grafana dashboard (2h).
  - Terraform EC2: Write Terraform script for EC2 instance (2h).
  - Grafana Visualizations: Customize dashboard with graphs and gauges (2h).
  - Terraform Variables: Use variables in Terraform scripts (2h).
- **Day 5** (Ubuntu: Geth/Grafana, M3: Go):
  - Go Commit: Commit Geth metrics tool to GitHub (2h).
  - Ethereum Sync Check: Verify Geth light node sync status (2h).
  - Grafana Finalization: Finalize Ethereum metrics dashboard (2h).
  - GitHub README: Add dashboard setup to repo README (2h).
- **Weekend Project**: `eth-monitor` – Go tool + Grafana dashboard for Ethereum metrics, GitHub repo with blog post draft.
- **Definition of Done**: Dashboard shows Geth metrics, tool exports data, repo has README with screenshots, blog post drafted.
- **Resources**: Grafana Labs, Ethereum.org, HashiCorp “Learn Terraform”.
- **Hardware**: Ubuntu MacBook (Geth, Grafana), M3 MacBook (Terraform, Go).
- **Outcome**: Ethereum monitoring, Terraform basics, blogging skills.

---

## Week 8: Terraform EC2, Go, and Kubernetes Intro
- **Goals**: Provision EC2 with Terraform, enhance Go CLI, start Kubernetes with ArgoCD.
- **Day 1** (M3):
  - Terraform EC2 Setup: Write Terraform script for EC2 and security groups (2h).
  - Go SSH CLI: Build Go CLI to SSH into EC2 instances (2h).
  - Terraform Security Groups: Define inbound rules for EC2 (2h).
  - Go CLI Planning: Plan CLI structure for EC2 interaction (2h).
- **Day 2** (M3):
  - Kubernetes Minikube: Install Minikube for local Kubernetes cluster (2h).
  - Terraform Variables: Add variables to EC2 Terraform script (2h).
  - Kubernetes Pods: Deploy a sample pod in Minikube (2h).
  - Terraform State: Understand Terraform state management (2h).
- **Day 3** (M3):
  - Go Log Tailing: Add log tailing functionality to Go CLI (2h).
  - Kubernetes Deployments: Create a deployment in Minikube (2h).
  - Go SDK Integration: Integrate AWS SDK for EC2 status checks (2h).
  - Kubernetes Services: Explore Service types for exposing pods (2h).
- **Day 4** (M3):
  - Terraform Deployment: Test EC2 deployment with Terraform (2h).
  - Go CLI Enhancements: Add status checks to CLI for EC2 (2h).
  - ArgoCD Intro: Install ArgoCD in Minikube for GitOps (2h).
  - Kubernetes Networking: Configure networking for deployments (2h).
- **Day 5** (M3):
  - Kubernetes Service Exposure: Expose deployment via Service in Minikube (2h).
  - Go Commit: Commit CLI code to GitHub (2h).
  - Terraform Commit: Push Terraform scripts to GitHub (2h).
  - GitHub README: Add EC2 CLI setup to repo README (2h).
- **Weekend Project**: `aws-infra` – Terraform EC2 deployment + Go CLI for log tailing, ArgoCD auto-deployment, GitHub repo.
- **Definition of Done**: Terraform deploys EC2, CLI tails logs, ArgoCD deploys app, repo has README with setup.
- **Resources**: HashiCorp Terraform, Kubernetes.io “Basics”, ArgoCD Docs.
- **Hardware**: M3 MacBook (all tasks, Minikube RAM-intensive).
- **Outcome**: IaC with Terraform, Kubernetes intro, GitOps basics.

---

## Week 9: Kubernetes, Ethereum, and Go Health Checks
- **Goals**: Deploy to Kubernetes with Helm, optimize Ethereum node, build Go health checker.
- **Day 1** (Ubuntu: Geth, M3: Kubernetes):
  - Kubernetes App Deployment: Deploy Go app to Minikube (2h).
  - Ethereum Optimization: Adjust Geth cache for light node performance (2h).
  - Helm Intro: Install Helm and create a basic chart (2h).
  - kubectl Debugging: Use `kubectl describe`, `logs` to debug pods (2h).
- **Day 2** (M3):
  - Go Health Checker: Build Go CLI to check Geth RPC health (2h).
  - Kubernetes Services: Create Service for Go app in Minikube (2h).
  - Helm Chart Development: Convert Go app to a Helm chart (2h).
  - Kubernetes ConfigMaps: Configure app settings with ConfigMaps (2h).
- **Day 3** (Ubuntu: Geth, M3: Go):
  - Ethereum RPC Queries: Query Geth for health metrics (e.g., peers) (2h).
  - Go Health Integration: Add health check logic to CLI (2h).
  - kubectl Exec: Use `kubectl exec` to interact with pods (2h).
  - Go Testing: Write tests for health checker functionality (2h).
- **Day 4** (M3):
  - Kubernetes StatefulSet: Write StatefulSet for Geth with Helm chart (2h).
  - Go Pod Checks: Enhance CLI to check Kubernetes pod health (2h).
  - ArgoCD Deployment: Configure ArgoCD to auto-deploy Helm chart (2h).
  - Kubernetes PVCs: Configure PersistentVolumeClaims for Geth data (2h).
- **Day 5** (Ubuntu: Geth, M3: Kubernetes/Go):
  - Kubernetes Testing: Test StatefulSet and Helm chart in Minikube (2h).
  - Go Commit: Commit health checker CLI to GitHub (2h).
  - Ethereum Verification: Verify Geth node stability (2h).
  - GitHub README: Add Kubernetes setup to repo README (2h).
- **Weekend Project**: `kube-eth` – Kubernetes Helm deployment + Go health checker, ArgoCD auto-deployment, GitHub repo.
- **Definition of Done**: Helm chart deploys Geth, CLI checks health, ArgoCD automates, repo has README with setup.
- **Resources**: Kubernetes.io, Ethereum.org, Helm Docs, ArgoCD Docs.
- **Hardware**: M3 MacBook (Kubernetes, Go), Ubuntu MacBook (Geth).
- **Outcome**: Kubernetes Helm deployment, advanced node ops, GitOps.

---

## Week 10: AWS KMS, Go Secrets, and Solana Intro
- **Goals**: Secure secrets with KMS, build Go secrets CLI, explore Solana testnet.
- **Day 1** (M3):
  - AWS KMS Setup: Store a fake validator key in AWS KMS (2h).
  - Go KMS CLI: Build Go CLI to fetch KMS secrets (2h).
  - AWS Secrets Manager: Explore secrets handling with Secrets Manager (2h).
  - AWS IAM Policies: Create IAM policy for KMS access (2h).
- **Day 2** (M3: KMS, AWS: Solana):
  - Solana CLI Setup: Install Solana CLI on AWS Free Tier for testnet (2h).
  - AWS KMS IAM: Configure IAM role for KMS access (2h).
  - Solana Testnet Node: Run Solana testnet node on AWS (2h).
  - Go KMS Integration: Integrate KMS fetch into CLI (2h).
- **Day 3** (M3: Go, AWS: Solana):
  - Go Encryption: Add encryption functionality to CLI using KMS (2h).
  - Solana Node Queries: Query Solana testnet node for status (2h).
  - Go Secrets Manager: Integrate Secrets Manager into CLI (2h).
  - Solana Monitoring: Monitor Solana node logs for performance (2h).
- **Day 4** (M3):
  - AWS KMS Testing: Test secret retrieval with KMS CLI (2h).
  - Go Decryption: Add decryption functionality to CLI (2h).
  - EKS IAM Roles: Explore IAM roles for EKS service accounts (2h).
  - EKS Load Balancer: Study load balancer integration for EKS (2h).
- **Day 5** (M3: Go/KMS, AWS: Solana):
  - Solana Node Management: Monitor and manage Solana testnet node (2h).
  - Go Commit: Commit KMS CLI to GitHub (2h).
  - EKS Helm Prep: Plan Helm chart for Go + Solana node on EKS (2h).
  - GitHub README: Add KMS and Solana setup to repo README (2h).
- **Weekend Project**: `secure-secrets` – Go CLI for KMS/Secrets Manager-encrypted keys, Solana testnet script, EKS Helm chart prep, GitHub repo.
- **Definition of Done**: CLI encrypts/decrypts, Solana node runs, Helm chart planned, repo has README with setup.
- **Resources**: AWS KMS Docs, Solana Docs, “Go AWS SDK” (YouTube).
- **Hardware**: M3 MacBook (KMS, Go), AWS Free Tier (Solana).
- **Outcome**: Secure key management, Solana intro, EKS prep.

---

## Week 11: Final Project and Terraform Advanced
- **Goals**: Build Web3 toolkit with EKS + Helm, use advanced Terraform, polish portfolio.
- **Day 1** (M3):
  - Terraform EKS Setup: Write Terraform script for AWS EKS cluster (2h).
  - Go EKS CLI: Build Go CLI for EKS deployment management (2h).
  - Terraform S3 Backend: Configure S3 + DynamoDB for state management (2h).
  - Terraform Modules: Organize EKS script into reusable modules (2h).
- **Day 2** (M3):
  - EKS Geth Deployment: Deploy Geth + Go app to EKS with Helm chart (2h).
  - Portfolio README: Write README for `eth-monitor` project (2h).
  - Terraform DRY Patterns: Use modules for reusable EKS code (2h).
  - EKS Networking: Configure EKS networking for Geth (2h).
- **Day 3** (M3):
  - EKS Monitoring: Add Prometheus/Grafana to EKS via Helm (2h).
  - Portfolio Enhancement: Add screenshots to `kube-eth` README (2h).
  - Terraform State Locking: Implement state locking with DynamoDB (2h).
  - Prometheus Operator: Install Prometheus operator in EKS (2h).
- **Day 4** (M3):
  - EKS Helm Testing: Test Geth + Go app Helm deployment in EKS (2h).
  - Portfolio Refinement: Update `aws-infra` README with details (2h).
  - Terraform Workspaces: Explore Terraform workspaces for environments (2h).
  - EKS Load Balancer: Configure load balancer for EKS services (2h).
- **Day 5** (M3):
  - EKS Project Commit: Commit `web3-toolkit` to GitHub (2h).
  - Portfolio Finalization: Finalize all repo READMEs and diagrams (2h).
  - EKS kubectl Debugging: Use `kubectl rollout` to manage deployments (2h).
  - GitHub README: Add EKS setup to `web3-toolkit` README (2h).
- **Weekend Project**: `web3-toolkit` – EKS Helm deployment of Geth + Go app, Prometheus/Grafana, Terraform S3 backend, GitHub repo with blog post.
- **Definition of Done**: EKS runs Geth/app via Helm, dashboard works, Terraform uses S3 backend, repos have READMEs/diagrams, blog post published.
- **Resources**: Terraform EKS Module, Helm Docs, draw.io.
- **Hardware**: M3 MacBook (all tasks).
- **Outcome**: Production-ready EKS deployment, advanced Terraform, polished portfolio.

---

## Week 12: AWS Certification Prep
- **Goals**: Prepare for AWS Certified Solutions Architect – Associate, polish portfolio.
- **Day 1** (M3):
  - AWS EC2 Study: Learn EC2 instance types, AMIs, and scaling (2h).
  - AWS S3 Study: Study S3 storage classes and bucket policies (2h).
  - Portfolio Certification: Add certification prep section to GitHub (2h).
  - AWS VPC Basics: Understand VPC, subnets, and routing (2h).
- **Day 2** (M3):
  - AWS IAM Study: Explore IAM users, roles, and policies (2h).
  - AWS ECS Basics: Learn ECS for container orchestration (2h).
  - Portfolio README: Refine READMEs for all projects (2h).
  - AWS Route 53: Study DNS management with Route 53 (2h).
- **Day 3** (M3):
  - AWS Well-Architected: Study Well-Architected Framework pillars (2h).
  - Portfolio Screenshots: Add screenshots to project READMEs (2h).
  - AWS SQS/SNS: Understand messaging and notification services (2h).
  - AWS Exam Practice: Practice exam questions (Udemy) (2h).
- **Day 4** (M3):
  - AWS Practice Exam: Take full-length practice exam (Whizlabs) (2h).
  - Portfolio Diagrams: Finalize architecture diagrams for repos (2h).
  - AWS Exam Strategy: Learn to dissect exam questions and manage time (2h).
  - AWS RDS Basics: Study relational database service basics (2h).
- **Day 5** (M3):
  - AWS Weak Areas: Review weak areas from practice exam (2h).
  - Portfolio Blog: Publish blog on “Deploying Ethereum Node on EKS” (Medium) (2h).
  - AWS Exam Scheduling: Schedule certification exam (2h).
  - Portfolio Final Check: Ensure all repos are polished (2h).
- **Weekend Project**: Updated portfolio with blog: “Deploying Ethereum Node on EKS”.
- **Definition of Done**: Score 80%+ on practice exams, blog published, portfolio complete.
- **Resources**: Udemy “AWS Solutions Architect” (Stephane Maarek), Whizlabs.
- **Hardware**: M3 MacBook.
- **Outcome**: Ready for AWS certification, polished portfolio with blog.

---

## Week 13: Job Applications, Networking, and System Design
- **Goals**: Apply to jobs, network, finalize portfolio with system design.
- **Day 1** (M3):
  - Job Applications: Apply to 10 DevOps/Web3 jobs on Web3.career, LinkedIn (2h).
  - X Networking: Engage with @ethereum posts on X (2h).
  - System Design Prep: Study CI/CD pipeline design for Ethereum nodes (2h).
  - Resume Tailoring: Customize resume for DevOps roles (2h).
- **Day 2** (M3):
  - Job Applications: Apply to 10 jobs on CryptoJobsList, Turing (2h).
  - LinkedIn Networking: Join Web3 groups on LinkedIn (2h).
  - System Design Questions: Practice “Secrets in Kubernetes” design (2h).
  - Cover Letter Writing: Write cover letters for top 5 jobs (2h).
- **Day 3** (M3):
  - Job Applications: Apply to 5 jobs with tailored cover letters (2h).
  - LinkedIn Blog Sharing: Share blog post on LinkedIn (2h).
  - Mock Interview Prep: Practice STAR method for behavioral questions (2h).
  - System Design Practice: Practice “Node crash monitoring” scenario (2h).
- **Day 4** (M3):
  - Job Follow-Ups: Follow up on previous job applications (2h).
  - X Engagement: Comment on @awscloud posts on X (2h).
  - System Design Practice: Practice “Recovering failed Terraform apply” scenario (2h).
  - Portfolio Video Prep: Plan 5-minute `web3-toolkit` demo video (2h).
- **Day 5** (M3):
  - Job Applications: Apply to 5 jobs with custom cover letters (2h).
  - Recruiter Outreach: Connect with 10 recruiters on LinkedIn (2h).
  - Portfolio Video: Upload 5-minute `web3-toolkit` demo video to GitHub/YouTube (2h).
  - System Design Doc: Draft doc for “Ethereum CI/CD” design (2h).
- **Weekend Project**: Final portfolio, 5-minute video demo of `web3-toolkit`, system design doc for “Ethereum CI/CD”.
- **Definition of Done**: 30+ applications submitted, video demo uploaded, 10+ connections made, system design doc completed.
- **Resources**: Web3.career, CryptoJobsList, STAR Method for Resumes, “System Design Interview” (YouTube).
- **Hardware**: M3 MacBook.
- **Outcome**: Active job applications, strong online presence, system design readiness.

---

## Outcomes
- **Skills**: Git, Docker (secure), Terraform (advanced), Kubernetes (Helm, EKS), AWS, Go, Ethereum light node, Solana testnet, Polygon RPC, Prometheus/Grafana, ArgoCD.
- **Portfolio**: 7 GitHub projects (`sys-monitor`, `docker-check`, `docker-stack`, `aws-monitor`, `eth-ping`, `kube-eth`, `web3-toolkit`), blog posts, system design doc.
- **Certification**: AWS Solutions Architect – Associate ($150).

