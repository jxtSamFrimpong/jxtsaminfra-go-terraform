locals {
  instances = {
    for k,v in var.instances : k => {
      instance_id          = v.instance_id
      target_instance_type = v.target_instance_type
    }
  }

  outputs = {
    for k,v in ec2instancetype_instance_type_changer.main:
        k => {
        instance_id          = v.instance_id
        current_type         = v.current_instance_type
        target_type          = v.target_instance_type
        state               = v.instance_state
    }
  }
}