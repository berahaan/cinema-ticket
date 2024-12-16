<script setup>
import { useRoute } from "vue-router";
import { useRouter } from "vue-router";
import { ArrowLeftIcon, CheckCircleIcon } from "@heroicons/vue/solid";
import { jwtDecode } from "jwt-decode";
import { onMounted, ref } from "vue";
import { useNuxtApp } from "#app";
import { GET_TICKET_INFO } from "./graphql/queries/GET_TICKET_INFO.graphql";

import Loading from "./admin/Loading.vue";
const route = useRoute();
const { $apollo } = useNuxtApp();
const router = useRouter();
const { formatScheduleDateTime } = useFormatSchedule();
const { selectClass } = useThemeColor();
const noticket = ref(false);
const ticketFound = ref(false);
const tx_ref = route.query.tx_ref;
const ticket = ref(null);
const statusClass = ref("");
const loading = ref(false);
const backhome = () => {
  const accessToken = localStorage.getItem("accessToken");
  const decodedToken = jwtDecode(accessToken);
  const defaultRole =
    decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
  if (defaultRole === "admin" || defaultRole == "managers") {
    router.push("/admin/Movielist");
  } else {
    router.push("/user/Movielist");
  }
};
const fetchTicket = async () => {
  loading.value = true;

  try {
    const response = await $apollo.query({
      query: GET_TICKET_INFO,
      variables: { tx_ref },
    });
    if (response.data.ticket.length > 0) {
     
      ticketFound.value = true;
      ticket.value = response.data.ticket[0];
      
      statusClass.value =
        ticket.value.payment_status === "paid"
          ? "text-green-500"
          : "text-red-500";
    } else {
      noticket.value = true;
      ticketFound.value = false;
    }
  } catch (error) {
   
    throw error
  } finally {
    loading.value = false;
  }
};

onMounted(fetchTicket);
</script>
<template>
  <div
    class="min-h-screen flex items-center justify-center py-10 border mt-4"
    :class="selectClass"
  >
    <div class="shadow-md rounded-lg max-w-4xl p-6" :class="selectClass">
      <div class="flex items-center mx-24" v-if="ticketFound">
        <h1
          v-if="
            ticket.payment_status === 'paid' ||
            ticket.payment_status === 'pending'
          "
          class="text-2xl font-semibold text-green-600"
        >
          🎉 Thank You For purchase!
        </h1>
      </div>
      <!-- Payment Status Messages -->
      <div v-if="noticket" class="text-center">
        <p class="text-red-500">
          Unfortunately, your ticket is not valid or the payment was
          unsuccessful. Please try again.
        </p>
        <button
          @click="backhome"
          class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200 text-lg border-b pb-2 mb-4"
        >
          <span>Back Home</span>
        </button>
      </div>

      <div v-if="loading" class="text-center mt-6">
        <Loading />
      </div>

      <div v-if="ticketFound" :class="selectClass">
        <!-- Payment Status -->
        <div v-if="ticket.payment_status === 'paid'" class="mt-6">
          <!-- Back Home Button -->
          <div class="flex justify-between">
            <button
              @click="backhome"
              class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200 text-lg border-b pb-2 mb-4"
            >
              <ArrowLeftIcon class="h-5 w-5" />
              <span>Back Home</span>
            </button>
          </div>

          <!-- Success Message -->
          <div class="text-center">
            <div class="flex items-center justify-center space-x-2 mt-4">
              <CheckCircleIcon class="text-green-500 w-8 h-8" />
              <p class="text-lg font-semibold text-gray-800">
                Success! Your ticket has been successfully ordered. 🎉
              </p>
            </div>
            <p class="mt-2 text-gray-600">
              Please check your <strong>Email</strong> for payment details and
              download your ticket attachment.
            </p>
          </div>

          <!-- Schedule Information -->
          <div class="space-y-3">
            <ul class="space-y-2">
              <h2
                class="text-lg font-medium border-b pb-2 mb-4 mt-6 text-gray-800"
              >
                🗓️ You have bought movies that Scheduled for :
              </h2>
              <li
                v-if="ticket.movie_schedule"
                class="rounded-lg p-3 shadow-md bg-gray-50"
              >
                <p class="flex items-center">
                  <span class="font-semibold text-gray-700">
                    {{
                      formatScheduleDateTime(
                        ticket.movie_schedule.schedule_date,
                        ticket.movie_schedule.start_time,
                        ticket.movie_schedule.end_time
                      )
                    }}
                  </span>
                  <span class="ml-2 text-gray-600"
                    >at {{ ticket.movie_schedule.cinema_hall }}</span
                  >
                </p>
              </li>
            </ul>
          </div>

          <!-- Additional Instructions -->
          <p class="mt-4 text-gray-700">
            <strong>Reminder:</strong> Please arrive at the cinema on time and
            bring your ticket attachment for smooth entry. We look forward to
            seeing you!
          </p>
        </div>

        <div v-else-if="ticket.payment_status === 'pending'" class="mt-6">
          <div class="text-center text-yellow-600">
            <p>
              💛 Your payment is still pending. Please wait for 15 or more
              minutes for confirmation.
            </p>
            <p>Thank You for your patience !!</p>
          </div>
        </div>

        <!-- Failed Payment Status -->
        <div v-else-if="ticket.payment_status === 'failed'" class="mt-6">
          <div class="text-center text-red-600">
            <p>❌ Unfortunately, your payment has failed. Please try again.</p>
            <button
              @click="backhome"
              class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200 text-lg border-b pb-2 mb-4"
            >
              <span>Try again</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
