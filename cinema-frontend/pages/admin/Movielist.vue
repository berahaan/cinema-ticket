<template>
  <AllMovieParts />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import AllMovieParts from "~/components/admin/AllMovieParts.vue";
</script>
