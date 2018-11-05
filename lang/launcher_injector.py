from __future__ import print_function

import argparse
import json

parser = argparse.ArgumentParser(
    fromfile_prefix_chars='@',
    description='')

parser.add_argument(
    '--entrypoint_prefix', action='append', default=[],
    help='Prefix the docker config entrypoint with the specified value(s)')

parser.add_argument(
    '--input', action='store', required=True,
    help='Source docker config file.')

parser.add_argument(
    '--output', action='store', required=True,
    help='Write modified config to this file.')


def main():
    args = parser.parse_args()
    # open input file
    # prefix 'config["Entrypoint"]' with configured values
    # write to output file
    with open(args.input) as f:
        config = json.load(f)
    entrypoint = []
    entrypoint.extend(args.entrypoint_prefix)
    entrypoint.extend(config["config"].get("Entrypoint", []))
    config["config"]["Entrypoint"] = entrypoint

    #print(json.dumps(config, indent=2))
    with open(args.output, "w") as f:
        f.write(json.dumps(config, indent=2))


if __name__ == '__main__':
    main()
