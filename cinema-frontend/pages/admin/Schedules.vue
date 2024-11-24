<template>
  <Schedules />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import Schedules from "~/components/admin/Schedules.vue";
</script>
