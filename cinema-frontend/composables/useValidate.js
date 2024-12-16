import * as yup from "yup";
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
  const validationSchema = yup.object({
    name: yup
      .string()
      .required("name is required")
      .min(3, "Star name should be atleaset 3 characters"),
    first_name: yup.string().required("First Name is required"),
    last_name: yup.string().required("Last Name is required"),
    phone_number: yup
      .string()
      .matches(/^(09|07)\d{8}$/, "Phone number is not valid")
      .required("Phone Number is required"),
    email: yup
      .string()
      .email("Email must be valid")
      .matches(
        /^[a-zA-Z0-9._%+-]+@(gmail|yahoo)\.com$/,
        "Email must be a valid Gmail or Yahoo address "
      )
      .required("Email is required"),
    selectedSchedule: yup.object().required("Please select a schedule"),
    ticketQuantity: yup
      .number()
      .typeError("Ticket number is required ")
      .required("At least one ticket is required")
      .integer("Ticket quantity must be a whole number")
      .min(1, "Ticket quantity must be at least 1"),
  });
  return {
    validateDescription,
    validateDirector,
    validateFeaturedImages,
    validateDuration,
    validateOtherImages,
    validateGenre,
    validationSchema,
    validateTitle,
    validateStar,
  };
};
