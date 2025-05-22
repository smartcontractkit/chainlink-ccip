import { PublicKey, Connection } from "@solana/web3.js"
import BN from "bn.js" // eslint-disable-line @typescript-eslint/no-unused-vars
import * as borsh from "@coral-xyz/borsh" // eslint-disable-line @typescript-eslint/no-unused-vars
import * as types from "../types" // eslint-disable-line @typescript-eslint/no-unused-vars
import { PROGRAM_ID } from "../programId"

export interface ExternalExecutionConfigFields {}

export interface ExternalExecutionConfigJSON {}

export class ExternalExecutionConfig {
  static readonly discriminator = Buffer.from([
    159, 157, 150, 212, 168, 103, 117, 39,
  ])

  static readonly layout = borsh.struct([])

  constructor(fields: ExternalExecutionConfigFields) {}

  static async fetch(
    c: Connection,
    address: PublicKey,
    programId: PublicKey = PROGRAM_ID
  ): Promise<ExternalExecutionConfig | null> {
    const info = await c.getAccountInfo(address)

    if (info === null) {
      return null
    }
    if (!info.owner.equals(programId)) {
      throw new Error("account doesn't belong to this program")
    }

    return this.decode(info.data)
  }

  static async fetchMultiple(
    c: Connection,
    addresses: PublicKey[],
    programId: PublicKey = PROGRAM_ID
  ): Promise<Array<ExternalExecutionConfig | null>> {
    const infos = await c.getMultipleAccountsInfo(addresses)

    return infos.map((info) => {
      if (info === null) {
        return null
      }
      if (!info.owner.equals(programId)) {
        throw new Error("account doesn't belong to this program")
      }

      return this.decode(info.data)
    })
  }

  static decode(data: Buffer): ExternalExecutionConfig {
    if (!data.slice(0, 8).equals(ExternalExecutionConfig.discriminator)) {
      throw new Error("invalid account discriminator")
    }

    const dec = ExternalExecutionConfig.layout.decode(data.slice(8))

    return new ExternalExecutionConfig({})
  }

  toJSON(): ExternalExecutionConfigJSON {
    return {}
  }

  static fromJSON(obj: ExternalExecutionConfigJSON): ExternalExecutionConfig {
    return new ExternalExecutionConfig({})
  }
}
