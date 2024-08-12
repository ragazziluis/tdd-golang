# 🙋‍ Implementando TDD na Prática - 1 [Luis Miranda]

## :mag: Introdução

Este repositório contém o código do tutorial "Golang Testing with TDD". O tutorial ensina como aplicar o Desenvolvimento Orientado a Testes (TDD) em Go, destacando o ciclo Red-Green-Refactor, a importância dos testes unitários e de integração, e boas práticas para manter um código modular e testável.

## :dart: Objetivo

O objetivo é fornecer uma implementação prática do tutorial, demonstrando a execução de exemplos, documentando os resultados e ampliando os comentários do código para explicar as técnicas e conceitos do TDD.

## :jigsaw: Estrutura do Repositório

- `src/`: Contém o código fonte dos exemplos do tutorial.
- `assets/`: Contém prints das execuções dos testes.
- `docs`: Arquivo de documentação.

Obs: deixei um *easter egg* em `assets/` :) 

## Passos Realizados

### 1. Criação do Repositório no GitHub

Criamos este repositório para hospedar os códigos do tutorial aqui. Acesse o artigo fonte do tutotial [aqui](https://williaminfante.medium.com/golang-testing-with-tdd-e548d8be776).

### 2. Execução dos Exemplos e Documentação

Todos os exemplos foram executados e documentados com prints das execuções, disponíveis no diretório `assets/`.

#### Exemplos Executados

- Exemplo 1: Soma de dois números
- Exemplo 2: Verificação de igualdade

### 3. Ampliação dos Comentários do Código

Os comentários foram expandidos para explicar as técnicas e conceitos do TDD. Veja um exemplo abaixo:

```go
package starter

import (
    "fmt"
    "net/http"
)

// SayHello retorna uma mensagem de boas-vindas de acordo com o nome do usuário (personalizada)
// Testa a função com diferentes nomes para garantir que a mensagem de boas-vindas está correta.
func SayHello(name string) string {
    return fmt.Sprintf("Hello %s. Welcome!", name)
}
```

## Conclusão

Este repositório fornece um guia prático sobre a aplicação do TDD em Go, com exemplos documentados e comentários ampliados para explicar as técnicas e conceitos. Seguindo este tutorial, você entenderá a importância dos testes e como manter um design de código limpo e testável.
