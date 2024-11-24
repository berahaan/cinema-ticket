import { useNuxtApp } from "#app";
import { GET_DIRECTORS } from "../components/graphql/queries/GET_DIRECTORS.graphql";
export const useFetchDirector = () => {
  const { $apollo } = useNuxtApp();
  const directors = ref([]);
  const fetchDirectors = async () => {
    try {
      const response = await $apollo.query({ query: GET_DIRECTORS });
      directors.value = response.data.directors;
    } catch (error) {
      console.error("Error fetching directors:", error);
    }
  };
  return {
    fetchDirectors,
    directors,
  };
};
