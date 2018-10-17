// Copyright 2018 ING Bank N.V.
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

/*
This file contains the implementation of the Bulletproofs scheme proposed in the paper:
Bulletproofs: Short Proofs for Confidential Transactions and More
Benedikt Bunz, Jonathan Bootle, Dan Boneh, Andrew Poelstra, Pieter Wuille and Greg Maxwell
Asiacrypt 2008
*/

package zkrangeproof

import (
	"math"
	"math/big"
	"crypto/rand"
	"crypto/sha256"
	"github.com/ing-bank/zkrangeproof/go-ethereum/byteconversion"
	"errors"
)

var (
	ORDER = CURVE.N 
	SEEDH = "BulletproofsDoesNotNeedTrustedSetupH"
	SEEDU = "BulletproofsDoesNotNeedTrustedSetupU"
)

/*
Bulletproofs parameters.
*/
type bp struct {
	n int64
	G *p256
	H *p256
	g []*p256  
	h []*p256  
	zkip bip
}

/*
Bulletproofs proof.
*/
type proofBP struct {
	V *p256
	A *p256
	S *p256
	T1 *p256
	T2 *p256
	taux *big.Int
	mu *big.Int
	tprime *big.Int
	proofip proofBip
}
 
/*
VectorCopy returns a vector composed by copies of a.
*/
func VectorCopy(a *big.Int, n int64) ([]*big.Int, error) {
	var (
		i int64
		result []*big.Int
	)
	result = make([]*big.Int, n)
	i = 0
	for i<n {
		result[i] = a
		i = i + 1
	}
	return result, nil
}

/*
VectorCopy returns a vector composed by copies of a.
*/
func VectorG1Copy(a *p256, n int64) ([]*p256, error) {
	var (
		i int64
		result []*p256
	)
	result = make([]*p256, n)
	i = 0
	for i<n {
		result[i] = a
		i = i + 1
	}
	return result, nil
}

/*
VectorConvertToBig converts an array of int64 to an array of big.Int.
*/
func VectorConvertToBig(a []int64, n int64) ([]*big.Int, error) {
	var (
		i int64
		result []*big.Int
	)
	result = make([]*big.Int, n)
	i = 0
	for i<n {
		result[i] = new(big.Int).SetInt64(a[i])
		i = i + 1
	}
	return result, nil
}

/*
VectorAdd computes vector addition componentwisely.
*/
func VectorAdd(a, b []*big.Int) ([]*big.Int, error) {
	var (
		result []*big.Int
		i,n,m int64
	)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return nil, errors.New("Size of first argument is different from size of second argument.")
	}
	i = 0
	result = make([]*big.Int, n)
	for i<n {
		result[i] = Add(a[i], b[i])	
		result[i] = Mod(result[i], ORDER) 
		i = i + 1
	}
	return result, nil
}

/*
VectorSub computes vector addition componentwisely.
*/
func VectorSub(a, b []*big.Int) ([]*big.Int, error) {
	var (
		result []*big.Int
		i,n,m int64
	)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return nil, errors.New("Size of first argument is different from size of second argument.")
	}
	i = 0
	result = make([]*big.Int, n)
	for i<n {
		result[i] = Sub(a[i], b[i])	
		result[i] = Mod(result[i], ORDER) 
		i = i + 1
	}
	return result, nil
}

/*
VectorScalarMul computes vector scalar multiplication componentwisely.
*/
func VectorScalarMul(a []*big.Int, b *big.Int) ([]*big.Int, error) {
	var (
		result []*big.Int
		i,n int64
	)
	n = int64(len(a))
	i = 0
	result = make([]*big.Int, n)
	for i<n {
		result[i] = Multiply(a[i], b)	
		result[i] = Mod(result[i], ORDER) 
		i = i + 1
	}
	return result, nil
}

/*
VectorMul computes vector multiplication componentwisely.
*/
func VectorMul(a, b []*big.Int) ([]*big.Int, error) {
	var (
		result []*big.Int
		i,n,m int64
	)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return nil, errors.New("Size of first argument is different from size of second argument.")
	}
	i = 0
	result = make([]*big.Int, n)
	for i<n {
		result[i] = Multiply(a[i], b[i])	
		result[i] = Mod(result[i], ORDER) 
		i = i + 1
	}
	return result, nil
}

