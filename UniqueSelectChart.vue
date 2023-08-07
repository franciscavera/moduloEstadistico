<template>
    <div>
      <!-- 
      <p>ESTA PREGUNTA ES DE TIPO UNIQUE SELECT</p>
      <p> Frecuencia {{ frecuencias }}</p>
      <p> Respuestas {{ answers }}</p>
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
        //questionValues: this.questionCharacterization.texts,
        question: this.questionCharacterization,
        answers: this.answersReports,
        frecuencias: [],
        chartOptions: {},
        /* 
        chartOptions: {
          
          chart: {
            type: "bar",
            height: 800,
          },
          
          plotOptions: {
            bar: {
              borderRadius: 4,
              horizontal: true,
            },
          },
          dataLabels: {
            enabled: false,
          },
          xaxis: {
            categories: this.questionCharacterization.texts,
          },
          title: {
            text: this.questionCharacterization.subtitle,
            align: "center",
            style: {
              fontSize: "8px",
              fontWeight: "bold",
              color: "#444",
  
            },
          },
          
        },*/
       
      };
    },
methods: {
      calcularFrecuencias() {
        this.frecuencias = [];
        //console.log("esto tiene questionValues", this.questionValues);
        this.question.texts.forEach((numeroFijo) => {
          const frecuencia = this.answers.filter((numeroAzar) => numeroAzar === numeroFijo).length;
          this.frecuencias.push(frecuencia);
        });

      
        this.series = [
    { name: 'Frecuencia', data: this.frecuencias},
  ];
  },
      graficar(){
        var options = {
          chart: {
            type: "bar",
            height: 400,
          },
          
          plotOptions: {
            bar: {
              borderRadius: 4,
              horizontal: true,
            },
          },
          dataLabels: {
            enabled: false,
          },
          xaxis: {
            categories: this.question.texts,
            title: {
            text: "Frecuencias",
          }
          },
          yaxis: {
          title: {
            text: this.question.name,
          }},
          title: {
            text: this.question.subtitle,
            align: "center",
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
      this.calcularFrecuencias();
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
