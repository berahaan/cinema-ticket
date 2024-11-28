<template>
  <detailMovies />
</template>

<script setup>
definePageMeta({
  middleware: "auth",
  layout: "user",
});
import detailMovies from "~/components/user/detailMovies.vue";
</script>
