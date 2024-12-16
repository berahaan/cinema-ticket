<script setup>
import { ref, computed, onMounted } from "vue";
import { useNuxtApp } from "#app";
import { GET_MOVIE_ALL } from "./graphql/queries/GET_MOVIE_ALL.graphql";
import { useRouter } from "vue-router";
const router = useRouter();
const { $apollo } = useNuxtApp();
const movies = ref([]);
const fetchMovies = async () => {
  try {
    const { data } = await $apollo.query({
      query: GET_MOVIE_ALL,
    });

    movies.value = data.movies;
  } catch (error) {
    console.error("Error fetching movies:", error);
  }
};
const recentMovies = computed(() => movies.value.slice(0, 3));
const currentlyShowingMovies = computed(() => movies.value.slice(4, 6));
const buyTicket = (movieId) => {
  router.push(`/admin/ticket/${movieId}`);
};
onMounted(() => {
  fetchMovies();
});
</script>
<template>
  <div>
    <div class="min-h-screen mt-16">
      <section
        class="text-center py-16 mb-12 text-black rounded-lg shadow-2xl mt-10 w-full"
      >
        <h2 class="text-xl md:text-4xl font-extrabold mb-6 drop-shadow-lg">
          <!-- 🎬 -->
          Welcome to Ethio Cinema!
          <!-- 🌟 -->
        </h2>
        <p class="text-xl md:text-xl font-light">
          Immerse yourself in the ultimate cinematic experience, where every
          movie is a masterpiece!
        </p>
        <div>
          <p class="text-teal-600">
            Explore Our Movies and Enjoy your Experience Here
          </p>
        </div>
      </section>

      <!-- Cinema Vision Statement -->
      <p
        class="text-base md:text-xl py-2 md:py-6 text-center font-semibold bg-teal-300"
      >
        At Ethio Cinema, we’re redefining entertainment. From blockbuster hits
        to indie gems, every seat feels like the best in the house!
      </p>

      <!-- Top 3 Recent Movies Section -->
      <section class="my-12">
        <h2
          class="md:text-4xl text-xl font-bold mb-8 text-center text-gray-800"
        >
          🔥 Top 3 Recent Releases
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="movie in recentMovies"
            :key="movie.movie_id"
            class="rounded-lg p-6 bg-gradient-to-r from-gray-100 to-gray-200 transition-transform transform hover:scale-105"
          >
            <h3 class="md:text-2xl text-xl font-bold text-gray-800 mb-3">
              {{ movie.title }}
            </h3>
            <p class="text-gray-600 text-sm">
              🎥 Duration:
              <span class="font-medium">{{ movie.duration }} minutes</span>
            </p>
            <img
              v-if="movie.featured_images"
              :src="movie.featured_images"
              alt="Featured Image"
              class="w-full h-48 object-cover mt-4 rounded-lg shadow-md"
            />
            <button
              @click="buyTicket(movie.movie_id)"
              class="px-4 mt-6 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-10"
            >
              🎟️Buy Ticket
            </button>
          </div>
        </div>
      </section>

      <!-- Currently Showing Section -->
      <section class="my-16">
        <h2
          class="md:text-4xl text-xl font-bold mb-8 text-center text-gray-800"
        >
          🎥 Currently Showing
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="movie in currentlyShowingMovies"
            :key="movie.movie_id"
            class="rounded-lg p-6 border border-t-4"
          >
            <h3 class="text-2xl font-bold text-gray-800 mb-3">
              {{ movie.title }}
            </h3>
            <p class="text-gray-600 text-sm">
              ⏳ Duration:
              <span class="font-medium">{{ movie.duration }} minutes</span>
            </p>
            <div class="flex flex-col md:flex-row mt-4">
              <img
                v-if="movie.featured_images"
                :src="movie.featured_images"
                alt="Featured Image"
                class="w-full md:w-1/2 object-cover mb-4 md:mr-4 rounded-lg shadow-md"
              />
              <div class="md:w-1/2">
                <h4 class="font-semibold text-gray-800">🎭 Genres:</h4>
                <ul class="list-disc pl-5 text-gray-600 text-sm">
                  <li
                    v-for="genre in movie.movie_genres"
                    :key="genre.genre.name"
                  >
                    {{ genre.genre.name }}
                  </li>
                </ul>
                <h4 class="font-semibold text-gray-800 mt-4">⭐ Cast:</h4>
                <ul class="list-disc pl-5 text-gray-600 text-sm">
                  <li v-for="star in movie.movie_stars" :key="star.star.name">
                    {{ star.star.name }}
                  </li>
                </ul>
                <h4 class="font-semibold text-gray-800 mt-4">📅 Showtimes:</h4>
                <button
                  @click="buyTicket(movie.movie_id)"
                  class="px-4 py-2 mt-6 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-4"
                >
                  🎟️Buy Ticket
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>

    <!-- Footer Section -->
    <footer
      class="text-white py-6 bg-gradient-to-r from-gray-900 via-gray-800 to-gray-900 shadow-lg"
    >
      <div class="container mx-auto text-center">
        <p class="text-sm">
          &copy; 2024 Ethio Cinema. Your Gateway to the Best Movies.
        </p>
        <div class="flex justify-center space-x-4 mt-2">
          <a href="#" class="text-gray-400 hover:text-white text-sm"
            >Privacy Policy</a
          >
          <a href="#" class="text-gray-400 hover:text-white text-sm"
            >Terms of Use</a
          >
          <a href="#" class="text-gray-400 hover:text-white text-sm"
            >Contact Us</a
          >
        </div>
      </div>
    </footer>
  </div>
</template>
