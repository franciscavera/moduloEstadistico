<template> 
    <div >
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
     <!-- 
     <p> pregunta del index1 name: {{ questions[this.uno].name }}</p>
     <p> pregunta del index1 name: {{ questions[this.dos].name }}</p>
     <p> pregunta del index1: {{ questions[this.uno].values }}</p>
     <p> pregunta del index2: {{ questions[this.dos].values }}</p>
     <p> pregunta del index: {{ questions[this.uno] }}</p> -->
     
     <h2 style="text-align: center;"><u> Tabla de Contingencia</u> </h2>
    <div v-if="matrizFrecuencias!=[]">
      <table class="styled-table">
      <tr>
        <td rowspan="2" style="text-align: center;">{{ questions[this.uno].name }}</td>
        <td :colspan="questions[this.dos].texts.length" style="text-align: center;">{{ questions[this.dos].name }}</td>
        <td rowspan="2"> Total</td>

      </tr>
      
      <tr>
        
        <th  v-for="(nombreColumna, columnIndex) in questions[this.dos].texts" :key="columnIndex">
          {{ nombreColumna }}
        </th>
       
      </tr>
      
      <tbody>
        
       
        <tr v-for="(nombreFila, rowIndex) in questions[this.uno].texts" :key="rowIndex">
          <td><strong>{{ nombreFila }}</strong></td> <!-- Mostrar el nombre de la fila -->
          <td v-for="(frecuencia, columnIndex) in matrizFrecuencias[rowIndex]" :key="columnIndex">
            <p style="text-align: center;">{{ frecuencia }}</p>
           
          </td>
          <td > 
        <p> {{ rowTotals[rowIndex] }}</p>
      </td>
          
        </tr>
        <tr>
          <td> Total</td>
          <td v-for="(total, columnIndex) in columnTotals" :key="columnIndex"> {{ total }}</td>
          <td> {{ totalFrecuencias }}</td>
        </tr>
     
        
        
      </tbody>
     
    </table>

     </div>
    <div v-if="todosMayoresIgualCinco == false" > 
      <h2 style="color: red;"> En el calculo de las frecuencias esperadas, hay valores menores a 5 (ver tabla que se muestra abajo)
         Por lo tanto, no hay suficientes datos para realizar la Prueba Chi-Cuadrado y que sea estadísticamente válida.
        Se recomienda aumentar la cantidad de formularios respondidos para tener un tamaño de muestra más grande
          con el fin de obtener frecuencias esperadas más confiables.
     </h2>
     <h2 style="text-align: center;"><u> Tabla con Frecuencias Esperadas </u> </h2>
     <table class="styled-table">
      <tr>
        <td rowspan="2" style="text-align: center;">{{ questions[this.uno].name }}</td>
        <td :colspan="questions[this.dos].texts.length" style="text-align: center;">{{ questions[this.dos].name }}</td>
      </tr>
      <tr>
        <th  v-for="(nombreColumna, columnIndex) in questions[this.dos].texts" :key="columnIndex">
          {{ nombreColumna }}
        </th>
      </tr>
      <tbody>
        <tr v-for="(nombreFila, rowIndex) in questions[this.uno].texts" :key="rowIndex">
          <td><strong>{{ nombreFila }}</strong></td> <!-- Mostrar el nombre de la fila -->
          <td v-for="(frecuencia, columnIndex) in frecuenciasEsperadas[rowIndex]" :key="columnIndex">
            <p style="text-align: center;">{{ frecuencia.toFixed(3) }}</p>
           
          </td>
        </tr>
      </tbody>
    </table>

    </div>
    <div v-else>
    
      <h2 style="text-align: left;"><u> Prueba Chi-Cuadrado de Independencia </u> </h2>
     

      <p> Para ver si las variables seleccionadas estan relacionadas se utilizará la "Prueba chi-cuadrado de independencia X²"  </p>
      <p> Las Hipótesis a contrastar son: </p>
      <p> H0: La variable <strong> {{ questions[this.uno].name }}</strong> y <strong> {{ questions[this.dos].name }}</strong> son independientes. </p>
      <p> HA: La variable <strong> {{ questions[this.uno].name }}</strong>  y <strong> {{ questions[this.dos].name }}</strong> son dependientes.</p>
     
      
    <div class="text-with-button">
    <h3> <u> Interpretación de los resultados</u></h3>
    <button class="btn btn-state" @click="openModal">Ayuda</button>
    <ModalInformationChi ref="modalInformationChi"> </ModalInformationChi>
    </div>
    <h4> <u> Resultado Prueba Chi-Cuadrado</u></h4>
    <h3> Valor α = 0.05</h3>
    <h3> Grados de Libertad = {{ gradosDeLibertad }}</h3>
    <h3> Valor P obtenido = {{ valorP }}</h3>
    
    <div v-if="valorP>0.05"> 
      <p class="p-text"> Al observar que Valor P = {{ valorP}} > α (0.05), no tenemos suficiente evidencia para rechazar
        la hipótesis nula (H0) y no podemos concluir que hay una diferencia significativa entre variables que estamos comparando 
        ({{ questions[this.uno].name }} y {{ questions[this.dos].name }}). Por lo tanto {{ questions[this.uno].name }} y {{ questions[this.dos].name }} 
        no estan relacionadas (son variables independientes).
      </p>
    
    </div>
    <div v-if="valorP<=0.05"> 
      
      <p class="p-text"> Al observar que Valor P = {{ valorP}} < α (0.05), se considera que el resultado es estadísticamente significativo entre
        las variables que estamos comparando echazamos ({{ questions[this.uno].name }} y {{ questions[this.dos].name }}). Por lo tanto rechazamos 
        la hipótesis nula y aceptamos la hipótesis alternativa por lo que podemos decir que {{ questions[this.uno].name }} y {{ questions[this.dos].name }}
        estan relacionadas (son variables dependientes).
      </p>
    
    </div>
       
     
     <h4> <u> Coeficiente V de Cramer </u></h4>
     <h3> Coeficiente de Cramer (V) = {{cramer}}</h3>
     <div v-if="cramer==0"> 
      <p class="p-text"> Como V = {{cramer}}, es igual a 0, por lo que 
      no existe una relación entre las variables categóricas analizadas.
    </p>
     </div>
     <div v-else-if="cramer>0 && cramer <=0.1"> 
    
    <p class="p-text"> Como V = {{cramer}}. Esto indica una asociación muy débil o casi nula
       entre las variables categóricas que se están evaluando. En este caso, la relación entre las
        variables es prácticamente inexistente o es extremadamente débil.
    </p>
  </div>
     <div v-else-if="cramer> 0.1 && cramer < 0.2 "> 
    
    <p class="p-text"> Como V = {{cramer}}. Indica una asociación débil entre las variables
      categóricas que se están evaluando. En este caso, la relación entre las variables es prácticamente
      inexistente o es extremadamente débil.
    </p>
  </div>
  <div v-else-if="cramer>=0.2 && cramer <0.4 "> 
    <p class="p-text"> Como V = {{cramer}}. Indica una asociación moderada entre las variables categóricas 
      que se están evaluando. Es decir, existe una relación significativa entre las variables, pero no es extremadamente
      fuerte.</p>
  </div>
  <div v-else-if="cramer>=0.4 && cramer<0.6  "> 
    <p class="p-text"> Como V = {{cramer}}. Indica una asociación moderadamente fuerte entre las variables categóricas
      que se están evaluando. Esta asociación sugiere que las variables están relacionadas de manera más marcada que
      una asociación débil o moderada, pero aún no alcanza el nivel de una asociación fuerte.</p>
  </div>
  <div v-else-if="cramer>=0.6 && crame<0.8  "> 
    <p class="p-text"> Como V = {{cramer}}. Indica una asociación fuerte entre las variables categóricas que se
      están evaluando. Una asociación fuerte sugiere que las variables están relacionadas de manera considerable
       y que hay un patrón o tendencia muy clara entre las categorías de las variables</p>
  </div>
  <div v-else-if="cramer>=0.8 && crame<1 "> 
    <p class="p-text"> Como V = {{cramer}}. Indica una asociación muy fuerte entre las variables categóricas que se están 
      evaluando. Una asociación muy fuerte sugiere que las variables están altamente relacionadas. La fuerza de esta asociación
      es muy alta, lo que indica que las variables están fuertemente influenciándose mutuamente.

