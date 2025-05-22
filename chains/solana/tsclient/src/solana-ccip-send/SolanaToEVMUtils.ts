export class SolanaToEVMUtils {
  /**
   * Converts an EVM address string (0x-prefixed) into a 32-byte left-padded Uint8Array.
   * Required for EVM-to-Solana compatibility in programs expecting 32-byte addresses.
   */
  static evmAddressToSolanaBytes(evmAddress: string): Uint8Array {
    return this.leftPadBytes(this.hexToBytes(evmAddress), 32);
  }

  private static hexToBytes(hex: string): Uint8Array {
    if (hex.startsWith('0x')) hex = hex.slice(2);
    if (hex.length !== 40) throw new Error('Invalid Ethereum address length');
    const bytes = new Uint8Array(20);
    for (let i = 0; i < 40; i += 2) {
      bytes[i / 2] = parseInt(hex.slice(i, i + 2), 16);
    }
    return bytes;
  }

  private static leftPadBytes(data: Uint8Array, length: number): Uint8Array {
    if (data.length > length) throw new Error('Data too long to pad');
    const padded = new Uint8Array(length);
    padded.set(data, length - data.length);
    return padded;
  }
}
  