# VNPrider

VNPrider is a minimal Layer 1 blockchain written in Go.  The project provides a
simple reference implementation featuring a proof-of-authority style consensus
engine, in-memory storage and a small RPC server.  It is intended as a teaching
project rather than a production system.

## Setup

```bash
bash scripts/build.sh
bash scripts/test.sh
```

## Usage

See `cmd/vnprider-cli` and `cmd/vnprider-node` for entry points.
