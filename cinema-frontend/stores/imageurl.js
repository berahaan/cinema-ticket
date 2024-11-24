import { defineStore } from "pinia";
export const useImageStore = defineStore("imageurl", {
  state: () => ({
    featuredImageURL: "",
    otherImageURLs: [],
  }),
  actions: {
    setImageURLs(featuredURL, otherURLs) {
      console.log("Featured image url here in stores ", featuredURL);
      this.featuredImageURL = featuredURL;
      this.otherImageURLs = otherURLs;
    },
  },
});
