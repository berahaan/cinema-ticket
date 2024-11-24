import { defineNuxtPlugin } from "#app";
import { useToast } from "vue-toastification";
import { onError } from "@apollo/client/link/error";
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
  ApolloLink,
} from "@apollo/client/core";
const toast = useToast();
const httpLink = createHttpLink({
  uri: "http://localhost:8080/v1/graphql",
});

const authLink = new ApolloLink((operation, forward) => {
  const accessToken = localStorage.getItem("accessToken");
  if (accessToken) {
    operation.setContext({
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
  }
  return forward(operation);
});

const getAccessToken = async () => {
  const refreshToken = localStorage.getItem("refreshToken");
  if (!refreshToken) {
    console.log("No refresh token available.");
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
    return null; // Return null if no refresh token
  }

  try {
    console.log("Sending refreshToken to backend", refreshToken);
    const response = await fetch("http://localhost:4000/refreshToken", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ refreshToken }),
    });

    if (!response.ok) {
      if (response.status === 403) {
        console.error("Invalid refresh token. Logging out...");
        localStorage.removeItem("accessToken");
        localStorage.removeItem("refreshToken");

        return navigateTo("/login");
      }
      throw new Error(
        "Failed to refresh access token please re login again Thanks "
      );
    }

    const data = await response.json();
    console.log("Here is the response from backend:", data.accessToken);
    localStorage.setItem("accessToken", data.accessToken);
    return data.accessToken; // Return the new access token
  } catch (error) {
    console.error("Error while sending to backend:", error.message);
    navigateTo("/login"); // Optionally handle navigation
    return null; // Return null in case of error
  }
};
const errorLink = onError(
  ({ graphQLErrors, networkError, operation, forward }) => {
    if (graphQLErrors) {
      for (const err of graphQLErrors) {
        // Intercept only invalid JWT errors
        if (err.extensions.code === "invalid-jwt") {
          localStorage.removeItem("accessToken");
          toast.warning(
            "Ur session is expired please login again to access .."
          );
          navigateTo("/login");
          // return getAccessToken().then((newAccessToken) => {
          //   if (newAccessToken) {
          //     operation.setContext({
          //       headers: {
          //         ...operation.getContext().headers,
          //         Authorization: `Bearer ${newAccessToken}`,
          //       },
          //     });
          //     return forward(operation); // Retry the request with a new token
          //   } else {
          //     console.log("Token refresh failed; redirecting to login.");
          //     navigateTo("/login");
          //   }
          // });
        }
      }
    }

    if (networkError) {
      console.error("Network error:", networkError);
    }

    // Allow other errors to propagate so the component can handle them
  }
);

const apolloClient = new ApolloClient({
  link: ApolloLink.from([errorLink, authLink, httpLink]),
  cache: new InMemoryCache(),
});

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.provide("apollo", apolloClient);
});
