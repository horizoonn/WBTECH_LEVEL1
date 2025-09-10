package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

type Calculator struct {
	prec uint
}

func NewCalculator(num uint) *Calculator {
	return &Calculator{prec : num}
}

func (c *Calculator) Sum(num1, num2 string) (string, error) {
	a := new(big.Float).SetPrec(c.prec)
	b := new(big.Float).SetPrec(c.prec)

	_, okA := a.SetString(num1)
	_, okB := b.SetString(num2)

	if !okA || !okB {
		return "", fmt.Errorf("Переданы неверные числа\n")
	}

	res := new(big.Float).SetPrec(c.prec).Add(a, b)
	return res.Text('f', -1), nil
}

func (c *Calculator) Subtract(num1, num2 string) (string, error) {
	a := new(big.Float).SetPrec(c.prec)
	b := new(big.Float).SetPrec(c.prec)

	_, okA := a.SetString(num1)
	_, okB := b.SetString(num2)

	if !okA || !okB {
		return "", fmt.Errorf("Переданы неверные числа\n")
	}

	res := new(big.Float).SetPrec(c.prec).Sub(a, b)
	return res.Text('f', -1), nil
}

func (c *Calculator) Multiply(num1, num2 string) (string, error) {
	a := new(big.Float).SetPrec(c.prec)
	b := new(big.Float).SetPrec(c.prec)

	_, okA := a.SetString(num1)
	_, okB := b.SetString(num2)

	if !okA || !okB {
		return "", fmt.Errorf("Переданы неверные числа\n")
	}

	res := new(big.Float).SetPrec(c.prec).Mul(a, b)
	return res.Text('f', -1), nil
}

func (c *Calculator) Division(num1, num2 string) (string, error) {
	a := new(big.Float).SetPrec(c.prec)
	b := new(big.Float).SetPrec(c.prec)
	zero := new(big.Float).SetFloat64(0)

	_, okA := a.SetString(num1)
	_, okB := b.SetString(num2)

	if !okA || !okB {
		return "", fmt.Errorf("Переданы неверные числа\n")
	}

	if b.Cmp(zero) == 0 {
		return "", fmt.Errorf("Деление на 0 запрещено\n")
	}

	res := new(big.Float).SetPrec(c.prec).Quo(a, b)
	return res.Text('f', -1), nil
}

func process(calc *Calculator, input string) (string, error) {
	input = strings.TrimSpace(input)

	operators := []string{"+", "-", "*", "/"}

	for _, operator := range operators {
		index := strings.LastIndex(input, operator)
		if index > 0 {
			a := strings.TrimSpace(input[:index])
			b := strings.TrimSpace(input[index+1:])

			if a != "" && b != "" {
				switch operator {
				case "+" :
					return calc.Sum(a, b)
				case "-":
					return calc.Subtract(a, b)
				case "*":
					return calc.Multiply(a, b)
				case "/":
					return calc.Division(a, b)
				}
			}
		}
	}
	return "", fmt.Errorf("Неподдерживаемая операция\n")
}

func main() {
	calc := NewCalculator(200)

	log.Println("Калькулятор для больших чисел")
	fmt.Println("Поддерживаемые операции: +, -, *, /")
	fmt.Println("Для выхода введите exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите выражение:")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if input == "exit" {
			break
		}

		result, err := process(calc, input)
		if err != nil {
			fmt.Printf("Ошибка: %v", err)
			continue
		}

		fmt.Println("Результат:", result)
	}
}