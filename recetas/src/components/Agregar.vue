<template>
  <div class="receta">
    <h2>Agregar recetas</h2>
    <form>
        <label for="inputNombre"> Nombre </label>
        <input type="text" id="inputNombre" v-model="Receta.nombre">
        <br>
        <label for="inputIngrediente"> ingrediente </label>
        <input type="text" id="inputIngrediente" v-model="ingrediente">
        <button @click="agregarIngrediente">Add</button>
        <br>
        <ul>
      <li v-for="(data, index) in ingredientes" :key='index'>{{index}}. {{ data.ingrediente }}</li>
        </ul>
        <br>
        <label for="inputInstrucciones"> instrucciones </label>
        <input type="text" id="inputInstrucciones" v-model="Receta.instrucciones">
        <br>

        <button @click="crearReceta">Crear!</button>

    </form>
    
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Agregar",
  data: function() {
    return {
      ingrediente: '',
      ingredientes: [],
      Receta: { nombre: "", ingredientes: [], instrucciones: "" }
    };
  },
  methods: {
    agregarIngrediente: function(){
        this.ingredientes.push({ingrediente: this.ingrediente})
        this.ingrediente = '';
    },
    crearReceta: function() {
      let nuevaReceta = {
        nombre: this.Receta.nombre,
        ingredientes: this.ingredientes,
        instrucciones: this.Receta.instrucciones
      };

      console.log(nuevaReceta);

      axios.post("http://localhost:8000/api/listaRecetas", nuevaReceta).then(
        response => {
          console.log(response);
        },
        error => {
          console.log(error);
        }
      );
    }
  },
};
</script>