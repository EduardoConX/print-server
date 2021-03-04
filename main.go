package main

import (
	/* "bufio" */
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	/* "github.com/knq/escpos" */)

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
	fmt.Println(impresora)
	/* f, err := os.OpenFile(impresora, os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewReadWriter(f)
	p := escpos.New(w)
	p.Init()*/

	//Por cada operacion recibida
	for _, o := range operaciones {
		manejarOperaciones(o)
		//TODO Switch case con funciones para cada tipo de operacion permitida
	}
	/* p.End()
	   w.Flush() */
}

func manejarOperaciones(operacion Operacion) {
	switch operacion.Accion {
	case "TextoPlano":
		fmt.Println("Escribe de manera plana:", operacion.Datos)
		/* p.Write(operacion.Datos) */
	case "TextoConAcentos":
		fmt.Println("Escribe con acentos:", operacion.Datos)
		/* p.Write(operacion.Datos) */
	case "Feed":
		fmt.Println("Feed:", operacion.Datos)
		/* p.FormfeedN(5) */
	case "TamanioFuente":
		valores := strings.Split(operacion.Datos, ",")
		fmt.Println("Ancho:", valores[0])
		fmt.Println("Alto:", valores[1])
		/* p.SetFontSize(valores[0], valores[1]) */
	case "Fuente":
		fmt.Println("Fuente:", operacion.Datos)
		/* p.SetFont("C") */
	case "Enfasis":
		fmt.Println("Enfasis:", operacion.Datos)
		/* p.SetEmphasize(1) */
	case "Alineacion":
		fmt.Println("Alineacion:", operacion.Datos)
	case "Cortar":
		fmt.Println("Cortar")
		/* p.Cut() */
	case "CortarParcialmente":
		fmt.Println("Cortar parcialmente")
	}
}
