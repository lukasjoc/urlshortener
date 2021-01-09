import Vue from "vue"
import App from "./App.vue"
import "./registerServiceWorker"
import Buefy from "buefy"
import "buefy/dist/buefy.css"

Vue.config.productionTip = false
Vue.use(Buefy)

const app = new Vue({
  render: h => h(App),
})

// mount it
app.$mount("#app")
