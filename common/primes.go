package common

import "fmt"

/* ListPrimes retorna la lista de los 'size' primeros numeros primos. Criba de Eratostenes */
func ListPrimes(size int) []bool {
	list := make([]bool, size)
	list[1] = true
	p := 2
	for {
		pSq := p * p
		if pSq >= size {
			break
		}

		for i := pSq; i < size; i += p {
			list[i] = true // I multiple de p
		}

		// siguiente primo
		for {
			p++
			if !list[p] {
				break
			}
		}
	}
	return list
}

/* PrimeFactors: return the list of prime factors of n*/
func PrimeFactors(n int) []int {
	var factors []int
	lim := n
	for i := 2; i*i <= lim; {
		if (n % i) != 0 {
			i++
		} else {
			fmt.Println(i)
			factors = append(factors, i)
			n = n / i
		}
	}
	return factors
}
