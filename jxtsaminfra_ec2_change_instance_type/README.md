// ================================================================
// Build and Installation Instructions
// ================================================================

/*
To build and use this provider:

1. Initialize the Go module:
   go mod init terraform-provider-ec2-instance-type
   go mod tidy

2. Build the provider:
   go build -o terraform-provider-ec2-instance-type

3. Install locally for Terraform to find:
   # Create the plugins directory
   mkdir -p ~/.terraform.d/plugins/local/custom/ec2-instance-type/1.0.0/linux_amd64/
   
   # Copy the binary
   cp terraform-provider-ec2-instance-type ~/.terraform.d/plugins/local/custom/ec2-instance-type/1.0.0/linux_amd64/

4. Create a terraform configuration file (main.tf):
*/

// ================================================================
// Example Terraform Configuration (main.tf)
// ================================================================

/*
terraform {
  required_providers {
    jxtsaminfra = {
      source  = "yourusername/jxtsaminfra"
      version = "~> 1.0"
    }
  }
}

provider "jxtsaminfra" {
  region = "us-east-1"  # Change to your preferred region
  
  # Optional: Specify AWS credentials if not using environment/IAM roles
  # access_key    = "your-access-key"
  # secret_key    = "your-secret-key"
  # session_token = "your-session-token"  # Only needed for temporary credentials
}

resource "jxtsaminfra_ec2_change_instance_type" "example" {
  instance_id          = "i-1234567890abcdef0"  # Replace with your instance ID
  target_instance_type = "t3.medium"            # Replace with desired type
}

output "instance_info" {
  value = {
    instance_id          = jxtsaminfra_ec2_change_instance_type.example.instance_id
    current_type         = jxtsaminfra_ec2_change_instance_type.example.current_instance_type
    target_type          = jxtsaminfra_ec2_change_instance_type.example.target_instance_type
    state               = jxtsaminfra_ec2_change_instance_type.example.instance_state
  }
}
*/

// ================================================================
// Usage Commands
// ================================================================

/*
After creating main.tf:

1. Initialize Terraform:
   terraform init

2. Plan the changes:
   terraform plan

3. Apply the changes:
   terraform apply

4. To destroy (removes from state only, doesn't affect the instance):
   terraform destroy
*/

// ================================================================
// Additional Files Required for Terraform Registry
// ================================================================

// .goreleaser.yml - For automated releases
---
before:
  hooks:
    - go mod tidy
builds:
  - env:
      - CGO_ENABLED=0
    mod_timestamp: '{{ .CommitTimestamp }}'
    flags:
      - -trimpath
    ldflags:
      - '-s -w -X main.version={{.Version}} -X main.commit={{.Commit}}'
    goos:
      - freebsd
      - windows
      - linux
      - darwin
    goarch:
      - amd64
      - '386'
      - arm
      - arm64
    ignore:
      - goos: darwin
        goarch: '386'
    binary: '{{ .ProjectName }}_v{{ .Version }}'
archives:
  - format: zip
    name_template: '{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}'
checksum:
  name_template: '{{ .ProjectName }}_{{ .Version }}_SHA256SUMS'
  algorithm: sha256
signs:
  - artifacts: checksum
    args:
      - "--batch"
      - "--local-user"
      - "{{ .Env.GPG_FINGERPRINT }}"
      - "--output"
      - "${signature}"
      - "--detach-sign"
      - "${artifact}"
release:
  draft: false
changelog:
  skip: true

// ================================================================
// README.md - Provider documentation
// ================================================================

# Terraform Provider JxtSamInfra

A Terraform provider for managing AWS EC2 instance type changes with proper state management.

## Features

- Change EC2 instance types safely
- Automatic instance stop/start handling
- Comprehensive error handling and validation
- Support for multiple AWS authentication methods

## Usage

```hcl
terraform {
  required_providers {
    jxtsaminfra = {
      source  = "yourusername/jxtsaminfra"
      version = "~> 1.0"
    }
  }
}

provider "jxtsaminfra" {
  region = "us-east-1"
}

resource "jxtsaminfra_ec2_change_instance_type" "example" {
  instance_id          = "i-1234567890abcdef0"
  target_instance_type = "t3.medium"
}
```

