export const useRefresh = () => {
  const refreshPage = async (page) => {
    await page();
  };
  return {
    refreshPage,
  };
};
