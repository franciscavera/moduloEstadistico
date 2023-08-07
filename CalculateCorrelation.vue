<template> 
  <div class="container">
  <!-- 
  <p> pregunta del index:{{ uno}} {{ dos}} </p>
 
   <p> pregunta del index: {{ questions[this.uno].subtitle }}</p>
  <p> UNO: answers {{answersReports[this.uno].respuestas }}</p>
  <p> DOS: answers {{answersReports[this.dos].respuestas }}</p>
  -->
  <!-- Se valida que el largo del arreglo de las respuestas sea mayor que dos, si es menor no 
  es posible calcular el coeficiente de correlación-->
  <!--<p> UNO: answers {{answersReports[this.uno].respuestas }}</p>
  <p> DOS: answers {{answersReports[this.dos].respuestas }}</p>
   -->
  <div v-if="answersReports[this.uno].respuestas.length >=2 "> 
    <h2 style="text-align: center;"><u> Análisis Coeficiente de Correlación de Pearson</u> </h2>
   
  <div> 
    <apexchart type="scatter" height="300%" width="80%" :options="chartOptions" :series="series"></apexchart>
  </div>

  
  <div class="text-with-button">
  <h2> <u> Interpretación de los resultados</u></h2>
  <button class="btn btn-state" @click="openModal">Ayuda</button>
  <ModalInformation ref="modalInformation"> </ModalInformation>
  </div>
  <h3> Coeficiente de Correlacion de Pearson (r) = {{ correlacion }}</h3>
  <div v-if="correlacion == -1"> 
    <p class="p-text"> El valor es igual a [-1]</p>
    <p class="p-text"> Como r = {{correlacion}}. Se puede observar que existe una relación lineal perfecta
        negativa entre las dos variables, lo que significa que en la medida que una variable 
        aumenta,la otra siempre disminuye de manera proporcional. Esta relación se puede 
        observar claramente en el diagrama de dispersión, los puntos de datos estarán perfectamente
        alineados en una línea recta descendente.</p>
  </div>
  <div v-else-if="correlacion > -1 && correlacion  < -0.5"> 
    <p class="p-text"> El valor se encuentra entre [-1, -0.5]</p>
    <p class="p-text"> Al observar que r = {{ correlacion }}, podemos interpretar que hay una asociación moderada
       entre las variables. Este Coeficiente de Correlación negativo sugiere que cuando una variable
        aumenta, la otra tiende a disminuir, y viceversa. En el diagrama de disperción 
        los puntos de datos tienden a mostrar una tendencia hacia abajo, lo que significa que los 
        puntos están más o menos distribuidos en una dirección descendente a medida que se mueven 
        desde el extremo superior izquierdo al extremo inferior derecho plano.</p>
  </div>
  <div v-else-if="correlacion >= -0.5 && correlacion < 0">
    <p class="p-text"> El valor se encuentra entre [-0.5, 0]</p>
    <p class="p-text"> Al observar que r = {{ correlacion }}. Se puede concluir que existe una relación lineal
       negativa débil entre las variables analizadas. Cuando una variable aumenta, la otra tiende a disminuir, y viceversa.
       Sin embargo, la magnitud del coeficiente es cercana a cero, lo que indica que la relación 
       es débil y no muy pronunciada. En diagrama de dispersión, los puntos de datos se dispersarán
       alrededor del plano cartesiano, mostrando una tendencia leve hacia abajo en la distribución de los puntos, 
       sin seguir una direccion específica </p>
  </div>
  <div v-else-if="correlacion == 0">
    <p class="p-text"> El valor es igual a [0]</p>
    <p class="p-text">  Al observar que r = {{ correlacion }}. Se puede concluir que NO existe relación lineal
    entre las variables.  Esto significa que los cambios en una variable  no están 
    relacionados con los cambios en la otra variable en una forma lineal. En el diagrama 
    de dispersión  los puntos de datos se dispersarán aleatoriamente alrededor de toda
    el área del plano, sin mostrar una agrupación en ninguna dirección específica.</p>
  </div>
  <div v-else-if="correlacion > 0 && correlacion <= 0.5">
    <p class="p-text"> El valor se encuentra entre a [0, 0.5]</p>
    <p class="p-text"> Al observar que r = {{ correlacion }}. Se puede concluir que existe una relación lineal
       débil entre las variables analizadas. Es importante destacar que un Coeficiente de 
       Correlación lineal cercano a 0 no implica que no haya relación entre las variables, 
       puede existir otro tipo de relación entre ellas que no es capturada por el coeficiente
      de correlación lineal. En el diagrama de dispersión se puede visualizar que los puntos se
      encuentran dispersos aleatoriamente en el plano cartesiano sin seguir una dirección clara.</p>
  </div>
  <div v-else-if="correlacion > 0.5 && correlacion < 1">
    <p class="p-text"> El valor se encuentra entre [0.5, 1]</p>
    <p class="p-text"> Al observar que r = {{ correlacion }}, podemos interpretar que hay una asociación moderada
       entre las variables. Este Coeficiente de Correlación positivo sugiere que, en general, 
       a medida que una variable aumenta, la otra tiende a aumentar también, aunque no de manera 
       perfecta. En el diagrama de dispersión se puede visualizar que los puntos se ajustan de forma
       creciente en lugar de formar una línea recta exacta,los puntos se dispersan alrededor de una
        línea de tendencia ascendente.</p>
    <p class="p-text"> Este tipo de relación entre variables es típico en muchos conjuntos de datos de muestras reales, 
      ya que no es frecuente que  encontremos relaciones lineales perfectas entre variables.</p>
  
  </div>
  <div v-else-if="correlacion == 1">
    <p class="p-text"> El valor se encuentra entre [1]</p>
    <p class="p-text"> Como r = {{correlacion}}. Se puede observar que existe una relación lineal perfecta
        positiva entre las dos variables, lo que significa que en la medida que una variable 
        aumenta, la otra también lo hace en una relación proporcional. Esta relación se puede 
        observar claramente en el diagrama de dispersión, ya que todos los puntos se ajustan a
         una línea recta creciente. </p>
  </div>
  <div v-else>
    <p class="p-text"> El valor se encuentra fuera de rango</p>
    <p class="p-text">Cuando el resultado del Coeficiente de Correlación es invalido o NaN:</p>
    <p class="p-text"> 1. Falta de variabilidad en los datos, (no hay variabilidad en los datos o una variable
       tiene valores constantes) </p>
    <p class="p-text"> 2. Existen valores extremos o atípicos, cuando hay este tipo de valores el Coeficiente de 
      Correlación se puede ver alterado.
    </p>
  </div>

  </div>
  
  <div v-else> 

  </div>
  
     
   </div>
   
    
  
      
