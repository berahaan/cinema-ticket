<script setup>
import { onMounted } from "vue";
import { jwtDecode } from "jwt-decode";
import {
  StarIcon as StarSolid,
  ArrowLeftIcon,
  ArrowRightIcon,
} from "@heroicons/vue/solid";
import { StarIcon as StarOutline } from "@heroicons/vue/outline";
import Loading from "../admin/Loading.vue";
let userId = null;
const movie = useMoviesStore();
const {
  fetchMovies,
  searchQuery,
  genreQuery,
  directorQuery,
  goToNextPage,
  isloading,
  scheduleQuery,
  goToPreviousPage,
  currentPage,
  noMoviesFound,
} = useFetchMovies();
const { scheduleOptions } = useMovies();
const { formatScheduleDateTime,isScheduleExpired,isvalid} = useFormatSchedule();
const { fetchGenres, genres } = useFetchGenres();
const { fetchDirectors, directors } = useFetchDirector();
const { refreshPage } = useRefresh();
const { buyTicket } = useTicket();
const {
  nextImage,
  prevImage,
  openModal,
  isModalOpen,
  closeModal,
  currentImageIndex,
  currentOtherImages,
} = useModals();

const {
  openRatingModal,
  isModal,
  currentMovieId,
  selectedRating,
  submitRating,
} = useRating();
const { selectClass, optionClass } = useThemeColor();
const { toggleBookmark } = useBookmarks();

