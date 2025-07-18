# Self-Development

Hey there! This is my journey to getting into Devops,Cloud infrastructure,web3.


# Roadmap:

**Goal**: Land a remote/international job in DevOps or Web3 infrastructure.  
**Hardware**: Ubuntu MacBook (8GB/256GB, EC2-like), M3 MacBook (24GB/512GB).  
**Focus**: Blend Git, Docker, Terraform, Kubernetes, AWS, Go, Ethereum light node, Solana testnet, and monitoring weekly to stay engaged.  
**Structure**: Each week mixes 2-3 skills with daily tasks (Day 1-5), projects, and resources.

---

## Week 1: Linux, Git, and Go Basics
- **Goals**: Master Linux CLI, Git, and Go fundamentals.
- **Day 1** (Ubuntu: Linux, M3: Git):
  - Linux CLI Navigation: Learn commands (`ls`, `cd`, `pwd`) to navigate and manage files.
  - File Permissions: Understand `chmod`, `chown` for file access control.
  - Git Setup: Create GitHub repo and push a sample file to start version control.
  - SSH Configuration: Setup passwordless SSH for secure remote access.
- **Day 2** (Ubuntu: Linux, M3: Go):
  - Firewall Setup: Configure `ufw` to secure network ports.
  - Go Installation: Install Go and write a program to print system info.
  - Go Syntax: Learn variables and functions via “A Tour of Go” (go.dev/tour).
  - System Services: Explore `systemd` for managing background processes.
- **Day 3** (M3):
  - Git Branching: Master `branch` and `pull request` for collaborative workflows.
  - Go Control Flow: Study conditionals and loops in Go.
  - Go Structs: Understand structs for structured data in Go.
  - GitHub Workflow: Practice commits and pull requests on GitHub.
- **Day 4** (Ubuntu: Linux, M3: Go):
  - Bash Scripting: Write a Bash script to monitor CPU usage.
  - Go CLI Development: Build a CLI to log system memory usage.
  - Linux Process Management: Use `ps`, `top` to monitor running processes.
  - Go Error Handling: Implement error handling in Go CLI.
- **Day 5** (Ubuntu: Linux, M3: Git/Go):
  - Git Commits: Commit Go CLI to GitHub with clear messages.
  - Go CLI Flags: Add `--cpu`, `--mem` flags to CLI for flexibility.
  - Systemd Service: Create a systemd service for auto-running CLI.
  - GitHub README: Write a README for the project repo.
- **Weekend Project**: `sys-monitor` – Go CLI + Bash script to log CPU/memory, systemd service, GitHub repo.
- **Definition of Done**: CLI runs with flags, systemd service auto-starts, repo has README with setup steps.
- **Resources**: FreeCodeCamp “Linux Basics” (YouTube), GitHub Docs, “A Tour of Go”.
- **Hardware**: Ubuntu MacBook (Linux, systemd), M3 MacBook (Git, Go).
- **Outcome**: Linux CLI proficiency, Git workflow, basic Go CLI.

---

## Week 2: Docker, Go CLI, and AWS Intro
- **Goals**: Learn Docker basics, build Go CLI, explore AWS Free Tier.
- **Day 1** (Ubuntu: Docker, M3: Go):
  - Docker Installation: Install Docker to run containerized apps.
  - Docker Containers: Run `hello-world` container to understand basics.
  - Go HTTP Requests: Add HTTP request functionality to `sys-monitor` CLI.
  - Docker Images: Learn image creation and management with `docker build`.
- **Day 2** (Ubuntu: Docker, M3: AWS):
  - AWS Free Tier: Sign up and explore EC2/S3 in AWS console.
  - Dockerfiles: Write a Dockerfile for a Go app.
  - AWS CLI Setup: Install and configure AWS CLI for command-line access.
  - Docker Run: Practice `docker run` with custom options.
- **Day 3** (M3):
  - Go CLI for Docker: Build a CLI to check Docker container status.
  - AWS EC2 Basics: Explore EC2 instance types and security groups.
  - Go Modules: Use Go modules for dependency management.
  - AWS S3 Commands: Run `aws s3 ls` to list buckets.
- **Day 4** (Ubuntu: Docker, M3: Go):
  - Docker Multi-Container: Run an NGINX container with Docker.
  - Go CLI Enhancements: Add container start/stop functionality to CLI.
  - Docker Networking: Understand Docker bridge and host networks.
  - Go JSON Parsing: Parse JSON responses in CLI for Docker stats.
