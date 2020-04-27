package main

import (
	log "github.com/sirupsen/logrus"
	"math/big"
)

func Euler(p *big.Int, q *big.Int) *big.Int {
	fi := new(big.Int)
	return fi.Mul(p.Sub(p, big.NewInt(1)), q.Sub(q, big.NewInt(1)))
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
	aLatest := true
	for r[len(r)-1].Cmp(big.NewInt(0)) != 0 {
		tempQ = big.NewInt(0)
		tempR = big.NewInt(0)
		if a.Cmp(b) == 1 {
			tempQ.DivMod(a, b, tempR)
			*a = *tempR
			aLatest = true
		} else {
			tempQ.DivMod(b, a, tempR)
			*b = *tempR
			aLatest = false
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
	//if a.Cmp(big.NewInt(0)) == 0 {
	x1 = big.NewInt(0)
	y1 = big.NewInt(1)
	//} else {
	//	x1 = big.NewInt(0)
	//	y1 = big.NewInt(1)
	//}
	for i := 0; i < len(q); i++ {
		//log.Infoln("На входе x1:", x1, "iteration: ", i)
		//log.Infoln("На входе y1:", y1, "iteration: ", i)
		x = big.NewInt(0)
		temp.Mul(x1, &q[len(q)-1-i])
		x.Sub(y1, temp)
		*y = *x1
		//log.Infoln("x:", x, "iteration: ", i)
		//log.Infoln("y:", y, "iteration: ", i)
		*x1 = *x
		*y1 = *y
		//log.Infoln("На выходе x1:", x1, "iteration: ", i)
		//log.Infoln("На выходе y1:", y1, "iteration: ", i)
	}
	if !aLatest {
		log.Info("Coefficient X: ", x)
		log.Info("Coefficient Y: ", y)
		return r[len(r)-2], x
	} else {
		log.Info("Coefficient X: ", y)
		log.Info("Coefficient Y: ", x)
		return r[len(r)-2], y
	}
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
func binPowBigIntModN(x *big.Int, e *big.Int, n *big.Int) *big.Int {
	res := big.NewInt(1)
	m := big.NewInt(0)
	temp := big.NewInt(0)
	temp.DivMod(x, n, x)
	for e.Cmp(big.NewInt(0)) == 1 {
		if temp.DivMod(e, big.NewInt(2), m); m.Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, x)
			temp.DivMod(res, n, res)
		}

		x.Mul(x, x)
		temp.DivMod(x, n, x)
		e.Rsh(e, 1)
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
	log.Info("e: ", e)

	log.Info("---TEST---")
	gcd, x := modifiedGCD(big.NewInt(1071), big.NewInt(462))
	log.Infoln("GCD 1071 and 462:", gcd, ", D =", x)
	gcd, x = modifiedGCD(big.NewInt(462), big.NewInt(1071))
	log.Infoln("GCD 462 and 1071:", gcd, ", D =", x)
	gcd, x = modifiedGCD(big.NewInt(3), big.NewInt(5))
	log.Infoln("GCD 3 and 5:", gcd, ", D =", x)
	log.Info("In big.Int 11^5(mod 7)=", binPowBigIntModN(big.NewInt(11), big.NewInt(5), big.NewInt(7)))
	log.Info("In big.Int 32333^7(mod 11)=", binPowBigIntModN(big.NewInt(32333), big.NewInt(7), big.NewInt(11)))
	log.Info("---TEST---")

	temp := new(big.Int)
	*temp = *e
	nod, d := modifiedGCD(temp, fi)

	log.Infoln("GCD:", nod, ", D =", d)
	X1 := new(big.Int)
	X1.SetString("832192044845038413443817859011", 10)
	log.Info("X1: ", X1)
	log.Info("e: ", e)
	log.Info("n: ", n)
	Y1 := binPowBigIntModN(X1, e, n)
	log.Info("Y1 = ", Y1)
	//tempD:=new(big.Int)
	tempD := d.String()
	temp.SetString(tempD, 10)
	log.Info("d: ", d)
	log.Info("X1 = ", binPowBigIntModN(Y1, d, n))
	Y2 := new(big.Int)
	Y2.SetString("381868705201087633178499417569", 10)
	log.Info("Y2: ", Y2)
	//log.Info("d: ", d)
	log.Info("temp: ", temp)
	log.Info("n: ", n)
	X2 := binPowBigIntModN(Y2, temp, n)
	log.Info("X2 = ", X2)
}

// 1005505123313855881766634063697   n
// 1005505123313853867461474914416   fi
// 4136825437389004471896504395      d
// 832192044845038413443817859011    x1
