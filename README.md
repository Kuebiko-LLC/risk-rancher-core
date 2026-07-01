# 🐴 RiskRancher Core (Community Edition)

RiskRancher Core is an open-source **Risk-Based Vulnerability Management (RBVM)** and **ASPM** platform built for modern 
DevSecOps teams. 

Compiled as a lightning-fast, **air-gapped single Go binary** with an embedded SQLite database, it ingests, deduplicates,
and routes millions of security findings from your CI/CD pipelines and scanners.

No external databases to spin up, no Docker swarms to manage, and zero complex microservices. 
Just drop the binary on a server and start triaging.

## Getting Started

### Option A: Download the Binary

1. Go to the [Releases](https://code.riskrancher.com/RiskRancher/core/releases) tab and download the compiled executable for your OS (Windows/macOS/Linux).
2. Place the binary in a dedicated directory and execute it.
3. Visit `http://localhost:8080` in your browser.

### Option B: Compile from Source

Ensure you have **Go 1.26+** installed (*CGO is required for the native `mattn/go-sqlite3` driver*).

```bash
git clone https://code.riskrancher.com/RiskRancher/core
cd core
go build -o rr ./cmd/rr/main.go
./rr
```

## License

Apache License 2.0