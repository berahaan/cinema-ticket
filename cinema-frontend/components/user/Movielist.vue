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
const filterMovieInMobile = ref(false);
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
const { formatScheduleDateTime, isScheduleExpired } = useFormatSchedule();
const { fetchGenres, genres } = useFetchGenres();
const { fetchDirectors, directors } = useFetchDirector();
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
  currentMovieId,
  selectedRating,
  submitRating,
  israte,
} = useRating();
const { selectClass, optionClass } = useThemeColor();
const { toggleBookmark, isBookmarkedCheck, getBookmarks } = useBookmarks();
const filterMovieMobile = () => {
  filterMovieInMobile.value = !filterMovieInMobile.value;
};
if (import.meta.client) {
  const token = localStorage.getItem("accessToken");
  if (token) {
    const decodedToken = jwtDecode(token);
    userId = decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
  }
}

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
  await getBookmarks();
  await fetchDirectors();
});
</script>
<template>
  <div class="w-full mx-auto mt-10">
    <div class="flex justify-center fixed top-28 right-2 left-2 mb-40">
      <div
        class="hidden md:flex justify-center fixed top-28 right-2 left-2 mb-40"
      >
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
          <option
            v-for="option in scheduleOptions"
            :key="option"
            :value="option"
          >
            {{ option }}
          </option>
        </select>
      </div>
      <div class="md:hidden">
        <div class="relative flex">
          <input
            type="text"
            v-model="searchQuery"
            placeholder="Search by title..."
            class="border p-1 pl-8 pr-2 rounded w-48 text-gray-600 outline-none focus:ring-2 focus:ring-blue-500 transition mb-4"
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

          <button
            @click="filterMovieMobile"
            class="text-black md:hidden flex mt-2"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 24 24"
              class="w-5 h-5"
            >
              <path
                fill-rule="evenodd"
                clip-rule="evenodd"
                d="M2 5.5A1.5 1.5 0 0 1 3.5 4h17a1.5 1.5 0 0 1 1.5 1.5 1.5 1.5 0 0 1-.297.9l-5.703 7.604a2 2 0 0 0-.375 1.161v3.085a1.5 1.5 0 0 1-.879 1.372l-3 1.364a1.5 1.5 0 0 1-2.121-1.372v-3.085a2 2 0 0 0-.375-1.161L2.297 6.4A1.5 1.5 0 0 1 2 5.5z"
              />
            </svg>

            <span class="mb-4">Filter</span>
          </button>
        </div>
        <div class="grid grid-cols-3 gap-2" v-if="filterMovieInMobile">
          <select
            v-model="genreQuery"
            id="genre"
            class="border border-gray-300 rounded-md p-1 h-8 text-sm w-full focus:outline-none focus:ring-2 focus:ring-blue-500"
            :class="selectClass"
          >
            <option disabled value="">Genre</option>
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
            class="border border-gray-300 rounded-md p-1 h-8 text-sm w-full focus:outline-none focus:ring-2 focus:ring-blue-500"
            :class="selectClass"
          >
            <option disabled value="">Director</option>
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
            id="schedule"
            class="border border-gray-300 rounded-md p-1 h-8 text-sm w-full focus:outline-none focus:ring-2 focus:ring-blue-500"
            :class="selectClass"
          >
            <option disabled value="">Schedule</option>
            <option
              v-for="option in scheduleOptions"
              :key="option"
              :value="option"
            >
              {{ option }}
            </option>
          </select>
        </div>
      </div>
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
          <p>Unfortunaely no movies is found with you searches</p>
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
              <h3
                class="md:text-4xl text-xl mb-4 font-bold"
                :class="selectClass"
              >
                {{ movie.title }}
              </h3>
              <div class="relative inline-flex items-center">
                <div
                  class="relative inline-flex items-center"
                  :class="selectClass"
                >
                  <button
                    v-if="isBookmarkedCheck(movie.movie_id)"
                    @click="toggleBookmark(movie.movie_id)"
                    class="flex items-center space-x-2"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      viewBox="0 0 24 24"
                      class="w-6 h-6"
                    >
                      <path
                        d="M17 3H7a2 2 0 0 0-2 2v16l7-4 7 4V5a2 2 0 0 0-2-2z"
                      />
                    </svg>
                    saved
                  </button>

                  <button
                    v-else
                    @click="toggleBookmark(movie.movie_id)"
                    class="flex items-center space-x-2"
                  >
                    <svg
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
                    save
                  </button>
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

            <div class="grid grid-cols-1 md:grid-cols-3 mt-4">
              <p class="text-gray-600 mb-2" :class="selectClass">
                ⏳ Duration:
                <span class="font-medium" :class="selectClass"
                  >{{ movie.duration }} minutes</span
                >
                <span class="font-medium mx-6 md:mx-4">
                  average Rate:
                  {{
                    movie.average_rating
                      ? movie.average_rating.average_rating.toFixed(1) + "stars"
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
                <span class="font-medium">{{
                  movie.tickets_aggregate?.aggregate?.sum?.ticket_quantity || 0
                }}</span>
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
                  <div
                    class="flex items-center font-semibold text-blue-800 md:text-lg text-base mb-2"
                    :class="selectClass"
                  >
                    🎭Genres
                    <span>
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
                    </span>
                  </div>

                  <div>
                    <ul v-if="getCurrentUserRating(movie)">
                      <li
                        :key="getCurrentUserRating(movie).rating"
                        class="rounded-lg p-3 shadow-md"
                        :class="selectClass"
                      >
                        ⭐ You have Rated:
                        {{ getCurrentUserRating(movie).rating }}
                        stars
                      </li>
                    </ul>
                    <!-- If the current user hasn't rated the movie -->
                    <p v-else class="text-gray-500 italic">
                      ⭐ Rate: Not rated
                    </p>
                  </div>
                </div>

                <div>
                  <div
                    class="flex items-center font-semibold text-blue-800 text-lg mb-4 mx-3"
                    :class="selectClass"
                  >
                    Stars
                    <span>
                      <ul
                        class="pl-8 mb-4 space-y-1 grid-cols-1 md:grid-cols-4 text-base"
                      >
                        <li
                          v-for="star in movie.movie_stars"
                          :key="star.star.name"
                          class="rounded-lg p-3 shadow-md mt-2"
                          :class="selectClass"
                        >
                          {{ star.star.name }}
                        </li>
                      </ul>
                    </span>
                  </div>
                </div>
                <!-- 
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
                </div> -->

                <div
                  class="mt-4 bg-gray-100 rounded-lg shadow-md"
                  :class="selectClass"
                  v-if="movie.movie_schedules?.length > 0"
                >
                  <h4
                    class="flex items-center font-semibold text-blue-800 dark:text-blue-400 text-lg mb-2"
                  >
                    📅 Scheduled Time
                  </h4>
                  <ul>
                    <li
                      v-for="schedule in movie.movie_schedules"
                      :key="schedule.start_time"
                      class="text-gray-700 rounded-lg p-3 shadow-md"
                      :class="selectClass"
                    >
                      <div
                        class="flex flex-col md:flex-row items-start md:items-center gap-2 text-sm md:text-base p-2 rounded-md shadow-sm"
                      >
                        <!-- Date and Time -->
                        <p class="font-semibold text-gray-700">
                          {{
                            formatScheduleDateTime(
                              schedule.schedule_date,
                              schedule.start_time,
                              schedule.end_time
                            )
                          }}
                        </p>

                        <!-- Status: Valid/Expired -->
                        <span
                          v-if="isScheduleExpired(schedule)"
                          class="text-green-600 dark:text-green-400 font-semibold"
                        >
                          Valid
                        </span>
                        <span
                          v-else
                          class="text-red-500 dark:text-red-400 font-semibold"
                        >
                          Expired
                        </span>

                        <!-- Ticket Price -->
                        <span class="text-gray-600 dark:text-gray-400">
                          | Price:
                          <span
                            class="font-semibold text-green-700 dark:text-green-500"
                          >
                            {{ schedule.ticket_price }} birr
                          </span>
                        </span>
                      </div>

                      <div
                        class="grid grid-cols-2 md:flex-row items-center gap-4 mt-2"
                      >
                        <button
                          v-if="isScheduleExpired(schedule)"
                          @click="buyTicket(movie.movie_id)"
                          class="text-sm px-3 py-2 bg-teal-600 text-white rounded shadow-md w-full md:w-auto"
                        >
                          <span class="hidden md:block">Buy Ticket</span>
                          <span class="md:hidden">Buy</span>
                        </button>
                        <!-- Buy Ticket Button -->
                        <button
                          v-else
                          @click="buyTicket(movie.movie_id)"
                          disabled
                          class="text-sm px-3 py-2 cursor-not-allowed bg-gray-600 text-white rounded shadow-md w-full md:w-auto"
                        >
                          <span class="hidden md:block">Buy Ticket</span>
                          <span class="md:hidden">Buy</span>
                        </button>

                        <!-- Rate Button -->
                        <!-- @click="openRatingModal(movie.movie_id)" -->
                        <button
                          @click="openRatingModal(movie.movie_id)"
                          class="text-sm px-3 py-2 bg-teal-600 text-white rounded shadow-md w-full md:w-auto"
                        >
                          <span v-if="!israte">Rate</span>
                          <span v-else>x</span>
                        </button>
                        <div v-if="israte && currentMovieId === movie.movie_id">
                          <div class="flex justify-center space-x-2">
                            <div
                              v-for="star in 5"
                              :key="star"
                              @click="selectedRating = star"
                              class="cursor-pointer"
                              :class="selectClass"
                            >
                              <component
                                :is="
                                  star <= selectedRating
                                    ? StarSolid
                                    : StarOutline
                                "
                                class="w-8 h-8 text-yellow-500"
                                :class="selectClass"
                              />
                            </div>
                            <div class="flex justify-end space-x-2">
                              <button
                                @click="submitRating(currentMovieId)"
                                class="px-2 py-1 bg-green-600 dark:bg-green-800 text-white rounded-lg"
                              >
                                Submit
                              </button>
                              <button
                                @click="
                                  israte = false;
                                  currentMovieId = null;
                                "
                                class="px-2 py-1 bg-red-600 dark:bg-red-800 text-white rounded-lg"
                              >
                                Cancel
                              </button>
                            </div>
                          </div>
                        </div>
                      </div>
                    </li>
                  </ul>
                </div>
                <div v-else>
                  <span
                    class="flex items-center text-sm md:text-xl shadow-md text-teal-400 mx-4"
                    >Not scheduled yet please wait for schedule date</span
                  >
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
            Previous
          </button>
          <span class="font-semibold" :class="selectClass">
            {{
              noMoviesFound
                ? "Page 1 of total 1"
                : `Page ${currentPage} of ${movie.totalPages}`
            }}
          </span>
          <button
            @click="goToNextPage"
            :disabled="currentPage === movie.totalPages || noMoviesFound"
            class="flex items-center px-4 py-2 rounded-md shadow-md transition-all bg-teal-500 text-white font-medium hover:bg-teal-600 disabled:bg-gray-300 disabled:text-gray-500 disabled:cursor-not-allowed"
          >
            <ArrowRightIcon class="h-5 w-5 ml-2" />
            Next
          </button>
        </div>
      </template>
    </div>
  </div>
</template>
