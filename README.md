# 🏗️ devctl

`devctl` is a Go-based CLI utility designed to streamline developer and SRE workflows by automating common tasks across Git, Kubernetes, networking, and beyond. It supports a plugin-based architecture, enabling seamless extensibility for custom use cases.

---

## 🔧 Features

- **Plugin Support** – Easily extend functionality through modular CLI plugins.
- **Network Tools** – Diagnose connectivity issues with fast net-check utilities.
- **Git Helpers** – Automate Git workflows like branch cleanup, squash commits, and more.
- **Kubernetes Utilities** – Rapid K8s context switching, resource summaries, and debugging helpers.
- **Extensible Architecture** – Build your own tools into the CLI using the plugin interface.

---

## 🗂 Project Structure
The project is organized into several directories, each serving a specific purpose. Below is a high-level overview of the structure:
```
devctl/
├── cmd/
│   └── devctl/          # Main CLI application entry point
│       └── main.go
├── internal/              # Internal packages (not intended for external use)
│   ├── netcheck/          # Network utility checks
│   ├── githelper/         # Git-related utilities
│   ├── kubehelper/        # Kubernetes-related utilities
│   └── ...                # Additional internal packages
├── plugins/               # Directory for CLI plugins
│   ├── hello-world/       # Sample plugin
│   └── ...                # Additional plugins
├── pkg/                   # Shared packages (can be imported by other projects)
│   └── ...                # Shared utility packages
├── configs/               # Configuration files (e.g., YAML, JSON)
├── scripts/               # Helper scripts (e.g., build, deploy)
├── Makefile               # Build automation
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```
---
## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://golang.org/doc/install)
- `make` (for build automation)

### Build & Run

```bash
git clone https://github.com/nijogeorgep/devctl.git
cd devctl
make build
./bin/devctl help
```

## 📦 Contributing
Contributions are welcome! To add new plugins or core features:

- Fork the repo

- Create a new branch

- Submit a pull request

Please follow the existing coding style and structure.

## 🛡 License
This project is licensed under the MIT License. See LICENSE for more details.

## 👥 Maintainers
@nijogeorgep – Core Developer

## 👥 Contributors
@nijogeorgep – Core Developer
