<template>
<div>
      <h2>Consultar recetas</h2>
    <br>
    <form>
      <input class="buscadorID" type="text" placeholder="Ingrese la ID" v-model="id"> 
      <button class="botonAzul" @click="buscarReceta(id)">Buscar</button>
      <button class="botonAzul" @click="fetchRecetas">Lista</button> <br>
    </form>
    <div>   
    <ul>
      <li class="listaRecetas"
       v-for="(receta, index) in recetas" :key='index'> 
        <span class="recetaID">{{ receta.id }}</span> <span class="nombreReceta"> {{ receta.nombre }} </span>
        <div class="botones">
      <button class="botonActualizar" @click="openModalActualizar(receta.id, receta.nombre, receta.ingredientes, receta.instrucciones)"></button>
      <button class="botonEliminar" @click="openModalConfirmacion(receta.id)"></button> 
        </div>   
      <br>
      Ingredientes:
      <ul>
      <li class="listaIngredientes"
       v-for="(ingrediente, index) in receta.ingredientes" :key="index"> {{ingrediente}}</li></ul>
      <br>
      Instrucciones:
      <br>
      <span class="instruccion"> {{ receta.instrucciones }} </span>
      <br></li>  
    </ul>
    </div>
    <button class="buttonAgregar" @click="openModalAgregar"></button> 

    


<div id="modalAgregar" v-if="showModalAgregar">
  <div class="modal-backdrop">
    <div class="modal" role="dialog">
      <header class="modal-header">
        <slot name="header">

          <div>
    <h2>Agregar Receta</h2>
    <form>

        <label> Nombre </label>
        <input type="text" v-model="Receta.nombre">
        <br>
        <label> Ingredientes </label>
        <input type="text" v-model="ingrediente">
        <button class="agregarIngrediente" @click="agregarIngrediente"></button>
        <br>
        <ul>
      <li class="ingredientes" v-for="(data, index) in Receta.ingredientes" :key='index'>{{index}}. {{ data }}
        <button class="eliminarIngrediente" @click="eliminarIngrediente(index)"></button> 
      </li>
        </ul>
        <br>
        <label> Instrucciones </label>
        <textarea class="instrucciones" v-model="Receta.instrucciones"> </textarea>
        <br>
          <button class="botonAzul" @click="crearReceta">Agregar</button>
          <button class="botonRojo" @click="modalClose">Cancelar</button>
    </form> 
  </div>

        </slot>
      </header>
  </div>
  </div>
</div>


    
<div id="modalActualizar" v-if="showModalActualizar">
  <div class="modal-backdrop">
    <div class="modal" role="dialog">
      <header class="modal-header">
        <slot name="header">
          <div>
    <h2>Actualizar</h2>
    <form>
        Id: {{ id }}
        <br>
        <br>
        <label> Nombre </label>
        <input type="text" v-model="Receta.nombre">
        <br>
        <label> Ingredientes </label>
        <input type="text" v-model="ingrediente">
        <button class="agregarIngrediente" @click="agregarIngrediente"></button>
        <br>
        <ul>
      <li class="ingredientes" v-for="(data, index) in Receta.ingredientes" :key='index'>{{index}}. {{ data }}
        <button class="eliminarIngrediente" @click="eliminarIngrediente(index)"></button> 
      </li>
        </ul>
        <br>
        <label> Instrucciones </label>
        <textarea class="instrucciones" v-model="Receta.instrucciones"> </textarea>
        <br>
          
          <button class="botonAzul" @click="actualizarReceta(id)">Actualizar</button>
          <button class="botonRojo" @click="modalClose">Cancelar</button>
    </form>
  </div>
        </slot>
      </header>
  </div>
  </div>
</div>
<div id="modalConfirmacion" v-if="showModalConfirmacion">
  <div class="modal-backdrop">
    <div class="modal" role="dialog">
      <header class="modal-header">
        <slot name="header">

          <div class="receta">
    <p>Desea eliminar esta receta?</p>
          
          <button class="botonAzul2" @click="eliminarReceta(id)">Sí</button>
          <button class="botonRojo2" @click="showModalConfirmacion = false">No</button>

  </div>

        </slot>
      </header>
  </div>
  </div>
</div>
</div>

</template>

<script>
import axios from "axios";

