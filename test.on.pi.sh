#!/bin/bash
docker context use bob-remote
docker compose --env-file pi.env -f compose.remote.yml up --build
docker context use desktop-linux