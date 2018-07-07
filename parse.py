import yaml
import binascii
import struct
import sys
import argparse
from pathlib import Path
from colorama import Fore, Style, init, Back

# colorama.init()
init(autoreset=True)
# default endian
byte_order = 'big'

def print_red(s):
    print(Fore.RED+str(s)+Style.RESET_ALL)

parser = argparse.ArgumentParser()
parser.add_argument("-f", "--file", type=str, help="file path for parsing the yaml message")
# TODO send over socket, dump to file and byte order
parser.add_argument("-ip", "--ip-address", help="ip address the message is sent to, if not given print to stout")
parser.add_argument("-p", "--port-number", help="port number the message is sent to")
parser.add_argument("-out", "--output-file", help="if given the message content will written to the file")
parser.add_argument("-ord", "--byte-order", help="define big or little endian", choices=['big','little'])

args = parser.parse_args()

# if no parameter given print usage
if not len(sys.argv) > 1:
    parser.print_usage()
    quit()

# check file exist
input_file = args.file
if not Path(input_file).is_file():
    print_red('The file is not there!')
    quit()

yaml_object = {}

try:
    yaml_object = yaml.load(open(input_file, 'r'))
except IOError as e:
    print_red(e)
except yaml.YAMLError as e:
    print_red(e)

buffer = bytearray()

for (key,value) in yaml_object.items():
    if key == 'ascii':
        buffer.extend(value.encode('ascii'))
    elif key == 'utf8':
        buffer.extend(value.encode('utf-8'))
    elif key == 'int8':
        buffer.extend(int(value).to_bytes(1, byteorder=byte_order,signed=True))
    elif key == 'int16':
        buffer.extend(int(value).to_bytes(2, byteorder=byte_order,signed=True))
    elif key == 'int32':
        buffer.extend(int(value).to_bytes(4, byteorder=byte_order,signed=True))
    elif key == 'int64':
        buffer.extend(int(value).to_bytes(8, byteorder=byte_order,signed=True))
    elif key == 'uint8':
        buffer.extend(int(value).to_bytes(1, byteorder=byte_order,signed=False))
    elif key == 'uint16':
        buffer.extend(int(value).to_bytes(2, byteorder=byte_order,signed=False))
    elif key == 'uint32':
        buffer.extend(int(value).to_bytes(4, byteorder=byte_order,signed=False))
    elif key == 'uint64':
        buffer.extend(int(value).to_bytes(8, byteorder=byte_order,signed=False))
    elif key == 'float':
        buffer.extend(struct.pack('>f',float(value)))

print (buffer.hex())