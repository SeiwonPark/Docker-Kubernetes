/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START container_hello_app]
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", getLargestPrimeFactor)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// get largest prime factor of the given input from URL
func getLargestPrimeFactor(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	n, _ := strconv.Atoi(input) // convert string to int
	var maxPrime int = -1

	if n < 2 {
		fmt.Fprintf(w, "The input %d has no prime factor\n", n)
	} else {
		for n%2 == 0 {
			maxPrime = 2
			n /= 2
		}
		m := float64(n)
		var sqrt float64 = math.Sqrt(m)
		var abs int = int(sqrt)
		for i := 3; i <= abs+1; i += 2 {
			for n%i == 0 {
				maxPrime = i
				n = n / i
			}
		}
		if n > 2 {
			maxPrime = n
		}

		fmt.Fprintf(w, "The input is %s\n", input)
		fmt.Fprintf(w, "The largest prime factor is %d\n", int(maxPrime))
	}
}

// [END container_hello_app]
