package estructurasrepetitivas

// Escribir una función que indique si un número entero, ingresado por un usuario, es primo (devuelve verdadero o falso).
// Pre: n > 0
// Post: Devuelve verdadero si n es primo, falso en caso contrario.

func EsPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true

}
