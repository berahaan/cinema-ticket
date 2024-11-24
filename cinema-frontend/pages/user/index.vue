<template>
  <!-- <Movielist /> -->
   <Home/>
</template>
<script setup>
definePageMeta({
  middleware:'auth',
  layout: "user",
});
// import Movielist from "~/components/regular/Movielist.vue";
import Home from '~/components/regular/Home.vue';
</script>
