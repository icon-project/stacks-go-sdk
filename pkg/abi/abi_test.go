package abi

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/clarity"
)

func TestClarityTypeDescriptor_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		jsonData string
		expected string
	}{
		{
			name: "uint128",
			jsonData: `{
                "uint128": {}
            }`,
			expected: "uint128",
		},
		{
			name: "string-ascii",
			jsonData: `{
                "string-ascii": { "length": 128 }
            }`,
			expected: "string-ascii(128)",
		},
		{
			name: "optional list",
			jsonData: `{
                "optional": {
                    "type": {
                        "list": {
                            "type": {
                                "string-ascii": {
                                    "length": 10
                                }
                            },
                            "length": 5
                        }
                    }
                }
            }`,
			expected: "optional(list[5](string-ascii(10)))",
		},
		{
			name: "response",
			jsonData: `{
                "response": {
                    "ok": "bool",
                    "error": "uint128"
                }
            }`,
			expected: "response(ok: bool, error: uint128)",
		},
		{
			name: "principal",
			jsonData: `{
                "principal": {}
            }`,
			expected: "principal",
		},
		{
			name: "trait_reference",
			jsonData: `{
                "trait_reference": {}
            }`,
			expected: "trait_reference",
		},
		{
			name: "response with none",
			jsonData: `{
				"response": {
					"ok": "bool",
					"error": "none"
				}
			}`,
			expected: "response(ok: bool, error: none)",
		},
		{
			name: "tuple",
			jsonData: `{
				"tuple": [
					{
						"name": "network-id",
						"type": {
							"string-ascii": {
								"length": 128
							}
						}
					},
					{
						"name": "conn-sn",
						"type": "int128"
					}
				]
			}`,
			expected: "tuple(network-id: string-ascii(128), conn-sn: int128)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var descriptor ClarityTypeDescriptor
			err := json.Unmarshal([]byte(tc.jsonData), &descriptor)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			typeName := getTypeName(descriptor)
			if typeName != tc.expected {
				t.Errorf("Expected type '%s', got '%s'", tc.expected, typeName)
			}
		})
	}
}

func TestABI_UnmarshalXcallProxy(t *testing.T) {
	samplePath := filepath.Join("testdata", "xcall-proxy.abi")

	data, err := os.ReadFile(samplePath)
	if err != nil {
		t.Fatalf("Failed to read sample.json: %v", err)
	}

	var abi ABI
	err = json.Unmarshal(data, &abi)
	if err != nil {
		t.Fatalf("Failed to unmarshal sample.json: %v", err)
	}

	t.Run("ClarityVersion", func(t *testing.T) {
		expectedClarityVersion := "Clarity2"
		if abi.ClarityVersion != expectedClarityVersion {
			t.Errorf("Expected Clarity version '%s', got '%s'", expectedClarityVersion, abi.ClarityVersion)
		}
	})

	t.Run("Epoch", func(t *testing.T) {
		expectedEpoch := "Epoch30"
		if abi.Epoch != expectedEpoch {
			t.Errorf("Expected epoch '%s', got '%s'", expectedEpoch, abi.Epoch)
		}
	})

	t.Run("Functions", func(t *testing.T) {
		expectedFunctionNames := []string{
			"execute-call",
			"execute-rollback",
			"get-fee",
			"get-network-address",
			"get-network-id",
			"get-protocol-fee",
			"handle-error",
			"handle-message",
			"send-call",
			"send-call-message",
			"set-admin",
			"set-contract-owner",
			"set-data",
			"set-default-connection",
			"set-protocol-fee",
			"set-protocol-fee-handler",
			"set-trusted-protocols",
			"upgrade",
			"verify-success",
			"get-contract-owner",
			"get-current-implementation",
			"get-current-proxy",
			"get-data",
			"is-contract-owner",
			"is-current-implementation",
		}

		if len(abi.Functions) != len(expectedFunctionNames) {
			t.Errorf("Expected %d functions, got %d", len(expectedFunctionNames), len(abi.Functions))
		}

		functionNameSet := make(map[string]bool)
		for _, fn := range abi.Functions {
			functionNameSet[fn.Name] = true
		}

		for _, name := range expectedFunctionNames {
			if !functionNameSet[name] {
				t.Errorf("Expected function '%s' not found in ABI", name)
			}
		}
	})

	t.Run("Maps", func(t *testing.T) {
		expectedMapNames := []string{
			"data-storage",
		}

		if len(abi.Maps) != len(expectedMapNames) {
			t.Errorf("Expected %d maps, got %d", len(expectedMapNames), len(abi.Maps))
		}

		mapNameSet := make(map[string]bool)
		for _, m := range abi.Maps {
			mapNameSet[m.Name] = true
		}

		for _, name := range expectedMapNames {
			if !mapNameSet[name] {
				t.Errorf("Expected map '%s' not found in ABI", name)
			}
		}
	})

	t.Run("Variables", func(t *testing.T) {
		expectedVariableNames := []string{
			"CONTRACT_NAME",
			"err-not-current-implementation",
			"err-not-owner",
			"contract-owner",
			"current-logic-implementation",
			"current-proxy",
		}

		if len(abi.Variables) != len(expectedVariableNames) {
			t.Errorf("Expected %d variables, got %d", len(expectedVariableNames), len(abi.Variables))
		}

		variableNameSet := make(map[string]bool)
		for _, v := range abi.Variables {
			variableNameSet[v.Name] = true
		}

		for _, name := range expectedVariableNames {
			if !variableNameSet[name] {
				t.Errorf("Expected variable '%s' not found in ABI", name)
			}
		}
	})

	t.Run("Tokens", func(t *testing.T) {
		if len(abi.FungibleTokens) != 0 {
			t.Errorf("Expected 0 fungible tokens, got %d", len(abi.FungibleTokens))
		}

		if len(abi.NonFungibleTokens) != 0 {
			t.Errorf("Expected 0 non-fungible tokens, got %d", len(abi.NonFungibleTokens))
		}
	})

	t.Run("TypeDescriptors", func(t *testing.T) {
		typeDescriptors := collectAllTypeDescriptors(&abi)

		for i, desc := range typeDescriptors {
			typeName := getTypeName(desc)
			if typeName == "unknown" {
				t.Errorf("Type descriptor at index %d is unknown: %+v", i, desc)
			}
		}
	})
}

