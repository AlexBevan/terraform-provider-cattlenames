# Terraform Provider for Cattle Names

This Terraform provider generates unique cattle names for non-idempotent resources, such as static services like databases. Below you will find usage instructions, examples, and configuration details.

## Usage Instructions

To use the cattle names provider in your Terraform configuration, you need to declare the provider in your Terraform files. Here is an example of how to do that:

```hcl
terraform {
  required_providers {
    cattlenames = {
      source  = "alexbevan/cattlenames"
      version = "1.0.0"
    }
  }
}

provider "cattlenames" {}

resource "cattlenames" "example" {
  count = 1
}
```

## Resource Configuration

The provider currently supports the following resource:

### `cattlenames`

This resource generates a unique name based on the selected family each time it is applied. The resource does not support idempotency, meaning that applying the same configuration multiple times will result in different names being generated.

#### Arguments

- `family` (Optional): Specifies the family of names to use. Valid options are:
  - `cattle` (default): Generates names like "Bessie" and "MooMoo".
  - `GOT`: Generates names of Game of Thrones characters like "Jon Snow" and "Daenerys".
  - `27Club`: Generates names of 27 Club members like "Kurt Cobain" and "Amy Winehouse".
- `count` (Optional): The number of names to generate. Defaults to 1.

#### Example

```hcl
resource "cattlenames" "example" {
  family = "GOT"
  count  = 1
}
```

## Examples

For a complete example of how to use the cattle names provider, please refer to the `examples/main.tf` file in this repository.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request. We welcome contributions of all kinds!

## License

This project is licensed under the MIT License. See the LICENSE file for more details.