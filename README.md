# Terraform Provider: vaultdecode

The [vaultdecode](https://registry.terraform.io/providers/michaelkosir/vaultdecode/latest/docs) provider is used to decode Vault root tokens with.

* [Terraform Registry](https://registry.terraform.io/providers/michaelkosir/vaultdecode/latest/docs)

To use provider functions, declare the provider as a required provider in your Terraform configuration:

```hcl
terraform {
  required_providers {
    vaultdecode = {
      source = "michaelkosir/vaultdecode"
    }
  }
}
```

## Decoding a Vault root token

The `root_token` function decodes a Vault root token using the encoded token, one-time password, and its length.

```hcl
output "decoded_root_token" {
  value = provider::vaultdecode::root_token("PzVHSRMXERYIFCl9IhIMKwUoMz8OJBEgXAIIdQ", "WC4gXgKrEUZIRDxHRNTiiUIf9GO8", 28)
}
```

## License

[Mozilla Public License v2.0](./LICENSE)
