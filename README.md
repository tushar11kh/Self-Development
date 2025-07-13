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
- **Core Concepts**:
  - Linux: `ls`, `cd`, `chmod`, `systemd`, `ufw`, SSH.
  - Git: `clone`, `commit`, `push`, `branch`.
  - Go: Variables, functions, error handling, goroutines.
- **Daily Breakdown (8 hours/day)**:
  - **Day 1**: Linux (4h): Setup SSH (`sudo apt install openssh-server`), practice `ls`, `cd`, `chmod`. Git (4h): Create GitHub repo, push sample file. (Ubuntu: Linux, M3: Git)
  - **Day 2**: Linux (3h): Configure `ufw` (`sudo ufw allow 22`). Go (5h): Start “A Tour of Go” (go.dev/tour), write program to print system info. (Ubuntu: Linux, M3: Go)
  - **Day 3**: Git (3h): Learn `branch`, `pull request`. Go (5h): Continue “A Tour of Go” (functions, structs). (M3)
  - **Day 4**: Linux (3h): Write Bash script for CPU usage. Go (5h): Build CLI to log memory usage. (Ubuntu: Linux, M3: Go)
  - **Day 5**: Git (2h): Commit CLI to GitHub. Go (4h): Add flags to CLI (`--cpu`, `--mem`). Linux (2h): Setup `systemd` for CLI. (Ubuntu: Linux, M3: Git/Go)
- **Weekend Project**: `sys-monitor` – Go CLI + Bash script to log CPU/memory, systemd service, GitHub repo with README.
- **Definition of Done**: CLI runs with flags, systemd service auto-starts, repo has README with setup steps.
- **Resources**: FreeCodeCamp “Linux Basics” (YouTube), GitHub Docs, “A Tour of Go”.
- **Hardware**: Ubuntu MacBook (Linux, systemd), M3 MacBook (Git, Go).
- **Outcome**: Linux CLI proficiency, Git workflow, basic Go CLI.

---

## Week 2: Docker, Go CLI, and AWS Intro
- **Goals**: Learn Docker basics, build Go CLI, explore AWS Free Tier.
- **Core Concepts**:
  - Docker: Dockerfile, containers, images, `docker run`.
  - Go: CLI flags, HTTP requests.
  - AWS: Free Tier, EC2, S3, AWS CLI.
- **Daily Breakdown**:
  - **Day 1**: Docker (4h): Install (`sudo apt install docker.io`), run `hello-world`. Go (4h): Add HTTP request to `sys-monitor`. (Ubuntu: Docker, M3: Go)
  - **Day 2**: AWS (3h): Create Free Tier account, explore EC2/S3. Docker (5h): Write Dockerfile for Go app. (Ubuntu: Docker, M3: AWS)
  - **Day 3**: Go (4h): Build CLI to check Docker container status. AWS (4h): Install AWS CLI (`sudo apt install awscli`), run `aws s3 ls`. (M3)
  - **Day 4**: Docker (3h): Run multi-container app (NGINX). Go (5h): Enhance CLI with container start/stop. (Ubuntu: Docker, M3: Go)
  - **Day 5**: AWS (3h): Launch EC2 (t3.micro), SSH in. Go (3h): Commit CLI to GitHub. Docker (2h): Push image to Docker Hub. (Ubuntu: Docker, M3: AWS/Go)
- **Weekend Project**: `docker-check` – Go CLI to manage Docker containers, Dockerized, pushed to GitHub/Docker Hub.
- **Definition of Done**: CLI starts/stops containers, Dockerfile builds image, repo has README with usage.
- **Resources**: TechWorld with Nana “Docker Tutorial” (YouTube), AWS Docs, Docker Docs.
- **Hardware**: Ubuntu MacBook (Docker), M3 MacBook (AWS, Go).
- **Outcome**: Docker basics, Go CLI, AWS console familiarity.

---

## Week 3: Docker Compose, Git CI/CD, and Go HTTP
- **Goals**: Master Docker Compose, setup CI/CD, build Go HTTP client.
- **Core Concepts**:
  - Docker Compose: `docker-compose.yml`, services, networks.
  - Git CI/CD: GitHub Actions, workflows, secrets.
  - Go: HTTP clients, JSON parsing.
