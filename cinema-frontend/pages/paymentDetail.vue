<template>
  <Thankyou />
</template>

<script setup>
definePageMeta({
  middleware: "auth",
  layout: "user",
});
import Thankyou from "~/components/Paymentdetail.vue";
</script>
