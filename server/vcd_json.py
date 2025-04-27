import json

def read_vcd_file(file_path):
    try:
        with open(file_path, 'r') as file:
            return file.read()
    except Exception as e:
        print(f"Error reading VCD file: {e}")
        return None

def parse_vcd(vcd_content):
    signals = {}
    signal_ids = {}
    current_scope = []
    time_unit = "ns"  # Default time unit
    time = 0

    lines = vcd_content.splitlines()
    in_dumpvars = False
    
    for line in lines:
        if '$scope' in line:
            _, scope_type, scope_name, _ = line.split()
            current_scope.append(scope_name)
        elif '$upscope' in line:
            current_scope.pop()
        elif '$timescale' in line:
            parts = line.split()
            if len(parts) > 1:
                time_unit = parts[1]
        elif '$var' in line:
            parts = line.split()
            if len(parts) >= 6:
                _, var_type, var_size, var_id, var_name, _ = parts[:6]
                full_var_name = var_name + ':' + ':'.join(current_scope)
                if var_id not in signal_ids:
                    signal_ids[var_id] = []
                signal_ids[var_id].append(full_var_name)
                if full_var_name not in signals:
                    signals[full_var_name] = []
        elif line.startswith('$dumpvars'):
            in_dumpvars = True
        elif line.startswith('$end'):
            in_dumpvars = False
        elif line.startswith('#'):
            time = int(line[1:])
        else:
            data_line = line.strip()
            if in_dumpvars or (data_line and not data_line.startswith('$')):
                if ' ' in data_line:
                    parts = data_line.split()
                    sig_value = parts[0]
                    sig_id = parts[1]
                else:
                    sig_value = data_line[:-1]
                    sig_id = data_line[-1]
                
                if sig_id in signal_ids:
                    for signal_name in signal_ids[sig_id]:
                        value = sig_value
                        if len(sig_value) > 1 and sig_value.startswith('b'):
                            try:
                                # Convert binary to decimal
                                value = str(int(sig_value[1:], 2))
                            except ValueError:
                                value = "x"
                        if value in {'x', 'z'}:
                            value = "x"
                        signals[signal_name].append({'time': time, 'value': value, 'unit': time_unit})

    return signals

def write_json(data, file_path):
    if not data:
        print("No data to write to JSON.")
    else:
        with open(file_path, 'w') as f:
            json.dump({"signal": data}, f, indent=4)
        print("JSON file successfully generated")

def vcd2json(vcd_file_path):
    vcd_content = read_vcd_file(vcd_file_path)
    if vcd_content:
        parsed_data = parse_vcd(vcd_content)
        return json.dumps({"signal": parsed_data}, indent=4)
    return None

def main():
    vcd_file_path = 'test_final.vcd'
    output_json_path = 'output.json'
    output_json = vcd2json(vcd_file_path)
    write_json(output_json, output_json_path)

if __name__ == '__main__':
    main()
