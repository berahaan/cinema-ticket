import { jwtDecode } from "jwt-decode"; // Import without destructuring
import { useNuxtApp } from "#app";
import { useToast } from "vue-toastification";
import { INSERT_RATING } from "../components/graphql/mutations/INSERT_RATING.graphql";
import { UPDATE_RATES } from "../components/graphql/mutations/UPDATE_RATES.graphql";
import { GET_RATE_INFO } from "../components/graphql/queries/GET_RATE_INFO.graphql";
export function useRating() {
  const toast = useToast();
  const userRatings = ref([]);
  const { fetchMovies } = useFetchMovies();
  const isModal = ref(false);
  const israte = ref(false);
  const selectedRating = ref(0);
  const currentMovieId = ref(null);
  const { $apollo } = useNuxtApp();

  const openRatingModal = (movieId) => {
    const rate = userRatings.value[movieId];
    currentMovieId.value = movieId;
    israte.value = !israte.value;
    console.log("Rated", rate);
    selectedRating.value = rate || 0;
  };

  const submitRating = async (movieId) => {
    const accessToken = localStorage.getItem("accessToken");

    if (!accessToken) {
      console.error("Access token is missing or invalid.");
      return;
    }

    try {
      const decodedToken = jwtDecode(accessToken);
      const userId =
        decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

      israte.value = false;
      currentMovieId.value = null;

      try {
        const checkRate = await $apollo.query({
          query: GET_RATE_INFO,
          variables: { userId: userId, movieId: movieId },
          fetchPolicy: "network-only",
        });

        const existingRate = checkRate.data.rate_movie;
        const rate = selectedRating.value;

        if (existingRate && existingRate.length > 0) {
          // Update existing rating
          const rateId = existingRate[0].rating_id;
          const updateResponse = await $apollo.mutate({
            mutation: UPDATE_RATES,
            variables: { ratingId: rateId, rate: rate },
          });

          if (updateResponse && updateResponse.data) {
            // Update userRatings locally
            userRatings.value[movieId] = rate;
            await fetchMovies();
            toast.success("Updated successfully", {
              position: "top-center",
              duration: 2000,
            });
          }
        } else {
          // Insert new rating
          const responseRate = await $apollo.mutate({
            mutation: INSERT_RATING,
            variables: {
              userId: userId,
              movieId: movieId,
              rating: rate,
            },
          });

          if (responseRate && responseRate.data) {
            // Add new rating to userRatings locally
            userRatings.value[movieId] = rate;
            await fetchMovies();
            toast.success("Rated successfully", {
              position: "top-center",
              duration: 4000,
            });
          }
        }
      } catch (error) {
        console.error("Error:", error);
      }
    } catch (error) {
      console.error("Error decoding token:", error);
    }
  };

  return {
    openRatingModal,
    isModal,
    userRatings,
    selectedRating,
    submitRating,
    israte,
    currentMovieId,
  };
}
