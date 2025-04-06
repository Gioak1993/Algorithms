package main

import "fmt"


func main () {

	EuclidsAlgorithm(15, 12)
	EuclidsAlgorithm(35, 14)
	EuclidsAlgorithm(10, 10)
	EuclidsAlgorithm(75, 15)
	EuclidsAlgorithm(7, 9)
	EuclidsAlgorithm(1090901401, 1092091)
	EuclidsAlgorithm(111, 11)
	EuclidsAlgorithm(999, 729)
	Extended_euclid(21,15)
	Extended_euclid(19, 17)
	Extended_euclid(60, 13)
	
}

func EuclidsAlgorithm (m int32, n int32) (int32, error) {
	if m < 0 {
		m = -m  //make them positive 
	}
	if n < 0 {
		n = -n
	}

	if m < n {
		m, n = n, m //make the biggest value be m 
	}

	fmt.Printf("Euclides Algorithm for (%d, %d)\n" , m, n)

	if n == 0 && m != 0 {
		return m, nil
	}else if n == 0 && m == 0{
		return 0, fmt.Errorf("the number was zero") 
	}

	for n > 0 {

		m, n = n, m % n  //make m = n and n = m % n
		fmt.Printf("GDC(%d, %d)\n", m, n)
	}

	fmt.Printf(" = %d\n", m)

	return m, nil 
}

// Bezout's Theorem and Extended Euclid's Algorithm
// Next, we will go over a simple extension to Euclid's algorithm that will be quite helpful in developing the RSA cryptosystem. This is a theorem by the famous French mathematician Etienne Bezout.

// Given two numbers say  16
//   and  12
//  , we consider all numbers that can be written in the form  16𝑠+12𝑡
//   for integers  16,12
//  .

// With  𝑠=𝑡=1
//  , we get  28
//  .
// With  𝑠=1,𝑡=−1
//  , we get  4
//  .
// With  𝑠=−2,𝑡=3
//  , we get  4
//  .
// With  𝑠=2,𝑡=−2
//  , we obtain  8
//  .
// ...
// What can we say about the (non-negative) numbers that we can generate in this fashion? Notice that they are all multiples of  4
//   and in fact,  4
//   is the smallest such natural number. In fact, we cannot write  16𝑠+12𝑡=1
//   for any choice of integers  𝑠,𝑡
//  .

// Given  𝑚,𝑛
//   where  𝑚≥1
//   and  𝑛≥0
//  , we can write  𝐺𝐶𝐷(𝑚,𝑛)=𝑠×𝑚+𝑡×𝑛
//   for integers  𝑠,𝑡
//  .
// We will give an "algorithmic" proof by extending Euclid's algorithm to not just generate the GCD of two numbers but also provide the "Bezout coefficients".

func Extended_euclid(m int , n int) (int, int, int){

	m0, n0 := m, n 	// lets remember the initial values

	s, t := 1, 0 	// m = s * m0 + t * n0
	s_hat, t_hat := 0, 1 // n = s_hat * m0 + t_hat * n0

	for n > 0 {


		q := m / n // compute the integer quotient 
		r := m % n // the reminder 

		// m = q * n + r, or alternatively, r = m - q * n = (s-q*s_hat) * m_0 + (t - q * t_hat)* n_0  

		a, b := s - q * s_hat, t - q * t_hat 
		m, n, s, t, s_hat, t_hat = n, r, s_hat, t_hat, a, b

		fmt.Printf("GCD(%d, %d) = GCD(%d, %d)\n", m0, n0, m, n)
		fmt.Printf("\t%d = %d*%d + %d*%d\n", m, s, m0, t, n0)
		fmt.Printf("\t%d = %d*%d + %d*%d\n", n, s_hat, m0, t_hat, n0)

	}

	return m, s, t

}
