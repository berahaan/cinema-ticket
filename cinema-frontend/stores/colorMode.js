// stores/colorMode.js
import { defineStore } from "pinia";

export const useColorModeStore = defineStore("colorMode", {
  state: () => ({
    colorMode: "light",
  }),
  actions: {
    toggleColorMode() {
      console.log("toggle color mode is called now");
      this.colorMode = this.colorMode === "light" ? "dark" : "light";
      if (import.meta.client) {
        console.log("here color is setted now ");
        localStorage.setItem("colorMode", this.colorMode);
      }
    },
    loadColorMode() {
      if (import.meta.client) {
        const savedMode = localStorage.getItem("colorMode") || "light";
        if (savedMode) {
          this.colorMode = savedMode;
        }
      } else {
        this.colorMode = "light";
      }
    },
    resetColorMode() {
      this.colorMode = "light";
      localStorage.removeItem("colorMode");
      document.body.className = "";
    },
  },
});
