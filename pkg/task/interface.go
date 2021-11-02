package task

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/infraboard/cmdb/pkg/resource"
)

var (
	validate = validator.New()
)

type Service interface {
	CreatTask(context.Context, *CreateTaskRequst) (*Task, error)
}

func NewCreateTaskRequst() *CreateTaskRequst {
	return &CreateTaskRequst{}
}

type CreateTaskRequst struct {
	SecretId     string `validate:"required,lte=100"`
	Region       string
	ResourceType resource.Type
}

func (req *CreateTaskRequst) Validate() error {
	return validate.Struct(req)
}
