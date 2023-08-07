<template>
  <div v-if="showModal" class="modal-container">
  
    <div class="modal-window">
    <h1> Ayuda aclaración de conceptos......</h1>
    <p style="text-align: left;"><strong>  Correlación Lineal </strong></p>
    <p style="text-align: justify;">Cuando se emplea el análisis de correlación,
    es posible determinar si existe una relación entre las variables, así como
    la forma, el sentido y el grado en que se relacionan entre sí (Grasso, 2006). </p>

    <p style="text-align: left;"><strong>¿Qué es el Coeficiente de Correlación de Pearson? </strong> </p>
    <p style="text-align: justify;">"Es una medida numérica de la fuerza de la relación
    lineal entre dos variables. El coeficiente refleja la consistencia del efecto
    que un cambio en una variable tiene sobre la otra”  (Johnson & Kuby, 2015, p.136). </p>

    <p style="text-align: justify;"> Para calcular el Coeficiente de Correlación de Pearson, necesitamos al menos dos pares de datos, 
      ejemplo (X1,Y1) y (X2,Y2). Mientras más grande el tamaño de la muestra el Coeficiente de Correlación Lineal 
      nos entrega una resultado más estable.
    </p>

    <p style="text-align: left;"><strong>¿Qué valores podemos obtener como resultado del Coefiente de Correlación? </strong> </p>
    <p style="text-align: justify;">El resultado del Coefiente de Correlación de Pearson tendra 
        un valor entre -1 y 1.  </p>

    <p style="text-align: left;"><strong>  ¿Qué tipo de variables necesitamos? </strong> </p>
    <p style="text-align: justify;">Para calcular el Coeficiente de Correlación Lineal,
    necesitamos dos variables cuantitativas, en otras palabras se necesitan variables
    que se midan en una escala númerica. Comunmente las variables involucradas reciben
    el nombre de variable independiente (también llamada X) y variable dependiente
    (también llamada Y) </p>

    <p style="text-align: left;"><strong> ¿Qué representa el Gráfico o Diagrama de Dispersión? </strong> </p>
    <p style="text-align: justify;">Nos permite representar la relación entre dos variables,
    debido a que cada punto corresponde a un par de valores (X,Y) de las variables analizadas. En
    el eje x se presentan los valores de la variable independiente o "X"y en el eje y los valores
     de la variable dependiente o "Y"</p>

    <p style="text-align: left;"><strong> ¿Comó interpretar un Gráfico o Diagrama de Dispersión? </strong> </p>
    <p style="text-align: justify;">La interpretación del Diagrama de Disperción va a depender
    del patrón y forma general de los puntos.</p>

    <p style="text-align: justify;"><strong> Relación lineal positiva:</strong> Los pares de puntos (X,Y) tenderán agruparse de forma lineal y 
    ascenderan desde la esquina inferior izquierda hasta la esquina superior derecha. Además 
    que a medida que X aumenta Y tambien aumenta de forma proporcional.</p>
    
    <p style="text-align: justify;"> <strong>Relación lineal negativa:</strong>Los pares de puntos (X,Y) tenderán agruparse de forma lineal y 
    descendente desde la esquina superior izquierda hasta la esquina inferior derecha. Además 
    que a medida que X aumenta Y tiende a disminuir.  </p>

    <p style="text-align: justify;"><strong>Relación no lineal: </strong>Los pares de puntos (X,Y) no siguen un patron definido, lo 
    que implica que no existe relación lineal entre las variables. </p>

    <img src="@/assets/img/NubeDePuntos.png" />
    <p style="text-align: center;">Ejemplos de Diagramas de Dispersión y Coefiente de Correlación
        Lineal (r) entre dos variables. Fuente (Dagnino, 2014).</p>
    
    <h3 style="text-align: justify;"> Es importante considerar que una correlación entre
    dos variables, no implica una relación de causalidad entre ellas. El resultado que obtenemos
    al calcular el coefiente de correlación solo mide la fuerza y dirección de la relación lineal, pero 
    no podemos asegurar que una variable cause cambios en la otra. </h3>
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