- **Daily Breakdown**:
  - **Day 1**: Docker Compose (4h): Write `docker-compose.yml` for NGINX + Redis. Go (4h): Build HTTP client for public API (jsonplaceholder.typicode.com). (Ubuntu: Compose, M3: Go)
  - **Day 2**: Git CI/CD (4h): Setup GitHub Actions for `docker-check` (build/test). Docker Compose (4h): Run stack locally. (M3: Git, Ubuntu: Compose)
  - **Day 3**: Go (4h): Add HTTP client to `docker-check` for container stats. Git CI/CD (4h): Add linting (`golangci-lint`). (M3)
  - **Day 4**: Docker Compose (3h): Add Go app to Compose stack. Go (5h): Parse API response in CLI. (Ubuntu: Compose, M3: Go)
  - **Day 5**: Git CI/CD (3h): Push Docker image via Actions. Go (3h): Commit to GitHub. Docker Compose (2h): Test stack. (M3: Git/Go, Ubuntu: Compose)
- **Weekend Project**: `docker-stack` – Compose file + Go HTTP client, CI/CD pipeline, GitHub repo.
- **Definition of Done**: Stack runs NGINX + Redis + Go app, CI/CD builds/pushes image, repo has README with setup.
- **Resources**: Docker Compose Docs, GitHub Actions Docs, “Go HTTP Clients” (YouTube).
- **Hardware**: Ubuntu MacBook (Compose), M3 MacBook (Git, Go).
- **Outcome**: Multi-container apps, CI/CD pipeline, Go HTTP skills.

---

## Week 4: AWS EC2, Go AWS SDK, and Linux Networking
- **Goals**: Launch EC2, use Go AWS SDK, configure Linux networking.
- **Core Concepts**:
  - AWS: EC2 instances, security groups, IAM, AWS CLI.
  - Go: AWS SDK, API calls.
  - Linux: Networking (`ss`, `dig`), `fail2ban`.
- **Daily Breakdown**:
  - **Day 1**: AWS (4h): Launch EC2 (t3.micro), configure security groups. Linux (4h): Setup `ufw` (`sudo ufw allow 22,80`). (Ubuntu: Linux, M3: AWS)
  - **Day 2**: Go (4h): Use AWS SDK to list EC2 instances. AWS (4h): SSH into EC2, test `aws s3 ls`. (M3)
  - **Day 3**: Linux (3h): Install `fail2ban`. Go (5h): Build CLI to list S3 buckets. (Ubuntu: Linux, M3: Go)
  - **Day 4**: AWS (4h): Create IAM role for EC2-S3 access. Go (4h): Add S3 listing to CLI. (M3)
  - **Day 5**: Linux (3h): Test network tools (`ss`, `dig`). Go (3h): Commit CLI to GitHub. AWS (2h): Terminate EC2. (Ubuntu: Linux, M3: Go/AWS)
- **Weekend Project**: `aws-monitor` – Go CLI to list EC2/S3 resources, GitHub repo with README.
- **Definition of Done**: CLI lists EC2/S3, IAM role works, repo has README with usage.
- **Resources**: AWS EC2 Docs, AWS SDK Go Docs, Ubuntu Server Guide.
- **Hardware**: Ubuntu MacBook (Linux), M3 MacBook (AWS, Go).
- **Outcome**: EC2 management, Go SDK usage, Linux networking.

---

## Week 5: Ethereum Light Node, Docker, and Go
- **Goals**: Run Ethereum light node, Dockerize Go app, enhance CLI.
- **Core Concepts**:
  - Ethereum: Geth, light sync, JSON-RPC.
  - Docker: Dockerizing Go apps, multi-stage builds.
  - Go: RPC queries, CLI enhancements.
- **Daily Breakdown**:
  - **Day 1**: Ethereum (4h): Install Geth (`sudo apt install geth`), run `geth --syncmode light`. Go (4h): Query RPC (`eth_blockNumber`). (Ubuntu: Geth, M3: Go)
  - **Day 2**: Docker (4h): Write Dockerfile for Go app. Ethereum (4h): Test Geth RPC with `curl`. (Ubuntu: Docker/Geth, M3: Go)
  - **Day 3**: Go (4h): Build CLI to log Ethereum block height. Docker (4h): Build/run Docker image. (Ubuntu: Docker, M3: Go)
  - **Day 4**: Ethereum (3h): Setup Geth systemd service. Go (5h): Add logging to CLI. (Ubuntu: Geth, M3: Go)
  - **Day 5**: Docker (3h): Push image to Docker Hub. Go (3h): Commit CLI to GitHub. Ethereum (2h): Verify node sync. (Ubuntu: Docker/Geth, M3: Go)
