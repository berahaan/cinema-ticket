<template><Userhome /></template>
<script setup>
definePageMeta({
  middleware: "auth",
});
import Userhome from "~/components/user/Userhome.vue";
</script>
