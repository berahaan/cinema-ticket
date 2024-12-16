<script setup>
import { ref, onMounted } from "vue";
import { useNuxtApp } from "#app";
import { jwtDecode } from "jwt-decode";
import { useToast } from "vue-toastification";
import { ArrowLeftIcon } from "@heroicons/vue/solid";
import { GET_USERINFO } from "./graphql/queries/GET_USERINFO.graphql";
import { GET_EMAIL } from "./graphql/queries/GET_EMAIL.graphql";
import { UPDATE_USERINFO } from "./graphql/mutations/UPDATE_USERINFO.graphql";
const currentEmail = ref("");
const { goBack } = useGobackArrow();
const { selectClass } = useThemeColor();
const { removeBookmarks, getBookmarks, getBookmarkDetail, bookmarks } =
  useBookmarks();
const { handleFeaturedImageUpload } = useFile();
const { uploadProfile } = HandleImageUpload();
const { Logout } = useLogout();
const toast = useToast();
const userSelected = ref(false);
const isNotSelected = ref(true);
const { $apollo } = useNuxtApp();
const isBookmarkExist = ref(false);
const userinfo = ref({});
const currentProfile = ref("");
const formData = ref({
  firstname: "",
  lastname: "",
  email: "",
  role: "",
  featured_image: "",
  profilePicture: "",
});

const checkUserchanges = () => {
  userSelected.value = !userSelected.value;
  isNotSelected.value = !isNotSelected.value;
};
const CheckEmailExist = async (email) => {
  
  try {
    const { data } = await $apollo.query({
      query: GET_EMAIL,
      variables: { email }, // Pass the email as a variable
      fetchPolicy: "network-only", // Ensure the query fetches fresh data
    });
    return data.users.length > 0; // Return true if email exists, false otherwise
  } catch (error) {
    console.error("Error checking email existence", error);
    return false;
  }
};

const saveProfile = async () => {
  const decodedToken = jwtDecode(localStorage.getItem("accessToken"));
  const userId =
    decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

  // If email has changed, validate it
  if (formData.value.email && formData.value.email !== currentEmail.value) {
  

    const emailExists = await CheckEmailExist(formData.value.email);

    if (emailExists) {
     
      toast.error("The email is already in use. Please choose another.", {
        position: "top-center",
        duration: 2000,
      });
      formData.value.email = currentEmail.value;
      return;
    }
  } else {
    console.log("Email is unchanged.");
  }
 
  const { profileImage64 } = useProfileImageBase64Store();
  try {
    // Upload profile image if required
    await uploadProfile(profileImage64);
    const { featuredImageURL } = useImageStore();
    const UpdateResponse = await $apollo.mutate({
      mutation: UPDATE_USERINFO,
      variables: {
        id: userId,
        firstname: formData.value.firstname || userinfo.value.firstname,
        lastname: formData.value.lastname || userinfo.value.lastname,
        email: formData.value.email || userinfo.value.email,
        profilePictures: featuredImageURL || currentProfile.value,
      },
      refetchQueries: [{ query: GET_USERINFO }],
    });
    toast.success("Profile updated successfully!", {
      position: "top-center",
      duration: 2000,
    });

    formData.value = {
      firstname: "",
      lastname: "",
      email: "",
      profilePicture: "",
    };
  } catch (error) {
    
    toast.error("Failed to update profile. Please try again.", {
      position: "top-center",
      duration: 2000,
    });
  }
};

const fetchUserInfo = async () => {
  const decodedToken = jwtDecode(localStorage.getItem("accessToken"));
  const userId =
    decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
 
  try {
    const response = await $apollo.query({
      query: GET_USERINFO,
      variables: { userId: userId },
      fetchpolicy: "network-only",
    });
    userinfo.value = response.data.users[0];
    formData.value = {
      firstname: userinfo.value.firstname || "",
      lastname: userinfo.value.lastname || "",
      email: userinfo.value.email,
      role: userinfo.value.role || "",
      profilePicture: userinfo.value.profile_picture,
    };
    currentEmail.value = userinfo.value.email;
    currentProfile.value = userinfo.value.profile_picture;
  } catch (error) {
    console.error("Error fetching user info:", error);
  }
};

