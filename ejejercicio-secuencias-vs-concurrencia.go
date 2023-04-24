package main

import (
	"fmt"
	"sync"
	"time"
)

// Función sumar que recibe un slice de números, una referencia a la variable resultado
// y un puntero al grupo de espera (WaitGroup) para coordinar la concurrencia.
// 1 - 500000 // 500001 - 1000000
func sumar(numeros []int64, resultado *int64, wg *sync.WaitGroup) {
	defer wg.Done() // Al finalizar la función, se indica al WaitGroup que esta goroutine ha terminado.

	suma := int64(0)
	for _, num := range numeros {
		suma += num
	}
	*resultado = suma
}

func sumarSecuencial(numeros []int64) int64 {
	suma := int64(0)
	for _, num := range numeros {
		suma += num
	}
	return suma
}

func main() {
	datos := make([]int64, 1000000)
	for i := 0; i < len(datos); i++ {
		datos[i] = int64(i + 1)
	}
	tamanoFragmento := len(datos) / 2

	// Versión con goroutines
	inicio := time.Now()

	var wg sync.WaitGroup
	wg.Add(2) // Se indica al WaitGroup que se esperará a 2 goroutines.

	var suma1, suma2 int64
	// Se crea una goroutine para sumar la primera mitad de los datos.
	go sumar(datos[:tamanoFragmento], &suma1, &wg)
	// Se crea otra goroutine para sumar la segunda mitad de los datos.
	go sumar(datos[tamanoFragmento:], &suma2, &wg)

	wg.Wait() // Se espera a que ambas goroutines terminen de ejecutarse.

	sumaTotal := suma1 + suma2
	fmt.Printf("Suma total con goroutines: %d\n", sumaTotal)

	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución con goroutines: %s\n\n", tiempoTranscurrido)

	// Versión secuencial
	inicio = time.Now()
	sumaTotal = sumarSecuencial(datos)
	fmt.Printf("Suma total secuencial: %d\n", sumaTotal)

	tiempoTranscurrido = time.Since(inicio)
	fmt.Printf("Tiempo de ejecución secuencial: %s\n", tiempoTranscurrido)
}
