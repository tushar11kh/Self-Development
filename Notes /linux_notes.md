# Week1

## SSH Server Initialization

Install the SSH server:
`sudo apt install openssh-server`.

Start the SSH service:
`sudo systemctl start ssh`.

Enable SSH to run on boot:
`sudo systemctl enable ssh`.

Verify SSH is running:
`sudo systemctl status ssh`. Look for “active (running)” in the output.

## User Creation: `deployuser1`

Where password is disabled and `gecos` helps to ignore other fields and gives access to a particular folder only:

```bash
sudo adduser --disabled-password --gecos "" deployuser1

sudo mkdir -p /var/www/myapp
sudo chown -R deployuser:deployuser /var/www/myapp
sudo chmod -R 750 /var/www/myapp  # -> Owner   group   others
                                  #      7       5        0
```


## File Permission Number and Binary

| Number | Binary | Permissions | Description |
| :-- | :-- | :-- | :-- |
| 0 | 000 | --- | No permissions |
| 1 | 001 | --x | Execute only |
| 2 | 010 | -w- | Write only |
| 3 | 011 | -wx | Write + Execute |
| 4 | 100 | r-- | Read only |
| 5 | 101 | r-x | Read + Execute |
| 6 | 110 | rw- | Read + Write |
| 7 | 111 | rwx | Read + Write + Execute |

## Disabling Shell Access (SCP Allowed)

```bash
sudo usermod -s /usr/sbin/nologin deployuser1
```


## Network Setup

You can fix the IP address so that whenever you SSH, you don't need to change the IP address in the command every time.
Example used in Ubuntu TPLink network: `192.168.1.17`

## What is ufw?

**ufw (Uncomplicated Firewall)** is a command-line interface to manage iptables—the built-in Linux firewall. It simplifies firewall configuration for beginners and admins by providing an easy way to allow, deny, or limit network traffic.

### Why use ufw?

- **Security:** Block unwanted access (e.g., malicious IPs, ports).
- **Control:** Allow only specific services (e.g., SSH, HTTP).
- **Simplicity:** Easier than manually editing iptables rules.


#### Basic Concepts

- Incoming traffic: Connections to your machine.
- Outgoing traffic: Connections from your machine.
- Default policies: What to do when no rule matches (e.g., deny all).
- Rules: Specific permissions or blocks for IPs, ports, or services.


## What is CPU Time?

CPU time refers to the actual time for which a CPU core was actively working on a process or task (as opposed to being idle or waiting).

**Types of CPU time:**

- **User time:** Time spent running user-space processes (e.g., your Go app).
- **System time:** Time spent running kernel-space operations (e.g., system calls, driver operations).
- **Idle time:** CPU is not doing any task.
- **I/O wait:** Time spent waiting for I/O (like disk or network).


### How to Check CPU Time

Use the `top`, `htop`, or `vmstat` command.

**Example output from top:**

```
%Cpu(s):  5.5 us,  2.0 sy,  0.0 ni, 92.0 id,  0.3 wa,  0.0 hi,  0.2 si,  0.0 st
```

Here, `us` is user time, `sy` is system time, `id` is idle, `wa` is I/O wait.

### Why it matters (for DevOps/Infra)

- High system time may indicate kernel-level issues or too many syscalls.
- High user time shows CPU-bound processes (like compression or Go computations).
- Helps diagnose performance bottlenecks.


## What is Load Average?

Load average represents the average number of processes waiting for CPU over a period of time.

### How to check load average

Use either of:

```bash
uptime
# or
cat /proc/loadavg
```

**Example:**

```
load average: 1.23, 0.75, 0.60
```

These are 1-minute, 5-minute, and 15-minute averages.

#### How to interpret

Compare the numbers against the number of logical CPU cores.

- If you have 4 cores:
    - Load average of 4.00 means full capacity.
    - Load average of >4.00 = overloaded.
    - Load average of <4.00 = under-utilized.


#### Why it matters

- Helps in autoscaling decisions on cloud (e.g., AWS EC2 CPU alarms).
- Identifies CPU pressure and queuing of processes.
- Important for alerting in DevOps and observability tools like Prometheus.


## Difference Between MemFree and MemAvailable

These come from `/proc/meminfo` and `free -m`.


