package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
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
	/* Copy("impresion", impresora) */
	copy("impresion", "\\\\EduardX-PC\\\\POS-58-Series")
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
	//Convierte el texto a byte
	b := []byte(texto)

	//Decodifica los acentos y las ñ
	c, e := charmap.CodePage850.NewEncoder().Bytes(b)
	if e != nil {
		log.Fatal(e)
	}

	//Imprime el texto codificado
	return []byte(c)
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

func copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}
