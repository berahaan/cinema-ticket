<template>
  <Dashboard />
  <NuxtPage />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import Dashboard from "~/components/admin/Dashboard.vue";
</script>