- **Weekend Project**: `eth-ping` – Dockerized Go CLI to ping Ethereum node, logs to S3, GitHub repo.
- **Definition of Done**: CLI queries block height, Docker image runs, repo has README with setup.
- **Resources**: Ethereum.org “Running a Node”, Docker Docs.
- **Hardware**: Ubuntu MacBook (Geth, Docker), M3 MacBook (Go).
- **Outcome**: Ethereum light node, Dockerized Go app.

---

## Week 6: Prometheus, AWS S3, and Go Metrics
- **Goals**: Setup Prometheus, use S3, build Go metrics tool.
- **Core Concepts**:
  - Prometheus: Scraping, Node Exporter, PromQL.
  - AWS: S3 buckets, IAM policies.
  - Go: Prometheus client, metrics export.
- **Daily Breakdown**:
  - **Day 1**: Prometheus (4h): Install Prometheus, Node Exporter. AWS (4h): Create S3 bucket, sync logs. (Ubuntu: Prometheus, M3: AWS)
  - **Day 2**: Go (4h): Build tool to export system metrics to Prometheus. Prometheus (4h): Configure scraping. (Ubuntu: Prometheus, M3: Go)
  - **Day 3**: AWS (3h): Setup IAM for S3 access. Go (5h): Add S3 logging to tool. (M3)
  - **Day 4**: Prometheus (4h): Query metrics with PromQL. Go (4h): Enhance tool with custom metrics. (Ubuntu: Prometheus, M3: Go)
  - **Day 5**: AWS (3h): Test S3 sync. Go (3h): Commit to GitHub. Prometheus (2h): Build basic dashboard. (Ubuntu: Prometheus, M3: AWS/Go)
- **Weekend Project**: `sys-metrics` – Go tool + Prometheus for system metrics, logs in S3, GitHub repo.
- **Definition of Done**: Tool exports metrics, Prometheus scrapes, S3 stores logs, repo has README with setup.
- **Resources**: Prometheus.io, AWS S3 Docs, “Prometheus Go Client” (YouTube).
- **Hardware**: Ubuntu MacBook (Prometheus), M3 MacBook (AWS, Go).
- **Outcome**: Monitoring basics, S3 integration, Go metrics.

---

## Week 7: Grafana, Ethereum Monitoring, and Terraform Intro
- **Goals**: Setup Grafana, monitor Ethereum node, start Terraform.
- **Core Concepts**:
  - Grafana: Dashboards, Prometheus data source.
  - Ethereum: Node metrics (block height, peers).
  - Terraform: HCL, providers, resources.
- **Daily Breakdown**:
  - **Day 1**: Grafana (4h): Install (`sudo apt install grafana`), add Prometheus data source. Ethereum (4h): Monitor Geth logs. (Ubuntu)
  - **Day 2**: Terraform (4h): Install (`sudo apt install terraform`), write S3 bucket script. Grafana (4h): Build CPU/memory dashboard. (Ubuntu: Grafana, M3: Terraform)
  - **Day 3**: Ethereum (3h): Query Geth metrics. Go (5h): Build tool to export Geth metrics to Prometheus. (Ubuntu: Geth, M3: Go)
  - **Day 4**: Grafana (4h): Add Geth metrics to dashboard. Terraform (4h): Write EC2 script. (Ubuntu: Grafana, M3: Terraform)
  - **Day 5**: Go (3h): Commit tool to GitHub. Ethereum (3h): Optimize Geth sync. Grafana (2h): Finalize dashboard. (Ubuntu: Geth/Grafana, M3: Go)
- **Weekend Project**: `eth-monitor` – Go tool + Grafana dashboard for Ethereum metrics, GitHub repo.
- **Definition of Done**: Dashboard shows Geth metrics, tool exports data, repo has README with screenshots.
- **Resources**: Grafana Labs, Ethereum.org, HashiCorp “Learn Terraform”.
- **Hardware**: Ubuntu MacBook (Geth, Grafana), M3 MacBook (Terraform, Go).
- **Outcome**: Ethereum monitoring, Terraform basics.

---

