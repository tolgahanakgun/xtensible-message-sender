package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"

	"gopkg.in/yaml.v2"
)

func main() {
	file, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	data := yaml.MapSlice{}
	err = yaml.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- example.yaml:\n%v\n\n", data)
	/*
		var bytes []byte
		bytes = append(bytes, []byte("a")...)
		bytes = append(bytes, []byte("s")...)
		bytes = append(bytes, []byte("m")...)
		bytes = append(bytes, []byte("l")...)
	*/
	var bytes []byte
	for index := 0; index < len(data); index++ {
		if t, ok := data[index].Key.(string); ok {
			switch t {
			case "ascii":
				// control non-ascii characters
				if str, ok := data[index].Value.(string); ok {
					bytes = append(bytes, []byte(str)...)
					fmt.Println("ascii->" + str)
				}
			case "int8":
				if i8, ok := data[index].Value.(int); ok {
					if i8 <= math.MaxInt8 && i8 >= math.MinInt8 {
						bytes = append(bytes, byte(i8))
						fmt.Println("int8->" + strconv.Itoa(int(i8)))
					}
				}
			case "int16":
				if i16, ok := data[index].Value.(int); ok {
					if i16 <= math.MaxInt16 && i16 >= math.MinInt16 {
						bytes = append(bytes, byte(int16(i16)))
						fmt.Println("int16->" + strconv.Itoa(int(i16)))
					}
				}
			case "int32":
				if i32, ok := data[index].Value.(int); ok {
					if i32 <= math.MaxInt32 && i32 >= math.MinInt32 {
						bytes = append(bytes, byte(int32(i32)))
						fmt.Println("int32->" + strconv.Itoa(int(i32)))
					}
				}
			case "int64":
				if i64, ok := data[index].Value.(int); ok {
					if i64 <= math.MaxInt64 && i64 >= math.MinInt64 {
						bytes = append(bytes, byte(int64(i64)))
						fmt.Println("int64->" + strconv.Itoa(int(i64)))
					}
				}
			case "uint8":
				//okk := data[index].Value.(uint8)
				//fmt.Println(okk)
				if ui8, ok := data[index].Value.(int); ok {
					if ui8 <= math.MaxUint8 && ui8 >= 0 {
						bytes = append(bytes, byte(uint8(ui8)))
						fmt.Println("uint8->" + strconv.Itoa(int(ui8)))
					}
				}
			case "uint16":
				if ui16, ok := data[index].Value.(int); ok {
					if ui16 <= math.MaxUint16 && ui16 >= 0 {
						bytes = append(bytes, byte(uint16(ui16)))
						fmt.Println("uint16->" + strconv.Itoa(int(ui16)))
					}
				}
			case "uint32":
				if ui32, ok := data[index].Value.(int); ok {
					if ui32 <= math.MaxUint32 && ui32 >= 0 {
						bytes = append(bytes, byte(uint32(ui32)))
						fmt.Println("uint32->" + strconv.Itoa(int(ui32)))
					}
				}
			case "uint64":
				if ui64, ok := data[index].Value.(int); ok {
					//if ui64 <= uint64(math.MaxUint64) && ui64 >= 0 {
					bytes = append(bytes, byte(uint64(ui64)))
					fmt.Println("ui64->" + strconv.Itoa(int(ui64)))
					//}
				}
			case "utf8":
			case "float":
			case "double":
			case "encoding":
				continue
			default:
				log.Fatalf("Wrong Type !!!")
			}
		}
		/*
			if str, ok := data[index].Value.(string); ok {
				fmt.Println("this is a string " + string(str))
			}
			if integer, ok := data[index].Value.(int); ok {
				fmt.Println("this is an integer " + strconv.Itoa(int(integer)))
			}
		*/
	}
	fmt.Println(bytes)
}
