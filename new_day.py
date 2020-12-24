#!/usr/bin/env python3

import argparse
import os.path
import shutil
import sys

TEMPLATE = 'template'

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('day', type=int)
    args = parser.parse_args()

    root_dir = os.path.dirname(os.path.abspath(sys.argv[0]))
    template_dir = os.path.join(root_dir, TEMPLATE)
    day_dir = f"day{args.day:02}"
    filename = f"day{args.day:02}.go"

    dir_fullpath = os.path.join(root_dir, day_dir)

    go_file_fullpath = os.path.join(dir_fullpath, filename)

    if os.path.exists(dir_fullpath):
        sys.exit("already exists!")

    shutil.copytree(template_dir, day_dir)
    shutil.move(os.path.join(day_dir, 'template.go'), go_file_fullpath)



if __name__ == '__main__':
    main()
