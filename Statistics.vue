<template>
  <main>
      <div class="dashboard-container">
          <h3>Gráficas de caracterización</h3>
          <!-- Se verifica si allForms > 0 (En allForms se guardan todos los formularios que hay en la base de datos)-->
          <div v-if="allForms.length > 0 "> 
          <!-- <h3> si hay formularios</h3>-->
          <!--Componente que muestra mediante un select los formularios -->
          <select style="width: 65%;"
        @click="getAnswersReports(idFormSelected), getQuestionCharacterization(idFormSelected)"
        v-model="idFormSelected"
        >
        <option disabled value="">Seleccione un formulario</option>
        <!--  <option  v-for="(opt, i) in allForms" :value="opt.id" >{{ opt.title }} {{ opt.id }}</option> -->
         <option  v-for="(opt, i) in allForms" :value="opt.id" >{{ opt.title }}</option> 
        </select>
          <!-- 
          <selected-form 
          :allForms = allForms
          @getAnswersReports="getAnswersReports"
          @getQuestionCharacterization = "getQuestionCharacterization"
           > 
          </selected-form>
          -->
           <!-- Cuando elige un formulario, se obtienen las respuestas de las preguntas de caracterización
          asociadas al formulario seleccionado, se pregunta si answersReports > 0 -->
          <div v-if="answersReports.length > 0 && idFormSelected!=''"> 
            <h3>Gráficas de caracterización generadas a partir de {{ answersReports[0].respuestas.length }} reportes que fueron obtenidos hasta {{ currentDate }},  a las  {{ currentTime }} hrs.</h3>
            <!-- 
            <p> arreglo de respuestas</p>
             {{ answersReports }}
            <p> arreglo de preguntas</p>
            {{ questionCharacterization }}
              -->
            <div class="dashboard-columns"> 
        <!--Con este for, se recorre el arreglo que contiene las preguntas de caracterizacion-->
        <div class="graph-container unique-column" v-for="(question, i) in questionCharacterization" :key="i">
          <!-- Con este if se pregunta si question.type es igual a un type especifico. 
               Los type pueden ser: unique, uniqueHorizontal, uniqueSelect, number-->
         
          <!-- type === unique -->
          
          <div class="unique-graph-content" style="background-color: #F3F3F3; height:503;"  v-if="question.type === 'unique'" >
            <!-- <p> la pregunta index: {{question.index}} es de tipo UNIQUE</p> -->
            <!-- Componente que graficará las preguntas de type=unique-->
            <p> la pregunta index: {{question.type}} </p>
            <unique-chart
            :answersReports="answersReports[i].respuestas"
            :questionCharacterization="questionCharacterization[i]">
           </unique-chart>
          </div>

          <!-- type === uniqueHorizontal -->
          <div style="background-color:#F3F3F3"  v-if="question.type === 'uniqueHorizontal'" >
            <p> la pregunta index: {{question.type}} </p>
             <!-- 
            <p> la pregunta index: {{question.index}} es de tipo uniqueHorizontal</p>
            <p> respuestas: {{answersReports[i]}}</p>
            <p> pregunta: {{ questionCharacterization[i]}}</p>
            -->
          <!-- Componente que graficará las preguntas de type=unique-->
            <unique-horizontal-chart
            :answersReports="answersReports[i].respuestas"
            :questionCharacterization="questionCharacterization[i]"> 
            </unique-horizontal-chart>
          </div>

          <!-- type === uniqueSelect -->
          <div style="background-color: #F3F3F3"  v-if="question.type === 'uniqueSelect'"> 
            <p> la pregunta index: {{question.type}} </p>
            <!-- 
            <p> la pregunta index: {{question.index}} </p>
            <p> respuestas: {{answersReports[i]}}</p>
            <p> pregunta: {{ questionCharacterization[i]}}</p>
            -->
             <!-- Componente que graficará las preguntas de type=uniqueSelect-->
            <unique-select-chart
            :answersReports="answersReports[i].respuestas"
            :questionCharacterization="questionCharacterization[i]"
            >
            </unique-select-chart>
          </div>

          <!-- type === number -->
          <div style="background-color: #F3F3F3" v-if="question.type === 'number'">
            <p> la pregunta index: {{question.type}} </p>
            <!-- 
            <p> la pregunta index: {{question.index}} Number</p>
            <p> respuestas: {{answersReports[i]}}</p>
            <p> pregunta: {{ questionCharacterization[i]}}</p>
            -->
             <!-- Componente que graficará las preguntas de type= number-->
            <number-chart
            :answersReports="answersReports[i].respuestas"
            :questionCharacterization="questionCharacterization[i]"
            >
            </number-chart>


          </div>
          <!-- type === multiple -->
          <div style="background-color:#F3F3F3"  v-if="question.type === 'multiple'">
            <p> la pregunta index: {{question.type}} </p>
            <!-- 
            <p> la pregunta index: {{question.index}} multiple</p>
            <p> respuestas: {{answersReports[i]}}</p>
            <p> pregunta: {{ questionCharacterization[i]}}</p>
            -->
             <!-- Componente que graficará las preguntas de type= multiple-->
            <multiple-chart
            :answersReports="answersReports[i].respuestas"
            :questionCharacterization="questionCharacterization[i]"
            >
            </multiple-chart>


          </div>

           <!-- type === likert -->
           <div style="background-color: #F3F3F3" v-if="question.type === 'likert'">
            <p> la pregunta index: {{question.type}} </p>
            <!-- 
            <p> la pregunta index: {{question.index}} likert</p>
            <p> respuestas: {{answersReports[i]}}</p>
            <p> pregunta: {{ questionCharacterization[i]}}</p>
            -->
             <!-- Componente que graficará las preguntas de type= likert-->
            <likert-chart
            :answersReports="answersReports[i].respuestas"
            :questionCharacterization="questionCharacterization[i]"
            >
            </likert-chart>


          </div>
        </div>
        </div>
        <!--
          <my-graphs
          :answersReports="answersReports"
          :questionCharacterization="questionCharacterization">
          
          </my-graphs> -->


         <!--  <div>
        <p> arreglo de respuestas</p>
        {{ answersReports}}
        
        </div>-->
          </div>
          
         
          </div>

      </div>
      
      <div>
      </div>
      
  </main>
