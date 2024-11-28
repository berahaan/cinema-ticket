export const useFormatSchedule = () => {
  const isvalid = ref(false);
  const formatScheduleDateTime = (date, startTime, endTime) => {
    const scheduleDate = new Date(date);
    const start = new Date(startTime);
    const end = new Date(endTime);
    const optionsDate = {
      weekday: "long",
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
    };
    const optionsTime = { hour: "numeric", minute: "2-digit", hour12: true };
    const formattedDate = scheduleDate.toLocaleDateString(
      undefined,
      optionsDate
    );
    const formattedStartTime = start.toLocaleTimeString(undefined, optionsTime);
    const formattedEndTime = end.toLocaleTimeString(undefined, optionsTime);

    return `${formattedDate} from ${formattedStartTime} to ${formattedEndTime}`;
  };
  const isScheduleExpired = (schedule) => {
    console.log("isSchedules called now .....");
    const now = new Date();
    const scheduleEndDateTime = parseScheduleDateTime(
      schedule.schedule_date,
      schedule.end_time
    );
    scheduleEndDateTime < now
      ? (isvalid.value = false)
      : (isvalid.value = true);
    console.log(
      "Parded going to be check here ischeduleExpired Func ",
      scheduleEndDateTime < now
    );
    return scheduleEndDateTime < now; // Return true if the schedule has expired
  };

  // Function to parse `schedule_date` and `time` into a valid Date object
  const parseScheduleDateTime = (scheduleDate, time) => {
    if (!scheduleDate || !time) {
      return new Date(NaN); // Return an invalid date to prevent unexpected behavior
    }

    // Check if `time` is already a full ISO date-time string
    if (time.includes("T")) {
      return new Date(time); // Parse directly if it's an ISO string
    }
    const combinedDateTime = `${scheduleDate} ${time}`;
    const parsedDate = new Date(combinedDateTime);

    if (isNaN(parsedDate)) {
      console.error("Failed to parse combined datetime", { combinedDateTime });
    }

    return parsedDate;
  };
  return {
    isScheduleExpired,
    formatScheduleDateTime,
    isvalid,
  };
};
