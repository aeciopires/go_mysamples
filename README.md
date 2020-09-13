# go_mysamples

[![Codeac.io](https://static.codeac.io/badges/2-231147601.svg)](https://app.codeac.io/github/aeciopires/go_mysamples)

<!-- TOC -->

- [go_mysamples](#go_mysamples)
- [Português](#português)
- [English](#english)
- [Tutorials](#tutorials)
  - [Learn Go](#learn-go)
  - [Go in Docker](#go-in-docker)
  - [Install Go from source](#install-go-from-source)
  - [Developers](#developers)
  - [License](#license)

<!-- TOC -->


# Português

Meus primeiros códigos desenvolvidos com a liguagem [Go](https://golang.org).

* `src`: contém o código fonte, bem como o código fonte de pacotes de terceiros importados nos programas;

# English

My first codes developed with the [Go](https://golang.org).

* `src`: contains the source code and source code of third party packages imported into programs;

# Tutorials

## Learn Go

Tutorials and documentation about **Go** language:

* http://aprendago.com
* https://golang.org/doc
* https://www.geeksforgeeks.org/golang/?ref=lbp
* https://astaxie.gitbooks.io/build-web-application-with-golang/content/pt-br/
* https://medium.com/@denis_santos/primeiros-passos-com-golang-c3368f6d707a
* https://sensedia.com/conteudo/introducao-ao-golang/
* https://sensedia.com/conteudo/motivos-para-utilizar-go
* https://sensedia.com/conteudo/estrutura-da-linguagem-go
* https://www.linode.com/docs/development/go/install-go-on-ubuntu/
* https://tecadmin.net/install-go-on-macos/
* https://sourabhbajaj.com/mac-setup/Go/README.html
* https://tour.golang.org
* https://www.youtube.com/watch?v=YS4e4q9oBaU
* https://www.youtube.com/watch?v=Q0sKAMal4WQ
* https://www.youtube.com/watch?v=G3PvTWRIhZA
* https://www.youtube.com/watch?v=_MkQLDMak-4
* https://www.youtube.com/watch?v=JepHr8egvBI
* https://stackify.com/learn-go-tutorials
* https://golangbot.com
* https://www.tutorialspoint.com/go/index.htm
* https://www.guru99.com/google-go-tutorial.html
* http://www.golangbr.org/doc/codigo
* https://golang.org/doc/articles/wiki
* https://gobyexample.com
* https://hackr.io/tutorials/learn-golang
* https://hackernoon.com/basics-of-golang-for-beginners-6bd9b40d79ae
* https://medium.com/hackr-io/learn-golang-best-go-tutorials-for-beginners-deb6cab45867
* https://github.com/dariubs/GoBooks#readme
* https://www.digitalocean.com/community/books/how-to-code-in-go-ebook-pt
* https://dzone.com/articles/golang-tutorial-learn-golang-by-examples
* https://dzone.com/articles/structure-of-a-go-program 
* https://gobyexample.com/
* https://awesome-go.com/
* https://research.swtch.com/
* https://levelup.gitconnected.com/get-a-taste-of-concurrency-in-go-625e4301810f

## Go in Docker

Use Go in Docker: https://hub.docker.com/_/golang

```
git clone https://github.com/aeciopires/go_mysamples

cd go_mysamples/src/helloworld

VERSION=1.15-alpine
GO_WORKSPACE=/usr/local/go/src
docker run --rm -v $GO_WORKSPACE/:/go/src/ -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:$VERSION go build -v

./myapp
```

This will add your current directory as a volume to the container, set the working directory to the volume, and run the command go build which will tell go to compile the project in the working directory and output the executable to myapp.

## Install Go from source

Install **Go** from source. See [REQUIREMENTS](REQUIREMENTS.md#Go):


Executing programs with **go**. Example:

```
git clone https://github.com/aeciopires/go_mysamples

cd go_mysamples/src/helloworld

go build helloworld.go

./helloworld
```

## Developers

developer: Aécio dos Santos Pires<br>
mail: http://blog.aeciopires.com/contato

## License

GPL-3.0 2020 Aécio dos Santos Pires


