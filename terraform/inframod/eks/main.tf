locals {
  cluster_security_group_additional_rules = {
    ingress_cluster_api = {
      description = "Allow cluster API access from network"
      type        = "ingress"
      protocol    = "TCP"
      from_port   = 443
      to_port     = 443
      cidr_blocks = var.cluster_api_access_cidr_blocks
    }
  }

  node_security_group_additional_rules = {
    egress_all = {
      description = "Egress all"
      type        = "egress"
      protocol    = "-1"
      from_port   = 0
      to_port     = 0
      cidr_blocks = ["0.0.0.0/0"]
    }

    ingress_self_all = {
      description = "Node to node all ports/protocols"
      protocol    = "-1"
      from_port   = 0
      to_port     = 0
      type        = "ingress"
      self        = true
    }

    ingress_healthcheck = {
      description = "Allow access inside vpc for healthcheck"
      type        = "ingress"
      protocol    = "TCP"
      from_port   = 8080
      to_port     = 32354
      cidr_blocks = [data.aws_vpc.this.cidr_block]
    }

    ingress_9443 = {
      description = "Allow access from control plane to webhook port of AWS load balancer controller"
      type        = "ingress"
      protocol    = "TCP"
      from_port   = 9443
      to_port     = 9443

      source_cluster_security_group = true
    }

    ingress_metrics_server = {
      description = "Cluster API to metrics server"
      type        = "ingress"
      protocol    = "TCP"
      from_port   = 4443
      to_port     = 4443

      source_cluster_security_group = true
    }
  }

  eks_managed_node_groups = {
    for i in keys(var.eks_managed_node_groups) : i =>
    merge(var.eks_managed_node_groups[i], {
      subnet_ids = data.aws_subnets.eks_subnets[var.eks_managed_node_groups[i]["subnet_type"]].ids
    })
  }
}

module "eks" {
  source = "../../../../../modules/eks"

  cluster_name                    = var.cluster_name
  cluster_version                 = var.cluster_version
  cluster_endpoint_private_access = true
  cluster_endpoint_public_access  = false
  cluster_enabled_log_types       = var.cluster_enabled_log_types
  cluster_ip_family               = "ipv4"

  enable_irsa = true
  vpc_id      = data.aws_vpc.this.id
  subnet_ids  = data.aws_subnets.eks_subnets["private"].ids

  node_security_group_additional_rules    = local.node_security_group_additional_rules
  cluster_security_group_additional_rules = local.cluster_security_group_additional_rules

  eks_managed_node_group_defaults = var.eks_managed_node_group_defaults
  eks_managed_node_groups         = local.eks_managed_node_groups

  cluster_addons = {
    coredns    = {}
    kube-proxy = {}
    vpc-cni    = {}
  }

  tags = {
    region    = var.region
    env       = var.env
    service   = "eks"
    dr        = "no"
    data_type = "confidential"
    managed   = "terraform"
  }
}
