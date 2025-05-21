use std::path::Path;

/// This function is called by Cargo to set up the build environment.
pub fn cargo_instructions(source_file: &str, version: &str) {
    let program_name = Path::new(source_file)
        .parent()
        .and_then(|p| p.file_name())
        .and_then(|name| name.to_str())
        .unwrap()
        .trim();

    // ensure it is re-run when the build script changes
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed={}", source_file);
    println!("cargo:rerun-if-changed={}", file!());

    // set environment variables for the build
    println!("cargo:rustc-env=CCIP_BUILD_PROGRAM_NAME={}", program_name);
    println!(
        "cargo:rustc-env=CCIP_BUILD_TYPE_VERSION={} {}",
        program_name, version,
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cargo_instructions() {
        // This test is meant to be manually run and inspect the output.
        // Not asserting anything programmatically, just checking the output.
        cargo_instructions("<some_prefix>/chainlink-ccip/chains/solana/contracts/programs/burnmint-token-pool/build.rs", env!("CARGO_PKG_VERSION"));
    }
}
