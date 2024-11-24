package models

type ImageUploadRequest struct {
	FeaturedImage string   `json:"featuredImage"`
	OtherImages   []string `json:"otherImages"`
}
type ProfileImageInput struct {
	ProfileImage string `json:"profileImages"`
}
