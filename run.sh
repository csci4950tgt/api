#!/bin/sh

# Start the proxy
./cloud_sql_proxy -instances=$CLOUDSQL_INSTANCE=tcp:5432 -credential_file=$CLOUDSQL_CREDENTIALS &

# wait for the proxy to spin up
while ! nc -z localhost 5432; do
  sleep 0.1 # wait for 1/10 of the second before check again
done
# sleep 10

# Start the server
./api