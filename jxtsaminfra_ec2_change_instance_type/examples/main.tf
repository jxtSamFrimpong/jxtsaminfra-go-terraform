resource "ec2instancetype_instance_type_changer" "main" {
    for_each = local.instances
  instance_id          = each.value.instance_id 
  target_instance_type = each.value.target_instance_type
}