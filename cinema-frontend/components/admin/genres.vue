<script setup>
import { ref, onMounted } from "vue";
import { GET_GENRE_ID } from "../components/graphql/queries/GET_GENRE_ID.graphql";
import { ADD_GENRE } from "../components/graphql/mutations/ADD_GENRE.graphql";
import { DELETE_GENRE } from "../components/graphql/mutations/DELETE_GENRE.graphql";
import { UPDATE_GENRE } from "../components/graphql/mutations/UPDATE_GENRE.graphql";
import { GET_GENRES } from "../components/graphql/queries/GET_GENRES.graphql";
import { Form, Field, ErrorMessage } from "vee-validate";
import { useToast } from "vue-toastification";
import { useNuxtApp } from "#app";
import Loading from "./Loading.vue";
const { selectClass } = useThemeColor();
const isLoading = ref(false);
const toast = useToast();
const { $apollo } = useNuxtApp();
const formData = ref({ name: "", genre_id: null });
const loading = ref(false);
const isEditing = ref(false);
const notification = ref({ message: "", type: "" });
const { fetchGenres, genres } = useFetchGenres();
const handleClick = (genre) => {
  // Scroll to the top of the page
  window.scrollTo({
    top: 0,
    behavior: "smooth",
  });

  // Call the editGenre function
  // Replace 'genre' with the actual data you want to edit
  editGenre(genre);
};
const onSubmit = async () => {
  loading.value = true;
  isLoading.value = true;
  try {
    if (isEditing.value) {
      const responseUpdate = await $apollo.mutate({
        mutation: UPDATE_GENRE,
        variables: { id: formData.value.genre_id, name: formData.value.name },
        refetchQueries: [{ query: GET_GENRES }],
      });
      if (responseUpdate && responseUpdate.data) {
        await fetchGenres();
        toast.success("updated successfully ", {
          position: "top-right",
          timeout: 3000,
        });
      }
    } else {
      if (formData.value.name) {
        const { data } = await $apollo.query({ query: GET_GENRES });
        const nameExists = data.genres.some(
          (genre) =>
            genre.name.trim().toLowerCase() ===
            formData.value.name.trim().toLowerCase()
        );
        if (nameExists) {
          // Check if any stars were found
          toast.error("genre name already exists. Please try another name.", {
            position: "top-right",
            timeout: 3000, // Toast stays visible for 5 seconds
          });
          return;
        } else {
          const repsonseGenre = await $apollo.mutate({
            mutation: ADD_GENRE,
            variables: { name: formData.value.name },
            refetchQueries: [{ query: GET_GENRES }],
          });
          if (repsonseGenre && repsonseGenre.data) {
            await fetchGenres();
            toast.success("Successfully added ", {
              position: "top-center",
              timeout: 3000,
            });
          }
        }
      }
    }
  } catch (error) {
    showNotification("Failed to save genre.", "error");
  } finally {
    setTimeout(() => {
      loading.value = false;
      isLoading.value = false;
    }, 200);
  }
};

// Edit genre
const editGenre = (genre) => {
  formData.value = { name: genre.name, genre_id: genre.genre_id };
  isEditing.value = true;
};

const deleteGenre = async (id) => {
  try {
    const { data } = await $apollo.query({
      query: GET_GENRE_ID,
      variables: { genre_id: id },
      fetchPolicy: "network-only",
    });

    if (data.genres.length > 0) {
      const response = await $apollo.mutate({
        mutation: DELETE_GENRE,
        variables: { genre_id: id },
        refetchQueries: [
          {
            query: GET_GENRES,
          },
        ],
      });

      if (response && response.data) {
        await fetchGenres();
        toast.success("Genres deleted successfully.", {
          position: "top-center",
          timeout: 3000,
        });
      }
    } else {
      toast.error("genre is already deleted ", {
        position: "top-center",
        timeout: 3000,
      });
    }
  } catch (error) {
    throw error;
  }
};
// Show notification
const showNotification = (message, type) => {
  notification.value = { message, type };
  setTimeout(() => {
    notification.value = { message: "", type: "" };
  }, 3000);
};
const validateName = () => {
  // Check if name is empty
  if (!formData.value.name) {
    return "Name is required";
  }
  if (formData.value.name.length < 5) {
    return "The name should be at least 5 letters";
  }
  return true;
};

onMounted(fetchGenres);
</script>

<template>
  <div class="max-w-4xl mx-auto mt-16">
    <h2 class="text-xl font-semibold mb-4">
      {{ isEditing ? "Update Genre" : "Add Genre" }}
    </h2>
    <div v-if="isLoading" class="flex justify-center mt-4 bg-white">
      <Loading />
    </div>
    <div class="flex justify-between mb-2">
      <label for="name"> Add Genre type here below</label>
    </div>
    <Form @submit="onSubmit" class="space-y-4">
      <div class="flex flex-col">
        <Field
          name="name"
          :rules="validateName"
          type="text"
          placeholder="Enter genre name"
          v-model="formData.name"
          class="border border-gray-300 rounded-md p-2"
          :class="selectClass"
        />
        <ErrorMessage name="name" class="text-red-600" />
      </div>

      <button
        type="submit"
        class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition duration-200 ease-in-out"
      >
        <span v-if="loading">{{
          isEditing ? "Updating..." : "Adding..."
        }}</span>
        <span v-else>{{ isEditing ? "Update Genre" : "Add Genre" }}</span>
      </button>
    </Form>

    <!-- Genre List -->
    <h2 class="text-xl font-semibold mt-10 mb-4">Genres List</h2>
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
            Genre Name
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
          v-for="(genre, index) in genres"
          :key="genre.genre_id"
          class="border-b transition duration-200 ease-in-out"
          :class="selectClass"
        >
          <td class="px-4 py-2">{{ index + 1 }}</td>
          <td class="px-4 py-2">{{ genre.name }}</td>
          <td class="px-4 py-2 text-right">
            <!-- Align cell content to the right -->
            <button
              @click="handleClick(genre)"
              class="bg-yellow-500 text-white px-3 py-1 rounded-md hover:bg-yellow-600 transition duration-200 ease-in-out mr-2"
            >
              ✏️ Edit
            </button>
            <button
              @click="deleteGenre(genre.genre_id)"
              class="bg-red-500 text-white px-3 py-1 rounded-md hover:bg-red-600 transition duration-200 ease-in-out"
            >
              ❌Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
