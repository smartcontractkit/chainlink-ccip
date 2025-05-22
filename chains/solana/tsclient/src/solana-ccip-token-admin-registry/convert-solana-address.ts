import { bs58 } from "@coral-xyz/anchor/dist/cjs/utils/bytes";

const solanaAddress = "2NcHzmcxjAnrq2KWtreLxvoF3vi2ciWabvnXoLQUYGTW";
const bytes = bs58.decode(solanaAddress);
console.log("0x" + Buffer.from(bytes).toString("hex"));