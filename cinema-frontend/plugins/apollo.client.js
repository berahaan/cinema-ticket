import { defineNuxtPlugin } from "#app";
import { useToast } from "vue-toastification";
import { onError } from "@apollo/client/link/error";
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
  ApolloLink,
} from "@apollo/client/core";

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig();
  const toast = useToast();

  const httpLink = createHttpLink({
    uri: config.public.graphqlEndpoint,
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

  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors) {
      for (const err of graphQLErrors) {
        // Intercept only invalid JWT errors
        if (err.extensions.code === "invalid-jwt") {
          localStorage.removeItem("accessToken");
          toast.warning(
            "Your session has expired. Please log in again to access.",
            {
              position: "top-right",
              timeout: 2000,
            }
          );
          navigateTo("/auth/login");
        }
      }
    }

    if (networkError) {
      toast.warning("Network error...", {
        position: "top-right",
        timeout: 2000,
      });
    }
  });

  const apolloClient = new ApolloClient({
    link: ApolloLink.from([errorLink, authLink, httpLink]),
    cache: new InMemoryCache(),
    defaultOptions: {
      query: {
        fetchPolicy: "network-only",
      },
      watchQuery: {
        fetchPolicy: "network-only",
      },
    },
  });
  nuxtApp.provide("apollo", apolloClient);
});
