package funciones

// MinMax devuelve el menor y el mayor número de una lista de enteros.
// Pre: La lista no debe estar vacía.
// Post: Devuelve el menor y el mayor número de la lista.
func EncontrarMinimoMaximo(lista []int) (int, int) {
	i := 0
	for i < len(lista) {
		j := i
		for j > 0 && lista[j-1] > lista[j] {
			swap(j-1, j, lista)
			j--
		}
		i++
	}
	minimo, maximo := lista[0], lista[len(lista)-1]
	return minimo, maximo
}

func swap(i int, j int, lista []int) {
	aux := lista[i]
	lista[i] = lista[j]
	lista[j] = aux
}
