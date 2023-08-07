<template>
    <div v-if="showModal" class="modal-container">
    
      <div class="modal-window">
      <h1> Ayuda</h1>
      <p style="text-align: left;"><strong>  Prueba Chi-Cuadrado de Pearson </strong></p>
      <p style="text-align: justify;"> La Prueba Chi-Cuadrado de Pearson, es una prueba que nos permite
        determinar si existe una relacion entre dos variables categóricas. Esta prueba se encarga de comparar
        las frecuencias observadas con las frecuencias esperadas.
     </p>
  
      <p style="text-align: left;"><strong> Tabla de Contingencia </strong> </p>
      <p style="text-align: justify;">La tabla de contingencia es una representación
        visual de cómo se distribuyen conjuntamente los datos en las diferentes categorías
        de ambas variables</p>
  
  
      <p style="text-align: left;"><strong>¿Existen condiciones para realizar la Prueba Chi-Cuadrado? </strong> </p>
      <p style="text-align: justify;">Si, existen dos condiciones que se deben cumplir. La primera es que las observaciones
    deben ser independientes entre sí y la segunda es que deben haber por lo menos 5 observaciones esperadas por conjunto. </p>
  
     
  
      <p style="text-align: left;"><strong> ¿Por qué se usan hipótesis? </strong> </p>
      <p style="text-align: justify;">Las hipótesis son muy importantes en pruebas estadísticas, ya que,
       con ella definimos el objetivo y proposito de la prueba. Siempre existirá una hipótesis nula (H0) y la hipótesis altern(HA) </p>
  
      <p style="text-align: left;"><strong> ¿Siempre se usan las mismas hipótesis? </strong> </p>
      <p style="text-align: justify;"> En este caso si, debido a que estamos utilizando la Prueba Chi-Cuadrado
        de Independencia . La hipótesis nula (H0) establece que no existe una relación significativa entre las dos variables
        categóricas seleccionadas (variables independientes) y la hipotesis alternativa (HA) establece que si existe una relación significativa
        entre las variables seleccionadas, es decir son "dependientes". </p>
  
      <p style="text-align: justify;"><strong>¿Qué son los grados de libertad?</strong> 
        Es un concepto estadístico que se refiere al número de valores que son libres de
        variar en un cálculo o estimación estadística. los grados de libertad están asociados con el número
         de observaciones o muestras que pueden variar independientemente en el cálculo del estadístico de prueba.</p>
      
      <p style="text-align: justify;"> <strong>¿Cómo se calculan los grados de libertad?</strong>En la prueba Chi-cuadrado de Pearson, 
        los grados de libertad se calculan a partir del número de filas (r) y el número de columnas (c) de la tabla de contingencia 
        utilizada en la prueba. La formula para calcular los grados de libertad es g=(r-1)*(c-1).</p>
  
      <p style="text-align: justify;"><strong>¿Cómo se interpretan los grados de libertad?</strong>
        Los grados de libertad en la Prueba Chi-Cuadrado representan el número de categorías que pueden variar
        de forma independiente en una tabla de frecuencias. Por ejemplo grados de libertad = 1, esto significa 
        que solo hay una categoría que puede variar de forma independiente. 
      </p>

        <p style="text-align: justify;"><strong>¿Quién define el valor de alfa (α)?</strong>
        El valor de alfa (α), lo define quien esta realizando la prueba estadística, depende del contexto y la naturaleza
        del estudio. Un nivel de significación de 0.05 es comunmente utilizado y se considera un estándar aceptado en 
        la investigación científica. </p>

        <p style="text-align: justify;"><strong>¿Qué significa α=0.05?</strong>
        El nivel de significación o mas bien conocido como alfa (α), es un valor comúnmente utilizado en estadística
        y pruebas de hipótesis. Representa la probabilidad máxima que estamos dispuestos a aceptar para cometer un error
        de tipo I (error de rechazo de la hipótesis nula cuando es verdadera). El nivel de significación se utiliza para
        determinar el punto de corte para rechazar la hipótesis nula</p>

        <p style="text-align: justify;"><strong>¿Qué es el Valor P?</strong>
        El valor P (probabilidad) se obtiene al realizar la prueba Chi-Cuadrado. Si el 
        valor p obtenido es menor o igual α (α=0.05) entonces se considera que el resultado es estadísticamente significativo, entonces
        significa que la probabilidad de obtener los resultados observados si la hipótesis nula es verdadera es menor o igual al 5%,
        por lo tanto rechazamos la hipótesis nula y aceptamos la hipótesis alternativa. En cambio si el valor p obtenido es 
        mayor a α (α=0.05) entonces no tenemos suficiente evidencia para rechazar la hipótesis nula y no podemos concluir 
        que hay una diferencia significativa entre  variables que estamos comparando.</p>

        <p style="text-align: justify;"><strong>¿Qué es el Coeficiente de Cramer (V) ?</strong>
         El coeficiente V de Cramer, también conocido como coeficiente de correlación nominal, es una medida de
         asociación utilizada para determinar qué tan fuerte es la relación entre dos variables categóricas.
         Es una extensión de la Prueba Chi-Cuadrado y proporciona información adicional sobre la magnitud o fuerza
        de la asociación entre las variables, más allá de simplemente indicar si están o no relacionadas.
        La Prueba Chi-Cuadrado permite determinar si existe una asociación significativa entre las variables categóricas, 
        pero no nos dice qué tan fuerte es esa asociación.
        </p>
         <p style="text-align: justify;"> Cuanto mayor sea el valor del Coeficiente (V) de Cramer, mayor será la asociación entre las variables. 
        Los valores de V se pueden encontrar entre 0 y 1.</p>

      
     
        <button class="btn btn-state" @click="closeModal">Cerrar</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        showModal: false,
      };
    },
    methods: {
      show() {
        this.showModal = true;
      },
      closeModal() {
          this.showModal = false;
      },
    },
  };
  </script>
  
  <style>
  /* Estilo para el container del modal */
  .modal-container {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5); /* Fondo semitransparente */
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
  }
  
  /* Estilo para la ventana modal */
  .modal-window {
    background-color: #ffffff;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
    max-width: 1000px;
    width: 1000px;
    text-align: center;
    max-height: 80vh; /* Altura máxima del modal (80% de la altura visible del viewport) */
    overflow-y: auto; /* Habilitar el scroll vertical*/
    /* ... */
    
  }
  img {
    width: 80%;
    height: auto; /* Ajusta automáticamente la altura proporcionalmente */
  }
  
  </style>
  