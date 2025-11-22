package service

import "context"

type IAdminService interface {
	HelloWorld(c context.Context) string
}
