package impl

import (
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
)

func InjectBaseFromSecret(b *resource.Meta, s *secret.Secret) {
	// 补充管理信息
	b.CredentialId = s.Id
	b.Domain = s.Data.Domain
	b.Namespace = s.Data.Namespace
}
