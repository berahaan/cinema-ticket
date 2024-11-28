<template>
  <Thankyou />
</template>

<script setup>
import Thankyou from "~/components/Paymentdetail.vue";
import { jwtDecode } from "jwt-decode"; // Ensure correct import path for jwtDecode

// Determine layout dynamically
if (import.meta.env.SSR === false) {
  const token = localStorage.getItem("accessToken");
  if (token) {
    try {
      const decodedToken = jwtDecode(token);
      const defaultRole =
        decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
      const layout =
        defaultRole === "admin" || defaultRole === "managers"
          ? "custom"
          : "user";

      // Set layout dynamically
      definePageMeta({
        layout,
      });

      console.log("Here is the default role layout:", layout);
      console.log("Here is the default role:", defaultRole);
    } catch (error) {
      console.error("Failed to decode token:", error);
    }
  } else {
    // Fallback for missing token
    definePageMeta({
      layout: "user",
    });
  }
}
</script>
