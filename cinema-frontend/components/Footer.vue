<template>
    <footer class="text-white py-6 bg-gradient-to-r from-teal-600 via-gray-800 to-gray-400 shadow-lg">
      <div class="container mx-auto text-center" v-for="x in arr" v-if="re">
        <p class="text-sm">
          &copy; 2024 Ethio Cinema. Your Gateway to the Best Movies.
        </p>
        <div class="flex justify-center space-x-4 mt-2 ">
          <a href="#" class="text-gray-400 hover:text-white text-sm">Privacy Policy</a>
          <a href="#" class="text-gray-400 hover:text-white text-sm">Terms of Use</a>
          <a href="#" class="text-gray-400 hover:text-white text-sm">Contact Us</a>
        </div>
      </div>
    </footer>
</template>


<script setup>
import { ref } from 'vue';
const arr =ref([])
</script>