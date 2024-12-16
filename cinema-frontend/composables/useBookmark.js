import { jwtDecode } from "jwt-decode";
import { useNuxtApp } from "#app";
import { useToast } from "vue-toastification";
import { DELETE_BOOKMARKS } from "../components/graphql/mutations/DELETE_BOOKMARKS.graphql";
import { ADD_BOOKMARKS } from "../components/graphql/mutations/ADD_BOOKMARKS.graphql";
import { GET_BOOKMARKS } from "../components/graphql/queries/GET_BOOKMARKS.graphql";
export const useBookmarks = () => {
  const { $apollo } = useNuxtApp();
  const bookmarkedMovies = ref([]);
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
    try {
      const response = await $apollo.query({
        query: GET_BOOKMARKS,
        variables: { userId: userId },
        fetchPolicy: "network-only",
      });
      bookmarks.value = response.data.bookmarks;
      bookmark.setBookmarks(response.data.bookmarks);
      bookmarkedMovies.value = response.data.bookmarks.map(
        (bookmark) => bookmark.movie_id
      );
      isBookmarkExist.value = bookmarks.value.length === 0;
      isBookmarked.value = bookmarks.value.length === 0;
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
    try {
      const response = await $apollo.mutate({
        mutation: DELETE_BOOKMARKS,
        variables: { userId: userId, movieId: movieId },
        refetchQueries: [
          {
            query: GET_BOOKMARKS,
            variables: { userId: userId },
          },
        ],
      });

      if (response.data.delete_bookmarks.affected_rows > 0) {
        isBookmarked.value = true;
        isBookmarkExist.value = false;
        bookmarks.value = bookmarks.value.filter(
          (movie) => movie.movie_id !== movieId
        );
        bookmark.setBookmarks(
          bookmark.bookmarks.filter((bookmark) => bookmark.movie_id !== movieId)
        );
        toast.success("Bookmarks removed successfully", {
          position: window.innerWidth < 768 ? "top-center" : "top-right",
          timeout: 1000,
          className: `
            bg-green-500 dark:bg-green-700 text-white 
            font-medium 
            ${
              window.innerWidth < 768
                ? "text-[10px] p-1 max-w-[150px]"
                : "text-sm p-3 max-w-[300px]"
            } 
            rounded-md shadow-md
          `,
        });
      } else {
        console.log("No bookmark found to delete.");
      }
    } catch (error) {
      console.error("Errors while deleting Bookmarks:", error);
    }
  };
  const isBookmarkedCheck = (movieId) => {
    return bookmarkedMovies.value.includes(movieId);
  };
  const toggleBookmark = async (movieId) => {
    const accessToken = localStorage.getItem("accessToken");
    if (!accessToken) {
      return toast.error("You must be logged in to toggle bookmarks.", {
        position: "top-center",
      });
    }

    const decodedToken = jwtDecode(accessToken);
    const userId =
      decodedToken["https://hasura.io/jwt/claims"]["x-hasura-user-id"];

    try {
      if (isBookmarkedCheck(movieId)) {
        await removeBookmarks(movieId);
        bookmarkedMovies.value = bookmarkedMovies.value.filter(
          (id) => id !== movieId
        );
      } else {
        await $apollo.mutate({
          mutation: ADD_BOOKMARKS,
          variables: {
            UserId: userId,
            movieId: movieId,
          },
        });
        bookmarkedMovies.value.push(movieId);
        toast.success("Bookmark added successfully!", {
          position: "top-center",
        });
      }
    } catch (error) {
      toast.error("An error occurred while processing your request.", {
        position: "top-center",
      });
    }
  };
  return {
    toggleBookmark,
    isBookmarked,
    removeBookmarks,
    getBookmarks,
    bookmarks,
    loading,
    isBookmarkedCheck,
    isBookmarkExist,
  };
};
