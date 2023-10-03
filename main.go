package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//AbrirArchivo()
	//CrearEscribir()
	LeerArchivo()
	Crear_Escribir()
	Abrir_Leer()
}

// abre el archivo
func AbrirArchivo() {
	file, err := os.Open("./hola.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

// crea y escribe el archivo
func CrearEscribir() {
	file, err := os.Create("./hola.txt")
	if err != nil {
		panic(err)

	}
	defer file.Close()

	text := "hola com oestas "
	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}
}

// lee el archivo
func LeerArchivo() {
	file, err := os.Open("hola.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes := make([]byte, 1024)
	//.Read() lee el archivo "pero en bytes ""
	n, err := file.Read(bytes) // n es la cantidad de bytes q hay el el archivo
	if err != nil {
		panic(err)
	}
	fmt.Printf("Leí %d bytes: %s\n", n, string(bytes[:n]))

}

func BorraArchivo() {
	err := os.Remove("hola.txt")
	if err != nil {
		panic(err)
	}

}

// con archivo json
type Persona struct {
	Name       string `json:"name"`
	Edad       int    `json:"edad"`
	Completado bool   `json:"completado"`
}

func Abrir_Leer() {
	//abrir el archivo
	file, err := os.Open("persona.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//leer el archivo
	personas := make(map[string]interface{})

	manejo := json.NewDecoder(file)
	err = manejo.Decode(&personas)
	if err != nil {
		panic(err)
	}
	fmt.Println(personas)
	for i, conte := range personas {
		fmt.Println(i, ":", conte)
	}
}
func Crear_Escribir() {
	file, err := os.Create("persona.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lucas := Persona{
		Name:       "lucas",
		Edad:       22,
		Completado: false,
	}
	manejo := json.NewEncoder(file)

	err = manejo.Encode(lucas)
	if err != nil {
		panic(err)
	}
}
