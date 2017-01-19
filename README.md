## getCep

Projeto baseado no projeto do ![Elton Minetto](https://github.com/eminetto/goCep).

Agradecimento especial ao Elton pelas dicas \o/

## Compilar

Além de ter o Go instalado no sistema operacional é necessário executar:

`export GOPATH=/path/getCep`

`go get github.com/gorilla/mux`

`go build`

## Executar

O binário chamado getCep será criado. Basta executá-lo e ele ficará ouvindo na porta 8080 por novas requisições
Testando

`./getCep`

## Uso

Basta acessar a URL como no exemplo abaixo

`http://localhost:8080/15600000`

O retorno será um JSON com o conteúdo 
