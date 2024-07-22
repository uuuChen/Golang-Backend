#!/bin/sh

docker-compose --env-file .env -f build/docker-compose.yml up --build -d