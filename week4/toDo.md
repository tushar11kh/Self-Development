### ✅ 1. Set Up and Configure EC2  
Launch a t3.micro EC2 instance via AWS Console.  
Create and assign security groups (allowing inbound SSH (22), HTTP (80)).  
Create a key pair, download it securely.  
Configure inbound/outbound rules in the security group.  

### ✅ 2. SSH into EC2 Instance  
From your Mac M3, SSH into the EC2 instance using the key.  
Ensure the connection works.  
(Optional but good practice) Add the EC2 public IP to ~/.ssh/config for convenience.  

### ✅ 3. Linux Firewall Setup (on EC2)  
Inside the EC2 (Ubuntu) instance:  
Install and configure ufw.  
Allow ports 22 and 80.  
Enable ufw and test connectivity.  

### ✅ 4. Linux Security Hardening  
Install and configure fail2ban on EC2.  
Verify it's protecting SSH from brute-force attempts.  

### ✅ 5. Linux Networking and Diagnostics  
Use ping, traceroute, and dig for network diagnostics.  
Explore active connections with ss, and understand ports/services.  
Practice tools to inspect your EC2 network behavior.  

### ✅ 6. Linux File Management  
Practice find, grep, xargs for searching, filtering files/logs.  
Explore /var/log, try finding login attempts, or banned IPs (fail2ban).  

### ✅ 7. Set Up AWS CLI on M3 MacBook  
Install AWS CLI.  
Configure aws configure with access keys and default region.  
Verify setup with aws ec2 describe-instances.  

### ✅ 8. Create IAM Role for EC2-S3  
Create an IAM Role with S3 ReadOnly access.  
Attach the role to your EC2 instance.  
Verify the instance can access S3 using CLI without credentials (aws s3 ls).  

### ✅ 9. Write Go Code to Use AWS SDK (EC2)  
Initialize a Go project.  
Configure AWS SDK for Go (v2).  
List EC2 instances using the SDK.  
Handle errors gracefully.  

### ✅ 10. Extend Go CLI for S3 Bucket Listing  
Use AWS SDK to list S3 buckets.  
Implement error handling and retry logic.  
Print buckets in a clean CLI-friendly format.  

### ✅ 11. Refine Go CLI for Usability  
Add flags (e.g. --list-ec2, --list-s3) for modular usage.  
Structure output with formatting or JSON as needed.  
Document usage in --help command.  

### ✅ 12. Commit Code to GitHub  
Push the Go CLI project to a public/private GitHub repo.  
Include proper Go module files, structure, and code comments.  

### ✅ 13. Write README with AWS CLI + Go CLI Usage  
Add clear sections:  
Project description  
Setup instructions  
AWS CLI setup guide  
Go CLI usage examples  
IAM role setup explanation  
Cleanup instructions  

### ✅ 14. Clean Up AWS Resources  
Terminate the EC2 instance.  
Delete unnecessary security groups, IAM roles, key pairs.  

### ✅ 15. Weekend Project: aws-monitor  
Finalize your Go CLI into a tool named aws-monitor.  
Ensure it:  
Lists EC2 instances  
Lists S3 buckets  
Handles errors  
Has clean output  
Make the repo production-readiness level with a complete README.  
```
