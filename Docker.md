# Docker

#### Comandos para executar um container
```shell

# Executando um container
docker run nginx

# atributos para executar um container
--name=nginx                      # nome do container
--rm                              # remove o container ao fazer o stop
-p 8080:80                        # mapeia a porta 8080 do nosso computador para a porta 80 do container
-d                                # roda o container em background
-v $(pwd):/usr/share/nginx/html   # Volume - compartilhando a pasta atual do computador com a pasta do container
-v ./data:/var/lib/mysql          # compartilhando uma pasta para manter o banco de dados

# Executando um container com todos os parametros
docker run --name=nginx --rm -d -p 8080:80 -v $(pwd):/usr/share/nginx/html nginx

```

#### Comandos para gerenciar imagens e containers
```shell

# Logando em um container ativo e acessando o bash do mesmo
docker exec -it container bash

## Listar os containers do docker
docker ps

# Listar todos os containers ativos e inativos
docker ps -aq

# Parando e removendo todos os containers
docker rm $(docker ps -aq)

# Listar as imagens
docker images -a

# Remover uma imagem
docker rmi imagem-a-ser-removida

# Remover todas as imagens
docker system prune -a

```

#### Criando uma imagem
Para criar uma imagem devemos criar um arquivo chamado **Dockerfile**
```shell

# Criando uma imagem a partir da imagem do nginx 
FROM nginx:latest

# Executando comando dentro do shell, atualizando os pacotes e a instalando o vim
RUN apt-get update && apt-get install -y vim

# Copiando o arquivo index.html que esta na minha máquina para o container
COPY index.html /usr/share/nginx/html

# Cria uma pasta no container
WORKDIR /go/src

```

>Para a criação da imagem temos também o comando **COMMAND** ele é utilizado para executar 
um comando no container para manter um processo rodando

#### Compilando a imagem
```shell
docker build -t arnaudsalves/nginx-fc:latest .

# -t = tag  definimos a tag e a versão da imagem
# Após a tag terminamos com ponto(.), pois o Dockerfile esta no diretório raiz caso estivesse em outro 
# diretorio teriamos que indica-lo
 
```

#### Subindo a imagem no docker hub
```shell
docker push nome-da-imagem:versao-imagem
```