/*
VectorECMul computes vector EC addition componentwisely.
*/
func VectorECAdd(a,b []*p256) ([]*p256, error) {
	var (
		result []*p256
		i,n,m int64
	)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return nil, errors.New("Size of first argument is different from size of second argument.")
	}
	result = make([]*p256, n)
	i = 0
	for i<n {
		result[i] = new(p256).Multiply(a[i], b[i])	
		i = i + 1
	}
	return result, nil
}
/*
ScalarProduct return the inner product between a and b.
*/
func ScalarProduct(a, b []*big.Int) (*big.Int, error) {
	var (
		result *big.Int
		i,n,m int64
	)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return nil, errors.New("Size of first argument is different from size of second argument.")
	}
	i = 0
	result = GetBigInt("0")
	for i<n {
		ab := Multiply(a[i], b[i])	
		result.Add(result, ab)	
		result = Mod(result, ORDER) 
		i = i + 1
	}
	return result, nil
}

/*
VectorExp computes Prod_i^n{a[i]^b[i]}.
*/
func VectorExp(a []*p256, b []*big.Int) (*p256, error) {
	var (
		result *p256
		i,n,m int64
	)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return nil, errors.New("Size of first argument is different from size of second argument.")
	}
	i = 0
	result = new(p256).SetInfinity()
	for i<n {
		result.Multiply(result, new(p256).ScalarMult(a[i], b[i]))	
		i = i + 1
	}
	return result, nil
}

/*
VectorScalarExp computes a[i]^b for each i.
*/
func VectorScalarExp(a []*p256, b *big.Int) ([]*p256, error) {
	var (
		result []*p256
		i,n int64
	)
	n = int64(len(a))
	result = make([]*p256, n)
	i = 0
	for i<n {
		result[i] = new(p256).ScalarMult(a[i], b)	
		i = i + 1
	}
	return result, nil
}

/*
PowerOf returns a vector composed by powers of x.
*/
func PowerOf(x *big.Int, n int64) ([]*big.Int, error) {
	var (
		i int64
		result []*big.Int
	)
	result = make([]*big.Int, n)
	current := GetBigInt("1")
	i = 0
	for i<n {
		result[i] = current 
		current = Multiply(current, x)
		current = Mod(current, ORDER) 
		i = i + 1
	}
	return result, nil
}

/*
aR = aL - 1^n
*/
func ComputeAR(x []int64) ([]int64, error) {
	var (
		i int64
		result []int64
	)
	result = make([]int64, len(x))
	i = 0
	for i<int64(len(x)) {
		if x[i] == 0 {
			result[i] = -1 
		} else if x[i] == 1 {
			result[i] = 0 
		} else {
			return nil, errors.New("input contains non-binary element") 
		}
		i = i + 1
	}
	return result, nil
}

/*
Hash is responsible for the computing a Zp element given elements from GT and G1.
*/
func HashBP(A, S *p256) (*big.Int, *big.Int, error) {
	digest1 := sha256.New()
	digest1.Write([]byte(A.String()))
	digest1.Write([]byte(S.String()))
	output1 := digest1.Sum(nil)
	tmp1 := output1[0: len(output1)]
	result1, err1 := byteconversion.FromByteArray(tmp1)
	
	digest2 := sha256.New()
	digest2.Write([]byte(S.String()))
	digest2.Write([]byte(A.String()))
	output2 := digest2.Sum(nil)
	tmp2 := output2[0: len(output2)]
	result2, err2 := byteconversion.FromByteArray(tmp2)
	
	if err1 != nil {
		return nil, nil, err1
	} else if err2 != nil {
		return nil, nil, err2
	}
	return result1, result2, nil
}

