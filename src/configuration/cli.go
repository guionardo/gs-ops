package configuration

type CLIConfiguration struct {
	Name          string `yaml:"name"`
	Description   string `yaml:"description"`
	RepositoryURL string `yaml:"repository_url"`
	DeployBranch  string `yaml:"deploy_branch"`
	DeployTrigger string `yaml:"deploy_trigger"`
}
