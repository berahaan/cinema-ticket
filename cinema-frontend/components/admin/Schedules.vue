<script setup>
import { onMounted } from "vue";
import Loading from "../admin/Loading.vue";
import { ArrowLeftIcon } from "@heroicons/vue/solid";
const { formatScheduleDateTime } = useFormatSchedule();
const { selectClass } = useThemeColor();
const { refreshPage } = useRefresh();
const { goBack } = useGobackArrow();
const movie = useMoviesStore();
const {
  addSchedule,
  schedule,
  openScheduleModal,
  closeScheduleModal,
  isScheduleModalOpen,
  isloading,
  selectedMovieId,
} = useAddSchedule();
const {
  fetchMovies,
  searchQuery,
  goToNextPage,
  goToPreviousPage,
  currentPage,
  totalPages,
  noMoviesFound,
} = useFetchMovies();

const handleSchedule = async () => {
  console.log("Now handle schedules is processing... ");
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
    console.log("Errors ", error);
  }
};

watch(currentPage, async () => {
  console.log("current pages is changing ");
  await fetchMovies();
});

onMounted(async () => {
  await fetchMovies();
});
</script>

<template>
  <div class="mt-6">
    <h2 class="text-2xl font-bold mb-6 text-center">Movies List</h2>

    <div
      v-if="isloading"
      class="flex justify-center items-center min-h-screen bg-white"
    >
      <Loading />
    </div>

    <div class="mb-10 flex justify-between items-center">
      <!-- Search Input (Center-aligned) -->
      <div class="flex-grow flex justify-center">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Search by title..."
          class="border p-3 rounded w-full md:w-1/2 text-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition mb-4"
          :class="selectClass"
        />
      </div>

      <!-- Container for Refresh and Go Back buttons (Right-aligned) -->
      <div class="flex items-center space-x-4 ml-4">
        <!-- Refresh Button (Right-aligned) -->
        <button @click="refreshPage(fetchMovies)" class="mb-4">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-6 h-6 text-blue-500 hover:text-blue-700 transition-colors"
            viewBox="0 0 24 24"
            fill="currentColor"
          >
            <path
              d="M17.65 6.35A7.95 7.95 0 0012 4V1L8 5l4 4V6c1.78 0 3.4.68 4.65 1.8a6 6 0 011.29 6.12l1.45 1.45c1.19-2.32 1.11-5.25-.6-7.57zm-11.3 11.3A7.95 7.95 0 0012 20v3l4-4-4-4v3c-1.78 0-3.4-.68-4.65-1.8a6 6 0 01-1.29-6.12l-1.45-1.45c-1.19 2.32-1.11 5.25.6 7.57z"
            />
          </svg>
        </button>

        <!-- Go Back Button (Right-aligned) -->
        <button
          @click="goBack"
          class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200 mb-4"
        >
          <ArrowLeftIcon class="h-5 w-5" />
          <span>Go Back</span>
        </button>
      </div>
    </div>
    <div
      v-if="noMoviesFound"
      class="flex justify-center items-center"
      :class="selectClass"
    >
      No movies match your search please browse our other search ...
    </div>
    <div
      v-for="movie in movie.movies"
      :key="movie.title"
      class="border rounded-lg shadow-lg p-6 transition-transform transform hover:shadow-xl"
      :class="selectClass"
    >
      <!--  -->
      <h3 class="text-4xl mb-4 font-bold" :class="selectClass">
        {{ movie.title }}
      </h3>
      <p class="text-gray-700 mb-2" :class="selectClass">
        {{ movie.description }}
      </p>
      <p class="font-medium" :class="selectClass">
        Duration: {{ movie.duration }} minutes
      </p>
      <div class="flex flex-col md:flex-row mt-4">
        <img
          v-if="movie.featured_images"
          :src="movie.featured_images"
          alt="Featured Image"
          class="w-full md:w-1/3 lg:w-1/3 max-h-96 object-cover mb-4 rounded-lg shadow-md cursor-pointer hover:scale-105 transition-transform duration-200"
          @click="openModal(movie.movie_images)"
        />
        <div class="md:w-1/2">
          <h4 class="font-semibold mx-2">Genres:</h4>
          <ul class="list-disc pl-5 mx-2">
            <li v-for="genre in movie.movie_genres" :key="genre.genre.name">
              {{ genre.genre.name }}
            </li>
          </ul>
          <h4 class="font-semibold mt-4 mx-2">Stars:</h4>
          <ul class="list-disc pl-5 mx-2">
            <li v-for="star in movie.movie_stars" :key="star.star.name">
              {{ star.star.name }}
            </li>
          </ul>
          <div
            class="mt-4 p-4 bg-gray-100 rounded-lg shadow-sm"
            :class="selectClass"
          >
            <h4
              class="flex items-center font-semibold text-blue-800 text-lg mb-2"
            >
              📅 Scheduled Time
            </h4>
            <ul class="space-y-1">
              <li
                v-for="schedule in movie.movie_schedules"
                :key="schedule.start_time"
                class="rounded-lg p-3 shadow-md"
                :class="selectClass"
              >
                <p class="flex items-center">
                  <span class="font-semibold mr-2" :class="selectClass">
                    {{
                      formatScheduleDateTime(
                        schedule.schedule_date,
                        schedule.start_time,
                        schedule.end_time
                      )
                    }}
                  </span>
                  <span class="text-sm">
                    | Price:
                    <span class="text-green-600"
                      >{{ schedule.ticket_price }} birr</span
                    >
                  </span>
                </p>
              </li>
            </ul>
          </div>
          <button
            class="px-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-6"
            @click="openScheduleModal(movie)"
          >
            Add Schedule
          </button>
          <div
            v-if="isScheduleModalOpen && selectedMovieId === movie.movie_id"
            class="absolute top-0 left-0 w-full h-auto bg-opacity-70 flex items-center justify-center z-50 transition-opacity"
          >
            <div
              class="p-6 rounded-lg max-w-lg w-full shadow-lg"
              :class="selectClass"
            >
              <h3 class="text-lg font-semibold mb-4">Schedule Movie</h3>
              <div class="mb-4">
                <label class="block font-semibold mb-2">Schedule Date</label>
                <input
                  type="date"
                  v-model="schedule.schedule_date"
                  :class="selectClass"
                  class="w-full p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:outline-none"
                />
              </div>
              <div class="mb-4">
                <label class="block font-semibold mb-2">Start Time</label>
                <input
                  type="time"
                  v-model="schedule.start_time"
                  :class="selectClass"
                  class="w-full p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:outline-none"
                />
              </div>
              <div class="mb-4">
                <label class="block font-semibold mb-2">End Time</label>
                <input
                  type="time"
                  v-model="schedule.end_time"
                  :class="selectClass"
                  class="w-full p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:outline-none"
                />
              </div>
              <div class="mb-4">
                <label class="block font-semibold mb-2">Hall</label>
                <input
                  type="text"
                  :class="selectClass"
                  v-model="schedule.cinema_hall"
                  class="w-full p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:outline-none"
                />
              </div>
              <div class="mb-4">
                <label class="block font-semibold mb-2">Price (Birr)</label>
                <input
                  type="number"
                  :class="selectClass"
                  v-model="schedule.ticket_price"
                  class="w-full p-2 border rounded focus:ring-2 focus:ring-blue-500 focus:outline-none"
                />
              </div>

              <div class="flex justify-between">
                <button
                  @click="closeScheduleModal"
                  class="bg-red-500 text-white px-4 py-2 rounded transition-all duration-200 transform hover:scale-105 focus:outline-none"
                >
                  Close
                </button>
                <button
                  @click="handleSchedule(movie.movie_id)"
                  class="bg-blue-600 text-white px-4 py-2 rounded transition-all duration-200 transform hover:scale-105 focus:outline-none"
                >
                  <span v-if="isloading">
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
                    Scheduling...
                  </span>
                  <span v-else>Schedule</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div
    class="pagination-controls bg-teal-600 py-4 flex justify-center items-center space-x-4 rounded-lg mt-8"
    v-if="movie.totalPages"
  >
    <button
      @click="goToPreviousPage"
      :disabled="currentPage === 1"
      class="text-teal-600 px-4 py-2 rounded-md shadow-md disabled:opacity-50 disabled:cursor-not-allowed transition-all"
      :class="selectClass"
    >
      Previous
    </button>
    <span class="font-semibold">
      {{
        noMoviesFound
          ? "Page 1 of total 1"
          : `page ${currentPage} of ${movie.totalPages}`
      }}
    </span>
    <button
      @click="goToNextPage"
      :disabled="currentPage === movie.totalPages || noMoviesFound"
      class="text-teal-600 px-4 py-2 rounded-md shadow-md disabled:opacity-50 disabled:cursor-not-allowed transition-all"
      :class="selectClass"
    >
      Next
    </button>
  </div>
</template>