export default {
  name: "Recetas",
  data: function() {
    return {
      recetas: [],
      ingrediente: "",
      id: "",
      Receta: { id: "", nombre: "", ingredientes: [], instrucciones: "" },
      showModalActualizar: false,
      showModalAgregar: false,
      showModalConfirmacion: false
    };
  },



  methods: {
    openModalActualizar: function(id, nombre, ingredientes, instrucciones) {
      this.showModalActualizar = true;
      (this.id = id),
      (this.Receta.nombre = nombre),
      (this.Receta.instrucciones = instrucciones),
      (this.Receta.ingredientes = ingredientes);
    },

    openModalAgregar: function() {
      this.showModalAgregar = true;
    },

    openModalConfirmacion: function(id) {
      this.showModalConfirmacion = true;
      this.id = id;
    },

    modalClose: function() {
      this.showModalActualizar = false;
      this.showModalAgregar = false;

      this.ingrediente = "";

      (this.Receta.id = ""),
        (this.Receta.nombre = ""),
        (this.Receta.instrucciones = []),
        (this.Receta.ingredientes = "");
    },

    agregarIngrediente: function() {
      console.log(this.ingrediente);
      this.Receta.ingredientes = this.Receta.ingredientes.concat(
        this.ingrediente
      );
      this.ingrediente = "";
    },

    eliminarIngrediente: function(index) {
      this.Receta.ingredientes.splice(index, 1);
    },

    crearReceta: function() {
      let nuevaReceta = {
        nombre: this.Receta.nombre,
        ingredientes: this.Receta.ingredientes,
        instrucciones: this.Receta.instrucciones
      };
      console.log(nuevaReceta);
      axios.post("http://localhost:8000/api/listaRecetas", nuevaReceta).then(
        response => {
          console.log(response);
          window.location.reload();
        },
        error => {
          console.log(error);
        }
      );
    },

    fetchRecetas: function() {
      axios.get("http://localhost:8000/api/listaRecetas").then(
        response => {
          this.recetas = response.data;
          console.log(response.data);
        },
        error => {
          console.log(error);
        }
      );
    },

    buscarReceta: function(id) {
      this.recetas = [];
      axios.get("http://localhost:8000/api/listaRecetas/" + id).then(
        response => {
          this.recetas = this.recetas.concat(response.data);

          console.log(this.recetas);
        },
        error => {
          console.log(error);
        }
      );
    },

    eliminarReceta: function(id) {
      axios.delete("http://localhost:8000/api/listaRecetas/" + id);
      window.location.reload();
    },

    actualizarReceta: function(id) {
      let nuevaReceta = {
        nombre: this.Receta.nombre,
        ingredientes: this.Receta.ingredientes,
        instrucciones: this.Receta.instrucciones
      };
      axios
        .put("http://localhost:8000/api/listaRecetas/" + id, nuevaReceta)
        .then(
          response => {
            window.location.reload();
          },
          error => {
            console.log(error);
          }
        );
    }
  },
  mounted: function() {
    this.fetchRecetas();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
ul {
  margin: 0;
  padding: 0;
  list-style-type: none;
  margin-bottom: 0.1em;
  margin-bottom: 0.4em;
}

.listaRecetas {
  padding: 15px;
  font-size: 1em;
  background-color: #5eadee;
  border-style: groove;
  margin-top: 0.8em;
  margin-bottom: 0.8em;
  color: #110f10;
  position: relative;
}

.listaIngredientes {
  padding: 5px;
  font-size: 1em;
  background-color: #4e87d3;
  margin-top: 0.2em;

  color: #110f10;
}

.instruccion {
  display: block;
  padding: 5px;
  font-size: 1em;
  background-color: #4e87d3;
  margin-top: 0.2em;

  color: #110f10;
}

p {
  text-align: center;
  padding: 30px 0;
  color: gray;
}

.buscadorID {
  width: 20%;
  padding: 5px 0;
}

.nombreReceta {
  font-size: 1.5em;
}

.recetaID {
  padding: 10px;
  font-size: 2.3em;
}

.instrucciones {
  width: 600px;
  height: 150px;
  vertical-align: top;
  margin-bottom: 10px;
}

.agregarIngrediente {
  background-image: url("add.png");
  background-repeat: no-repeat;
  background-position: center;
  background-color: white;
  height: 25px;
  width: 25px;
  margin-left: 10px;
}
.eliminarIngrediente {
  background-image: url("delete.png");
  background-repeat: no-repeat;
  background-position: center;
  background-color: white;
  height: 25px;
  width: 25px;
  margin-left: 10px;
}

.botonAzul {
  background-color: #6185d4;
  height: 30px;
  margin-left: 40px;
}

.botonAzul2 {
  background-color: #6185d4;
  height: 30px;
  margin-left: 10px;
}

.botonEliminar {
  background-image: url("deleteButton.png");
  background-color: transparent;
  height: 35px;
  width: 105px;
  border: none;
  background-repeat: no-repeat;
}

.botonActualizar {
  background-image: url("editButton.png");
  background-color: transparent;
  height: 35px;
  width: 104px;
  border: none;
  background-repeat: no-repeat;
  margin-right: 10px;
}

.botonRojo {
  background-color: #f05439;
  height: 30px;
  margin-left: 370px;
}

.botonRojo2 {
  background-color: #f05439;
  height: 30px;
  margin-left: 140px;
}

.buttonAgregar {
  background-image: url("addIcon.png");
  border: none;
  height: 70px;
  width: 70px;
  background-color: white;
  position: fixed;
  bottom: 50px;
  right: 60px;
}

.ingredientes {
  margin-left: 10px;
  margin-top: 5px;
  border-bottom: 1px solid;
}

.botones {
  position: absolute;
  right: 40px;
  top: 30px;
}

label {
  display: block;
  color: black;
}

input[type="text"] {
  width: 600px;
}

form {
  margin: 5px;
}

.modal-backdrop {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal {
  background: #ffffff;
  box-shadow: 2px 2px 20px 1px;
  overflow-x: auto;
  display: flex;
  flex-direction: column;
}

.modal-header,
.modal-footer {
  padding: 15px;
  display: flex;
  position: relative;
}

.modal-header {
  border-bottom: 1px solid #eeeeee;
  color: #4aae9b;
  justify-content: space-between;
}

.modal-footer {
  border-top: 1px solid #eeeeee;
  justify-content: flex-end;
}
</style>
