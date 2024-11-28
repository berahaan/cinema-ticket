import { useNuxtApp } from "#app";
import { GET_MOVIES } from "../components/graphql/queries/GET_MOVIES.graphql";
import { GET_TITLE } from "../components/graphql/queries/GET_TITLE.graphql";
import { GET_MOVIE_BYDIRECTOR } from "../components/graphql/queries/GET_MOVIE_BYDIRECTOR.graphql";
import { GET_MOVIE_GENRES } from "../components/graphql/queries/GET_MOVIE_GENRES.graphql";
import { GET_MOVIE_TITLE_GENRE } from "../components/graphql/queries/GET_MOVIE_TITLE_GENRE.graphql";
import { GET_MOVIE_DIR_GENRE } from "../components/graphql/queries/GET_MOVIE_DIR_GENRE.graphql";
import { GET_MOVIE_TITLE_DIRECTOR } from "../components/graphql/queries/GET_MOVIE_TITLE_DIRECTOR.graphql";
import { GET_MOV_TIT_DIR_GEN } from "../components/graphql/queries/GET_MOV_TIT_DIR_GEN.graphql";
import { GET_MOVIEBY_SCHEDULE } from "../components/graphql/queries/GET_MOVIEBY_SCHEDULE.graphql";
import { GET_ALLMOVIEPARTS } from "../components/graphql/queries/GET_ALLMOVIEPARTS.graphql";
import { GET_MOVIE_TGDS } from "../components/graphql/queries/GET_MOVIE_T_G_D_Sch.graphql";
export const useFetchMovies = () => {
  const movieStore = useMoviesStore();
  const noMoviesFound = ref(false);
  const isloading = ref(false);
  const movies = ref([]);
  const searchQuery = ref("");
  const directorQuery = ref("");
  const genreQuery = ref("");
  const scheduleQuery = ref("");
  const currentPage = ref(1);
  const limit = 4;
  const totalPages = ref(1);
  const offset = computed(() => (currentPage.value - 1) * limit);
  const goToNextPage = () => {
    console.log("Here go nextpage is clicked now ");
    if (currentPage.value < totalPages.value) {
      currentPage.value += 1;
      fetchMovies();
    }
  };
  const goToPreviousPage = () => {
    console.log("Here go previous is clicked now ");
    if (currentPage.value > 1) {
      currentPage.value -= 1;
      fetchMovies();
    }
  };

  const fetchMovies = async () => {
    const { $apollo } = useNuxtApp();
    isloading.value = true;

    try {
      const today = new Date();
      let scheduleStart = null;
      let scheduleEnd = null;

      const variables = {
        limit,
        offset: offset.value,
      };
      if (scheduleQuery.value && scheduleQuery.value !== "All") {
        console.log("Schedule is selected:", scheduleQuery.value);
        if (scheduleQuery.value === "Today") {
          scheduleStart = scheduleEnd = today.toISOString().split("T")[0];
        } else if (scheduleQuery.value === "This Week") {
          const startOfWeek = new Date(today);
          startOfWeek.setDate(today.getDate() - today.getDay());
          const endOfWeek = new Date(startOfWeek);
          endOfWeek.setDate(startOfWeek.getDate() + 6);

          scheduleStart = startOfWeek.toISOString().split("T")[0];
          scheduleEnd = endOfWeek.toISOString().split("T")[0];
        } else if (scheduleQuery.value === "This Month") {
          scheduleStart = new Date(today.getFullYear(), today.getMonth(), 1)
            .toISOString()
            .split("T")[0];
          scheduleEnd = new Date(today.getFullYear(), today.getMonth() + 1, 0)
            .toISOString()
            .split("T")[0];
        }
        variables.scheduleStart = scheduleStart;
        variables.scheduleEnd = scheduleEnd;
      }

      const conditions = [];
      if (searchQuery.value) {
        conditions.push(`title: "%${searchQuery.value}%"`);
        variables.title = `%${searchQuery.value}%`;
      }
      if (genreQuery.value && genreQuery.value !== "All") {
        conditions.push(`genre: "${genreQuery.value}"`);
        variables.genres = genreQuery.value;
      }
      if (directorQuery.value && directorQuery.value !== "All") {
        conditions.push(`director: "${directorQuery.value}"`);
        variables.director = directorQuery.value;
      }
      if (
        scheduleQuery.value &&
        scheduleQuery.value !== "All" &&
        scheduleStart &&
        scheduleEnd
      ) {
        // Schedule filter condition to match dates between scheduleStart and scheduleEnd
        conditions.push(
          `scheduleDate_gte: "${scheduleStart}", scheduleDate_lte: "${scheduleEnd}"`
        );
        variables.scheduleStart = scheduleStart;
        variables.scheduleEnd = scheduleEnd;
      }

      let query;
      if (
        searchQuery.value &&
        genreQuery.value &&
        directorQuery.value &&
        scheduleQuery.value
      ) {
        console.log(
          "Fetching movies with title, genre, director, and schedule"
        );
        query = GET_MOVIE_TGDS;
      }
      // Check for different combinations of filters
      if (conditions.length == 4 && scheduleStart) {
        console.log("Fetching movies with title, genre, and schedule");
        query = GET_MOVIE_TGDS;
      }
      if (conditions.length === 3 && scheduleStart) {
        console.log("Fetching movies with title, genre, and schedule");
        query = GET_ALLMOVIEPARTS;
      } else if (conditions.length === 3) {
        console.log("Fetching movies with title, genre, and director");
        query = GET_MOV_TIT_DIR_GEN;
      } else if (conditions.length === 2) {
        console.log("Fetching movies with two active filters:", conditions);
        if (variables.title && variables.genres) {
          query = GET_MOVIE_TITLE_GENRE;
        } else if (variables.title && variables.director) {
          query = GET_MOVIE_TITLE_DIRECTOR;
        } else if (variables.genres && variables.director) {
          query = GET_MOVIE_DIR_GENRE;
        }
      } else if (conditions.length === 1) {
        console.log("Fetching movies with one active filter:", conditions);
        if (variables.title) {
          query = GET_TITLE;
        } else if (variables.genres) {
          query = GET_MOVIE_GENRES;
        } else if (variables.director) {
          query = GET_MOVIE_BYDIRECTOR;
        } else if (variables.scheduleStart && variables.scheduleEnd) {
          console.log("Fetching movies with selected schedule only");
          console.log(
            "here is shcedule ",
            variables.scheduleStart,
            "End in ",
            variables.scheduleEnd
          );
          query = GET_MOVIEBY_SCHEDULE; // Adjust query for schedule only
        }
      } else {
        console.log("Fetching all movies");
        query = GET_MOVIES;
      }

      const { data } = await $apollo.query({
        query,
        variables,
        fetchPolicy: "network-only",
      });
      movies.value = data.movies;
      console.log("Fetched movies:", movies.value);
      const totalMoviesCount = data.movies_aggregate.aggregate.count;
      totalPages.value = Math.ceil(totalMoviesCount / limit);
      if (movies.value.length === 0) {
        noMoviesFound.value = true;
        console.log(
          "Current pages",
          currentPage.value,
          "total Pages ",
          totalPages.value
        );
        movieStore.setMovies([]);
      } else {
        movieStore.setMovies(data.movies);
        movieStore.setTotalPages(Math.ceil(totalMoviesCount / limit));
        noMoviesFound.value = false; // Hide the message if movies are found
      }
    } catch (error) {
      console.error("Error fetching movies:", error);
    } finally {
      isloading.value = false;
    }
  };

  watch(
    [searchQuery, genreQuery, scheduleQuery, currentPage, directorQuery],
    fetchMovies,
    {
      immediate: true,
    }
  );
  return {
    fetchMovies,
    goToNextPage,
    scheduleQuery,
    genreQuery,
    currentPage,
    directorQuery,
    searchQuery,
    noMoviesFound,
    totalPages,
    isloading,
    limit,
    offset,
    goToPreviousPage,
  };
};
