package router

type Customer struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// maybe create  methods to update the data
