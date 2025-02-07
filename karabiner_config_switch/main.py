import hashlib
import time
import sys
import json


CONFIG_PATH = '/Users/bytedance/.config/karabiner/karabiner.json'
wf = {"items" = []}

with open('{}/{}'.format(CONFIG_PATH)) as json_data:
    config = json.load(json_data)
    for profile in config['profiles']:
        wf(
            profile['name'], 'Keyboard Preset Profile',
            arg=profile['name'], valid=True)

args = sys.argv
result = {
    "items": [{
        "title": "result",
        "subtitle": url,
        "arg": url
    }]
}

sys.stdout.write(json.dumps(result))