- **Day 5** (Ubuntu: Docker, M3: AWS/Go):
  - AWS EC2 Launch: Launch a t3.micro EC2 instance and SSH in.
  - Go CLI Commit: Commit CLI code to GitHub.
  - Docker Hub: Push Docker image to Docker Hub.
  - AWS Cleanup: Terminate EC2 instance to avoid costs.
- **Weekend Project**: `docker-check` – Go CLI to manage Docker containers, Dockerized, pushed to GitHub/Docker Hub.
- **Definition of Done**: CLI starts/stops containers, Dockerfile builds image, repo has README with usage.
- **Resources**: TechWorld with Nana “Docker Tutorial” (YouTube), AWS Docs, Docker Docs.
- **Hardware**: Ubuntu MacBook (Docker), M3 MacBook (AWS, Go).
- **Outcome**: Docker basics, Go CLI, AWS console familiarity.

---

## Week 3: Docker Compose, Git CI/CD, and Go HTTP
- **Goals**: Master Docker Compose, setup CI/CD, build Go HTTP client.
- **Day 1** (Ubuntu: Compose, M3: Go):
  - Docker Compose Basics: Write `docker-compose.yml` for NGINX + Redis stack.
  - Go HTTP Client: Build an HTTP client for a public API (jsonplaceholder.typicode.com).
  - Compose Services: Define services and networks in Compose file.
  - Go API Integration: Integrate API calls into existing CLI.
- **Day 2** (M3: Git, Ubuntu: Compose):
  - GitHub Actions Setup: Create a CI/CD pipeline for `docker-check` to build/test.
  - Docker Compose Execution: Run NGINX + Redis stack locally.
  - GitHub Workflows: Understand workflow syntax (`on`, `jobs`, `steps`).
  - Compose Volumes: Configure volumes for persistent data.
- **Day 3** (M3):
  - Go Container Stats: Add container stats fetching to `docker-check` via HTTP.
  - GitHub Actions Linting: Add `golangci-lint` to CI/CD pipeline.
  - Go Error Handling: Enhance HTTP client with robust error handling.
  - Git Secrets: Setup GitHub Secrets for secure CI/CD credentials.
- **Day 4** (Ubuntu: Compose, M3: Go):
  - Compose Go App: Add Go app to Docker Compose stack.
  - Go Response Parsing: Parse API responses in CLI for structured output.
  - Compose Networking: Configure custom networks in Compose file.
  - Go Testing: Write unit tests for HTTP client functionality.
- **Day 5** (M3: Git/Go, Ubuntu: Compose):
  - GitHub Actions Push: Configure Actions to push Docker image to registry.
  - Go Commit: Commit updated CLI to GitHub.
  - Docker Compose Testing: Test full stack functionality locally.
  - GitHub README Update: Add CI/CD details to repo README.
- **Weekend Project**: `docker-stack` – Compose file + Go HTTP client, CI/CD pipeline, GitHub repo.
- **Definition of Done**: Stack runs NGINX + Redis + Go app, CI/CD builds/pushes image, repo has README with setup.
- **Resources**: Docker Compose Docs, GitHub Actions Docs, “Go HTTP Clients” (YouTube).
- **Hardware**: Ubuntu MacBook (Compose), M3 MacBook (Git, Go).
- **Outcome**: Multi-container apps, CI/CD pipeline, Go HTTP skills.

---

## Week 4: AWS EC2, Go AWS SDK, and Linux Networking
- **Goals**: Launch EC2, use Go AWS SDK, configure Linux networking.
- **Day 1** (Ubuntu: Linux, M3: AWS):
  - AWS EC2 Setup: Launch a t3.micro EC2 instance with security groups.
  - Linux Firewall: Configure `ufw` to allow ports 22, 80.
  - AWS Security Groups: Define inbound/outbound rules for EC2.
  - Linux File Management: Practice `find`, `grep` for file operations.
- **Day 2** (M3):
  - Go AWS SDK: Use AWS SDK to list EC2 instances in Go.
  - AWS SSH Access: SSH into EC2 instance with key pair.
  - Go SDK Setup: Configure AWS SDK with credentials and region.
  - AWS CLI EC2: Run `aws ec2 describe-instances` to verify setup.
