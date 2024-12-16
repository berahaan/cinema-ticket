<script setup>
import { onMounted } from "vue";
import Loading from "./Loading.vue";
import { ArrowLeftIcon } from "@heroicons/vue/solid";
const { formatScheduleDateTime } = useFormatSchedule();
const { selectClass } = useThemeColor();
const { goBack } = useGobackArrow();
const bookmark = useBookmarkStore();
const { removeBookmarks, getBookmarks, loading } = useBookmarks();

onMounted(async () => {
  await getBookmarks();
});
</script>
<template>
  <div class="container mx-auto p-6 rounded-lg shadow-lg">
    <h2 class="text-3xl mb-6 text-center text-teal-400">
      Your Bookmarked Movies
    </h2>
    <button
      @click="goBack"
      class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200"
    >
      <ArrowLeftIcon class="h-5 w-5" />
      <span>Go Back</span>
    </button>
    <div v-if="loading" class="text-center">
      <Loading />
    </div>

    <div v-if="bookmark.bookmarks?.length > 0">
      <ul class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <li
          v-for="bookmark in bookmark.bookmarks"
          :key="bookmark.movie_id"
          class="border rounded-lg shadow-lg p-5 flex flex-col"
          :class="selectClass"
        >
          <!-- Movie Info Container -->
          <div class="flex flex-col lg:flex-row lg:space-x-4">
            <!-- Movie Poster -->
            <div class="w-full lg:w-1/2 mb-4 lg:mb-0">
              <img
                :src="bookmark.movie.featured_images"
                alt="Movie Poster"
                class="w-full h-48 object-cover rounded-lg shadow-md cursor-pointer"
              />
            </div>
            <div class="flex-1 flex flex-col space-y-3">
              <div>
                <h3 class="text-lg" :class="selectClass">
                  {{ bookmark.movie.title }}
                </h3>
                <p class="text-sm">
                  <span class="font-medium">Duration:</span>
                  {{ bookmark.movie.duration }} minutes
                </p>
                <p class="text-sm" :class="selectClass">
                  <span class="font-medium">Director:</span>
                  {{ bookmark.movie.director.name }}
                </p>
                <p class="text-sm" :class="selectClass">
                  <span class="font-medium">Genres:</span>
                  {{
                    bookmark.movie.movie_genres
                      .map((g) => g.genre.name)
                      .join(", ")
                  }}
                </p>
                <p class="text-sm">
                  <span class="font-medium">Stars:</span>
                  {{
                    bookmark.movie.movie_stars
                      .map((s) => s.star.name)
                      .join(", ")
                  }}
                </p>
              </div>

              <!-- Schedule List -->
            </div>
          </div>
          <div class="p-4 rounded-lg shadow-inner" :class="selectClass">
            <h4 class="text-base font-semibold text-blue-800 mb-2">
              Scheduled Times
            </h4>
            <ul class="space-y-2">
              <li
                v-for="schedule in bookmark.movie.movie_schedules"
                :key="schedule.start_time"
                class="rounded-lg p-3 shadow-md"
                :class="selectClass"
              >
                <p class="text-sm">
                  <span class="font-semibold" :class="selectClass">
                    {{
                      formatScheduleDateTime(
                        schedule.schedule_date,
                        schedule.start_time,
                        schedule.end_time
                      )
                    }}
                  </span>
                  <span class="text-sm text-gray-500 ml-2">
                    | Hall: {{ schedule.cinema_hall }} | Price:
                    <span class="font-semibold text-green-600"
                      >{{ schedule.ticket_price }} birr</span
                    >
                  </span>
                </p>
              </li>
            </ul>
          </div>
          <button
            @click="removeBookmarks(bookmark.movie_id)"
            class="px-4 py-2 bg-red-600 dark:bg-red-800 text-white rounded-lg mx-6"
          >
            Remove Movie
          </button>
        </li>
      </ul>
    </div>
    <div v-else class="mx-10 mb-96">
      Unfortunately, no bookmarks found at the moment Please try to add Movies
      here...
    </div>
  </div>
</template>
