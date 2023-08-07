<template>

  <div>
    <!-- 

    <p> Este es el arreglo de respuestas</p> {{answersReports}}
    <p> este es el arreglo de las preguntas </p> {{questionCharacterization.values}}
    <p> Este es el texto de las preguntas </p> {{questionCharacterization.texts}}
    <p> la frecuencia es</p>
    {{ resultados }}
    <h3> {{questionCharacterization.subtitle}} {{ type }}</h3>
     <p> la frecuencia es</p>
    {{ frecuencias }}
    -->
   
    <apexchart type="donut" height="500" width="600"  :options="chartOptions" :series="series"></apexchart>
  </div>
</template>
<script>
import VueApexCharts from "vue-apexcharts";
export default {
props:["answersReports", "questionCharacterization"],
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
              type: 'donut',
            },
            plotOptions: {
              pie: {
                startAngle: -90,
                endAngle: 90,
                offsetY: 10
              }
            },
            grid: {
              padding: {
                bottom: -80
              }
            },
            labels: this.question.texts,
            responsive: [{
              breakpoint: 480,
              options: {
                chart: {
                  width: 200
                },
                legend: {
                  position: 'bottom'
                }
              }
            }],
            title: {
          text: this.question.subtitle,
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
    this.question.values.forEach((numeroFijo) => {
      const frecuencia = this.answers.filter((numeroAzar) => numeroAzar === numeroFijo).length;
      this.frecuencias.push(frecuencia);
    });
    this.series = this.frecuencias;
    
    },
    graficar() {
      var options = {
        
      chart: {
              type: 'donut',
            },
            plotOptions: {
              pie: {
                startAngle: -90,
                endAngle: 90,
                offsetY: 10
              }
            },
            grid: {
              padding: {
                bottom: -80
              }
            },
            labels: this.question.texts,
            responsive: [{
              breakpoint: 480,
              options: {
                chart: {
                  width: 600,
                },
                legend: {
                  position: 'bottom'
                }
              }
            }],
            title: {
          text: this.question.subtitle,
          offsetX: -50,
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
  console.log(this.frecuencias);
  this.graficar();
},
watch: {
  answersReports(newValue) {
    console.log('newValue',newValue)
    console.log('Este es el valor anterior',newValue)
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