- **Day 3** (Ubuntu: Linux, M3: Go):
  - Linux Security: Install and configure `fail2ban` for intrusion prevention.
  - Go S3 SDK: Build CLI to list S3 buckets using AWS SDK.
  - Linux Networking Tools: Use `ss`, `dig` to inspect network connections.
  - Go Error Handling: Implement error handling for AWS SDK calls.
- **Day 4** (M3):
  - AWS IAM Roles: Create IAM role for EC2-S3 access.
  - Go S3 Integration: Add S3 bucket listing to CLI.
  - AWS CLI S3: Run `aws s3 ls` with IAM role credentials.
  - Go CLI Refinement: Optimize CLI for user-friendly output.
- **Day 5** (Ubuntu: Linux, M3: Go/AWS):
  - Linux Network Diagnostics: Test network connectivity with `ping`, `traceroute`.
  - Go Commit: Commit CLI code to GitHub.
  - AWS Cleanup: Terminate EC2 instance to avoid costs.
  - GitHub README: Add AWS CLI usage details to repo README.
- **Weekend Project**: `aws-monitor` – Go CLI to list EC2/S3 resources, GitHub repo with README.
- **Definition of Done**: CLI lists EC2/S3, IAM role works, repo has README with usage.
- **Resources**: AWS EC2 Docs, AWS SDK Go Docs, Ubuntu Server Guide.
- **Hardware**: Ubuntu MacBook (Linux), M3 MacBook (AWS, Go).
- **Outcome**: EC2 management, Go SDK usage, Linux networking.

---

## Week 5: Ethereum Light Node, Docker, and Go
- **Goals**: Run Ethereum light node, Dockerize Go app, enhance CLI.
- **Day 1** (Ubuntu: Geth, M3: Go):
  - Ethereum Geth Setup: Install Geth and run light node (`geth --syncmode light`).
  - Go RPC Query: Write Go code to query Ethereum RPC (`eth_blockNumber`).
  - Ethereum Sync Modes: Understand light vs. full sync modes.
  - Go CLI Structure: Plan CLI structure for Ethereum node interaction.
- **Day 2** (Ubuntu: Docker/Geth, M3: Go):
  - Docker Go App: Write Dockerfile for existing Go CLI.
  - Ethereum RPC Testing: Test Geth RPC endpoint with `curl`.
  - Docker Multi-Stage Builds: Optimize Dockerfile with multi-stage builds.
  - Go RPC Integration: Integrate RPC query into CLI.
- **Day 3** (Ubuntu: Docker, M3: Go):
  - Go Ethereum CLI: Build CLI to log Ethereum block height.
  - Docker Image Build: Build and test Docker image for Go app.
  - Go Logging: Add file logging to CLI for metrics.
  - Docker Volumes: Configure volumes for persistent CLI data.
- **Day 4** (Ubuntu: Geth, M3: Go):
  - Ethereum Systemd: Setup Geth as a systemd service for auto-start.
  - Go CLI Enhancements: Add flags for customizable RPC queries.
  - Ethereum Node Monitoring: Monitor Geth logs for sync status.
  - Go Error Handling: Handle RPC errors in CLI.
- **Day 5** (Ubuntu: Docker/Geth, M3: Go):
  - Docker Hub Push: Push Go app image to Docker Hub.
  - Go Commit: Commit CLI code to GitHub.
  - Ethereum Node Verification: Verify light node sync progress.
  - GitHub README: Add Ethereum CLI setup to repo README.
- **Weekend Project**: `eth-ping` – Dockerized Go CLI to ping Ethereum node, logs to S3, GitHub repo.
- **Definition of Done**: CLI queries block height, Docker image runs, repo has README with setup.
- **Resources**: Ethereum.org “Running a Node”, Docker Docs.
- **Hardware**: Ubuntu MacBook (Geth, Docker), M3 MacBook (Go).
- **Outcome**: Ethereum light node, Dockerized Go app.

---

## Week 6: Prometheus, AWS S3, and Go Metrics
- **Goals**: Setup Prometheus, use S3, build Go metrics tool.
- **Day 1** (Ubuntu: Prometheus, M3: AWS):
  - Prometheus Installation: Install Prometheus and Node Exporter for metrics.
  - AWS S3 Setup: Create S3 bucket and sync sample logs.
  - Prometheus Scraping: Configure Prometheus to scrape Node Exporter.
  - AWS IAM Basics: Setup IAM user for S3 access.
