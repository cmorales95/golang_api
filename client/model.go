package main

type (
	//Login .....
	Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GeneralResponse struct {
		MessageType string `json:"message_type"`
		Message     string `json:"message"`
	}

	LoginResponse struct {
		GeneralResponse
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	Community struct {
		Name string `json:"name"`
	}

	Communities []Community

	//Person .....
	Person struct {
		Name string `json:"name"`
		Age uint8 `json:"age"`
		Communities Communities `json:"communities"`
	}
)
