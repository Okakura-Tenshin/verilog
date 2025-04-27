from vcdvcd import VCDVCD

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

# Get the VCD file names from user input or command line arguments
'''
vcd_filename1 = input('Enter the first VCD file name (with double quotes): ')
vcd_filename2 = input('Enter the second VCD file name (with double quotes): ')
'''
vcd_filename1 = "./d470.vcd"
vcd_filename2 = "./d471.vcd"

# Parse the first VCD file
result1 = parse_vcd_file(vcd_filename1)

# Parse the second VCD file
result2 = parse_vcd_file(vcd_filename2)

'''
vcd_filename3 = "./timer.vcd"
result3 = parse_vcd_file(vcd_filename3)
print("\nResults for", vcd_filename3)
for signal_name, signal_data in result3.items():
    print(f"Signal: {signal_name}")
    print("Data:", signal_data)
'''

'''
# Print results for the first VCD file
print("\nResults for", vcd_filename1)
for signal_name, signal_data in result1.items():
    print(f"Signal: {signal_name}")
    print("Data:", signal_data)

# Print results for the second VCD file
print("\nResults for", vcd_filename2)
for signal_name, signal_data in result2.items():
    print(f"Signal: {signal_name}")
    print("Data:", signal_data)

print("\n")
'''
###########################################

from termcolor import colored

# 定义两个具有相同键的字典
dict1 = result1
dict2 = result2

def isKeyPrinted(i,j):
    if i==j==0:
        return key
    else:
        return ''

# 遍历字典，按三列输出，并将第二个字典中不一致的值标为红色
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
                print(f"{isKeyPrinted(i,j):<25}\033[92m{val1:<20}\033[91m{val2:<20}\033[0m")
            else:
                print(f"{isKeyPrinted(i,j):<25}{val1:<20}{val2:<20}")
            i=i+1
            j=j+1
        elif value1[i][0]<=value2[j][0]:
            print(f"{isKeyPrinted(i,j):<25}\033[92m{val1:<20}\033[91m{'-':<20}\033[0m")
            i=i+1
        else:
            print(f"{isKeyPrinted(i,j):<25}\033[92m{'-':<20}\033[91m{val2:<20}\033[0m")
            j=j+1

    if i<len(value1):
        print(f"{isKeyPrinted(i,j):<25}\033[92m{val1:<20}\033[91m{'-':<20}\033[0m")
        i=i+1
    if j<len(value2):
        print(f"{isKeyPrinted(i,j):<25}\033[92m{'-':<20}\033[91m{val2:<20}\033[0m")
        j=j+1
