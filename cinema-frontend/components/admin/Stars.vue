<script setup>
import { ref, onMounted } from "vue";
import { useNuxtApp } from "#app";
import { GET_STARS } from "../components/graphql/queries/GET_STARS.graphql";
import { GET_STAR_ID } from "../components/graphql/queries/GET_STAR_ID.graphql";
import { ADD_STAR } from "../components/graphql/mutations/ADD_STAR.graphql";
import { UPDATE_STAR } from "../components/graphql/mutations/UPDATE_STAR.graphql";
import { DELETE_STAR } from "../components/graphql/mutations/DELETE_STAR.graphql";
import { useToast } from "vue-toastification";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import Loading from "./Loading.vue";
const {validationSchema}=useValidate()
const { refreshPage } = useRefresh();
const { selectClass } = useThemeColor();
const { $apollo } = useNuxtApp();
const stars = ref([]);
const formData = ref({ name: "", star_id: null });
const loading = ref(false);
const success = ref(false);
const isEditing = ref(false);
const toast = useToast();
const colorStore = useColorModeStore();
const selectClassn = computed(() =>
  colorStore.colorMode === "dark"
    ? "bg-gray-800 text-white border-gray-100"
    : "bg-gray-50 text-gray-900 border-gray-200"
);
const fetchStars = async () => {
  loading.value = true;
  try {
    const response = await $apollo.query({
      query: GET_STARS,
      fetchPolicy: "network-only",
    });
    // await new Promise((resolve) => setTimeout(resolve, 4000));
    stars.value = response.data.stars;
  } catch (error) {
    console.error("Error fetching stars:", error);
  } finally {
    loading.value = false;
  }
};
const onSubmit = async () => {
  success.value = true;
  loading.value = true;
  try {
    if (isEditing.value) {
      await $apollo.mutate({
        mutation: UPDATE_STAR,
        variables: {
          id: formData.value.star_id,
          name: formData.value.name,
        },
        refetchQueries: [{ query: GET_STARS }],
      });
      toast.success("updated successfully ", {
        position: "top-right",
        timeout: 3000,
      });
    } else {
      success.value = true;
      if (formData.value.name) {
        console.log("Checking database for existing star name...");
        try {
          const { data } = await $apollo.query({ query: GET_STARS });
          const nameExists = data.stars.some(
            (star) =>
              star.name.trim().toLowerCase() ===
              formData.value.name.trim().toLowerCase()
          );
          if (nameExists) {
            // Check if any stars were found
            toast.error("Star name already exists. Please try another name.", {
              position: "top-right",
              timeout: 3000, // Toast stays visible for 5 seconds
            });
            // formData.value.name = "";
            return;
          } else {
            success.value = true;
            console.log(
              "Star does not exist in the database. Proceeding to add."
            );
            const response = await $apollo.mutate({
              mutation: ADD_STAR,
              variables: { name: formData.value.name },
              refetchQueries: [{ query: GET_STARS }],
            });
            if (response && response.data) {
              toast.success("Star added successfully.", {
                position: "top-right",
                timeout: 3000,
              });
              formData.value.name=""
            }
          }
        } catch (error) {
          console.log("Error:", error); // Use error, not errors
        }
      }
    }

    // formData.value = { name: "", star_id: null };
    isEditing.value = false;
  } catch (error) {
    console.error("Error saving star:", error);
    toast.error("Failed to save star", {
      position: "top-right",
      duration: 2000,
    });
  } finally {
    loading.value = false;
  }
};
const editStar = (star) => {
  formData.value = { name: star.name, star_id: star.star_id };
  isEditing.value = true;
};

