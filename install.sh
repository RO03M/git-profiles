#!/bin/sh
set -e

wget https://github.com/RO03M/git-profiles/releases/download/0.0.3/git-profiles_0.0.3_linux_386.tar.gz -O /tmp/gp.tar.gz -q --show-progress

echo "Asking permission to move the binary to /usr/local/bin, which need admin rights"

sudo tar -xzf /tmp/gp.tar.gz -C /usr/local/bin

sudo mv /usr/local/bin/git-profiles /usr/local/bin/gp -f

echo "Ready! Use gp to use the git profiles tool"