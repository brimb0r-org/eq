output "cluster_endpoint" {
  value = module.eks.cluster_endpoint
}

output "cluster_id" {
  value = module.eks.cluster_id
}

output "cluser_nodes_release_versions" {
  value = { 
   for node_group in  module.eks.eks_managed_node_groups : node_group.node_group_id => node_group.node_group_node_release_version
  }
}