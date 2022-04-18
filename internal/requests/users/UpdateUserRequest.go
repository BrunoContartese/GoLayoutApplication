package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/go-errors"
)

func HandleUpdateUserRequest(request *fiber.Request) (bool, int, error) {
	if authorizeUpdate(request) {
		_, err := validateUpdate(request)

		if err != nil {
			return false, 422, err
		}

		return true, 200, nil

	} else {

		return false, 401, errors.New("You are unauthorized to perform this action.").Format("StoreUserRequest")

	}

}

/*Handle user authorization*/
func authorizeUpdate(request *fiber.Request) bool {
	return true
}

func validateUpdate(request *fiber.Request) (bool, error) {

	/*Acá pensar como utilizar fiber validation*/

	return true, nil
}

func updateMessages() {

}

func updateRules() {

}
