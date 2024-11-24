<template>
  <ticket />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import ticket from "~/components/ticket.vue";
</script>
