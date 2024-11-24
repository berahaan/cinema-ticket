import { reactive } from "vue";

export default defineNuxtPlugin((nuxtApp) => {
  const loadingState = reactive({ isLoading: false });
  nuxtApp.provide("loading", {
    loadingState,
    setLoading(value) {
      loadingState.isLoading = value;
    },
  });
});
