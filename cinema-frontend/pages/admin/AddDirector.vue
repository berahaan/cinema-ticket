<template>
  <Directors />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "custom",
});
import Directors from "~/components/admin/Directors.vue";
</script>
