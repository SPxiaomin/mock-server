package service

import "context"

type AppServiceInterface interface {
	GetApps(c context.Context) (string, error)
	GetGMs(c context.Context) (string, error)
	GetGroups(c context.Context) (string, error)
	GetUserInfo(c context.Context) (string, error)
}
