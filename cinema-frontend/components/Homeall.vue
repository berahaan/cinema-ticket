<script setup>
import { ref, computed, onMounted } from "vue";
import { useNuxtApp } from "#app";
import { jwtDecode } from "jwt-decode";
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
    console.log("Data here ",data);
    movies.value = data.movies;
   
  } catch (error) {
    console.error("Error fetching movies:", error);
  }
};
onMounted(() => {
  fetchMovies();
});
const recentMovies = computed(() => movies.value.slice(0, 3));
const currentlyShowingMovies = computed(() => movies.value.slice(3, 6));
const buyTicket = (movieId) => {
  console.log("Ticket is clicked now ");
  router.push(`/admin/ticket/${movieId}`);
};
const moveToMovielist =()=>{
  console.log("Clicked ..");
      const accessToken=localStorage.getItem("accessToken");
      if(accessToken){
        const decodedToken = jwtDecode(accessToken);
        const defaultRole =
          decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
        console.log("Default Role for this pages :", defaultRole);
        if (defaultRole === "admin" || defaultRole === "managers") {
          router.push("/admin/Movielist");
        } else if (defaultRole === "regular") {
          router.push("/user/Movielist");
        } else {
          router.push("/");
        }
      }
      else{
        router.push("/auth/login")
      }
  
}
</script>
<template>
  <div>
    <div class="container mx-auto py-12">
      <!-- Welcome Section -->
      <section
        class="text-center py-16 mb-12 bg-gradient-to-br from-purple-600 via-pink-500 to-red-500 text-white rounded-lg shadow-2xl mt-10 w-full"
      >
        <h2 class="text-5xl font-extrabold mb-6 drop-shadow-lg">
          🎬 Welcome to Ethio Cinema! 🌟
        </h2>
        <p class="text-2xl font-light">
          Immerse yourself in the ultimate cinematic experience, where every movie is a masterpiece!
        </p>
        <button
        @click="moveToMovielist"
          class="mt-8 bg-yellow-400 text-gray-800 px-8 py-4 font-bold text-lg rounded-full transition-transform transform hover:scale-110 hover:bg-yellow-500 shadow-md focus:outline-none"
        >
          🎥 Explore Our Outstanding Movies
        </button>
      </section>

      <!-- Cinema Vision Statement -->
      <p
        class="text-xl py-6 text-center bg-gradient-to-br from-blue-500 via-teal-400 to-green-400 text-white rounded-lg shadow-lg font-semibold"
      >
        At Ethio Cinema, we’re redefining entertainment. From blockbuster hits to indie gems, every seat feels like the best in the house!
      </p>

      <!-- Top 3 Recent Movies Section -->
      <section class="my-12">
        <h2 class="text-4xl font-bold mb-8 text-center text-gray-800">
          🔥 Top 3 Recent Releases
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="movie in recentMovies"
            :key="movie.movie_id"
            class="rounded-lg shadow-xl p-6 bg-gradient-to-r from-gray-100 to-gray-200 transition-transform transform hover:scale-105"
          >
            <h3 class="text-2xl font-bold text-gray-800 mb-3">
              {{ movie.title }}
            </h3>
            <p class="text-gray-600 text-sm">
              🎥 Duration: <span class="font-medium">{{ movie.duration }} minutes</span>
            </p>
            <img
              v-if="movie.featured_images"
              :src="movie.featured_images"
              alt="Featured Image"
              class="w-full h-48 object-cover mt-4 rounded-lg shadow-md"
            />
            <button
              @click="buyTicket(movie.movie_id)"
              class="mt-6 bg-teal-500 text-white px-6 py-3 rounded-full text-lg font-medium transition-transform transform hover:scale-110  focus:outline-none mx-12"
            >
              🎟️ Buy Ticket
            </button>
          </div>
        </div>
      </section>

      <!-- Currently Showing Section -->
      <section class="my-16">
        <h2 class="text-4xl font-bold mb-8 text-center text-gray-800">
          🎥 Currently Showing
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="movie in currentlyShowingMovies"
            :key="movie.movie_id"
            class="rounded-lg shadow-xl p-6 bg-gradient-to-r from-gray-100 to-gray-200 transition-transform transform hover:scale-105"
          >
            <h3 class="text-2xl font-bold text-gray-800 mb-3">
              {{ movie.title }}
            </h3>
            <p class="text-gray-600 text-sm">
              ⏳ Duration: <span class="font-medium">{{ movie.duration }} minutes</span>
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
                  class="mt-6 bg-teal-500 text-white px-4 py-3 rounded-full text-lg font-medium transition-transform transform hover:scale-110  focus:outline-none" 
                >
                  🎟️ Buy Ticket
                </button>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>

    <!-- Footer Section -->
    <footer class="text-white py-6 bg-gradient-to-r from-gray-900 via-gray-800 to-gray-900 shadow-lg">
      <div class="container mx-auto text-center">
        <p class="text-sm">
          &copy; 2024 Ethio Cinema. Your Gateway to the Best Movies.
        </p>
        <div class="flex justify-center space-x-4 mt-2">
          <a href="#" class="text-gray-400 hover:text-white text-sm">Privacy Policy</a>
          <a href="#" class="text-gray-400 hover:text-white text-sm">Terms of Use</a>
          <a href="#" class="text-gray-400 hover:text-white text-sm">Contact Us</a>
        </div>
      </div>
    </footer>
  </div>
</template>
