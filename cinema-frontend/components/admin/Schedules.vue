<script setup>
import { onMounted } from "vue";
import Loading from "../admin/Loading.vue";
import { ArrowLeftIcon, ArrowRightIcon } from "@heroicons/vue/solid";
const { formatScheduleDateTime } = useFormatSchedule();
const { selectClass } = useThemeColor();
const { refreshPage } = useRefresh();
const { goBack } = useGobackArrow();
const movie = useMoviesStore();
const { openScheduleModal, isloading } = useAddSchedule();
const {
  fetchMovies,
  searchQuery,
  goToNextPage,
  goToPreviousPage,
  currentPage,
  noMoviesFound,
} = useFetchMovies();

watch(currentPage, async () => {
  console.log("current pages is changing ");
  await fetchMovies();
});

onMounted(async () => {
  await fetchMovies();
});
</script>

<template>
  <div class="mt-16">
    <!-- <div class="mt-1"></div> -->
    <div
      v-if="isloading"
      class="flex justify-center items-center min-h-screen bg-white"
    >
      <Loading />
    </div>

    <div class="mb-10 flex justify-between items-center mt-32">
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
        </div>
      </div>
    </div>
  </div>

  <div
    class="pagination-controls py-4 flex justify-center items-center space-x-4 rounded-lg mt-8"
    v-if="movie.totalPages"
  >
    <button
      @click="goToPreviousPage"
      :disabled="currentPage === 1"
      class="text-teal-600 flex px-4 py-2 rounded-md shadow-md disabled:opacity-50 disabled:cursor-not-allowed transition-all"
    >
      <ArrowLeftIcon class="w-6 h-6 mx-2" />
      previous
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
      :disabled="currentPage === movie.totalPages"
      class="text-teal-600 flex px-4 py-2 rounded-md shadow-md disabled:opacity-50 disabled:cursor-not-allowed transition-all"
    >
      Next
      <ArrowRightIcon class="w-6 h-6 mx-2" />
    </button>
  </div>
</template>
