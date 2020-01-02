# go_mysamples

[![Codeac.io](https://static.codeac.io/badges/2-231147601.svg)](https://app.codeac.io/github/aeciopires/go_mysamples)

[Português]: #português
[English]: #english
[Contributing]: #contributing
[Developers]: #developers
[Tutorials]: #tutorials
[License]: #license

#### Menu

1. [Português][Português]
2. [English][English]
3. [Contributing][Contributing]
4. [Developers][Developers]
5. [Tutorials][Tutorials]
6. [License][License]

# Português

Minhas amostras de binários desenvolvidos com a liguagem Go.

* `src`: contém o código fonte, bem como o código fonte de pacotes de terceiros importados nos programas;

## Contribuindo

1. Instale os pacotes: `git`, `go`, `curl`. Ou use Go com Docker: https://hub.docker.com/_/golang
2. Clone o repositório em seu computador com o comando: `git clone https://github.com/aeciopires/go_mysamples`
3. Acesse a pasta do projeto com o comando: `cd go_mysamples`
4. Crie uma branch usando o padrão: `git checkout -b US-${DEV_NAME}`. Exemplo: *git checkout -b US-AECIO*
5. Desenvolva a tarefa
6. Execute o código nos ambientes de desenvolvimento e teste
7. Faça commit e push dos arquivos para o repositório remoto com o comando: `git push --set-upstream origin US-${DEV_NAME}`. Exemplo: *git push --set-upstream origin US-AECIO*
8. Crie o Pull Request (PR) para a branch `master`, usando o padrão de notação do processo de desenvolvimento

**ATENÇÂO:** Antes de iniciar a sua contribuição, execute o comando: `git pull origin master` para obter o conteúdo mais recente da branch master e evite conflitos que possam fazer você perder tempo.

# English

My samples of binaries developed with Go language.

* `src`: contains the source code, as well as the source code of third party packages imported into programs;

## Contributing

1. Install packages: `git`, `go`, `curl`. Or use Go in Docker: https://hub.docker.com/_/golang
2. Clone the repository to your machine through the command: `git clone https://github.com/aeciopires/go_mysamples`
3. Go to the project folder: `cd go_mysamples`
4. Create a branch using the pattern: `git checkout -b US-${DEV_NAME}`. Example: *git checkout -b US-AECIO*
5. Develop the task
6. Execute the code in 'development' and 'test' environments
7. Do commit and push files to repository remote with command `git push --set-upstream origin US-${DEV_NAME}`. Example: *git push --set-upstream origin US-AECIO*
8. Create Pull Request (PR) to the `master` branch, using the notations patterns of the development process

**WARNING:** Before start to contribute, run the command: `git pull origin master` to fetch the newest content of the main branch and avoid conflicts that can make you waste time.

## Developers

developer: Aécio dos Santos Pires<br>
mail: http://blog.aeciopires.com/contato

## Tutorials

Use Go in Docker: https://hub.docker.com/_/golang

```
git clone https://github.com/aeciopires/go_mysamples

cd go_mysamples/src/helloworld

VERSION=1.13.5-alpine
GO_WORKSPACE=/usr/local/go/src
docker run --rm -v $GO_WORKSPACE/:/go/src/ -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:$VERSION go build -v

./myapp
```

This will add your current directory as a volume to the container, set the working directory to the volume, and run the command go build which will tell go to compile the project in the working directory and output the executable to myapp.

Or

Install go from source:

```
VERSION=1.13.5

curl https://dl.google.com/go/go$VERSION.linux-amd64.tar.gz -o go.tar.gz

sudo tar -C /usr/local -xzf go.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version
```

Executing programs with go. Exemple:

```
git clone https://github.com/aeciopires/go_mysamples

cd go_mysamples/src/helloworld

go run helloworld.go
```

Tutorials and documentation about go language:

* https://medium.com/@denis_santos/primeiros-passos-com-golang-c3368f6d707a
* https://tour.golang.org
* https://www.youtube.com/watch?v=YS4e4q9oBaU
* https://www.youtube.com/watch?v=Q0sKAMal4WQ
* https://www.youtube.com/watch?v=G3PvTWRIhZA
* https://stackify.com/learn-go-tutorials/
* https://golangbot.com/
* https://www.tutorialspoint.com/go/index.htm
* https://www.guru99.com/google-go-tutorial.html
* http://www.golangbr.org/doc/codigo
* https://golang.org/doc/articles/wiki/
* https://gobyexample.com/
* https://hackr.io/tutorials/learn-golang
* https://hackernoon.com/basics-of-golang-for-beginners-6bd9b40d79ae
* https://medium.com/hackr-io/learn-golang-best-go-tutorials-for-beginners-deb6cab45867


## License

GPL-3.0 2020 Aécio dos Santos Pires


