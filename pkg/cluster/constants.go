package cluster

const (
	defaultResource            = "contrail-cluster"
	defaultResourcePath        = "/" + defaultResource
	defaultNodeRes             = "node"
	defaultNodeResPath         = "/" + defaultNodeRes
	defaultConfigNodeRes       = "contrail-config-node"
	defaultConfigNodeResPath   = "/" + defaultConfigNodeRes
	defaultConfigDBNodeRes     = "contrail-config-database-node"
	defaultConfigDBNodeResPath = "/" + defaultConfigDBNodeRes
	defaultControlNodeRes      = "contrail-control-node"
	defaultControlNodeResPath  = "/" + defaultControlNodeRes
	//defaultWebuiNodeRes        = "contrail-webui-node"
	//defaultWebuiNodeResPath       = "/" + defaultWebuiNodeRes
	defaultAnalyticsNodeRes       = "contrail-analytics-node"
	defaultAnalyticsNodeResPath   = "/" + defaultAnalyticsNodeRes
	defaultAnalyticsDBNodeRes     = "contrail-analytics-database-node"
	defaultAnalyticsDBNodeResPath = "/" + defaultAnalyticsDBNodeRes
	defaultVrouterNodeRes         = "contrail-vrouter-node"
	defaultVrouterNodeResPath     = "/" + defaultVrouterNodeRes
	defaultWorkRoot               = "/var/tmp/contrail_cluster"
	defaultTemplateRoot           = "./pkg/cluster/configs"
	defaultInstanceTemplate       = "instances.tmpl"
	defaultInstanceFile           = "instances.yml"
	defaultProvisioner            = "ansible"
	defaultAnsibleRepo            = "contrail-ansible-deployer"
	defaultAnsibleRepoURL         = "https://github.com/Juniper/" + defaultAnsibleRepo + ".git"
	defaultClusterProvPlay        = "playbooks/install_contrail.yml"
	defaultInstanceProvPlay       = "playbooks/provision_instances.yml"
	defaultInstanceConfPlay       = "playbooks/configure_instances.yml"
)
