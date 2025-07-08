# Deployment Guide

Docker compose files are provided to run the entire stack locally or in production.

## Development
Run the full stack including monitoring tools:
```bash
docker-compose -f deployment/docker-compose.yml up -d
```

## Production
A simplified production compose file is available:
```bash
docker-compose -f deployment/docker-compose.prod.yml up -d
```

For Kubernetes deployments see the manifests under `deployment/k8s/`.
