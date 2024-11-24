<script setup>
import { ref } from "vue";
import {
  MenuIcon,
  SunIcon,
  MoonIcon,
  ChevronDownIcon,
} from "@heroicons/vue/solid";
import { computed, onMounted, watch } from "vue";
import cinemaImage from "../image/cinema.jpg";
const colorStore = useColorModeStore();
const { selectClass } = useThemeColor();
const { Logout } = useLogout();
const isDark = ref(true);
const isMobileMenuOpen = ref(false);
const activeMenu = ref(null);
const isMoviesOpen = ref(false);
const isScheduleOpen = ref(false);
const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};
onMounted(() => {
  if (import.meta.client) {
    if (isDark.value) {
      document.body.classList.add("dark");
      document.body.classList.remove("light");
    } else {
      document.body.classList.add("light");
      document.body.classList.remove("dark");
    }
  }
});

const toggleMenu = (menuName) => {
  if (menuName === "movies") {
    activeMenu.value = activeMenu.value === "movies" ? null : "movies";
    isMoviesOpen.value = !isMoviesOpen.value;
    isScheduleOpen.value = false; // close other menus
  } else if (menuName === "schedule") {
    activeMenu.value = activeMenu.value === "schedule" ? null : "schedule";
    isScheduleOpen.value = !isScheduleOpen.value;
    isMoviesOpen.value = false; // close other menus
  }
};

const toggleColor = () => {
  colorStore.toggleColorMode();
  // colorStore.colorMode = colorStore.colorMode === "dark" ? "light" : "dark";
};

// Load initial color mode only on client side
onMounted(() => {
  colorStore.loadColorMode();
  applyBodyClass(colorStore.colorMode);
});

watch(
  () => colorStore.colorMode,
  (newMode) => {
    applyBodyClass(newMode);
  }
);

function applyBodyClass(mode) {
  if (import.meta.client) {
    document.body.classList.toggle("bg-gray-900", mode === "dark");
    document.body.classList.toggle("text-white", mode === "dark");
    document.body.classList.toggle("bg-gray-100", mode === "light");
    document.body.classList.toggle("text-black", mode === "light");
  }
}
const colorModeClasses = computed(() =>
  colorStore.colorMode === "dark"
    ? "bg-gray-900 text-white"
    : "bg-gray-100 text-black"
);
</script>

<style scoped>
.rotate-180 {
  transform: rotate(180deg);
}
body.bg-gray-900,
body.text-white,
body.bg-gray-100,
body.text-black {
  transition: background-color 0.3s ease, color 0.3s ease;
}
</style>

