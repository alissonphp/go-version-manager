#!/usr/bin/env bash
docker build -t alissonphp/go-version-manager .
docker stop go-version-manager && docker rm go-version-manager
docker run --name go-version-manager -p 8000:8000 --mount type=bind,source="$(pwd)",target=/app alissonphp/go-version-manager