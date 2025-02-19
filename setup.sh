#!/bin/bash

if [ ! -f .env ] || [ -z "$(grep ACME_EMAIL .env | cut -d= -f2)" ]; then
    echo "Error: ACME_EMAIL not configured in .env"
    echo "Please create .env and set ACME_EMAIL=your-email@domain.com"
    exit 1
fi

mkdir -p {traefik,letsencrypt,middleware}

# required SSL permissions
touch letsencrypt/acme.json
chmod 600 letsencrypt/acme.json

docker compose up -d

echo "Setup complete. Check logs with: docker compose logs -f"
