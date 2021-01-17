import Vue from "vue"
import App from "./App.vue"
import "./registerServiceWorker"
import Buefy from "buefy"
import "buefy/dist/buefy.css"
import axios from "axios"
import VueAxios from "vue-axios"

Vue.config.productionTip = false
Vue.use(Buefy)

Vue.use(VueAxios, axios.create({
	baseURL: "http://localhost:5050"
}))

const app = new Vue({
	render: h => h(App),
})

// mount it
app.$mount("#app")
