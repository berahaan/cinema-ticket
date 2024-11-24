import { useRouter } from "vue-router";
export const useLogout = () => {
  const router = useRouter();
  const Logout = () => {
    const color = useColorModeStore();
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
    document.documentElement.classList.remove("dark");
    color.resetColorMode();
    router.push("/");
  };
  return {
    Logout,
  };
};