</template>
<script>
  import { formatter } from "@/components/mixin/formatter.js";
  import { sampleCorrelation } from 'simple-statistics';
  import ModalInformation from "@/components/ModalInformation.vue";
  import VueApexCharts from "vue-apexcharts";
  
  export default {
    props: ["idQuestionSelected1", "idQuestionSelected2", "questionsForm", "answersReports", "idFormSelected"],
    mixins:[formatter],
    components:{
      ModalInformation,
    },
    data() {
     
       return {
        matriz: [], 
        questions: this.questionsForm,
        answers: this.answersReports,
        uno: this.idQuestionSelected1,
        dos: this.idQuestionSelected2,
        correlacion: "",
        arregloParesOrdenados:[],
        chartOptions: {},
        showModal:"",
        }; 
    },
methods: {
  calcularCorrelacion(){
    this.correlacion="";
    //let arreglo1=this.answersReports[this.uno];
    //let arreglo2=this.answersReports[this.dos];
    if(this.answers[this.uno].respuestas.length>=2){
      //this.alertError("Para calcular el coeficiente de correlación se necesitan al menos dos pares de datos");
    //}  
    //else{
    //let arreglo1=[22, 28, 35, 42, 48, 55, 62, 69, 75, 82, 89, 95];
    //let arreglo2=[60, 68, 75, 82, 88, 95, 102, 109, 116, 122, 129, 136];
    this.correlacion = "";
    console.log("Arreglo1", this.answers[this.uno].respuestas);
    console.log("Arreglo2", this.answers[this.dos].respuestas);
    const correlacion = sampleCorrelation(this.answers[this.uno].respuestas, this.answers[this.dos].respuestas).toFixed(2);
    //const correlacion = sampleCorrelation(arreglo1, arreglo2);
    this.correlacion = correlacion;
    console.log("esto tiene correlacion", this.correlacion);          
    }  
    else{
      this.alertError("Para calcular el coeficiente de correlación se necesitan al menos dos pares de datos");
    }  
  },
  paresOrdenados(){
    this.arregloParesOrdenados = [];
    this.series = [];
    //let arregloX = [22, 28, 35, 42, 48, 55, 62, 69, 75, 82, 89, 95];
    //let arregloY = [60, 68, 75, 82, 88, 95, 102, 109, 116, 122, 129, 136];
    if(this.answers[this.uno].respuestas.length>=2){
    for(let i=0;i<this.answers[this.uno].respuestas.length;i++)
    //for(let i=0;i<arregloX.length;i++)
    {
      const x = this.answers[this.uno].respuestas[i];
      const y = this.answers[this.dos].respuestas[i];
      //const x = arregloX[i];
      //const y = arregloY[i];
      const parXY= [x,y];
      this.arregloParesOrdenados.push(parXY);
    }
    console.log("contenido de arregloXY", this.arregloParesOrdenados);
    this.series = [{data: this.arregloParesOrdenados }];
    /*this.series = [
      { name: this.questions[this.idQuestionSelected1].subtitle, data: this.arregloParesOrdenados }
    ];*/
    }
    
    },
  graficar() {
    var variableX = this;
    var variableY = this;
    var options = {
    chart: {
  
      type: "scatter",

    },
    xaxis: {
      min: 0, //(valor minimo del eje x)
       tickAmount: 10, // divisiones del eje x
      // type: "numeric",
      title: { text:` ${this.questions[this.idQuestionSelected1].name}  (${this.questions[this.idQuestionSelected1].umedida})`} ,
      labels: {
        formatter: function (val) {
          return parseFloat(val).toFixed(1);
        },
      },
    },
    yaxis: {
      tickAmount: 7,
      title: { text:` ${ this.questions[this.idQuestionSelected2].name}  (${ this.questions[this.idQuestionSelected2].umedida})`} ,

    },
    title: {
      text: "Diagrama dispersión",
      align: "center",
      style: {
        fontSize: "8px",
        fontWeight: "bold",
        color: "#444",
      },
    },
    tooltip: {
      y: {
        formatter: function (val) {
          return val;
        },
      },
      //Para que la etiqueta que se ve cuando uno se posiciona sobre el punto quede del estilo (X,Y)
      custom: function ({ series, seriesIndex, dataPointIndex, w }) {
        const xValue = w.globals.seriesX[seriesIndex][dataPointIndex]; 
        const yValue = series[seriesIndex][dataPointIndex];
        return `<div>(${variableX.questions[variableX.idQuestionSelected1].name}: ${xValue} ${variableX.questions[variableX.idQuestionSelected1].umedida}), (${variableY.questions[variableY.idQuestionSelected2].name}: ${yValue} ${variableY.questions[variableY.idQuestionSelected2].umedida})</div>`;
        //return `<div> {{this.questions[this.idQuestionSelected1].name:}} (${xValue}, Y:${yValue})</div>`;
      },
    },
  };
  this.chartOptions = options;
},
openModal() {
  this.$refs.modalInformation.show();
},
closeModal() {
  this.showModal = false;
},
},    
created() {
  this.calcularCorrelacion();
  this.paresOrdenados();
  this.graficar(); 
},
watch: {
  answersReports(newValue) {
    this.answers = newValue;
    this.calcularCorrelacion();
    this.paresOrdenados();
    this.graficar();

  },       
  idQuestionSelected1(newValue) {
    this.uno = newValue;    
    this.graficar(); 
    this.calcularCorrelacion();
    this.paresOrdenados();
  },
  idQuestionSelected2(newValue){
    this.dos = newValue;
    this.graficar();
    this.calcularCorrelacion();
    this.paresOrdenados();
  },
  questionsForm(newValue){
  this.questions = newValue;
  this.graficar();
  }
}
}
</script>    
<style>
.text-with-button {
  display: flex;
  align-items: center;
}
.text-with-button button {
  margin-left: 10px; /* Espacio entre el texto y el botón (ajusta el valor según tu preferencia) */
}
.p-text{
  text-align: justify;
  font-size: 22px;
}
.container {
  display: flex;
  flex-wrap: wrap; /* Permite que los elementos se envuelvan en filas o columnas según el espacio disponible */
  justify-content: center; /* Centra los elementos horizontalmente */
  align-items: center; /* Centra los elementos verticalmente */
  max-width: 2000px; /* Establece un ancho máximo para el contenedor */
  margin: 0 auto; /* Centra el contenedor en la pantalla */
  padding: 20px; /* Agrega un espacio interno alrededor del contenedor (opcional) */
}

</style>  