#!/bin/bash

echo "postgres setup start ..."

# create role
echo "creating role ..."
sudo -u postgres createuser --inherit --superuser --no-password dummydb
# create db
echo "creating db ..."
sudo -u postgres createdb dummydb
# create user
echo "creating user ..."
sudo adduser --disabled-password --gecos "" dummydb
# set password
echo "setting user password ..."
sudo -u dummydb psql -c "ALTER USER dummydb PASSWORD 'dummydbps';"
echo "postgresql ready <3"
