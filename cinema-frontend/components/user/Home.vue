<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { jwtDecode } from "jwt-decode";
import { useNuxtApp } from "#app";
import { GET_MOVIES_QUERY } from "../graphql/queries/GET_DETAIL_QUERY.graphql";
const router = useRouter();
const { $apollo } = useNuxtApp();
const recommendedMovies = ref([]);
const featuredMovies = ref([]);
const loading = ref(false);
const error = ref(null);
const { selectClass } = useThemeColor();
const ViewDetail = (id) => {
  const accessToken = localStorage.getItem("accessToken");
  if (accessToken) {
    const decodedToken = jwtDecode(accessToken);
    const defaultRole =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
    if (defaultRole === "regular") {
      router.push(`/user/${id}`);
    } else {
      router.push("/");
    }
  } else {
    router.push("/login");
  }
};
const BookMovies = (id) => {
  const accessToken = localStorage.getItem("accessToken");
  if (accessToken) {
    const decodedToken = jwtDecode(accessToken);
    const defaultRole =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
   
    if (defaultRole === "admin" || defaultRole === "managers") {
      router.push(`/admin/ticket/${id}`);
    } else if (defaultRole === "regular") {
      router.push(`/user/ticket/${id}`);
    } else {
      router.push("/");
    }
  } else {
    router.push("/login");
  }
};
onMounted(async () => {
  loading.value = true;
  try {
    const { data } = await $apollo.query({
      query: GET_MOVIES_QUERY,
      variables: {
        limit: 12,
        offset: 0,
      },
    });

    recommendedMovies.value = data.movies.slice(0, 4);
    featuredMovies.value = data.movies.slice(0, 4);
  } catch (err) {
    error.value = err;
  } finally {
    loading.value = false;
  }
});
</script>
<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <div
      class="relative bg-gradient-to-r from-teal-500 to-indigo-600 py-20 px-6"
      :class="selectClass"
    >
      <div class="max-w-6xl mx-auto text-center">
        <h1 class="md:text-4xl text-base font-bold mb-4 mt-10">
          Welcome to the MovieZone 🎥
        </h1>
        <p class="text-lg mb-6">
          Explore your favorite movies, bookmark them and book tickets in just a
          few clicks.
        </p>
        <NuxtLink
          to="/user/Movielist"
          class="px-6 py-3 bg-white text-blue-600 font-semibold rounded-lg shadow-lg hover:bg-gray-100"
        >
          Discover Movies
        </NuxtLink>
      </div>
    </div>

    <!-- Recommended Movies Section -->
    <section class="py-10">
      <div class="max-w-6xl mx-auto px-6">
        <h2 class="md:text-2xl text-xl font-bold mb-6">Recommended For You</h2>
        <div v-if="loading" class="text-center text-gray-600">
          Loading movies...
        </div>
        <div v-else-if="error" class="text-center text-red-600">
          Error loading movies: {{ error.message }}
        </div>
        <div
          v-else
          class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6"
        >
          <div
            v-for="movie in recommendedMovies"
            :key="movie.movie_id"
            class="rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition"
          >
            <img
              :src="movie.featured_images"
              alt="Movie Poster"
              class="w-full h-40 object-cover"
            />
            <div class="p-4">
              <h3 class="text-lg font-semibold truncate">
                {{ movie.title }}
              </h3>
              <p class="text-sm mt-2 truncate">
                {{ movie.description }}
              </p>
              <button
                @click="ViewDetail(movie.movie_id)"
                class="mt-4 inline-block text-blue-600 font-semibold hover:underline"
              >
                view details
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Section -->
    <section class="py-10">
      <div class="max-w-6xl mx-auto px-6">
        <h2 class="text-2xl font-bold mb-6">Featured This Week</h2>
        <div v-if="loading" class="text-center text-gray-600">
          Loading featured movies...
        </div>
        <div v-else-if="error" class="text-center text-red-600">
          Error loading movies: {{ error.message }}
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div
            v-for="featured in featuredMovies"
            :key="featured.movie_id"
            class="flex items-center rounded-lg shadow-lg overflow-hidden"
          >
            <img
              :src="featured.featured_images"
              alt="Featured Movie"
              class="w-40 h-40 object-cover"
            />
            <div class="p-4">
              <h3 class="text-xl font-semibold">
                {{ featured.title }}
              </h3>
              <p class="mt-2 truncate">
                {{ featured.description }}
              </p>
              <button
                @click="BookMovies(featured.movie_id)"
                class="mt-4 inline-block text-blue-600 font-semibold hover:underline"
              >
                Book now
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
