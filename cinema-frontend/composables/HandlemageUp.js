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
      console.log("Both are selected ..");
      if (!Array.isArray(otherImages) || otherImages.length === 0) {
        console.error("Invalid otherImages: must be a non-empty array");
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
          console.log(
            "response.data.featuredurl:",
            response.data.fileUpload.featuredurl,
            "Other image urls is ",
            response.data.fileUpload.otherurl
          );
          imageStore.setImageURLs(
            response.data.fileUpload.featuredurl,
            response.data.fileUpload.otherurl
          );
        }
      } catch (error) {
        console.error("Error while uploading images:", error);
      }
    } else if (featuredImage != null && otherImages == null) {
      console.log("Here featured Images only is uploaded now.....");
      const responseOther = await $apollo.mutate({
        mutation: INSERT_IMAGE,
        variables: {
          featuredImage,
          otherImages: [],
        },
      });
      console.log(
        "Response after inserting featured Image from Golang backned ",
        responseOther.data.fileUpload.featuredurl
      );
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
      console.log(
        "Response after inserting otherImage in backned ",
        responseOther.data.fileUpload.otherurl
      );
      if (responseOther && responseOther.data) {
        toast.success("Other images uploaded successfully", {
          position: "top-center",
          duration: 2000,
        });
        imageStore.setImageURLs(null, responseOther.data.fileUpload.otherurl);
      }
    } else if (otherImages == null && featuredImage) {
      console.log("Here other images is selected now.....");
      const responseOther = await $apollo.mutate({
        mutation: INSERT_IMAGE,
        variables: {
          featuredImage,
          otherImages,
        },
      });
      console.log(
        "Response after inserting otherImage in backned ",
        responseOther.data
      );
    } else {
      console.warn("No images provided for upload.");
    }
  };
  const uploadProfile = async (profilePhoto) => {
    console.log(
      "Only profile photo is being uploaded and featured Images to be sent here is:::",
      profilePhoto
    );
    if (!profilePhoto || profilePhoto.length === 0) {
      console.error("Invalid featuredImage (Profile Photo):");
      return;
    }
    try {
      const profileResponse = await $apollo.mutate({
        mutation: SEND_PROFILE_PHOTO,
        variables: { profilePhoto },
      });
      console.log("Profile Response:", profileResponse.data);
      console.log(
        "Profile Response:",
        profileResponse.data.profileUpload.profileurl
      );
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
