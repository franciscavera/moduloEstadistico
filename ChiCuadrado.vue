<template> 
    <div>
      <h2> <u>Prueba Chi-Cuadrado y Coeficiente V de Cramer</u></h2>
      <h3> Formularios disponibles</h3>
        
        <select style="width: 50%;" @click="getQuestionsForm(idFormSelected)" v-model="idFormSelected">
        <option disabled value="">Seleccione un formulario</option>
        <!-- <option  v-for="(opt, i) in allForms" :value="opt.id" >{{ opt.title }} {{ opt.id }}</option>  -->
         <option  v-for="(opt, i) in allForms" :value="opt.id" >{{ opt.title }}</option>   
        </select>
       
        <div class="col-6">
          <!--
    <h2>Id del form seleccionado vista correlacion:</h2>
    {{idFormSelected}} -->
    <div v-if="idFormSelected !=''">
      <div v-if="questionsTypeUnique.length > 0">
   <!--<h2> {{ answersReports }}</h2>-->
      <h3> Seleccione una pregunta</h3>
    <!--<h2> {{ questionsForm }} </h2>-->
    <select style="width: 50%;" v-model="idQuestionSelected1">
        <option disabled value="">Seleccione una pregunta </option>
        <option  v-for="(question, i) in questionsTypeUnique" :value="question.index" >{{ question.subtitle }}</option> 
    </select>
        <!-- 
        <h2>Id del form seleccionado vista correlacion:</h2>
    {{idQuestionSelected1}}-->
  
    <h3> Seleccione una pregunta </h3>
    <!-- <h2> {{ questionsForm }} </h2>-->
    <select style="width: 50%;" v-model="idQuestionSelected2">
        <option disabled value="">Seleccione una pregunta </option>
        <option  v-for="(question, i) in questionsTypeUnique" :value="question.index" >{{ question.subtitle }}</option> 
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
    <h3>Analizando los datos de  {{ answersReports[0].respuestas.length }} reportes recopilados hasta Fecha: {{ currentDate }} Hora: {{ currentTime }}</h3>
    <!-- <h3> Ambos están llenos</h3>-->
    <calculate-chi-squared
    :idQuestionSelected1 = idQuestionSelected1
    :idQuestionSelected2 = idQuestionSelected2
    :questionsForm = questionsAllForm
    :answersReports = answersReports
    >
    </calculate-chi-squared>
  </div>
  </div>
  <div v-else>
    <h2> El formulario seleccionado no tiene preguntas categóricas</h2>
  </div>
  </div>
  </div>
  </div>
  </template>
  <script>
    
    import PaginatedList from "@/components/PaginatedList.vue";
    import CalculateChiSquared from "../components/CalculateChiSquared.vue";
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
      CalculateChiSquared,
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
          questionsTypeUnique: [],
          currentDate: '',
          currentTime: '',
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
        this.questionsTypeUnique=[];
        for(let i=0;i<this.questionsAllForm.length;i++)
        {
          if(this.questionsAllForm[i].type == "unique")
          {
            this.questionsTypeUnique.push(this.questionsAllForm[i]);
          }
        }
        this.currentDate='';
    this.currentTime= '';
      const now = new Date();
      this.currentDate = now.toLocaleDateString();
      this.currentTime = now.toLocaleTimeString();
      
        console.log("Question type NUMBER", this.questionsTypeUnique);
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
        this.questionsTypeUnique = [];
      },
    },
    };
    </script>
  