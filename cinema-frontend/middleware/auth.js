import { useNuxtApp, navigateTo } from "#app";
import { VALIDATE_TOKEN } from "../components/graphql/mutations/auth.graphql";
import { useToast } from "vue-toastification";
export default defineNuxtRouteMiddleware(async (to, from) => {
  const { $apollo } = useNuxtApp();
  const toast = useToast();
  if (import.meta.client) {
    const token = localStorage.getItem("accessToken");

    if (!token) {
      toast.warning("You need to login first to accesss it Thank you", {
        position: "top-right",
        duration: 2000,
      });
      return navigateTo("/");
    }

    try {
      const { data } = await $apollo.mutate({
        mutation: VALIDATE_TOKEN,
        variables: { token },
      });
      const { isvalid } = data.authenthicate;
      if (!isvalid) {
        toast.warning(
          "Your session is not valid, you will be redirected to login."
        );
        return navigateTo("/auth/login");
      }
    } catch (error) {
      console.error("Error validating token:", error);
      toast.warning("Server is down now please try again in a few minutes ", {
        position: "top-right",
        duration: 2000,
      });
      return navigateTo("/auth/login");
    } finally {
      console.log("Finished validation.");
    }
  } else {
    console.log("Server side rendering middleware...");
  }
});
