variable "env" {
  type = string
}

variable "region" {
  type = string
}

variable "vpc_name" {
  type = string
}

variable "cluster_api_access_cidr_blocks" {
  type = list(string)
}

variable "cluster_name" {
  type = string
}

variable "cluster_version" {
  type = string
}

variable "cluster_enabled_log_types" {
  type = list(string)
}

variable "subnet_types" {
  description = "Types of subnets for lookup to use in cluster"
  type        = list(string)
}

variable "eks_managed_node_groups" {
  type = any
}

variable "eks_managed_node_group_defaults" {
  type = any
}
