package funciones

// Post: Se muestra un mensaje indicando la opción seleccionada o "Opción incorrecta" si la opción no es válida.
func elegirOpcion(opcion int) int {
	if opcion >= 1 && opcion <= 4 {
		return opcion
	}
	return 0
}
