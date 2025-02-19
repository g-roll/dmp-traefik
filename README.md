# dmp-traefik

Public because of fubar. Should **NOT** be used in prod by anyone. Really.

A dead-simple reverse proxy setup for Docker container web apps with DNS
validation middleware that ensures DNS propagation before SSL certificate
requests.

## Prerequisites

- Docker
- Docker Compose

## Quick Start

1. Clone repository
2. Configure in `.env`:
   ```bash
   ACME_EMAIL=your-email@domain.com
   SERVER_IP=your-server-ip
   ```
3. Run setup:
   ```bash
   ./setup.sh
   ```

## Configuration

### Environment Variables

| Variable   | Description       | Required |
| ---------- | ----------------- | -------- |
| ACME_EMAIL | LetsEncrypt email | Yes      |
| SERVER_IP  | Server IP         | Yes      |

### Adding Domains

Add containers to the `proxy` network and configure Traefik labels:

```yaml
services:
    webapp:
        networks:
            - proxy
        labels:
            - "traefik.enable=true"
            - "traefik.http.routers.webapp.rule=Host(`example.com`)"
```
