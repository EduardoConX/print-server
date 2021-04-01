package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		imprimir(c.Operaciones, c.Impresora)

	}
}

func configurarCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func imprimir(operaciones []Operacion, impresora string) {
	//Inicializa la impresora
	f, err := os.Create("impresion")

	//Checa si hay errores, en caso de haber se cierra el programa
	if err != nil {
		panic(err)
	}

	//Es buena practica cerrar un archivo inmediatamente despues de crearlo
	defer f.Close()

	w := bufio.NewWriter(f)
	w.Write(inicia())

	//Por cada operacion recibida
	for _, operacion := range operaciones {
		w.Write(manejarOperaciones(operacion))
	}

	w.Flush()
}

func manejarOperaciones(operacion Operacion) []byte {
	switch operacion.Accion {
	case "TextoPlano":
		return texto(operacion.Datos)
	case "Feed":
		return feed(operacion.Datos)
	case "TamanioFuente":
		return fontSize(operacion.Datos)
	case "Alinear":
		return align(operacion.Datos)
	case "Enter":
		return enter()
	}

	return []byte("")
}

func inicia() []byte {
	return []byte("\x1B@")
}

func fontSize(datos string) []byte {
	valores := strings.Split(datos, ",")
	ancho, err := strconv.Atoi(valores[0])
	if err != nil {
		panic(err)
	}
	alto, err := strconv.Atoi(valores[1])
	if err != nil {
		panic(err)
	}
	return []byte(fmt.Sprintf("\x1D!%c", ((ancho-1)<<4)|(alto-1)))
}

func align(align string) []byte {
	switch align {
	case "I":
		return []byte(fmt.Sprintf("\x1Ba%c", 0))
	case "C":
		return []byte(fmt.Sprintf("\x1Ba%c", 1))
	case "D":
		return []byte(fmt.Sprintf("\x1Ba%c", 2))
	}
	return []byte(fmt.Sprintf("\x1Ba%c", 0))
}

func texto(texto string) []byte {
	return []byte(texto)
}

func feed(nLineas string) []byte {
	n, err := strconv.Atoi(nLineas)
	if err != nil {
		panic(err)
	}
	return []byte(fmt.Sprintf("\x1Bd%c", n))
}

func enter() []byte {
	return []byte("\n")
}