func TestABI_UnmarshalCentralizedConnection(t *testing.T) {
	samplePath := filepath.Join("testdata", "centralized-connection.abi")

	data, err := os.ReadFile(samplePath)
	if err != nil {
		t.Fatalf("Failed to read centralized-connection.abi: %v", err)
	}

	var abi ABI
	err = json.Unmarshal(data, &abi)
	if err != nil {
		t.Fatalf("Failed to unmarshal centralized-connection.abi: %v", err)
	}

	t.Run("ClarityVersion", func(t *testing.T) {
		expectedClarityVersion := "Clarity2"
		if abi.ClarityVersion != expectedClarityVersion {
			t.Errorf("Expected Clarity version '%s', got '%s'", expectedClarityVersion, abi.ClarityVersion)
		}
	})

	t.Run("Epoch", func(t *testing.T) {
		expectedEpoch := "Epoch30"
		if abi.Epoch != expectedEpoch {
			t.Errorf("Expected epoch '%s', got '%s'", expectedEpoch, abi.Epoch)
		}
	})

	t.Run("Maps", func(t *testing.T) {
		expectedMapNames := []string{
			"message-fees",
			"receipts",
			"response-fees",
		}

		if len(abi.Maps) != len(expectedMapNames) {
			t.Errorf("Expected %d maps, got %d", len(expectedMapNames), len(abi.Maps))
		}

		mapNameSet := make(map[string]bool)
		for _, m := range abi.Maps {
			mapNameSet[m.Name] = true
		}

		for _, name := range expectedMapNames {
			if !mapNameSet[name] {
				t.Errorf("Expected map '%s' not found in ABI", name)
			}
		}
	})

	t.Run("Functions", func(t *testing.T) {
		expectedFunctionNames := []string{
			"emit-message-event",
			"is-admin",
			"is-xcall",
			"claim-fees",
			"initialize",
			"recv-message",
			"send-message",
			"set-admin",
			"set-fee",
			"get-admin",
			"get-conn-sn",
			"get-fee",
			"get-receipt",
			"get-xcall",
		}

		if len(abi.Functions) != len(expectedFunctionNames) {
			t.Errorf("Expected %d functions, got %d", len(expectedFunctionNames), len(abi.Functions))
		}

		functionNameSet := make(map[string]bool)
		for _, fn := range abi.Functions {
			functionNameSet[fn.Name] = true
		}

		for _, name := range expectedFunctionNames {
			if !functionNameSet[name] {
				t.Errorf("Expected function '%s' not found in ABI", name)
			}
		}
	})

	t.Run("Variables", func(t *testing.T) {
		expectedVariableNames := []string{
			"ERR_DUPLICATE_MESSAGE",
			"ERR_INVALID_FEE",
			"ERR_UNAUTHORIZED",
			"ERR_XCALL_NOT_SET",
			"admin",
			"conn-sn",
			"xcall",
		}

		if len(abi.Variables) != len(expectedVariableNames) {
			t.Errorf("Expected %d variables, got %d", len(expectedVariableNames), len(abi.Variables))
		}

		variableNameSet := make(map[string]bool)
		for _, v := range abi.Variables {
			variableNameSet[v.Name] = true
		}

		for _, name := range expectedVariableNames {
			if !variableNameSet[name] {
				t.Errorf("Expected variable '%s' not found in ABI", name)
			}
		}
	})

	t.Run("Tokens", func(t *testing.T) {
		if len(abi.FungibleTokens) != 0 {
			t.Errorf("Expected 0 fungible tokens, got %d", len(abi.FungibleTokens))
		}

		if len(abi.NonFungibleTokens) != 0 {
			t.Errorf("Expected 0 non-fungible tokens, got %d", len(abi.NonFungibleTokens))
		}
	})

	t.Run("TypeDescriptors", func(t *testing.T) {
		typeDescriptors := collectAllTypeDescriptors(&abi)

		for i, desc := range typeDescriptors {
			typeName := getTypeName(desc)
			if typeName == "unknown" {
				t.Errorf("Type descriptor at index %d is unknown: %+v", i, desc)
			}
		}
	})

	t.Run("SpecificFunctionOutputs", func(t *testing.T) {
		var initializeFn *Function
		for _, fn := range abi.Functions {
			if fn.Name == "initialize" {
				initializeFn = &fn
				break
			}
		}
		if initializeFn == nil {
			t.Fatalf("Function 'initialize' not found")
		}

		expectedOutput := "response(ok: bool, error: uint128)"
		typeName := getTypeName(initializeFn.Outputs.Type)
		if typeName != expectedOutput {
			t.Errorf("Expected output type '%s', got '%s'", expectedOutput, typeName)
		}
	})

	t.Run("SpecificMapTypes", func(t *testing.T) {
		var messageFeesMap *Map
		for _, m := range abi.Maps {
			if m.Name == "message-fees" {
				messageFeesMap = &m
				break
			}
		}
		if messageFeesMap == nil {
			t.Fatalf("Map 'message-fees' not found")
		}

		expectedKeyType := "tuple(network-id: string-ascii(128))"
		keyTypeName := getTypeName(messageFeesMap.Key)
		if keyTypeName != expectedKeyType {
			t.Errorf("Expected key type '%s', got '%s'", expectedKeyType, keyTypeName)
		}

		expectedValueType := "uint128"
		valueTypeName := getTypeName(messageFeesMap.Value)
		if valueTypeName != expectedValueType {
			t.Errorf("Expected value type '%s', got '%s'", expectedValueType, valueTypeName)
		}
	})
}

