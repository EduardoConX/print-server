class Conector {
  constructor(impresora = "POS-58-SERIES", ruta = "http://localhost:8080/") {
    this.impresora = impresora;
    this.ruta = ruta;
    this.operaciones = []; //Arreglo donde se iran acumulando las operaciones a realizar
  }

  textoPlano(texto) {
    this.operaciones.push({
      accion: "TextoPlano",
      datos: texto,
    });
  }

  textoConAcentos(texto) {
    this.operaciones.push({
      accion: "TextoConAcentos",
      datos: texto,
    });
  }

  enter() {
    this.operaciones.push({
      accion: "Enter",
      datos: "",
    });
  }

  feed(nLineas) {
    if (Number.isInteger(nLineas) && nLineas > 0) {
      this.operaciones.push({
        accion: "Feed",
        datos: nLineas.toString(),
      });
    }
  }

  tamanioFuente(ancho, alto) {
    if (Number.isInteger(ancho) && Number.isInteger(alto)) {
      if (ancho >= 1 && ancho <= 8) {
        if (alto >= 1 && alto <= 8) {
          //Solo se ejecutara si ancho y alto son numeros enteros entre 1 y 8
          this.operaciones.push({
            accion: "TamanioFuente",
            datos: `${ancho},${alto}`,
          });
        }
      }
    }
  }

  fuente(fuente) {
    if (typeof fuente === "string") {
      fuente = fuente.toUpperCase();
      if (fuente === "A" || fuente === "B" || fuente === "C") {
        //Solo se ejecutara si se selecciona una fuente valida
        this.operaciones.push({
          accion: "Fuente",
          datos: fuente,
        });
      }
    }
  }

  enfasis(valor) {
    if (valor === true || valor === false) {
      //Si los datos son numeros los intentara convertir a valores true y false
      if (valor === true) {
        valor = 1;
      } else if (valor === false) {
        valor = 0;
      }
    }

    if (valor === 1 || valor === 0) {
      //Solo se ejecutara si se selecciona un valor valido
      this.operaciones.push({
        accion: "Enfasis",
        datos: valor.toString(),
      });
    }
  }

  alineacion(valor) {
    if (typeof valor === "string") {
      valor = valor.toUpperCase();
      if (valor === "I" || valor === "C" || valor === "D") {
        //Solo se ejecutara si se selecciona un valor valido
        this.operaciones.push({
          accion: "Alinear",
          datos: valor,
        });
      }
    }
  }

  cortar() {
    this.operaciones.push({
      accion: "Cortar",
      datos: "",
    });
  }

  cortarParcialmente() {
    this.operaciones.push({
      accion: "CortarParcialmente",
      datos: "",
    });
  }

  //TODO imagenes, codigos de barra y qr

  imprimir() {
    const cuerpo = {
      operaciones: this.operaciones,
      impresora: this.impresora,
    };

    fetch(this.ruta, {
      method: "POST",
      body: JSON.stringify(cuerpo),
    });
    console.log(JSON.stringify(cuerpo));
  }
}
