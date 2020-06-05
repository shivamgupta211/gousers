#!/bin/bash

# docker hub username i.e. shivamgupta211
export DOCKER_HUB_USERNAME="shivamgupta211"

# server env "production"
export SERVER_ENV="production"


# ----------------- DATABASE related ENVs -----------------
# databse connection url for psql i.e. localhost:5432
export DB_CONNECTION_URL=""

# databse user for psql
export DB_USER=""

# databse user's password for psql
export DB_PASSWORD=""

# databse name for psql
export DB_DATABASE=""

# ----------------- Sentry related ENVs -----------------
# Sentry link for ravevn-go to push errors
export SENTRY=""
