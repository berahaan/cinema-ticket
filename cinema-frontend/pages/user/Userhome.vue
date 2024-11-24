<template><Userhome /></template>
<script setup>
definePageMeta({
  middleware: "auth",
});
import Userhome from "~/components/regular/Userhome.vue";
</script>
