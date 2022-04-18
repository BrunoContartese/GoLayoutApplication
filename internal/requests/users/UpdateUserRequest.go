package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/go-errors"
	"github.com/thedevsaddam/govalidator"
)

func Handle(request *fiber.Request) (bool, int, error) {
	if authorize(request) {
		_, err := validate(request)

		if err != nil {
			return false, 422, err
		}

		return true, 200, nil

	} else {

		return false, 401, errors.New("You are unauthorized to perform this action.").Format("StoreUserRequest")

	}

}

/*Handle user authorization*/
func authorize(request *fiber.Request) bool {
	return true
}

func validate(request *fiber.Request) (bool, error) {
	return true, nil
}

func messages() govalidator.MapData {
	return govalidator.MapData{
		"first_name": []string{"required: Debe ingresar el nombre del usuario", "string: El nombre del usuario no es válido."},
		"last_name":  []string{"required: Debe ingresar el apellido del usuario", "string: El apellido del usuario no es válido."},
		"email":      []string{"required: Debe ingresar el email del usuario", "email: El email no tiene un formato válido."},
	}
}

func rules() govalidator.MapData {
	return govalidator.MapData{
		"first_name": []string{"required", "string"},
		"last_name":  []string{"required", "string"},
		"email":      []string{"required", "email"},
		"password":   []string{"required", "string"},
	}
}
