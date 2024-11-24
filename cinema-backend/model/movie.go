package models

type FeatureMovieDeletePayload struct {
	Event struct {
		Data struct {
			Old struct {
				FeaturedImage string `json:"featured_images"`
			} `json:"old"`
		} `json:"data"`
	} `json:"event"`
}
type OtherMovieImagesDeletePayload struct {
	Event struct {
		Data struct {
			Old struct {
				OtherImage string `json:"other_images"` // single image URL
			} `json:"old"`
		} `json:"data"`
	} `json:"event"`
}
