package impl

import (
	"context"

	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CreateSecret(ctx context.Context, req *secret.CreateSecretRequest) (
	*secret.Secret, error) {
	ins, err := secret.NewSecret(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create secret error, %s", err)
	}

	// 入库之前先加密
	if err := ins.Spec.EncryptAPISecret(conf.C().App.EncryptKey); err != nil {
		s.log.Warnf("encrypt api key error, %s", err)
	}

	if err := s.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) QuerySecret(ctx context.Context, req *secret.QuerySecretRequest) (
	*secret.SecretSet, error) {
	query := s.db.WithContext(ctx)
	if req.Keywords != "" {
		query = query.Where("description LIKE ? OR api_key = ?",
			"%"+req.Keywords+"%",
			req.Keywords,
		)
	}

	set := secret.NewSecretSet()
	offset := req.Page.ComputeOffset()
	err := query.
		Order("create_at DESC").
		Limit(int(req.Page.PageSize)).
		Offset(int(offset)).
		Scan(&set.Items).
		Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (s *service) DescribeSecret(ctx context.Context, req *secret.DescribeSecretRequest) (
	*secret.Secret, error) {
	query := s.db.WithContext(ctx).Where("id = ?", req.Id)

	ins := secret.NewDefaultSecret()
	if err := query.First(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DeleteSecret(ctx context.Context, req *secret.DeleteSecretRequest) (
	*secret.Secret, error) {
	ins, err := s.DescribeSecret(ctx, secret.NewDescribeSecretRequest(req.Id))
	if err != nil {
		return nil, err
	}

	err = s.db.WithContext(ctx).Delete(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}
