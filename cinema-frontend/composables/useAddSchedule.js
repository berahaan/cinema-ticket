import { ref } from "vue";
import { useToast } from "vue-toastification";
import { ADD_SCHEDULE } from "../components/graphql/mutations/ADD_SCHEDULE.graphql";
import { useNuxtApp } from "#app";
import { useRouter } from "vue-router";
export const useAddSchedule = () => {
  const success = ref(false);
  const { $apollo } = useNuxtApp();
  const router = useRouter();
  const { movies, fetchMovies } = useMoviesSchedule();
  const toast = useToast();
  const selectedMovieId = ref(null);
  const isScheduleModalOpen = ref(false);
  const isloading = ref(false);
  const test = ref(false);
  const schedule = ref({
    start_time: "",
    schedule_date: "",
    end_time: "",
    cinema_hall: "",
    ticket_price: null,
  });
  const openScheduleModal = (movie) => {
    selectedMovieId.value = movie.movie_id;
    router.push(`/admin/Addsch/${movie.movie_id}`);
    isScheduleModalOpen.value = false;
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
