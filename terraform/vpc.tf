resource "aws_eip" "nat_gw_elastic_ip" {
  vpc = true

  tags = {
    Name = "${local.cluster_name}-nat-eip"
  }
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.44.0"

  name = "eks-${terraform.workspace}-vpc"
  cidr = "10.0.0.0/16"
  azs  = data.aws_availability_zones.available_azs.names

  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]

  public_subnets = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  enable_nat_gateway     = true
  single_nat_gateway     = true
  one_nat_gateway_per_az = false
  enable_dns_hostnames   = true
  reuse_nat_ips          = true

  external_nat_ip_ids = [ aws_eip.nat_gw_elastic_ip.id ]

  tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    environment                                   = terraform.workspace
  }

  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/elb"                      = "1"
    environment                                   = terraform.workspace
  }

  private_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/internal-elb"             = "1"
    environment                                   = terraform.workspace
  }
}
