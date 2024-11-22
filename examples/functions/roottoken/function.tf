terraform {
  required_providers {
    vaultdecode = {
      source = "michaelkosir/vaultdecode"
    }
  }
}

output "decoded_root_token" {
  value = provider::vaultdecode::root_token("PzVHSRMXERYIFCl9IhIMKwUoMz8OJBEgXAIIdQ", "WC4gXgKrEUZIRDxHRNTiiUIf9GO8", 28)
}

