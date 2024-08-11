package starter

import (
	"fmt"
	"net/http"
)

// SayHello retorna uma saudação personalizada
func SayHello(name string) string {
	return fmt.Sprintf("Hello %s. Welcome!", name)
}

// OddOrEven verifica se um número é par ou ímpar
func OddOrEven(number int) string {
	if number%2 == 0 {
		return fmt.Sprintf("%d is an even number", number)
	}
	return fmt.Sprintf("%d is an odd number", number)
}

// CheckHealth é um handler HTTP que retorna um status de saúde
func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("health check passed"))
}