<template>
  <header
    class="w-full flex fixed top-0 right-2 left-0 z-10 items-center justify-between p-4"
    :class="colorModeClasses"
  >
    <button @click="toggleMobileMenu" class="md:hidden mr-2">
      <MenuIcon class="w-6 h-6" />
    </button>

    <div class="flex items-center">
      <img
        :src="cinemaImage"
        alt="Cinema Image"
        class="w-10 h-10 object-cover rounded-full mr-4"
      />

      <nuxt-link to="/admin" class="text-xl font-semibold">
        Admin Dashboard
      </nuxt-link>
    </div>

    <nav class="hidden md:flex items-center space-x-8 relative">
      <!-- Movie Management -->
      <div class="relative">
        <button
          @click="toggleMenu('movies')"
          class="flex items-center justify-between transition-colors"
        >
          🎬 Movie
          <ChevronDownIcon
            class="w-6 h-6 transition-transform duration-200"
            :class="{ 'rotate-180': isMoviesOpen }"
          />
        </button>
        <div
          v-if="isMoviesOpen"
          class="absolute left-0 mt-2 w-40 shadow-lg border border-gray-300 rounded-lg py-2 z-50"
          :class="selectClass"
        >
          <nuxt-link
            to="/admin/addMovie"
            class="px-4 py-2 rounded-lg flex items-center space-x-2 relative group hover:text-blue-500"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-green-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4v16m8-8H4"
              />
            </svg>
            <span>Add Movie</span>

            <span
              class="absolute bottom-0 left-0 w-full h-[2px] bg-blue-500 scale-x-0 group-hover:scale-x-100 transition-transform duration-300"
            ></span>
          </nuxt-link>

          <nuxt-link
            to="/admin/Movielist"
            class="px-4 py-2 rounded-lg flex items-center space-x-2 relative group"
            :class="selectClass"
          >
            <!-- Film Icon -->
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-green-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
            <span>Movielist</span>
            <!-- Underline effect -->
            <span
              class="absolute bottom-0 left-0 w-full h-[2px] bg-blue-500 scale-x-0 group-hover:scale-x-100 transition-transform duration-300"
            ></span>
          </nuxt-link>
        </div>
      </div>

      <!-- Schedule Management -->
      <div class="relative">
        <button
          @click="toggleMenu('schedule')"
          class="flex items-center justify-between transition-colors hover:text-teal-300"
        >
          📅 Schedule
          <ChevronDownIcon
            class="w-6 h-6 transition-transform duration-200"
            :class="{ 'rotate-180': isScheduleOpen }"
          />
        </button>

        <div
          v-if="isScheduleOpen"
          class="absolute left-0 mt-2 w-48 shadow-lg border border-gray-300 rounded-lg py-2 z-50"
          :class="selectClass"
        >
          <nuxt-link
            to="/admin/Schedules"
            class="px-4 py-2 rounded-lg flex items-center space-x-2 relative group"
            :class="selectClass"
          >
            <!-- Calendar Plus Icon for Add Schedules -->
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-green-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4v16m8-8H4"
              />
            </svg>

            <span>Add Schedules</span>
            <!-- Underline effect -->
            <span
              class="absolute bottom-0 left-0 w-full h-[2px] bg-blue-500 scale-x-0 group-hover:scale-x-100 transition-transform duration-300"
            ></span>
          </nuxt-link>

          <!-- Update Schedules Link -->
          <nuxt-link
            to="/admin/MovieSchedule"
            class="px-4 py-2 rounded-lg flex items-center space-x-2 relative group"
            :class="selectClass"
          >
            <!-- Calendar Edit Icon for Update Schedules -->
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-green-600"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4v16m8-8H4m14-6l4 4m-4-4l-4 4"
              />
            </svg>
            <span>Update Schedules</span>
            <!-- Underline effect -->
            <span
              class="absolute bottom-0 left-0 w-full h-[2px] bg-blue-500 scale-x-0 group-hover:scale-x-100 transition-transform duration-300"
            ></span>
          </nuxt-link>
        </div>
      </div>

      <!-- Star Management -->
      <nuxt-link
        to="/admin/Star"
        class="block px-4 py-2 transition-colors hover:text-teal-300"
      >
        ⭐ Star Management
      </nuxt-link>

      <!-- Genre Management -->
      <nuxt-link
        to="/admin/Genre"
        class="block px-4 py-2 transition-colors hover:text-teal-300"
      >
        🎭 Genres Management
      </nuxt-link>

      <!-- Director Management -->
      <nuxt-link
        to="/admin/AddDirector"
        class="block px-4 py-2 transition-colors hover:text-teal-300"
      >
        🎥 Director
      </nuxt-link>
    </nav>

    <div
      v-if="isMobileMenuOpen"
      :class="selectClass"
      class="absolute top-16 left-0 w-full text-white p-4 md:hidden rounded-lg shadow-lg"
    >
      <nav class="flex flex-col space-y-4">
        <div class="relative">
          <button
            @click="toggleMenu('movies')"
            :class="selectClass"
            class="flex items-center justify-between w-full px-4 py-2 text-lg font-semibold rounded-md transition-colors hover:text-teal-600"
          >
            <span>🎬 Movie Management</span>
            <span
              :class="`transform ${
                activeMenu === 'movies' ? 'rotate-90' : 'rotate-0'
              } transition-transform ${selectClass}`"
            >
              ➤
            </span>
          </button>
          <div v-if="activeMenu === 'movies'" class="ml-6 mt-2 space-y-1">
            <nuxt-link
              to="/admin/Addmovie"
              class="block px-2 py-1 text-sm rounded-lg hover:text-teal-600 transition-colors"
              :class="selectClass"
              >Add Movies</nuxt-link
            >
            <nuxt-link
              to="/admin/Movielist"
              class="block px-2 py-1 text-sm rounded-lg hover:text-teal-600 transition-colors"
              :class="selectClass"
              >Movie List</nuxt-link
            >
          </div>
        </div>

        <!-- Schedule Management -->
        <div class="relative">
          <button
            @click="toggleMenu('schedule')"
            class="flex items-center justify-between w-full px-4 py-2 text-lg font-semibold rounded-md transition-colors hover:text-teal-600"
            :class="selectClass"
          >
            <span>📅 Schedule Management</span>
            <span
              :class="`transform ${
                activeMenu === 'schedule' ? 'rotate-90' : 'rotate-0'
              } transition-transform  ${selectClass}`"
            >
              ➤
            </span>
          </button>
          <div
            v-if="activeMenu === 'schedule'"
            class="ml-6 mt-2 space-y-1"
            :class="selectClass"
          >
            <nuxt-link
              to="/admin/Schedules"
              class="block px-2 py-1 text-sm rounded-lg hover:text-teal-600 transition-colors"
              >Add Schedule</nuxt-link
            >
            <nuxt-link
              to="/admin/MovieSchedule"
              class="block px-2 py-1 text-sm rounded-lg hover:text-teal-600 transition-colors"
              >Update Schedule</nuxt-link
            >
          </div>
        </div>

        <!-- Star Management -->
        <nuxt-link
          to="/admin/Star"
          class="flex items-center px-4 py-2 text-lg font-semibold rounded-md transition-colors hover:text-teal-600"
          :class="selectClass"
        >
          Star Management
        </nuxt-link>

        <!-- Genre Management -->
        <nuxt-link
          to="/admin/Genre"
          class="flex items-center px-4 py-2 text-lg font-semibold rounded-md transition-colors hover:text-teal-600"
          :class="selectClass"
        >
          Genres
        </nuxt-link>

        <!-- Director Management -->
        <nuxt-link
          to="/admin/AddDirector"
          class="flex items-center px-4 py-2 text-lg font-semibold rounded-md transition-colors hover:text-teal-600"
          :class="selectClass"
        >
          Director Management
        </nuxt-link>
      </nav>
    </div>

    <!-- Logout Button (Always visible) -->
    <button @click="Logout">
      <div class="flex">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="size-6"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15"
          />
        </svg>
        <span> Logout</span>
      </div>
    </button>
    <nuxt-link to="/profile"> ⚙️ </nuxt-link>
    <button
      @click="toggleColor"
      class="p-2 rounded-full focus:outline-none transition-colors duration-300"
      :class="{
        'bg-gray-800': colorStore.colorMode === 'dark',
        'bg-yellow-500': colorStore.colorMode !== 'dark',
      }"
    >
      <MoonIcon
        v-if="colorStore.colorMode === 'dark'"
        class="w-6 h-6 text-white"
      />
      <SunIcon v-else class="w-6 h-6 text-white" />
    </button>
  </header>
</template>
