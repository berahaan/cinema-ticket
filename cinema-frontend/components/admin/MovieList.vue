<template>
  <div class="max-w-4xl mx-auto mt-10">
    <h2 class="text-2xl font-bold mb-6 text-center">Movies List</h2>
    <div
      v-for="movie in movies"
      :key="movie.title"
      class="border rounded-lg shadow-lg p-6 mb-6 bg-white"
    >
      <h3 class="text-xl font-semibold">{{ movie.title }}</h3>
      <p class="text-gray-700 mb-2">{{ movie.description }}</p>
      <p class="text-gray-600">Duration: {{ movie.duration }} minutes</p>
      <div class="flex flex-col md:flex-row mt-4">
        <img
          v-if="movie.featured_images"
          :src="movie.featured_images"
          alt="Featured Image"
          class="w-full md:w-1/2 object-cover mb-4 md:mr-4 rounded cursor-pointer transition-transform transform hover:scale-105"
          @click="openModal(movie.movie_images)"
        />
        <div class="md:w-1/2">
          <h4 class="font-semibold">Genres:</h4>
          <ul class="list-disc pl-5">
            <li v-for="genre in movie.movie_genres" :key="genre.genre.name">
              {{ genre.genre.name }}
            </li>
          </ul>
          <h4 class="font-semibold mt-4">Stars:</h4>
          <ul class="list-disc pl-5">
            <li v-for="star in movie.movie_stars" :key="star.star.name">
              {{ star.star.name }}
            </li>
          </ul>
        </div>
      </div>
    </div>

    <!-- Modal for Image Slideshow -->
    <div
      v-if="isModalOpen"
      class="fixed inset-0 bg-black bg-opacity-70 flex items-center justify-center z-50 transition-opacity"
    >
      <div
        class="bg-white p-6 rounded-lg max-w-lg w-full align-center shadow-lg"
      >
        <h3 class="text-lg font-semibold mb-2">Other Images</h3>

        <div class="relative">
          <!-- Display the current image -->
          <img
            :src="currentOtherImages[currentImageIndex]"
            alt="Other Image"
            class="w-full h-auto rounded-md shadow-lg"
          />
          <button
            v-if="currentOtherImages.length > 1"
            @click="prevImage"
            class="absolute left-0 top-1/2 bg-gray-800 text-white px-3 py-1 rounded-full transition-transform transform hover:scale-110 focus:outline-none"
          >
            &lt;
          </button>
          <button
            v-if="currentOtherImages.length > 1"
            @click="nextImage"
            class="absolute right-0 top-1/2 bg-gray-800 text-white px-3 py-1 rounded-full transition-transform transform hover:scale-110 focus:outline-none"
          >
            &gt;
          </button>
        </div>

        <div class="mt-4 flex justify-between">
          <button
            @click="closeModal"
            class="w-full bg-red-500 text-white px-4 py-2 rounded transition-all duration-200 transform hover:scale-105 focus:outline-none"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { gql } from "@apollo/client/core";
import { useNuxtApp } from "#app";

const { $apollo } = useNuxtApp(); // Access Apollo Client

const movies = ref([]); // State to hold movies data
const isModalOpen = ref(false); // State to handle modal visibility
const currentOtherImages = ref([]); // Store the images to display
const currentImageIndex = ref(0); // State to track the current image index

// GraphQL Query to fetch movies
const GET_MOVIES = gql`
  query GetMovies {
    movies {
      title
      description
      duration
      featured_images
      movie_images {
        other_images
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
    }
  }
`;

// Fetch movies data
const fetchMovies = async () => {
  try {
    const response = await $apollo.query({ query: GET_MOVIES });
    movies.value = response.data.movies;
    console.log("Fetched movies:", movies.value);
  } catch (error) {
    console.error("Error fetching movies:", error);
  }
};

// Open Modal and set currentOtherImages
const openModal = (movieImages) => {
  const otherImages = movieImages.map((image) => image.other_images);

  if (otherImages && otherImages.length > 0) {
    currentOtherImages.value = otherImages; // Set the other images for the modal
    currentImageIndex.value = 0; // Reset to first image
    isModalOpen.value = true; // Open the modal
  } else {
    alert("No other images available.");
  }
};

// Function to go to the next image
const nextImage = () => {
  currentImageIndex.value =
    (currentImageIndex.value + 1) % currentOtherImages.value.length;
};

// Function to go to the previous image
const prevImage = () => {
  currentImageIndex.value =
    (currentImageIndex.value - 1 + currentOtherImages.value.length) %
    currentOtherImages.value.length;
};

// Close Modal
const closeModal = () => {
  isModalOpen.value = false;
  currentOtherImages.value = [];
};

// Fetch movies when the component is mounted
onMounted(fetchMovies);
</script>

<style scoped>
body {
  background-color: #f8f9fa; /* Light background for better contrast */
}

h2,
h3 {
  color: #343a40; /* Dark text for better readability */
}

ul {
  list-style-type: disc;
  padding-left: 20px;
}

ul li {
  color: #495057; /* Medium gray for list items */
}
</style>
