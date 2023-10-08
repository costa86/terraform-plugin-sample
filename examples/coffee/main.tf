terraform {
  required_providers {
    hashicups = {
      version = "0.3"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

data "hashicups_coffees" "all" {}


data "hashicups_order" "order"{
  id=1
}

resource "hashicups_order" "edu" {
  items {
    coffee{}
    quantity = 2
  }

}

output "edu_order" {
  value = hashicups_order.edu
}


output "order" {
  value = data.hashicups_order.order
}

# Returns all coffees
output "all_coffees" {
  value = data.hashicups_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in data.hashicups_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}
