
pub const CHAIN_CONFIG_SEED: &[u8] = b"remote_chain_config";
pub const CCIP_SENDER: &[u8] = b"ccip_sender";

pub const CCIP_SEND_DISCRIMINATOR: [u8; 8] = [108, 216, 134, 191, 249, 234, 33, 84]; // ccip_send
pub const CCIP_GET_FEE_DISCRIMINATOR: [u8; 8] = [115, 195, 235, 161, 25, 219, 60, 29]; // get_fee