- **Day 2** (Ubuntu: Prometheus, M3: Go):
  - Go Metrics Tool: Build Go tool to export system metrics to Prometheus.
  - Prometheus Configuration: Setup `prometheus.yml` for custom metrics.
  - Go Prometheus Client: Use Prometheus Go client library for metrics.
  - Prometheus UI: Explore Prometheus UI for scraped data.
- **Day 3** (M3):
  - AWS S3 IAM Role: Create IAM role for S3 write access.
  - Go S3 Logging: Add S3 logging to Go metrics tool.
  - AWS CLI S3 Sync: Practice `aws s3 sync` for log uploads.
  - Go CLI Flags: Add flags for customizable metrics output.
- **Day 4** (Ubuntu: Prometheus, M3: Go):
  - Prometheus PromQL: Query metrics using PromQL basics.
  - Go Custom Metrics: Enhance tool with custom system metrics.
  - Prometheus Alerts: Setup basic alert rules in Prometheus.
  - Go Testing: Write tests for metrics tool functionality.
- **Day 5** (Ubuntu: Prometheus, M3: AWS/Go):
  - AWS S3 Testing: Verify log sync to S3 bucket.
  - Go Commit: Commit metrics tool to GitHub.
  - Prometheus Dashboard: Build a basic Prometheus dashboard.
  - GitHub README: Add metrics tool setup to repo README.
- **Weekend Project**: `sys-metrics` – Go tool + Prometheus for system metrics, logs in S3, GitHub repo.
- **Definition of Done**: Tool exports metrics, Prometheus scrapes, S3 stores logs, repo has README with setup.
- **Resources**: Prometheus.io, AWS S3 Docs, “Prometheus Go Client” (YouTube).
- **Hardware**: Ubuntu MacBook (Prometheus), M3 MacBook (AWS, Go).
- **Outcome**: Monitoring basics, S3 integration, Go metrics.

---

## Week 7: Grafana, Ethereum Monitoring, and Terraform Intro
- **Goals**: Setup Grafana, monitor Ethereum node, start Terraform.
- **Day 1** (Ubuntu):
  - Grafana Installation: Install Grafana for visualization.
  - Ethereum Log Monitoring: Monitor Geth logs for sync status.
  - Grafana Data Source: Add Prometheus as a Grafana data source.
  - Ethereum RPC Queries: Query Geth RPC for block height.
- **Day 2** (Ubuntu: Grafana, M3: Terraform):
  - Terraform Installation: Install Terraform for infrastructure as code.
  - Grafana Dashboard: Build CPU/memory dashboard in Grafana.
  - Terraform S3: Write Terraform script for S3 bucket.
  - Grafana Panels: Create basic panels for system metrics.
- **Day 3** (Ubuntu: Geth, M3: Go):
  - Ethereum Metrics: Query Geth for metrics (e.g., peer count).
  - Go Prometheus Exporter: Build Go tool to export Geth metrics.
  - Ethereum Node Optimization: Adjust Geth cache for performance.
  - Go Metrics Integration: Integrate Geth metrics into Prometheus.
- **Day 4** (Ubuntu: Grafana, M3: Terraform):
  - Grafana Ethereum Dashboard: Add Geth metrics to Grafana dashboard.
  - Terraform EC2: Write Terraform script for EC2 instance.
  - Grafana Visualizations: Customize dashboard with graphs and gauges.
  - Terraform Variables: Use variables in Terraform scripts.
- **Day 5** (Ubuntu: Geth/Grafana, M3: Go):
  - Go Commit: Commit Geth metrics tool to GitHub.
  - Ethereum Sync Check: Verify Geth light node sync status.
  - Grafana Finalization: Finalize Ethereum metrics dashboard.
  - GitHub README: Add dashboard setup to repo README.
- **Weekend Project**: `eth-monitor` – Go tool + Grafana dashboard for Ethereum metrics, GitHub repo.
- **Definition of Done**: Dashboard shows Geth metrics, tool exports data, repo has README with screenshots.
- **Resources**: Grafana Labs, Ethereum.org, HashiCorp “Learn Terraform”.
- **Hardware**: Ubuntu MacBook (Geth, Grafana), M3 MacBook (Terraform, Go).
- **Outcome**: Ethereum monitoring, Terraform basics.

---

