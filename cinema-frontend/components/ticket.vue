<script setup>
import { onMounted } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import { ArrowLeftIcon, ArrowRightIcon } from "@heroicons/vue/solid";
const { goBack } = useGobackArrow();
const { formatScheduleDateTime, isScheduleExpired, isvalid } =
  useFormatSchedule();
const {
  movie,
  fetchInfo,
  purchaseTickets,
  fetchUserName,
  form,
  islengthGreater,
  totalPrice,
  ticketQuantity,
  selectedSchedule,
  isloading,
} = useTicket();
const { selectClass } = useThemeColor();
const handleTicket = async (values) => {
  console.log("Form data information:", values);
  try {
    await purchaseTickets(form);
  } catch (error) {
    console.log("Error", error);
  }
};

// Function to check if the schedule date/time has passed
const validationSchema = yup.object({
  first_name: yup.string().required("First Name is required"),
  last_name: yup.string().required("Last Name is required"),
  phone_number: yup
    .string()
    .matches(/^(09|07)\d{8}$/, "Phone number is not valid")
    .required("Phone Number is required"),
  email: yup
    .string()
    .email("Email must be valid")
    .matches(
      /^[a-zA-Z0-9._%+-]+@(gmail|yahoo)\.com$/,
      "Email must be a valid Gmail or Yahoo address "
    )
    .required("Email is required"),
  selectedSchedule: yup.object().required("Please select a schedule"),
  ticketQuantity: yup
    .number()
    .typeError("Ticket number is required ")
    .required("At least one ticket is required")
    .integer("Ticket quantity must be a whole number")
    .min(1, "Ticket quantity must be at least 1"),
});

onMounted(async () => {
  await fetchInfo();
  await fetchUserName();
});
</script>

