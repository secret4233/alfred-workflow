// use alfred;
// use std::env;
use alfred;
use serde_json;
use std::ffi::OsStr;
use std::fs;
use std::io;
use std::io::Error;
use std::path::PathBuf;
use structopt::StructOpt;

mod config_struct;

use config_struct::KarabinerConfig; // 引入 Person 结构体

fn read_json_file(file_path: &OsStr) -> Result<KarabinerConfig, Error> {
    // 读取文件内容
    let data = fs::read_to_string(file_path)?;

    // 将 JSON 字符串解析为 `Person` 结构体
    let config: KarabinerConfig = serde_json::from_str(&data)?;

    Ok(config)
}

fn write_json_file(file_path: &OsStr, config: &KarabinerConfig) -> Result<(), Error> {
    fs::write(file_path, serde_json::to_string_pretty(&config)?)
}
// try to use library https://suibianxiedianer.github.io/rust-cli-book-zh_CN/tutorial/cli-args_zh.html
#[derive(Debug, Clone, StructOpt)]
#[structopt(name = "example", about = "An example of StructOpt usage.")]
struct IoArg {
    #[structopt(short = "l")]
    list: bool,

    #[structopt(short = "u")]
    config_use: String,

    #[structopt(short = "p", parse(from_os_str))]
    config_path: PathBuf,
}

fn main() {
    let opt = IoArg::from_args();
    // let file_path = "~/.config/karabiner/karabiner.json";
    let file_path = opt.config_path.as_os_str();
    let mut result: KarabinerConfig;

    // 读取并解析 JSON 文件
    match read_json_file(file_path) {
        Ok(config) => result = config,
        Err(e) => panic!("Error reading JSON file: {}", e),
    }
    println!("{}", serde_json::to_string_pretty(&result).unwrap());

    if opt.list {
        workflow_output(
            result
                .profiles
                .unwrap()
                .iter()
                .map(|item| item.name.clone())
                .flatten()
                .collect(),
            true,
        );
    } else if opt.config_use.len() > 0 {
        if let Some(profiles) = result.profiles.as_mut() {
            profiles.iter_mut().for_each(|item| {
                item.selected = match &item.name {
                    Some(name) if name == &opt.config_use => Some(true), // 如果 `name` 匹配 `opt.config_use`，则设为 `Some(true)`
                    Some(_) => None, // 如果 `name` 不匹配，则设为 `Some(false)`
                    None => None,    // 如果 `name` 是 `None`，则保留 `item.selected` 为 `None`
                };
            });
            // 将修改后的 result 写回文件
            write_json_file(file_path, &result).unwrap();
        } else {
            panic!("No profiles found in result.");
        }
    }
}

fn workflow_output(list: Vec<String>, json: bool) {
    let items: Vec<alfred::Item> = list
        .into_iter()
        .map(|item| {
            alfred::ItemBuilder::new(item.clone())
                .arg(item.clone())
                .subtitle(item.clone())
                .into_item()
        })
        .collect();
    if json {
        // alfred::json::Builder::with_items(&items)
        //     .write(io::stdout())
        //     .expect("Couldn't write items to Alfred");
        // alfred::json::write_items(io::stdout(), &items).expect("Couldn't write items to Alfred");
    } else {
        alfred::xml::write_items(io::stdout(), &items).expect("Couldn't write items to Alfred");
    }
}
