terraform {
  required_providers {
    vaultx = {
      source = "michaelkosir/vaultx"
    }
  }
}

output "token" {
  value = provider::vaultx::decode("PzVHSRMXERYIFCl9IhIMKwUoMz8OJBEgXAIIdQ", "WC4gXgKrEUZIRDxHRNTiiUIf9GO8", 28)
}