<template>
  <div
    class="container mx-auto p-6 rounded-lg shadow-lg mt-10"
    :class="selectClass"
  >
    <div class="movie-ticket rounded-lg shadow-lg p-6">
      <div
        class="movie-info flex flex-col md:flex-row items-start mb-6"
        :class="selectClass"
      >
        <!-- Movie Image -->
        <div class="w-full md:w-1/3 mb-4 md:mb-0">
          <img
            :src="movie.featured_images"
            alt="Movie Featured Image"
            class="w-full h-auto max-h-96 object-cover rounded-lg shadow-md transform transition-transform hover:scale-105 cursor-pointer"
          />
        </div>
        <!-- Movie Text Info -->
        <div class="w-full md:w-2/3 md:ml-6">
          <div class="flex justify-between">
            <h1 class="text-4xl font-extrabold mb-4" :class="selectClass">
              {{ movie.title }}
            </h1>
            <button
              @click="goBack"
              class="flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium py-2 px-4 rounded-md transition duration-200"
            >
              <ArrowLeftIcon class="h-5 w-5" />
              <span>Go Back</span>
            </button>
          </div>
          <p class="text-lg mb-2" :class="selectClass">
            Description: {{ movie.description }}
          </p>
          <div class="genres flex flex-wrap">
            Duration
            <span
              class="px-3 mx-2 py-1 rounded-full text-sm font-medium"
              :class="selectClass"
              >{{ movie.duration }} mins</span
            >
          </div>
          <div class="genres flex flex-wrap mb-2">
            Genre:
            <span
              v-for="genre in movie.movie_genres"
              :key="genre.genre.name"
              class="mr-2 px-3 py-1 rounded-full text-sm font-medium mx-2"
              :class="selectClass"
            >
              {{ genre.genre.name }}
            </span>
          </div>
          <!-- Stars -->
          <div class="stars flex flex-wrap mb-2">
            Available Star:
            <span
              v-for="star in movie.movie_stars"
              :key="star.star.name"
              class="mr-2 px-3 mx-2 py-1 rounded-full text-sm font-medium"
              :class="selectClass"
            >
              {{ star.star.name }}
            </span>
          </div>
        </div>
      </div>

      <div class="schedules mb-6" v-if="islengthGreater">
        <h2 class="text-2xl font-bold mb-4" :class="selectClass">
          Select a Schedule
        </h2>

        <Form
          :validation-schema="validationSchema"
          @submit="handleTicket"
          class="grid grid-cols-1 sm:grid-cols-2 gap-5"
        >
          <!-- First Name Field -->
          <div>
            <label class="block font-semibold mb-1" :class="selectClass"
              >First Name</label
            >
            <Field
              v-model="form.first_name"
              name="first_name"
              type="text"
              class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2"
              :class="selectClass"
            />
            <ErrorMessage name="first_name" class="text-red-500" />
          </div>

          <!-- Last Name Field -->
          <div>
            <label class="block font-semibold mb-1" :class="selectClass"
              >Last Name</label
            >
            <Field
              v-model="form.last_name"
              name="last_name"
              type="text"
              class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              :class="selectClass"
            />
            <ErrorMessage name="last_name" class="text-red-500" />
          </div>

          <!-- Phone Number Field -->
          <div>
            <label class="block font-semibold mb-1" :class="selectClass"
              >Phone Number</label
            >
            <Field
              v-model="form.phone_number"
              name="phone_number"
              type="text"
              class="w-full p-3 border border-gray-300 rounded-lg"
              :class="selectClass"
            />
            <ErrorMessage name="phone_number" class="text-red-500" />
          </div>
          <div>
            <label class="block font-semibold mb-1" :class="selectClass"
              >Email</label
            >
            <Field
              v-model="form.email"
              name="email"
              type="email"
              class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              :class="selectClass"
            />
            <ErrorMessage name="email" class="text-red-500" />
          </div>

          <!-- Schedule Dropdown -->
          <div class="mb-4">
            <label class="block font-semibold mb-1" :class="selectClass"
              >Select a Schedule</label
            >
            <Field
              as="select"
              name="selectedSchedule"
              class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              :class="selectClass"
              v-model="selectedSchedule"
            >
              <option
                v-for="schedule in movie.movie_schedules"
                :key="schedule.schedule_id"
                :value="schedule"
                :class="selectClass"
              >
                <span>
                  {{
                    formatScheduleDateTime(
                      schedule.schedule_date,
                      schedule.start_time,
                      schedule.end_time
                    )
                  }}
                  <span
                    v-if="isScheduleExpired(schedule)"
                    class="text-red-500 font-semibold mr-2"
                    :class="selectClass"
                  >
                    (Expired)
                  </span>
                  <span
                    v-else
                    class="text-green-500 font-semibold mr-2"
                    :class="selectClass"
                  >
                    (Valid)
                  </span>
                  (Hall: {{ schedule.cinema_hall }})
                </span>
              </option>
            </Field>
            <ErrorMessage name="selectedSchedule" class="text-red-500" />
          </div>

          <!-- Ticket Quantity Field -->
          <div class="mb-4">
            <label
              for="ticketQuantity"
              class="block font-semibold mb-2"
              :class="selectClass"
              >Number of Tickets</label
            >
            <Field
              name="ticketQuantity"
              v-model.number="ticketQuantity"
              type="number"
              id="ticketQuantity"
              min="1"
              class="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              :class="selectClass"
            />
            <ErrorMessage name="ticketQuantity" class="text-red-500" />
          </div>

          <div class="ticket-price text-lg font-semibold" :class="selectClass">
            Total Price: <span :class="selectClass">{{ totalPrice }}</span> birr
          </div>
          <div v-if="islengthGreater">
            <button
              type="submit"
              class="flex items-center px-10 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-400 transition"
            >
              <span v-if="isloading" class="flex items-center justify-center">
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
              <span
                v-else
                class="flex items-center rounded-md shadow-md transition-all text-white font-medium"
                >Pay <ArrowRightIcon class="h-5 w-5 ml-2"
              /></span>
            </button>
          </div>
        </Form>
      </div>
      <div v-else>
        <p>
          No available schedules. Please choose a scheduled movie. Thank you!
        </p>
      </div>
    </div>
  </div>
</template>
