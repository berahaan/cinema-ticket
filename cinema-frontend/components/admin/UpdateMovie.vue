<script setup>
import { onMounted } from "vue";
import { useRoute } from "vue-router";
import { ArrowLeftIcon } from "@heroicons/vue/solid";
const route = useRoute();
const movie_id = parseInt(route.params.id);
const { goBack } = useGobackArrow();
const {
  fetchMovieDetails,
  formData,
  previewOtherImages,
  otherImages,
  currentImage,
  currentOtherImage,
  availableDirectors,
  availableGenres,
  availableStars,
} = useMovies();
const {
  handleFeaturedImageUpload,
  handleOtherImagesUpload,
  previewFeaturedImage,
} = useFile();

const { updateMovies, isUpdated, isUpdateloading, loading } = useUpdateMovie();
const { autoResize } = useAddMovies();
const { UploadImage } = HandleImageUpload();
const { selectClass, labelClass } = useThemeColor();
const { refreshPage } = useRefresh();
const { fetchMovies } = useFetchMovies();
watch(
  () => currentOtherImage.value,
  (newVal) => {
    console.log("Updated currentOtherImage in components :", newVal);
  }
);
// First, handles image base64 from stores
const handleUpdateMovie = async () => {
  const { featuredImage64, otherImage64 } = useImageBase64Store();
  if (!featuredImage64 && otherImage64.length == 0) {
    console.log("Both images is not selected ...");
    console.log("featuredImages is now ", currentImage.value);
    console.log("otherImages is now ", otherImages.value);
    await updateMovies(
      movie_id,
      formData.value.title,
      formData.value.description,
      formData.value.director_id,
      formData.value.genres,
      formData.value.stars,
      null,
      null,
      formData.value.duration
    );
  } else if (featuredImage64 != "" && otherImage64.length == 0) {
    alert(
      "Only featured images is being saved now and other images is remain same "
    );
    await UploadImage(featuredImage64, null);
    const { featuredImageURL } = useImageStore();
    await updateMovies(
      movie_id,
      formData.value.title,
      formData.value.description,
      formData.value.director_id,
      formData.value.genres,
      formData.value.stars,
      featuredImageURL,
      null,
      formData.value.duration
    );
  } else if (featuredImage64 == "" && otherImage64.length != 0) {
    alert(
      "Only other images is being saved now and other images is remain same "
    );
    await UploadImage(null, otherImage64);
    const { otherImageURLs } = useImageStore();
    await updateMovies(
      movie_id,
      formData.value.title,
      formData.value.description,
      formData.value.director_id,
      formData.value.genres,
      formData.value.stars,
      null,
      otherImageURLs,
      formData.value.duration
    );
  } else {
    // if both are selected here then fix both of them
    try {
      await UploadImage(featuredImage64, otherImage64);
      const { featuredImageURL, otherImageURLs } = useImageStore();
      await updateMovies(
        movie_id,
        formData.value.title,
        formData.value.description,
        formData.value.director_id,
        formData.value.genres,
        formData.value.stars,
        featuredImageURL,
        otherImageURLs,
        formData.value.duration
      );
    } catch (error) {
      console.error("Error updating movie:", error);
    }
  }
};

onMounted(() => {
  fetchMovieDetails(movie_id);
});
</script>

