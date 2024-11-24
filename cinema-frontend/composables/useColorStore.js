export const useThemeColor = () => {
  const colorStore = useColorModeStore();
  const selectClass = computed(() =>
    colorStore.colorMode === "dark"
      ? "bg-gray-800 text-white dark:bg-gray-800"
      : "bg-gray-50 text-gray-900 "
  );
  const optionClass = computed(() =>
    colorStore.colorMode === "dark"
      ? "bg-gray-800 text-white"
      : "bg-gray-100 text-gray-900"
  );
  const deleteOption = computed(() =>
    colorStore.colorMode === "dark"
      ? "bg-gray-800 text-white border-gray-700"
      : "bg-gray-50 text-gray-900 border-gray-300"
  );
  const labelClass = computed(() =>
    colorStore.colorMode === "dark"
      ? "text-white font-semibold"
      : "text-gray-900 font-semibold"
  );
  return {
    selectClass,
    labelClass,
    deleteOption,
    optionClass,
  };
};
