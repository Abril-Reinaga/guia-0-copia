package estructurasrepetitivas

// Escribir una función que recibe un número entero n, no negativo, y devuelve su factorial.
// Pre: n >= 0
// Post: Devuelve el factorial de n.
func Factorial(n int) int {
	if n < 0 {
		panic("Debe ser un número positivo o 0")
	}

	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado
}
