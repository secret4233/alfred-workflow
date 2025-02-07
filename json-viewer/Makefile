.PHONY: all clean build package install

BUILD_DIR = target/release
WORKFLOW_DIR = workflow
WORKFLOW_NAME = json_test.alfredworkflow
ALFRED_WORKFLOW_DIR = "/Users/bytedance/Nutstore Files/data/Alfred.alfredpreferences/workflows"

all: build package

build:
	cargo build --release
	cp $(BUILD_DIR)/alfred-json-viewer $(WORKFLOW_DIR)/alfred-json-viewer

package: build
	cd $(WORKFLOW_DIR) && zip -r ../$(WORKFLOW_NAME) .

install: package
	cp $(WORKFLOW_NAME) "$(ALFRED_WORKFLOW_DIR)/"

clean:
	cargo clean
	rm -f $(WORKFLOW_DIR)/alfred-json-viewer
	rm -f $(WORKFLOW_NAME)