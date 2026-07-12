# RiskRancher Core

Open-source **vulnerability management** in a single Go binary.

Ingest findings from any scanner (Qualys, Nessus, Trivy, Dependabot, or custom JSON/CSV), map fields in a no-code Adapter Builder, and track remediation in one air-gapped ticket dashboard — no Postgres, Redis, Docker Compose, or Kubernetes required.

**DefectDojo without the Docker tax.**

Apache License 2.0 · [Releases](https://github.com/Kuebiko-LLC/risk-rancher-core/releases) · [Website](https://www.riskrancher.com) · [Pricing](https://www.riskrancher.com/pricing)

## What it is

RiskRancher Core is a self-hosted vulnerability management platform for security engineers and pentesters who are tired of:

- Standing up heavy stacks just to store scanner output
- Writing one-off scripts to normalize every new tool’s JSON
- Spreadsheets as the “system of record” for findings

Drop one binary on a laptop or air-gapped server, open the UI, build a connector, and start triaging.

## Why teams use it

- **Single binary + SQLite** — download, run, open `http://localhost:8080`. No microservices.
- **No-code scanner connectors** — upload any export, map title / asset / severity in the UI, ingest. No custom parsers.
- **100% air-gapped** — zero telemetry, zero outbound API calls. Findings stay on your machine.
- **Ticket workflow built in** — asset grouping, severities, SLAs, and a practitioner dashboard so work actually gets done.

## Quick start

### 1. Download a release

Grab the binary for your OS from [Releases](https://github.com/Kuebiko-LLC/risk-rancher-core/releases/latest):

| Platform | Asset |
|----------|--------|
| Linux (amd64) | `rr-linux-amd64` |
| macOS (Apple Silicon) | `rr-darwin-arm64` |
| Windows (amd64) | `rr-windows-amd64.exe` |

```bash
chmod +x rr-linux-amd64   # or rr-darwin-arm64
./rr-linux-amd64
# open http://localhost:8080
```

### 2. Ingest findings in five minutes

**Option A — sample Trivy file**

1. Register the first user (admin).
2. Go to **Ingest** → select the built-in **Trivy Container Scan** adapter (or create one).
3. Upload [`examples/trivy-sample.json`](examples/trivy-sample.json).
4. Open the ticket dashboard — findings are grouped by asset.

**Option B — your own scanner (Qualys, Nessus, etc.)**

1. Go to **Build New Adapter** (`/admin/adapters/new`).
2. Drop in a sample JSON or CSV export.
3. Map title, asset, severity (and optional description / remediation).
4. Save the adapter and ingest the full export — no code required.

### 3. Build from source

Requires **Go 1.26+** (pure Go SQLite via `modernc.org/sqlite` — no CGO).

```bash
git clone https://github.com/Kuebiko-LLC/risk-rancher-core.git
cd risk-rancher-core
go build -o rr ./cmd/rr/main.go
./rr
```

## Who it’s for

- Security engineers drowning in scanner noise
- Pentesters who need findings in a real tracker (not only a Word report)
- Teams that need on-prem / air-gapped vulnerability management without a DevOps project

## Upgrade path

Core is free forever (Apache 2.0). When you need engagement reports or team automation, drop in the commercial binary — same SQLite database, zero ETL.

- [Auditor — $1,999/yr](https://www.riskrancher.com/auditor) — findings ↔ branded reports for pentesters
- [Pro — $4,999/yr](https://www.riskrancher.com/pro) — auto-assign, exceptions, suppressions, flat fee unlimited assets

## License

Apache License 2.0 — see [LICENSE](LICENSE).
