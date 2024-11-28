<template>
  <AddSchedule />
</template>
<script setup>
definePageMeta({
  layout: "custom",
  middleware: "auth",
});
import AddSchedule from "~/components/admin/AddSchedule.vue";
</script>
