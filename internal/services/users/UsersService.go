package users

import "layoutapp/internal/models"

func Index() []models.User {
	var users []models.User

	users = append(users, models.User{
		ID:        "asdfioasdf-sdafasdf-123",
		Username:  "bcontartese",
		Password:  "estoesunpassword",
		FirstName: "Bruno",
		LastName:  "Contartese",
		CreatedAt: "2022-04-17",
		UpdatedAt: "2022-04-17",
		DeletedAt: "",
	})

	return users
}
