package main

import (
	log "github.com/sirupsen/logrus"
	"math/big"
)

func Euler(p *big.Int, q *big.Int) *big.Int {
	fi := new(big.Int)
	return fi.Mul(p.Div(p, big.NewInt(1)), q.Div(q, big.NewInt(1)))
}
func modifiedGCD(a *big.Int, b *big.Int) (*big.Int, *big.Int) {
	q := make([]big.Int, 0)
	r := make([]*big.Int, 0)
	temp := new(big.Int)
	tempQ := big.NewInt(0)
	tempR := new(big.Int)
	if a.Cmp(b) == 1 {
		tempQ.DivMod(a, b, tempR)
		*a = *tempR
	} else {
		tempQ.DivMod(b, a, tempR)
		*b = *tempR
	}
	r = append(r, tempR)
	q = append(q, *tempQ)
	for r[len(r)-1].Cmp(big.NewInt(0)) != 0 {
		tempQ = big.NewInt(0)
		tempR = big.NewInt(0)
		if a.Cmp(b) == 1 {
			tempQ.DivMod(a, b, tempR)
			*a = *tempR
		} else {
			tempQ.DivMod(b, a, tempR)
			*b = *tempR
		}
		r = append(r, tempR)
		q = append(q, *tempQ)
	}
	for i := range q {
		log.Info("q[", i, "]: ", q[i].String())
	}
	for i := range r {
		log.Info("r[", i, "]: ", r[i].String())
	}
	x1 := new(big.Int)
	y1 := new(big.Int)
	x := new(big.Int)
	y := new(big.Int)
	if a.Cmp(big.NewInt(0)) == 0 {
		x1 = big.NewInt(0)
		y1 = big.NewInt(1)
	} else {
		x1 = big.NewInt(1)
		y1 = big.NewInt(0)
	}
	for i := 0; i < len(q); i++ {
		log.Info()
		x = big.NewInt(0)
		temp.Mul(x1, &q[len(q)-1-i])
		x.Sub(y1, temp)
		*y = *x1
		log.Infoln("x:", x, "iteration:", i)
		log.Infoln("y:", y, "iteration:", i)
		*x1 = *x
		*y1 = *y
		log.Info()
	}
	return r[len(r)-2], x
}

func binPow(a int, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res *= a
		}

		a *= a
		n >>= 1
	}
	return res
}
func binPowBigInt(a *big.Int, n *big.Int) *big.Int {
	res := big.NewInt(1)
	m := big.NewInt(0)
	temp := big.NewInt(0)
	for n.Cmp(big.NewInt(0)) == 1 {
		if temp.DivMod(n, big.NewInt(2), m); m.Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, a)
		}

		a.Mul(a, a)
		n.Rsh(n, 1)
	}
	return res
}
func main() {
	p := new(big.Int)
	p.SetString("1101233370547069", 10)
	log.Info("p: ", p)

	q := new(big.Int)
	q.SetString("913071788602213", 10)
	log.Info("q: ", q)

	e := new(big.Int)
	e.SetString("441294907009469893617176298995", 10)
	log.Info("e: ", e)

	n := new(big.Int)
	n.Mul(p, q)
	log.Info("n: ", n)
	fi := Euler(p, q)
	log.Info("fi: ", fi)
	gcd, x := modifiedGCD(big.NewInt(1071), big.NewInt(462))
	log.Infoln("GCD 1071 and 462:", gcd, ", D =", x)
	log.Info("2^9=", binPow(2, 9))
	log.Info("7^5=", binPow(7, 5))
	log.Info("In big.Int 7^5=", binPowBigInt(big.NewInt(7), big.NewInt(5)))

}
