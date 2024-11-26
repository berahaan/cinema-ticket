<template>
  <header
    class="py-4 fixed top-0 right-0 left-0 z-20 shadow-lg bg-gradient-to-r from-black via-gray-800 to-black"
  >
    <div class="container mx-auto flex justify-between items-center">
      <!-- Logo or Title -->
      <h1 class="text-4xl font-bold text-yellow-400 flex items-center">
        🎥 <span class="ml-2 tracking-widest">Ethio Cinema</span>
      </h1>

      <!-- Navigation Menu -->
      <nav>
        <ul class="flex space-x-8 text-lg font-medium">
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
      </nav>
    </div>
  </header>
</template>
