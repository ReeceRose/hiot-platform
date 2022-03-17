package types

//StandardResponse is used to standardize the returns for all HTTP endpoint
type StandardResponse struct {
	Data       interface{}
	StatusCode int
	Error      string
	Success    bool
}
