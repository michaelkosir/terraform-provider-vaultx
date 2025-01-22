# Terraform Provider: vaultx

The [vaultx](https://registry.terraform.io/providers/michaelkosir/vaultx/latest/docs) provider is used to experiment with new Terraform features before they enter the official HashiCorp supported `vault` provider. This project is not owned, maintained or supported by HashiCorp. Use at your own risk.

* [Terraform Registry](https://registry.terraform.io/providers/michaelkosir/vaultx/latest/docs)

To use provider functions, declare the provider as a required provider in your Terraform configuration:

```hcl
terraform {
  required_providers {
    vaultx = {
      source = "michaelkosir/vaultx"
    }
  }
}
```

## Decoding a Vault root token

The `decode` function decodes a Vault root token using the encoded token, one-time password, and its length.

```hcl
output "token" {
  value = provider::vaultx::decode("PzVHSRMXERYIFCl9IhIMKwUoMz8OJBEgXAIIdQ", "WC4gXgKrEUZIRDxHRNTiiUIf9GO8", 28)
}
```

## License

[Mozilla Public License v2.0](./LICENSE)
