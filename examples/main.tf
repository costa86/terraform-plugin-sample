terraform {
  required_providers {
    hashicups = {
      version = "0.3"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  # username = "costa"
  # password = "wrong"
}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

# output "coffees" {
#   value = module.psl.all_coffees
# }



# output "my_order" {
#   value = module.psl.order
# }

output "edu_oder" {
  value = module.psl.edu_order
}
