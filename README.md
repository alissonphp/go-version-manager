# Go Version Manager
Repositório de versões de pacotes com persistência no sistema de arquivos local escrito em golang.

## Propósito
Ter um serviço REST onde possa ser possível enviar artefatos (plugins, libs, pacotes, etc.) e gerenciar suas versões de forma prática, garantindo performance e disponibilidade.
## Tecnologias
* [Golang](https://golang.org/project)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Hashicorp Go version](https://github.com/hashicorp/go-version)
* [swaggo](https://github.com/swaggo/swag)
* [Air](https://github.com/cosmtrek/air)
## Imagem de Desenvolvimento
* Faça o build da imagem de desenvolvimento
```shell
docker build -t alissonphp/go-version-manager .
```
* Suba o container da aplicação
```shell
docker run --name go-version-manager -p 8000:8000 --mount type=bind,source="$(pwd)",target=/app alissonphp/go-version-manager
```

O diretório da aplicação está em modo bind com o container, ou seja, cada alteração feita no fonte é refletida no container e o serviço é reiniciado (live reload) chamando o gerador do swagger.

P.s.: o script `./start.sh` faz o rebuild da imagem, para o container, remove, e sobe um novo.

* o [swagger](http://localhost:8000/docs/index.html#/) é acessível em `http://localhost:8000/docs/index.html#/`

## Imagem de Produção
* Faça o build da imagem de produção
```shell
docker build -t alissonphp/go-version-manager:production . -f Dockerfile.prod
```
* Suba o container
```shell
docker run -d --name go-version-manager-prod -p 80:8000 alissonphp/go-version-manager:production
```
P.s.: os arquivos são persistidos no diretório `./downloads/` para ambiente de produção, é recomendável que seja criado um volume no docker e seja feito o bind com o diretório interno do container, assim os arquivos não são perdidos no restart do pod/container

## Roadmap
* [ ] adicionar integração com serviços de storage (s3, azure files, etc)
* [ ] testar as unidades
* [ ] implementar esteira de integração com análise estática de código