/*
Commitvector computes a commitment to the bit of the secret. 
TODO: Maybe the common interface could have Commit method, but must take care of the different 
secret types though...
*/
func CommitVector(aL,aR []int64, alpha *big.Int, G,H *p256, g,h []*p256, n int64) (*p256, error) {
	var (
		i int64
		R *p256
	)
	// Compute h^alpha.vg^aL.vh^aR
	R = new(p256).ScalarMult(H, alpha)
	i = 0
	for i<n {
		gaL := new(p256).ScalarMult(g[i], new(big.Int).SetInt64(aL[i]))
		haR := new(p256).ScalarMult(h[i], new(big.Int).SetInt64(aR[i]))
		R.Multiply(R, gaL)
		R.Multiply(R, haR)
		i = i + 1
	}
	return R, nil
}

func CommitVectorBig(aL,aR []*big.Int, alpha *big.Int, G,H *p256, g,h []*p256, n int64) (*p256, error) {
	var (
		i int64
		R *p256
	)
	// Compute h^alpha.vg^aL.vh^aR
	R = new(p256).ScalarMult(H, alpha)
	i = 0
	for i<n {
		R.Multiply(R, new(p256).ScalarMult(g[i], aL[i]))
		R.Multiply(R, new(p256).ScalarMult(h[i], aR[i]))
		i = i + 1
	}
	return R, nil
}

/*
delta(y,z) = (z-z^2) . < 1^n, y^n > - z^3 . < 1^n, 2^n >
*/
func (zkrp *bp) Delta(y, z *big.Int) (*big.Int, error) {
	var (
		result *big.Int
	)
	// delta(y,z) = (z-z^2) . < 1^n, y^n > - z^3 . < 1^n, 2^n >
	z2 := Multiply(z, z)
	z2 = Mod(z2, ORDER) 
	z3 := Multiply(z2, z)
	z3 = Mod(z3, ORDER)

	// < 1^n, y^n >
	v1, _ := VectorCopy(new(big.Int).SetInt64(1), zkrp.n)
	vy, _ := PowerOf(y, zkrp.n) 
	sp1y, _ := ScalarProduct(v1, vy)

	// < 1^n, 2^n >
	p2n, _ := PowerOf(new(big.Int).SetInt64(2), zkrp.n)
	sp12, _ := ScalarProduct(v1, p2n)

	result = Sub(z, z2)
	result = Multiply(result, sp1y)
	result = Sub(result, Multiply(z3, sp12))
	result = Mod(result, ORDER)

	return result, nil
}

/* 
Setup is responsible for computing the common parameter. 
This is STILL a trusted setup.
*/
func (zkrp *bp) Setup(a,b int64) {
	var (
		i int64
	)
	zkrp.G = new(p256).ScalarBaseMult(new(big.Int).SetInt64(1))
	// TODO: change to avoid trusted setup
	zkrp.H, _ = MapToGroup(SEEDH)
	zkrp.n = int64(math.Log2(float64(b)))
	zkrp.g = make([]*p256, zkrp.n)
	zkrp.h = make([]*p256, zkrp.n)
	i = 0
	for i<zkrp.n {
		eg, _ := rand.Int(rand.Reader, ORDER)
		eh, _ := rand.Int(rand.Reader, ORDER)
		zkrp.g[i] = new(p256).ScalarBaseMult(eg)
		zkrp.h[i] = new(p256).ScalarMult(zkrp.H, eh)
		i = i + 1
	}
	// Setup Inner Product
	zkrp.zkip.Setup(zkrp.H, zkrp.g, zkrp.h, new(big.Int).SetInt64(0))
}

