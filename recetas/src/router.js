import Vue from 'vue'
import Router from 'vue-router'
import Recetas from './components/Recetas.vue'
import Agregar from './components/Agregar.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/agregar',
      name: 'agregar',
      component: Agregar
    },
    {
      path: '/lista',
      name: 'lista',
      component: Recetas
    }
  ]
})