/*
Copyright © 2023 Jan Lauinger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aes

import (
	"encoding/hex"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

func TestAES128(t *testing.T) {
	assert := test.NewAssert(t)

	key := "7E24067817FAE0D743D6CE1F32539163"
	plaintext := "000102030405060708090A0B0C0D0E0F"  // +"101112131415161718191A1B1C1D1E1F"
	ciphertext := "5104A106168A72D9790D41EE8EDAD388" // +"EB2E1EFC46DA57C8FCE630DF9141BE28"
	Nonce := "006CB6DBC0543B59DA48D90B"
	Counter := 1

	keyAssign := StrToIntSlice(key, true)
	ptAssign := StrToIntSlice(plaintext, true)
	ctAssign := StrToIntSlice(ciphertext, true)
	nonceAssign := StrToIntSlice(Nonce, true)

	// witness values preparation
	assignment := AES128Wrapper{
		Key:        [16]frontend.Variable{},
		Counter:    Counter,
		Nonce:      [12]frontend.Variable{},
		Plaintext:  [16]frontend.Variable{},
		Ciphertext: [16]frontend.Variable{},
	}

	// assign values here because required to use make in assignment
	for i := 0; i < len(keyAssign); i++ {
		assignment.Key[i] = keyAssign[i]
	}
	for i := 0; i < len(ptAssign); i++ {
		assignment.Plaintext[i] = ptAssign[i]
	}
	for i := 0; i < len(ctAssign); i++ {
		assignment.Ciphertext[i] = ctAssign[i]
	}

	for i := 0; i < len(nonceAssign); i++ {
		assignment.Nonce[i] = nonceAssign[i]
	}

	// var circuit SHA256
	var circuit AES128Wrapper

	assert.CheckCircuit(&circuit, test.WithValidAssignment(&assignment), test.WithBackends(backend.GROTH16), test.WithCurves(ecc.BN254))
}

func StrToIntSlice(inputData string, hexRepresentation bool) []int {
	var byteSlice []byte
	if hexRepresentation {
		hexBytes, _ := hex.DecodeString(inputData)
		byteSlice = hexBytes
	} else {
		byteSlice = []byte(inputData)
	}

	var data []int
	for i := 0; i < len(byteSlice); i++ {
		data = append(data, int(byteSlice[i]))
	}
	return data
}
