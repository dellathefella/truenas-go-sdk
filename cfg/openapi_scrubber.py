import yaml
import os
import argparse
import ruamel.yaml
from pathlib import Path
import re


def remove_key_with_value(dct, value):
    new_dct = {}
    for k, v in dct.items():
        if isinstance(v, dict):
            new_v = remove_key_with_value(v, value)
            if new_v:
                new_dct[k] = new_v
        elif v != value:
            new_dct[k] = v
    return new_dct


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('-i','--input_file_path',required=True)
    parser.add_argument('-c','--config_file_path',required=True)      
    parser.add_argument('-o','--output_file_path',required=True)

    args = parser.parse_args()

    input_file = Path(args.input_file_path)
    config_file = Path(args.config_file_path)
    output_file = Path(args.output_file_path)

    api_methods = ["delete","get","post","put"]

    # Open input file to begin processing
    with open(input_file, 'r') as input_file:
        input = yaml.load(input_file, Loader=yaml.SafeLoader)
    input_file.close()

    with open(config_file, 'r') as config_file:
        config = yaml.load(config_file, Loader=yaml.SafeLoader)
    config_file.close()

    if config["regex_remove_expressions"] != None:
        try:
            for component in list(input["components"]["schemas"]):
                for remove_path_regex_expression in config["regex_remove_expressions"]:
                    if re.search(remove_path_regex_expression,component):
                        del input["components"]["schemas"][component]
        except:
            print(Exception)
        try:
            for path in list(input["paths"]):
                for remove_path_regex_expression in config["regex_remove_expressions"]:
                    if re.search(remove_path_regex_expression,path):
                        print(f'{input["paths"][path]}')
                        del input["paths"][path]
        except:
            print(Exception)            
    # for path in list(input["paths"]):
    #     if re.search("^.*{{id}}.*$",path):
    #         input["paths"][path] = input["paths"][f'{path}']
    #         del input["paths"][path]
    #     for api_method in api_methods:
    #         try:
    #             if input["paths"][path][api_method]["responses"]["200"]["$ref"] != None:
    #                 ref = input["paths"][path][api_method]["responses"]["200"]["$ref"]
    #                 del input["paths"][path][api_method]["responses"]["200"]["$ref"]
    #                 # Create new dict structure
    #                 dict_200 = {"description" : f"placeholder for {path}","content" : {"application/json": { "schema" : { "$ref":ref }}}}
    #                 input["paths"][path][api_method]["responses"]["200"].update(dict_200)
    #                 print(f'{input["paths"][path][api_method]["responses"]["200"]} corrected')
    #         except:
    #             print(f'{api_method} not found on {path}')

    # Write yaml file
    with open(output_file, 'w') as output_file:
        yaml.dump(input, output_file)
    output_file.close()

    #print(input)
    # Print the values as a dictionary
