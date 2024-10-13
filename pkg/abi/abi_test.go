package abi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
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

func getTypeName(descriptor ClarityTypeDescriptor) string {
	if descriptor.Uint128 != nil {
		return "uint128"
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

func TestABI_UnmarshalSample(t *testing.T) {
	samplePath := filepath.Join("testdata", "xcall-proxy-abi.json")

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
