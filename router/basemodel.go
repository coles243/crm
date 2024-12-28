package router

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"Phone"`
	Contacted bool   `json:"contacted"`
}

// maybe create  methods to update the data
