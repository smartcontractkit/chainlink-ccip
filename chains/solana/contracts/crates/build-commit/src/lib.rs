use std::process::Command;

pub fn cargo_instructions() {
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

    let head_output = Command::new("cat")
        .args(&[format!("{}/HEAD", git_dir.trim())])
        .output()
        .expect("Failed to execute git");
    let head = String::from_utf8(head_output.stdout).expect("Invalid UTF-8");

    if head.starts_with("ref: ") {
        let reference = head.split_whitespace().nth(1).unwrap();
        println!(
            "cargo:rerun-if-changed={}/{}",
            git_dir.trim(),
            reference.trim()
        );
    }

    println!("cargo:rustc-env=CCIP_BUILD_GIT_HASH={}", git_hash.trim());
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed={}/HEAD", git_dir.trim());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cargo_instructions() {
        cargo_instructions();
    }
}
