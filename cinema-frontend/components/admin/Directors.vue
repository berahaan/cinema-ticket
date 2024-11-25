<script setup>
import { useToast } from "vue-toastification";
import { ref, onMounted } from "vue";
import { UPDATE_DIRECTOR } from "../components/graphql/mutations/UPDATE_DIRECTOR.graphql";
import { DELETE_DIRECTOR } from "../components/graphql/mutations/DELETE_DIRECTOR.graphql";
import { ADD_DIRECTOR } from "../components/graphql/mutations/ADD_DIRECTOR.graphql";
import { GET_DIRECTORS } from "../components/graphql/queries/GET_DIRECTORS.graphql";
import { GET_DIRECTORS_ID } from "../components/graphql/queries/GET_DIRECTORS_ID.graphql";
import { useNuxtApp } from "#app";
import { Form, Field, ErrorMessage } from "vee-validate";
const toast = useToast();
const { selectClass } = useThemeColor();
const { refreshPage } = useRefresh();
const { $apollo } = useNuxtApp();
const directors = ref([]);
const formData = ref({ name: "", director_id: null });
const loading = ref(false);
const isEditing = ref(false);
const errors = ref({});
const fetchDirectors = async () => {
  try {
    const response = await $apollo.query({
      query: GET_DIRECTORS,
      fetchPolicy: "network-only",
    });
    directors.value = response.data.directors;
  } catch (error) {
    console.error("Error fetching directors:", error);
  }
};

const validateName = (value) => {
  // Check if name is empty
  if (!value) {
    return "Name is required";
  }
  // Check if name length is less than 4
  if (value.length < 4) {
    return "The name should be at least 4 letters";
  }

  return true; // No errors
};

// Add or update a director
const onSubmit = async () => {
  loading.value = true;
  try {
    if (isEditing.value) {
      // Update director
      const updateResponse = await $apollo.mutate({
        mutation: UPDATE_DIRECTOR,
        variables: {
          id: formData.value.director_id,
          name: formData.value.name,
        },
      });
      if (updateResponse && updateResponse.data) {
        isEditing.value = false;
        toast.success("Updated successfully ", {
          position: "top-center",
          durations: 3000,
        });
        // formData.value.name = ""; // Clear the form
      }
    } else {
      // Add director
      if (formData.value.name) {
        // Check if director already exists
        const { data } = await $apollo.query({ query: GET_DIRECTORS });
        const nameExists = data.directors.some(
          (director) =>
            director.name.trim().toLowerCase() ===
            formData.value.name.trim().toLowerCase()
        );
        if (nameExists) {
          toast.error(
            "Director name already exists. Please try another name.",
            {
              position: "top-right",
              timeout: 3000, // Toast stays visible for 3 seconds
            }
          );
          // formData.value.name = ""; // Clear name if already exists
          return;
        } else {
          console.log("Adding director to database...");
          const responseDirector = await $apollo.mutate({
            mutation: ADD_DIRECTOR,
            variables: { name: formData.value.name },
            refetchQueries: [{ query: GET_DIRECTORS }],
          });
          if (responseDirector && responseDirector.data) {
            toast.success("Successfully added ", {
              position: "top-center",
              timeout: 3000,
            });
            // formData.value.name = ""; // Clear the form
          }
        }
      }
    }
  } catch (error) {
    console.error("Error saving director:", error);
    alert("Failed to save director");
  } finally {
    loading.value = false;
  }
};

const editDirector = (director) => {
  formData.value = { name: director.name, director_id: director.director_id };
  isEditing.value = true;
};

const deleteDirector = async (id) => {
  try {
    console.log("To be deleted Id of directors ", id);
    const { data } = await $apollo.query({
      query: GET_DIRECTORS_ID,
      variables: { directors_id: id },
      fetchPolicy: "network-only",
    });
    // console.log(data);
    if (data.directors.length > 0) {
      console.log(
        "Id to be deleted for this genres in data.stars.length > 0 ",
        id
      );
      const response = await $apollo.mutate({
        mutation: DELETE_DIRECTOR,
        variables: { director_id: id },
        refetchQueries: [
          {
            query: GET_DIRECTORS,
          },
        ],
      });

      if (response && response.data) {
        toast.success("directors deleted successfully.", {
          position: "top-center",
          timeout: 3000,
        });
      }
    } else {
      toast.error("directors is already deleted ", {
        position: "top-center",
        timeout: 3000,
      });
    }
  } catch (error) {
    console.error("Error deleting star:", error);
  }
};
onMounted(fetchDirectors);
</script>

<template>
  <div class="max-w-4xl mx-auto mt-16">
    <!-- Add or Update Director Form -->
    <h2 class="text-xl font-semibold mb-4">
      {{ isEditing ? "Update Director" : "Add Director" }}
    </h2>
    <div class="flex justify-between mb-2">
      <label for="name">Director Name:</label>
      <button @click="refreshPage(fetchDirectors)" class="ml-4">
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
    <Form @submit="onSubmit" class="space-y-4">
      <div class="flex flex-col">
        <Field
          v-model="formData.name"
          type="text"
          name="name"
          :rules="validateName"
          placeholder="Enter director's name"
          class="border border-gray-300 rounded-md p-2 focus:outline-none"
          :class="selectClass"
        />
        <span v-if="errors.name" class="text-red-600 text-sm mt-1">{{
          errors.name
        }}</span>
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
        <span v-else>{{ isEditing ? "Update Director" : "Add Director" }}</span>
      </button>
    </Form>

    <!-- Director List -->
    <h2 class="text-xl font-semibold mt-10 mb-4">Directors List</h2>

    <table class="min-w-full border border-gray-300">
      <thead>
        <tr>
          <th
            class="px-4 py-2 text-left font-semibold border-b"
            :class="selectClass"
          >
            index
          </th>
          <th
            class="px-4 py-2 text-left font-semibold border-b"
            :class="selectClass"
          >
            Director Name
          </th>
          <th
            class="px-6 py-2 font-semibold border-b text-right"
            :class="selectClass"
          >
            <!-- Align text to the right -->
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(director, index) in directors"
          :key="director.director_id"
          class="border-b transition duration-200 ease-in-out"
          :class="selectClass"
        >
          <td class="px-4 py-2">{{ index + 1 }}</td>
          <td class="px-4 py-2">{{ director.name }}</td>
          <td class="px-4 py-2 text-right">
            <!-- Align cell content to the right -->
            <button
              @click="editDirector(director)"
              class="bg-yellow-500 text-white px-3 py-1 rounded-md hover:bg-yellow-600 transition duration-200 ease-in-out mr-2"
            >
              ✏️ Edit
            </button>
            <button
              @click="deleteDirector(director.director_id)"
              class="bg-red-500 text-white px-3 py-1 rounded-md hover:bg-red-600 transition duration-200 ease-in-out"
            >
              ❌ Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
