<template>
  <div>
    <!-- 
    <p>ESTA PREGUNTA ES DE TIPO UNIQUE HORIZONTAL</p>
    <p> Este es el arreglo de respuestas</p> {{answersReports}}
    <p> este es el arreglo de las preguntas </p> {{questionCharacterization.values}}
    <p> Este es el texto de las preguntas </p> {{questionCharacterization}}
    <p> la frecuencia es</p>
    {{ frecuencias }}
    -->
    <apexchart type="bar" height="400" width="600"  :options="chartOptions" :series="series"></apexchart>
  </div>
</template>

<script>
import VueApexCharts from "vue-apexcharts";

export default {
  props: ["answersReports", "questionCharacterization"],
  components: {
    apexchart: VueApexCharts,
  },
  data() {
    return {
      questionValues: this.questionCharacterization.values,
      question : this.questionCharacterization,
      answers: this.answersReports,
      frecuencias: [],
      chartOptions: {},
      /*
      chartOptions: {
        
        chart: {
          type: "bar",
          height: 350,
        },
        
        plotOptions: {
          bar: {
            borderRadius: 4,
            horizontal: false,
          },
        },
        dataLabels: {
          enabled: false,
        },
        xaxis: {
          categories: this.questionCharacterization.texts,
        },
        title: {
          text: "¿Cómo evalúa el desempeño de la plataforma?",
          align: "center",
          style: {
            fontSize: "8px",
            fontWeight: "bold",
            color: "#444",

          },
        },
        
      }, */
    };
  },
  methods: {
    calcularFrecuencias() {
      this.frecuencias = [];
      this.question.values.forEach((numeroFijo) => {
        const frecuencia = this.answers.filter((numeroAzar) => numeroAzar === numeroFijo).length;
        this.frecuencias.push(frecuencia);
      });
      this.series = [ { name: 'Frecuencia', data: this.frecuencias }];
    },
    graficar(){
      var options = {
        chart: {
          type: "bar",
          height: 350,
        },
        
        plotOptions: {
          bar: {
            borderRadius: 4,
            horizontal: false,
          },
        },
        dataLabels: {
          enabled: false,
         
        },
        xaxis: {
          categories: this.question.texts,
          title: { text: this.question.name } ,
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
          title: {
            text: this.question.name,
          }
          
        },
        yaxis: {
              title: {
                text: 'Frecuencias'
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
    
  },
},
  created() {
    this.calcularFrecuencias();
    console.log(this.frecuencias);
    this.graficar();
  },
  watch: {
  answersReports(newValue) {
    console.log('newValue',newValue)
    this.answers = newValue
    this.calcularFrecuencias()
  },
  questionCharacterization(newValue) {
    console.log('QC NewValue',newValue)
    this.question = newValue
    this.graficar();
  },
},
};
</script>
