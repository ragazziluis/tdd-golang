# 📚 Documentação da Atividade Ponderada - Testes TDD em Go

## :mag: Introdução

Este documento fornece uma visão geral dos testes realizados na ponderada de TDD em Go. Ele inclui a execução dos testes e os resultados obtidos. 

## :dart: Objetivo

O objetivo desta documentação é detalhar os testes realizados para garantir que o código esteja funcionando corretamente e para ilustrar a aplicação do Desenvolvimento Orientado a Testes (TDD).

## :jigsaw: Estrutura do Projeto

- **`src/`**: Contém o código fonte e os testes do projeto.
- **`assets/`**: Diretório reservado para prints das execuções dos testes.
- **`docs/`**: Diretório reservado para este arquivo de documentação.

## 📋 Testes Executados

Abaixo estão os detalhes dos testes realizados, com os prints dos resultados.

### 1. Teste da Função `SayHello`

A função `SayHello` retorna uma mensagem de boas-vindas personalizada. O teste verifica se a mensagem gerada está correta para diferentes inputs.

#### Código do Teste

```go
func TestSayHello(t *testing.T) {
    greeting := starter.SayHello("William")
    assert.Equal(t, "Hello William. Welcome!", greeting)

    anotherGreeting := starter.SayHello("asdf ghjkl")
    assert.Equal(t, "Hello asdf ghjkl. Welcome!", anotherGreeting)
}
```

#### Print dos Resultados

![test_sayhello](https://github.com/user-attachments/assets/7a4b54d0-826d-4747-92de-2a07ee00b7fb)

No teste de `SayHello/`, utilizamos a técnica de *assertion* para comparar a saída esperada com a saída real da função. Este é um exemplo padrão de TDD, onde primeiro escrevemos o teste que define o comportamento esperado e, em seguida, implementamos a função para que o teste passe.

### 2. Teste da Função `OddOrEven`

A função `OddOrEven` verifica se um número é par ou ímpar. Os testes cobrem números positivos, negativos e zero.

#### Código do Teste

```go
func TestOddOrEven(t *testing.T) {
    t.Run("Check Non Negative Numbers", func(t *testing.T) {
        assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
        assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
        assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
    })
    t.Run("Check Negative Numbers", func(t *testing.T) {
        assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
        assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
    })
}
```

#### Print dos Resultados

![test_oddoreven](https://github.com/user-attachments/assets/9d9c721a-16a5-4de5-b32f-78fe02f7a372)

No teste de `OddOrEven`, utilizamos a técnica de *test case grouping* com `t.Run` para agrupar testes relacionados. Isso ajuda a organizar os testes e a manter o código de teste limpo e legível. Cada caso de teste cobre um cenário específico, como números negativos e positivos, seguindo o ciclo TDD de escrever o teste antes de implementar a função.

### 3. Teste do Handler `CheckHealth`

O handler `CheckHealth` verifica o status de saúde do servidor. O teste assegura que o status HTTP e a resposta estejam corretos.

#### Código do Teste

```go
func TestCheckHealth(t *testing.T) {
    t.Run("Check health status", func(t *testing.T) {
        req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
        writer := httptest.NewRecorder()
        starter.CheckHealth(writer, req)
        response := writer.Result()
        body, err := io.ReadAll(response.Body)

        assert.Equal(t, "health check passed", string(body))
        assert.Equal(t, 200, response.StatusCode)
        assert.Equal(t,
            "text/plain; charset=utf-8",
            response.Header.Get("Content-Type"))
        assert.Equal(t, nil, err)
    })
}
```

#### Print dos Resultados

![test_checkhealth](https://github.com/user-attachments/assets/68544ebf-f52b-4d5f-8fb3-37e4b5b4e555)

No teste de `CheckHealth`, utilizamos a técnica de *HTTP testing* com as bibliotecas `httptest` para simular uma requisição HTTP e validar a resposta do servidor. Utilizamos este tipo de teste para garantir que os endpoints da aplicação estejam respondendo conforme esperado, especialmente em um contexto de TDD, onde o comportamento é definido antes da implementação.

## :memo: Conclusão

Esta documentação fornece um resumo dos testes realizados e dos resultados obtidos. A inclusão dos prints ajuda a verificar visualmente o sucesso dos testes e a garantir a qualidade do código.
