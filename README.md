# RiskRancher Core

**DefectDojo without the Docker tax.** Air-gapped vulnerability management in a single Go binary — with a no-code Adapter Builder that maps any scanner export into tickets.

Apache License 2.0 · [Releases](https://github.com/Kuebiko-LLC/risk-rancher-core/releases) · [Website](https://www.riskrancher.com) · [Pricing](https://www.riskrancher.com/pricing)

## Why Core

- **Single binary + SQLite** — no Postgres, Redis, Celery, or Kubernetes. Download, run, open `http://localhost:8080`.
- **No-code scanner connectors** — upload Qualys, Nessus, Trivy, Dependabot, or any JSON/CSV export; map fields in the UI; ingest. No custom parsers.
- **100% air-gapped** — zero telemetry, zero outbound API calls. Your findings stay on your machine.

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

1. Register the first user (Sheriff / admin).
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

## Upgrade path

Core is free forever. When you need engagement reports (Auditor) or team automation (Pro), drop in the commercial binary — same SQLite database, zero ETL.

- [Auditor — $1,999/yr](https://www.riskrancher.com/auditor) — findings ↔ branded reports for pentesters
- [Pro — $4,999/yr](https://www.riskrancher.com/pro) — auto-assign, exceptions, suppressions, flat fee unlimited assets

## License

Apache License 2.0 — see [LICENSE](LICENSE).
