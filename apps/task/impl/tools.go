package impl

import (
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
)

func InjectBaseFromSecret(b *resource.Meta, s *secret.Secret) {
	// 补充管理信息
	b.CredentialId = s.Meta.Id
	b.Domain = s.Spec.Domain
	b.Namespace = s.Spec.Namespace
}
