
pub fn is_valid_on_ramp(&SourceChainConfig) -> bool {
    let len = self.on_ramp.len();

    len == 0 || len == 64 || len == 128
}

pub fn get_on_ramps(&SourceChainConfig) -> Result<Vec<[u8; 64]>> {
    require!(self.is_valid_on_ramp(), CcipRouterError::InvalidInputs);

    let valid_on_ramps: Vec<[u8; 64]> = match self.on_ramp.len() {
        0 => Vec::new(),
        64 => vec![self.on_ramp[0..64]
            .try_into()
            .map_err(|_| CcipRouterError::InvalidInputs)?],
        128 => vec![
            self.on_ramp[0..64]
                .try_into()
                .map_err(|_| CcipRouterError::InvalidInputs)?,
            self.on_ramp[64..128]
                .try_into()
                .map_err(|_| CcipRouterError::InvalidInputs)?,
        ],
        _ => return Err(CcipRouterError::InvalidInputs.into()),
    };
    Ok(valid_on_ramps)
}

pub fn is_on_ramp_configured(&SourceChainConfig, on_ramp: &[u8]) -> bool {
    let mut on_ramp_bytes = [0u8; 64];
    let len = on_ramp.len().min(64);
    on_ramp_bytes[..len].copy_from_slice(&on_ramp[..len]);

    let valid_on_ramps = self.get_on_ramps();
    if valid_on_ramps.is_err() {
        return false;
    }
    valid_on_ramps.unwrap().contains(&on_ramp_bytes)
}