import { useNuxtApp } from "#app";
import { ref, computed } from "vue";
import { useRouter } from "vue-router";
import { GET_MOVIES } from "../components/graphql/queries/GET_MOVIES.graphql";
import { GET_TITLE } from "../components/graphql/queries/GET_TITLE.graphql";
export const useMoviesSchedule = () => {
  const { $apollo } = useNuxtApp();
  const router = useRouter();
  const searchQuery = ref("");
  const movies = ref([]);
  const currentPage = ref(1);
  const limit = 4;
  const totalPages = ref(1);
  const isLoading = ref(false);
  const offset = computed(() => (currentPage.value - 1) * limit);
  const fetchMovies = async () => {
    isLoading.value = true;
    if (searchQuery.value !== "") {
      isLoading.value = true;
      try {
        const { data } = await $apollo.query({
          query: GET_TITLE,
          variables: {
            limit,
            offset: offset.value,
            title: `%${searchQuery.value}%`,
          },
          fetchPolicy: "network",
        });

        movies.value = data.movies;
        totalPages.value = Math.ceil(
          data.movies_aggregate.aggregate.count / limit
        );
      } catch (error) {
        console.error("Error fetching movies:", error);
      } finally {
        isLoading.value = false;
      }
    } else {
      isLoading.value = true;
      try {
        const { data } = await $apollo.query({
          query: GET_MOVIES,
          variables: {
            limit,
            offset: offset.value,
          },
          fetchPolicy: "network",
        });
        movies.value = data.movies;
        totalPages.value = Math.ceil(
          data.movies_aggregate.aggregate.count / limit
        );
        useRequestFetch();
      } catch (error) {
        console.error("Error fetching movies:", error);
      } finally {
        isLoading.value = false;
      }
    }
  };
  watch([searchQuery], fetchMovies, {
    immediate: true,
  });
  const UpdateSchedules = (movieId) => {
    router.push(`/admin/schedule/${movieId}`);
  };
  const goToNextPage = () => {
    if (currentPage.value < totalPages.value) {
      currentPage.value += 1;
      fetchMovies();
    }
  };

  const goToPreviousPage = () => {
    if (currentPage.value > 1) {
      currentPage.value -= 1;
      fetchMovies();
    }
  };

  return {
    movies,
    currentPage,
    totalPages,
    isLoading,
    searchQuery,
    fetchMovies,
    goToNextPage,
    UpdateSchedules,
    goToPreviousPage,
  };
};
