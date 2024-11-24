export function useModals() {
  const currentImageIndex = ref(0);
  const currentOtherImages = ref([]);
  const isModalOpen = ref(false);
  const isNotherImagesExist = ref(false);
  const openModal = (movieImages) => {
    const otherImages = movieImages.map((image) => image.other_images);
    const flattenedImages = otherImages.flat();
    if (flattenedImages.length > 0) {
      currentOtherImages.value = flattenedImages;
      currentImageIndex.value = 0;
      isNotherImagesExist.value = false;
    } else {
      isNotherImagesExist.value = true;
    }
    isModalOpen.value = true;
  };
  const nextImage = () => {
    if (currentImageIndex.value < currentOtherImages.value.length - 1) {
      currentImageIndex.value++;
    } else {
      currentImageIndex.value = 0; // Loop back to the first image
    }
  };

  // Function to go to the previous image
  const prevImage = () => {
    if (currentImageIndex.value > 0) {
      currentImageIndex.value--;
    } else {
      currentImageIndex.value = currentOtherImages.value.length - 1; // Loop to the last image
    }
  };
  const closeModal = () => {
    isModalOpen.value = false;
    isNotherImagesExist.value = false;
  };
  return {
    closeModal,
    prevImage,
    isNotherImagesExist,
    nextImage,
    openModal,
    isModalOpen,
    currentImageIndex,
    currentOtherImages,
  };
}
