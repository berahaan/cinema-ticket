import { ref, computed } from "vue";
import { useToast } from "vue-toastification";
import { useNuxtApp } from "#app";
import { GET_MOVIES } from "../components/graphql/queries/GET_MOVIES.graphql";
import { GET_MOVIE_DETAILS } from "../components/graphql/queries/GET_MOVIE_DETAILS.graphql";
import { DELETE_MOVIE } from "../components/graphql/mutations/DELETE_MOVIE.graphql";
import { useRouter } from "vue-router";
export function useMovies() {
  const toast = useToast();
  const movieToDelete = ref(null);
  const isDeleteModalOpen = ref(false);
  const movie = useMoviesStore();
  const otherImages = ref([]);
  const movies = ref([]);
  const previewOtherImages = ref([]);
  const scheduleOptions = ref(["All", "Today", "This Week", "This Month"]);
  const currentOtherImage = ref([]);
  const successMessage = ref("");
  const currentImage = ref("");
  const isloading = ref(false);
  const loading = ref(true);
  const availableGenres = ref([]);
  const router = useRouter();
  const availableStars = ref([]);
  const availableDirectors = ref([]);
  const { $apollo } = useNuxtApp();
  const currentPage = ref(1);
  const limit = 3;
  const totalPages = ref(1);
  const offset = computed(() => (currentPage.value - 1) * limit);
  const formData = ref({
    title: "",
    description: "",
    duration: 0,
    featured_images: "",
    other_images: [],
    genres: [],
    stars: [],
    director_id: null,
  });

  const fetchMovieDetails = async (movie_id) => {
    try {
      const response = await $apollo.query({
        query: GET_MOVIE_DETAILS,
        variables: { movie_id },
      });

      movie.setMovies(response.data.movies[0]);
      formData.value = {
        title: movie.movies.title,
        description: movie.movies.description,
        duration: movie.movies.duration,
        genres: movie.movies.movie_genres.map((g) => g.genre.genre_id),
        stars: movie.movies.movie_stars.map((s) => s.star.star_id),
        director_id: movie.movies.director_id,
      };

      availableGenres.value = response.data.genres;
      availableStars.value = response.data.stars;
      availableDirectors.value = response.data.directors;
      currentImage.value = movie.movies.featured_images;
      otherImages.value = movie.movies.movie_images.map(
        (other) => other.other_images
      );
      loading.value = false;
      currentOtherImage.value = movie.movies.movie_images.map((other) =>
        other.other_images.trim()
      );
    } catch (error) {
      console.error("Error fetching movie details:", error);
    }
  };

  const UpdateMovie = (id) => {
    router.push(`/admin/movielist/${id}`);
  };

  const confirmDelete = (movieId) => {
    movieToDelete.value = movieId;
    isDeleteModalOpen.value = true;
  };
  const cancelDelete = () => {
    movieToDelete.value = null;
    isDeleteModalOpen.value = false;
  };
  const deleteMovie = async () => {
    try {
      await $apollo.mutate({
        mutation: DELETE_MOVIE,
        variables: { movieId: movieToDelete.value },
        refetchQueries: [
          {
            query: GET_MOVIES,
            variables: { limit: limit, offset: offset.value },
          },
        ],
      });
      movies.value = movies.value.filter(
        (movie) => movie.movie_id !== movieToDelete.value
      );
      movie.setMovies(
        movie.movies.filter((movie) => movie.movie_id !== movieToDelete.value)
      );
      toast.success("Movie deleted successfully ", {
        position: "top-center",
        timeout: 5000,
      });
    } catch (error) {
      console.error("Error deleting movie:", error);
    } finally {
      cancelDelete();
    }
  };
  return {
    movies,
    isloading,
    isDeleteModalOpen,
    confirmDelete,
    cancelDelete,
    totalPages,
    fetchMovieDetails,
    deleteMovie,
    formData,
    UpdateMovie,
    otherImages,
    currentImage,
    availableDirectors,
    availableGenres,
    availableStars,
    successMessage,
    previewOtherImages,
    offset,
    currentOtherImage,
    scheduleOptions,
    limit,
  };
}
