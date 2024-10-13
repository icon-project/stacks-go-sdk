package abi

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
)

type ABI struct {
	Maps              []Map         `json:"maps"`
	Epoch             string        `json:"epoch"`
	Functions         []Function    `json:"functions"`
	Variables         []Variable    `json:"variables"`
	ClarityVersion    string        `json:"clarity_version"`
	FungibleTokens    []interface{} `json:"fungible_tokens"`
	NonFungibleTokens []interface{} `json:"non_fungible_tokens"`
}

type Map struct {
	Key   ClarityTypeDescriptor `json:"key"`
	Name  string                `json:"name"`
	Value ClarityTypeDescriptor `json:"value"`
}

type Function struct {
	Args    []Argument       `json:"args"`
	Name    string           `json:"name"`
	Access  string           `json:"access"`
	Outputs OutputDescriptor `json:"outputs"`
}

type Argument struct {
	Name string                `json:"name"`
	Type ClarityTypeDescriptor `json:"type"`
}

type OutputDescriptor struct {
	Type ClarityTypeDescriptor `json:"type"`
}

type Variable struct {
	Name   string                `json:"name"`
	Type   ClarityTypeDescriptor `json:"type"`
	Access string                `json:"access"`
}

type ClarityTypeDescriptor struct {
	Raw json.RawMessage `json:"-"`

	StringASCII    *StringASCIIDescriptor `json:"string-ascii,omitempty"`
	StringUTF8     *StringUTF8Descriptor  `json:"string-utf8,omitempty"`
	Uint128        *struct{}              `json:"uint128,omitempty"`
	Bool           *struct{}              `json:"bool,omitempty"`
	Buffer         *BufferDescriptor      `json:"buffer,omitempty"`
	Optional       *OptionalDescriptor    `json:"optional,omitempty"`
	List           *ListDescriptor        `json:"list,omitempty"`
	Response       *ResponseDescriptor    `json:"response,omitempty"`
	Principal      *struct{}              `json:"principal,omitempty"`
	TraitReference *struct{}              `json:"trait_reference,omitempty"`
	None           *struct{}              `json:"none,omitempty"`
}

type StringASCIIDescriptor struct {
	Length int `json:"length"`
}

type StringUTF8Descriptor struct {
	Length int `json:"length"`
}

type BufferDescriptor struct {
	Length int `json:"length"`
}

type OptionalDescriptor struct {
	Type ClarityTypeDescriptor `json:"type"`
}

type ListDescriptor struct {
	Type   ClarityTypeDescriptor `json:"type"`
	Length int                   `json:"length"`
}

type TraitReferenceDescriptor struct {
}

type ResponseDescriptor struct {
	Ok    *ClarityTypeDescriptor `json:"ok,omitempty"`
	Error *ClarityTypeDescriptor `json:"error,omitempty"`
}

