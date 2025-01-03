<script setup>
import { ref } from "vue";
import { MenuIcon, XIcon } from "@heroicons/vue/outline";
const isMobileMenuOpen = ref(false);
const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};
</script>

<template>
  <!-- bg-gradient-to-r from-black via-gray-800 to-black -->
  <header class="py-4 fixed top-0 right-0 left-0 z-20 shadow-lg bg-gray-900">
    <div class="container mx-auto flex justify-between items-center">
      <!-- Logo or Title -->
      <h1 class="text-4xl font-bold text-yellow-400 flex items-center">
        🎥
        <span class="text-xl md:text-3xl ml-2 tracking-widest">
          <nuxt-link to="/">Ethio Cinema</nuxt-link></span
        >
      </h1>
      <nav>
        <ul class="text-lg font-medium hidden md:flex">
          <li>
            <nuxt-link
              to="/"
              class="transition duration-200 text-gray-300 hover:text-yellow-400"
            >
              Home
            </nuxt-link>
          </li>
          <li>
            <nuxt-link
              to="/auth/login"
              class="text-gray-300 px-4 py-2 transition-all duration-200 hover:text-white hover:border hover:border-yellow-400 rounded-full"
            >
              Login
            </nuxt-link>
          </li>
          <li>
            <nuxt-link
              to="/auth/signup"
              class="text-gray-300 px-4 py-2 transition-all duration-200 hover:text-white hover:border hover:border-yellow-400 rounded-full"
            >
              Signup
            </nuxt-link>
          </li>
        </ul>
        <button
          class="md:hidden p-2 rounded focus:outline-none"
          @click="toggleMobileMenu"
        >
          <MenuIcon
            v-if="!isMobileMenuOpen"
            class="w-6 h-6 bg-gray-800 text-white"
          />
          <XIcon v-else class="w-6 h-6 bg-gray-800 text-white" />
        </button>
        <transition name="fade">
          <div
            v-if="isMobileMenuOpen"
            class="sm:hidden fixed top-14 right-1 shadow-lg rounded-md w-48 bg-gray-900"
          >
            <ul class="flex flex-col space-y-2 p-4 text-lg font-medium">
              <li>
                <nuxt-link
                  to="/"
                  class="text-gray-300 px-4 py-2 transition-all duration-200 hover:text-white hover:border hover:border-yellow-400 rounded-full"
                >
                  Home
                </nuxt-link>
              </li>
              <li>
                <nuxt-link
                  to="/auth/login"
                  class="text-gray-300 px-4 py-2 transition-all duration-200 hover:text-white hover:border hover:border-yellow-400 rounded-full"
                >
                  Login
                </nuxt-link>
              </li>
              <li>
                <nuxt-link
                  to="/auth/signup"
                  class="text-gray-300 px-4 py-2 transition-all duration-200 hover:text-white hover:border hover:border-yellow-400 rounded-full"
                >
                  Signup
                </nuxt-link>
              </li>
            </ul>
          </div>
        </transition>
      </nav>
    </div>
  </header>
</template>
