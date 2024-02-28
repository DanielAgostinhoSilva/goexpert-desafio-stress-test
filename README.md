# goexpert-desafio-tecnico

## Sistema de Stress test

`Objetivo`: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

`Entrada de Parâmetros via CLI:`: 

`--url`: URL do serviço a ser testado.
`--requests`: Número total de requests.
`--concurrency`: Número de chamadas simultâneas.

### Execução do Teste:

* Realizar requests HTTP para a URL especificada.
* Distribuir os requests de acordo com o nível de concorrência definido.
* Garantir que o número total de requests seja cumprido.

### Geração de Relatório:

* Apresentar um relatório ao final dos testes contendo:
  * Tempo total gasto na execução
  * Quantidade total de requests realizados.
  * Quantidade de requests com status HTTP 200.
  * Distribuição de outros códigos de status HTTP (como 404, 500, etc.).


1. Execução da aplicação:
* Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
  * docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10

# Resumo do projeto

Este código implementa testes de carga para um serviço Web. Ele realizará um número específico de solicitações a uma URL especificada, possivelmente de forma concorrente, e reportará os resultados.

## Estrutura do Código

O código contém uma estrutura chamada `StressTestReport` que mantém informações sobre os testes realizados. Isso inclui informações como o tempo total de execução, o número total de solicitações, o número de solicitações que resultaram em erro e uma distribuição dos códigos de status HTTP recebidos.

A estrutura `StressTestReport` tem os seguintes métodos:

- `Execute:` Este método é responsável por iniciar o teste de carga. Ele aceita a URL a ser testada, o número total de solicitações a serem feitas e o número de solicitações a serem feitas simultaneamente. Este método configura e dispara as goroutines necessárias para fazer as solicitações, processa os resultados e exibe uma barra de progresso para indicar o progresso do teste.

- `worker:` Este método é usado como a função de trabalho para as goroutines disparadas no método `Execute`. Ele faz um número específico de solicitações para a URL especificada e envia os resultados para o canal fornecido.

- `printReport:` Este método imprime o relatório do teste de carga. Ele exibe informações como o tempo total de execução, o número total de solicitações feitas, o número de solicitações que resultaram em erro e uma distribuição dos códigos de status HTTP recebidos.

- `displayProgressBar:` Este método atualiza e exibe uma barra de progresso com base no número de solicitações concluídas e no número total de solicitações.


## Como executar a aplicação

## Passos

1. **Build da imagem docker**
   Acessa a pasta raiz do projeto e execute o seguinte comando para fazer o build da imagem:

    ```shell
    docker build -t stress-test .
    ```

2. **Executar a aplicação**

   Executando via terminal :
   Acessa a pasta raiz do projeto e execute o seguinte comando:

    ```shell
    go run main.go test --url=http://google.com --requests=1000 --concurrency=100
    ```

   Executando via imagem docker :

    ```shell
    docker run stress-test ./main  test --url=http://google.com --requests=200 --concurrency=100
    ```


 