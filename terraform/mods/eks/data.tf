data "aws_eks_addon_version" "eks_addons_recent_version" {
  for_each   = toset(local.eks_default_addons)
  addon_name         = each.key
  most_recent        = true
  kubernetes_version = var.cluster_version
}