| Metric | Meaning |
| :-- | :-- |
| MemFree | Total **unallocated** RAM (not used by anything at all) |
| MemAvailable | Total memory that is **available for use**, including cache, buffers, and reclaimable memory |

### Why is MemAvailable often higher?

Linux uses unused memory as disk cache and buffers to speed up I/O.
This memory can be reclaimed if needed.
So even if MemFree is low, the system isn’t necessarily low on memory.

**Example:**

```bash
$ free -h
              total        used        free      shared  buff/cache   available
Mem:           7.7G        1.2G        300M        200M        6.2G        6.9G
```

- free: 300MB = very low
- available: 6.9GB = system is healthy, because 6.2GB is cache and reclaimable


### Why it matters

- Monitoring tools (like Prometheus node exporter) should use MemAvailable.
- Don’t panic just by looking at low MemFree — Linux is using RAM effectively.


## Cheat Sheet

| Term | Meaning | How to Use It |
| :-- | :-- | :-- |
| CPU Time | How long the CPU was active (user/system time) | Use `top`, `vmstat`, `pidstat` |
| Load Average | Processes waiting for CPU (1, 5, 15 min) | Use `uptime`, compare to core count |
| MemFree | Truly free memory, not used for anything | Rarely useful in isolation |
| MemAvailable | Free + cache + reclaimable memory | Use this for memory health |

## Linux Files

`/proc/stat`: A virtual file in Linux that contains statistics about the system, especially CPU usage.
It’s updated frequently. The line starting with cpu shows cumulative CPU time since boot, broken into different categories.
Should read -> https://www.linuxhowtos.org/System/procstat.htm

## Bash and Script File Notes

- `#!/bin/bash`
Shebang — Tells the system that this script should be run using the Bash shell.
- `function_name() {}`
Defines a function in Bash.
- `awk '/^cpu / {print ...}'`
Reads `/proc/stat`, finds the line that starts with `cpu `, and prints columns 2 to 8, which represent CPU time in various modes (user, system, idle, etc.).
- `read -r var1 var2 ...`
Reads space-separated input into multiple variables. `-r` tells Bash not to treat backslashes `\` as escape characters.
- `< <(command)`
Process substitution. Runs the command inside the parentheses in a sub-shell and provides its output as a file-like input.
- `sleep N`
Pauses script execution for N seconds. In this script, `sleep 1` gives a 1-second delay between two CPU readings.
- `echo`
Prints output to the terminal.
- `#`
Comment in Bash; everything after `#` on the line is ignored by the shell.


## Connect SSH to VSCode

In config file in `.ssh` write this:

```
Host ubuntu-mac
    HostName 192.168.1.17
    User tk
    IdentityFile ~/.ssh/ubuntu_mac_key
    IdentitiesOnly yes
```


## Bash Script to Get CPU Usage

https://github.com/tushar11kh/Self-Development/blob/main/week1/cpu_usage.sh

## SYSTEMD

### What is systemd?

systemd is the default init system on most modern Linux distributions (like Ubuntu, CentOS, Debian, Fedora).
It is responsible for booting the system, managing services, mounting filesystems, and handling processes.

#### 1. What is an Init System?

An init system is the first process (PID 1) started by the Linux kernel during boot.
Its job:

- Set up the system environment
- Start services (like SSH, networking, Docker)
- Manage dependencies between them
- Handle shutdown and reboot

Historically, systems used:

- SysVinit
- Upstart
- Now, mostly systemd


#### 2. Why is systemd Important?

| Feature | What it Means |
| :-- | :-- |
| Parallel Service Start | Speeds up boot by not starting services one by one |
| Dependency Mgmt | Starts services only when their dependencies are ready |
| Unified Tooling | `systemctl`, `journalctl`, etc. — consistent interface |
| Service Monitoring | Automatically restarts services if they crash |
| Logging | Uses built-in journal (no need for separate syslog) |

#### 3. Where Is systemd Configured?

| Type | Location |
| :-- | :-- |
| Default Unit Files | `/usr/lib/systemd/system/` |
| Custom/Override Unit Files | `/etc/systemd/system/` |
| Active Services at Boot | Linked in `/etc/systemd/system/default.target.wants/` |

#### 4. What Are systemd Units?

Units are configuration files that tell systemd how to start/manage something.


