/*******************************************************************************
 * Copyright 2018 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package define

import (
	"errors"
	"strconv"
	"strings"
)



const (
	// Bool indicates that the value is a bool,
	// stored in CommandValue's boolRes member.
	Bool  = "Bool"
	// String indicates that the value is a string,
	// stored in CommandValue's stringRes member.
	String = "String"
	// Uint8 indicates that the value is a uint8 that
	// is stored in CommandValue's NumericRes member.
	Uint8 ="Uint8"
	// Uint16 indicates that the value is a uint16 that
	// is stored in CommandValue's NumericRes member.
	Uint16 ="Uint16"
	// Uint32 indicates that the value is a uint32 that
	// is stored in CommandValue's NumericRes member.
	Uint32 ="Uint32"
	// Uint64 indicates that the value is a uint64 that
	// is stored in CommandValue's NumericRes member.
	Uint64 ="Uint64"
	// Int8 indicates that the value is a int8 that
	// is stored in CommandValue's NumericRes member.
	Int8 ="Int8"
	// Int16 indicates that the value is a int16 that
	// is stored in CommandValue's NumericRes member.
	Int16 ="Int16"
	// Int32 indicates that the value is a int32 that
	// is stored in CommandValue's NumericRes member.
	Int32 ="Int32"
	// Int64 indicates that the value is a int64 that
	// is stored in CommandValue's NumericRes member.
	Int64 ="Int64"
	// Float32 indicates that the value is a float32 that
	// is stored in CommandValue's NumericRes member.
	Float32 ="Float32"
	// Float64 indicates that the value is a float64 that
	// is stored in CommandValue's NumericRes member.
	Float64 ="Float64"
	// Binary indicates that the value is a binary payload that
	// is stored in CommandValue's ByteArrRes member.
	Binary ="Binary"
)




func CommandValueTypeTrans64String(RsType string, valueStr string) (interface{}, error){

	var RSValue interface{}
	//var valueType sdkModel.ValueType
	//valueType := GetsdkModelTypeInt(RsType)
	RSValue = valueStr

	if RsType == String {
		//var res string
		RSValue = valueStr
		return RSValue, nil
	}

	switch RsType {
	case Bool:
		//var res bool
		if strings.ToUpper(valueStr) == "TRUE" {
			RSValue = true
			return RSValue, nil
		}else if strings.ToUpper(valueStr) == "FALSE" {
			RSValue = true
			return RSValue, nil
		}
	case Uint8:
		//var res uint8
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = uint8(v)
			return RSValue, nil
		}
	case Uint16:
		//var res uint16
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = uint16(v)
			return RSValue, nil
		}
	case Uint32:
		//var res uint32
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = uint32(v)
			return RSValue, nil
		}
	case Uint64:
		//var res uint64
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = uint64(v)
			return RSValue, nil
		}
	case Int8:
		//var res int8
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = int8(v)
			return RSValue, nil
		}
	case Int16:
		//var res int16
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = int16(v)
			return RSValue, nil
		}
	case Int32:
		//var res int32
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = int32(v)
			return RSValue, nil
		}
	case Int64:
		//var res int64
		if v,err := strconv.Atoi(valueStr);err == nil{
			RSValue = int64(v)
			return RSValue, nil
		}
	case Float32:
		//var res float32
		v, err := strconv.ParseFloat(valueStr, 32)
		if err == nil{
			RSValue = float32(v)
			//RSValue = math.Float32frombits(binary.BigEndian.Uint32(byteValue))
			//fmt.Fprintf(os.Stdout, fmt.Sprintf("IN CommandValueFromBase64String reading[%v] => [%v] \n", valueStr, RSValue))
			return RSValue, nil
		}
	case Float64:
		//var res float64
		v, err := strconv.ParseFloat(valueStr, 64)
		if err == nil{
			RSValue = float64(v)
			//RSValue = math.Float64frombits(binary.BigEndian.Uint64(byteValue))
			return RSValue, nil
		}
	default:
		return RSValue, errors.New("default string value")
	}

	return RSValue, errors.New("Failed")
}





