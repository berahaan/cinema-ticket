import { gql } from "@apollo/client/core";
import { useNuxtApp, navigateTo } from "#app";
import { useToast } from "vue-toastification";
export default defineNuxtRouteMiddleware(async (to, from) => {
  const { $apollo } = useNuxtApp();
  const color = useColorModeStore();
  const toast = useToast();
  // Check if we're on the client side to access localStorage
  if (import.meta.client) {
    const token = localStorage.getItem("accessToken");
    console.log("Here is the auth.js authentication token:", token);

    // If no token is found, redirect to the login page
    if (!token) {
      toast.warning("You need to login first to access site Thank you", {
        position: "top-right",
        duration: 2000,
      });
      return navigateTo("/");
    }
    console.log("Starting token validation...");

    try {
      const VALIDATE_TOKEN = gql`
        mutation ($token: String!) {
          validateToken(token: $token) {
            isvalid
          }
        }
      `;

      // Execute the mutation to validate the token
      const { data } = await $apollo.mutate({
        mutation: VALIDATE_TOKEN,
        variables: { token },
      });

      const { isvalid } = data.validateToken;
      console.log("Backend validation response:", isvalid);

      if (!isvalid) {
        toast.warning("Your session is not valid, redirecting to login.");
        return navigateTo("/login");
      }
    } catch (error) {
      console.error("Error validating token:", error);
      toast.warning("Server is down now please try again in a few minutes ", {
        position: "top-right",
        duration: 2000,
      });
      return navigateTo("/login");
    } finally {
      console.log("Finished token validation.");
    }
  } else {
    console.log("Server side rendering ");
    // return navigateTo("/loading");
  }
});
