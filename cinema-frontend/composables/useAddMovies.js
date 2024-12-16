import { ADD_MOVIES } from "../components/graphql/mutations/ADD_MOVIES.graphql";
import { GET_MOVIES } from "../components/graphql/queries/GET_MOVIES.graphql";
import { useNuxtApp } from "#app";
import { useToast } from "vue-toastification";
import { ref, onMounted } from "vue";
export const useAddMovies = () => {
  const toast = useToast();
  const { fetchStars, stars } = useFetchStars();
  const { fetchGenres, genres } = useFetchGenres();
  const selectedStars = ref([]);
  const loading = ref(false);
  const isSpinning = ref(false);
  const selectedStar = ref("");
  const selectedGenres = ref([]);
  const selectedGenre = ref("");
  const { $apollo } = useNuxtApp();
  const formData = ref({
    title: "",
    description: "",
    duration: 0,
    director_id: null,
    star_id: null,
    selected_genre: null,
    featured_image: null,
    other_images: [],
  });

  onMounted(async () => {
    await fetchStars();
    await fetchGenres();
  });
  const addMovies = async (
    title,
    description,
    duration,
    director_id,
    selectedStars,
    selectedGenres,
    featuredImageurl,
    otherImageURL
  ) => {
    loading.value = true;
    isSpinning.value = true;
    try {
      const response = await $apollo.mutate({
        mutation: ADD_MOVIES,
        variables: {
          title,
          description,
          duration,
          director_id,
          movie_star: selectedStars.map((star) => ({ star_id: star })),
          movie_genres: selectedGenres.map((genre) => ({
            genre_id: genre,
          })),
          featured_image: featuredImageurl,
          movie_images: otherImageURL.map((url) => ({
            other_images: url,
          })),
        },
        refetchQueries: [{ query: GET_MOVIES }],
      });
      if (response && response.data) {
        toast.success("Movie added successfully!", {
          position: "top-center",
          timeout: 5000, // Toast stays visible for 5 seconds
        });
      }
    } catch (error) {
      toast.error("Error while adding movies ");
    } finally {
      loading.value = false;
      isSpinning.value = false;
    }
  };

  const autoResize = (event) => {
    const textarea = event.target;
    textarea.style.height = "auto"; // Reset height to let it shrink if text
    textarea.style.height = `${textarea.scrollHeight}px`;
  };
  const addStarToArray = () => {
    const starId = selectedStar.value;

    if (starId && !selectedStars.value.includes(starId)) {
      selectedStars.value.push(starId);

      formData.value.star_id = [...selectedStars.value];
    }
  };

  const removeStarFromArray = (id) => {
    selectedStars.value = selectedStars.value.filter((starId) => starId !== id);
  };

  const getStarNameById = (id) => {
    const star = stars.value.find((star) => star.star_id === id);
    return star ? star.name : "Unknown Star";
  };

  const addGenreToArray = () => {
    const genreId = selectedGenre.value;
    if (genreId && !selectedGenres.value.includes(genreId)) {
      selectedGenres.value.push(genreId);
      formData.value.selected_genre = [...selectedGenres.value];
    }
  };

  const removeGenreFromArray = (id) => {
    selectedGenres.value = selectedGenres.value.filter(
      (genreId) => genreId !== id
    );
  };

  const getGenreNameById = (id) => {
    const genre = genres.value.find((genre) => genre.genre_id === id);
    return genre ? genre.name : "Unknown Genre";
  };

  return {
    addMovies,
    addGenreToArray,
    removeGenreFromArray,
    getGenreNameById,
    getStarNameById,
    removeStarFromArray,
    selectedStars,
    loading,
    selectedStar,
    selectedGenre,
    autoResize,
    isSpinning,
    selectedGenres,
    addStarToArray,
  };
};
