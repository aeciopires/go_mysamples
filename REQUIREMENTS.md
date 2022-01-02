<!-- TOC -->

- [Requirements](#requirements)
  - [General Packages](#general-packages)
- [Docker](#docker)
- [Go](#go)

<!-- TOC -->

# Requirements

## General Packages

Install the follow packages.

Debian/Ubuntu:

```bash
sudo apt-get install -y git curl vim
```

CentOS:

```bash
sudo yum install -y git curl vim
```

# Docker

Install Docker CE (Community Edition) following the instructions of pages.

CentOS: https://docs.docker.com/install/linux/docker-ce/centos/

Debian: https://docs.docker.com/install/linux/docker-ce/debian/

Ubuntu Server: https://docs.docker.com/install/linux/docker-ce/ubuntu/

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
VERSION=1.17.5

mkdir -p $HOME/go/bin $HOME/go/pkg $HOME/go/src

cd /tmp

curl https://dl.google.com/go/go$VERSION.linux-amd64.tar.gz -o go.tar.gz

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
