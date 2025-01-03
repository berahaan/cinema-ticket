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
import Loading from "./Loading.vue";
const { selectClass } = useThemeColor();
const { $apollo } = useNuxtApp();
const { fetchStars } = useFetchStars();
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
const handleClick = (star) => {
  window.scrollTo({
    top: 0,
    behavior: "smooth",
  });

  editStar(star);
};

const validateName = (event) => {
  if (!event) {
    return "Please add name of star";
  } else if (event?.length < 3) {
    return "start name should be at least 3 char";
  }
  return true;
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
      });
      await fetchStars();
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

            const response = await $apollo.mutate({
              mutation: ADD_STAR,
              variables: { name: formData.value.name },
            });
            if (response && response.data) {
              await fetchStars();

              toast.success("Star added successfully.", {
                position: "top-right",
                timeout: 3000,
              });
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
        stars.value = stars.value.filter((star) => star.star_id !== id);
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
    toast.error("Failed to delete star ", {
      position: "top-right",
      timeout: 2000,
    });
  }
};
onMounted(fetchStars);
</script>

<template>
  <div class="max-w-4xl mx-auto mt-16" :class="selectClassn">
    <!-- Add or Update Star Form -->
    <h2 class="text-xl mb-4">
      {{ isEditing ? "Update Star" : "Add Star" }}
    </h2>
    <div v-if="loading" class="flex justify-center mb-4 bg-white">
      <Loading />
    </div>
    <div class="flex justify-between mb-2">
      <label for="name">Add Stars Here</label>
    </div>
    <Form @submit="onSubmit" class="space-y-4 px-2">
      <div class="flex flex-col">
        <Field
          v-model="formData.name"
          type="text"
          name="name"
          placeholder="Enter star's name"
          class="border border-gray-300 rounded-md p-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
          :class="selectClass"
          :rules="validateName"
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
              @click="handleClick(star)"
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
