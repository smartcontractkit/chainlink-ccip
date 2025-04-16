use std::process::Command;

pub fn main() {
    let hash_output = Command::new("git")
        .args(&["rev-parse", "HEAD"])
        .output()
        .expect("Failed to execute git");

    let git_hash = String::from_utf8(hash_output.stdout).expect("Invalid UTF-8");

    let dir_output = Command::new("git")
        .args(&["rev-parse", "--git-dir"])
        .output()
        .expect("Failed to execute git");
    let git_dir = String::from_utf8(dir_output.stdout).expect("Invalid UTF-8");

    println!("cargo:rustc-env=CCIP_BUILD_GIT_HASH={}", git_hash.trim());
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed={}/HEAD", git_dir.trim());
}
