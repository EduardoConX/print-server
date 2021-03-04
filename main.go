package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Operacion struct {
	Accion string
	Datos  string
}

type Cuerpo struct {
	Operaciones []Operacion
	Impresora   string
}

func main() {
	http.HandleFunc("/", manejador) //(ruta, funcion)

	fmt.Printf("Iniciando servidor...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func manejador(w http.ResponseWriter, r *http.Request) {
	configurarCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var c Cuerpo
		err := decoder.Decode(&c)
		if err != nil {
			panic(err)
		}

		//Por cada operacion recibida
		for _, o := range c.Operaciones {
			fmt.Println(o)
			//TODO Switch case con funciones para cada tipo de operacion permitida
		}
	}
}

func configurarCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
