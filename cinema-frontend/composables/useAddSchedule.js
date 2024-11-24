import { ref } from "vue";
import { useToast } from "vue-toastification";
import { ADD_SCHEDULE } from "../components/graphql/mutations/ADD_SCHEDULE.graphql";
import { GET_MOVIES } from "../components/graphql/queries/GET_MOVIES.graphql";
import { useNuxtApp } from "#app";

export const useAddSchedule = () => {
  const success = ref(false);
  const { $apollo } = useNuxtApp();

  const { movies, fetchMovies } = useMoviesSchedule();
  const toast = useToast();
  const { offset, limit, testAddmovies } = useMovies();
  const selectedMovieId = ref(null);
  const isScheduleModalOpen = ref(false);
  const isloading = ref(false);
  const loading = ref(false);
  const test = ref(false);
  const schedule = ref({
    start_time: "",
    schedule_date: "",
    end_time: "",
    cinema_hall: "",
    ticket_price: null,
  });

  const openScheduleModal = (movie) => {
    selectedMovieId.value = movie.movie_id; // S
    isScheduleModalOpen.value = true;
  };

  const closeScheduleModal = () => {
    isScheduleModalOpen.value = false;
    isloading.value = false;
    test.value = true;
  };

  const resetScheduleForm = () => {
    schedule.value = {
      start_time: "",
      schedule_date: "",
      end_time: "",
      cinema_hall: "",
      ticket_price: null,
    };
    selectedMovieId.value = null;
  };

  const addSchedule = async (
    movieId,
    scheduleDate,
    startTime,
    endTime,
    cinemahall,
    ticket_price
  ) => {
    isloading.value = true;
    const startDate = `${scheduleDate}T${startTime}:00`;
    const endDate = `${scheduleDate}T${endTime}:00`;

    try {
      const variables = {
        movie_id: movieId,
        start_time: startDate,
        end_time: endDate,
        schedule_date: scheduleDate,
        cinema_hall: cinemahall,
        ticket_price: parseFloat(ticket_price),
      };
      const response = await $apollo.mutate({
        mutation: ADD_SCHEDULE,
        variables: variables,
        fetchPolicy: "network-only",
      });

      if (response && response.data) {
        // testAddmovies.value = "Added successfully ";
        toast.success("Schedule added successfully!", {
          position: "top-center",
          timeout: 4000,
        });
        success.value = true;
        resetScheduleForm();
      }
    } catch (error) {
      console.error("Error adding schedule:", error);
    } finally {
      isloading.value = false;
    }
  };
  watch([isScheduleModalOpen], (newVal) => {
    if (newVal) {
      fetchMovies();
    }
  });

  return {
    addSchedule,
    openScheduleModal,
    schedule,
    success,
    test,
    movies,
    isloading,
    closeScheduleModal,
    selectedMovieId,
    isScheduleModalOpen,
    resetScheduleForm,
  };
};
