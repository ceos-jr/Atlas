# Backend

Código do backend da aplicação Orb, a qual utiliza as seguintes tecnologias:

- Go
- Fiber (Framework Web para Go)
- GORM (ORM para Go)
- Postgres (Banco de dados relacional)
- Docker

## Desenvolvimento 

Para configurar o ambiente de desenvolvimento, siga os passos abaixo:

- Crie um diretório chamado **bin** para guardar os binários do projeto.
- Crie uma arquivo **.env** com base no arquivo **.env.example** fornecido.
- Baixe o docker 
- Inicie a imagem do banco de dados, com o comando: 

```
docker compose up db
```

## Compilação

Certifique-se de criar o diretório **bin**. Compile o backend com o seguinte comando: 

```
go build -o /bin/server .
```

Para rodar o servidor:

```
./bin/server
```

## Formatação 

Para manter o padrão do código durante o desenvolvimento utilizamos o linter **revive**

Para formatar o código use:
```
revive -config revive.toml -formatter friendly **/*.go
```
