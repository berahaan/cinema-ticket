<template>
  <HomeBar />
  <div>
    <slot />
  </div>
</template>
<script setup>
import HomeBar from "~/components/HomeBar.vue";
</script>
