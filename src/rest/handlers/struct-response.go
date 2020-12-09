package handlers

type ResponseData struct {
	Hash	string
	Level	float64
	Skills	[]float64
}

type ResponseJson struct {
	Status	bool
	Message	string
	Data	ResponseData
}