const deleteStar = async (id) => {
  try {
    console.log("To be deleted Id of stars ", id);
    const { data } = await $apollo.query({
      query: GET_STAR_ID,
      variables: { star_id: id },
      fetchPolicy: "network-only",
    });
    if (data.stars.length > 0) {
      const response = await $apollo.mutate({
        mutation: DELETE_STAR,
        variables: { id },
        refetchQueries: [
          {
            query: GET_STARS,
          },
        ],
      });

      if (response && response.data) {
        toast.success("Star deleted successfully.", {
          position: "top-center",
          timeout: 3000,
        });
      }
    } else {
      toast.error("Star is already deleted ", {
        position: "top-center",
        timeout: 3000,
      });
    }
  } catch (error) {
    console.error("Error deleting star:", error);
    toast.error("Failed to delete star ",{
      position:"top-right",
      timeout:2000
    })
  }
};
// const validationSchema = yup.object({
//     name: yup.string().required("name is required from yup"),
    
//   });
onMounted(fetchStars);
</script>

<template>
  <div class="max-w-4xl mx-auto mt-10" :class="selectClassn">
    <!-- Add or Update Star Form -->
    <h2 class="text-xl mb-4">
      {{ isEditing ? "Update Star" : "Add Star" }}
    </h2>
    <div v-if="loading" class="flex justify-center mb-4 bg-white">
      <Loading />
    </div>
    <div class="flex justify-between mb-2">
          <label for="name" class="" :class="selectClass"> Star Name: </label>
          <button @click="refreshPage(fetchStars)" class="ml-4">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-6 h-6 text-blue-500 hover:text-blue-700 transition-colors"
              viewBox="0 0 24 24"
              fill="currentColor"
            >
              <path
                d="M17.65 6.35A7.95 7.95 0 0012 4V1L8 5l4 4V6c1.78 0 3.4.68 4.65 1.8a6 6 0 011.29 6.12l1.45 1.45c1.19-2.32 1.11-5.25-.6-7.57zm-11.3 11.3A7.95 7.95 0 0012 20v3l4-4-4-4v3c-1.78 0-3.4-.68-4.65-1.8a6 6 0 01-1.29-6.12l-1.45-1.45c-1.19 2.32-1.11 5.25.6 7.57z"
              />
            </svg>
          </button>
        </div>
    <Form @submit="onSubmit" :validation-schema="validationSchema"   class="space-y-4 px-2">
      <div class="flex flex-col">
       
        <Field
          v-model="formData.name"
          type="text"
          name="name"
          placeholder="Enter star's name"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="selectClass"
        />
        <ErrorMessage name="name" class="text-red-600" />
      </div>

      <button
        type="submit"
        :disabled="loading"
        class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition duration-200 ease-in-out"
      >
        <span v-if="loading">{{
          isEditing ? "Updating..." : "Adding..."
        }}</span>
        <span v-else>{{ isEditing ? "Update Star" : "Add Star" }}</span>
      </button>
    </Form>

    <!-- Add spinner before stars list if loading -->
    <div v-if="loading" class="flex justify-center mt-4 bg-white">
      <Loading />
    </div>
    <!-- Stars List as Table -->
    <h2 class="text-xl mt-10 mb-4" :class="selectClass">Stars List</h2>
    <p :class="selectClass">Total Stars: {{ stars.length }}</p>
    <table class="w-full border border-gray-300" :class="selectClass">
      <thead>
        <tr>
          <th class="px-4 py-2 text-left border-b" :class="selectClass">
            Index
          </th>
          <th class="px-4 py-2 text-left border-b" :class="selectClass">
            Star Name
          </th>
          <th class="px-6 py-2 border-b text-right" :class="selectClass">
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(star, index) in stars"
          :key="star.star_id"
          class="border-b transition duration-200 ease-in-out"
        >
          <td class="px-4 py-2">{{ index + 1 }}</td>
          <td class="px-4 py-2">{{ star.name }}</td>
          <td class="px-4 py-2 text-right">
            <button
              @click="editStar(star)"
              class="bg-yellow-500 text-white px-3 py-1 rounded-md hover:bg-yellow-600 transition duration-200 ease-in-out mr-2"
              :class="selectClass"
            >
              ✏️ Edit
            </button>
            <button
              @click="deleteStar(star.star_id)"
              class="bg-red-500 text-white px-3 py-1 rounded-md hover:bg-red-600 transition duration-200 ease-in-out"
              :class="selectClass"
            >
              ❌ Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
