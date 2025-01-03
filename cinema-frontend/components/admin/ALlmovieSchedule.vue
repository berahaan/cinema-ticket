<script setup>
import { ref, onMounted } from "vue";
import Loading from "./Loading.vue";
const { formatScheduleDateTime } = useFormatSchedule();
const movie = useMoviesStore();
const {
  fetchMovies,
  searchQuery,
  goToNextPage,
  isloading,
  goToPreviousPage,
  currentPage,
} = useFetchMovies();

const { selectClass } = useThemeColor();
const { UpdateSchedules } = useMoviesSchedule();
const errorMessage = ref("");
onMounted(async () => {
  await fetchMovies();
});
</script>

<template>
  <div class="max-w-6xl mx-auto mt-10 px-4">
    <div class="mb-10 flex justify-center">
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search by title..."
        class="border p-2 rounded w-full md:w-1/3 text-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition mb-4 mt-20"
        :class="selectClass"
      />
    </div>
    <div
      v-if="isloading"
      class="flex justify-center items-center min-h-screen bg-white"
    >
      <Loading />
    </div>

    <template v-else>
      <div v-if="errorMessage" class="text-red-600 text-center mb-6">
        {{ errorMessage }}
      </div>
      <div
        v-if="movie.movies.length > 0"
        class="grid grid-cols-1 md:grid-cols-2 gap-4"
      >
        <!-- Movie Card -->
        <div
          v-for="movie in movie.movies"
          :key="movie.movie_id"
          class="rounded-lg shadow-lg overflow-hidden transform"
        >
          <!--  -->
          <div class="p-6">
            <h3 class="text-xl mb-4">
              {{ movie.title }}
            </h3>
            <div class="flex">
              <img
                v-if="movie.featured_images"
                :src="movie.featured_images"
                alt="Featured Image"
                class="w-full object-cover max-h-72 rounded-lg shadow-md cursor-pointer"
              />
            </div>

            <div
              class="p-2 bg-gray-100 rounded-lg shadow-sm"
              :class="selectClass"
            >
              <h4
                class="flex items-center font-semibold text-blue-800 text-lg mb-2"
              >
                📅 Scheduled Time
              </h4>
              <ul class="space-y-2">
                <li
                  v-for="schedule in movie.movie_schedules"
                  :key="schedule.start_time"
                  class="text-gray-700 rounded-lg p-3 shadow-md"
                  :class="selectClass"
                >
                  <p class="flex items-center">
                    <span
                      class="font-semibold text-sm mr-2"
                      :class="selectClass"
                    >
                      {{
                        formatScheduleDateTime(
                          schedule.schedule_date,
                          schedule.start_time,
                          schedule.end_time
                        )
                      }}
                      <span class="text-sm">
                        | Price(birr):
                        <span class="font-semibold text-green-600">{{
                          schedule.ticket_price
                        }}</span>
                      </span>
                      <span class="text-sm">
                        Cinema Hall:
                        <span class="font-semibold">{{
                          schedule.cinema_hall
                        }}</span>
                      </span>
                    </span>
                  </p>
                </li>
              </ul>
            </div>
            <button
              @click="UpdateSchedules(movie.movie_id)"
              class="px-4 mt-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg flex items-center justify-center"
            >
              Update Schedule
            </button>
          </div>
        </div>
      </div>

      <div
        v-else
        class="flex items-center justify-center p-6 bg-gray-100 border border-gray-300 rounded-md shadow-md md:mx-20 mb-40"
      >
        <p class="font-semibold text-gray-600 text-center">
          Unfortunately, we didn't find any movies matching your search.
        </p>
      </div>
    </template>

    <div
      class="pagination-controls py-4 flex justify-center items-center space-x-4 rounded-lg mt-8"
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
      <span>page {{ currentPage }} of {{ movie.totalPages }}</span>
      <button
        @click="goToNextPage"
        :disabled="currentPage === movie.totalPages"
        class="text-teal-600 px-4 py-2 rounded-md shadow-md disabled:opacity-50 disabled:cursor-not-allowed transition-all"
        :class="selectClass"
      >
        Next
      </button>
    </div>
  </div>
</template>
