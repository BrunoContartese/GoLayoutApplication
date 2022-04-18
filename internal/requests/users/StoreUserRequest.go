package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kataras/go-errors"
)

func HandleStoreUserRequest(request *fiber.Request) (bool, int, error) {
	if authorizeStore(request) {
		_, err := validateStore(request)

		if err != nil {
			return false, 422, err
		}

		return true, 200, nil

	} else {

		return false, 401, errors.New("You are unauthorized to perform this action.").Format("StoreUserRequest")

	}

}

/*Handle user authorization*/
func authorizeStore(request *fiber.Request) bool {
	return true
}

func validateStore(request *fiber.Request) (bool, error) {
	/*Acá pensar como utilizar fiber validation*/

	return true, nil
}

func storeMessages() {

}

func storeRules() {

}