/* 
Prove computes the ZK proof. 
*/
func (zkrp *bp) Prove(secret *big.Int) (proofBP, error) {
	var (
		i int64
		sL []*big.Int
		sR []*big.Int
		proof proofBP
	)
	//////////////////////////////////////////////////////////////////////////////
	// First phase
	//////////////////////////////////////////////////////////////////////////////
	
	// commitment to v and gamma
	gamma, _ := rand.Int(rand.Reader, ORDER)
	V, _ := CommitG1(secret, gamma, zkrp.H) 

	// aL, aR and commitment: (A, alpha)
	aL, _ := Decompose(secret, 2, zkrp.n)	
	aR, _ := ComputeAR(aL)
	alpha, _ := rand.Int(rand.Reader, ORDER)
	A, _ := CommitVector(aL, aR, alpha, zkrp.G, zkrp.H, zkrp.g, zkrp.h, zkrp.n) 

	// sL, sR and commitment: (S, rho)
	rho, _ := rand.Int(rand.Reader, ORDER)
	sL = make([]*big.Int, zkrp.n)
	sR = make([]*big.Int, zkrp.n)
	i = 0
	for i<zkrp.n {
		sL[i], _ = rand.Int(rand.Reader, ORDER)
		sR[i], _ = rand.Int(rand.Reader, ORDER)
		i = i + 1
	}
	S, _ := CommitVectorBig(sL, sR, rho, zkrp.G, zkrp.H, zkrp.g, zkrp.h, zkrp.n) 

	// Fiat-Shamir heuristic to compute challenges y, z
	y, z, _ := HashBP(A, S)

	//////////////////////////////////////////////////////////////////////////////
	// Second phase
	//////////////////////////////////////////////////////////////////////////////
	tau1, _ := rand.Int(rand.Reader, ORDER) // page 20 from eprint version
	tau2, _ := rand.Int(rand.Reader, ORDER)
	
	// compute t1: < aL - z.1^n, y^n . sR > + < sL, y^n . (aR + z . 1^n) > 
	vz, _ := VectorCopy(z, zkrp.n)
	vy, _ := PowerOf(y, zkrp.n) 

	// aL - z.1^n
	naL, _ := VectorConvertToBig(aL, zkrp.n)
	aLmvz, _ := VectorSub(naL, vz)
	
	// y^n .sR
	ynsR, _ := VectorMul(vy, sR) 	

	// scalar prod: < aL - z.1^n, y^n . sR >
	sp1, _ := ScalarProduct(aLmvz, ynsR)

	// scalar prod: < sL, y^n . (aR + z . 1^n) >
	naR, _ := VectorConvertToBig(aR, zkrp.n)
	aRzn, _ := VectorAdd(naR, vz)
	ynaRzn, _ := VectorMul(vy, aRzn) 

	// Add z^2.2^n to the result
	// z^2 . 2^n
	p2n, _ := PowerOf(new(big.Int).SetInt64(2), zkrp.n)
	zsquared := Multiply(z, z)
	z22n, _ := VectorScalarMul(p2n, zsquared)
	ynaRzn, _ = VectorAdd(ynaRzn, z22n)
	sp2, _ := ScalarProduct(sL, ynaRzn)
	
	// sp1 + sp2
	t1 := Add(sp1, sp2)
	t1 = Mod(t1, ORDER)
	

	// compute t2: < sL, y^n . sR >
	t2, _ := ScalarProduct(sL, ynsR)
	t2 = Mod(t2, ORDER)

	// compute T1
	T1, _ := CommitG1(t1, tau1, zkrp.H)

	// compute T2
	T2, _ := CommitG1(t2, tau2, zkrp.H)

	// Fiat-Shamir heuristic to compute 'random' challenge x
	x, _, _ := HashBP(T1, T2)

	//////////////////////////////////////////////////////////////////////////////
	// Third phase                                                              //
	//////////////////////////////////////////////////////////////////////////////

	// compute bl
	sLx, _ := VectorScalarMul(sL, x)
	bl, _ := VectorAdd(aLmvz, sLx)

	// compute br
	// y^n . ( aR + z.1^n + sR.x )
	sRx, _ := VectorScalarMul(sR, x)
	aRzn, _ = VectorAdd(aRzn, sRx)
	ynaRzn, _ = VectorMul(vy, aRzn) 
	// y^n . ( aR + z.1^n sR.x ) + z^2 . 2^n
	br, _ := VectorAdd(ynaRzn, z22n)

	// Compute t` = < bl, br >
	tprime, _ := ScalarProduct(bl, br)

	// Compute taux = tau2 . x^2 + tau1 . x + z^2 . gamma
	taux := Multiply(tau2, Multiply(x, x))
	taux = Add(taux, Multiply(tau1, x)) 
	taux = Add(taux, Multiply(Multiply(z, z), gamma))
	taux = Mod(taux, ORDER) 

	// Compute mu = alpha + rho.x
	mu := Multiply(rho, x)
	mu = Add(mu, alpha)
	mu = Mod(mu, ORDER) 

	// Inner Product over (g, h', P.h^-mu, tprime)
	// Compute h'
	hprime := make([]*p256, zkrp.n)
	// Switch generators
	yinv := ModInverse(y, ORDER)
	expy := yinv
	hprime[0] = zkrp.h[0]	
	i = 1
	for i<zkrp.n {
		hprime[i] = new(p256).ScalarMult(zkrp.h[i], expy)	
		expy = Multiply(expy, yinv)
		i = i + 1
	}

	// Update Inner Product Proof Setup
	zkrp.zkip.h = hprime
	zkrp.zkip.c = tprime

	commit, _ := CommitInnerProduct(zkrp.g, hprime, bl, br)
	proofip, _ := zkrp.zkip.Prove(bl, br, commit)	

	// Remove unnecessary variables
	proof.V = V
	proof.A = A
	proof.S = S
	proof.T1 = T1
	proof.T2 = T2
	proof.taux = taux
 	proof.mu = mu
	proof.tprime = tprime
	proof.proofip = proofip

	return proof, nil
}

