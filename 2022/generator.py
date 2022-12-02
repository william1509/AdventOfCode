import os
import argparse
import shutil

if __name__ == "__main__":
    
    parser = argparse.ArgumentParser(description='Generates a folder following the template')

    parser.add_argument("--day", required=True, type=int, help='Number corresponding to the day')
    args = parser.parse_args()
   
    dir_name = args.day if args.day > 10 else f"0{args.day}"

    dirs = os.listdir()

    if dir_name in dirs:
        y = input(f"This will remove the directory: {dir_name}. Do you want to continue (y/n) ?")
        if (y == "y"): 
            shutil.rmtree(dir_name)

    shutil.copytree("template", dir_name)
    
    
