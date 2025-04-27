import subprocess
import os

def vcd_gene_iverilog(verilog_file_path, testbench_file_path):

    # testbench_file_name = verilog_file_name.join("_tb.v")
    # testbench_file_path = os.path.join(os.getcwd(), testbench_file_name)

    # 生成iverilog编译命令
    iverilog_cmd = f"iverilog -o wave {testbench_file_path} {verilog_file_path}"
    subprocess.run(iverilog_cmd, shell=True, check=True)

    # 生成vvp执行命令
    vvp_cmd = f"vvp -n wave -vcd"
    subprocess.run(vvp_cmd, shell=True, check=True)

    # 确定生成的vcd文件路径
    vcd_path = f"wave.vcd"

    os.remove("./wave")

    return vcd_path