func (o *OptionalDescriptor) UnmarshalJSON(data []byte) error {
	var typeStr string
	if err := json.Unmarshal(data, &typeStr); err == nil {
		var desc ClarityTypeDescriptor
		if err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, typeStr)), &desc); err != nil {
			return fmt.Errorf("failed to unmarshal optional type string '%s': %v", typeStr, err)
		}
		o.Type = desc
		return nil
	}

	var temp struct {
		Type ClarityTypeDescriptor `json:"type"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to unmarshal optional as object: %v", err)
	}
	o.Type = temp.Type
	return nil
}

func (c *ClarityTypeDescriptor) UnmarshalJSON(data []byte) error {
	c.Raw = data

	var typeStr string
	if err := json.Unmarshal(data, &typeStr); err == nil {
		switch typeStr {
		case "uint128":
			c.Uint128 = &struct{}{}
		case "bool":
			c.Bool = &struct{}{}
		case "principal":
			c.Principal = &struct{}{}
		case "trait_reference":
			c.TraitReference = &struct{}{}
		case "none":
			c.None = &struct{}{}
		default:
			return fmt.Errorf("unknown Clarity type string: %s", typeStr)
		}
		return nil
	}

	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return fmt.Errorf("ClarityTypeDescriptor should be a JSON object: %v", err)
	}

	for key, value := range tempMap {
		switch key {
		case "string-ascii":
			var desc StringASCIIDescriptor
			if err := json.Unmarshal(value, &desc); err != nil {
				return fmt.Errorf("failed to unmarshal string-ascii: %v", err)
			}
			c.StringASCII = &desc
		case "string-utf8":
			var desc StringUTF8Descriptor
			if err := json.Unmarshal(value, &desc); err != nil {
				return fmt.Errorf("failed to unmarshal string-utf8: %v", err)
			}
			c.StringUTF8 = &desc
		case "uint128":
			c.Uint128 = &struct{}{}
		case "bool":
			c.Bool = &struct{}{}
		case "buffer":
			var desc BufferDescriptor
			if err := json.Unmarshal(value, &desc); err != nil {
				return fmt.Errorf("failed to unmarshal buffer: %v", err)
			}
			c.Buffer = &desc
		case "optional":
			var desc OptionalDescriptor
			if err := json.Unmarshal(value, &desc); err != nil {
				return fmt.Errorf("failed to unmarshal optional: %v", err)
			}
			c.Optional = &desc
		case "list":
			var desc ListDescriptor
			if err := json.Unmarshal(value, &desc); err != nil {
				return fmt.Errorf("failed to unmarshal list: %v", err)
			}
			c.List = &desc
		case "response":
			var desc ResponseDescriptor
			if err := json.Unmarshal(value, &desc); err != nil {
				return fmt.Errorf("failed to unmarshal response: %v", err)
			}
			c.Response = &desc
		case "principal":
			c.Principal = &struct{}{}
		case "trait_reference":
			c.TraitReference = &struct{}{}
		default:
			return fmt.Errorf("unknown Clarity type key: %s", key)
		}
	}

	return nil
}

func ClarityTypeToClarityValue(descriptor ClarityTypeDescriptor, value interface{}) (clarity.ClarityValue, error) {
	switch {
	case descriptor.Uint128 != nil:
		valStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for uint128 value")
		}
		return clarity.NewUInt(valStr)
	case descriptor.Bool != nil:
		valBool, ok := value.(bool)
		if !ok {
			return nil, fmt.Errorf("expected bool for bool value")
		}
		return clarity.NewBool(valBool), nil
	case descriptor.Buffer != nil:
		valBytes, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for buffer value")
		}
		decoded, err := hex.DecodeString(strings.TrimPrefix(valBytes, "0x"))
		if err != nil {
			return nil, fmt.Errorf("invalid hex string for buffer: %v", err)
		}
		return clarity.NewBuffer(decoded), nil
	case descriptor.StringASCII != nil:
		valStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for string-ascii value")
		}
		return clarity.NewStringASCII(valStr)
	case descriptor.StringUTF8 != nil:
		valStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for string-utf8 value")
		}
		return clarity.NewStringUTF8(valStr)
	case descriptor.Optional != nil:
		if value == nil {
			return clarity.NewOptionNone(), nil
		}
		someValue, err := ClarityTypeToClarityValue(descriptor.Optional.Type, value)
		if err != nil {
			return nil, fmt.Errorf("failed to create OptionSome: %v", err)
		}
		return clarity.NewOptionSome(someValue), nil
	case descriptor.List != nil:
		listItems, ok := value.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected slice for list value")
		}
		clarityList := make([]clarity.ClarityValue, 0, len(listItems))
		for _, item := range listItems {
			clarityVal, err := ClarityTypeToClarityValue(descriptor.List.Type, item)
			if err != nil {
				return nil, fmt.Errorf("failed to create list item: %v", err)
			}
			clarityList = append(clarityList, clarityVal)
		}
		return clarity.NewList(clarityList), nil
	case descriptor.Response != nil:
		respMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected map for response value")
		}
		if okVal, exists := respMap["ok"]; exists {
			clarityOk, err := ClarityTypeToClarityValue(*descriptor.Response.Ok, okVal)
			if err != nil {
				return nil, fmt.Errorf("failed to create response ok: %v", err)
			}
			return clarity.NewResponseOk(clarityOk), nil
		}
		if errVal, exists := respMap["error"]; exists {
			clarityErr, err := ClarityTypeToClarityValue(*descriptor.Response.Error, errVal)
			if err != nil {
				return nil, fmt.Errorf("failed to create response error: %v", err)
			}
			return clarity.NewResponseErr(clarityErr), nil
		}
		return nil, fmt.Errorf("response must have either 'ok' or 'error'")
	case descriptor.Principal != nil:
		valStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for principal value")
		}
		return clarity.StringToPrincipal(valStr)
	case descriptor.TraitReference != nil:
		valStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected string for trait_reference value")
		}
		return clarity.StringToPrincipal(valStr)
	case descriptor.None != nil:
		return clarity.NewOptionNone(), nil
	default:
		return nil, fmt.Errorf("unsupported Clarity type descriptor")
	}
}
