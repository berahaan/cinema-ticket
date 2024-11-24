import { useNuxtApp } from "#app";
import { GET_GENRES } from "../components/graphql/queries/GET_GENRES.graphql";
export const useFetchGenres = () => {
  const { $apollo } = useNuxtApp();
  const genres = ref([]);
  const fetchGenres = async () => {
    try {
      const response = await $apollo.query({ query: GET_GENRES });
      genres.value = response.data.genres;
    } catch (error) {
      console.error("Error fetching genres:", error);
    }
  };
  return {
    fetchGenres,
    genres,
  };
};
