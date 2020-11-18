package models

//Image is the struct of the assets of each profile
type Image struct {
	IDImage int64  `json:"idImage"`
	Path    string `json:"path"`
	Name    string `json:"name"`
}

//Images list of images
type Images []Image
