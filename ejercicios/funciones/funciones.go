package funciones

import "fmt"

// Pre: El parámetro coeficientes no debe ser nulo.
// Post: La cadena devuelta representa el polinomio formado por los coeficientes.
func ImprimirPolinomio(coeficientes []float64) string {
	if len(coeficientes) == 0 {
		return "El array está vacío"
	}

	a, b, c := fmt.Sprintf("%v", coeficientes[0]), fmt.Sprintf("%v", coeficientes[1]), fmt.Sprintf("%v", coeficientes[2])

	polinomio := "" + a + " + " + b + " x + " + c + " x**2"

	return polinomio
}