/* 
Verify returns true if and only if the proof is valid.
*/
func (zkrp *bp) Verify (proof proofBP) (bool, error) {
	var (
		i int64
		hprime []*p256
	)
	hprime = make([]*p256, zkrp.n)
	y, z, _ := HashBP(proof.A, proof.S)
	x, _, _ := HashBP(proof.T1, proof.T2)

	// Switch generators
	yinv := ModInverse(y, ORDER)
	expy := yinv
	hprime[0] = zkrp.h[0]	
	i = 1
	for i<zkrp.n {
		hprime[i] = new(p256).ScalarMult(zkrp.h[i], expy)	
		expy = Multiply(expy, yinv)
		i = i + 1
	}

	//////////////////////////////////////////////////////////////////////////////
	// Check that tprime  = t(x) = t0 + t1x + t2x^2  ----------  Condition (65) //
	//////////////////////////////////////////////////////////////////////////////
	
	// Compute left hand side
	lhs, _ := CommitG1(proof.tprime, proof.taux, zkrp.H)
	
	// Compute right hand side
	z2 := Multiply(z, z)
	z2 = Mod(z2, ORDER) 
	x2 := Multiply(x, x)
	x2 = Mod(x2, ORDER) 

	rhs := new(p256).ScalarMult(proof.V, z2)

	delta, _ := zkrp.Delta(y,z)

	gdelta := new(p256).ScalarBaseMult(delta)

	rhs.Multiply(rhs, gdelta)

	T1x := new(p256).ScalarMult(proof.T1, x) 
	T2x2 := new(p256).ScalarMult(proof.T2, x2) 

	rhs.Multiply(rhs, T1x)
	rhs.Multiply(rhs, T2x2)

	// Subtract lhs and rhs and compare with poitn at infinity
	lhs.Neg(lhs)
	rhs.Multiply(rhs, lhs)
	c65 := rhs.IsZero() // Condition (65), page 20, from eprint version

	// Verify Inner Product Proof ################################################
	ok, _ := zkrp.zkip.Verify(proof.proofip)

	result := c65 && ok

	return result, nil
}

//////////////////////////////////// Inner Product ////////////////////////////////////

/*
Base struct for the Inner Product Argument.
*/
type bip struct {
	n int64
	c *big.Int
	u *p256
	H *p256
	g []*p256  
	h []*p256  
	P *p256
}

/*
Struct that contains the Inner Product Proof.
*/
type proofBip struct {
	Ls []*p256
	Rs []*p256
	u *p256
	P *p256
	g *p256
	h *p256
	a *big.Int
	b *big.Int
	n int64
}

/*
HashIP is responsible for the computing a Zp element given elements from GT and G1.
*/
func HashIP(g,h []*p256, P *p256, c *big.Int, n int64) (*big.Int, error) {
	var (
		i int64
	)

	digest := sha256.New()
	digest.Write([]byte(P.String()))
	
	i = 0
	for i<n {
		digest.Write([]byte(g[i].String()))
		digest.Write([]byte(h[i].String()))
		i = i + 1
	}
	
	digest.Write([]byte(c.String()))
	output := digest.Sum(nil)
	tmp := output[0: len(output)]
	result, err := byteconversion.FromByteArray(tmp)
	
	return result, err
}

