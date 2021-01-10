<template>
  <main id="app">
    <header class="hero is-primary is-bold">
      <div class="hero-body">
        <div class="container">
          <h1 class="title"> urlshortener </h1>
          <h2 class="subtitle"> Simple url shortener mono repo vue2 gui + laravel service </h2>
          <a href="https://github.com/lukasjoc/urlshortener" target="_blank">
            <b-button class="is-light is-text is-link">GitHub</b-button>
          </a>
        </div>
      </div>
    </header>
    <div class="section">
      <UrlBox />
    </div>
  </main>
</template>

<script>
import UrlBox from "./components/UrlBox"

export default {
	name: "App",
	components: {
		UrlBox,
	},
}
</script>

