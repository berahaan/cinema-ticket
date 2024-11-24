<template>
  <AddMovies />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import AddMovies from "~/components/admin/AddMovies.vue";
</script>
