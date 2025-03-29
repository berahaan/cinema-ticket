import { defineStore } from "pinia";

export const useColorModeStore = defineStore("colorMode", {
  state: () => ({
    colorMode: "light",
  }),
  actions: {
    toggleColorMode() {
      this.colorMode = this.colorMode === "light" ? "dark" : "light";
      if (import.meta.client) {
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