onMounted(() => {
  fetchUserInfo();
  getBookmarks();
});
</script>

<template>
  <div class="max-w-4xl mx-auto p-8 rounded-lg" :class="selectClass">
    <button
      @click="goBack"
      class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200"
    >
      <ArrowLeftIcon class="h-5 w-5" />
      <span>Go Back</span>
    </button>
    <header
      class="flex justify-between items-center border-b pb-4"
      :class="selectClass"
    >
      <h1 class="text-3xl font-bold">Basic Profile pages</h1>

      <button @click="Logout">
        <div class="flex">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M8.25 9V5.25A2.25 2.25 0 0 1 10.5 3h6a2.25 2.25 0 0 1 2.25 2.25v13.5A2.25 2.25 0 0 1 16.5 21h-6a2.25 2.25 0 0 1-2.25-2.25V15m-3 0-3-3m0 0 3-3m-3 3H15"
            />
          </svg>
          <nuxt-link to="/login">Logout</nuxt-link>
        </div>
      </button>
    </header>

    <section class="mt-6" v-if="userinfo.firstname" :class="selectClass">
      <h2 class="text-2xl">Profile Information</h2>
      <div class="flex flex-col md:flex-row items-center md:space-x-4 mt-4">
        <div class="flex flex-col items-center">
          <img
            :src="formData.profilePicture"
            alt="Profile Picture"
            class="w-32 h-32 rounded-full object-cover mb-2 shadow-md border-2 border-blue-200"
          />
          <button
            v-if="isNotSelected"
            @click="checkUserchanges"
            class="mt-4 px-4 py-2 rounded transition bg-teal-500"
            :class="selectClass"
          >
            change profiles
          </button>
          <input
            v-if="userSelected"
            @change="handleFeaturedImageUpload"
            type="file"
            accept="image/*"
            :class="selectClass"
            class="border border-gray-300 rounded-md p-2 focus:outline-none"
          />
        </div>
        <div class="flex-1 mt-4 md:mt-0">
          <input
            v-model="formData.firstname"
            type="text"
            :class="selectClass"
            class="border p-2 rounded w-full mt-2"
            placeholder="Enter your first name"
          />
          <input
            v-model="formData.lastname"
            type="text"
            class="border p-2 rounded w-full mt-2"
            :class="selectClass"
            placeholder="Enter your last name"
          />
          <input
            v-model="formData.email"
            type="email"
            class="border p-2 rounded w-full mt-2"
            :class="selectClass"
            placeholder="Enter your email"
          />
          <p class="mt-2" :class="selectClass">Role:{{ formData.role }}</p>
          <button
            @click="saveProfile"
            class="mt-4 hover:last:bg-blue-500 px-4 py-2 rounded transition"
            :class="selectClass"
          >
            Save Changes
          </button>
        </div>
      </div>
    </section>

    <section class="mt-6">
      <h2 class="text-2xl font-semibold text-gray-800">Bookmarked Movies</h2>
      <div v-if="isBookmarkExist" class="text-center mt-8 text-gray-600">
        Unfortunately, we didn't find any bookmarks here. Bookmark from the
        available lists.
      </div>
      <ul class="mt-4">
        <li
          v-for="bookmark in bookmarks"
          :key="bookmark.movie_id"
          class="flex items-center justify-between border-b py-2"
        >
          <div class="flex items-center">
            <img
              @click="getBookmarkDetail"
              :src="bookmark.movie.featured_images"
              alt="Movie Image"
              class="w-12 h-12 rounded mr-4 object-cover cursor-pointer"
            />
            <span class="font-medium text-gray-800">{{
              bookmark.movie.title || "Movie Title"
            }}</span>
          </div>
          <button
            @click="removeBookmarks(bookmark.movie_id)"
            class="bg-red-500 text-white px-2 py-1 rounded hover:bg-red-600 transition"
          >
            ❌ Remove
          </button>
        </li>
      </ul>
    </section>
  </div>
</template>
