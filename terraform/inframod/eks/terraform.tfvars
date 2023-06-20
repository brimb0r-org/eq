env                            = "dev"
region                         = "us-west-2"
vpc_name                       = "us-west-2.dev.int-cloud.altium.com"
cluster_api_access_cidr_blocks = [
  "10.0.0.0/8",
]
cluster_name              = "dev-us-west-2-regional"
cluster_version           = "1.24"
cluster_enabled_log_types = ["api", "audit", "authenticator", "controllerManager", "scheduler"]
subnet_types              = ["private", "private_unrestricted"]
eks_managed_node_groups   = {
  dev-eks-app = {
    ami_type              = "AL2_x86_64"
    ami_release_version   = "1.24.9-20230203"
    capacity_type         = "ON_DEMAND"
    instance_types        = ["t3a.large"]
    min_size              = 7
    desired_size          = 7
    max_size              = 9
    create_security_group = false
    subnet_type           = "private"
    labels                = {
      role = "app"
    }
  }

  dev-eks-unlimited = {
    ami_type              = "AL2_x86_64"
    ami_release_version   = "1.24.9-20230203"
    capacity_type         = "ON_DEMAND"
    instance_types        = ["t3a.medium"]
    min_size              = 2
    desired_size          = 2
    max_size              = 4
    create_security_group = false
    subnet_type           = "private_unrestricted"
    labels                = {
      role = "unrestricted-app"
    }
  }

  dev-eks-sys = {
    ami_type              = "AL2_x86_64"
    ami_release_version   = "1.24.9-20230203"
    capacity_type         = "ON_DEMAND"
    instance_types        = ["t3a.medium"]
    min_size              = 2
    desired_size          = 2
    max_size              = 2
    create_security_group = false
    subnet_type           = "private"
    labels                = {
      role = "system"
    }
  }
}

eks_managed_node_group_defaults = {
  disk_size                    = 60
  iam_role_additional_policies = ["arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"]
  pre_bootstrap_user_data      = <<-EOT
      yum install -y amazon-ssm-agent
      systemctl enable amazon-ssm-agent
      systemctl start amazon-ssm-agent
      EOT
}