func TestClarityTypeToClarityValue(t *testing.T) {
	testCases := []struct {
		name       string
		descriptor ClarityTypeDescriptor
		value      interface{}
		expected   clarity.ClarityValue
	}{
		{
			name: "uint128",
			descriptor: ClarityTypeDescriptor{
				Uint128: &struct{}{},
			},
			value: "123456789012345678901234567890",
			expected: func() clarity.ClarityValue {
				u, err := clarity.NewUInt("123456789012345678901234567890")
				if err != nil {
					t.Fatalf("Failed to create expected UInt: %v", err)
				}
				return u
			}(),
		},
		{
			name: "int128",
			descriptor: ClarityTypeDescriptor{
				Int128: &struct{}{},
			},
			value: "-123456789012345678901234567890",
			expected: func() clarity.ClarityValue {
				i, err := clarity.NewInt("-123456789012345678901234567890")
				if err != nil {
					t.Fatalf("Failed to create expected Int: %v", err)
				}
				return i
			}(),
		},
		{
			name: "bool true",
			descriptor: ClarityTypeDescriptor{
				Bool: &struct{}{},
			},
			value:    true,
			expected: clarity.NewBool(true),
		},
		{
			name: "bool false",
			descriptor: ClarityTypeDescriptor{
				Bool: &struct{}{},
			},
			value:    false,
			expected: clarity.NewBool(false),
		},
		{
			name: "string-ascii",
			descriptor: ClarityTypeDescriptor{
				StringASCII: &StringASCIIDescriptor{
					Length: 128,
				},
			},
			value: "Hello, Clarity!",
			expected: func() clarity.ClarityValue {
				s, err := clarity.NewStringASCII("Hello, Clarity!")
				if err != nil {
					t.Fatalf("Failed to create expected StringASCII: %v", err)
				}
				return s
			}(),
		},
		{
			name: "string-utf8",
			descriptor: ClarityTypeDescriptor{
				StringUTF8: &StringUTF8Descriptor{
					Length: 256,
				},
			},
			value: "こんにちは、Clarity!",
			expected: func() clarity.ClarityValue {
				s, err := clarity.NewStringUTF8("こんにちは、Clarity!")
				if err != nil {
					t.Fatalf("Failed to create expected StringUTF8: %v", err)
				}
				return s
			}(),
		},
		{
			name: "buffer",
			descriptor: ClarityTypeDescriptor{
				Buffer: &BufferDescriptor{
					Length: 2048,
				},
			},
			value: "0xabcdef123456",
			expected: func() clarity.ClarityValue {
				decoded, err := hex.DecodeString("abcdef123456")
				if err != nil {
					t.Fatalf("Failed to decode buffer hex: %v", err)
				}
				return clarity.NewBuffer(decoded)
			}(),
		},
		{
			name: "optional some",
			descriptor: ClarityTypeDescriptor{
				Optional: &OptionalDescriptor{
					Type: ClarityTypeDescriptor{
						Bool: &struct{}{},
					},
				},
			},
			value: true,
			expected: func() clarity.ClarityValue {
				return clarity.NewOptionSome(clarity.NewBool(true))
			}(),
		},
		{
			name: "optional none",
			descriptor: ClarityTypeDescriptor{
				Optional: &OptionalDescriptor{
					Type: ClarityTypeDescriptor{
						Bool: &struct{}{},
					},
				},
			},
			value:    nil,
			expected: clarity.NewOptionNone(),
		},
		{
			name: "list",
			descriptor: ClarityTypeDescriptor{
				List: &ListDescriptor{
					Type: ClarityTypeDescriptor{
						Bool: &struct{}{},
					},
					Length: 5,
				},
			},
			value: []interface{}{true, false, true, true, false},
			expected: func() clarity.ClarityValue {
				clarityList := []clarity.ClarityValue{
					clarity.NewBool(true),
					clarity.NewBool(false),
					clarity.NewBool(true),
					clarity.NewBool(true),
					clarity.NewBool(false),
				}
				return clarity.NewList(clarityList)
			}(),
		},
		{
			name: "tuple",
			descriptor: ClarityTypeDescriptor{
				Tuple: &TupleDescriptor{
					Fields: []TupleField{
						{
							Name: "field1",
							Type: ClarityTypeDescriptor{
								Bool: &struct{}{},
							},
						},
						{
							Name: "field2",
							Type: ClarityTypeDescriptor{
								Int128: &struct{}{},
							},
						},
					},
				},
			},
			value: map[string]interface{}{
				"field1": true,
				"field2": "-98765432109876543210",
			},
			expected: func() clarity.ClarityValue {
				tupleData := map[string]clarity.ClarityValue{
					"field1": clarity.NewBool(true),
					"field2": func() clarity.ClarityValue {
						i, err := clarity.NewInt("-98765432109876543210")
						if err != nil {
							t.Fatalf("Failed to create expected Int for tuple field2: %v", err)
						}
						return i
					}(),
				}
				return clarity.NewTuple(tupleData)
			}(),
		},
		{
			name: "response ok",
			descriptor: ClarityTypeDescriptor{
				Response: &ResponseDescriptor{
					Ok: &ClarityTypeDescriptor{
						Bool: &struct{}{},
					},
					Error: &ClarityTypeDescriptor{
						Uint128: &struct{}{},
					},
				},
			},
			value: map[string]interface{}{
				"ok": true,
			},
			expected: func() clarity.ClarityValue {
				return clarity.NewResponseOk(clarity.NewBool(true))
			}(),
		},
		{
			name: "response error",
			descriptor: ClarityTypeDescriptor{
				Response: &ResponseDescriptor{
					Ok: &ClarityTypeDescriptor{
						Bool: &struct{}{},
					},
					Error: &ClarityTypeDescriptor{
						Uint128: &struct{}{},
					},
				},
			},
			value: map[string]interface{}{
				"error": "999999999999999999999999999999",
			},
			expected: func() clarity.ClarityValue {
				return clarity.NewResponseErr(func() clarity.ClarityValue {
					u, err := clarity.NewUInt("999999999999999999999999999999")
					if err != nil {
						t.Fatalf("Failed to create expected UInt for response error: %v", err)
					}
					return u
				}())
			}(),
		},
		{
			name: "principal",
			descriptor: ClarityTypeDescriptor{
				Principal: &struct{}{},
			},
			value: "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH", // Updated to 42 characters
			expected: func() clarity.ClarityValue {
				return func() clarity.ClarityValue {
					p, err := clarity.StringToPrincipal("ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH")
					if err != nil {
						t.Fatalf("Failed to create expected Principal: %v", err)
					}
					return p
				}()
			}(),
		},
		{
			name: "trait_reference",
			descriptor: ClarityTypeDescriptor{
				TraitReference: &struct{}{},
			},
			value: "ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH",
			expected: func() clarity.ClarityValue {
				return func() clarity.ClarityValue {
					p, err := clarity.StringToPrincipal("ST15C893XJFJ6FSKM020P9JQDB5T7X6MQTXMBPAVH")
					if err != nil {
						t.Fatalf("Failed to create expected TraitReference: %v", err)
					}
					return p
				}()
			}(),
		},
		{
			name: "none",
			descriptor: ClarityTypeDescriptor{
				None: &struct{}{},
			},
			value:    nil,
			expected: clarity.NewOptionNone(),
		},
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ClarityTypeToClarityValue(tc.descriptor, tc.value)
			if err != nil {
				t.Fatalf("ClarityTypeToClarityValue failed: %v", err)
			}

			expectedBytes, err := tc.expected.Serialize()
			if err != nil {
				t.Fatalf("Failed to serialize expected value: %v", err)
			}

			actualBytes, err := actual.Serialize()
			if err != nil {
				t.Fatalf("Failed to serialize actual value: %v", err)
			}

			if !bytes.Equal(expectedBytes, actualBytes) {
				t.Errorf("Serialized ClarityValue instances do not match.\nExpected: %x\nActual:   %x", expectedBytes, actualBytes)
			}
		})
	}
}

