# Chameleon Privacy Module

## Overview

The **Chameleon Privacy Module** is a privacy-focused module that provides stealth address functionality and sanction list management. It allows users to:

- **Generate stealth accounts** using elliptic-curve cryptography (ECDH).
- **Recover stealth private keys** for the recipient.
- **Manage sanctioned addresses** by adding, removing, and checking addresses against a sanction list.

This module leverages elliptic curve Diffie-Hellman (ECDH) for secure, private transactions and uses a simple, in-memory storage for the sanctions list.

### Key Features:
1. **Sanctions Management**:
   - Add an address to the sanctions list.
   - Remove an address from the sanctions list.
   - Check if an address is sanctioned.

2. **Stealth Wallet Operations**:
   - Generate a stealth account using a recipient's public key and an ephemeral public key from the sender.
   - Recover the stealth private key for the recipient using their private key and the sender's ephemeral public key.

---

## Prerequisites

To run the project, make sure you have the following installed:

- Go 1.18+
- Git
- `curl` and `jq` for testing

---

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/PrikshitKumar/chameleon-privacy-module
   cd chameleon-privacy-module
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Unit testing**:

   ```bash
    go test ./...
   ```

---

## Running the Project

1. **Build and start the server**:

   ```bash
   go run main.go
   ```

   The server will start running on `http://localhost:8080`.

2. **Set up environment variables (optional)**:

   If you need to specify a custom port, you can do so by setting the `PORT` environment variable.

   ```bash
   export PORT=8081  # Optional, default is 8080
   ```

---

## API Endpoints

### 1. **Stealth Wallet Endpoints**

#### a. **Generate Account**
Generates a new account.
```bash
curl http://localhost:8080/generate-account | jq
```

#### b. **Generate Stealth Account**
Generates a stealth account by accepting the recipient's public key.
```bash
curl -X POST "http://localhost:8080/generate-stealth" -H "Content-Type: application/json" -d '{"pub_key": "0x04244fc9ce4c29334b372faee3f692e49d6dcc7824b5b54afd6b1233bad5db2d368109325c6b99a95f9e9cf8d5eba0a967c5ebee08381f07b3f31b7e562964ec5d"}' | jq
```

#### c. **Recover Stealth Private Key**
Recovers the stealth private key based on the recipient's private key and the sender's ephemeral public key.
```bash
curl -X POST http://localhost:8080/recover-stealth-priv-key -H "Content-Type: application/json" -d '{ "recipient_privkey": "0xb97f63a1825de57f8551245b6a2de926465700809a3397bf0b4987c3310c3a2f", "ephemeral_pubkey": "0x04a844128b99f7eb00283a8882cd43d8484c61abe4c1d8bf3f3fad2c6992c779ed1cbbd95d48210456dd4fc57c5e8649414f19cff404eb961cbd82a03461d23332"}' | jq
```

### 2. **Sanctions Endpoints**

#### a. **Check if Address is Sanctioned**
Checks if an address is sanctioned.
```bash
curl -X POST http://localhost:8080/sanctions/check -H "Content-Type: application/json" -d '{"address": "0xAbc123"}' | jq
```

#### b. **Add Address to Sanctioned List**
Adds an address to the sanctioned list.
```bash
curl -X POST http://localhost:8080/sanctions/add -H "Content-Type: application/json" -d '{"address": "0xAbc123"}' | jq
```

#### c. **Remove Address from Sanctioned List**
Removes an address from the sanctioned list.
```bash
curl -X POST http://localhost:8080/sanctions/remove -H "Content-Type: application/json" -d '{"address": "0xAbc123"}' | jq
```

---

## Stealth Wallet Explanation (ECDH Algorithm)

### Key Concepts

The **Elliptic-curve Diffie–Hellman (ECDH)** algorithm is used to securely generate shared secrets. It is based on elliptic curve cryptography (ECC) and enables two parties to share a secret key without explicitly transmitting it.

In this context:

- **Recipient's private key (`d_r`)** and **public key (`P_r`)** are used by the recipient.
- **Sender's ephemeral private key (`d_e`)** and **public key (`P_e`)** are used by the sender.
- The shared secret (`s`) is derived from these values.

### Mathematical Explanation

1. **Key Exchange**:
   - The **ephemeral private key** (`d_e`) is combined with the **recipient's public key** (`P_r`), or the **recipient's private key** (`d_r`) is combined with the **sender's ephemeral public key** (`P_e`).
   
   ```plaintext
   s = d_e * P_r = d_r * P_e
   ```

2. **Stealth Address**:
   - The recipient can then recover the **stealth private key** (`d_s`) using their private key (`d_r`) and the sender's ephemeral public key (`P_e`).

   ```plaintext
   d_s = (d_r + H(d_r * P_e)) mod n
   ```

3. **Verification**:
   - To verify correctness, the recovered private key (`d_s`) should match the **stealth public key** generated by the sender.

### Key Insights:

- **d_s** (stealth private key) is **different from** the sender’s ephemeral private key (`d_e`), but both allow for the same **stealth public key** (`P_s`) to be used.
- The **ephemeral private key** (`d_e`) is never intended to be recovered by the recipient.

---

## Testing and Troubleshooting

- Make sure the server is running on the correct port (`8080` by default).
- Use `curl` and `jq` for testing and parsing responses.
- Verify that all endpoints return expected JSON responses.
- Check server logs for detailed error reporting if issues arise.

---

## References Used: 
1. https://eips.ethereum.org/EIPS/eip-5564
2. https://research.csiro.au/blockchainpatterns/general-patterns/stealth-address/
3. https://chainstack.com/stealth-addresses-blockchain-transaction-privacy/