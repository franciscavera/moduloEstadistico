<template> 
  <div>
    <h2><u>Cálculo Coeficiente de Correlación de Pearson</u></h2>
    <h3> Formularios disponibles</h3>
      
      <select style="width: 50%;" @click="getQuestionsForm(idFormSelected)" v-model="idFormSelected">
      <option disabled value="">Seleccione un formulario</option>
      <!-- <option  v-for="(opt, i) in allForms" :value="opt.id" >{{ opt.title }} {{ opt.id }}</option>  -->
       <option  v-for="(opt, i) in allForms" :value="opt.id" >{{ opt.title }}</option>   
      </select>
     
      <div class="col-6">
        <!--
  <h2>Id del form seleccionado vista correlacion:</h2>
  {{idFormSelected}} 
<h2> {{ answersReports }}</h2>
-->
<div v-if="idFormSelected !=''">


  <div v-if="questionsTypeNumber.length > 0">
    <h3> Seleccione una pregunta</h3>
  <!--<h2> {{ questionsForm }} </h2>-->
  <select style="width: 50%;" v-model="idQuestionSelected1">
      <option disabled value="">Seleccione una pregunta </option>
      <option  v-for="(question, i) in questionsTypeNumber" :value="question.index" >{{ question.subtitle }}</option> 
  </select>
      <!-- 
      <h2>Id del form seleccionado vista correlacion:</h2>
  {{idQuestionSelected1}}-->

  <h3> Seleccione una pregunta </h3>
  <!-- <h2> {{ questionsForm }} </h2>-->
  <select style="width: 50%;" v-model="idQuestionSelected2">
      <option disabled value="">Seleccione una pregunta </option>
      <option  v-for="(question, i) in questionsTypeNumber" :value="question.index" >{{ question.subtitle }}</option> 
  </select>
      <!-- 
      <h2>Id del form seleccionado vista correlacion:</h2>
  {{idQuestionSelected2}}-->


  
<div v-if="idQuestionSelected1 === '' && idQuestionSelected2 === ''">
  <!--<h3> Ambos están vacíos</h3>-->
</div>
<div v-else-if="idQuestionSelected1 !== '' && idQuestionSelected2 === ''">
  <!-- <h3> idQuestionSelected1 no está vacío, idQuestionSelected2 está vacío</h3>-->
</div>
<div v-else-if="idQuestionSelected1 === '' && idQuestionSelected2 !== ''">
  <!--<h3> idQuestionSelected1 está vacío, idQuestionSelected2 no está vacío</h3> -->
</div>
<div v-else-if="idQuestionSelected1 == idQuestionSelected2"> 
<h2 style="color: red;"> Las preguntas seleccionadas deben ser distintas</h2>
</div>
<div v-else>
  <h3>Cálculo del coeficiente de correlacion realizado con datos de  {{ answersReports[0].respuestas.length }} reportes que fueron obtenidos hasta {{ currentDate }}, a las {{ currentTime }} hrs.</h3>
  <!-- <h3> Ambos están llenos</h3>-->
  <calculate-correlation
  :idQuestionSelected1 = idQuestionSelected1
  :idQuestionSelected2 = idQuestionSelected2
  :questionsForm = questionsAllForm
  :answersReports = answersReports
  >
  </calculate-correlation>
</div>
</div>
<div v-else>
<h2> El formulario seleccionado no tiene preguntas con respuestas numéricas</h2>
</div>
</div>
</div>
</div>
</template>
<script>
  
  import PaginatedList from "@/components/PaginatedList.vue";
  import CalculateCorrelation from "../components/CalculateCorrelation.vue";
  import SimplePie from "../components/SimplePie.vue";
  import UniqueChart from "../components/UniqueChart.vue";
  import UniqueHorizontalChart from "../components/UniqueHorizontalChart.vue";
  import UniqueSelectChart from "../components/UniqueSelectChart.vue";
  import NumberChart from "../components/NumberChart.vue";
  import MultipleChart from "../components/MultipleChart.vue";
  import LikertChart from "../components/LikertChart.vue";
  import { params } from "@/components/mixin/params.js";
  import { pushNotification } from "../components/mixin/pushNotification.js";
  export default {
  mixins: [params, pushNotification],
  components: {
    PaginatedList,
    CalculateCorrelation,
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
      questionsAllForm: "",
        idFormSelected: "",
        idQuestionSelected1: "",
        idQuestionSelected2: "",
        answersReports: [],
        questionsTypeNumber: [],
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
    getQuestionsForm: async function(idFormSelected){
      this.$emit("getQuestionsForm", idFormSelected);
      
      try{
        let response = await this.$http.get(`questionsForm/` + idFormSelected);
        console.log("esto tiene response.data", response.data);
        this.questionsAllForm = response.data
        console.log("esto tiene QuestionForm", this.questionsAllForm);
        let responses = await this.$http.get(`allAnswersReports/` + idFormSelected);
      this.answersReports = responses.data
       console.log("esto tiene answers", this.answersReports);
      }catch(e){
        this.error = true;
        console.log('error while fetching forms', e)
      }
      this.questionsTypeNumber=[];
      for(let i=0;i<this.questionsAllForm.length;i++)
      {
        if(this.questionsAllForm[i].type == "number")
        {
          this.questionsTypeNumber.push(this.questionsAllForm[i]);
        }
      }
      this.currentTime= '';
      const now = new Date();
      this.currentDate = now.toLocaleDateString();
      this.currentTime = now.toLocaleTimeString();
    
      console.log("Question type NUMBER", this.questionsTypeNumber);
    }
  },
  beforeDestroy() {
  },
  computed: {
  },
  created() {
    this.getAllForms();
  },
  mounted(){
    this.getQuestionsForm();
  },
  watch: {
    idFormSelected(newValue) {
      this.idFormSelected = newValue;
      // Va a limpiar las selecciones del select 1 y select2
      this.idQuestionSelected1 = "";
      this.idQuestionSelected2 = "";
      this.questionsTypeNumber = [];
    },
  },
  };
  </script>
