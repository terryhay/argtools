package argValueTools

/*type ArgValueTools struct {
	yamlString2ValueType map[string]argParserConfig.ArgValueType
	valueConvertors      map[argParserConfig.ArgValueType]func(string) (variant.Variant, error)
}

func NewArgValueTools() ArgValueTools {
	return ArgValueTools{
		yamlString2ValueType: map[string]argParserConfig.ArgValueType{
			"s": argParserConfig.ArgValueTypeString,
			"i": argParserConfig.ArgValueTypeInt,
			"f": argParserConfig.ArgValueTypeFloat,
		},
		valueConvertors: map[argParserConfig.ArgValueType]func(string) (variant.Variant, error){
			argParserConfig.ArgValueTypeString: String2string,
			argParserConfig.ArgValueTypeInt:    String2int,
			argParserConfig.ArgValueTypeFloat:  String2float,
		},
	}
}

func (i ArgValueTools) CheckTypeAndValues(valuesType string, values []string) *argtoolsError.Error {
	_, contains := i.yamlString2ValueType[valuesType]
	if !contains {
		return argtoolsError.NewError(0, fmt.Errorf("err"))
	}

	return nil
}

func String2string(value string) (variant.Variant, error) {
	return variant.New(value), nil
}

func String2int(value string) (variant.Variant, error) {
	intValue, err := strconv.Atoi(value)
	return variant.New(intValue), err
}

func String2float(value string) (variant.Variant, error) {
	floatValue, err := strconv.ParseFloat(value, 64)
	return variant.New(floatValue), err
}

func StringValue2Variant(valuesType argParserConfig.ArgValueType, value string) (res variant.Variant, err error) {
	switch valuesType {
	case argParserConfig.ArgValueTypeString:
		res = variant.New(value)
	case argParserConfig.ArgValueTypeInt:
		var vd int

		vd, err = strconv.Atoi(value)
		if err != nil {
			return variant.Variant{}, fmt.Errorf("convertin string to int error: %v", err)
		}
		res = variant.New(vd)

	case argParserConfig.ArgValueTypeFloat:
		var vf float64

		vf, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return variant.Variant{}, fmt.Errorf("convertin string to float error: %v", err)
		}
		res = variant.New(vf)

	default:
		return variant.Variant{}, fmt.Errorf("unexpected value type: %v", valuesType)
	}

	return res, nil
}

func StringValues2Variants(valuesType argParserConfig.ArgValueType, values []string) (res []variant.Variant, err error) {
	valueCount := len(values)
	if valueCount == 0 {
		return nil, nil
	}
	res = make([]variant.Variant, 0, valueCount)

	switch valuesType {
	case argParserConfig.ArgValueTypeString:
		for _, v := range values {
			res = append(res, variant.New(v))
		}
	case argParserConfig.ArgValueTypeInt:
		var vd int
		for _, v := range values {
			vd, err = strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("convertin string to int error: %v", err)
			}
			res = append(res, variant.New(vd))
		}
	case argParserConfig.ArgValueTypeFloat:
		var vf float64
		for _, v := range values {
			vf, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, fmt.Errorf("convertin string to float error: %v", err)
			}
			res = append(res, variant.New(vf))
		}
	default:
		return nil, fmt.Errorf("unexpected value type: %v", valuesType)
	}

	return res, nil
}//*/
