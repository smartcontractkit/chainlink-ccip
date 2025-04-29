use std::env::var;
use std::path::Path;
use std::process::Command;

const GIT_HASH_ENV_VAR: &str = "CCIP_BUILD_GIT_HASH";
const PATH_TO_FILE: &str = "../../"; // so that this is in chainlink-ccip/chains/solana/contracts/

/// This function is called by Cargo to set up the build environment.
/// It reads the git hash from an environment variable (useful in CI), a custom file (useful in verifiable builds), or from the git repository (useful in local builds), in that order of preference.
pub fn cargo_instructions(source_file: &str) {
    let program_name = get_program_name(source_file);
    let custom_file_path = Path::new(PATH_TO_FILE).join(GIT_HASH_ENV_VAR);

    let git_hash = if let Ok(git_hash) = var(GIT_HASH_ENV_VAR) {
        // Reading it from the environment variable
        git_hash
    } else if custom_file_path.exists() {
        // Reading it from the custom file
        std::fs::read_to_string(custom_file_path.clone())
            .expect("Failed to read git hash from file")
            .trim()
            .to_string()
    } else {
        // Reading it from the git repository using the git binary that must be in the PATH
        let (git_hash, git_dir, git_ref) = git_strategy();

        // ensure it is re-run if git hash changes, when using the git strategy
        println!("cargo:rerun-if-changed={}/HEAD", git_dir); // for when it's read from git
        if let Some(reference) = git_ref {
            println!("cargo:rerun-if-changed={}/{}", git_dir, reference);
        }

        git_hash
    };

    assert!(
        git_hash.len() == 40,
        "CCIP_BUILD_GIT_HASH (env-var or file) must have a 40 character commit hash, or there must be a git repository in the current directory"
    );

    // ensure it is re-run if git hash changes
    println!("cargo:rerun-if-env-changed={}", GIT_HASH_ENV_VAR); // for when it's read from an env var
    println!(
        "cargo:rerun-if-changed={}", // for when it's read from a file
        Path::new(PATH_TO_FILE)
            .canonicalize() // canonicalize path without filename to ensure it exists first
            .unwrap()
            .join(GIT_HASH_ENV_VAR) // only once it's canonicalized to the dir, we can append the filename
            .display()
    );

    // ensure it is re-run when the build script changes
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed={}", source_file);
    println!("cargo:rerun-if-changed={}", file!());

    // set environment variables for the build
    println!("cargo:rustc-env=CCIP_BUILD_PROGRAM_NAME={}", program_name);
    println!("cargo:rustc-env=CCIP_BUILD_GIT_HASH={}", git_hash);
    println!(
        "cargo:rustc-env=CCIP_BUILD_TYPE_VERSION={} {}",
        program_name, git_hash,
    );
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
        std::env::set_var(GIT_HASH_ENV_VAR, "env4567890123456789012345678901234567890");

        // This test is meant to be manually run and inspect the output.
        // Not asserting anything programmatically, as the method prints directly to stdout and the output
        // varies depending on the current git branch/commit
        cargo_instructions("<some_prefix>/chainlink-ccip/chains/solana/contracts/programs/burnmint-token-pool/build.rs");
    }

    #[test]
    fn test_cargo_instructions_with_file() {
        // write to a file
        let git_hash = "file567890123456789012345678901234567890";
        let file_path = Path::new(PATH_TO_FILE).join(GIT_HASH_ENV_VAR);

        // move existing file to a temporary location
        let use_backup = file_path.exists();
        let temp_file_path = file_path.with_extension("bak");
        if use_backup {
            std::fs::rename(file_path.clone(), temp_file_path.clone())
                .expect("Unable to rename file");
        }

        std::fs::write(file_path.clone(), git_hash).expect("Unable to write file");

        // This test is meant to be manually run and inspect the output.
        // Not asserting anything programmatically, as the method prints directly to stdout and the output
        // varies depending on the current git branch/commit
        cargo_instructions("<some_prefix>/chainlink-ccip/chains/solana/contracts/programs/burnmint-token-pool/build.rs");

        std::fs::remove_file(file_path.clone()).expect("Unable to delete file");

        if use_backup {
            std::fs::rename(temp_file_path, file_path.clone()).expect("Unable to rename file");
        }
    }
}
