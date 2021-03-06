package actions

import (
	"github.com/getfider/fider/app"
	"github.com/getfider/fider/app/models"
	"github.com/getfider/fider/app/pkg/validate"
)

// Actionable is any action that the user can perform using the web app
type Actionable interface {
	Initialize() interface{}
	Validate(services *app.Services) *validate.Result
	IsAuthorized(user *models.User) bool
}
