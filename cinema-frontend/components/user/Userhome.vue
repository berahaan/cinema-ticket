<script setup>
import { ref, computed, watch, onMounted } from "vue";
import {
  SunIcon,
  MoonIcon,
  FilmIcon,
  BookmarkIcon,
  UserIcon,
  MenuIcon,
  XIcon,
} from "@heroicons/vue/outline";

const colorStore = useColorModeStore();
const { Logout } = useLogout();

const colorModeClasses = computed(() =>
  colorStore.colorMode === "dark"
    ? "bg-gray-900 text-white shadow-lg"
    : "bg-white text-gray-900 shadow-md"
);

const isMobileMenuOpen = ref(false);

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};

function applyBodyClass(mode) {
  if (import.meta.client) {
    document.body.classList.toggle("bg-gray-900", mode === "dark");
    document.body.classList.toggle("text-white", mode === "dark");
    document.body.classList.toggle("bg-white", mode === "light");
    document.body.classList.toggle("text-gray-900", mode === "light");
  }
}

const toggleColor = () => {
  colorStore.toggleColorMode();
};

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
</script>

<template>
  <header
    class="py-4 fixed top-0 right-0 left-0 z-20 border-b"
    :class="colorModeClasses"
  >
    <div class="container mx-auto flex justify-between items-center px-4">
      <div class="flex items-center space-x-3">
        <nuxt-link to="/profile" class="flex items-center">
          <UserIcon class="w-7 h-7 text-blue-500" />
          Profile
        </nuxt-link>
        <h1 class="text-xl md:text-2xl font-bold tracking-wide">
          Cinema Ticket
        </h1>
      </div>

      <nav class="hidden md:flex space-x-8 text-sm font-medium">
        <ul class="flex items-center space-x-6">
          <li
            class="flex items-center space-x-2 hover:text-blue-500 transition"
          >
            <nuxt-link to="/user" class="hover:underline">Home</nuxt-link>
          </li>
          <li
            class="flex items-center space-x-2 hover:text-blue-500 transition"
          >
            <FilmIcon class="w-5 h-5" />
            <nuxt-link to="/user/Movielist" class="hover:underline"
              >Movielist</nuxt-link
            >
          </li>
          <li
            class="flex items-center space-x-2 hover:text-blue-500 transition"
          >
            <BookmarkIcon class="w-5 h-5" />
            <nuxt-link to="/user/Bookmark" class="hover:underline"
              >Bookmark</nuxt-link
            >
          </li>
          <li>
            <button
              @click="Logout"
              class="flex items-center space-x-2 hover:text-red-500 transition"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-5 h-5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15"
                />
              </svg>
              <span>Logout</span>
            </button>
          </li>
        </ul>
      </nav>
      <button
        class="md:hidden p-2 rounded focus:outline-none"
        @click="toggleMobileMenu"
      >
        <MenuIcon v-if="!isMobileMenuOpen" class="w-6 h-6" />
        <XIcon v-else class="w-6 h-6" />
      </button>

      <transition name="fade" class="items-end">
        <div
          v-if="isMobileMenuOpen"
          :class="colorModeClasses"
          class="sm:hidden fixed top-16 right-4 shadow-lg rounded-md w-48"
        >
          <ul class="flex flex-col space-y-2 p-4 text-lg font-medium">
            <li>
              <nuxt-link to="/user" class="hover:underline">Home</nuxt-link>
            </li>
            <li>
              <nuxt-link to="/user/Movielist" class="hover:text-blue-400"
                >Movielist</nuxt-link
              >
            </li>
            <li>
              <nuxt-link to="/user/Bookmark" class="hover:text-blue-400"
                >Bookmark</nuxt-link
              >
            </li>
            <li>
              <button @click="Logout" class="hover:text-red-400">Logout</button>
            </li>
          </ul>
        </div>
      </transition>

      <!-- Theme Toggle Button -->
      <button
        @click="toggleColor"
        class="hidden md:block p-2 rounded-full focus:outline-none transition-colors duration-300"
        :class="{
          'bg-gray-800 hover:bg-gray-700': colorStore.colorMode === 'dark',
          'bg-yellow-500 hover:bg-yellow-400': colorStore.colorMode !== 'dark',
        }"
      >
        <MoonIcon
          v-if="colorStore.colorMode === 'dark'"
          class="w-6 h-6 text-white"
        />
        <SunIcon v-else class="w-6 h-6 text-white" />
      </button>
    </div>
  </header>
</template>

<style scoped>
/* Smooth fade for mobile menu */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
