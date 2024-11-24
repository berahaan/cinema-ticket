import { defineStore } from "pinia";

export const useImageBase64Store = defineStore("imageBase64", {
  state: () => ({
    featuredImage64: "",
    otherImage64: [],
  }),
  actions: {
    setImageBase64(featuredImage64, otherImage64) {
      this.featuredImage64 = featuredImage64;
      this.otherImage64 = otherImage64;
    },
  },
});