## Week 8: Terraform EC2, Go, and Kubernetes Intro
- **Goals**: Provision EC2 with Terraform, enhance Go CLI, start Kubernetes.
- **Day 1** (M3):
  - Terraform EC2 Setup: Write Terraform script for EC2 and security groups.
  - Go SSH CLI: Build Go CLI to SSH into EC2 instances.
  - Terraform Security Groups: Define inbound rules for EC2.
  - Go CLI Planning: Plan CLI structure for EC2 interaction.
- **Day 2** (M3):
  - Kubernetes Minikube: Install Minikube for local Kubernetes cluster.
  - Terraform Variables: Add variables to EC2 Terraform script.
  - Kubernetes Pods: Deploy a sample pod in Minikube.
  - Terraform State: Understand Terraform state management.
- **Day 3** (M3):
  - Go Log Tailing: Add log tailing functionality to Go CLI.
  - Kubernetes Deployments: Create a deployment in Minikube.
  - Go SDK Integration: Integrate AWS SDK for EC2 status checks.
  - Kubernetes Services: Explore Service types for exposing pods.
- **Day 4** (M3):
  - Terraform Deployment: Test EC2 deployment with Terraform.
  - Go CLI Enhancements: Add status checks to CLI for EC2.
  - Kubernetes Networking: Configure networking for deployments.
  - Go Error Handling: Handle SSH and AWS errors in CLI.
- **Day 5** (M3):
  - Kubernetes Service Exposure: Expose deployment via Service in Minikube.
  - Go Commit: Commit CLI code to GitHub.
  - Terraform Commit: Push Terraform scripts to GitHub.
  - GitHub README: Add EC2 CLI setup to repo README.
- **Weekend Project**: `aws-infra` – Terraform EC2 deployment + Go CLI for log tailing, GitHub repo.
- **Definition of Done**: Terraform deploys EC2, CLI tails logs, repo has README with setup.
- **Resources**: HashiCorp Terraform, Kubernetes.io “Basics”.
- **Hardware**: M3 MacBook (all tasks, Minikube RAM-intensive).
- **Outcome**: IaC with Terraform, Kubernetes intro.

---

## Week 9: Kubernetes, Ethereum, and Go Health Checks
- **Goals**: Deploy to Kubernetes, optimize Ethereum node, build Go health checker.
- **Day 1** (Ubuntu: Geth, M3: Kubernetes):
  - Kubernetes App Deployment: Deploy Go app to Minikube.
  - Ethereum Optimization: Adjust Geth cache for light node performance.
  - Kubernetes Pods: Monitor pod status in Minikube.
  - Ethereum RPC Testing: Test Geth RPC endpoints for reliability.
- **Day 2** (M3):
  - Go Health Checker: Build Go CLI to check Geth RPC health.
  - Kubernetes Services: Create Service for Go app in Minikube.
  - Go HTTP Client: Use HTTP client for Geth health checks.
  - Kubernetes ConfigMaps: Configure app settings with ConfigMaps.
- **Day 3** (Ubuntu: Geth, M3: Go):
  - Ethereum RPC Queries: Query Geth for health metrics (e.g., peers).
  - Go Health Integration: Add health check logic to CLI.
  - Ethereum Log Analysis: Analyze Geth logs for issues.
  - Go Testing: Write tests for health checker functionality.
- **Day 4** (M3):
  - Kubernetes StatefulSet: Write StatefulSet for Geth deployment.
  - Go Pod Checks: Enhance CLI to check Kubernetes pod health.
  - Kubernetes PVCs: Configure PersistentVolumeClaims for Geth data.
  - Go CLI Refinement: Optimize CLI for user-friendly output.
- **Day 5** (Ubuntu: Geth, M3: Kubernetes/Go):
  - Kubernetes Testing: Test StatefulSet deployment in Minikube.
  - Go Commit: Commit health checker CLI to GitHub.
  - Ethereum Verification: Verify Geth node stability.
  - GitHub README: Add Kubernetes setup to repo README.
- **Weekend Project**: `kube-eth` – Kubernetes deployment + Go health checker, GitHub repo.
- **Definition of Done**: StatefulSet runs Geth, CLI checks health, repo has README with setup.
- **Resources**: Kubernetes.io, Ethereum.org.
- **Hardware**: M3 MacBook (Kubernetes, Go), Ubuntu MacBook (Geth).
- **Outcome**: Kubernetes deployment, advanced node ops.

---

