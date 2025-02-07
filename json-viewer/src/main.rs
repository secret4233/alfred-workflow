use clap::Parser;
use serde_json::{Map, Value};
use std::io::Write;
use tempfile::NamedTempFile;
use std::borrow::Cow;
use std::collections::HashMap;
use serde::Serialize;

#[derive(Parser)]
#[command(name = "json-viewer")]
struct Cli {
    /// JSON 文本
    json_text: String,
}

#[derive(Serialize)]
struct ScriptFilterItem<'a> {
    title: Cow<'a, str>,
    subtitle: Option<Cow<'a, str>>,
    arg: Option<Cow<'a, str>>,
    variables: Option<HashMap<Cow<'a, str>, Cow<'a, str>>>
}

fn parse_json_recursive(value: &Value) -> Value {
    match value {
        Value::String(s) => {
            if let Ok(parsed) = serde_json::from_str(s) {
                parse_json_recursive(&parsed)
            } else {
                Value::String(s.clone())
            }
        }
        Value::Object(map) => {
            let mut new_map = Map::new();
            for (k, v) in map {
                new_map.insert(k.clone(), parse_json_recursive(v));
            }
            Value::Object(new_map)
        }
        Value::Array(arr) => {
            Value::Array(arr.iter().map(parse_json_recursive).collect())
        }
        _ => value.clone(),
    }
}

fn create_temp_html(json: &Value) -> std::io::Result<NamedTempFile> {
    let mut file = NamedTempFile::new()?;
    write!(
        file,
        r#"<!DOCTYPE html>
<html>
<head>
    <title>JSON Viewer</title>
    <link href="https://cdn.jsdelivr.net/npm/jsoneditor@9.10.0/dist/jsoneditor.min.css" rel="stylesheet" type="text/css">
    <script src="https://cdn.jsdelivr.net/npm/jsoneditor@9.10.0/dist/jsoneditor.min.js"></script>
    <style type="text/css">
        #jsoneditor {{ height: 100vh; }}
    </style>
</head>
<body>
    <div id="jsoneditor"></div>
    <script>
        const container = document.getElementById("jsoneditor");
        const options = {{mode: 'view'}};
        const editor = new JSONEditor(container, options);
        const json = {};
        editor.set(json);
    </script>
</body>
</html>"#,
        serde_json::to_string(json)?
    )?;
    Ok(file)
}

fn main() {
    let cli = Cli::parse();
    let mut items = vec![];

    match serde_json::from_str(&cli.json_text) {
        Ok(json) => {
            let parsed = parse_json_recursive(&json);
            
            // 创建两个选项
            let mut vars = HashMap::new();
            vars.insert(Cow::from("json"), Cow::from(serde_json::to_string(&parsed).unwrap()));

            items.push(ScriptFilterItem {
                title: Cow::from("查看解析后的 JSON"),
                subtitle: Some(Cow::from("在浏览器中打开 JSON 查看器")),
                arg: Some(Cow::from("browser")),
                variables: Some(vars)
            });

            items.push(ScriptFilterItem {
                title: Cow::from("复制解析后的 JSON"),
                subtitle: Some(Cow::from("复制格式化后的 JSON 文本")),
                arg: Some(Cow::from(serde_json::to_string_pretty(&parsed).unwrap())),
                variables: None
            });
        }
        Err(e) => {
            items.push(ScriptFilterItem {
                title: Cow::from("JSON 解析错误"),
                subtitle: Some(Cow::from(e.to_string())),
                arg: None,
                variables: None
            });
        }
    }

    if let Some(args) = std::env::args().skip(1).next() {
        if args == "browser" {
            if let Some(json_str) = std::env::var("json").ok() {
                if let Ok(json) = serde_json::from_str(&json_str) {
                    if let Ok(file) = create_temp_html(&json) {
                        let path = file.path().to_str().unwrap();
                        webbrowser::open(&format!("file://{}", path)).ok();
                        std::mem::forget(file); // 保持文件直到浏览器打开
                    }
                }
            }
            return;
        }
    }

    println!("{{\"items\":{}}}", serde_json::to_string(&items).unwrap());
}