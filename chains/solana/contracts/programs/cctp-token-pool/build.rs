fn main() {
    build_commit::cargo_instructions(file!(), env!("CARGO_PKG_VERSION"));
}
