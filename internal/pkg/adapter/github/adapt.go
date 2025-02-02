package github

import (
	"github.com/aquasecurity/defsec/provider/github"
	"github.com/aquasecurity/tfsec/internal/pkg/adapter/github/repositories"
	"github.com/aquasecurity/tfsec/internal/pkg/adapter/github/secrets"
	"github.com/aquasecurity/tfsec/internal/pkg/block"
)

func Adapt(modules block.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
	}
}
