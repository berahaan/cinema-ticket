<script setup>
import { ref, onMounted, watch } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import Loading from "./Loading.vue";
const errors = ref({});
const { formData } = useMovies();
const { fetchDirectors, directors } = useFetchDirector();
const { fetchGenres, genres } = useFetchGenres();
const { fetchStars, stars } = useFetchStars();
const { fetchMovies } = useFetchMovies();
const { previewOtherImages } = useFile();
const {
  validateDescription,
  validateDirector,
  validateFeaturedImages,
  validateDuration,
  validateOtherImages,
  validateGenre,
  validateTitle,
  validateStar,
} = useValidate();

const { handleFeaturedImageUpload, handleOtherImagesUpload } = useFile();
const { UploadImage } = HandleImageUpload();
const { selectClass, optionClass } = useThemeColor();
const {
  addMovies,
  addGenreToArray,
  removeGenreFromArray,
  getGenreNameById,
  getStarNameById,
  removeStarFromArray,
  autoResize,
  selectedStars,
  selectedStar,
  selectedGenre,
  loading,
  selectedGenres,
  addStarToArray,
} = useAddMovies();

const handleAddMovies = async () => {
  const { featuredImage64, otherImage64 } = useImageBase64Store();
  try {
    await UploadImage(featuredImage64, otherImage64);
    const { featuredImageURL, otherImageURLs } = useImageStore();
    const response = await addMovies(
      formData.value.title,
      formData.value.description,
      formData.value.duration,
      formData.value.director_id,
      selectedStars.value,
      selectedGenres.value,
      featuredImageURL,
      otherImageURLs
    );
    await fetchMovies();
    if (response) {
      console.log("Success ");
    }
  } catch (error) {
    console.error("Errors while adding movies:", error);
  }
};

onMounted(async () => {
  await fetchDirectors();
  await fetchGenres();
  await fetchStars();
});

watch(stars, (newStars) => {
  if (newStars.length > 0) {
    console.log("updated:");
  }
});
</script>

