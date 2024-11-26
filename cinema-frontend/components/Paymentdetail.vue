<script setup>
// ssh -R 80:localhost:4000 serveo.net
import { useRoute } from "vue-router";
import { useRouter } from "vue-router";
import { ArrowLeftIcon, ArrowDownIcon } from "@heroicons/vue/solid";
import { jwtDecode } from "jwt-decode";
import { onMounted, ref } from "vue";
import { useNuxtApp } from "#app";
import { GET_TICKET_INFO } from "./graphql/queries/GET_TICKET_INFO.graphql";
import { DOWNLOAD_TICKET } from "./graphql/mutations/DOWNLOAD_TICKET.graphql";
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
  console.log("Clicked now ");
  const accessToken = localStorage.getItem("accessToken");
  const decodedToken = jwtDecode(accessToken);
  const defaultRole =
    decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
  if (defaultRole === "admin" || defaultRole == "managers") {
    router.push("/admin/Movielist");
  } else {
    router.push("/user");
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
      console.log("response ", response.data);
      ticketFound.value = true;
      ticket.value = response.data.ticket[0];
      console.log("ticket.value", ticket.value);
      statusClass.value =
        ticket.value.payment_status === "paid"
          ? "text-green-500"
          : "text-red-500";
    } else {
      noticket.value = true;
      ticketFound.value = false;
    }
  } catch (error) {
    console.error("Error fetching ticket details:", error);
  } finally {
    loading.value = false;
  }
};
const downloadTicket = async () => {
  console.log("Tx_ref ", tx_ref);
  try {
    const { data } = await $apollo.mutate({
      mutation: DOWNLOAD_TICKET,
      variables: { tx_ref },
    });

    if (data && data.downloadTicket.pdf_base64) {
      console.log("PDF data received:", data.pdf_base64);
      // Decode the Base64 string to binary data
      const binaryData = atob(data.downloadTicket.pdf_base64);
      // Convert binary data to an array of bytes
      const byteArray = new Uint8Array(binaryData.length);
      for (let i = 0; i < binaryData.length; i++) {
        byteArray[i] = binaryData.charCodeAt(i);
      }
      // Create a Blob from the byteArray
      const blob = new Blob([byteArray], { type: "application/pdf" });

      // Create a download link and trigger the download
      const downloadUrl = URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = downloadUrl;
      link.download = "ticket.pdf";
      link.click();
      // Clean up
      URL.revokeObjectURL(downloadUrl);
    } else {
      console.log("Else part is executed ");
    }
  } catch (error) {
    console.error("Error downloading the ticket:", error);
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
          <div class="flex justify-between">
            <button
              @click="downloadTicket"
              :class="{
                'flex items-end text-lg font-medium border-b pb-2 mb-4': true,
                'text-teal-600': ticket.payment_status === 'paid',
                'text-gray-500': ticket.payment_status !== 'paid',
              }"
            >
              <ArrowDownIcon class="h-5 w-5 mr-2" />
              Download Ticket
            </button>
            <button
              @click="backhome"
              class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200 text-lg border-b pb-2 mb-4"
            >
              <ArrowLeftIcon class="h-5 w-5" />
              <span>Back Home</span>
            </button>
          </div>
          <div class="text-center">
            <p class="mt-2">
              You have successfully ordered your ticket. Here are your payment
              details:
            </p>
          </div>

          <h2 class="text-lg font-medium border-b pb-2 mb-4">
            🎫 Ticket Information
          </h2>
          <div class="space-y-3">
            <!-- Ticket Details -->
            <div class="flex justify-between">
              <span class="font-medium">Your Name:</span>
              <span>{{ ticket.firstname }} {{ ticket.lastname }}</span>
            </div>

            <div class="flex justify-between">
              <span class="font-medium">Movie Title:</span>
              <span>{{ ticket.movie.title }}</span>
            </div>
            <div class="flex justify-between">
              <span class="font-medium">Director:</span>
              <span>{{ ticket.movie.director.name }}</span>
            </div>
            <div class="flex justify-between">
              <span class="font-medium">Ticket Quantity:</span>
              <span>{{ ticket.ticket_quantity }}</span>
            </div>
            <div class="flex justify-between">
              <span class="font-medium">Amount:</span>
              <span>{{ ticket.amount }}</span>
            </div>
            <div class="flex justify-between">
              <span class="font-medium">Transaction Reference:</span>
              <span>{{ ticket.tx_ref }}</span>
            </div>
          </div>

          <!-- Schedule Information -->
          <div class="space-y-3">
            <ul class="space-y-2">
              <h2 class="text-lg font-medium border-b pb-2 mb-4 mt-6">
                🗓️ You bought Schedule for
              </h2>
              <li v-if="ticket.movie_schedule" class="rounded-lg p-3 shadow-md">
                <p class="flex items-center">
                  <span class="font-semibold mr-2">
                    {{
                      formatScheduleDateTime(
                        ticket.movie_schedule.schedule_date,
                        ticket.movie_schedule.start_time,
                        ticket.movie_schedule.end_time
                      )
                    }}
                    <span> at {{ ticket.movie_schedule.cinema_hall }}</span>
                  </span>
                </p>
              </li>
            </ul>
          </div>
          <p>
            Please make yourself available before deadlines and download This
            ticket
          </p>
          <p>and bring them with you</p>
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
