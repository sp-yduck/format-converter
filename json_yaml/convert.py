import os
import json
import yaml


def yaml_to_json(yaml_file:str, json_file:str):
    with open(yaml_file, "r") as f:
        dict = yaml.load(f, Loader=yaml.CLoader)
    with open(json_file, "w") as f:
        json_string = json.dumps(dict, indent=2)
        f.write(json_string)
    return


def json_to_yaml(json_file:str, yaml_file:str):
    with open(json_file, "r") as f:
        dict = json.load(f)
    with open(yaml_file, "w") as f:
        yaml_string = yaml.dump(dict)
        f.write(yaml_string)
    return


if __name__ == "__main__":
    yaml_file = "alerts.yml"
    json_file = "alerts.json"

    output_dir = "outputs"
    os.makedirs(output_dir, exist_ok=True)

    out_yaml = os.path.join(output_dir, "alerts.yml")
    out_json = os.path.join(output_dir, "alerts.json")
    
    yaml_to_json(yaml_file=yaml_file, json_file=json_file)
    json_to_yaml(json_file=json_file, yaml_file=out_yaml)
    yaml_to_json(yaml_file=out_yaml, json_file=out_json)