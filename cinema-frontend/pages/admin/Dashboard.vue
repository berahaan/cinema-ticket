<template>
  <div class="admin-dashboard">
    <h1>Admin Dashboard</h1>
    <!-- Admin-specific content goes here -->
  </div>
</template>
<script setup>
definePageMeta({
  middleware: "auth",
});
</script>
<style scoped></style>
