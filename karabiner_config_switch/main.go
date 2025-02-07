package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bytedance/sonic/ast"

	"github.com/bytedance/sonic"
	"github.com/drgrib/alfred"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Global struct {
		ShowInMenuBar            bool `json:"show_in_menu_bar"`
		ShowProfileNameInMenuBar bool `json:"show_profile_name_in_menu_bar"`
	} `json:"global"`
	Profiles []struct {
		ComplexModifications struct {
			Rules []struct {
				Manipulators []struct {
					Description string `json:"description,omitempty"`
					From        struct {
						KeyCode   string `json:"key_code"`
						Modifiers struct {
							Optional  []string `json:"optional"`
							Mandatory []string `json:"mandatory,omitempty"`
						} `json:"modifiers"`
					} `json:"from"`
					To []struct {
						KeyCode   string   `json:"key_code,omitempty"`
						Modifiers []string `json:"modifiers,omitempty"`
						MouseKey  struct {
							VerticalWheel int `json:"vertical_wheel,omitempty"`
							X             int `json:"x,omitempty"`
							Y             int `json:"y,omitempty"`
						} `json:"mouse_key,omitempty"`
						PointingButton string `json:"pointing_button,omitempty"`
					} `json:"to"`
					Type string `json:"type"`
				} `json:"manipulators"`
				Description string `json:"description,omitempty"`
			} `json:"rules"`
		} `json:"complex_modifications"`
		Name               string `json:"name"`
		VirtualHidKeyboard struct {
			KeyboardTypeV2 string `json:"keyboard_type_v2"`
		} `json:"virtual_hid_keyboard"`
		Devices []struct {
			Identifiers struct {
				IsKeyboard bool `json:"is_keyboard"`
				ProductId  int  `json:"product_id"`
				VendorId   int  `json:"vendor_id"`
			} `json:"identifiers"`
			Ignore              bool `json:"ignore,omitempty"`
			SimpleModifications []struct {
				From struct {
					KeyCode string `json:"key_code"`
				} `json:"from"`
				To []struct {
					KeyCode string `json:"key_code"`
				} `json:"to"`
			} `json:"simple_modifications,omitempty"`
		} `json:"devices,omitempty"`
		Selected            bool `json:"selected,omitempty"`
		SimpleModifications []struct {
			From struct {
				KeyCode string `json:"key_code"`
			} `json:"from"`
			To []struct {
				KeyCode string `json:"key_code"`
			} `json:"to"`
		} `json:"simple_modifications,omitempty"`
	} `json:"profiles"`
}

var (
	configFilePath = "/Users/bytedance/.config/karabiner/karabiner.json"
)

func run(args []string) {

	var err error

	if len(args) == 0 {
		return
	}

	app := &cli.App{
		Name:  "karabiner config switch",
		Usage: "A simple command line tool to manage karabiner configuration",
		Commands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "List all items",
				Action: listItems,
			},
			{
				Name:   "switch",
				Usage:  "Switch the selected item by name",
				Action: switchItem,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Name of the item to switch",
						Required: true,
					},
				},
			},
		},
	}

	err = app.Run(args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func listItems(c *cli.Context) error {
	config, err := readConfig(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	for _, item := range config.Profiles {
		selected := ""
		if item.Selected {
			selected = "selected"
		}
		alfred.Add(alfred.Item{
			Title:    item.Name,
			Subtitle: selected,
			Arg:      item.Name,
		})
	}
	alfred.Run()

	return nil
}

func switchItem(c *cli.Context) error {
	name := c.String("name")
	fmt.Println(name)
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	var config ast.Node
	err = sonic.Unmarshal(data, &config)

	found := false
	items, err := config.Get("profiles").ArrayUseNode()
	if err != nil {
		return fmt.Errorf("error parsing items: %w", err)
	}

	for _, item := range items {
		itemName, err := item.Get("name").String()
		if err != nil {
			return fmt.Errorf("error getting item name: %w", err)
		}

		if itemName == name {
			item.SetAny("selected", true)
			found = true
		} else {
			item.Unset("selected")
		}
	}
	fmt.Println("found:", found)
	configStr, _ := config.String()
	fmt.Println(configStr)

	if !found {
		return fmt.Errorf("item not found: %s", name)
	}

	data, err = sonic.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	err = writeConfig(configFilePath, data)
	if err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func readConfig(configFilePath string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	return config, err
}

func writeConfig(configFilePath string, data []byte) error {
	err := ioutil.WriteFile(configFilePath, data, 0644)
	return err
}

func main() {
	args := os.Args
	run(args)
}
