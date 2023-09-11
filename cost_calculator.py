# COST CALCULATOR
# See the instruction file for further instructions

import os

title = r"""* * * * * * * * * * 
* Cost Calculator *
* * * * * * * * * * 
"""
print(title)

# obtain directory
directory = input('\033[1;4m' + "Enter the path of the folder:" + '\033[0m' + " ") 

# obtain all files (including files in nested folders) in directory)
def get_all_files(directory):
    all_files = []
    for root, _, files in os.walk(directory):
        for file in files:
            relative_path = os.path.relpath(os.path.join(root, file), directory)
            all_files.append(relative_path)
    return all_files

files = get_all_files(directory)

# print directory name for confirmation purposes
folder_name = directory.split("/")[-1]
print('\033[1;4m' + "\nThis is the name of the folder you have specified:" + '\033[0m' + " " + folder_name)

# obtain list of file names in the directory
files = [f for f in files if os.path.isfile(os.path.join(directory, f))]

# separate the costs from the file names and return a sum of the costs
def get_sum(file_list, separator="_"):
    float_costs = []
    formatted_correctly = []
    needs_reformatting = []

    # go through each file
    for file_path in file_list:
        # obtain base filename
        filename = os.path.basename(file_path)
        # obtain file's relative path
        relative_path = os.path.relpath(file_path, directory)
        relative_path = relative_path.replace("../", "")
        # split based on separator '_'
        base_filename, _ = os.path.splitext(filename) 
        if separator in base_filename: # if the file name contains separators
            # obtain last chunk 
            last_delimiter = base_filename.split(separator)[-1]
            try: # if last chunk can be converted to float, add to float_costs
                float_costs.append(float(last_delimiter))
                formatted_correctly.append(relative_path)
            except: # if not a float, then the file name is not formatted correctly
                needs_reformatting.append(relative_path)
        else: # if the file name has no separators
            try:
                float_costs.append(float(base_filename))
                formatted_correctly.append(relative_path)
            except: 
                needs_reformatting.append(relative_path)
    
    # print the relative paths of files included in calculation 
    print('\033[1;4m' + "\nThese are the files that are included in the sum:" + '\033[0m')
    for path in formatted_correctly:
        print(f" - {path}")

    # print the improperly formatted files not included in the sum for user reference
    if len(needs_reformatting) > 0: 
        print('\033[1;4m' + "\nThese files are not properly formatted and are not included in the sum:" + '\033[0m')
        print(' - ', end="")
        print(*needs_reformatting, sep="\n - ")
        print('\033[3m' + "\nIf these files are supposed to be included in the sum, reformat the file names and run the program again!" + '\033[0m')

    # obtain the total costs of the properly formatted files
    total = sum(float_costs)
    return total

# print the sum
total_costs = get_sum(files)
print('\033[1;4m' + "\nTOTAL COSTS:" + '\033[0m' + " " + "${:.2f}\n".format(total_costs))

# possible improvements:
#   - a more friendly interface that lets you select a folder may be easier to use for a non-programmer