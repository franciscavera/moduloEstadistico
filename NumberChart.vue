<template>
    <div>
      <!-- 
      <p>ESTA PREGUNTA ES DE TIPO NUMBER</p>
      <p>Este es el arreglo de respuestas: {{ answersReports }}</p>
      <p>Esta es la pregunta: {{ questionCharacterization.subtitle }}</p>
      <p> Frecuencia {{ this.resultado }}</p>
      -->
      <apexchart type="bar" height="400" width="600" :options="chartOptions" :series="series"></apexchart>
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
        question: this.questionCharacterization,
        answers: this.answersReports,
        chartOptions: {},
        resultado: [],
        categorias: [],
        series: [],
        frecuencias: {},
      /*  chartOptions: {
          chart: {
            type: "bar",
            height: 900,
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
            categories: [],
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
        this. frecuencias = {};
        this.resultado = [];
        this.categorias = [];
        this.series = [];
        for (let i = 0; i < this.answers.length; i++) {
          const num = this.answers[i];
          if (!this.frecuencias[num]) {
            this.frecuencias[num] = 1;
          } else {
            this.frecuencias[num]++;
          }
        }
  
        for (let num in this.frecuencias) {
          this.resultado.push({ x: num, y: this.frecuencias[num] });
        }
        
      
        this.series = [{ name: 'Frecuencia', data: this.resultado }];
      },
      graficar(){
        var options = {
          chart: {
            type: "bar",
            height: 900,
         
          },
          plotOptions: {
            bar: {
              borderRadius: 4,
              horizontal: false,
              barWidth: 20,
            },
          },
          dataLabels: {
            enabled: false,
          },
          xaxis:{
            title:  { text:` ${ this.question.name}  (${ this.question.umedida})`} ,
        },
        yaxis: {
              title: {
                text: 'Frecuencias'
              }
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
        }
        this.chartOptions = options;
      },
       
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
    this.categorias = newValue
    this.graficar();
  },
},
  };
  </script>
  