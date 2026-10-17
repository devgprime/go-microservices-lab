# Go microservices lab

An experimental multi-module Go application exploring HTTP services, RabbitMQ messaging, gRPC logging and container deployment. This is a learning/reference project, not a production-ready platform.

## Repository layout

| Module | Responsibility |
| --- | --- |
| `authentication-service` | Authentication and PostgreSQL access |
| `broker-service` | HTTP routing to authentication, mail and logging services |
| `listener-service` | RabbitMQ message consumption |
| `logger-service` | MongoDB-backed logging with HTTP/RPC/gRPC interfaces |
| `mail-service` | Email handling and rendering |

## Local verification

Install Go (CI uses Go 1.27) and Bash, then run from a fresh clone:

```bash
bash scripts/check.sh
# Optional convenience targets, if GNU Make is installed:
make test
make build
```

The check script runs tests and builds every module separately. Broker tests cover request parsing and invalid submissions without connecting to external services. Passing those checks does not verify the distributed deployment.

## Supporting containers

With Docker and the Compose plugin installed:

```bash
make up
make down
```

The checked-in Compose configuration currently starts **PostgreSQL and MongoDB only**. Application services, RabbitMQ and etcd entries are commented out. `make up` does not start a complete application. Kubernetes and Swarm files are historical deployment examples and need environment-specific validation. Demo database credentials are for isolated local development.

Service discovery is experimental: registration is disabled in some services. Do not assume the README implies a fully integrated discovery workflow.

## Benchmarking

```bash
bash scripts/benchmark.sh
```

This records the revision, Go/runtime environment and repeated broker request-decoding microbenchmarks under `benchmark-results/`. The results include allocation counts and time per operation. They include HTTP request/recorder construction and are **not** end-to-end latency or throughput measurements.

See [the roadmap](ROADMAP.md) for distributed-system tests, controlled load benchmarks and deployment automation still to be implemented.
