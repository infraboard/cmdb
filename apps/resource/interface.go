package resource

import "gorm.io/gorm"

type Service interface {
	PutResource(*gorm.DB, *ResourceSet) error
	DeleteResource(*gorm.DB, *ResourceSet) error
	RPCServer
}
