import yaml
import binascii

yaml_object = {}

try:
    f = open('example.yaml','r')
    yaml_object = yaml.load(f)
except IOError as e:
    print (e)
except yaml.YAMLError as e:
    print (e)

buffer = bytearray()

for (key,value) in yaml_object.items():
    if key == 'ascii':
        buffer.extend(value.encode())
    elif key == 'int8':
        buffer.extend(int(value).to_bytes(1, byteorder='big',signed=True))
    elif key == 'int16':
        buffer.extend(int(value).to_bytes(2, byteorder='big',signed=True))
    elif key == 'int32':
        buffer.extend(int(value).to_bytes(4, byteorder='big',signed=True))
    elif key == 'int64':
        buffer.extend(int(value).to_bytes(8, byteorder='big',signed=True))
    elif key == 'uint8':
        buffer.extend(int(value).to_bytes(1, byteorder='big',signed=False))
    elif key == 'uint16':
        buffer.extend(int(value).to_bytes(2, byteorder='big',signed=False))
    elif key == 'uint32':
        buffer.extend(int(value).to_bytes(4, byteorder='big',signed=False))
    elif key == 'uint64':
        buffer.extend(int(value).to_bytes(8, byteorder='big',signed=False))

print (buffer.hex())