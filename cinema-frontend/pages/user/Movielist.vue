<template>
  <Movielist />
</template>
<script setup>
definePageMeta({
  middleware: "auth",
  layout: "user",
});
import Movielist from "~/components/regular/Movielist.vue";
</script>
