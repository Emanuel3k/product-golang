# Projeto Go Web API

Este projeto é uma API web escrita em Go, que utiliza PostgreSQL para armazenamento de dados e manipulação de arquivos JSON para leitura e escrita de dados.

## Tecnologias Utilizadas

- **GoLang**: Linguagem de programação utilizada para desenvolver a API.
- **Postgres**: Banco de dados relacional utilizado para armazenamento de dados.
- **Docker**: Ferramenta de contêinerização utilizada para criar ambientes isolados.
- **Docker Compose**: Ferramenta para definir e gerenciar multi-contêineres Docker.
- **Chi Router**: Router utilizado para gerenciar as rotas da API.

## Arquitetura do Projeto

- **Injeção de Dependência**: Utilizada para facilitar a troca de implementações de serviços e repositórios.
- **Inversão de Controle**: Facilita a criação de testes unitários, permitindo maior flexibilidade e modularidade no código.
- **Camadas**: O projeto é dividido em camadas, seguindo o padrão de arquitetura Clean Architecture.
- **Packages**: O projeto é dividido em packages, seguindo o princípio do package orientd design (POD).

## Objetivo

Este projeto foi desenvolvido com fins de estudo, visando aprimorar conhecimentos em desenvolvimento de APIs web com Go e boas práticas de programação.