# Projeto API REST em Go com Concorrência para Consulta de CEP

Este projeto consiste em uma API REST em Go que utiliza concorrência para consulta de endereços através de CEP. Ele bate em duas rotas distintas para verificar o endereço da pessoa, escolhendo a resposta mais rápida entre elas. Além disso, há um mecanismo de timeout configurado para 3 segundos, caso nenhuma resposta seja recebida dentro desse período.

## Estrutura do Projeto

O projeto é estruturado da seguinte forma:

### Core

#### Controller
- Responsável pela comunicação entre o cliente e o serviço.
  
#### Business
- Contém a lógica de negócio da aplicação.

### Infrastructure

#### Server
- Configurações do servidor e roteamento de requisições.

#### Dependency
- Gerenciamento de dependências.

### Constants
- Constantes utilizadas no projeto.

### Model
- Contém as definições das estruturas de dados utilizadas em toda a aplicação. Essas estruturas representam os diferentes tipos de dados manipulados pela API, como informações de endereço.

### Service
- É responsável por facilitar a comunicação com serviços externos, como APIs de terceiros.

### main.go
- Ponto de entrada da aplicação.

## Funcionamento

A API é acessada através de endpoints específicos para consulta de endereço por CEP. Ao receber uma requisição, a API dispara concorrentemente consultas em duas APIs distintas. A resposta mais rápida é escolhida e retornada ao cliente. Caso nenhuma resposta seja recebida dentro de 3 segundos, a requisição é encerrada com um timeout.

## Como Executar

Para executar o projeto, siga os passos abaixo:

1. Certifique-se de ter o Go instalado em seu sistema.
2. Clone este repositório.
3. Navegue até o diretório do projeto.
4. Execute o comando `go run main.go` para iniciar o servidor.

## Exemplos de Uso

### Consulta de Endereço por CEP

**Endpoint:** `/api/v1/address/{CEP}`

**Método:** GET

**Exemplo de Requisição:**

GET /api/v1/address/12345678900 - HTTP/1.1
Host: http://localhost:8080


**Exemplo de Resposta:**
```json
{
  "cep": "12345678",
  "logradouro": "Rua Exemplo",
  "bairro": "Bairro Exemplo",
  "cidade": "Cidade Exemplo",
  "estado": "EX"
}
```

# Obrigado a todos espero que tenham gostado (: 