</template>
<script>

import PaginatedList from "@/components/PaginatedList.vue";
import SelectedForm from "../components/SelectedForm.vue";
import SimplePie from "../components/SimplePie.vue";
import UniqueChart from "../components/UniqueChart.vue";
import UniqueHorizontalChart from "../components/UniqueHorizontalChart.vue";
import UniqueSelectChart from "../components/UniqueSelectChart.vue";
import NumberChart from "../components/NumberChart.vue";
import MultipleChart from "../components/MultipleChart.vue";
import LikertChart from "../components/LikertChart.vue";
import { user } from "@/components/mixin/user.js";
import { params } from "@/components/mixin/params.js";
import { pushNotification } from "../components/mixin/pushNotification.js";

export default {
mixins: [params, pushNotification,user],
components: {
  PaginatedList,
  SelectedForm,
  SimplePie,
  UniqueChart,
  UniqueHorizontalChart,
  UniqueSelectChart,
  NumberChart,
  MultipleChart,
  LikertChart,
},
data() {
  return {
    allForms: [],
    answersReports: [],
    questionCharacterization:[],
    currentDate: '',
    currentTime: '',
    idFormSelected:"",

  };
},
methods: {
  getAllForms: async function() {
    try{
      let response = await this.$http.get(`forms`);
      console.log("esto tiene forms");
      this.allForms = response.data
    }catch(e){
      this.error = true;
      console.log('error while fetching forms', e)
    }
    //console.log("todos los form", this.allForms);
  },
  getAnswersReports: async function(id) {
    this.answersReports=[];
    try{
      let response = await this.$http.get(`reports/` + id);
      console.log("esto tiene forms");
      this.answersReports = response.data
      console.log("AnswersReports", this.answersReports);
    }catch(e){
      this.error = true;
      console.log('error while fetching forms', e)
    }
    if (this.answersReports.length > 0) {
        this.selectedFormId = id; // Establecer el ID del formulario seleccionado
      }
    this.currentDate='';
    this.currentTime= '';
      const now = new Date();
      this.currentDate = now.toLocaleDateString();
      this.currentTime = now.toLocaleTimeString();
    if(this.answersReports.length ==0 && this.idFormSelected!=''){
      this.alertError(" El formulario seleccionado no tiene reportes.");
    }
   
  }, 
  getQuestionCharacterization: async function(id){
    this.questionCharacterization=[];
    try{
      let response = await this.$http.get(`question-characterization/` + id);
      console.log("esto tiene forms");
      this.questionCharacterization = response.data
      console.log("esto tiene questionCharact", this.questionCharacterization);
    }catch(e){
      this.error = true;
      console.log('error while fetching forms', e)
    }
  },


},
beforeDestroy() {
},
computed: {
},
created() {
  this.getAllForms();
},
mounted(){
  this.getAnswersReports();
  this.getQuestionCharacterization();

  
}
};
</script>
<style>
.dashboard-container {
  margin: 20px;
}

.dashboard-columns {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  gap: 20px;

}

.graph-container {
  flex: 1 1 45%; /* Ocupa un 45% del espacio disponible en el contenedor */
}

.graph {
  width: 100%;
}

.unique-column {
  order: 1; /* Orden en el que se muestra */
}

.right-column {
  order: 2; /* Orden en el que se muestra */
}
.unique-graph-content {
  height: 100%; /* O la altura que desees para el contenido del gráfico unique */
}

</style>