if (import.meta.client) {
  const token = localStorage.getItem("accessToken");
  if (token) {
    const decodedToken = jwtDecode(token);
    userId = decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
  }
}
// Function to get the current user's rating for a specific movie
const getCurrentUserRating = (movie) => {
  // Check if rate_movies is defined and if the current user has rated the movie
  if (movie.rate_movies && Array.isArray(movie.rate_movies)) {
    const userRating = movie.rate_movies.find((rate) => rate.user_id == userId);
    return userRating; // Return the user's rating or undefined if not found
  }
  return undefined; // If rate_movies is undefined or not an array, return undefined
};
onMounted(async () => {
  await fetchMovies();
  await fetchGenres();
  await fetchDirectors();
});
</script>
<template>
  <div class="w-full mx-auto mt-10">
    <div class="flex justify-center fixed top-28 right-2 left-2 mb-60">
      <div class="relative">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Search by title..."
          class="border p-2 pl-8 pr-2 rounded w-full text-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition mb-4"
          :class="selectClass"
        />
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="w-5 h-5 absolute left-2 top-1/3 transform -translate-y-1/2 text-gray-600"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <circle cx="11" cy="11" r="8" />
          <line x1="21" y1="21" x2="16.65" y2="16.65" />
        </svg>
      </div>
      <select
        v-model="genreQuery"
        id="genre"
        class="border border-gray-300 rounded-md p-1 h-10 focus:outline-none focus:ring-2 focus:ring-blue-500 ml-2"
        :class="selectClass"
      >
        <option disabled value="">Select a Genre</option>
        <option value="All">All</option>
        <option
          v-for="genre in genres"
          :key="genre.genre_id"
          :value="genre.name"
          :class="optionClass"
        >
          {{ genre.name }}
        </option>
      </select>
      <select
        v-model="directorQuery"
        id="director"
        class="border border-gray-300 rounded-md p-1 h-10 focus:outline-none focus:ring-2 focus:ring-blue-500 ml-2"
        :class="selectClass"
      >
        <option disabled value="">select director</option>
        <option value="All">All</option>
        <option
          v-for="director in directors"
          :key="director.director_id"
          :value="director.name"
          :class="optionClass"
        >
          {{ director.name }}
        </option>
      </select>
      <select
        v-model="scheduleQuery"
        id="genre"
        class="border border-gray-300 rounded-md p-1 h-10 focus:outline-none focus:ring-2 focus:ring-blue-500 ml-2"
        :class="selectClass"
      >
        <option disabled value="">Select schedules</option>
        <option v-for="option in scheduleOptions" :key="option" :value="option">
          {{ option }}
        </option>
      </select>
      <button @click="refreshPage(fetchMovies)" class="ml-4 mb-8 mt-2">
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

    <div class="w-full mx-auto mt-10 pt-40" :class="selectClass">
      <div
        v-if="isloading"
        class="flex justify-center items-center min-h-screen bg-white mt-10"
      >
        <Loading />
      </div>
      <template v-else>
        <!-- <div class="max-w-2xl mx-auto mt-10 mb-10"></div> -->
        <div v-if="noMoviesFound" :class="selectClass" class="mx-10 mb-80">
          <p>Unfortunaely no movies is found with ur searches</p>
        </div>
        <div v-if="movie.movies.length > 0">
          
          <div
            v-for="movie in movie.movies"
            :key="movie.movie_id"
            class="border rounded-lg shadow-lg p-6 transition-transform transform hover:shadow-xl"
            :class="selectClass"
          >
            <div
              class="flex justify-between items-center border-b-4 dark:border-gray-600"
            >
              <h3 class="text-4xl mb-4 font-bold" :class="selectClass">
                {{ movie.title }}
              </h3>
              <div class="relative inline-flex items-center">
                <div
                  class="relative inline-flex items-center"
                  :class="selectClass"
                >
                  <svg
                    v-if="movie.bookmarks?.length > 0"
                    @click="toggleBookmark(movie.movie_id)"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    class="w-6 h-6 cursor-pointer"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M3 3l1.664 1.664M21 21l-1.5-1.5m-5.485-1.242L12 17.25 4.5 21V8.742m.164-4.078a2.15 2.15 0 0 1 1.743-1.342 48.507 48.507 0 0 1 11.186 0c1.1.128 1.907 1.077 1.907 2.185V19.5M4.664 4.664L19.5 19.5"
                    />
                  </svg>
                  <svg
                    v-else
                    @click="toggleBookmark(movie.movie_id)"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    class="w-6 h-6 cursor-pointer"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z"
                    />
                  </svg>
                </div>
              </div>
            </div>

            <h4
              class="font-semibold text-blue-800 dark:text-blue-400 text-lg mb-2"
            >
              📖 Description
            </h4>
            <p class="text-gray-700 mb-2" :class="selectClass">
              {{ movie.description }}
            </p>

            <div class="flex flex-col md:flex-row md:space-x-4 mt-4">
              <p class="text-gray-600 mb-2" :class="selectClass">
                ⏳ Duration:
                <span class="font-medium" :class="selectClass"
                  >{{ movie.duration }} minutes</span
                >
                <span class="font-medium mx-4">
                  average Rate:
                  {{
                    movie.average_rating
                      ? movie.average_rating.average_rating.toFixed(1) +
                        " stars"
                      : "Not rated yet"
                  }}
                </span>
              </p>
              <p class="text-gray-600 mb-2" :class="selectClass">
                🎬 Director:
                <span class="font-medium">{{ movie.director.name }}</span>
              </p>
              <p class="font-medium">
                Total Tickets Sold:
                {{
                  movie.tickets_aggregate?.aggregate?.sum?.ticket_quantity || 0
                }}
              </p>
            </div>

            <div class="flex flex-col md:flex-row mt-4">
              <div
                class="relative w-full md:w-1/2 lg:w-1/3 max-h-96 mb-4 rounded-lg shadow-md cursor-pointer hover:scale-105 transition-transform duration-200"
                @click="openModal(movie.movie_images)"
              >
                <div
                  class="absolute inset-0 flex items-center justify-center bg-teal-900 bg-opacity-50 text-white text-lg font-semibold opacity-0 hover:opacity-100 transition-opacity duration-200"
                >
                  Click to Expand
                </div>

                <img
                  v-if="movie.featured_images"
                  :src="movie.featured_images"
                  alt="Featured Image"
                  class="w-full h-full object-cover rounded-lg"
                />
              </div>

              <div class="md:w-1/2">
                <div class="mt-4 p-4 rounded-lg shadow-sm" :class="selectClass">
                  <h4
                    class="flex items-center font-semibold text-blue-800 text-lg mb-2"
                    :class="selectClass"
                  >
                    🎭 Genres
                  </h4>
                  <ul class="space-y-2">
                    <li
                      v-for="genre in movie.movie_genres"
                      :key="genre.genre.name"
                      class="text-gray-700 dark:text-gray-300 rounded-lg p-3 shadow-md"
                      :class="selectClass"
                    >
                      <span
                        class="font-medium text-gray-800"
                        :class="selectClass"
                      >
                        {{ genre.genre.name }}
                      </span>
                    </li>
                  </ul>
                  <div>
                    <!-- Only show the current user's rating -->
                    <ul v-if="getCurrentUserRating(movie)">
                      <li
                        :key="getCurrentUserRating(movie).rating"
                        class="rounded-lg p-3 shadow-md"
                        :class="selectClass"
                      >
                        ⭐ You have Rated:
                        {{ getCurrentUserRating(movie).rating }} stars
                      </li>
                    </ul>

                    <!-- If the current user hasn't rated the movie -->
                    <p v-else class="text-gray-500 italic">
                      ⭐ Rate: Not rated
                    </p>
                  </div>
                </div>

                <div>
                  <h4
                    class="flex items-center font-semibold text-blue-800 text-lg mb-2 mx-3"
                    :class="selectClass"
                  >
                    Stars
                  </h4>
                  <ul class="list-disc pl-8 mb-4 space-y-1">
                    <li
                      v-for="star in movie.movie_stars"
                      :key="star.star.name"
                      class="text-gray-700"
                      :class="selectClass"
                    >
                      {{ star.star.name }}
                    </li>
                  </ul>
                </div>

                <div
                  v-if="isModal && currentMovieId === movie.movie_id"
                  class="fixed inset-0 bg-opacity-50 flex items-center justify-center z-50"
                >
                  <div
                    class="p-6 rounded-lg shadow-lg space-y-4"
                    :class="selectClass"
                  >
                    <p class="text-lg font-semibold">Rate this movie:</p>

                    <div class="flex justify-center space-x-2">
                      <div
                        v-for="star in 5"
                        :key="star"
                        @click="selectedRating = star"
                        class="cursor-pointer"
                        :class="selectClass"
                      >
                        <component
                          :is="star <= selectedRating ? StarSolid : StarOutline"
                          class="w-8 h-8 text-yellow-500"
                          :class="selectClass"
                        />
                      </div>
                    </div>

                    <div class="flex justify-end space-x-2">
                      <button
                        @click="submitRating(currentMovieId)"
                        class="px-4 py-2 bg-green-600 dark:bg-green-800 text-white rounded-lg"
                      >
                        Submit
                      </button>
                      <button
                        @click="
                          isModal = false;
                          currentMovieId = null;
                        "
                        class="px-4 py-2 bg-red-600 dark:bg-red-800 text-white rounded-lg"
                      >
                        Cancel
                      </button>
                    </div>
                  </div>
                </div>

                <div
                  class="mt-4 p-4 bg-gray-100 rounded-lg shadow-sm"
                  :class="selectClass"
                >
                  <h4
                    class="flex items-center font-semibold text-blue-800 dark:text-blue-400 text-lg mb-2"
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
             
               <p :class="selectClass">
                <span class="font-semibold mr-2" :class="selectClass">
                {{
                  formatScheduleDateTime(
                    schedule.schedule_date,
                    schedule.start_time,
                    schedule.end_time
                  )
                }}
              </span>
              <span
            v-if="isScheduleExpired(schedule)"
             class="text-red-500 font-semibold mr-2"
              :class="selectClass"
                 >
                (Expired)
                     </span>

      <!-- Valid Schedule Label -->
                  <span
                     v-else
                     class="text-green-500 font-semibold mr-2"
                     :class="selectClass"
                   >
                     (Valid)
                   </span>
                 <span class="text-sm mr-2" :class="selectClass">
                   | Price:
                   <span class="font-semibold text-green-600 mr-2" :class="selectClass">
                     {{ schedule.ticket_price }} birr
                   </span>
                 </span>
               </p>
             </p>
           </li>
                  </ul>
                  <button
                    @click="openRatingModal(movie.movie_id)"
                    class="px-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-6 mt-2"
                  >
                    Rate
                  </button>
                  <!-- <div   v-for="schedule in movie.movie_schedules"
                  :key="schedule.start_time">
                </div> -->
                 <button
                  v-if="movie.movie_schedules?.length>0&& isvalid"
                  @click="buyTicket(movie.movie_id)"
                  
                  class="px-4 py-2 bg-blue-600 dark:bg-blue-800 text-white rounded-lg mx-6 mt-2"
                >
                  Buy Ticket
                </button>

                <button v-else  @click="buyTicket(movie.movie_id)"
                  disabled
                  class="px-4 py-2 bg-blue-600 disabled:bg-gray-400 dark:bg-gray-800 text-white rounded-lg mx-6 mt-2 disabled:cursor-not-allowed">
                  Buy Ticket
                </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div
          v-if="isModalOpen"
          class="fixed inset-0 bg-black bg-opacity-80 flex items-center justify-center z-50 transition-opacity duration-300 ease-in-out"
        >
          <div
            class="bg-teal-300 rounded-lg shadow-lg max-w-xl w-full p-2 relative"
          >
            <h3 class="text-lg font-semibold mb-2 text-center">
              Image Previews
            </h3>
            <div v-if="isNotherImagesExist" class="text-center text-gray-500">
              <p>No other images available for this Movies. Thank you!</p>
            </div>
            <div v-else class="relative">
              <img
                :src="currentOtherImages[currentImageIndex]"
                alt="Other Image"
                class="w-full h-auto max-h-96 rounded-lg shadow-md transition-transform duration-300 transform hover:scale-105"
              />
              <button
                v-if="currentOtherImages.length > 1"
                @click="prevImage"
                class="absolute left-4 top-1/2 transform -translate-y-1/2 bg-gray-800 text-white px-3 py-2 rounded-full transition-transform hover:scale-110 focus:outline-none"
              >
                &lt;
              </button>
              <button
                v-if="currentOtherImages.length > 1"
                @click="nextImage"
                class="absolute right-4 top-1/2 transform -translate-y-1/2 bg-gray-800 text-white px-3 py-2 rounded-full transition-transform hover:scale-110 focus:outline-none"
              >
                &gt;
              </button>
            </div>

            <p
              v-if="!isNotherImagesExist"
              class="mt-2 text-center text-gray-600"
            >
              Image {{ currentImageIndex + 1 }} of
              {{ currentOtherImages.length }}
            </p>

            <button
              @click="closeModal"
              class="absolute top-2 right-2 bg-teal-500 text-white px-2 py-1 rounded-full transition-all duration-200 transform hover:bg-red-600 focus:outline-none shadow-lg mb-4"
            >
              X
            </button>
          </div>
        </div>
        <div
          class="pagination-controls py-4 flex justify-center items-center space-x-4 rounded-lg mt-8"
          v-if="movie.totalPages"
        >
          <button
            @click="goToPreviousPage"
            :disabled="currentPage === 1"
            class="flex items-center px-4 py-2 rounded-md shadow-md transition-all bg-teal-500 text-white font-medium hover:bg-teal-600 disabled:bg-gray-300 disabled:text-gray-500 disabled:cursor-not-allowed"
          >
            <ArrowLeftIcon class="h-5 w-5 mr-2" />
          </button>
          <!-- Pagination Info -->
          <span class="font-semibold" :class="selectClass">
            {{
              noMoviesFound
                ? "Page 1 of total 1"
                : `Page ${currentPage} of ${movie.totalPages}`
            }}
          </span>

          <!-- Next Button -->
          <button
            @click="goToNextPage"
            :disabled="currentPage === movie.totalPages || noMoviesFound"
            class="flex items-center px-4 py-2 rounded-md shadow-md transition-all bg-teal-500 text-white font-medium hover:bg-teal-600 disabled:bg-gray-300 disabled:text-gray-500 disabled:cursor-not-allowed"
          >
            <ArrowRightIcon class="h-5 w-5 ml-2" />
          </button>
        </div>
      </template>
    </div>
  </div>
</template>