## Week 10: AWS KMS, Go Secrets, and Solana Intro
- **Goals**: Secure secrets with KMS, build Go secrets CLI, explore Solana testnet.
- **Day 1** (M3):
  - AWS KMS Setup: Store a fake validator key in AWS KMS.
  - Go KMS CLI: Build Go CLI to fetch KMS secrets.
  - AWS IAM Policies: Create IAM policy for KMS access.
  - Go SDK Setup: Configure AWS SDK for KMS integration.
- **Day 2** (M3: KMS, AWS: Solana):
  - Solana CLI Setup: Install Solana CLI on AWS Free Tier for testnet.
  - AWS KMS IAM: Configure IAM role for KMS access.
  - Solana Testnet Node: Run Solana testnet node on AWS.
  - Go KMS Integration: Integrate KMS fetch into CLI.
- **Day 3** (M3: Go, AWS: Solana):
  - Go Encryption: Add encryption functionality to CLI using KMS.
  - Solana Node Queries: Query Solana testnet node for status.
  - Go Error Handling: Handle KMS errors in CLI.
  - Solana Monitoring: Monitor Solana node logs for performance.
- **Day 4** (M3):
  - AWS KMS Testing: Test secret retrieval with KMS CLI.
  - Go Decryption: Add decryption functionality to CLI.
  - AWS KMS Best Practices: Study KMS security best practices.
  - Go CLI Optimization: Optimize CLI for secure key handling.
- **Day 5** (M3: Go/KMS, AWS: Solana):
  - Solana Node Management: Monitor and manage Solana testnet node.
  - Go Commit: Commit KMS CLI to GitHub.
  - AWS KMS Commit: Push KMS config to GitHub.
  - GitHub README: Add KMS and Solana setup to repo README.
- **Weekend Project**: `secure-secrets` – Go CLI for KMS-encrypted keys, Solana testnet script, GitHub repo.
- **Definition of Done**: CLI encrypts/decrypts, Solana node runs, repo has README with setup.
- **Resources**: AWS KMS Docs, Solana Docs, “Go AWS SDK” (YouTube).
- **Hardware**: M3 MacBook (KMS, Go), AWS Free Tier (Solana).
- **Outcome**: Secure key management, Solana intro.

---

## Week 11: Final Project and Portfolio
- **Goals**: Build Web3 toolkit, polish GitHub portfolio.
- **Day 1** (M3):
  - Terraform EKS Setup: Write Terraform script for AWS EKS cluster.
  - Go EKS CLI: Build Go CLI for EKS deployment management.
  - Terraform Modules: Organize EKS script into reusable modules.
  - Go CLI Planning: Plan CLI for EKS interaction.
- **Day 2** (M3):
  - EKS Geth Deployment: Deploy Geth + Go app to EKS with Terraform.
  - Portfolio README: Write README for `eth-monitor` project.
  - EKS Networking: Configure EKS networking for Geth.
  - Portfolio Diagrams: Create architecture diagram for `web3-toolkit`.
- **Day 3** (M3):
  - EKS Monitoring: Add Prometheus/Grafana to EKS deployment.
  - Portfolio Enhancement: Add screenshots to `kube-eth` README.
  - Prometheus Operator: Install Prometheus operator in EKS.
  - Grafana EKS Dashboard: Create EKS dashboard for Geth metrics.
- **Day 4** (M3):
  - EKS Testing: Test Geth + Go app deployment in EKS.
  - Portfolio Refinement: Update `aws-infra` README with details.
  - Terraform Validation: Validate EKS Terraform scripts.
  - Portfolio Organization: Pin 5-7 projects to GitHub profile.
- **Day 5** (M3):
  - EKS Project Commit: Commit `web3-toolkit` to GitHub.
  - Portfolio Finalization: Finalize all repo READMEs and diagrams.
  - Go CLI Testing: Test EKS CLI for functionality.
  - GitHub README: Add EKS setup to `web3-toolkit` README.
- **Weekend Project**: `web3-toolkit` – EKS deployment of Geth + Go app, monitoring, GitHub repo.
- **Definition of Done**: EKS runs Geth/app, dashboard works, repos have READMEs/diagrams.
- **Resources**: Terraform EKS Module, draw.io.
- **Hardware**: M3 MacBook (all tasks).
- **Outcome**: Job-ready portfolio with 7 projects.

---

