import * as fs from 'fs';
import {
  Connection,
  Keypair,
  PublicKey,
  sendAndConfirmTransaction,
  Transaction,
  AddressLookupTableProgram,
} from '@solana/web3.js';

async function main() {
  const keypairPath = './ccip-keypair.json';
  if (!fs.existsSync(keypairPath)) {
    throw new Error(`❌ Missing keypair file at ${keypairPath}`);
  }

  const secret = Uint8Array.from(JSON.parse(fs.readFileSync(keypairPath, 'utf-8')));
  const keypair = Keypair.fromSecretKey(secret);

  console.log(`🔑 Using keypair with public key: ${keypair.publicKey.toBase58()}`);

  const args = process.argv.slice(2);
  if (args.length === 0) {
    throw new Error('❌ No addresses provided. Usage: ts-node create-lookup-table.ts <ADDR1> <ADDR2> ...');
  }

  const addresses: PublicKey[] = args.map((addr) => new PublicKey(addr));

  const connection = new Connection('https://api.devnet.solana.com');
  const recentSlot = await connection.getSlot('finalized');

  const [createIx, lookupTableAddress] = AddressLookupTableProgram.createLookupTable({
    authority: keypair.publicKey,
    payer: keypair.publicKey,
    recentSlot,
  });

  const extendIx = AddressLookupTableProgram.extendLookupTable({
    payer: keypair.publicKey,
    authority: keypair.publicKey,
    lookupTable: lookupTableAddress,
    addresses,
  });

  const tx = new Transaction().add(createIx, extendIx);
  const sig = await sendAndConfirmTransaction(connection, tx, [keypair]);

  console.log('✅ Lookup Table created at:', lookupTableAddress.toBase58());
  console.log('📦 Included addresses:', addresses.length);
  console.log('🔗 Tx signature:', sig);
}

main().catch((err) => {
  console.error('❌ Error:', err);
  process.exit(1);
});
