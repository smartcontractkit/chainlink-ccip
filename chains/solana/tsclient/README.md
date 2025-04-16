# 🔁 Solana to EVM CCIP Send Example

This tool demonstrates how to perform a **CCIP send** from **Solana to an EVM chain**, with a step-by-step walkthrough.

---

## ⚙️ Setup NPM

```bash
npm install
```

## ⚙️ Setup Your Solana Environment

### 1. Generate a Solana Keypair (if you don't already have one)

```bash
solana-keygen new --outfile ./ccip-keypair.json
```

### 2. Configure Your Solana CLI

```bash
solana config set --keypair ./ccip-keypair.json
solana config set --url https://api.devnet.solana.com
```

### 3. Verify Your Configuration

```bash
solana config get
```

## ✅ Request Access

Ask the Solana CCIP team to add your public key as an allowed caller to the CCIPRouter program.

## 💰 Get Devnet Funds

```bash
solana airdrop 3
```

## 🔁 Convert SOL to Wrapped SOL (wSOL)

### 1. 1. Create a wSOL Token Account

```bash
spl-token create-account So11111111111111111111111111111111111111112
```

### 2. Transfer SOL into Your wSOL Account

Replace <YOUR_WSOL_ACCOUNT> with the account returned from the previous step.

```bash
solana transfer <YOUR_WSOL_ACCOUNT> 1 --allow-unfunded-recipient
```

Sample output:

```makefile
Signature: "Your signature here"
```

### 3. Check balances

```bach
spl-token accounts --verbose
```

Check that you see the corresponding balance on wSOL

## 🧾 Approve CCIP Router to Spend wSOL

`FEE_BILLING_SIGNER` is the current fee billing signer which is derived from the router program address on devnet

```bash
spl-token approve \
  <YOUR_WSOL_ACCOUNT> \
  0.09 \
  <FEE_BILLING_SIGNER> \
  --url=devnet
```

Sample output:

```yaml
Approve 0.09 tokens
  Account: <YOUR_WSOL_ACCOUNT>
  Delegate:

Signature:
```

## 🚀 Send a CCIP Message

Run your message-sending script:

```bash
npx ts-node src/ccip_send.ts --message='Testing this! --tokens'
```

Example output:

```yaml
🔍 Simulating transaction...
✅ Simulation passed
📤 Transaction sent:
✅ Done
```


### Setup your token, deploy your token pool and configure them through the Token Admin Registry

### 1. Deploy the token pool

First you need to modify the program id. After that, deploy it:

```bash
anchor deploy --program-name burnmint_token_pool
```

TODO: fix this
Upload the IDL:

```bash
anchor idl upgrade <PROGRAM_ID> --filepath target/idl/burnmint_token_pool.json
```


### 2. Create your token


```bash
spl-token create-token
```

```bash
spl-token create-account  <TOKEN_MINT>
```


```bash
spl-token mint <TOKEN_MINT> 100
```


## 3. Configure the token pool

Initialize the token pool

```bash
npx ts-node src/solana-ccip-token-admin-registry/InitializeTokenPool.ts
```

Set Router
```bash
npx ts-node src/solana-ccip-token-admin-registry/SetRouter.ts
```

## 4. Create the lookup table

```bash
npx ts-node src/create-lookup-table.ts <ACCOUNT_0> <ACCOUNT_1> ...
```

## 5. Configure it in the token admin registry

```bash
npx ts-node src/setup-token-admin-registry.ts
```