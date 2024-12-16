<script setup>
import { GET_ADMIN_STATS } from "./components/graphql/queries/GET_ADMIN_STATS.graphql";
import { GET_USER } from "../graphql/queries/GET_USER.graphql";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useNuxtApp } from "#app";
import { onMounted } from "vue";
import { jwtDecode } from "jwt-decode";
import Loading from "../admin/Loading.vue";
const { $apollo } = useNuxtApp();
const totalMovies = ref(0);
const router = useRouter();
const totalSchedules = ref(0);
const totalUsers = ref(0);
const upcomingEvents = ref(3);
const firstName = ref("");
const isloading = ref(true);
const getUserName = async () => {
  const accessToken = localStorage.getItem("accessToken");
  if (!accessToken) {
    router.push("/auth/login");
    return;
  }
  const decodedToken = jwtDecode(accessToken);
  const userid =
    decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

  try {
    const { data } = await $apollo.query({
      query: GET_USER,
      variables: { userid }, // Pass the userid as a variable
    });

    if (data.users.length > 0) {
      firstName.value = data.users[0].firstname; // Access the first user's firstname
    } else {
      console.error("No user found.");
    }
  } catch (error) {
    throw error;
  }
};
const fetchAdminStatus = async () => {
  try {
    const { data } = await $apollo.query({
      query: GET_ADMIN_STATS,
      fetchPolicy: "network-only",
    });
    totalMovies.value = data.total_movies[0]?.total_count || 0;
    totalSchedules.value = data.total_schedules[0]?.total_count || 0;
    totalUsers.value = data.total_users[0]?.total_count || 0;
    isloading.value = false;
  } catch (error) {
    console.error("Error fetching", error);
  } finally {
    isloading.value = false;
  }
};
onMounted(async () => {
  isloading.value = true;
  await getUserName();
  await fetchAdminStatus();
});
</script>

<style scoped>
.rotate-180 {
  transform: rotate(180deg);
}
</style>

<template>
  <div class="min-h-screen p-6 mt-20">
    <div class="mb-6">
      <h1 class="text-3xl">Welcome, {{ firstName }}!</h1>
      <p>Here’s a quick summary of what’s happening in the system today.</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
      <div class="p-6 rounded-lg shadow-md border">
        <h2 class="text-xl">Total Movies</h2>
        <p class="text-3xl">
          {{ totalMovies }}
        </p>
      </div>

      <!-- Total Schedules -->
      <div class="p-6 rounded-lg shadow-md border">
        <h2 class="text-xl">Total Schedules</h2>
        <p class="text-3xl">
          {{ totalSchedules }}
        </p>
      </div>

      <!-- Registered Users -->
      <div class="p-6 rounded-lg shadow-md border">
        <h2 class="text-xl font-semibold">Registered Users</h2>
        <p class="text-3xl">
          {{ totalUsers }}
        </p>
      </div>

      <!-- Upcoming Events -->
      <div class="p-6 rounded-lg shadow-md border">
        <h2 class="text-xl">Upcoming Events</h2>
        <p class="text-3xl">
          {{ upcomingEvents }}
        </p>
      </div>
    </div>
    <div class="mb-6">
      <h2 class="p-6 rounded-lg shadow-md border">Quick Links</h2>
      <div class="gap-4 mt-2 items-center flex justify-center">
        <nuxt-link
          to="/admin/Movielist"
          class="bg-green-500 hover:bg-green-600 text-white p-4 rounded-lg shadow-md text-center"
        >
          Movie Management
        </nuxt-link>
        <nuxt-link
          to="/admin/MovieSchedule"
          class="bg-green-500 hover:bg-green-600 text-white p-4 rounded-lg shadow-md text-center"
        >
          Schedule Management
        </nuxt-link>
      </div>
    </div>

    <!-- Recent Activity or Logs -->
    <div class="p-6 rounded-lg shadow-md border">
      <h2 class="text-2xl">Recent Activity</h2>
      <ul class="p-6 rounded-lg shadow-md border">
        <li>
          <strong>New Movie Added:</strong> "Inception" was added on Oct 20th.
        </li>
        <li>
          <strong>Schedule Updated:</strong> The schedule for "Avatar" was
          modified.
        </li>
        <li>
          <strong>Admin Action:</strong> User "John Doe" was promoted to admin.
        </li>
      </ul>
    </div>
    <div
      v-if="isloading"
      class="flex justify-center items-center min-h-screen bg-white mt-10"
    >
      <Loading />
    </div>
    <div class="p-6 rounded-lg shadow-md border">
      <h2 class="text-2xl mb-4">System Health</h2>
      <div class="p-6 rounded-lg shadow-md border">
        <p><strong>API Status:</strong> Operational</p>
        <p><strong>Database Health:</strong> 99% uptime in the last 24 hours</p>
      </div>
    </div>
  </div>
</template>
