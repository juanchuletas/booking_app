package models

// Holds data send to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //For data that can be structs  thal hold other data
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
