import json
import re

def parse_input_json(file_path):
    """解析输入 JSON 文件"""
    with open(file_path, 'r') as file:
        data = json.load(file)
    return data

def find_time_unit(signals):
    """根据所有信号确定最小时间单位"""
    min_time_diff = float('inf')
    for values in signals.values():
        if len(values) > 1:
            for i in range(1, len(values)):
                time_diff = values[i]["time"] - values[i-1]["time"]
                if time_diff > 0:
                    min_time_diff = min(min_time_diff, time_diff)
    return min_time_diff if min_time_diff != float('inf') else 10  # 默认时间单位

def process_signal_values(values, time_unit, last_time, flag):
    """处理信号值，生成波形和数据列表"""
    wave = ""
    data_list = []

    previous_value = "x"
    previous_time = 0

    if values:
        previous_value = values[0]["value"]
        previous_time = values[0]["time"]
        if flag and previous_value != "x":
            wave = "="
            data_list.append(previous_value)
        else:
            wave = str(previous_value)

    for entry in values[1:]:
        time = entry["time"]
        value = entry["value"]

        # 计算时间差
        time_diff = (time - previous_time) // time_unit

        # 填充时间间隙
        if time_diff > 1:
            wave += "." * (time_diff - 1)

        # 更新波形
        if value == "x":
            wave += "x"
        else:
            if value == previous_value:
                wave += "."
            else:
                if flag:
                    wave += "="
                    data_list.append(value)
                else:
                    wave += str(value)

        previous_value = value
        previous_time = time

    # 延续最后一个值到时钟的最后
    end_time_diff = (last_time - previous_time) // time_unit
    if end_time_diff > 1:
        wave += "." * (end_time_diff - 1)

    return wave, data_list

def generate_wavedrom_json(data):
    """生成 Wavedrom 格式的 JSON 数据"""
    wavedrom_data = {"signal": []}

    time_unit = find_time_unit(data["signal"])

    # 找到最大的时间戳
    last_time = 0
    for values in data["signal"].values():
        if values:
            last_time = max(last_time, values[-1]["time"])

    signal_length = (last_time // time_unit) + 1

    for key, values in data["signal"].items():
        signal_name = key.split(":")[0]
        flag = any(int(entry["value"]) > 1 for entry in values if entry["value"] != "x")

        wave, data_list = process_signal_values(values, time_unit, last_time, flag)

        # 调整信号长度与最大时间戳一致
        if len(wave) < signal_length:
            wave += "." * (signal_length - len(wave))

        if flag:
            wavedrom_data["signal"].append({"name": signal_name, "wave": wave, "data": data_list})
        else:
            wavedrom_data["signal"].append({"name": signal_name, "wave": wave})

    return wavedrom_data

def save_output_json(data, file_path):
    """保存输出 JSON 文件"""
    with open(file_path, 'w') as file:
        json.dump(data, file, indent=4)

def json2wave(json_data):
    wavedrom_data = generate_wavedrom_json(json_data)
    return wavedrom_data


if __name__ == "__main__":
    input_file = 'output.json'
    output_file = 'wavedrom_output.json'

    input_data = parse_input_json(input_file)
    wavedrom_data = json2wave(input_data)
    save_output_json(wavedrom_data, output_file)
