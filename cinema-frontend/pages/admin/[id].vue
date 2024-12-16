<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { useNuxtApp } from "#app";
import gql from "graphql-tag";

// Apollo instance
const { $apollo } = useNuxtApp();

// Query to fetch movie details
const GET_MOVIE_DETAILS = gql`
  query GetMovieDetails($movieId: Int!) {
    movies(where: { movie_id: { _eq: $movieId } }) {
      movie_id
      title
      description
      duration
      featured_images
      average_rating {
        average_rating
      }
      movie_genres {
        genre {
          name
        }
      }
      movie_stars {
        star {
          name
        }
      }
      director {
        name
      }
      movie_schedules {
        start_time
        end_time
        schedule_date
        cinema_hall
        ticket_price
      }
    }
  }
`;

// Reactive state
const route = useRoute();
const movieId = parseInt(route.params.id, 10);
const movie = ref({});
const loading = ref(false);
const error = ref(null);

// Fetch movie details
async function fetchMovieDetails(movieId) {
  loading.value = true;
  try {
    const { data } = await $apollo.query({
      query: GET_MOVIE_DETAILS,
      variables: { movieId },
    });
    movie.value = data.movies[0];
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }
}

// Fetch details when component is mounted
onMounted(() => {
  fetchMovieDetails(movieId);
});
</script>
<template>
  <div class="min-h-screen bg-gray-100">
    <div v-if="loading" class="text-center text-gray-600 py-10">
      Loading movie details...
    </div>
    <div v-else-if="error" class="text-center text-red-600 py-10">
      Error: {{ error.message }}
    </div>
    <div v-else class="max-w-4xl mx-auto p-6 bg-white shadow-lg rounded-lg">
      <!-- Movie Title and Poster -->
      <div class="flex flex-col md:flex-row gap-6">
        <img
          :src="movie.featured_images || '/placeholder.jpg'"
          alt="Movie Poster"
          class="w-full md:w-1/3 h-auto rounded-lg shadow-lg object-cover"
        />
        <div class="flex-1">
          <h1 class="text-3xl font-bold text-gray-800">{{ movie.title }}</h1>
          <p class="mt-4 text-gray-600">{{ movie.description }}</p>
          <p class="mt-4 text-gray-600">
            <strong>Director:</strong> {{ movie.director?.name || "Unknown" }}
          </p>
          <p class="mt-4 text-gray-600">
            <strong>Duration:</strong> {{ movie.duration }} minutes
          </p>
          <p class="mt-4 text-gray-600">
            <strong>Average Rating:</strong>
            {{ movie.average_rating?.average_rating || "Not Rated" }}
          </p>
        </div>
      </div>

      <!-- Genres -->
      <div v-if="movie.movie_genres.length > 0" class="mt-8">
        <h2 class="text-xl font-semibold text-gray-800">Genres</h2>
        <div class="flex gap-2 mt-2">
          <span
            v-for="genre in movie.movie_genres"
            :key="genre.genre.name"
            class="px-3 py-1 bg-gray-200 text-gray-700 rounded-full text-sm"
          >
            {{ genre.genre.name }}
          </span>
        </div>
      </div>

      <!-- Stars -->
      <div v-if="movie.movie_stars.length > 0" class="mt-8">
        <h2 class="text-xl font-semibold text-gray-800">Stars</h2>
        <ul class="list-disc list-inside text-gray-600">
          <li v-for="star in movie.movie_stars" :key="star.star.name">
            {{ star.star.name }}
          </li>
        </ul>
      </div>

      <!-- Schedules -->
      <div v-if="movie.movie_schedules.length > 0" class="mt-8">
        <h2 class="text-xl font-semibold text-gray-800">Schedules</h2>
        <table class="w-full border-collapse border border-gray-300 mt-4">
          <thead>
            <tr class="bg-gray-200 text-left">
              <th class="border border-gray-300 px-4 py-2">Date</th>
              <th class="border border-gray-300 px-4 py-2">Start Time</th>
              <th class="border border-gray-300 px-4 py-2">End Time</th>
              <th class="border border-gray-300 px-4 py-2">Cinema Hall</th>
              <th class="border border-gray-300 px-4 py-2">Ticket Price</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="schedule in movie.movie_schedules"
              :key="schedule.start_time"
              class="odd:bg-white even:bg-gray-100"
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
                ${{ schedule.ticket_price.toFixed(2) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
