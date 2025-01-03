import { UPDATE_MOVIE_NO_OTHER } from "../components/graphql/mutations/UPDATE_NO_OTHER.graphql";
import { UPDATE_MOVIE_NO_IMAGE } from "../components/graphql/mutations/UPDATE_NO_IMAGE.graphql";
import { UPDATE_OTHER_IMAGE } from "../components/graphql/mutations/UPDATE_OTHER_IMAGE.graphql";
import { useNuxtApp } from "#app";
import { useToast } from "vue-toastification";
export const useUpdateMovie = () => {
  const isUpdated = ref("");
  const toast = useToast();
  const isUpdateloading = ref(false);
  const loading = ref(false);
  const { $apollo } = useNuxtApp();
  const updateMovies = async (
    movie_id,
    title,
    description,
    director_id,
    genres,
    stars,
    featuredImageURL,
    otherImageURLs,
    duration
  ) => {
    if (featuredImageURL == null && otherImageURLs == null) {
      try {
        const response = await $apollo.mutate({
          mutation: UPDATE_MOVIE_NO_IMAGE,
          variables: {
            movie_id,
            title,
            description,
            director_id,
            genres: genres.map((genre_id) => ({ movie_id, genre_id })),
            stars: stars.map((star_id) => ({ movie_id, star_id })),
            duration,
          },
        });

        if (response && response.data) {
          toast.success("Movies without Image change Updated successfully", {
            position: "top-center",
            duration: 2000,
          });
        } else {
          isUpdated.value = "Error while updating man !!";
        }
        return response.data;
      } catch (error) {
        throw error;
      } finally {
        isUpdateloading.value = false;
        loading.value = false;
      }
    } else if (featuredImageURL != null && otherImageURLs == null) {
      try {
        const response = await $apollo.mutate({
          mutation: UPDATE_MOVIE_NO_OTHER,
          variables: {
            movie_id,
            title,
            description,
            director_id,
            genres: genres.map((genre_id) => ({ movie_id, genre_id })),
            stars: stars.map((star_id) => ({ movie_id, star_id })),
            featured_images: featuredImageURL,
            duration,
          },
        });

        if (response && response.data) {
          toast.success(
            "Movies with only featuredImages is Update along with other Fields",
            {
              position: "top-center",
              duration: 2000,
            }
          );
        } else {
          isUpdated.value = "Error while updating man !!";
        }
        return response.data;
      } catch (error) {
        console.error("Error updating movie:", error);
        throw error;
      } finally {
        isUpdateloading.value = false;
        loading.value = false;
      }
    } else if (featuredImageURL == null && otherImageURLs != null) {
      // alert("Now only other images is need to be updated ");

      const movieImagesData = otherImageURLs.map((imageUrl) => ({
        movie_id,
        other_images: imageUrl,
      }));
      try {
        const response = await $apollo.mutate({
          mutation: UPDATE_OTHER_IMAGE,
          variables: {
            movie_id,
            title,
            description,
            director_id,
            genres: genres.map((genre_id) => ({ movie_id, genre_id })),
            stars: stars.map((star_id) => ({ movie_id, star_id })),
            movie_images: movieImagesData,
            duration,
          },
        });

        if (response && response.data) {
          toast.success(
            "Movies with otherImages changes is Update along with other fields successfully",
            {
              position: "top-center",
              duration: 2000,
            }
          );
        } else {
          isUpdated.value = "Error while updating movies  !!";
        }
        return response.data;
      } catch (error) {
        throw error;
      } finally {
        isUpdateloading.value = false;
        loading.value = false;
      }
    } else {
      if (featuredImageURL && otherImageURLs) {
        const movieImagesData = otherImageURLs.map((imageUrl) => ({
          movie_id,
          other_images: imageUrl,
        }));
        isUpdateloading.value = true;
        loading.value = true;
        try {
          const response = await $apollo.mutate({
            mutation: UPDATE_MOVIE_NO_OTHER,
            variables: {
              movie_id,
              title,
              description,
              director_id,
              genres: genres.map((genre_id) => ({ movie_id, genre_id })),
              stars: stars.map((star_id) => ({ movie_id, star_id })),
              featured_images: featuredImageURL,
              movie_images: movieImagesData,
              duration,
            },
          });

          if (response && response.data) {
            toast.success("Movies Updated successfully ", {
              position: "top-center",
              duration: 2000,
            });
          } else {
            isUpdated.value = "Error while updating man !!";
          }
          return response.data;
        } catch (error) {
          throw error;
        } finally {
          isUpdateloading.value = false;
          loading.value = false;
        }
      }
    }
  };
  return {
    updateMovies,
    loading,
    isUpdated,
    isUpdateloading,
  };
};
