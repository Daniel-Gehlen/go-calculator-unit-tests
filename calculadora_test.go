package calculadora_test

import (
    "testing"

    "api-bolos-vovo/calculadora"
)

func TestSoma(t *testing.T) {
    resultado, err := calculadora.Soma(10, 5)
    if err != nil {
        t.Errorf("Soma falhou: %v", err)
        return
    }

    if resultado != 15 {
        t.Errorf("Soma esperada: 15, obtida: %v", resultado)
    }
}

func TestSubtrai(t *testing.T) {
    resultado, err := calculadora.Subtrai(10, 5)
    if err != nil {
        t.Errorf("Subtração falhou: %v", err)
        return
    }

    if resultado != 5 {
        t.Errorf("Subtração esperada: 5, obtida: %v", resultado)
    }
}

func TestMultiplica(t *testing.T) {
    resultado, err := calculadora.Multiplica(10, 5)
    if err != nil {
        t.Errorf("Multiplicação falhou: %v", err)
        return
    }

    if resultado != 50 {
        t.Errorf("Multiplicação esperada: 50, obtida: %v", resultado)
    }
}

func TestDivide(t *testing.T) {
    resultado, err := calculadora.Divide(10, 2)
    if err != nil {
        t.Errorf("Divisão falhou: %v", err)
        return
    }

    if resultado != 5 {
        t.Errorf("Divisão esperada: 5, obtida: %v", resultado)
    }

    // Teste de divisão por zero
    _, err = calculadora.Divide(10, 0)
    if err == nil {
        t.Errorf("Divisão por zero não deve retornar erro nulo")
    } else if err.Error() != "divisão por zero" {
        t.Errorf("Mensagem de erro de divisão por zero incorreta: %v", err)
    }
}
