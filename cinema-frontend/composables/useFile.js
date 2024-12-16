import { ref } from "vue";
import { useToast } from "vue-toastification";
export function useFile() {
  const previewFeaturedImage = ref(null);
  const previewOtherImages = ref([]);
  const featuredImagesBase64 = ref("");
  const otherImagesBase64 = ref([]);
  const currentFeatureImage = ref(null);
  const imageStore = useImageBase64Store();
  const profile = useProfileImageBase64Store();
  const toast = useToast();

  function convertImageToBase64(imageFile) {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onloadend = () => resolve(reader.result);
      reader.onerror = (error) => reject(error);
      reader.readAsDataURL(imageFile);
    });
  }
  const handleFeaturedImageUpload = async (event) => {
    const file = event.target.files[0];
    if (file) {
      featuredImagesBase64.value = await convertImageToBase64(file);
      if (featuredImagesBase64) {
        toast.success("feature images successfully selected", {
          position: "top-center",
          timeout: 1000,
        });
      }
      profile.setProfileImageBase64(featuredImagesBase64.value);
      imageStore.setImageBase64(
        featuredImagesBase64.value,
        imageStore.otherImage64
      );
    }
  };
  // Handle upload for other images
  const handleOtherImagesUpload = async (event) => {
    // this function purpose is to make ready the images to be sent to a hasura for other process
    const files = event.target.files;
    const selectedFiles = [...files];
    if (selectedFiles.length > 5) {
      toast.warning("You can upload a maximum of 5 images ", {
        position: "top-right",
        duration: 500,
      });
      return;
    }
    // Convert each selected image to Base64 and update previewOtherImages and otherImages arrays
    const base64Images = await Promise.all(
      selectedFiles.map((file) => convertImageToBase64(file))
    );
    previewOtherImages.value = selectedFiles.map((file) =>
      URL.createObjectURL(file)
    );
    if (base64Images) {
      toast.success("other images successfully selected", {
        position: "top-center",
        timeout: 1000,
      });
      imageStore.setImageBase64(imageStore.featuredImage64, base64Images);
      otherImagesBase64.value = base64Images;
    }
  };

  return {
    handleFeaturedImageUpload,
    handleOtherImagesUpload,
    previewFeaturedImage,
    currentFeatureImage,
    previewOtherImages,
  };
}
