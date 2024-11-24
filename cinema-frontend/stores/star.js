import { defineStore } from "pinia";
export const useStarStored = defineStore("star", {
  state: () => ({
    stars: [],
  }),
  actions: {
    setStars(star) {
      this.stars = star;
    },
  },
});
