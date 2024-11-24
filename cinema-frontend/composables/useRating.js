import { jwtDecode } from "jwt-decode"; // Import without destructuring
import { useNuxtApp } from "#app";
import { useToast } from "vue-toastification";
import { INSERT_RATING } from "../components/graphql/mutations/INSERT_RATING.graphql";
import { UPDATE_RATES } from "../components/graphql/mutations/UPDATE_RATES.graphql";
import { GET_RATE_INFO } from "../components/graphql/queries/GET_RATE_INFO.graphql";
export function useRating() {
  const toast = useToast();
  const userRatings = ref([]);
  const isModal = ref(false);
  const selectedRating = ref(0);
  const currentMovieId = ref(null);
  const { $apollo } = useNuxtApp();
  const openRatingModal = (movieId) => {
    console.log("movieId in modals", movieId);
    const rate = userRatings.value[movieId];
    currentMovieId.value = movieId;
    isModal.value = true;
    selectedRating.value = rate || 0;
  };

  const submitRating = async (movieId) => {
    const accessToken = localStorage.getItem("accessToken");
    // Check if the token exists and is a valid string
    if (!accessToken) {
      console.error("Access token is missing or invalid.");
      return;
    }
    try {
      const decodedToken = jwtDecode(accessToken);
      const userId =
        decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

      console.log(
        "movieId here ",
        movieId,
        "And star values",
        selectedRating.value,
        "User id here is ",
        userId
      );
      isModal.value = false;
      currentMovieId.value = null;

      try {
        const checkRate = await $apollo.query({
          query: GET_RATE_INFO,
          variables: { userId: userId, movieId: movieId },
          fetchPolicy: "network-only",
        });
        const existingRate = checkRate.data.rate_movie;
        if (existingRate && existingRate.length > 0) {
          console.log("existingrate ", existingRate);
          const rateId = existingRate[0].rating_id;
          const rate = selectedRating.value;
          console.log(
            "rate info to be updated here ",
            "rating id ",
            existingRate[0].rating_id,
            "rate value now ",
            rate
          );

          const updateResponse = await $apollo.mutate({
            mutation: UPDATE_RATES,
            variables: { ratingId: rateId, rate: rate },
          });
          if (updateResponse && updateResponse.data) {
            toast.success("Updated successfully", {
              position: "top-center",
              duration: 2000,
            });
          }
          return;
        } else {
          const responseRate = await $apollo.mutate({
            mutation: INSERT_RATING,
            variables: {
              userId: userId,
              movieId: movieId,
              rating: selectedRating.value,
            },
            refetchQueries: [
              {
                query: GET_RATE_INFO,
                variables: { UserId: userId, movieId: movieId },
              },
            ],
          });
          console.log(
            "Here is response for rating information",
            responseRate.data
          );
          if (responseRate && responseRate.data) {
            toast.success("Rated successfully", {
              position: "top-center",
              duration: 4000,
            });
          }
        }
      } catch (error) {
        console.log("Error", error);
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
    currentMovieId,
  };
}
