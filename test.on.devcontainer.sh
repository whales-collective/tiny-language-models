#!/bin/bash
docker context use desktop-linux
docker compose -f compose.remote.yml up --build
