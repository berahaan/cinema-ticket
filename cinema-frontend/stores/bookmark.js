import { defineStore } from "pinia";
export const useBookmarkStore = defineStore("useBookmarkStore", {
  state: () => ({
    bookmarks: [],
  }),
  actions: {
    setBookmarks(bookmark) {
      this.bookmarks = bookmark;
      console.log("Bookmarks setted successfully now is ", this.bookmarks);
    },
  },
});
