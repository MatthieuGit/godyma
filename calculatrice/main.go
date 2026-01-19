package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a, b int) (int, error) {
	return a + b, nil
}

func soustraction(a, b int) (int, error) {
	return a - b, nil
}

func multiplier(a, b int) (int, error) {
	return a * b, nil
}

func diviser(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division par Zéero impossible")
	}
	return a / b, nil
}

func main() {

	fmt.Println("--- Ma calculatrice en Go ---")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Saisissez votre opération (ou q pour quitter) : ")
		input, _ := reader.ReadString('\n')
		input = strings.ReplaceAll(input, "+", " + ")
		input = strings.ReplaceAll(input, "-", " - ")
		input = strings.ReplaceAll(input, "*", " * ")
		input = strings.ReplaceAll(input, "/", " / ")
		input = strings.TrimSpace(input)

		if input == "q" || input == "Q" {
			break
		}

		parts := strings.Split(input, " ")

		if len(parts) != 3 {
			fmt.Println("Erreur : format attendu A + B (n'oubliez pas les espaces)")
		}

		val1, err1 := strconv.Atoi(parts[0])
		val2, err2 := strconv.Atoi(parts[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Erreur : Veuillez entrer des nombres valides")
			continue
		}

		switch parts[1] {
		case "+":
			result, err := Sum(val1, val2)
			if err != nil {
				fmt.Println("Erreur lors du calcul :", err)
			} else {
				fmt.Println(result)
			}
		case "-":
			result, err := soustraction(val1, val2)
			if err != nil {
				fmt.Println("Erreur lors du calcul :", err)
			} else {
				fmt.Println(result)
			}
		case "*":
			result, err := multiplier(val1, val2)
			if err != nil {
				fmt.Println("Erreur lors du calcul :", err)
			} else {
				fmt.Println(result)
			}
		case "/":
			result, err := diviser(val1, val2)
			if err != nil {
				fmt.Println("Erreur lors du calcull :", err)
			} else {
				fmt.Println(result)
			}
		default:
			fmt.Println("l'operande saisie doit etre egale à : + - * / exemple (A + B)")
		}

	}

}
