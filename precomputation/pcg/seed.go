package pcg

import (
	"fmt"
	bls12381 "github.com/kilic/bls12-381"
	"github.com/perun-network/bbs-plus-threshold-wallet/precomputation/dspf"
	"math/big"
)

type seedExponents struct {
	aOmega [][]*big.Int // Exponents for a_i
	eEta   [][]*big.Int // Exponents for e_i
	sPhi   [][]*big.Int // Exponents for s_i
}

type seedCoefficients struct {
	aBeta    [][]*bls12381.Fr // Coefficients for a_i
	eGamma   [][]*bls12381.Fr // Coefficients for e_i
	sEpsilon [][]*bls12381.Fr // Coefficients for s_i
}

type DSPFKeyPair struct {
	Key0 dspf.Key
	Key1 dspf.Key
}

// Seed is the seed generated by the Gen function of the PCG.
// It allows to derive ECDSA tuples from the EvalAll function of the PCG.
type Seed struct {
	index        int
	ski          *bls12381.Fr
	exponents    seedExponents
	coefficients seedCoefficients
	U            [][][]*DSPFKeyPair   // U[i][j][r]
	C            [][][][]*DSPFKeyPair // C[i][j][r][s]
	V            [][][][]*DSPFKeyPair // V[i][j][r][s]
}

func (s *Seed) GetSki() *bls12381.Fr {
	return s.ski
}

func (s *Seed) GetIndex() int {
	return s.index
}

func (s *Seed) Serialize() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *Seed) Deserialize(data []byte) error {
	return fmt.Errorf("not implemented")
}

type oleSeed struct {
	index        int
	exponents    seedExponents
	coefficients seedCoefficients
	V            [][]*DSPFKeyPair
}

type voleSeed struct {
	index        int
	constant     *bls12381.Fr
	exponents    seedExponents
	coefficients seedCoefficients
	V            []*DSPFKeyPair
}