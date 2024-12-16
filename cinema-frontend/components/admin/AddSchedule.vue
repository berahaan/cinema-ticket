<script setup>
import { useRoute } from "vue-router";
import { ArrowLeftIcon } from "@heroicons/vue/solid";
const route = useRoute();
const movieId = parseInt(route.params.id);
const selectedMovieId = ref(null);
const { schedule, addSchedule, isloading } = useAddSchedule();
const { selectClass } = useThemeColor();
const { goBack } = useGobackArrow();
const handleSchedule = async () => {
  
  selectedMovieId.value = movieId;
  
  try {
    await addSchedule(
      selectedMovieId.value,
      schedule.value.schedule_date,
      schedule.value.start_time,
      schedule.value.end_time,
      schedule.value.cinema_hall,
      schedule.value.ticket_price
    );
  } catch (error) {
   throw error
  }
};
</script>
<template>
  <div class="flex justify-center items-center py-10 mt-6" :class="selectClass">
    <div class="p-6 rounded-lg max-w-md w-full shadow-lg">
      <h3 class="text-xl font-bold mb-4 text-center">Schedule a Movie</h3>

      <!-- Schedule Date -->
      <div class="mb-4">
        <label class="block font-medium mb-2"> Schedule Date </label>
        <input
          type="date"
          v-model="schedule.schedule_date"
          :class="selectClass"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
        />
      </div>

      <!-- Start Time -->
      <div class="mb-4">
        <label class="block font-medium mb-2"> Start Time </label>
        <input
          type="time"
          v-model="schedule.start_time"
          :class="selectClass"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
        />
      </div>

      <!-- End Time -->
      <div class="mb-4">
        <label class="block font-medium mb-2"> End Time </label>
        <input
          type="time"
          v-model="schedule.end_time"
          :class="selectClass"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
        />
      </div>

      <!-- Hall -->
      <div class="mb-4">
        <label class="block font-medium mb-2"> Cinema Hall </label>
        <input
          type="text"
          :class="selectClass"
          v-model="schedule.cinema_hall"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
          placeholder="e.g., Hall A"
        />
      </div>

      <!-- Price -->
      <div class="mb-6">
        <label class="block font-medium mb-2"> Price (Birr) </label>
        <input
          type="number"
          v-model="schedule.ticket_price"
          :class="selectClass"
          class="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:outline-none"
          placeholder="e.g., 150"
        />
      </div>

      <!-- Button -->
      <button
        @click="handleSchedule(movieId)"
        class="w-full bg-blue-600 text-white font-medium px-4 py-2 rounded-lg transition-all duration-200 transform hover:scale-105 hover:bg-blue-700 focus:outline-none disabled:opacity-50"
        :disabled="isloading"
      >
        <span v-if="isloading" class="flex items-center justify-center">
          <svg
            class="animate-spin h-5 w-5 mr-2 text-white"
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
          Scheduling...
        </span>
        <span v-else>Schedule</span>
      </button>
      <div class="fixed top-20 left-0 z-20">
        <button
          @click="goBack"
          class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 dark:text-blue-400 dark:hover:text-blue-500 font-medium py-2 px-4 rounded-lg transition-all duration-200 transform hover:scale-105 mb-6"
        >
          <ArrowLeftIcon class="h-5 w-5" />
          <span>Go Back</span>
        </button>
      </div>
    </div>
  </div>
</template>