</p>
  </div>
  <div v-else-if="cramer==1 "> 
    <p class="p-text"> Como V = {{cramer}}.Existe una asociación perfecta entre las variables. Es decir, las variables están
      completamente relacionadas</p>
  </div>
    
    </div>
    

  </div>
    
        
  </template>
  <script>
    import { formatter } from "@/components/mixin/formatter.js";
    import jStat from 'jstat';
    import ModalInformationChi from "@/components/ModalInformationChi.vue";
    import VueApexCharts from "vue-apexcharts";
    
    export default {
      props: ["idQuestionSelected1", "idQuestionSelected2", "questionsForm", "answersReports", "idFormSelected"],
      mixins:[formatter],
      components:{
        ModalInformationChi,
      },
      data() {
       
         return {
          matriz: [], 
          questions: this.questionsForm,
          answers: this.answersReports,
          uno: this.idQuestionSelected1,
          dos: this.idQuestionSelected2,
          arregloParesOrdenados:[],
          chartOptions: {},
          showModal:"",
          matriz: [],
          matrizFrecuencias:[],
          frecuenciasEsperadas:[],
          sumatoriaCuadrados:"",
          valorP:"",
          bandera:"",
          totalFrecuencias:"",
          cramer:"",
          todosMayoresIgualCinco: "",
          rowTotals: "",
          columnTotals:"",
          }; 
      },
  methods: {
   matrizFrecuenciasRespuestas(){
    this.matrizFrecuencias=[];
    //dependiendo del largo de answersReports, son las filas que se crean (horizontal)
    for (let i = 0; i < this.questions[this.uno].values.length; i++) {
      const fila = [];
      //dependiendo del largo de questionValues, son las columnas que se crean (vertical)
      for (let j = 0; j < this.questions[this.dos].values.length; j++) {
        fila.push(0);
      }
      this.matrizFrecuencias.push(fila);
    }
    let valuesQuestion1=this.questions[this.uno].values;
    let valuesQuestion2=this.questions[this.dos].values;
    // Llenar matriz con las frecuencias
    for (let i = 0; i <this.answers[this.uno].respuestas.length; i++) {
    /*se recorre cada uno de los arreglos de las preguntas seleccionadas, para ver 
    cuantas veces aparece la combinacion [valueQuestion1,valueQuestion2], debido a que 
    valuesQuestion1 y valuesQuestion2 son los values que tienen lar aletarnativas de las preguntas,
    en la BD se guarda el value asociado a la alternativa que escogio la persona*/
    const unoIndex = valuesQuestion1.indexOf(this.answers[this.uno].respuestas[i]);
    const dosIndex = valuesQuestion2.indexOf(this.answers[this.dos].respuestas[i]);
  
    this.matrizFrecuencias[unoIndex][dosIndex]++;
    // Calcular los totales por fila
    this.rowTotals=0;
    this.columnTotals=0;
    this.rowTotals = this.matrizFrecuencias.map(row => row.reduce((acc, val) => acc + val, 0));
    console.log("total por fila", this.rowTotals);
      // Calcular los totales por columna
      const numRows = this.questions[this.uno].values.length;
      const numCols = this.questions[this.dos].values.length;
      this.columnTotals = Array.from({ length: numCols }, (_, colIndex) =>
        this.matrizFrecuencias.reduce((acc, row) => acc + row[colIndex], 0)
      );
      console.log("total por columnas", this.columnTotals);
}

    console.log("Matriz Frecuencias",this.matrizFrecuencias);
    

   },
   calcularFrecuenciasEsperadas(){
    const numFilas = this.matrizFrecuencias.length;
    const numColumnas = this.matrizFrecuencias[0].length;
    const totalCount = this.matrizFrecuencias.flat().reduce((acc, val) => acc + val, 0);
    
    this.frecuenciasEsperadas = [];

    for (let i = 0; i < numFilas; i++) {
      const rowTotal = this.matrizFrecuencias[i].reduce((acc, val) => acc + val, 0);
      const rowExpected = [];
      
      for (let j = 0; j < numColumnas; j++) {
        const colTotal = this.matrizFrecuencias.reduce((acc, row) => acc + row[j], 0);
        const expected = (rowTotal * colTotal) / totalCount;
        rowExpected.push(expected);
      }
      
      this.frecuenciasEsperadas.push(rowExpected);
    }
    
   console.log("frecuencias esperadas",  this.frecuenciasEsperadas);
   this.todosMayoresIgualCinco = true;
   for (let i = 0; i < numFilas; i++){
    for (let j = 0; j < numColumnas; j++){
      if( this.frecuenciasEsperadas[i][j]<5){
           this.todosMayoresIgualCinco = false;
            break;
      }
    }
     // Si encontramos un valor menor que 5, no es necesario seguir recorriendo el resto de filas
    if (!this.todosMayoresIgualCinco) {
    break;
    }
   }
  
  
   },
   estadisticoPruebaChiCuadrado() {
   
if(this.todosMayoresIgualCinco == true)
{
   // Paso 2: Calcular el estimador puntual
const estimadorPuntual = this.matrizFrecuencias.map((fila, i) =>
  fila.map((valor, j) => valor -  this.frecuenciasEsperadas[i][j])
);

// Paso 4: Calcular la diferencia entre el estimador puntual y el valor nulo
const diferencia = this.matrizFrecuencias.map((fila, i) =>
  fila.map((valor, j) => valor -  this.frecuenciasEsperadas[i][j])
);
console.log("diferencia", diferencia);
// Paso 5: Calcular la raíz cuadrada del estimador puntual
const raizCuadradaEstimadorPuntual =  this.frecuenciasEsperadas.map((fila) =>
  fila.map((valor) => Math.sqrt(valor))
);

// Paso 6: Calcular el valor Z
const valorZ = diferencia.map((fila, i) =>
  fila.map((valor, j) => valor / raizCuadradaEstimadorPuntual[i][j])
);

console.log("Estimador Puntual:", estimadorPuntual);
console.log("Valor Z:", valorZ);
//Paso 7: Sumatoria de los valores Z al cuadrado
const cuadradoValorZ = valorZ.map((fila, i) =>
  fila.map((valor, j) => valor * valor)
);

console.log("valorZcuadrado", cuadradoValorZ);
// Sumar todos los elementos de la matriz cuadradoValorZ
this.sumatoriaCuadrados = 0;
for (let i = 0; i < cuadradoValorZ.length; i++) {
  for (let j = 0; j < cuadradoValorZ[i].length; j++) {
    this.sumatoriaCuadrados += cuadradoValorZ[i][j];
  }
}

console.log("Sumatoria de los cuadrados de los valores Z:", this.sumatoriaCuadrados);
}
else{
  console.log("Hay valores de la frecuencia esperada que son menores 5");
  //this.alertError("Hay valores de la frecuencia esperada que son menores 5");
}
  

  },
  calcularValorP(){
    //los grados de libertad se calcula como v=(m-1)*(n-1), donde m corresponde al número de filas y n al número de columnas
    const m = this.questions[this.uno].values.length;
    const n = this.questions[this.dos].values.length
    this.gradosDeLibertad=((m-1)*(n-1));
    console.log("grados de libertad", this.gradosDeLibertad);
    this.valorP= ( 1 - jStat.chisquare.cdf(this.sumatoriaCuadrados, this.gradosDeLibertad)).toFixed(5);
    console.log("el valor p es:", this.valorP);
  },
  //calcular el coeficiente de Cramer
  coefCramer(){
    this.totalFrecuencias=0;
    this.cramer=0;
    const minimo = (Math.min(this.questions[this.uno].values.length, this.questions[this.dos].values.length))-1;
    for (let i = 0; i <this.matrizFrecuencias.length; i++) {
      for (let j = 0; j <this.matrizFrecuencias[0].length; j++) {
        this.totalFrecuencias +=  this.matrizFrecuencias[i][j];
      }
    }
    console.log("HOLAAtotal", this.totalFrecuencias);
    console.log("minimo", minimo);
    this.cramer=Math.sqrt((this.sumatoriaCuadrados/(this.totalFrecuencias*minimo))).toFixed(5);
    console.log("coef Cramer", this.cramer);


  },
 
openModal() {
    this.$refs.modalInformationChi.show();
  },
  closeModal() {
    this.showModal = false;
  },
  },    
  created() {
    this.matrizFrecuenciasRespuestas();
    this.calcularFrecuenciasEsperadas();
    this.estadisticoPruebaChiCuadrado();
    this.calcularValorP();
    this.coefCramer();
  },
  watch: {
    answersReports(newValue) {
      this.answers = newValue;
      this.matrizFrecuenciasRespuestas();
      this.estadisticoPruebaChiCuadrado();
      this.calcularValorP();
      this.calcularFrecuenciasEsperadas();
  
    },       
    idQuestionSelected1(newValue) {
      this.uno = newValue;    
      this.estadisticoPruebaChiCuadrado();
      this.matrizFrecuenciasRespuestas();
      this.calcularValorP();
      this.calcularFrecuenciasEsperadas();
      

    },
    idQuestionSelected2(newValue){
      this.dos = newValue;
      this.estadisticoPruebaChiCuadrado();
      this.matrizFrecuenciasRespuestas();
      this.calcularValorP();
      this.calcularFrecuenciasEsperadas();

    },
    questionsForm(newValue){
    this.questions = newValue;
    this.estadisticoPruebaChiCuadrado();
    this.matrizFrecuenciasRespuestas();
    this.calcularValorP();
    this.calcularFrecuenciasEsperadas();
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
  .styled-table {
    border-collapse: collapse;
    width: 100%;
    max-width: 800px;
    margin: 0 auto;
    text-align: center;
  }

  .styled-table th, .styled-table td {
    border: 1px solid #ddd;
    padding: 8px;
  }

  .styled-table th {
    background-color: #f2f2f2;
  }

  .styled-table tr:nth-child(even) {
    background-color: #f2f2f2;
  }

  .styled-table tr:hover {
    background-color: #ddd;
  }
  
  </style>  