# 🙋‍ Implementando TDD na Prática - 1 [Luis Miranda]

## :mag: Introdução:

Este repositório contém os códigos do tutorial "Golang Testing with TDD", que explica como aplicar o Desenvolvimento Orientado a Testes (TDD) em Go. O tutorial descreve o ciclo Red-Green-Refactor, destaca a importância dos testes unitários e de integração, e fornece exemplos de código em Go. Também enfatiza as melhores práticas para escrever testes eficazes e manter um design de código modular e testável.

## :dart: O objetivo:

O objetivo deste repositório é fornecer uma implementação prática do tutorial, incluindo a execução de cada exemplo, a documentação dos resultados e a ampliação dos comentários do código para explicar as técnicas e conceitos do TDD.

## :jigsaw: Estrutura do Repositório

- `src/`: Diretório contendo os códigos fonte dos exemplos do tutorial.
- `screenshots/`: Diretório contendo prints das execuções dos testes.
- `README.md`: Este arquivo de documentação.

# Passos Realizados

## 1. Criação do Repositório no GitHub

Foi criado um repositório no GitHub para hospedar os códigos do tutorial. Você pode acessar o repositório aqui.  

## 2. Execução dos Exemplos e Documentação

Todos os exemplos do tutorial foram executados, e os resultados foram documentados por meio de prints das execuções. Estes prints estão disponíveis no diretório `screenshots/`.

### Exemplos Executados

Cada exemplo do tutorial foi executado usando o comando `go test/`. Abaixo estão os prints das execuções:

- Exemplo 1: Soma de dois números
- Exemplo 2: Verificação de igualdade

## 3. Ampliação dos Comentários do Código

Os comentários nos códigos foram ampliados para explicar as técnicas e conceitos do TDD utilizados. Abaixo está um exemplo de como os comentários foram expandidos: 

```go
package main

import (
    "testing"
)

// Soma retorna a soma de dois inteiros.
// Exemplo de uma função simples que vamos testar usando TDD.
func Soma(a int, b int) int {
    return a + b
}

// TestSoma testa a função Soma.
// No TDD, começamos escrevendo um teste que falha para definir o comportamento desejado.
func TestSoma(t *testing.T) {
    total := Soma(2, 3)
    esperado := 5

    if total != esperado {
        t.Errorf("Resultado da soma é incorreto, obtido: %d, esperado: %d", total, esperado)
    }
}

```

# Conclusão do Relatório

Este repositório apresenta um guia prático sobre como aplicar TDD em Go, com exemplos executados e documentados, além de comentários expandidos para explicar as técnicas e conceitos do TDD. Seguindo este tutorial, você aprenderá a importância dos testes unitários e de integração, e como manter um design de código modular e testável. 
