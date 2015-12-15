/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package imgtypes

import (
	"bytes"
	"fmt"
	//"encoding/gob"
	"strings"
)

/**
 * The Image struct will be what we use to represent a DRM-protected image in the
 * ledger.
 */
type Image struct{
	strVal string 		// The value we will use to represent the image
	url string			// The url that can be used to access this image online
	owner string 		// Represents the source of the image
	licensees []string 	// A list of keys representing those who have purchased a license to view the image
}

type Source struct {
	name string			// The name associated with the image
}

func (s Source) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, strings.Replace(s.name, " ", "&s&", -1))
	return b.Bytes(), nil
}

func (s *Source) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &s.name)
	s.name = strings.Replace(s.name, "&s&", " ", -1)
	return err
}

/*
func main() {
    src := Source{name: "Dale Photography"}
	fmt.Println("Source: ", src)
	fmt.Println("Marshalling source into binary...")
	encoded, err := src.MarshalBinary();
	if (err != nil) {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Marshalled:",encoded)
	fmt.Println("Demarshalling...")
	var newsrc Source
	err = newsrc.UnmarshalBinary(encoded)
	if(err != nil) {
		fmt.Println("Error:", err)
	}
	fmt.Println("Demarshalled:", newsrc)
}
*/