use std::env::var;
use std::path::Path;
use std::process::Command;

const GIT_HASH_ENV_VAR: &str = "CCIP_BUILD_GIT_HASH";

pub fn cargo_instructions(source_file: &str) {
    let program_name = get_program_name(source_file);

    let git_hash = if let Ok(git_hash) = var(GIT_HASH_ENV_VAR) {
        git_hash
    } else {
        let (git_hash, git_dir, git_ref) = git_strategy();
        println!("cargo:rerun-if-changed={}/HEAD", git_dir);
        if let Some(reference) = git_ref {
            println!("cargo:rerun-if-changed={}/{}", git_dir, reference);
        }

        git_hash
    };

    assert!(
        git_hash.len() == 40,
        "GIT_HASH_ENV_VAR must be a 40 character hash, or there must be a git repository in the current directory"
    );

    println!("cargo:rerun-if-env-changed={}", GIT_HASH_ENV_VAR);
    println!("cargo:rustc-env=CCIP_BUILD_PROGRAM_NAME={}", program_name);
    println!("cargo:rustc-env=CCIP_BUILD_GIT_HASH={}", git_hash);
    println!(
        "cargo:rustc-env=CCIP_BUILD_TYPE_VERSION={} {}",
        program_name, git_hash,
    );
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed={}", source_file);
    println!("cargo:rerun-if-changed={}", file!());
}

fn get_program_name(source_file: &str) -> &str {
    Path::new(source_file)
        .parent()
        .and_then(|p| p.file_name())
        .and_then(|name| name.to_str())
        .unwrap()
        .trim()
}

fn git_strategy() -> (String, String, Option<String>) {
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

    let reference = if head.starts_with("ref: ") {
        Some(head.split_whitespace().nth(1).unwrap().trim().to_string())
    } else {
        None
    };

    (git_hash, git_dir, reference)
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

    #[test]
    fn test_cargo_instructions_with_env_var() {
        std::env::set_var(GIT_HASH_ENV_VAR, "1234567890123456789012345678901234567890");
        // This test is meant to be manually run and inspect the output.
        // Not asserting anything programmatically, as the method prints directly to stdout and the output
        // varies depending on the current git branch/commit
        cargo_instructions("<some_prefix>/chainlink-ccip/chains/solana/contracts/programs/burnmint-token-pool/build.rs");
    }
}
