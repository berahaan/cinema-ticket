import { ref, computed } from "vue";
import { jwtDecode } from "jwt-decode";
import { useNuxtApp } from "#app";
import { useRoute } from "vue-router";
import { SEND_TICKET } from "../components/graphql/mutations/SEND_TICKET.graphql";
import { GET_SPECIFIC_MOVIE } from "../components/graphql/queries/GET_SPECIFIC_MOVIE.graphql";
import { GET_USER } from "../components/graphql/queries/GET_USER.graphql";
import { useToast } from "vue-toastification";
export const useTicket = () => {
  const route = useRoute();
  const router = useRouter();
  const toast = useToast();
  const movieId = parseInt(route.params.id, 10);
  const scheduleId = ref(null);
  const movie = ref({});
  const isloading = ref(false);
  const selectedSchedule = ref(null);
  const ticketQuantity = ref(1);
  const islengthGreater = ref(false);
  const { $apollo } = useNuxtApp();
  const form = ref({
    first_name: "",
    last_name: "",
    phone_number: "",
    email: "",
    currency: "ETB",
  });
  const totalPrice = computed(() => {
    scheduleId.value = selectedSchedule.value?.schedule_id;
    return selectedSchedule.value
      ? selectedSchedule.value.ticket_price * ticketQuantity.value
      : 0;
  });
  watch(scheduleId, (newId) => {
    console.log("Selected Schedule ID:", newId); // Access the value
  });

  const fetchInfo = async () => {
    try {
      const response = await $apollo.query({
        query: GET_SPECIFIC_MOVIE,
        variables: { movieId: movieId },
      });
      movie.value = response.data.movies[0];
      islengthGreater.value = movie.value.movie_schedules.length > 0;
    } catch (error) {
      console.log("Error fetching movie data: ", error);
    }
  };
  const fetchUserName = async () => {
    const token = jwtDecode(localStorage.getItem("accessToken"));
    const userId = token["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
    try {
      const { data } = await $apollo.query({
        query: GET_USER,
        variables: { userid: userId },
      });

      form.value.first_name = data.users[0].firstname;
      form.value.last_name = data.users[0].lastname;
      form.value.email = token.email;
    } catch (error) {
      console.log("Errors ", error);
    }
  };
  const buyTicket = (id) => {
    const token = jwtDecode(localStorage.getItem("accessToken"));
    const defaultRole =
      token["https://hasura.io/jwt/claims"]["x-hasura-default-role"];

    if (defaultRole === "admin") {
      router.push(`/admin/ticket/${id}`);
    } else {
      router.push(`/user/ticket/${id}`);
    }
  };
  const purchaseTickets = async (form) => {
    isloading.value = true;

    const Input = {
      TicketQuantity: ticketQuantity.value,
      movieId: movieId,
      scheduleId: scheduleId.value,
      amount: totalPrice.value,
      first_name: form.value.first_name,
      last_name: form.value.last_name,
      phone_number: form.value.phone_number,
      email: form.value.email,
      currency: form.value.currency,
    };

    try {
      const { data } = await $apollo.mutate({
        mutation: SEND_TICKET,
        variables: Input,
      });
      const accessUrl = data.initiatePayment.access_url;
      window.location.href = accessUrl;
    } catch (error) {
      if (!navigator.onLine) {
        toast.warning(
          "It seems you're offline. Please check your internet connection and try again.",
          {
            position: "top-right",
            timeout: 1000,
          }
        );
      } else {
        toast.warning(
          "look like offline please connect to internet and try again",
          {
            position: "top-right",
            duration: 2000,
          }
        );
      }
    } finally {
      isloading.value = false;
    }
  };

  return {
    movie,
    totalPrice,
    fetchInfo,
    selectedSchedule,
    islengthGreater,
    form,
    buyTicket,
    isloading,
    movieId,
    fetchUserName,
    purchaseTickets,
    ticketQuantity,
  };
};
