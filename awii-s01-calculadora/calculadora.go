package main

import "fmt"

func main() {
	var a int
	var b int
	var operacion string
	for {
		fmt.Println("===== CALCULADORA CIENTIFICA =====")
		fmt.Println("Ingrese el primer número: ")
		fmt.Scanln(&a)
		fmt.Println("Ingrese el segundo número: ")
		fmt.Scanln(&b)
		fmt.Print("Ingrese la operación (+, -, *, /, ^, !): ")
		fmt.Scanln(&operacion)
		switch operacion {
		case "+":
			resultado := a + b
			fmt.Printf("Resultado: %d + %d = %d", a, b, resultado)
		case "-":
			resultado := a - b
			fmt.Printf("Resultado: %d - %d = %d", a, b, resultado)
		case "*":
			resultado := a * b
			fmt.Printf("Resultado: %d * %d = %d", a, b, resultado)
		case "/":
			if b == 0 {
				fmt.Println("Error: No se puede dividir por cero.")
			} else {
				resultado := float64(a) / float64(b)
				fmt.Printf("Resultado: %d / %d = %.2f", a, b, resultado)
			}
		case "^":
			resultado := 1
			for i := 0; i < b; i++ {
				resultado = resultado * a
			}
			fmt.Printf("Resultado: %d ^ %d = %d", a, b, resultado)
		case "!":
			resultado := 1
			for i := 1; i <= a; i++ {
				resultado = resultado * i
			}
			fmt.Printf("Resultado %d! = %d", a, resultado)
		default:
			fmt.Println("Operación no válida.")
		}
		fmt.Println("\nDesea Continuar s/n")
		var respuesta string
		fmt.Scanln(&respuesta)

		if respuesta == "n" {
			println("Hasta luego")
			break
		}
	}
}
