#!/bin/bash
set -o pipefail # abort on errors in pipeline
set -e          # abort on errors

# download, for ubuntu 18.04 bionic
wget -O aerospike.tgz 'https://www.aerospike.com/download/server/latest/artifact/ubuntu18'
# untar
tar xvzf aerospike.tgz # will extract: aerospike-server-community-5.0.0.4-ubuntu18.04
# install
cd $(ls | grep "aerospike-server-community")
sudo ./asinstall
# start aerospike
sudo systemctl start aerospike
# go back to source
cd ..
