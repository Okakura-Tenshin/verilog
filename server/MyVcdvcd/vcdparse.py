import sys
sys.path.append('./MyVcdvcd/vcdvcd')
from vcdvcd import VCDVCD
from termcolor import colored

def parse_vcd_file(vcd_file):
    # Remove quotes from the file name
    vcd_file = vcd_file.strip('"\'')

    # Do the parsing.
    vcd = VCDVCD(vcd_file)

    # List all human readable signal names.
    signal_names = list(vcd.references_to_ids.keys())

    # Create a dictionary to store signal data for all signals
    signal_data = {}

    # Loop through each signal
    for signal_name in signal_names:
        # Get a signal by human readable name.
        signal = vcd[signal_name]

        # tv is a list of Time/Value delta pairs for this signal.
        tv = signal.tv

        # Store signal data in the dictionary
        signal_data[signal_name] = tv

    # Return a dictionary with the signal data
    return signal_data

def isKeyPrinted(i,j,key):
    if i==j==0:
        return key
    else:
        return ''
    
def compare_vcd_files(dict1,dict2):
    result=""
    for key in dict1:
        value1 = dict1[key]
        value2 = dict2[key]
        i=int(0)
        j=int(0)

        while i<len(value1) and j<len(value2):
            val1=str(value1[i])
            val2=str(value2[j])
            if value1[i][0]==value2[j][0]:
                if value1[i] != value2[j]:
                    result += f"{isKeyPrinted(i,j,key):<25}\033[92m{val1:<20}\033[91m{val2:<20}\033[0m\n"
                else:
                    result += f"{isKeyPrinted(i,j,key):<25}{val1:<20}{val2:<20}\n"
                i=i+1
                j=j+1
            elif value1[i][0]<=value2[j][0]:
                result += f"{isKeyPrinted(i,j,key):<25}\033[92m{val1:<20}\033[91m{'-':<20}\033[0m\n"
                i=i+1
            else:
                result += f"{isKeyPrinted(i,j,key):<25}\033[92m{'-':<20}\033[91m{val2:<20}\033[0m\n"
                j=j+1

        if i<len(value1):
            result += f"{isKeyPrinted(i,j,key):<25}\033[92m{val1:<20}\033[91m{'-':<20}\033[0m\n"
            i=i+1
        if j<len(value2):
            result += f"{isKeyPrinted(i,j,key):<25}\033[92m{'-':<20}\033[91m{val2:<20}\033[0m\n"
            j=j+1

    return result

def generate_result_table(dict1, dict2):
    result_table = []
    for key in dict1:
        value1 = dict1[key]
        value2 = dict2[key]
        i, j = 0, 0

        while i < len(value1) and j < len(value2):
            if value1[i][0] == value2[j][0]:
                result_table.append((isKeyPrinted(i, j, key), value1[i], value2[j]))
                i += 1
                j += 1
            elif value1[i][0] < value2[j][0]:
                result_table.append((isKeyPrinted(i, j, key), value1[i], ("-", "-")))
                i += 1
            else:
                result_table.append((isKeyPrinted(i, j, key), ("-", "-"), value2[j]))
                j += 1

        while i < len(value1):
            result_table.append((isKeyPrinted(i, j, key), value1[i], ("-", "-")))
            i += 1

        while j < len(value2):
            result_table.append((isKeyPrinted(i, j, key), ("-", "-"), value2[j]))
            j += 1

    return result_table