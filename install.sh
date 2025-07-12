wget https://github.com/RO03M/git-profiles/releases/download/0.0.3/git-profiles_0.0.3_linux_386.tar.gz -O /tmp/gp.tar.gz -q

sudo tar -xzf /tmp/gp.tar.gz -C /bin

sudo mv /bin/git-profiles /bin/gp -f

echo "Ready! Use gp to use the git profiles tool"