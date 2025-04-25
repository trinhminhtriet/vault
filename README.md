# Vault 🔐

Vault is a CLI tool for securely decrypting encrypted data using RSA private keys. Built with Go, it provides a simple and efficient way to manage sensitive information.

## Features

- 🛡️ RSA-based decryption
- 📜 PEM private key support
- 🧩 Base64 decoding
- ⚙️ Lightweight and fast CLI tool

## Installation

### Clone the repository

To clone the repository, use the following command:

```sh
git clone git@github.com:trinhminhtriet/vault.git
cd vault
```

### Environment Variables

```sh
export VAULT_PRIVATE_KEY_PATH="id_rsa path"
```

### Add new command

```sh
cobra-cli add timezone
```

### Run command

```sh
vault timezone EST
```

### Clone command

```sh
cobra-cli add mfa
cobra-cli add getPasswd

vault mfa <encrypted data>
vault getPasswd <encrypted data>
```

### Build

```sh
go build -o dist/vault main.go
```

## 🤝 How to contribute

We welcome contributions!

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m "feat: my new feature"`;
- Push to your branch: `git push origin my-feature`.

Once your pull request has been merged, you can delete your branch.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go](https://golang.org/) - The Go programming language.
- [Cobra](https://github.com/spf13/cobra) - A library for creating powerful command-line applications in Go.
