import { GET_STARS } from "../components/graphql/queries/GET_STARS.graphql";
import { useNuxtApp } from "#app";
export const useFetchStars = () => {
  const { $apollo } = useNuxtApp();
  const stars = ref([]);
  const fetchStars = async () => {
    try {
      const response = await $apollo.query({ query: GET_STARS });
      stars.value = response.data.stars;
    } catch (error) {
      console.error("Error fetching stars:", error);
    }
  };
  return {
    fetchStars,
    stars,
  };
};
