<script setup>
import { useNuxtApp } from "#app";
import * as Yup from "yup";
import { useToast } from "vue-toastification";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { SIGNUP } from "../graphql/mutations/SIGNUP.graphql";
import { MailIcon, EyeIcon, EyeOffIcon } from "@heroicons/vue/solid";
import { Field, Form, ErrorMessage } from "vee-validate";
const loading = ref(false);
const toast = useToast();
const { $apollo } = useNuxtApp();
const router = useRouter();
const firstname = ref("");
const lastname = ref("");
const email = ref("");
const password = ref("");
const confirmPassword = ref("");
const errorMessage = ref("");
const showPassword = ref(false);
const togglePassword = () => {
  showPassword.value = !showPassword.value;
};
const signupValidationSchema = Yup.object({
  firstname: Yup.string()
    .required("First name is required")
    .min(2, "First name must be at least 2 characters"),
  lastname: Yup.string()
    .required("Last name is required")
    .min(2, "Last name must be at least 2 characters"),
  email: Yup.string()
    .email("Invalid email address")
    .required("Email is required"),
  password: Yup.string()
    .required("Password is required")
    .min(8, "Password must be at least 8 characters")
    .matches(/[A-Z]/, "Password must contain at least one uppercase letter")
    .matches(/[a-z]/, "Password must contain at least one lowercase letter")
    .matches(/[0-9]/, "Password must contain at least one number")
    .matches(
      /[@$!%*?&#]/,
      "Password must contain at least one special character"
    ),
  confirmPassword: Yup.string()
    .oneOf([Yup.ref("password"), null], "Passwords is not match")
    .required("Please confirm your password"),
});

const signup = async () => {
  loading.value = true;
  errorMessage.value = "";

  const payload = {
    firstname: firstname.value,
    lastname: lastname.value,
    email: email.value,
    password: password.value,
  };

  try {
    const response = await $apollo.mutate({
      mutation: SIGNUP,
      variables: payload,
    });

    // Check if signup was successful
    const mess = response.data.signup.message;
    if (mess === "account successfully created!") {
      errorMessage.value = mess;
      toast.success(mess, {
        position: "top-right",
        duration: 2000,
      });
      setTimeout(() => {
        loading.value = false;
        router.push("/auth/login");
      }, 2000);
    } else {
      errorMessage.value = mess;
      toast.warning(mess, {
        position: "top-right",
        duration: 2000,
      });
      loading.value = false;
    }
  } catch (error) {
    errorMessage.value = "Signup failed: " + error.message;
    toast.error(errorMessage.value, {
      position: "top-right",
      duration: 2000,
    });
    loading.value = false;
  }
};
</script>

<style scoped>
input::placeholder {
  color: #a0aec0;
}
</style>

<template>
  <div
    class="min-h-screen flex items-center justify-center bg-gradient-to-br from-indigo-100 to-purple-100 py-10"
  >
    <div class="bg-white shadow-xl rounded-lg p-8 max-w-lg w-full">
      <!-- Header Section -->
      <h2
        class="md:text-4xl text-xl font-extrabold text-center mb-4 text-gray-800"
      >
        Join Us Today!
      </h2>
      <p class="text-center text-lg text-gray-600 mb-8">
        Unlock exclusive features and take your movie experience to the next
        level. <br />
        <span class="font-semibold text-purple-600"
          >It's free and takes less than a minute!</span
        >
      </p>

      <!-- Form Section -->
      <Form
        @submit="signup"
        class="space-y-6"
        autocomplete="on"
        :validation-schema="signupValidationSchema"
      >
        <!-- First Name Field -->
        <div>
          <label
            for="firstname"
            class="block text-lg font-semibold text-gray-700"
            >First Name</label
          >
          <Field
            name="firstname"
            type="text"
            v-model="firstname"
            placeholder="Enter your first name"
            class="mt-1 block w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-purple-500 focus:border-purple-500"
          />
          <ErrorMessage name="firstname" class="text-red-600 text-sm mt-1" />
        </div>

        <!-- Last Name Field -->
        <div>
          <label
            for="lastname"
            class="block text-lg font-semibold text-gray-700"
            >Last Name</label
          >
          <Field
            name="lastname"
            type="text"
            v-model="lastname"
            placeholder="Enter your last name"
            class="mt-1 block w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-purple-500 focus:border-purple-500"
          />
          <ErrorMessage name="lastname" class="text-red-600 text-sm mt-1" />
        </div>

        <!-- Email Field with Icon -->
        <div class="relative">
          <label for="email" class="block text-lg font-semibold text-gray-700">
            Email
          </label>
          <Field
            name="email"
            type="email"
            v-model="email"
            placeholder="Enter your email"
            class="mt-1 block w-full p-3 pl-10 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-purple-500 focus:border-purple-500"
          />
          <MailIcon class="absolute left-3 top-10 w-5 h-5 text-gray-400" />
          <ErrorMessage name="email" class="text-red-600 text-sm mt-1" />
        </div>

        <!-- Password Field with Toggle Icon -->
        <div class="relative">
          <label
            for="password"
            class="block text-lg font-semibold text-gray-700"
            >Password</label
          >
          <Field
            name="password"
            id="password"
            autocomplete="new-password"
            :type="showPassword ? 'text' : 'password'"
            v-model="password"
            placeholder="Create password"
            class="mt-1 block w-full p-3 pl-10 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-purple-500 focus:border-purple-500"
          />
          <span
            @click="togglePassword"
            class="absolute right-3 top-12 cursor-pointer"
          >
            <EyeIcon
              v-if="showPassword"
              class="w-2 h-4 md:w-5 md:h-5 text-gray-400"
            />
            <EyeOffIcon v-else class="w-5 h-5 text-gray-400" />
          </span>
          <ErrorMessage name="password" class="text-red-600 text-sm mt-1" />
        </div>

        <!-- Confirm Password Field -->
        <div class="relative">
          <label
            for="confirmPassword"
            class="block text-lg font-semibold text-gray-700"
            >Confirm Password</label
          >
          <Field
            id="confirmPassword"
            name="confirmPassword"
            :type="'password'"
            v-model="confirmPassword"
            placeholder="Confirm your password"
            class="mt-1 block w-full p-3 pl-10 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-purple-500 focus:border-purple-500"
          />
          <ErrorMessage
            name="confirmPassword"
            class="text-red-600 text-sm mt-1"
          />
        </div>

        <!-- Error Message -->
        <div v-if="errorMessage" class="text-red-500 text-center text-sm">
          {{ errorMessage }}
        </div>

        <!-- Submit Button -->
        <button
          type="submit"
          class="w-full bg-gradient-to-r from-purple-500 to-indigo-600 text-white py-3 rounded-lg text-lg font-semibold hover:from-purple-600 hover:to-indigo-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 transition duration-150"
        >
          <span v-if="loading">
            <svg
              class="animate-spin h-5 w-5 mr-2 text-white inline-block"
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
            Creating your account...
          </span>
          <span v-else>Sign Up</span>
        </button>
      </Form>

      <!-- Login Link -->
      <div class="text-center mt-6">
        <p class="text-sm text-gray-600">
          Already have an account?
          <NuxtLink
            to="/auth/login"
            class="text-indigo-600 font-semibold hover:underline transition duration-150"
            >Login</NuxtLink
          >
        </p>
      </div>
    </div>
  </div>
</template>
