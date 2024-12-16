<script setup>
import { ref } from "vue";
import {
  MenuIcon,
  SunIcon,
  MoonIcon,
  XIcon,
  ChevronDoubleRightIcon,
  ChevronRightIcon,
} from "@heroicons/vue/solid";
import { computed, onMounted, watch } from "vue";
// PencilSquareIcon
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
    class="w-full flex fixed top-0 right-2 left-0 z-10 items-center justify-between p-2 border-b"
    :class="colorModeClasses"
  >
    <button @click="toggleMobileMenu" class="md:hidden mr-2">
      <!-- <MenuIcon class="w-6 h-6" /> -->
      <MenuIcon v-if="!isMobileMenuOpen" class="w-6 h-6" />
      <XIcon v-else class="w-6 h-6" />
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

    <nav class="hidden md:flex items-center relative font-medium space-x-2">
      <div class="group relative cursor-pointer py-2">
        <div class="flex items-center justify-between">
          <a
            class="menu-hover my-2 py-2 text-base font-medium lg:mx-4"
            onClick=""
          >
            Movies
          </a>
          <span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-6 w-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M19.5 8.25l-7.5 7.5-7.5-7.5"
              />
            </svg>
          </span>
        </div>
        <div
          :class="selectClass"
          class="invisible absolute flex flex-col py-1 px-4 shadow-xl group-hover:visible w-40 min-w-[160px]"
        >
          <nuxt-link
            to="/admin/addMovie"
            :class="selectClass"
            class="my-2 block border-b border-gray-100 py-1 font-semibold md:mx-2 hover:text-teal-400"
          >
            Add Movies
          </nuxt-link>
          <nuxt-link
            to="/admin/Movielist"
            :class="selectClass"
            class="my-2 block border-b border-gray-100 py-1 font-semibold md:mx-2 hover:text-teal-400"
          >
            MovieList
          </nuxt-link>
        </div>
      </div>

      <div class="group relative cursor-pointer py-2">
        <div class="flex items-center justify-between">
          <a
            class="menu-hover my-2 py-2 text-base font-medium lg:mx-4"
            onClick=""
          >
            📅 Schedule
          </a>
          <span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="h-6 w-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M19.5 8.25l-7.5 7.5-7.5-7.5"
              />
            </svg>
          </span>
        </div>
        <div
          :class="selectClass"
          class="invisible absolute flex flex-col py-1 px-4 shadow-xl group-hover:visible w-40 min-w-[180px]"
        >
          <nuxt-link
            to="/admin/MovieSchedule"
            :class="selectClass"
            class="my-2 block border-b border-gray-100 py-1 font-semibold md:mx-2 hover:text-teal-400"
          >
            Update Schedule
          </nuxt-link>
          <nuxt-link
            to="/admin/Schedules"
            :class="selectClass"
            class="my-2 block border-b border-gray-100 py-1 font-semibold md:mx-2 hover:text-teal-400"
          >
            Add schedule
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
      class="sm:hidden fixed top-16 left-0 shadow-lg rounded-md w-60"
    >
      <nav class="flex flex-col items-start space-y-2 p-4 text-lg font-medium">
        <div class="relative">
          <button
            @click="toggleMenu('movies')"
            :class="selectClass"
            class="flex items-center justify-between w-full px-4 py-2 text-lg font-semibold rounded-md transition-colors hover:text-teal-600"
          >
            <span>Movie </span>
            <span
              :class="`transform ${
                activeMenu === 'movies' ? 'rotate-90' : 'rotate-0'
              } transition-transform ${selectClass}`"
            >
              <ChevronRightIcon class="w-6 h-6" />
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
            <span> Schedule</span>
            <span
              :class="`transform ${
                activeMenu === 'schedule' ? 'rotate-90' : 'rotate-0'
              } transition-transform  ${selectClass}`"
            >
              <ChevronRightIcon class="w-6 h-6" />
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
          Director
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
