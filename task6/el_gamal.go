package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	minBits := 2048
	maxBits := 4096

	// Генеруємо випадкове просте число p
	p, err := generatePrime(minBits, maxBits)
	if err != nil {
		fmt.Println("Помилка генерації простого числа:", err)
		return
	}

	// Знаходимо примітивний корінь g
	g, err := findPrimitiveRoot(p)
	if err != nil {
		fmt.Println("Помилка знаходження примітивного кореня:", err)
		return
	}

	fmt.Println("Модуль p:", p)
	fmt.Println("Примітивний корінь g:", g)
}

// Генерує випадкове просте число довжиною від minBits до maxBits бітів
func generatePrime(minBits, maxBits int) (*big.Int, error) {
	for {
		// Генеруємо випадкове число розміром maxBits бітів
		p, err := rand.Prime(rand.Reader, maxBits)
		if err != nil {
			return nil, err
		}

		// Перевіряємо, чи розмір p задовольняє вимогам
		if p.BitLen() >= minBits {
			// Використовуємо вірогіднісний тест Міллера-Рабіна для перевірки простоти
			isPrime := p.ProbablyPrime(10)
			if isPrime {
				return p, nil
			}
		}
	}
}

// Знаходить примітивний корінь модуля p
func findPrimitiveRoot(p *big.Int) (*big.Int, error) {
	one := big.NewInt(1)

	// Шукаємо примітивний корінь перебором
	for g := big.NewInt(2); g.Cmp(p) < 0; g.Add(g, one) {
		isPrimitiveRoot := true

		// Перевіряємо, чи g генерує всі елементи поля
		for i := big.NewInt(1); i.Cmp(p) < 0; i.Add(i, one) {
			result := big.NewInt(0).Exp(g, i, p)
			if result.Cmp(one) == 0 {
				isPrimitiveRoot = false
				break
			}
		}

		if isPrimitiveRoot {
			return g, nil
		}
	}

	return nil, fmt.Errorf("примітивний корінь не знайдено для модуля p")
}
