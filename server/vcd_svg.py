import json

import wavedrom

from vcd_json import vcd2json
from json_wave import json2wave


def vcd2svg(vcd_file):
    json_data_str = vcd2json(vcd_file)
    json_data = json.loads(json_data_str)
    wave_data = json2wave(json_data)
    wavedrom_json_str = json.dumps(wave_data)
    svg = wavedrom.render(wavedrom_json_str)
    # svg.saveas("output_wave.svg")
    return svg


def main():
    vcd_file = 'test_final.vcd'
    svg = vcd2svg(vcd_file)
    svg.saveas("output_wave.svg")


if __name__ == '__main__':
    main()
