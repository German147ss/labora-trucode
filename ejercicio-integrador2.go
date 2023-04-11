package main

import (
	"fmt"
	"sort"
	"strings"
)

type Estudiante struct {
	nombre string
	nota   float64
	codigo string
}

func main() {
	estudiantes := make([]Estudiante, 0)
	var opcion int

	for {
		fmt.Println("\n1. Crear estudiante")
		fmt.Println("2. Ordenar estudiantes")
		fmt.Println("3. Buscar estudiante")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")

		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var nombre, codigo string
			var nota float64
			fmt.Print("Nombre: ")
			fmt.Scan(&nombre)
			fmt.Print("Nota: ")
			fmt.Scan(&nota)
			fmt.Print("Código: ")
			fmt.Scan(&codigo)
			estudiantes = append(estudiantes, Estudiante{nombre, nota, codigo})
		case 2:
			var criterio string
			fmt.Println("Ordenar por: (nombre, nota, codigo)")
			fmt.Scan(&criterio)
			estudiantesOrdenados := ordenarEstudiantes(estudiantes, criterio)
			mostrarEstudiantes(estudiantesOrdenados)
		case 3:
			var valorBusqueda string
			fmt.Print("Ingrese el valor a buscar: ")
			fmt.Scan(&valorBusqueda)
			estudianteEncontrado := buscarEstudiante(estudiantes, valorBusqueda)
			if estudianteEncontrado != nil {
				mostrarEstudiantes([]Estudiante{*estudianteEncontrado})
			} else {
				fmt.Println("No se encontró ningún estudiante con ese valor.")
			}
		case 4:
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func ordenarEstudiantes(estudiantes []Estudiante, criterio string) []Estudiante {
	estudiantesOrdenados := make([]Estudiante, len(estudiantes))
	copy(estudiantesOrdenados, estudiantes)

	switch criterio {
	case "nombre":
		sort.Slice(estudiantesOrdenados, func(i, j int) bool {
			return strings.ToLower(estudiantesOrdenados[i].nombre) < strings.ToLower(estudiantesOrdenados[j].nombre)
		})
	case "nota":
		sort.Slice(estudiantesOrdenados, func(i, j int) bool {
			return estudiantesOrdenados[i].nota < estudiantesOrdenados[j].nota
		})
	case "codigo":
		sort.Slice(estudiantesOrdenados, func(i, j int) bool {
			return strings.ToLower(estudiantesOrdenados[i].codigo) < strings.ToLower(estudiantesOrdenados[j].codigo)
		})
	default:
		fmt.Println("Criterio de ordenamiento no válido.")
	}

	return estudiantesOrdenados
}

func buscarEstudiante(estudiantes []Estudiante, valor string) *Estudiante {
	for _, estudiante := range estudiantes {
		if strings.ToLower(estudiante.nombre) == strings.ToLower(valor) || fmt.Sprintf("%.2f", estudiante.nota) == valor || strings.ToLower(estudiante.codigo) == strings.ToLower(valor) {
			return &estudiante
		}
	}
	return nil
}

func mostrarEstudiantes(estudiantes []Estudiante) {
	for _, estudiante := range estudiantes {
		fmt.Printf("Nombre: %s | Nota: %.2f | Código: %s\n", estudiante.nombre, estudiante.nota, estudiante.codigo)
	}
}
