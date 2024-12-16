// This function uploads images and retrieves URLs to send to Hasura
import { useNuxtApp } from "#app";
import { INSERT_IMAGE } from "../components/graphql/mutations/INSERT_IMAGE.graphql";
import { SEND_PROFILE_PHOTO } from "../components/graphql/mutations/SEND_PROFILE_PHOTO.graphql";
import { useToast } from "vue-toastification";
export function HandleImageUpload() {
  const imageStore = useImageStore();
  const toast = useToast();
  const { $apollo } = useNuxtApp();
  const UploadImage = async (featuredImage, otherImages) => {
    if (featuredImage && otherImages) {
      if (!Array.isArray(otherImages) || otherImages.length === 0) {
        toast.error("Invalid otherImages: must be a non-empty array", {
          position: "top-right",
          timeout: 1000,
        });
        return;
      }
      try {
        const response = await $apollo.mutate({
          mutation: INSERT_IMAGE,
          variables: {
            featuredImage,
            otherImages,
          },
        });

        if (response && response.data) {
          toast.success(response.data.fileUpload.message, {
            position: "top-center",
            duration: 2000,
          });

          imageStore.setImageURLs(
            response.data.fileUpload.featuredurl,
            response.data.fileUpload.otherurl
          );
        }
      } catch (error) {
        console.error("Error while uploading images:", error);
      }
    } else if (featuredImage != null && otherImages == null) {
      const responseOther = await $apollo.mutate({
        mutation: INSERT_IMAGE,
        variables: {
          featuredImage,
          otherImages: [],
        },
      });

      if (responseOther && responseOther.data) {
        toast.success("feature image successfully uploaded.", {
          position: "top-center",
          duration: 2000,
        });
        imageStore.setImageURLs(
          responseOther.data.fileUpload.featuredurl,
          null
        );
      }
    } else if (featuredImage == null && otherImages != null) {
      const responseOther = await $apollo.mutate({
        mutation: INSERT_IMAGE,
        variables: {
          featuredImage: null,
          otherImages,
        },
      });

      if (responseOther && responseOther.data) {
        toast.success("Other images uploaded successfully", {
          position: "top-center",
          duration: 2000,
        });
        imageStore.setImageURLs(null, responseOther.data.fileUpload.otherurl);
      }
    } else if (otherImages == null && featuredImage) {
      const responseOther = await $apollo.mutate({
        mutation: INSERT_IMAGE,
        variables: {
          featuredImage,
          otherImages,
        },
      });
    } else {
      console.warn("No images provided for upload.");
    }
  };
  const uploadProfile = async (profilePhoto) => {
    if (!profilePhoto || profilePhoto.length === 0) {
      toast.error("Invalid featuredImage (Profile Photo)", {
        position: "top-right",
        timeout: 2000,
      });
      return;
    }
    try {
      const profileResponse = await $apollo.mutate({
        mutation: SEND_PROFILE_PHOTO,
        variables: { profilePhoto },
      });

      imageStore.setImageURLs(
        profileResponse.data.profileUpload.profileurl,
        null
      );
    } catch (error) {
      console.error("Error while uploading profile photo:", error);
    }
  };
  return {
    UploadImage,
    uploadProfile,
  };
}
