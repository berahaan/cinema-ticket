<template>
  <ALlmovieSchedule />
</template>

<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import ALlmovieSchedule from "~/components/admin/ALlmovieSchedule.vue";
</script>
