export const useValidate = () => {
  const validateTitle = (values) => {
    if (!values) return "Title is required";
    if (values.length < 4) {
      return "Please make a title at least 4 characters";
    }
    return true;
  };
  const validateDescription = (value) => {
    if (!value) {
      return "Description is required";
    }
    if (value.length < 10) {
      return "Description must be at least 10 characters long";
    }
    return true;
  };
  const validateDuration = (value) => {
    if (!value) {
      return "Duration is required";
    }
    if (value <= 0) {
      return "Duration must be a positive number";
    }
    return true; // Validation passed
  };

  const validateDirector = (value) => {
    if (!value) {
      return "Please choose Directors here";
    }
    return true; // Validation passed
  };
  const validateStar = (value) => {
    console.log("Validating star:", value); // Debugging log
    if (!value || value === "") {
      return "Please choose a star";
    }
    return true;
  };
  const validateFeaturedImages = (value) => {
    // Check if value is provided
    if (!value || value.length === 0) {
      return "Please choose an image here";
    }

    // Check if more than one image is selected
    if (value.length > 1) {
      return "Please choose only one image";
    }

    // Validate the file type if needed (optional)

    return true;
  };

  const validateOtherImages = (value) => {
    // Check if no value is provided
    if (!value || value.length === 0) {
      return "Please choose at least one image here";
    }

    // Check if more than two images are selected
    if (value.length > 5) {
      return "Please choose only up to 5 images";
    }

    // Validate file types if needed (optional)
    const validExtensions = ["jpg", "jpeg", "png", "gif"]; // Add valid extensions as necessary
    for (let i = 0; i < value.length; i++) {
      const extension = value[i].name.split(".").pop().toLowerCase();
      if (!validExtensions.includes(extension)) {
        return "Please choose valid image files (jpg, jpeg, png, gif)";
      }
    }

    // If all checks pass
    return true;
  };
  const validateGenre = (value) => {
    if (!value || value === "") {
      return "Please choose Genre";
    }
    return true;
  };
  return {
    validateDescription,
    validateDirector,
    validateFeaturedImages,
    validateDuration,
    validateOtherImages,
    validateGenre,
    validateTitle,
    validateStar,
  };
};
