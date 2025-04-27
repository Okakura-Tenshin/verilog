from flask import Flask, request, jsonify
from MyVcdvcd import vcdparse
from vcd_gene import vcd_gene_iverilog

app = Flask(__name__)

@app.route("/run_test", methods=["POST"])
def run_test():
    data = request.get_json()
    print("the data is",data)
    received = data["received"]
    tb = data["testbench"]
    std = data["standard"]

    try:
        vcd_path = vcd_gene_iverilog(received, tb)
        vcd1 = vcdparse.parse_vcd_file(vcd_path)
        vcd2 = vcdparse.parse_vcd_file(std)
        result = vcdparse.generate_result_table(vcd1, vcd2)
        return jsonify({"match": check_all_pairs_equal(result)})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

def check_all_pairs_equal(result_table):
    for entry in result_table:
        if entry[1] != entry[2]:  # 判断第二个和第三个是否完全一样
            return 0
    return 1


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)