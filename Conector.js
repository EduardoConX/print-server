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

  imprimir() {
    const cuerpo = {
      operaciones: this.operaciones,
      impresora: this.impresora,
    };

    fetch(this.ruta, {
      method: "POST",
      body: JSON.stringify(cuerpo),
    });
  }
}
