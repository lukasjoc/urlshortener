<template>
  <main id="app" class="section">
    <header class="level">
      <div class="level-left">
        <div class="level-item">
          <span class="title is-1">urlshortener</span>
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