## Week 8: Terraform EC2, Go, and Kubernetes Intro
- **Goals**: Provision EC2 with Terraform, enhance Go CLI, start Kubernetes.
- **Core Concepts**:
  - Terraform: EC2 provisioning, variables, modules.
  - Go: SSH integration, CLI enhancements.
  - Kubernetes: Minikube, pods, deployments.
- **Daily Breakdown**:
  - **Day 1**: Terraform (4h): Write EC2 + security group script. Go (4h): Build CLI to SSH into EC2. (M3)
  - **Day 2**: Kubernetes (4h): Install Minikube (M3), deploy sample pod. Terraform (4h): Add variables to script. (M3)
  - **Day 3**: Go (4h): Add log tailing to CLI. Kubernetes (4h): Create deployment. (M3)
  - **Day 4**: Terraform (4h): Test EC2 deployment. Go (4h): Enhance CLI with status checks. (M3)
  - **Day 5**: Kubernetes (3h): Expose deployment via Service. Go (3h): Commit to GitHub. Terraform (2h): Push script to GitHub. (M3)
- **Weekend Project**: `aws-infra` – Terraform EC2 deployment + Go CLI for log tailing, GitHub repo.
- **Definition of Done**: Terraform deploys EC2, CLI tails logs, repo has README with setup.
- **Resources**: HashiCorp Terraform, Kubernetes.io “Basics”.
- **Hardware**: M3 MacBook (all tasks, Minikube RAM-intensive).
- **Outcome**: IaC with Terraform, Kubernetes intro.

---

## Week 9: Kubernetes, Ethereum, and Go Health Checks
- **Goals**: Deploy to Kubernetes, optimize Ethereum node, build Go health checker.
- **Core Concepts**:
  - Kubernetes: StatefulSet, Services, PVCs.
  - Ethereum: Light node optimization, RPC calls.
  - Go: Health check logic, Kubernetes SDK.
- **Daily Breakdown**:
  - **Day 1**: Kubernetes (4h): Deploy Go app to Minikube. Ethereum (4h): Optimize Geth (`--cache`). (Ubuntu: Geth, M3: Kubernetes)
  - **Day 2**: Go (4h): Build health checker for Geth RPC. Kubernetes (4h): Create Service for app. (M3)
  - **Day 3**: Ethereum (3h): Test RPC calls. Go (5h): Add health check to CLI. (Ubuntu: Geth, M3: Go)
  - **Day 4**: Kubernetes (4h): Write StatefulSet for Geth. Go (4h): Enhance CLI with pod checks. (M3)
  - **Day 5**: Kubernetes (3h): Test StatefulSet. Go (3h): Commit to GitHub. Ethereum (2h): Verify node. (Ubuntu: Geth, M3: Kubernetes/Go)
- **Weekend Project**: `kube-eth` – Kubernetes deployment + Go health checker, GitHub repo.
- **Definition of Done**: StatefulSet runs Geth, CLI checks health, repo has README with setup.
- **Resources**: Kubernetes.io, Ethereum.org.
- **Hardware**: M3 MacBook (Kubernetes, Go), Ubuntu MacBook (Geth).
- **Outcome**: Kubernetes deployment, advanced node ops.

---

## Week 10: AWS KMS, Go Secrets, and Solana Intro
- **Goals**: Secure secrets with KMS, build Go secrets CLI, explore Solana testnet.
- **Core Concepts**:
  - AWS KMS: Secret storage, IAM policies.
  - Go: KMS SDK, encryption/decryption.
  - Solana: Testnet node, CLI setup.
- **Daily Breakdown**:
  - **Day 1**: AWS KMS (4h): Store fake validator key in KMS. Go (4h): Build CLI to fetch KMS secrets. (M3)
  - **Day 2**: Solana (4h): Setup Solana CLI on AWS Free Tier, run testnet node. AWS KMS (4h): Configure IAM. (M3: KMS, AWS: Solana)
  - **Day 3**: Go (4h): Add encryption to CLI. Solana (4h): Query testnet node. (M3: Go, AWS: Solana)
  - **Day 4**: AWS KMS (4h): Test secret retrieval. Go (4h): Enhance CLI with decryption. (M3)
  - **Day 5**: Solana (3h): Monitor testnet node. Go (3h): Commit CLI to GitHub. AWS KMS (2h): Push config to GitHub. (M3: Go/KMS, AWS: Solana)
