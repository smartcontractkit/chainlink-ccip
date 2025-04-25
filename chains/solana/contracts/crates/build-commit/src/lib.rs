use std::path::Path;
use std::process::Command;

pub fn cargo_instructions(source_file: &str) {
    let program_name = Path::new(source_file)
        .parent()
        .and_then(|p| p.file_name())
        .and_then(|name| name.to_str())
        .unwrap()
        .trim();

    let hash_output = Command::new("git")
        .args(["rev-parse", "HEAD"])
        .output()
        .expect("Failed to execute git");

    let git_hash = String::from_utf8(hash_output.stdout)
        .expect("Invalid UTF-8")
        .trim()
        .to_string();

    let dir_output = Command::new("git")
        .args(["rev-parse", "--git-dir"])
        .output()
        .expect("Failed to execute git");
    let git_dir = String::from_utf8(dir_output.stdout)
        .expect("Invalid UTF-8")
        .trim()
        .to_string();

    let head_output = Command::new("cat")
        .args([format!("{}/HEAD", git_dir.trim())])
        .output()
        .expect("Failed to execute git");
    let head = String::from_utf8(head_output.stdout)
        .expect("Invalid UTF-8")
        .trim()
        .to_string();

    if head.starts_with("ref: ") {
        let reference = head.split_whitespace().nth(1).unwrap().trim().to_string();
        println!("cargo:rerun-if-changed={}/{}", git_dir, reference);
    }

    println!("cargo:rustc-env=CCIP_BUILD_PROGRAM_NAME={}", program_name);
    println!("cargo:rustc-env=CCIP_BUILD_GIT_HASH={}", git_hash);
    println!(
        "cargo:rustc-env=CCIP_BUILD_TYPE_VERSION={} {}",
        program_name, git_hash,
    );
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed={}", source_file);
    println!("cargo:rerun-if-changed={}", file!());
    println!("cargo:rerun-if-changed={}/HEAD", git_dir);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cargo_instructions() {
        // This test is meant to be manually run and inspect the output.
        // Not asserting anything programmatically, as the method prints directly to stdout and the output
        // varies depending on the current git branch/commit
        cargo_instructions("<some_prefix>/chainlink-ccip/chains/solana/contracts/programs/burnmint-token-pool/build.rs");
    }
}
