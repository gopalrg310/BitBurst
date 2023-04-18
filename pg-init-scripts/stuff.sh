#! /bin/bash

#set -euo pipefail

#cp -p ~/artem/dockercreation/ca/server.* /var/lib/postgresql/
#chown postgres:postgres /docker-entrypoint-initdb.d/ca/server.*
#chmod 0600 /docker-entrypoint-initdb.d/ca/server.key
#cd /var/lib/postgresql/
#openssl genrsa -des3 -out server.key 2048
#penssl rsa -in server.key -out server.key.insecure
#mv server.key server.key.secure
#mv server.key.insecure server.key

#openssl req -new -key server.key -out server.csr
#openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

#chown postgres.postgres server.*
#chmod 640 server.*
export PGSSLMODE=disable
#pg_restore --verbose --clean --no-acl --no-owner -U configman -d postgresvideoDev /docker-entrypoint-initdb.d/restore.dump
psql -p 5432 -U postgres -W bitburstasses -f /docker-entrypoint-initdb.d/restore.dump