/*
CommitinnerProduct is responsible for calculating g^a.h^b.
*/
func CommitInnerProduct(g,h []*p256, a,b []*big.Int) (*p256, error) {
	var (
		result *p256
	)

	ga, _ := VectorExp(g, a)
	hb, _ := VectorExp(h, b)
	result = new(p256).Multiply(ga, hb)
	return result, nil
}

/*
Setup is responsible for computing the inner product basic parameters that are common to both
Prove and Verify algorithms.
*/
func (zkip *bip) Setup(H *p256, g,h []*p256, c *big.Int) (bip, error) {
	var (
		params bip
	)
	
	zkip.g = make([]*p256, zkip.n)
	zkip.h = make([]*p256, zkip.n)
	// TODO: not yet avoiding trusted setup...
	zkip.u, _ = MapToGroup(SEEDU)
	zkip.H = H
	zkip.g = g
	zkip.h = h
	zkip.c = c

	return params, nil
}


/*
Prove is responsible for the generation of the Inner Product Proof.
*/
func (zkip *bip) Prove(a,b []*big.Int, P *p256) (proofBip, error) {
	var (
		proof proofBip
		n,m int64
		Ls []*p256 
		Rs []*p256 
	)

	// Fiat-Shamir:
	// x = Hash(g,h,P,c)
	x, _ := HashIP(zkip.g, zkip.h, P, zkip.c, zkip.n)
	// Pprime = P.u^(x.c)		
	ux := new(p256).ScalarMult(zkip.u, x)  
	uxc := new(p256).ScalarMult(ux, zkip.c)  
	P.Multiply(P, uxc)
	n = int64(len(a))
	m = int64(len(b))
	if (n != m) {
		return proof, errors.New("Size of first array argument must be equal to the second")
	} else {
		// Execute Protocol 2 recursively
		zkip.P = P
		proof, err := BIP(a, b, zkip.g, zkip.h, ux, zkip.P, n, Ls, Rs)
		return proof, err
	}
		
	return proof, nil
}

