variable "ami_filter_name" {
  default = "chef-highperf-centos7-*"
}

variable "ami_filter_owner" {
  default = "446539779517"
}

variable "ami_filter_virt_type" {
  default = "hvm"
}

variable "automate_ebs_volume_iops" {
  default = 100
}

variable "automate_ebs_volume_size" {
  default = 50
}

variable "automate_ebs_volume_type" {
  default = "gp3"
}

variable "automate_fqdn" {
}

variable "automate_instance_count" {
  default = 1
}

variable "automate_lb_certificate_arn" {
}

variable "automate_server_instance_type" {
  default = "t3a.medium"
}

variable "aws_ami_id" {
  default     = ""
  description = "Setting this value overrides ami search features"
}

variable "aws_cidr_block_addr" {
}

variable "aws_instance_profile_name" {
  default = ""
}

variable "aws_os_snapshot_role_arn" {
  default = ""
}

variable "aws_region" {
  default     = ""
  description = "The name of the selected AWS region / datacenter."
}

variable "aws_ssh_key_file" {
}

variable "aws_ssh_key_pair_name" {
}

variable "aws_ssh_port" {
  default = 22
}

variable "aws_ssh_user" {
  default = "centos"
}

variable "aws_vpc_id" {
}

variable "chef_ebs_volume_iops" {
  default = 100
}

variable "chef_ebs_volume_size" {
  default = 50
}

variable "chef_ebs_volume_type" {
  default = "gp3"
}

variable "chef_server_instance_count" {
  default = 1
}

variable "chef_server_instance_type" {
  default = "t3a.medium"
}

variable "chef_server_lb_certificate_arn" {
}

variable "delete_on_termination" {
  default = true
}

variable "json_data" {
}

variable "kibana_listen_port" {
  default = 5601
}

variable "lb_access_logs" {
  default = false
}

variable "managed_opensearch_domain_name" {
  default = ""
}

variable "managed_opensearch_domain_url" {
  default = ""
}

variable "nfs_mount_path" {
  default = "/mnt/automate_backups"
}

variable "opensearch_ebs_volume_iops" {
  default = 300
}

variable "opensearch_ebs_volume_size" {
  default = 50
}

variable "opensearch_ebs_volume_type" {
  default = "gp3"
}

variable "opensearch_instance_count" {
  default = 3
}

variable "opensearch_listen_port" {
  default = 9200
}

variable "opensearch_server_instance_type" {
  default = "m5a.large"
}

variable "os_snapshot_user_access_key_id" {
  default = ""
}

variable "os_snapshot_user_access_key_secret" {
  default = ""
}

variable "pgleaderchk_listen_port" {
  default = 6432
}

variable "postgresql_ebs_volume_iops" {
  default = 150
}

variable "postgresql_ebs_volume_size" {
  default = 50
}

variable "postgresql_ebs_volume_type" {
  default = "gp3"
}

variable "postgresql_instance_count" {
  default = 3
}

variable "postgresql_listen_port" {
  default = 5432
}

variable "postgresql_server_instance_type" {
  default = "t3a.medium"
}

variable "private_custom_subnets" {
  default = []
  type    = list(string)
}

variable "proxy_listen_port" {
  default = 7432
}

variable "public_custom_subnets" {
  default = []
  type    = list(string)
}

variable "setup_managed_services" {
  default = false
}

variable "ssh_user_sudo_password" {
  default = ""
}

variable "sudo_cmd" {
  default = "sudo"
}

variable "tag_name" {
  default = "A2"
}

variable "tags" {
}

variable "tmp_path" {
  default = "/var/automate-ha"
}