func collectAllTypeDescriptors(abi *ABI) []ClarityTypeDescriptor {
	var descriptors []ClarityTypeDescriptor

	for _, m := range abi.Maps {
		descriptors = append(descriptors, m.Key, m.Value)
	}

	for _, fn := range abi.Functions {
		for _, arg := range fn.Args {
			descriptors = append(descriptors, arg.Type)
		}
		descriptors = append(descriptors, fn.Outputs.Type)
	}

	for _, v := range abi.Variables {
		descriptors = append(descriptors, v.Type)
	}

	return descriptors
}

func getTypeName(descriptor ClarityTypeDescriptor) string {
	if descriptor.Uint128 != nil {
		return "uint128"
	}
	if descriptor.Int128 != nil {
		return "int128"
	}
	if descriptor.Bool != nil {
		return "bool"
	}
	if descriptor.StringASCII != nil {
		return fmt.Sprintf("string-ascii(%d)", descriptor.StringASCII.Length)
	}
	if descriptor.StringUTF8 != nil {
		return fmt.Sprintf("string-utf8(%d)", descriptor.StringUTF8.Length)
	}
	if descriptor.Buffer != nil {
		return fmt.Sprintf("buffer(%d)", descriptor.Buffer.Length)
	}
	if descriptor.Optional != nil {
		return fmt.Sprintf("optional(%s)", getTypeName(descriptor.Optional.Type))
	}
	if descriptor.List != nil {
		return fmt.Sprintf("list[%d](%s)", descriptor.List.Length, getTypeName(descriptor.List.Type))
	}
	if descriptor.Tuple != nil {
		fields := []string{}
		for _, field := range descriptor.Tuple.Fields {
			fields = append(fields, fmt.Sprintf("%s: %s", field.Name, getTypeName(field.Type)))
		}
		return fmt.Sprintf("tuple(%s)", strings.Join(fields, ", "))
	}
	if descriptor.Response != nil {
		okType := "unknown"
		errorType := "unknown"
		if descriptor.Response.Ok != nil {
			okType = getTypeName(*descriptor.Response.Ok)
		}
		if descriptor.Response.Error != nil {
			errorType = getTypeName(*descriptor.Response.Error)
		}
		return fmt.Sprintf("response(ok: %s, error: %s)", okType, errorType)
	}
	if descriptor.Principal != nil {
		return "principal"
	}
	if descriptor.TraitReference != nil {
		return "trait_reference"
	}
	if descriptor.None != nil {
		return "none"
	}
	return "unknown"
}
