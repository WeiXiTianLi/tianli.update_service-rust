use serde::{Deserialize, Serialize};
use serde_json::{to_string, Result};
use std::{fs::File, io::Write, path::Path};

#[derive(Serialize, Deserialize)]
struct Config {
    host: String,
    port: u32,
    database: String,
}

fn check_config_file_is_exist(path: &str) -> bool {
    let path = Path::new(path);
    return path.exists();
}

fn create_default_config_file(path: &str) {
    let path = Path::new(path);
    let display = path.display();

    let mut file = match File::create(&path) {
        Err(why) => panic!("couldn't create {}:{}", display, why),
        Ok(file) => file,
    };

    let config = Config {
        host: "localhost".to_string(),
        port: 3306,
        database: "test".to_string(),
    };

    let json = serde_json::to_string(&config).unwrap();
    file.write_all(json.as_bytes()).unwrap();
}

fn check_or_create_config_file(path: &str) {
    if check_config_file_is_exist(path) {
        println!("Config file is exist");
    } else {
        println!("Config file is not exist");
        create_default_config_file(path);
    }
}

fn main() {
    let config_file_path = "config.json";
    check_or_create_config_file(config_file_path);
}
