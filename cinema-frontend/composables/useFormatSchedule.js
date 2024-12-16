export const useFormatSchedule = () => {
  const isvalid = ref(false);
  const isNeedTotalPrice =ref(false)
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
    const now = new Date();
    const currentTimeDb =new Date(schedule.end_time)
    currentTimeDb>now ? isNeedTotalPrice.value =true: isNeedTotalPrice.value =false
    return  currentTimeDb>now
  };
  // Function to parse `schedule_date` and `time` into a valid Date object

  return {
    isScheduleExpired,
    formatScheduleDateTime,
    isvalid,
    isNeedTotalPrice
  };
};