/*
BIP is the main recursive function that will be used to compute the inner product argument.
*/
func BIP(a,b []*big.Int, g,h []*p256, u,P *p256, n int64, Ls,Rs []*p256) (proofBip, error) {
	var (
		proof proofBip
		cL, cR, x, xinv, x2, x2inv *big.Int
		L, R, Lh, Rh, Pprime *p256
		gprime, hprime, gprime2, hprime2 []*p256
		aprime, bprime, aprime2, bprime2 []*big.Int
	)

	if (n == 1) {
		// recursion end
		proof.a = a[0]
		proof.b = b[0]
		proof.g = g[0]
		proof.h = h[0]
		proof.P = P
		proof.u = u
		proof.Ls = Ls
		proof.Rs = Rs

	} else {
		// recursion

		// nprime := n / 2
		nprime := n / 2

		// Compute cL = < a[:n'], b[n':] >
		cL, _ = ScalarProduct(a[:nprime], b[nprime:])
		// Compute cR = < a[n':], b[:n'] >
		cR, _ = ScalarProduct(a[nprime:], b[:nprime])
		// Compute L = g[n':]^(a[:n']).h[:n']^(b[n':]).u^cL
		L, _ = VectorExp(g[nprime:],a[:nprime])
		Lh, _ = VectorExp(h[:nprime], b[nprime:])
		L.Multiply(L, Lh)
		L.Multiply(L, new(p256).ScalarMult(u, cL))
		
		// Compute R = g[:n']^(a[n':]).h[n':]^(b[:n']).u^cR
		R, _ = VectorExp(g[:nprime],a[nprime:]) 
		Rh, _ = VectorExp(h[nprime:], b[:nprime])
		R.Multiply(R, Rh)
		R.Multiply(R, new(p256).ScalarMult(u, cR))

		// Fiat-Shamir:
		x, _, _ = HashBP(L, R)
		xinv = ModInverse(x, ORDER)

		// Compute g' = g[:n']^(x^-1) * g[n':]^(x)
		gprime, _ = VectorScalarExp(g[:nprime], xinv)
		gprime2, _ = VectorScalarExp(g[nprime:], x)
		gprime, _ = VectorECAdd(gprime, gprime2)
		// Compute h' = h[:n']^(x)    * h[n':]^(x^-1)
		hprime, _ = VectorScalarExp(h[:nprime], x)
		hprime2, _ = VectorScalarExp(h[nprime:], xinv)
		hprime, _ = VectorECAdd(hprime, hprime2)

		// Compute P' = L^(x^2).P.R^(x^-2)
		x2 = Mod(Multiply(x,x), ORDER)
		x2inv = ModInverse(x2, ORDER)
		Pprime = new(p256).ScalarMult(L, x2)
		Pprime.Multiply(Pprime, P)
		Pprime.Multiply(Pprime, new(p256).ScalarMult(R, x2inv))

		// Compute a' = a[:n'].x      + a[n':].x^(-1)
		aprime, _ = VectorScalarMul(a[:nprime], x)
		aprime2, _ = VectorScalarMul(a[nprime:], xinv)
		aprime, _ = VectorAdd(aprime, aprime2)
		// Compute b' = b[:n'].x^(-1) + b[n':].x
		bprime, _ = VectorScalarMul(b[:nprime], xinv)
		bprime2, _ = VectorScalarMul(b[nprime:], x)
		bprime, _ = VectorAdd(bprime, bprime2)

		Ls = append(Ls, L)
		Rs = append(Rs, R)
		// recursion BIP(g',h',u,P'; a', b')
		proof, _ = BIP(aprime, bprime, gprime, hprime, u, Pprime, nprime, Ls, Rs)
	}
	proof.n = n
	return proof, nil
}

/* 
Verify is responsible for the verification of the Inner Product Proof. 
*/
func (zkip *bip) Verify(proof proofBip) (bool, error) {
	
	logn := len(proof.Ls)
	var (
		i int64
		x, xinv, x2, x2inv *big.Int
		ngprime, nhprime, ngprime2, nhprime2 []*p256
	)

	i = 0
	gprime := zkip.g
	hprime := zkip.h
	Pprime := zkip.P
	nprime := proof.n 
	for i < int64(logn) {
		nprime = nprime / 2
		x, _, _ = HashBP(proof.Ls[i], proof.Rs[i])
		xinv = ModInverse(x, ORDER)
		// Compute g' = g[:n']^(x^-1) * g[n':]^(x)
		ngprime, _ = VectorScalarExp(gprime[:nprime], xinv)
		ngprime2, _ = VectorScalarExp(gprime[nprime:], x)
		gprime, _ = VectorECAdd(ngprime, ngprime2)
		// Compute h' = h[:n']^(x)    * h[n':]^(x^-1)
		nhprime, _ = VectorScalarExp(hprime[:nprime], x)
		nhprime2, _ = VectorScalarExp(hprime[nprime:], xinv)
		hprime, _ = VectorECAdd(nhprime, nhprime2)
		// Compute P' = L^(x^2).P.R^(x^-2)
		x2 = Mod(Multiply(x,x), ORDER)
		x2inv = ModInverse(x2, ORDER)
		Pprime.Multiply(Pprime, new(p256).ScalarMult(proof.Ls[i], x2))
		Pprime.Multiply(Pprime, new(p256).ScalarMult(proof.Rs[i], x2inv))

		i = i + 1
	}

	// c == a*b
	ab := Multiply(proof.a, proof.b)
	ab = Mod(ab, ORDER)

	rhs := new(p256).ScalarMult(gprime[0], proof.a)
	hb := new(p256).ScalarMult(hprime[0], proof.b)
	rhs.Multiply(rhs, hb)
	rhs.Multiply(rhs, new(p256).ScalarMult(proof.u, ab))

	nP := Pprime.Neg(Pprime)
	nP.Multiply(nP, rhs)
	c := nP.IsZero() 

	return c, nil
}