<template>
  <div
    v-if="loading"
    class="flex justify-center items-center min-h-screen bg-white mt-10"
  >
    <Loading />
  </div>

  <div class="max-w-3xl mx-auto mt-16 mb-10" v-else>
    <Form
      @submit="handleAddMovies"
      class="space-y-4 grid grid-cols-1 sm:grid-cols-1"
    >
      <div class="flex flex-col">
        <h2 class="text-xl mb-4">Add Movie</h2>
        <label for="title"> Title: </label>

        <Field
          v-model="formData.title"
          :rules="validateTitle"
          name="title"
          type="text"
          placeholder="Enter movie title"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="selectClass"
        />
        <ErrorMessage name="title" class="text-red-600 text-sm mt-1" />
      </div>

      <!-- Movie Description -->
      <div class="flex flex-col">
        <label for="description" class="font-semibold" :class="selectClass">
          Description:
        </label>
        <Field
          as="textarea"
          v-model="formData.description"
          :rules="validateDescription"
          name="description"
          placeholder="Enter movie description"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none overflow-hidden transition-all duration-200"
          @input="autoResize($event)"
          :class="selectClass"
        />
        <ErrorMessage name="description" class="text-red-600 text-sm mt-1" />
      </div>

      <div class="flex flex-col">
        <label for="duration" class="font-semibold" :class="selectClass"
          >Duration (in minutes):</label
        >
        <Field
          v-model.number="formData.duration"
          type="number"
          :rules="validateDuration"
          name="duration"
          placeholder="Enter movie duration"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="selectClass"
        />
        <ErrorMessage name="duration" class="text-red-600 text-sm mt-1" />
      </div>

      <!-- Director Dropdown -->
      <div class="flex flex-col">
        <label for="director" class="font-semibold" :class="selectClass"
          >Director:</label
        >
        <Field
          v-model="formData.director_id"
          as="select"
          name="director"
          :rules="validateDirector"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="selectClass"
        >
          <option disabled value="">select director</option>
          <option
            v-for="director in directors"
            :key="director.director_id"
            :value="director.director_id"
            :class="optionClass"
          >
            {{ director.name }}
          </option>
        </Field>
        <ErrorMessage name="director" class="text-red-600 text-sm mt-1" />
      </div>
      <div class="flex flex-col">
        <div class="flex flex-col">
          <label for="genre"> Genres: </label>
          <Field
            v-model="selectedGenre"
            as="select"
            name="genre"
            :rules="validateGenre"
            @change="addGenreToArray"
            class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            :class="selectClass"
          >
            <option disabled value="">Select a genre</option>
            <option
              v-for="genre in genres"
              :key="genre.genre_id"
              :value="genre.genre_id"
              :class="optionClass"
            >
              {{ genre.name }}
            </option>
          </Field>
          <ErrorMessage name="genre" class="text-red-600 text-sm mt-1" />
        </div>

        <div class="mb-4">
          <h3 class="font-semibold">Selected Genres:</h3>
          <ul>
            <li
              v-for="id in selectedGenres"
              :key="id"
              class="flex justify-between items-center"
            >
              {{ getGenreNameById(id) }}
              <button @click="removeGenreFromArray(id)" class="text-red-600">
                Remove
              </button>
            </li>
          </ul>
        </div>
      </div>
      <!-- Stars Dropdown -->
      <div class="flex flex-col">
        <label for="stars" class="font-semibold" :class="selectClass"
          >Stars:</label
        >
        <Field
          name="star"
          as="select"
          :rules="validateStar"
          v-model="selectedStar"
          @change="addStarToArray"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="selectClass"
        >
          <option disabled value="">select stars</option>
          <option
            v-for="star in stars"
            :key="star.star_id"
            :value="star.star_id"
          >
            {{ star.name }}
          </option>
        </Field>
        <ErrorMessage name="star" class="text-red-600 text-sm mt-1" />

        <div class="mb-4">
          <h3 class="font-semibold">Selected Stars:</h3>
          <ul>
            <li
              v-for="id in selectedStars"
              :key="id"
              class="flex justify-between items-center"
            >
              {{ getStarNameById(id) }}
              <button @click="removeStarFromArray(id)" class="text-red-600">
                Remove
              </button>
            </li>
          </ul>
        </div>
      </div>

      <div class="flex flex-col">
        <div class="flex flex-col">
          <div>
            <h3 class="text-sm font-semibold text-indigo-400">
              Upload Featured Image:
            </h3>

            <div class="flex items-center justify-center w-full">
              <label
                for="dropzone-featured"
                class="flex flex-col items-center justify-center w-full h-40 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer"
              >
                <div
                  class="flex flex-col items-center justify-center pt-5 pb-6"
                >
                  <svg
                    class="w-8 h-6 mb-4"
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
                  <p class="mb-2 text-sm">
                    <span class="font-semibold">Click to upload</span>
                  </p>
                  <p class="text-xs">choose file (PNG, JPG)</p>
                </div>

                <Field
                  id="dropzone-featured"
                  type="file"
                  :rules="validateFeaturedImages"
                  name="featuredImage"
                  accept="image/*"
                  @change="handleFeaturedImageUpload"
                  class="hidden"
                />
              </label>
            </div>
          </div>
          <ErrorMessage
            name="featuredImage"
            class="text-red-600 text-sm mt-1"
          />
        </div>
      </div>

      <div class="flex flex-col">
        <h3 class="text-sm font-semibold text-indigo-400">
          Upload New Other Images (Max 5):
        </h3>
        <div class="flex items-center justify-center w-full">
          <label
            for="dropzone-other"
            class="flex flex-col items-center justify-center w-full h-40 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer"
          >
            <div class="flex flex-col items-center justify-center pt-5 pb-6">
              <svg
                class="w-8 h-8 mb-4"
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
              <p class="mb-2 text-sm">
                <span class="font-semibold">Click to upload</span>
              </p>
              <p class="text-xs">choose (PNG, JPG)</p>
            </div>

            <Field
              id="dropzone-other"
              type="file"
              name="otherImages"
              accept="image/*"
              :rules="validateOtherImages"
              multiple
              @change="handleOtherImagesUpload"
              class="hidden"
            />
          </label>
        </div>

        <ErrorMessage name="otherImages" class="text-red-600 text-sm mt-1" />
      </div>
      <div v-if="errors.images" class="text-red-600 text-sm mt-1">
        <ul>
          <li v-for="(error, index) in errors.images" :key="index">
            {{ error }}
          </li>
        </ul>
      </div>

      <div v-if="previewOtherImages.length > 0" class="mt-2">
        <h3 class="font-semibold">Selected Other Images:</h3>
        <ul>
          <li>
            <img
              v-for="(image, index) in previewOtherImages"
              :key="index"
              :src="image"
              alt="New Other Image"
              class="w-60 h-auto rounded-md shadow-md border border-gray-500"
            />
            <button
              type="button"
              @click="removeSelectedImage(index)"
              class="text-red-500 ml-2"
            >
              Remove
            </button>
          </li>
        </ul>
      </div>
      <div class="flex flex-col">
        <button
          type="submit"
          class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition duration-200 ease-in-out"
        >
          <span v-if="loading">
            <svg
              class="animate-spin h-5 w-5 mr-2 text-white inline-block"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
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
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
              ></path>
            </svg>
            Adding...
          </span>
          <span v-else>Add Movie</span>
        </button>
      </div>
    </Form>
    <!-- <div
      v-if="notification.message"
      :class="
        notification.type === 'success' ? 'text-green-600' : 'text-red-600'
      "
      class="mt-4 mx-10 text-xl text-center"
    >
      {{ notification.message }}
    </div> -->
  </div>
</template>
