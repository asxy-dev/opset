<<<<<<< HEAD
# opset
Opset is a lightweight CLI automation tool that lets you run shell commands on multiple remote servers via SSH at scheduled intervals. Built in Go for speed and concurrency, with Java helpers for config generation and validation, Opset makes server task automation simple, reliable, and portable.
=======
# Opset (starter layout)

This repository contains a minimal starter layout for the "opset" project.

Structure created:

- `go/` - Go sources (CLI entrypoint, scheduler, SSH exec, storage, models)
  - `main.go` - CLI entrypoint (starter)
  - `scheduler.go` - scheduling placeholder
  - `ssh_exec.go` - SSH exec placeholder
  - `storage.go` - JSON load/save helpers
  - `models.go` - `Server` and `Task` structs
- `java/` - Java helper classes (optional)
  - `OpsetHelper.java` - JSON read/write helper (uses org.json)
  - `ConfigCheck.java` - lightweight validator placeholder
- `scripts/`
  - `install.sh` - builds Go binary and compiles Java helpers
  - `demo.sh` - shows sample demo steps
- `data/`
  - `servers.json` - sample servers list
  - `tasks.json` - sample tasks list

Quick start

1. Make the scripts executable:

```bash
chmod +x scripts/*.sh
```

2. Build and compile helpers:

```bash
./scripts/install.sh
```

3. Run the CLI (starter placeholder):

```bash
./opset -config data/servers.json
```

Notes

- The Go files are minimal placeholders to get you started. Implement SSH using `golang.org/x/crypto/ssh` or an external tool.
- The Java helpers reference `org.json` — add the library to your classpath if you plan to use them.
>>>>>>> 2f54c9a (file created)
