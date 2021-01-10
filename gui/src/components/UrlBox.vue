<template>
  <section>
    <ValidationObserver ref="observer" v-slot="{ invalid }">
    <form @submit.prevent="onSubmit()" novalidate>
      <validation-provider
        name="url"
        rules="required|url"
        v-slot="{ errors }"
        >
        <b-field
          label="Enter URL:"
          :type="errors.length > 0 ? 'is-danger' : ''"
          :message="errors.length > 0 ? 'URL is invalid' : ''"
          >
          <b-input
            v-model="urlModel"
            placeholder="https://example.com"
            type="url"
            size="is-large"
            >
          </b-input>
        </b-field>
      </validation-provider>
        <b-button type="is-dark" native-type="submit">Convert</b-button>
    </form>
    </ValidationObserver>

    <div v-if="urlsConverted" class="section content convertedBox">
      <div v-for="(url, index) of urlsConverted" :key="index" class="level">
        <span class="link">{{ url.url }}</span>
        <b-button
          class="is-link is-dark is-text"
          @click="copyToClipBoard(url.urlConverted)"
          >{{ url.urlConverted }}</b-button
        >
      </div>
    </div>
  </section>
</template>

<script>
import { nanoid } from "nanoid"
import { ValidationProvider, extend, ValidationObserver } from "vee-validate"
import { required } from "vee-validate/dist/rules"

extend("required", {
	...required,
	message: "This field is required",
})
extend("url", {
	validate: (url) => {
		return /^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([-.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$/.test(
			url
		)
	},
	message: "URL is invalid.",
})

export default {
	name: "UrlBox",
	components: {
		ValidationProvider,
		ValidationObserver,
	},
	data() {
		return {
			urlModel: "",
			urlsConverted: [],
		}
	},
	methods: {
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
			this.urlsConverted.push({ url, urlConverted: `${base}${urlSpec}` })
		},
		async onSubmit() {
			const valid = await this.$refs.observer.validate()
			if (!valid) return
			this.minifyUrl(this.urlModel)
		},
		async copyToClipBoard(url) {
			await navigator.clipboard.writeText(url)
			if ((await navigator.clipboard.readText()) === url) {
				this.$buefy.toast.open({
					duration: 442,
					message: "Copied to clipboard!",
					position: "is-top-right",
				})
			}
		}
	},
}
</script>


