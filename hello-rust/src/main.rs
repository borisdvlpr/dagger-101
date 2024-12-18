use std::{fs, io::Write, path::Path};

fn main() {
    // create the output directory if it doesn't exist
    let dir = "output";
    if !Path::new(dir).exists() {
        fs::create_dir_all(dir).expect("Failed to create directory");
    }

    // create or open the file and write to it
    let file_path = format!("{}/hello-rust.txt", dir);
    let mut file = fs::File::create(file_path).expect("Failed to create file");
    file.write_all(b"Hello from Rust in Dagger!\n").expect("Failed to write to file");
}
