import { jwtDecode } from "jwt-decode";
import { useNuxtApp } from "#app";
import { useToast } from "vue-toastification";
import { useRouter } from "vue-router";
import { GET_BOOKMARK_U_INFO } from "../components/graphql/queries/GET_BOOKMARK_U_INFO.graphql";
import { DELETE_BOOKMARKS } from "../components/graphql/mutations/DELETE_BOOKMARKS.graphql";
import { ADD_BOOKMARKS } from "../components/graphql/mutations/ADD_BOOKMARKS.graphql";
import { GET_BOOKMARKS } from "../components/graphql/queries/GET_BOOKMARKS.graphql";
export const useBookmarks = () => {
  const { $apollo } = useNuxtApp();
  const router = useRouter();
  const bookmark = useBookmarkStore();
  const toast = useToast();
  const isBookmarked = ref(false);
  const isBookmarkExist = ref(false);
  const loading = ref(true);
  const bookmarks = ref([]);
  const getBookmarks = async () => {
    const decodedToken = jwtDecode(localStorage.getItem("accessToken"));
    const userId =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
    console.log("user id now for regular users ", userId);

    try {
      const response = await $apollo.query({
        query: GET_BOOKMARKS,
        variables: { userId: userId },
        fetchPolicy: "network-only",
      });
      bookmarks.value = response.data.bookmarks;
      bookmark.setBookmarks(response.data.bookmarks);

      isBookmarkExist.value = bookmarks.value.length === 0;
    } catch (error) {
      console.error("Error fetching bookmarks:", error);
    } finally {
      loading.value = false;
    }
  };
  const removeBookmarks = async (movieId) => {
    const token = localStorage.getItem("accessToken");
    if (!token) return alert("You must be logged in to delete bookmarks.");
    const decodedToken = jwtDecode(token);
    const userId =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
    console.log("user id here", userId);
    try {
      const response = await $apollo.mutate({
        mutation: DELETE_BOOKMARKS,
        variables: { userId: userId, movieId: movieId },
        refetchQueries: [
          {
            query: GET_BOOKMARK_U_INFO,
            variables: { UserId: userId, movieId: movieId },
          },
        ],
      });

      if (response.data.delete_bookmarks.affected_rows > 0) {
        isBookmarked.value = false;
        console.log("Bookmark successfully removed from list.");
        toast.success("Bookmarks removed successfully ", {
          position: "top-right",
          duration: 4000,
        });
      } else {
        console.log("No bookmark found to delete.");
      }
    } catch (error) {
      console.error("Errors while deleting Bookmarks:", error);
    }
  };

  const getBookmarkDetail = () => {
    console.log("Clicked here ");
    const accessToken = localStorage.getItem("accessToken");
    const decodedToken = jwtDecode(accessToken);
    const defaultRole =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-default-role"];
    if (defaultRole === "admin") {
      console.log("Role for profiles is admin here ");
      router.push("/admin/Bookmark");
    } else {
      console.log("Role for Profiles is regular here");
      router.push("/user/Bookmark");
    }
  };

  const toggleBookmark = async (movieId) => {
    console.log("Selected Movie ID for Bookmarking:", movieId);
    const decodedToken = jwtDecode(localStorage.getItem("accessToken"));
    const userId =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

    try {
      // Check if the bookmark already exists
      const queryResponse = await $apollo.query({
        query: GET_BOOKMARK_U_INFO,
        variables: { userId, movieId },
        fetchPolicy: "network-only",
      });

      const existingBookmark = queryResponse.data.bookmarks;
      if (existingBookmark && existingBookmark.length > 0) {
        console.log("Bookmark exist and try to remove it now ");
        await removeBookmarks(movieId);
      } else {
        const mutationResponse = await $apollo.mutate({
          mutation: ADD_BOOKMARKS,
          variables: { UserId: userId, movieId: movieId },
          refetchQueries: [
            {
              query: GET_BOOKMARK_U_INFO,
              variables: { UserId: userId, movieId: movieId },
            },
          ],
        });

        if (mutationResponse && mutationResponse.data) {
          isBookmarked.value = true;
          toast.success("Bookmark added successfully!", {
            position: "top-center",
            timeout: 5000,
          });
        }
      }
    } catch (error) {
      console.error("Error handling bookmark:", error);
    }
  };

  return {
    toggleBookmark,
    isBookmarked,
    removeBookmarks,
    getBookmarks,
    bookmarks,
    getBookmarkDetail,
    loading,
    isBookmarkExist,
  };
};
