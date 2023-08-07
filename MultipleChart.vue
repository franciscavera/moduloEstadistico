<template>
    <div>
      <!--
      <p>Este es el texto de las preguntas: {{ question.type }}</p>
      <p>La frecuencia es:</p>
      {{ frecuenciaSi }} {{ frecuenciaNo}}
      <h3>{{ questionCharacterization.subtitle }}</h3>
       -->
      <apexchart type="bar" height="400" width="100%" :options="chartOptions" :series="series"></apexchart>
    </div>
  </template>
  
  <script>
  import VueApexCharts from "vue-apexcharts";
  export default {
    props: ["answersReports", "questionCharacterization"],
    data() {
      return {
        //questionValues: this.questionCharacterization.values,
        question: this.questionCharacterization,
        answers: this.answersReports,
        frecuenciaSi: [],
        frecuenciaNo: [],
        matriz: [],
        chartOptions: {},
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
              },
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
                categories: this.questionCharacterization.texts,
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
                  return  val 
                }
              }
            }
          },*/
          
          
        };
      },
    methods: {
  matrizRespuestas() {
    this.matriz = [];
    //dependiendo del largo de answersReports, son las filas que se crean (horizontal)
    for (let i = 0; i < this.answers.length; i++) {
      const fila = [];
      //dependiendo del largo de questionValues, son las columnas que se crean (vertical)
      for (let j = 0; j < this.question.values.length; j++) {
        const state = this.answers[i][j].state;
        const valor = state ? 1 : 0;
        fila.push(valor);
      }
      this.matriz.push(fila);
    }
    console.log("Matriz de multiple",this.matriz);
  },
  calcularFrecuencias() {
    this.frecuenciaSi = [];
    this.frecuenciaNo = [];
    //Largo de las columnas (hacia al lado ->, en horizontal)
    for (let columna = 0; columna < this.question.values.length; columna++) {
      let frecuencia1Columna = 0;
      let frecuencia0Columna = 0;
      //Largo de la filas (hacia abajo, en vertical) 
      for (let fila = 0; fila < this.matriz.length; fila++) {
        if (this.matriz[fila][columna] === 1) {
          frecuencia1Columna++;
        } else {
          frecuencia0Columna++;
        }
      }
      this.frecuenciaSi.push(frecuencia1Columna);
      this.frecuenciaNo.push(frecuencia0Columna);
      
    }
    // Crear el arreglo 'series'
  this.series = [
    { name: 'Si', data: this.frecuenciaSi },
    { name: 'No', data: this.frecuenciaNo }
  ];
  console.log("Frecuencias Multiple", this.series);
  },
  graficar(){
    var options = {
      chart: {
              type: 'bar',
              height: 350
            },
            plotOptions: {
              bar: {
                horizontal: false,
                columnWidth: '35%',
                endingShape: 'rounded', 
              },
            },
            dataLabels: {
              enabled: false,
            },
            stroke: {
              show: true,
              width: 2,
              colors: ['transparent']
            },
            xaxis: {
                categories: this.question.texts,
                labels: {
        rotate: -35, // Ajusta el ángulo de rotación para mejorar la legibilidad
        maxHeight: 80, // Define la altura máxima para cada etiqueta
        style: {
          fontSize: "10px", // Ajusta el tamaño de fuente para evitar superposiciones
        },
        // Ajustes para habilitar el ajuste de texto en varias líneas
        maxHeight: 100,
        style: {
          fontSize: '10px',
          cssClass: 'apexcharts-xaxis-label',
        },
        offsetX: 0,
        offsetY: 5,
        formatter: function (value, timestamp, opts) {
          const words = value.split(' ');
          const maxWordsPerLine = 2; // Define el número máximo de palabras por línea
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
                text: 'Frecuencia'
              }
            },
            fill: {
              opacity: 1
            },
            tooltip: {
              y: {
                formatter: function (val) {
                  return  val 
                }
              }
            },
            title: {
          align: "center",
          text: this.question.subtitle,
          style: {
            fontSize: "8px",
            fontWeight: "bold",
            color: "#444",

          },
        },
    }
    this.chartOptions = options;

  }
},

    created() {
      this.matrizRespuestas();
      this.calcularFrecuencias();
      this.graficar();
    },
    watch: {
      answersReports(newValue) {
        this.answers = newValue;
        this.matrizRespuestas();
        this.calcularFrecuencias();
      },
      questionCharacterization(newValue) {
        this.question = newValue;
        this.graficar();
      },
    },
  };
  </script>