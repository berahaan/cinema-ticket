import { useRouter } from "vue-router";
export const useGobackArrow = () => {
  const router = useRouter();
  const goBack = () => {
    router.back();
  };
  return {
    goBack,
  };
};
