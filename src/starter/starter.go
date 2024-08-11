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

// OddOrEven irá verificar se um número é par ou ímpar
// Testa a função com valores positivos, negativos e zero.
func OddOrEven(number int) string {
	if number%2 == 0 {
		return fmt.Sprintf("%d is an even number", number)
	}
	return fmt.Sprintf("%d is an odd number", number)
}

// CheckHealth retorna o status de saúde via HTTP.
// Testa a resposta para garantir que o status e o corpo da resposta estão corretos.
func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("health check passed"))
}
