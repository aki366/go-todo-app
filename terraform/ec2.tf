resource "aws_instance" "example" {
  ami           = "ami-01bef798938b7644d"
  instance_type = "t2.micro"
}
