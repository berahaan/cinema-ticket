<script setup>
import { ArrowLeftIcon } from "@heroicons/vue/solid";
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
const { goBack } = useGobackArrow();
const movie = useMoviesStore();
const { fetchMovieDetails } = useMovies();
const route = useRoute();
const movieId = parseInt(route.params.id, 10);
const loading = ref(false);
const error = ref(null);
onMounted(() => {
  fetchMovieDetails(movieId);
});
</script>

<template>
  <div class="min-h-screen w-full mt-16">
    <div v-if="loading" class="text-center py-10">Loading movie details...</div>
    <div v-else-if="error" class="text-center text-red-600 py-10">
      Error: {{ error.message }}
    </div>
    <div v-else class="max-w-4xl mx-auto p-6 shadow-lg rounded-lg">
      <!-- Movie Title and Poster -->
      <button
        @click="goBack"
        class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200"
      >
        <ArrowLeftIcon class="h-5 w-5" />
        <span>Go Back</span>
      </button>
      <div class="flex flex-col md:flex-row gap-6">
        <img
          :src="movie.movies.featured_images"
          alt="Movie Poster"
          class="w-full md:w-1/3 h-auto rounded-lg shadow-lg object-cover"
        />
        <div class="flex-1">
          <h1 class="text-3xl font-bold">{{ movie.movies.title }}</h1>
          <p class="mt-4">{{ movie.movies.description }}</p>
          <p class="mt-4">
            <strong>Director:</strong>
            {{ movie.movies.director?.name || "Unknown" }}
          </p>
          <p class="mt-4">
            <strong>Duration:</strong> {{ movie.movies.duration }} minutes
          </p>
          <p class="mt-4">
            <strong>Average Rating:</strong>
            {{ movie.movies.average_rating?.average_rating || "Not Rated" }}
          </p>
        </div>
      </div>

      <!-- Genres -->
      <div v-if="movie.movies.movie_genres?.length > 0" class="mt-8">
        <h2 class="text-xl">Genres</h2>
        <div class="flex gap-2 mt-2">
          <span
            v-for="genre in movie.movies.movie_genres"
            :key="genre.genre.name"
            class="px-3 py-1 rounded-full text-sm"
          >
            {{ genre.genre.name }}
          </span>
        </div>
      </div>

      <!-- Stars -->
      <div v-if="movie.movies.movie_stars?.length > 0" class="mt-8">
        <h2 class="text-xl">Stars</h2>
        <ul class="list-disc list-inside">
          <li v-for="star in movie.movies.movie_stars" :key="star.star.name">
            {{ star.star.name }}
          </li>
        </ul>
      </div>

      <!-- Schedules -->
      <div v-if="movie.movies.movie_schedules?.length > 0" class="mt-8">
        <h2 class="text-xl font-semibol">Schedules</h2>
        <table class="w-full border-collapse border border-gray-300 mt-4">
          <thead>
            <tr class="text-left">
              <th class="border border-gray-300 px-4 py-2">Date</th>
              <th class="border border-gray-300 px-4 py-2">Start Time</th>
              <th class="border border-gray-300 px-4 py-2">End Time</th>
              <th class="border border-gray-300 px-4 py-2">Cinema Hall</th>
              <th class="border border-gray-300 px-4 py-2">Ticket Price</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="schedule in movie.movies.movie_schedules"
              :key="schedule.start_time"
            >
              <td class="border border-gray-300 px-4 py-2">
                {{ schedule.schedule_date }}
              </td>
              <td class="border border-gray-300 px-4 py-2">
                {{ schedule.start_time }}
              </td>
              <td class="border border-gray-300 px-4 py-2">
                {{ schedule.end_time }}
              </td>
              <td class="border border-gray-300 px-4 py-2">
                {{ schedule.cinema_hall }}
              </td>
              <td class="border border-gray-300 px-4 py-2">
                {{ schedule.ticket_price.toFixed(2) }} birr
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