<template>
  <div class="text-gray-200 py-10 mt-6">
    <div class="mt-6"></div>
    <div
      class="max-w-4xl mx-auto p-8 rounded-lg shadow-lg border border-gray-600"
    >
      <div class="flex justify-between">
        <h1 class="text-3xl font-bold mb-8 text-indigo-400">
          Edit Movie Details
        </h1>
        <div class="flex">
          <button @click="refreshPage(fetchMovies)" class="ml-4 mb-8 mx-10">
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
          <button
            @click="goBack"
            class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200 mb-10"
          >
            <ArrowLeftIcon class="h-5 w-5" />
            <span>Go Back</span>
          </button>
        </div>
      </div>
      <div v-if="loading" class="flex justify-center items-center">
        <svg
          class="animate-spin h-5 w-5 text-gray-400"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
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
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 2.8 1.2 5.3 3.1 7l2.9-1.709z"
          ></path>
        </svg>
        <span class="ml-2 text-gray-400">Loading movie details...</span>
      </div>

      <form v-else @submit.prevent="handleUpdateMovie" class="space-y-8">
        <div>
          <label
            for="title"
            class="block text-sm font-semibold text-gray-300"
            :class="labelClass"
            >Title:</label
          >
          <input
            v-model="formData.title"
            type="text"
            id="title"
            required
            class="mt-2 block w-full px-4 py-2 border rounded-md focus:ring-indigo-500 focus:border-indigo-500 text-gray-200"
            placeholder="Enter movie title"
            :class="selectClass"
          />
        </div>

        <!-- Duration -->
        <div>
          <label
            for="duration"
            class="block text-sm font-semibold text-gray-300"
            :class="labelClass"
            >Duration (mins):</label
          >
          <input
            v-model="formData.duration"
            type="number"
            id="duration"
            required
            class="mt-2 block w-full px-4 py-2 border rounded-md focus:ring-indigo-500 focus:border-indigo-500 text-gray-200"
            placeholder="Enter duration"
            :class="selectClass"
          />
        </div>

        <!-- Description -->
        <div>
          <label
            for="description"
            class="block text-sm font-semibold text-gray-300"
            :class="labelClass"
            >Description:</label
          >
          <textarea
            v-model="formData.description"
            id="description"
            rows="4"
            class="mt-2 block w-full px-4 py-2 border rounded-md focus:ring-indigo-500 focus:border-indigo-500 text-gray-200"
            placeholder="Enter movie description"
            @input="autoResize($event)"
            :class="selectClass"
          ></textarea>
        </div>

        <!-- Genres Dropdown -->
        <div>
          <label
            for="genres"
            class="block text-sm font-semibold text-gray-300"
            :class="labelClass"
            >Genres:</label
          >
          <select
            v-model="formData.genres"
            id="genres"
            multiple
            class="mt-2 block w-full px-4 py-2 border-gray-600 rounded-md focus:ring-indigo-500 focus:border-indigo-500 text-gray-200"
            :class="selectClass"
          >
            <option
              v-for="genre in availableGenres"
              :key="genre.genre_id"
              :value="genre.genre_id"
            >
              {{ genre.name }}
            </option>
          </select>
        </div>

        <div>
          <label
            for="stars"
            class="block text-sm font-semibold text-gray-300"
            :class="labelClass"
            >Stars:</label
          >
          <select
            v-model="formData.stars"
            id="stars"
            multiple
            class="mt-2 block w-full px-4 py-2 border rounded-md focus:ring-indigo-500 focus:border-indigo-500 text-gray-200"
            :class="selectClass"
          >
            <option
              v-for="star in availableStars"
              :key="star.star_id"
              :value="star.star_id"
            >
              {{ star.name }}
            </option>
          </select>
        </div>

        <!-- Director Dropdown -->
        <div>
          <label
            for="director"
            class="block text-sm font-semibold text-gray-300"
            :class="labelClass"
            >Director:</label
          >
          <select
            v-model="formData.director_id"
            id="director"
            class="mt-2 block w-full px-4 py-2 rounded-md text-gray-200"
            :class="selectClass"
          >
            <option
              v-for="director in availableDirectors"
              :key="director.director_id"
              :value="director.director_id"
            >
              {{ director.name }}
            </option>
          </select>
        </div>
        <div>
          <h3 class="text-sm font-semibold text-indigo-400">
            Current Featured Image:
          </h3>
          <img
            v-if="currentImage"
            :src="currentImage"
            alt="Featured Image"
            class="mt-4 w-80 h-auto rounded-md shadow-md border border-gray-500"
          />
        </div>

        <div v-if="previewFeaturedImage">
          <h3 class="text-sm font-semibold text-indigo-400">
            New Featured Image Preview:
          </h3>
          <img
            :src="previewFeaturedImage"
            alt="New Featured Image"
            class="mt-4 w-full h-auto rounded-md shadow-md border border-gray-500"
          />
        </div>

        <div>
          <h3 class="text-sm font-semibold text-indigo-400">
            Upload New Featured Image:
          </h3>
          <div class="flex items-center justify-center w-full">
            <label
              for="dropzone-featured"
              class="flex flex-col items-center justify-center w-full h-64 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer dark:border-gray-600 dark:hover:border-gray-500"
              :class="selectClass"
            >
              <div class="flex flex-col items-center justify-center pt-5 pb-6">
                <svg
                  class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 20 16"
                >
                  <path
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
                  />
                </svg>
                <p class="mb-2 text-sm text-gray-500 dark:text-gray-400">
                  <span class="font-semibold">Click to upload</span> or drag and
                  drop
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400">PNG, JPG</p>
              </div>
              <input
                id="dropzone-featured"
                type="file"
                accept="image/*"
                @change="handleFeaturedImageUpload"
                class="hidden"
              />
            </label>
          </div>
        </div>

        <div>
          <h3 class="text-sm font-semibold text-indigo-400">
            Current Other Images:
          </h3>
          <div class="grid grid-cols-2 gap-4 mt-4">
            <img
              v-if="currentOtherImage.length > 0"
              v-for="(image, index) in currentOtherImage"
              :key="index"
              :src="image"
              alt="Other Image"
              class="w-40 h-auto rounded-md shadow-md border border-gray-500"
            />
          </div>
        </div>
        <div>
          <h3 class="text-sm font-semibold text-indigo-400">
            Upload New Other Images (Max 5):
          </h3>
          <div class="flex items-center justify-center w-full">
            <label
              for="dropzone-other"
              class="flex flex-col items-center justify-center w-full h-64 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer dark:border-gray-600 dark:hover:border-gray-500"
              :class="selectClass"
            >
              <div class="flex flex-col items-center justify-center pt-5 pb-6">
                <svg
                  class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 20 16"
                >
                  <path
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
                  />
                </svg>
                <p class="mb-2 text-sm text-gray-500 dark:text-gray-400">
                  <span class="font-semibold">Click to upload</span> or drag and
                  drop
                </p>
                <p class="text-xs text-gray-500 dark:text-gray-400">
                  SVG, PNG, JPG or GIF (MAX. 800x400px)
                </p>
              </div>
              <input
                id="dropzone-other"
                type="file"
                accept="image/*"
                multiple
                @change="handleOtherImagesUpload"
                class="hidden"
              />
            </label>
          </div>
          <p class="text-sm text-gray-400 mt-1">
            You can upload up to 5 images for "Other Images".
          </p>
        </div>

        <!-- Submit button -->
        <div class="text-right">
          <button
            type="submit"
            :disabled="isUpdateloading"
            class="px-6 py-2 bg-indigo-600 hover:bg-indigo-700 text-white font-medium rounded-md shadow-md focus:ring-indigo-500 focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-700 disabled:bg-gray-600"
          >
            {{ isUpdateloading ? "Updating..." : "Update Movie" }}
          </button>
        </div>
        <div v-if="isUpdated" class="text-xl">Movies updated successfully</div>
        <!-- <div v-if="!isUpdated" class="text-xl">
          Movies not updated successfully please try again...
        </div> -->
      </form>
    </div>
  </div>
</template>
