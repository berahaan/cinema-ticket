<template>
  <Stars />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import Stars from "~/components/admin/Stars.vue";
</script>
