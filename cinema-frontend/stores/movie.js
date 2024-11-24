import { defineStore } from "pinia";
export const useMoviesStore = defineStore("useMovieStore", {
  state: () => ({
    movies: [],
    totalPages: 0,
  }),
  actions: {
    setMovies(movie) {
      this.movies = movie;
      console.log("Movie setted successfully now is ", this.movies);
    },
    setTotalPages(pages) {
      this.totalPages = pages;
    },
  },
});
