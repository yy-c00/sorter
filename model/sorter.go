package model

//Numbers struct used to organize numbers
type Numbers struct {
	//Unsorted contains unsorted numbers
	Unsorted []int `json:"unsorted,omitempty"`
	//Sorted contains sorted numbers
	Sorted []int `json:"sorted,omitempty"`
}

//Sorter interface implemented by all structs to sort information
type Sorter interface {
	Sort(*Numbers)
}
