import { defineStore } from "pinia";

export const useProfileImageBase64Store = defineStore(
  "useProfileImageBase64Store",
  {
    state: () => ({
      profileImage64: "",
    }),
    actions: {
      setProfileImageBase64(profileImage) {
        console.log(
          "Here profile image is being setted in store ",
          profileImage
        );
        this.profileImage64 = profileImage;
      },
    },
  }
);
