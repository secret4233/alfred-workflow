use serde::Deserialize;
use serde::Serialize;

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Global {
    #[serde(rename = "show_in_menu_bar")]
    pub show_in_menu_bar: Option<bool>,

    #[serde(rename = "show_profile_name_in_menu_bar")]
    pub show_profile_name_in_menu_bar: Option<bool>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Modifiers {
    #[serde(rename = "optional")]
    pub optional: Option<Vec<String>>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct From {
    #[serde(rename = "key_code")]
    pub key_code: Option<String>,

    #[serde(rename = "modifiers")]
    pub modifiers: Option<Modifiers>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct To {
    #[serde(rename = "key_code")]
    pub key_code: Option<String>,

    #[serde(rename = "modifiers", skip_serializing_if = "Option::is_none")]
    pub modifiers: Option<Vec<String>>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Manipulators {
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,

    #[serde(rename = "from")]
    pub from: Option<From>,

    #[serde(rename = "to")]
    pub to: Option<Vec<To>>,

    #[serde(rename = "type")]
    pub manipulators_type: Option<String>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Rules {
    #[serde(rename = "description")]
    pub description: Option<String>,

    #[serde(rename = "manipulators")]
    pub manipulators: Option<Vec<Manipulators>>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ComplexModifications {
    #[serde(rename = "rules")]
    pub rules: Option<Vec<Rules>>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Identifiers {
    #[serde(rename = "is_keyboard")]
    pub is_keyboard: Option<bool>,

    #[serde(rename = "product_id")]
    pub product_id: Option<i32>,

    #[serde(rename = "vendor_id")]
    pub vendor_id: Option<i32>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SimpleModifications {
    #[serde(rename = "from")]
    pub from: Option<From>,

    #[serde(rename = "to")]
    pub to: Option<Vec<To>>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Devices {
    #[serde(rename = "disable_built_in_keyboard_if_exists")]
    pub disable_built_in_keyboard_if_exists: Option<bool>,

    #[serde(rename = "identifiers")]
    pub identifiers: Option<Identifiers>,

    #[serde(rename = "ignore")]
    pub ignore: Option<bool>,

    #[serde(rename = "simple_modifications")]
    pub simple_modifications: Option<Vec<SimpleModifications>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VirtualHidKeyboard {
    #[serde(rename = "keyboard_type_v2")]
    pub keyboard_type_v2: Option<String>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Profiles {
    #[serde(rename = "complex_modifications")]
    pub complex_modifications: Option<ComplexModifications>,

    #[serde(rename = "devices")]
    pub devices: Option<Vec<Devices>>,

    #[serde(rename = "name")]
    pub name: Option<String>,

    #[serde(rename = "selected")]
    pub selected: Option<bool>,

    #[serde(rename = "simple_modifications")]
    pub simple_modifications: Option<Vec<SimpleModifications>>,

    #[serde(rename = "virtual_hid_keyboard")]
    pub virtual_hid_keyboard: Option<VirtualHidKeyboard>,
}
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct KarabinerConfig {
    #[serde(rename = "global")]
    pub global: Option<Global>,

    #[serde(rename = "profiles")]
    pub profiles: Option<Vec<Profiles>>,
}
