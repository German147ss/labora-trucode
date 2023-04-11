package main

import (
	"fmt"
	"sort"
	"strings"
)

type Persona struct {
	nombre string
	edad   int
	altura int
	peso   int
}

func main() {
	personas := make([]Persona, 0, 5)
	for i := 1; i <= 5; i++ {
		var nombre string
		var edad, altura, peso int
		fmt.Printf("Ingrese los datos de la persona %d:\n", i)
		for {
			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			fmt.Print("Edad: ")
			fmt.Scan(&edad)
			fmt.Print("Altura (en cm): ")
			fmt.Scan(&altura)
			fmt.Print("Peso (en kg): ")
			fmt.Scan(&peso)

			if nombre != "" && edad > 0 && altura > 0 && peso > 0 {
				personas = append(personas, Persona{nombre, edad, altura, peso})
				break
			} else {
				fmt.Println("Datos inválidos, por favor ingrese los datos nuevamente.")
			}
		}
	}

	for {
		fmt.Println("\n1. Ordenar por nombre")
		fmt.Println("2. Ordenar por edad")
		fmt.Println("3. Ordenar por altura")
		fmt.Println("4. Ordenar por peso")
		fmt.Println("5. Buscar persona")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			personasOrdenadas := ordenarPersonas(personas, "nombre")
			mostrarPersonas(personasOrdenadas)
		case 2:
			personasOrdenadas := ordenarPersonas(personas, "edad")
			mostrarPersonas(personasOrdenadas)
		case 3:
			personasOrdenadas := ordenarPersonas(personas, "altura")
			mostrarPersonas(personasOrdenadas)
		case 4:
			personasOrdenadas := ordenarPersonas(personas, "peso")
			mostrarPersonas(personasOrdenadas)
		case 5:
			var valorBusqueda string
			fmt.Print("Ingrese el valor a buscar: ")
			fmt.Scan(&valorBusqueda)
			personaEncontrada := buscarPersona(personas, valorBusqueda)
			if personaEncontrada != nil {
				mostrarPersonas([]Persona{*personaEncontrada})
			} else {
				fmt.Println("No se encontró ninguna persona con ese valor.")
			}
		case 6:
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func ordenarPersonas(personas []Persona, criterio string) []Persona {
	personasOrdenadas := make([]Persona, len(personas))
	copy(personasOrdenadas, personas)

	switch criterio {
	case "nombre":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return strings.ToLower(personasOrdenadas[i].nombre) < strings.ToLower(personasOrdenadas[j].nombre)
		})
	case "edad":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return personasOrdenadas[i].edad < personasOrdenadas[j].edad
		})
	case "altura":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return personasOrdenadas[i].altura < personasOrdenadas[j].altura
		})
	case "peso":
		sort.Slice(personasOrdenadas, func(i, j int) bool {
			return personasOrdenadas[i].peso < personasOrdenadas[j].peso
		})
	default:
		fmt.Println("Criterio de ordenamiento no válido.")
	}

	return personasOrdenadas
}

func buscarPersona(personas []Persona, valor string) *Persona {
	for _, persona := range personas {
		if strings.ToLower(persona.nombre) == strings.ToLower(valor) || fmt.Sprint(persona.edad) == valor || fmt.Sprint(persona.altura) == valor || fmt.Sprint(persona.peso) == valor {
			return &persona
		}
	}
	return nil
}

func mostrarPersonas(personas []Persona) {
	for _, persona := range personas {
		imc := calcularIMC(persona.altura, persona.peso)
		categoria := clasificarIMC(imc)

		fmt.Printf("Nombre: %s | Edad: %d | Altura: %d cm | Peso: %d kg | IMC: %.2f (%s)\n", persona.nombre, persona.edad, persona.altura, persona.peso, imc, categoria)
	}
}

func calcularIMC(altura int, peso int) float64 {
	alturaMetros := float64(altura) / 100
	return float64(peso) / (alturaMetros * alturaMetros)
}

func clasificarIMC(imc float64) string {
	switch {
	case imc < 18.5:
		return "Bajo peso"
	case imc >= 18.5 && imc <= 24.9:
		return "Peso normal"
	case imc >= 25 && imc <= 29.9:
		return "Sobrepeso"
	default:
		return "Obesidad"
	}
}
