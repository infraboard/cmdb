package impl

import (
	"github.com/infraboard/cmdb/apps/credential"
	"github.com/infraboard/cmdb/apps/resource"
)

func InjectBaseFromSecret(b *resource.Meta, s *credential.Secret) {
	// 补充管理信息
	b.CredentialId = s.Id
	b.Domain = s.Data.Domain
	b.Namespace = s.Data.Namespace
}
