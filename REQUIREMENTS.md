<!-- TOC -->

- [Requirements](#requirements)
- [General Packages](#general-packages)
- [Docker](#docker)
- [Go](#go)

<!-- TOC -->

# Requirements

# General Packages

Install the follow packages.

* git
* curl
* vim

# Docker

Install Docker following the instructions of pages.

* CentOS: https://docs.docker.com/engine/install/centos/
* Debian: https://docs.docker.com/engine/install/debian/
* Ubuntu: https://docs.docker.com/engine/install/ubuntu/

Save the Docker service, configure the Docker to boot at the operating system and add the user to the Docker group.

```bash
sudo systemctl start docker

sudo systemctl enable docker

sudo usermod -aG docker $USER

sudo setfacl -m user:$USER:rw /var/run/docker.sock

docker version
```

For more information about Docker Compose visit:

* https://docs.docker.com
* http://blog.aeciopires.com/primeiros-passos-com-docker

# Go

Install Go with the follow commands.

```bash
VERSION=1.19.5

mkdir -p $HOME/go/bin

cd /tmp
curl -L https://go.dev/dl/go$VERSION.linux-amd64.tar.gz -o go.tar.gz

sudo rm -rf /usr/local/go 
sudo tar -C /usr/local -xzf go.tar.gz
rm /tmp/go.tar.gz

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

go version

echo "GOPATH=$HOME/go" >> ~/.bashrc
echo "PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> ~/.bashrc
```

For more information about Go visit:

* https://golang.org/doc/install#install
* https://golang.org/doc
* http://aprendago.com

For troubleshooting with Go modules:

```bash
go env -w GO111MODULE=off
```

Reference: https://stackoverflow.com/questions/66894200/go-go-mod-file-not-found-in-current-directory-or-any-parent-directory-see-go