| Unit Type | Purpose | File Extension |
| :-- | :-- | :-- |
| `service` | Defines a system service | `.service` |
| `target` | Groups other units (like runlevels) | `.target` |
| `mount` | Manages a mount point | `.mount` |
| `socket` | Listens for sockets (like systemd-socket activation) | `.socket` |
| `timer` | Cron-like scheduling | `.timer` |

#### 5. Boot Process with systemd

Here's what happens during a typical boot:

- Kernel loads and executes `/sbin/init` → this is a symlink to systemd.
- systemd loads its unit files.
- It looks for its default target (usually default.target, linked to graphical.target or multi-user.target).
- It starts all the services required by that target and its dependencies.

To view your default boot target:

```bash
systemctl get-default
```

Common targets:


| Target | Description |
| :-- | :-- |
| `multi-user.target` | CLI-only, networking enabled |
| `graphical.target` | GUI + networking |
| `rescue.target` | Single-user mode |
| `poweroff.target` | System shutdown |
| `reboot.target` | System restart |

#### 6. Example: Service Unit File

File: `/etc/systemd/system/foobar.service`

```
[Unit]
Description=Run foobar service
After=network.target

[Service]
ExecStart=/usr/local/bin/foobar
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

| Section | What It Does |
| :-- | :-- |
| `[Unit]` | Describes the unit and when to start |
| `[Service]` | Defines how to start, stop, restart |
| `[Install]` | Used when enabling service (boot link) |

#### 7. Common systemctl Commands

| Action | Command |
| :-- | :-- |
| View all running units | `systemctl list-units` |
| Start a service | `sudo systemctl start nginx.service` |
| Stop a service | `sudo systemctl stop nginx.service` |
| Restart a service | `sudo systemctl restart nginx.service` |
| View service status | `systemctl status nginx.service` |
| Enable at boot | `sudo systemctl enable nginx.service` |
| Disable at boot | `sudo systemctl disable nginx.service` |
| View logs | `journalctl -u nginx.service` |

#### 8. Check if You're Using systemd

```bash
ps -p 1 -o comm=
```

Output will be:

```
systemd
```

Also, look for: `/usr/lib/systemd/`

#### 9. Key Concepts to Remember

| Concept | Meaning |
| :-- | :-- |
| `systemctl` | Command-line tool to manage systemd |
| `journalctl` | Log viewer for systemd services |
| `unit` | A resource systemd manages (service, mount, socket...) |
| `target` | A group of units (like runlevels) |
| `enable`/`disable` | Add/remove from default target |
| `WantedBy` | Used in `[Install]` section to link service to target |

#### 10. Pro Tips for DevOps

- 📁 Store custom unit files in `/etc/systemd/system/`
- 🛠 After editing a unit, always reload:

```bash
sudo systemctl daemon-reexec   # Full reload
sudo systemctl daemon-reload   # Reload unit files
```

- 🚀 Want to run a Go/CLI app on boot? Make a `.service` file!
- 🔄 Use `Restart=on-failure` to auto-recover crashed services.
- 🧪 Debug slow boot:

```bash
systemd-analyze blame
```


## Using a Go flag_cli Tool with systemd

| Step | Action |
| :-- | :-- |
| 1️⃣ | Build: `go build -o name_of_program_or_service main.go` |
| 2️⃣ | Move: `sudo mv name_of_prog /usr/local/bin/` |
| 3️⃣ | Create service: `/etc/systemd/system/name_of_prog.service` |
| 4️⃣ | Enable \& start: `systemctl enable/start` |
| 5️⃣ | Monitor: `systemctl status`, `journalctl` |

### Format of a .service file

A `.service` file is a text-based configuration file with three main sections:

- `[Unit]`      → Metadata and dependencies
- `[Service]`   → Execution settings (how to run it)
- `[Install]`   → When to start it (boot-time integration)


#### Sample

```
[Unit]
Description=My CLI system monitor tool
After=network.target

[Service]
ExecStart=/usr/local/bin/mycli --all
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

**Common Unit dependencies:**

- `network.target` — if your service needs networking
- `local-fs.target` — waits for filesystems to be ready
- `syslog.target` — if you rely on logging early

**Most common install values:**

- `multi-user.target` → boots in CLI mode (like runlevel 3)
- `graphical.target` → boots with GUI
- `default.target` → usually links to graphical or multi-user

