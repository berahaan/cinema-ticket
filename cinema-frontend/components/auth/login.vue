<script setup>
import { ref } from "vue";
import { useToast } from "vue-toastification";
import { Form, Field, ErrorMessage } from "vee-validate";
import { MailIcon, EyeIcon, EyeOffIcon } from "@heroicons/vue/solid";
import { useRouter } from "vue-router";
import { jwtDecode } from "jwt-decode";
import { useNuxtApp } from "#app";
import { LOGIN } from "../graphql/mutations/LOGIN.graphql";
const toast = useToast();
const loading = ref(false);
const showPassword = ref(false);
const { $apollo } = useNuxtApp();
const email = ref("");
const password = ref("");
const errorMessage = ref("");
const router = useRouter();
const validateEmail = (value) => {
  const regex = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i;
  if (!value) return "Email is required";
  if (!regex.test(value)) {
    return "This field must be a valid email";
  }
  return true;
};
const validatePassword = (value) => {
  if (!value) return "Password is required";
  return true;
};

const login = async () => {
  loading.value = true;

  try {
    const response = await $apollo.mutate({
      mutation: LOGIN,
      variables: { email: email.value, password: password.value },
    });

    const { accessToken } = response.data.login;

    if (accessToken) {
      localStorage.setItem("accessToken", accessToken);
      const decodedToken = jwtDecode(accessToken);
      const defaultRole =
        decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
      // Navigate based on the default role

      if (defaultRole === "admin" || defaultRole === "managers") {
        router.push("/admin");
      } else if (defaultRole === "regular") {
        router.push("/user");
      } else {
        router.push("/");
      }
    } else {
      router.push("/");
      errorMessage.value = "Login failed: No tokens received";
    }
  } catch (error) {
    if (error.graphQLErrors && error.graphQLErrors[0]) {
      errorMessage.value = error.graphQLErrors[0].message;
      toast.warning(errorMessage.value, {
        position: "top-right",
        duration: 2000,
      });
      setTimeout(() => {
        loading.value = false;
      }, 500);
    } else {
      errorMessage.value = "Login failed: Unknown error";
    }
  } finally {
    setTimeout(() => {
      errorMessage.value = "";
      loading.value = false;
    }, 3000);
  }
};
</script>
<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gradient-to-r from-gray-100 via-gray-200 to-gray-300"
  >
    <div class="bg-white shadow-2xl rounded-lg p-8 max-w-md w-full">
      <!-- Heading -->
      <h1 class="text-3xl font-extrabold text-center mb-4 text-gray-700">
        Welcome Back!
      </h1>
      <h2 class="text-lg font-medium text-center mb-8 text-gray-600">
        Login to your account and enjoy your experience.
      </h2>

      <Form @submit="login" class="space-y-6">
        <!-- Email Input -->
        <div class="relative">
          <label for="email" class="block text-sm font-medium text-gray-700">
            Email Address
          </label>
          <div class="relative mt-1">
            <MailIcon
              class="w-5 h-5 text-gray-400 absolute left-3 top-1/2 transform -translate-y-1/2 pointer-events-none"
            />
            <Field
              id="email"
              name="email"
              type="email"
              v-model="email"
              placeholder="Enter your email"
              class="pl-10 pr-3 py-2 w-full border border-gray-300 rounded-lg shadow-sm text-gray-700 focus:ring-blue-500 focus:border-blue-500 focus:outline-none transition duration-200"
              :rules="validateEmail"
            />
          </div>
          <ErrorMessage name="email" class="text-sm text-red-600 mt-1" />
        </div>

        <!-- Password Input -->
        <div class="relative">
          <label for="password" class="block text-sm font-medium text-gray-700">
            Password
          </label>
          <div class="relative mt-1">
            <Field
              id="password"
              name="password"
              :type="showPassword ? 'text' : 'password'"
              v-model="password"
              placeholder="Enter your password"
              class="pl-3 pr-10 py-2 w-full border border-gray-300 rounded-lg shadow-sm text-gray-700 focus:ring-blue-500 focus:border-blue-500 focus:outline-none transition duration-200"
              :rules="validatePassword"
            />
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute right-3 top-1/2 transform -translate-y-1/2 focus:outline-none"
            >
              <EyeIcon v-if="!showPassword" class="w-5 h-5 text-gray-400" />
              <EyeOffIcon v-else class="w-5 h-5 text-gray-400" />
            </button>
            <ErrorMessage name="password" class="text-sm text-red-600 mt-1" />
          </div>
        </div>

        <!-- Submit Button -->
        <button
          type="submit"
          class="w-full py-3 bg-gradient-to-r from-teal-500 to-purple-600 text-white font-bold rounded-lg shadow-md hover:from-teal-600 hover:to-purple-700 focus:outline-none focus:ring-4 focus:ring-teal-400 focus:ring-opacity-50 transition duration-200"
        >
          <span v-if="loading" class="flex items-center justify-center">
            <svg
              class="animate-spin h-5 w-5 mr-2 text-white"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <circle
                class="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                stroke-width="4"
              ></circle>
              <path
                class="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
              ></path>
            </svg>
            please wait...
          </span>
          <span v-else>Login</span>
        </button>

        <!-- Footer Links -->
        <div class="flex items-center justify-center mt-6">
          <p class="text-sm text-gray-500">
            Don't have an account?
            <nuxt-link
              to="/auth/signup"
              class="ml-1 text-sm font-bold text-teal-500 hover:text-teal-600 transition-colors duration-200"
            >
              Sign Up
            </nuxt-link>
          </p>
        </div>
      </Form>
    </div>
  </div>
</template>
