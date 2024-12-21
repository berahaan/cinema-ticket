export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      graphqlEndpoint: process.env.GRAPHQL_ENDPOINT,
    },
  },
  modules: ["@pinia/nuxt"],
  plugins: ["~/plugins/apollo.client.js", "~/plugins/loading.js"],
  css: ["~/assets/css/main.css"],
  build: {
    transpile: ["vue-toastification"],
  },
  nitro: {
    preset: "node-server",
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
});
