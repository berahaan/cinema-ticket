<template>
  <genres />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout:'custom'
});
import genres from "~/components/admin/genres.vue";
</script>
