<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useNuxtApp } from "#app";
import { ArrowLeftIcon } from "@heroicons/vue/solid";
import { UPDATE_SCHEDULE } from "../components/graphql/mutations/UPDATE_SCHEDULE.graphql";
import { DELETE_SCHEDULE } from "../components/graphql/mutations/DELETE_SCHEDULE.graphql";
import { FETCH_SCHEDULES } from "../components/graphql/queries/FETCH_SCHEDULES.graphql";
import { GET_MOVIES } from "../components/graphql/queries/GET_MOVIES.graphql";
import { GET_SCHEDULE_INFO } from "../components/graphql/queries/GET_SCHEDULE_INFO.graphql";
import { useToast } from "vue-toastification";
const { $apollo } = useNuxtApp();
const route = useRoute();
const router = useRouter();
const toast = useToast();
const movie_id = parseInt(route.params.id);
const { offset, limit } = useFetchMovies();
const { selectClass } = useThemeColor();
const schedules = ref([]);
const loadingStates = ref([]);
const { goBack } = useGobackArrow();
const fetchMovieSchedules = async () => {
  const { data } = await $apollo.query({
    query: FETCH_SCHEDULES,
    variables: { movie_id },
    fetchPolicy: "network-only",
  });
  console.log("Fetched schedules now ", data.movie_schedules);
  schedules.value = data.movie_schedules.map((schedule) => ({
    schedule_id: schedule.schedule_id, // Ensure each schedule has a unique ID
    start_time: extractTime(schedule.start_time),
    end_time: extractTime(schedule.end_time),
    schedule_date: schedule.schedule_date,
    ticket_price: schedule.ticket_price,
    cinema_hall: schedule.cinema_hall,
  }));
  loadingStates.value = Array(schedules.value.length).fill(false);
};
const removeSchedule = async (id) => {
  console.log("To be deleted schedules Id ", id);
  try {
    // first check if the schedules is already removed and warn them as its removed
    const { data } = await $apollo.query({
      query: GET_SCHEDULE_INFO,
      variables: { schedule_id: id },
      fetchPolicy: "network-only",
    });
    console.log("checkSchedulesDelete", data);
    if (data.movie_schedules.length > 0) {
      console.log("Need to be deleted now....");
      const deleteSchedules = await $apollo.mutate({
        mutation: DELETE_SCHEDULE,
        variables: { schedule_id: id },
        refetchQueries: [
          {
            query: GET_MOVIES,
            variables: { limit: limit, offset: offset.value },
          },
        ],
      });
      if (deleteSchedules && deleteSchedules.data) {
        await fetchMovieSchedules();
        toast.success("Schedules deleted successfully ", {
          position: "top-right",
          timeout: 2000,
        });
      }
    } else {
      console.log("ALready removed ");
      toast.error("Schedules ALready removed try other one", {
        position: "top-right",
        timeout: 2000,
      });
    }
  } catch (error) {
    console.log("Errors ", error);
  }
};

const extractTime = (dateTime) => {
  return dateTime.substring(11, 16);
};
const updateSchedule = async (schedule, index) => {
  const {
    schedule_id,
    start_time,
    end_time,
    schedule_date,
    ticket_price,
    cinema_hall,
  } = schedule;
  const startDateTime = `${schedule_date} ${start_time}`;
  const endDateTime = `${schedule_date} ${end_time}`;

  loadingStates.value[index] = true; // Track loading state for this specific schedule

  try {
    const response = await $apollo.mutate({
      mutation: UPDATE_SCHEDULE,
      variables: {
        schedule_id: schedule_id, // Use unique ID to update only this schedule
        start_time: startDateTime,
        end_time: endDateTime,
        schedule_date,
        ticket_price: parseFloat(ticket_price),
        cinema_hall,
      },
      refetchQueries: [
        {
          query: GET_MOVIES,
          variables: { limit: limit, offset: offset.value },
        },
      ],
    });
    if (response && response.data) {
      await fetchMovieSchedules();
      toast.success(`Schedule ${index + 1} updated successfully!`, {
        position: "top-center",
        timeout: 4000,
      });
    }
  } catch (error) {
    console.error("Error updating schedule:", error);
    alert("An error occurred while updating the schedule.");
  } finally {
    loadingStates.value[index] = false;
  }
};

const goToAddSchedulePage = () => {
  router.push("/admin/Schedules");
};

onMounted(async () => {
  await fetchMovieSchedules();
});
</script>

<template>
  <div class="container mx-auto p-4 mt-20">
    <button
      @click="goBack"
      class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200"
    >
      <ArrowLeftIcon class="h-5 w-5" />
      <span>Go Back</span>
    </button>
    <div v-if="schedules.length">
      <div
        v-for="(schedule, index) in schedules"
        :key="index"
        :class="selectClass"
        class="mb-6 p-6 border border-gray-300 rounded-lg shadow-md transition-shadow hover:shadow-lg"
      >
        <h3 class="font-bold text-lg mb-2">Schedule {{ index + 1 }}</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label
              for="start_time"
              :class="selectClass"
              class="block text-sm font-medium"
              >Start Time</label
            >
            <input
              id="start_time"
              type="time"
              v-model="schedule.start_time"
              placeholder="HH:MM"
              :class="selectClass"
              class="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
          </div>
          <div>
            <label
              for="end_time"
              class="block text-sm font-medium"
              :class="selectClass"
              >End Time</label
            >
            <input
              id="end_time"
              type="time"
              v-model="schedule.end_time"
              placeholder="HH:MM"
              :class="selectClass"
              class="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
          </div>

          <!-- Scheduled Date -->
          <div>
            <label
              for="schedule_date"
              :class="selectClass"
              class="block text-sm font-medium"
              >Scheduled Date</label
            >
            <input
              id="schedule_date"
              type="date"
              :class="selectClass"
              v-model="schedule.schedule_date"
              class="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
          </div>

          <!-- Ticket Price -->
          <div>
            <label
              for="ticket_price"
              class="block text-sm font-medium"
              :class="selectClass"
              >Ticket Price</label
            >
            <input
              id="ticket_price"
              type="number"
              :class="selectClass"
              v-model="schedule.ticket_price"
              class="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
          </div>

          <!-- Cinema Hall -->
          <div>
            <label
              for="cinema_hall"
              class="block text-sm font-medium"
              :class="selectClass"
              >Cinema Hall</label
            >
            <input
              id="cinema_hall"
              type="text"
              :class="selectClass"
              v-model="schedule.cinema_hall"
              class="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-600"
            />
          </div>
        </div>

        <button
          @click="updateSchedule(schedule, index)"
          class="px-4 mt-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-6"
        >
          <span v-if="loadingStates[index]">
            <svg
              class="animate-spin h-5 w-5 mr-2 text-white inline-block"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <circle
                class="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                stroke-width="4"
              ></circle>
              <path
                class="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
              ></path>
            </svg>
            Updating...
          </span>
          <span v-else>Update Schedule</span>
        </button>
        <button
          @click="removeSchedule(schedule.schedule_id)"
          class="px-4 mt-4 py-2 bg-red-600 dark:bg-red-800 text-white rounded-lg mx-6"
        >
          Remove Schedule
        </button>
      </div>
    </div>

    <!-- Notification for No Schedules -->
    <div
      v-else
      class="mt-4 p-4 bg-yellow-100 text-yellow-700 border border-yellow-500 rounded-lg"
    >
      <p>This movie is not yet scheduled. Schedule it here:</p>
      <button
        @click="goToAddSchedulePage"
        class="px-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-6"
      >
        Add Schedule
      </button>
    </div>
  </div>
</template>
