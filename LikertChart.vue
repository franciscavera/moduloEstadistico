<template>
  <div>
    <!-- 
    <p>Estamos en el componente de Likert</p>
    <p>CATEGORIAS {{ this.questionCharacterization.texts }}</p>
    -->
    <div style="margin-bottom: 2%;">
      <label>
        <input type="radio" v-model="seleccion" value="true" @change="mostrarValoresNormalizados" /> <strong>Gráfico con valores originales </strong>
      </label>
      <label>
        <input type="radio" v-model="seleccion" value="false" @change="mostrarValoresNormalizados" /> <strong> Gráfico con valores normalizados</strong>
      </label>
    </div>
    <div> 
      <!--  <p> {{ seleccion }}</p>-->   
      <apexchart type="bar" height="400" width="1000" :options="chartOptions" :series="series"></apexchart>
    </div>
  </div>
</template>

<script>
import VueApexCharts from "vue-apexcharts";
export default {
  props: ["answersReports", "questionCharacterization"],
  components: {
    apexchart: VueApexCharts
  },
  data() {
    return {
      questionValues: this.questionCharacterization.values,
      question: this.questionCharacterization,
      answers: this.answersReports,
      type: this.questionCharacterization.type,
      title: this.questionCharacterization.subtitle,
      texts: this.questionCharacterization.texts,
      categorias: this.questionCharacterization.sentence,
      matriz: [],
      resultados: [],
      chartOptions: {},
      series: [],
      mySerie: [],
      frecNormalizadas:[],
      seriesNormalizadas:[],
      seleccion: "true",
      /*
      chartOptions: {
        chart: {
          type: 'bar',
          height: 350
        },
        plotOptions: {
          bar: {
            horizontal: false,
            columnWidth: '55%',
            endingShape: 'rounded'
          }
        },
        dataLabels: {
          enabled: false
        },
        stroke: {
          show: true,
          width: 2,
          colors: ['transparent']
        },
        xaxis: {
          categories: this.questionCharacterization.sentence.map((s) => s.text)
        },
        yaxis: {
          title: {
            text: 'Frecuencia'
          }
        },
        fill: {
          opacity: 1
        },
        tooltip: {
          y: {
            formatter: function (val) {
              return val;
            }
          }
        }, 
        legend: {
          position: 'right',
          offsetY: 40
        },
        title: {
        text: this.questionCharacterization.subtitle,
        align: "center",
        style: {
          fontSize: "8px",
          fontWeight: "bold",
          color: "#444",this.graficar();
        },
        },    
      }, */

    };
  },
  methods: {
    matrizRespuestas() {
      this.matriz = [];
      //for para crear filas (horizontal) del largo de answers reports 
      for (let i = 0; i < this.answers.length; i++) {
        const fila = [];
        //for para crear columnas del largo de sentence (enunciados de las preguntas de tipo likert)
        for (let j = 0; j < this.question.sentence.length; j++) {
          //se obtiene la respuesta que se encuentre en la posicion [i][j]
          const respuesta = this.answers[i][j];
          //se guarda el valor en la fila
          fila.push(respuesta.option);  /*la variable respuesta tiene esta estructura [ { "option": 3, "sentence": 1 } ]
          se necesita guardar el valor de "option", por lo que respuesta.option = 3 */
          
        }
        //se inserta la fila en la matriz
        this.matriz.push(fila);
      }
      console.log("Matriz Likert", this.matriz)
    },
    calcularFrecuencias() {
      this.resultados = [];
  const numFilas = this.question.values.length;
  const numColumnas = this.question.sentence.length;
  // for para recorrer las filas
  for (let i = 0; i < numFilas; i++) {
    const fila = [];
    //for para recorrer las columnas
    for (let j = 0; j < numColumnas; j++) {
      //en la variable valor se guarda el valor que tiene questionCharacterization.values[i],
      // se debe acceder a cada elemento de  questionCharacterization.values porque es un arreglo con la siguiente estructura
      // "values": [ 1, 2, 3, 4 ], estos values son los que se guardan en la BD dependiendo de la respuesta que eligio una persona
      // si eligio por ejemplo la primera respuesta, se guarda un 1 y asi con las demas
      const valor = this.question.values[i];
      // contador de la frecuencia
      let frecuencia = 0;
      //se hace otro for para acceder a las respuestas
      for (let k = 0; k < this.answers.length; k++) {
        //se accede a matriz[k][j], k porque debe pasar por todas las respuestas y j por las columnas 
        const respuesta = this.matriz[k][j];
        if (respuesta === valor) { //if para preguntar si la respuesta es === a uno de los valores [ 1, 2, 3, 4 ]
          frecuencia++; // se suma
        }
      }
      fila.push(frecuencia); //  se inserta frecuencia al arreglo fila
    
    }
    this.resultados.push(fila); //  se inserta la fila a resultados
  }

  console.log("Matriz con las frecuencias:", this.resultados);
},
/*Método para normalizar las frecuencias, al finalizar las frecuencias se econtrarán entre
[0,1]*/
normalizarFrecuencias(){
  this.normalizados=[];
  const numFilas = this.question.values.length;
  const numColumnas = this.question.sentence.length;
  let minimo = this.resultados[0][0];
  let maximo = this.resultados[0][0];
  for (let i = 0; i < numFilas; i++) {
    //for para recorrer las columnas
    for (let j = 0; j < numColumnas; j++) {
      if(this.resultados[i][j] < minimo )
      {
        minimo = this.resultados[i][j];
      }
      if(this.resultados[i][j] > maximo)
      {
        maximo = this.resultados[i][j];
      }
    }
  }
  console.log("minimo", minimo);
  console.log("maximo", maximo);
  this.frecNormalizadas=[];
  
  for (let i = 0; i < numFilas; i++) {
    //for para recorrer las columnas
    const fila = [];
    for (let j = 0; j < numColumnas; j++) {
      //formula para normalizar la frecuencia en cada posicion 
      // poscionNormalizada=((posicion[i][j]-minimo)/(maximo-minimo))
      const posNormalizada= ((this.resultados[i][j]-minimo)/(maximo-minimo)).toFixed(3);
      fila.push(posNormalizada);
    }
    this.frecNormalizadas.push(fila);
  }
  console.log("Frecuencias normalizadas", this. frecNormalizadas);
  this.seriesNormalizadas = [];
    for (let i = 0; i < this.frecNormalizadas.length; i++) {
      //se crea la serie, que correponde a una fila de resultados y el nombre
      const serie = {
          data: this.frecNormalizadas[i],
          name: this.question.texts[i],
        };
        this.seriesNormalizadas.push(serie); //  se inserta la serie en series 
        console.log("esto tiene la fila de resultados", this.seriesNormalizadas);
      }
      
      console.log("Esto tiene la serieNormalizada", this.seriesNormalizadas);
},
//se construyen las series para que el grafico las pueda graficar
prepararSeries() {
  this.series = [];
    for (let i = 0; i < this.resultados.length; i++) {
      //se crea la serie, que correponde a una fila de resultados y el nombre
      const serie = {
          data: this.resultados[i],
          name: this.question.texts[i],
        };
        this.series.push(serie); //  se inserta la serie en series 
        console.log("esto tiene la fila de resultados", this.resultados)
      }
      
      console.log("Esto tiene la serie", this.series);
    },
    mostrarValoresNormalizados() {
      if (this.seleccion === "true") {
        // Mostrar "Valores Originales" (usar la serie original)
        this.series = this.series.map((serie, i) => ({
          name: serie.name,
          data: this.resultados[i] // Usar los datos originales
        }));
      } else {
        // Mostrar "Valores Normalizados"
        this.series = this.seriesNormalizadas;
      }
      this.graficar();
    },
    graficar() {
      var seleccionRadio = this.seleccion;
      console.log("SeleccionRadio", seleccionRadio);
      var options = {
        chart: {
          type: 'bar',
          height: 350, offset: {
        x: 1000,
        y: 2000 // Ajusta este valor para cambiar el espacio en blanco entre el eje x y el gráfico
      },
        },
        plotOptions: {
          bar: {
            horizontal: false,
            columnWidth: '55%',
            endingShape: 'rounded'
          }
        },
        dataLabels: {
          enabled: false
        },
        stroke: {
          show: true,
          width: 2,
          colors: ['transparent']
        },
        xaxis: {
          /* 
          title: {
            text: 'Sentencias'
          },*/
      categories: this.question.sentence.map((s) => s.text),
      labels: {
        maxHeight: 80, // Define la altura máxima para cada etiqueta
        style: {
          fontSize: "12px", // Ajusta el tamaño de fuente para evitar superposiciones
        },
        // Ajustes para habilitar el ajuste de texto en varias líneas
        maxHeight: 100,
        style: {
          fontSize: '12px',
          cssClass: 'apexcharts-xaxis-label',
        },
        offsetX: 0,
        offsetY: 5,
        formatter: function (value, timestamp, opts) {
          const words = value.split(' ');
          const maxWordsPerLine = 3; // Define el número máximo de palabras por línea
          const lines = [];
          let line = '';

          words.forEach((word) => {
            if (line.split(' ').length < maxWordsPerLine) {
              line += line.length ? ` ${word}` : word;
            } else {
              lines.push(line);
              line = word;
            }
          });

          if (line) {
            lines.push(line);
          }

          return lines;
        },
      },
    },

        yaxis: {
          title: {
            text: 'Frecuencias'
          },
          labels: {
            formatter: function (value) {
              if (this.seleccion==="true") {
                // Si se muestran valores normalizados, se muestran sin decimales
                return value;
              } else {
                // Si se muestran valores originales, se muestran con 5 decimales
                return value.toFixed(3);
              }
            }.bind(this),

          },
        },
        fill: {
          opacity: 1
        },
        tooltip: {
          y: {
            formatter: function (val) {
              return val;
            }
          }
        }, 
        legend: {
          position: 'right',
          offsetY: 40
        },
        title: {
          text: this.question.subtitle,
          align: "center",
          style: {
            fontSize: "8px",
            fontWeight: "bold",
            color: "#444",
          },
        },    
      };
      this.chartOptions = options;
    }
  },
  created() {
    this.matrizRespuestas();
    this.calcularFrecuencias();
    this.prepararSeries();
    this.normalizarFrecuencias();
    this.graficar();
  },
  watch: {
    answersReports(newValue) {
      this.answers = newValue;
      this.matrizRespuestas();
      this.calcularFrecuencias();
      this.prepararSeries();
      
    },
    questionCharacterization(newValue) {
      this.questionValues = newValue.values;
      this.question = newValue;
      this.graficar();
      
    }
  }
};
</script>