## Provider Configuration

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|----------|
| region | AWS region | string | us-east-1 | no |
| access_key | AWS access key ID | string | "" | no |
| secret_key | AWS secret access key | string | "" | no |
| session_token | AWS session token | string | "" | no |

## Resource: jxtsaminfra_ec2_change_instance_type

### Schema

| Name | Description | Type | Required |
|------|-------------|------|----------|
| instance_id | EC2 instance ID to modify | string | yes |
| target_instance_type | Target instance type | string | yes |
| current_instance_type | Current instance type (computed) | string | no |
| instance_state | Current instance state (computed) | string | no |

### Example

```hcl
resource "jxtsaminfra_ec2_change_instance_type" "web_server" {
  instance_id          = "i-1234567890abcdef0"
  target_instance_type = "t3.large"
}
```

## License

MIT License

// ================================================================
// .github/workflows/release.yml - GitHub Actions for releases
// ================================================================

name: release
on:
  push:
    tags:
      - 'v*'
permissions:
  contents: write
jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      - name: Unshallow
        run: git fetch --prune --unshallow
      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.21
      - name: Import GPG key
        id: import_gpg
        uses: crazy-max/ghaction-import-gpg@v5
        with:
          gpg_private_key: ${{ secrets.GPG_PRIVATE_KEY }}
          passphrase: ${{ secrets.PASSPHRASE }}
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v4
        with:
          version: latest
          args: release --rm-dist
        env:
          GPG_FINGERPRINT: ${{ steps.import_gpg.outputs.fingerprint }}
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

// ================================================================
// .github/workflows/test.yml - GitHub Actions for testing
// ================================================================

name: test
on:
  pull_request:
  push:
jobs:
  test:
    runs-on: ubuntu-latest
    timeout-minutes: 15
    strategy:
      fail-fast: false
      matrix:
        terraform:
          - '1.0.*'
          - '1.1.*'
          - '1.2.*'
    steps:
      - uses: actions/setup-go@v3
        with:
          go-version: 1.21
      - uses: hashicorp/setup-terraform@v2
        with:
          terraform_version: ${{ matrix.terraform }}
          terraform_wrapper: false
      - uses: actions/checkout@v3
      - run: go mod download
      - run: go build -v .
      - name: Run linters
        uses: golangci/golangci-lint-action@v3
        with:
          version: latest

// ================================================================
// Publishing Steps for Terraform Registry
// ================================================================

/*
Steps to publish to Terraform Registry:

1. **Prepare Repository:**
   - Create GitHub repository: terraform-provider-jxtsaminfra
   - Replace "yourusername" with your actual GitHub username in all files
   - Push all code to the repository

2. **Set up GPG Signing:**
   ```bash
   # Generate GPG key for signing releases
   gpg --batch --full-generate-key <<EOF
   Key-Type: RSA
   Key-Length: 4096
   Subkey-Type: RSA
   Subkey-Length: 4096
   Expire-Date: 0
   Name-Comment: terraform-provider-jxtsaminfra
   Name-Real: Your Name
   Name-Email: your-email@domain.com
   EOF
   
   # Export public key for Terraform Registry
   gpg --armor --export your-email@domain.com > public-key.asc
   
   # Export private key for GitHub Secrets
   gpg --armor --export-secret-keys your-email@domain.com
   ```

3. **Configure GitHub Secrets:**
   Add these secrets to your GitHub repository:
   - GPG_PRIVATE_KEY: Your private GPG key
   - PASSPHRASE: Your GPG key passphrase

4. **Create Initial Release:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

5. **Register on Terraform Registry:**
   - Go to https://registry.terraform.io/
   - Sign in with GitHub
   - Click "Publish Provider"
   - Select your repository: terraform-provider-jxtsaminfra
   - Upload your public GPG key (public-key.asc)
   - Complete the registration

6. **Verify Publication:**
   After successful registration, your provider will be available at:
   https://registry.terraform.io/providers/yourusername/jxtsaminfra

7. **Documentation:**
   The registry will automatically generate docs from your README.md
   and resource/provider schemas.

Note: Replace "yourusername" with your actual GitHub username throughout all files!
*/