variable "instances" {
  default = [
    {
        instance_id = "i-00413914761f0ccd9"
        target_instance_type = "t3.nano"
    },
    {
        instance_id = "i-0a14606d7266d5376"
        target_instance_type = "t3.micro"
    }
  ]
}