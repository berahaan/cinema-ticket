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
const { UpdateMovie } = useMoviesSchedule();
const { refreshPage } = useRefresh();
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
        class="border p-2 rounded w-full md:w-1/3 text-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition mb-4 mt-12"
        :class="selectClass"
      />

      <button @click="refreshPage(fetchMovies)" class="ml-4 mb-8 mt-16">
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
        class="grid grid-cols-1 md:grid-cols-2 gap-8"
      >
        <!-- Movie Card -->
        <div
          v-for="movie in movie.movies"
          :key="movie.movie_id"
          class="rounded-lg shadow-lg overflow-hidden transform"
        >
          <!--  -->
          <div class="p-6">
            <h3 class="text-xl mb-4 font-bold" :class="selectClass">
              {{ movie.title }}
            </h3>
            <div class="flex">
              <img
                v-if="movie.featured_images"
                :src="movie.featured_images"
                alt="Featured Image"
                class="w-full object-cover mb-4 max-h-72 rounded-lg shadow-md cursor-pointer"
              />
            </div>
            <div class="text-gray-600">
              <h4
                class="flex items-center font-semibold text-blue-800 text-lg mb-2"
                :class="selectClass"
              >
                Genres:
              </h4>
              <ul class="flex">
                <li
                  v-for="genre in movie.movie_genres"
                  :key="genre.genre.name"
                  class="rounded-lg p-3 shadow-md"
                  :class="selectClass"
                >
                  {{ genre.genre.name }}
                </li>
              </ul>
              <span
                class="flex items-center font-semibold text-blue-800 text-lg mb-2 mt-2"
                :class="selectClass"
                >Stars:</span
              >
              <ul class="flex gap-2 mt-1">
                <li
                  v-for="star in movie.movie_stars"
                  :key="star.star.name"
                  class="rounded-lg p-3 shadow-md"
                  :class="selectClass"
                >
                  {{ star.star.name }}
                </li>
              </ul>
            </div>

            <div
              class="mt-4 p-4 bg-gray-100 rounded-lg shadow-sm"
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
                      <span class="text-sm text-gray-500">
                        | Price(birr):
                        <span class="font-semibold text-green-600">{{
                          schedule.ticket_price
                        }}</span>
                      </span>
                      <span class="text-sm text-gray-500">
                        Hall:
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
              @click="UpdateMovie(movie.movie_id)"
               class="px-4 mt-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-6"
            >
            
              Update Schedule
            </button>
          </div>
        </div>
      </div>

      <div
        v-else
        class="flex items-center justify-center p-6 bg-gray-100 border border-gray-300 rounded-md shadow-md md:mx-20"
      >
        <p class="font-semibold text-gray-600 text-center">
          Unfortunately, we didn't find any movies matching your search.
        </p>
      </div>
    </template>
 
       <div
          class="pagination-controls  py-4 flex justify-center items-center space-x-4 rounded-lg mt-8"
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
  </div>
</template>
