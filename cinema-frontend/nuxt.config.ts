export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  apollo: {
    clients: {
      default: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
      },
    },
  },
  modules: ["@pinia/nuxt"],

  plugins: ["~/plugins/apollo.client.js", "~/plugins/loading.js"],

  css: ["~/assets/css/main.css"],

  build: {
    transpile: ["vue-toastification"],
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
});
