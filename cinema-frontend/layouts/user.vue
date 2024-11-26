<template>
  <Userhome />
  <div>
    <slot />
  </div>
  <Footer />
</template>
<script setup>
import Userhome from "~/components/user/Userhome.vue";
import Footer from "~/components/Footer.vue";
</script>