- **Weekend Project**: `secure-secrets` – Go CLI for KMS-encrypted keys, Solana testnet script, GitHub repo.
- **Definition of Done**: CLI encrypts/decrypts, Solana node runs, repo has README with setup.
- **Resources**: AWS KMS Docs, Solana Docs, “Go AWS SDK” (YouTube).
- **Hardware**: M3 MacBook (KMS, Go), AWS Free Tier (Solana).
- **Outcome**: Secure key management, Solana intro.

---

## Week 11: Final Project and Portfolio
- **Goals**: Build Web3 toolkit, polish GitHub portfolio.
- **Core Concepts**:
  - Integration: Terraform, Kubernetes, Go, Ethereum.
  - Portfolio: READMEs, diagrams, documentation.
- **Daily Breakdown**:
  - **Day 1**: Project (4h): Write Terraform for EKS cluster. Go (4h): Build CLI for EKS deployment. (M3)
  - **Day 2**: Project (4h): Deploy Geth + Go app to EKS. Portfolio (4h): Write README for `eth-monitor`. (M3)
  - **Day 3**: Project (4h): Add Prometheus/Grafana to EKS. Portfolio (4h): Create diagram for `web3-toolkit`. (M3)
  - **Day 4**: Project (4h): Test EKS deployment. Portfolio (4h): Update `kube-eth` README. (M3)
  - **Day 5**: Project (3h): Commit `web3-toolkit` to GitHub. Portfolio (5h): Finalize repos, pin 5-7 projects. (M3)
- **Weekend Project**: `web3-toolkit` – EKS deployment of Geth + Go app, monitoring, GitHub repo.
- **Definition of Done**: EKS runs Geth/app, dashboard works, repos have READMEs/diagrams.
- **Resources**: Terraform EKS Module, draw.io.
- **Hardware**: M3 MacBook (all tasks).
- **Outcome**: Job-ready portfolio with 7 projects.

---

## Week 12: AWS Certification Prep
- **Goals**: Prepare for AWS Certified Solutions Architect – Associate.
- **Core Concepts**:
  - AWS: EC2, S3, VPC, IAM, ECS, Well-Architected Framework.
  - Exam: Question analysis, time management.
- **Daily Breakdown**:
  - **Day 1**: AWS (6h): Study EC2, S3, VPC (Udemy). Portfolio (2h): Add certification prep to GitHub. (M3)
  - **Day 2**: AWS (6h): Study IAM, ECS, Route 53. Portfolio (2h): Refine READMEs. (M3)
  - **Day 3**: AWS (6h): Study Well-Architected Framework. Portfolio (2h): Add screenshots. (M3)
  - **Day 4**: AWS (6h): Take practice exam (Whizlabs). Portfolio (2h): Finalize diagrams. (M3)
  - **Day 5**: AWS (6h): Review weak areas, schedule exam. Portfolio (2h): Publish blog on Medium. (M3)
- **Weekend Project**: Updated portfolio with blog: “Deploying Ethereum Node on EKS”.
- **Definition of Done**: Score 80%+ on practice exams, blog published, portfolio complete.
- **Resources**: Udemy “AWS Solutions Architect” (Stephane Maarek), Whizlabs.
- **Hardware**: M3 MacBook.
- **Outcome**: Ready for AWS certification, polished portfolio.

---

## Week 13: Job Applications and Networking
- **Goals**: Apply to jobs, network, finalize portfolio.
- **Core Concepts**:
  - Job Hunt: Resume tailoring, cover letters.
  - Networking: X, LinkedIn engagement.
- **Daily Breakdown**:
  - **Day 1**: Applications (5h): Apply to 10 jobs (Web3.career, LinkedIn). Networking (3h): Engage on X (@ethereum). (M3)
  - **Day 2**: Applications (5h): Apply to 10 jobs (CryptoJobsList, Turing). Networking (3h): Join LinkedIn Web3 groups. (M3)
  - **Day 3**: Applications (4h): Apply to 5 jobs, write cover letters. Networking (4h): Share blog on LinkedIn. (M3)
  - **Day 4**: Applications (4h): Follow up on applications. Networking (4h): Comment on X posts (@awscloud). (M3)
  - **Day 5**: Applications (4h): Apply to 5 jobs. Networking (4h): Connect with recruiters. (M3)
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
- **Job Prospects**: Remote DevOps/Web3 roles.

