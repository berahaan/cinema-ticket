<template>
  <AdminBar />
  <div>
    <slot />
  </div>
<Footer/>
</template>
<script setup>
definePageMeta({
  middleware:'auth'
})
import AdminBar from "~/components/admin/AdminBar.vue";
import Footer from "~/components/Footer.vue";
</script>