## Week 12: AWS Certification Prep
- **Goals**: Prepare for AWS Certified Solutions Architect – Associate.
- **Day 1** (M3):
  - AWS EC2 Study: Learn EC2 instance types, AMIs, and scaling.
  - AWS S3 Study: Study S3 storage classes and bucket policies.
  - Portfolio Certification: Add certification prep section to GitHub.
  - AWS VPC Basics: Understand VPC, subnets, and routing.
- **Day 2** (M3):
  - AWS IAM Study: Explore IAM users, roles, and policies.
  - AWS ECS Basics: Learn ECS for container orchestration.
  - Portfolio README: Refine READMEs for all projects.
  - AWS Route 53: Study DNS management with Route 53.
- **Day 3** (M3):
  - AWS Well-Architected: Study Well-Architected Framework pillars.
  - Portfolio Screenshots: Add screenshots to project READMEs.
  - AWS SQS/SNS: Understand messaging and notification services.
  - Exam Question Practice: Practice AWS exam questions (Udemy).
- **Day 4** (M3):
  - AWS Practice Exam: Take full-length practice exam (Whizlabs).
  - Portfolio Diagrams: Finalize architecture diagrams for repos.
  - AWS Exam Strategy: Learn to dissect exam questions and manage time.
  - AWS RDS Basics: Study relational database service basics.
- **Day 5** (M3):
  - AWS Weak Areas: Review weak areas from practice exam.
  - Portfolio Blog: Write blog on “Deploying Ethereum Node on EKS”.
  - AWS Exam Scheduling: Schedule certification exam.
  - Portfolio Final Check: Ensure all repos are polished.
- **Weekend Project**: Updated portfolio with blog: “Deploying Ethereum Node on EKS”.
- **Definition of Done**: Score 80%+ on practice exams, blog published, portfolio complete.
- **Resources**: Udemy “AWS Solutions Architect” (Stephane Maarek), Whizlabs.
- **Hardware**: M3 MacBook.
- **Outcome**: Ready for AWS certification, polished portfolio.

---

## Week 13: Job Applications and Networking
- **Goals**: Apply to jobs, network, finalize portfolio.
- **Day 1** (M3):
  - Job Applications: Apply to 10 DevOps/Web3 jobs on Web3.career, LinkedIn.
  - X Networking: Engage with @ethereum posts on X.
  - Resume Tailoring: Customize resume for DevOps roles.
  - Portfolio Video Planning: Plan 5-minute `web3-toolkit` demo video.
- **Day 2** (M3):
  - Job Applications: Apply to 10 jobs on CryptoJobsList, Turing.
  - LinkedIn Networking: Join Web3 groups on LinkedIn.
  - Cover Letter Writing: Write cover letters for top 5 jobs.
  - Portfolio Video Script: Draft script for demo video.
- **Day 3** (M3):
  - Job Applications: Apply to 5 jobs with tailored applications.
  - LinkedIn Blog Sharing: Share blog post on LinkedIn.
  - Behavioral Interview Prep: Practice STAR method for interviews.
  - Portfolio Video Recording: Record `web3-toolkit` demo video.
- **Day 4** (M3):
  - Job Follow-Ups: Follow up on previous job applications.
  - X Engagement: Comment on @awscloud posts on X.
  - Technical Interview Prep: Prepare to explain `web3-toolkit` project.
  - Portfolio Video Editing: Edit and finalize demo video.
- **Day 5** (M3):
  - Job Applications: Apply to 5 jobs with custom cover letters.
  - Recruiter Outreach: Connect with 10 recruiters on LinkedIn.
  - Portfolio Video Upload: Upload demo video to GitHub/YouTube.
  - GitHub README: Add video link to `web3-toolkit` README.
- **Weekend Project**: Final portfolio, 5-minute video demo of `web3-toolkit`.
- **Definition of Done**: 30+ applications submitted, video demo uploaded, 10+ connections made.
- **Resources**: Web3.career, CryptoJobsList, STAR Method for Resumes.
- **Hardware**: M3 MacBook.
- **Outcome**: Active job applications, strong online presence.

---

## Outcomes
- **Skills**: Git, Docker, Terraform, Kubernetes, AWS, Go, Ethereum light node, Solana testnet, Prometheus/Grafana.
- **Portfolio**: 7 GitHub projects (`sys-monitor`, `docker-check`, `docker-stack`, `aws-monitor`, `eth-ping`, `kube-eth`, `web3-toolkit`).
- **Certification**: AWS Solutions Architect – Associate ($150).

