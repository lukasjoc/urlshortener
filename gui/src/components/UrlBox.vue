<template>
  <section class="section">
    <span class="subtitle is-3">Enter URL: </span>
    <b-field>
      <b-input v-model="urlModel" placeholder="https://" type="url"  size="is-large"></b-input>
    </b-field>
    <b-button @click.stop="minifyUrl(urlModel)" type="is-primary">Convert</b-button>
    <div v-if="urlsConverted" class="section content convertedBox">
      <div v-for="(url, index) of urlsConverted" :key="index" class="level">
        <span class="link">{{url.url}}</span>
        <b-button class="is-link is-text" @click="copyToClipBoard(url.urlConverted)">{{url.urlConverted}}</b-button>
      </div>
    </div>
  </section>
</template>

<script>
import { nanoid } from "nanoid"

export default {
  name: "UrlBox" ,
  data() {
    return {
      urlModel: "",
      urlsConverted: [],
    };
  },
  methods: {
    async copyToClipBoard(url) {
      await navigator.clipboard.writeText(url)
      if (await navigator.clipboard.readText() === url) {
        this.$buefy.toast.open({
          duration:  442,
          message: "Copied to clipboard!",
          position: "is-top-right",
        })
      }
    },
    minifyUrl(url) {
      // TODO: check if server is reachable first
      // TODO: check if url is available
      // TODO: add localstorage to save converted url
      const base = "http://localhost:6060/"
      const urlSpec = nanoid(8)
      // const payload = { url, urlSpec }
      // check availability
      // save if available else return false and return possible from server
      // let ok = checkAvailability(payload)
      this.urlsConverted.push({url, urlConverted: `${base}${urlSpec}`})
    },
  },
}
</script>


