package calculadora

import (
    "errors"
)

// Soma: Retorna a soma de dois números
func Soma(a, b float64) (float64, error) {
    return a + b, nil
}

// Subtrai: Retorna a diferença de dois números
func Subtrai(a, b float64) (float64, error) {
    return a - b, nil
}

// Multiplica: Retorna o produto de dois números
func Multiplica(a, b float64) (float64, error) {
    return a * b, nil
}

// Divide: Retorna o quociente da divisão de dois números
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}
