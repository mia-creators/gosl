// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"math"
	"testing"

	"github.com/cpmech/gosl/chk"
)

func TestJacobi01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Jacobi01")

	A := NewMatrixSlice([][]float64{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2},
	})

	Q := NewMatrix(3, 3)
	v := NewVector(3)
	err := Jacobi(Q, v, A)
	if err != nil {
		chk.Panic("Jacobi failed:\n%v", err)
	}

	chk.Matrix(tst, "Q", 1e-17, Q.GetSlice(), [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	chk.Vector(tst, "v", 1e-17, v, []float64{2, 2, 2})
}

func TestJacobi02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Jacobi02")

	A := NewMatrixSlice([][]float64{
		{2, 0, 0},
		{0, 3, 4},
		{0, 4, 9},
	})

	Q := NewMatrix(3, 3)
	v := NewVector(3)
	err := Jacobi(Q, v, A)
	if err != nil {
		chk.Panic("Jacobi failed:\n%v", err)
	}

	os3 := 1.0 / math.Sqrt(5.0)
	chk.Matrix(tst, "Q", 1e-17, Q.GetSlice(), [][]float64{
		{1, +0 * os3, 0 * os3},
		{0, +2 * os3, 1 * os3},
		{0, -1 * os3, 2 * os3},
	})
	chk.Vector(tst, "v", 1e-17, v, []float64{2, 1, 11})
}

func TestJacobi03(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Jacobi03")

	A := NewMatrixSlice([][]float64{
		{1, 2, 3},
		{2, 3, 2},
		{3, 2, 2},
	})

	Q := NewMatrix(3, 3)
	v := NewVector(3)
	err := Jacobi(Q, v, A)
	if err != nil {
		chk.Panic("Jacobi failed:\n%v", err)
	}

	chk.Matrix(tst, "Q", 1e-17, Q.GetSlice(), [][]float64{
		{+7.81993314738381295e-01, 5.26633230856907386e-01, +3.33382506832158143e-01},
		{-7.14394870018381645e-02, 6.07084171793832561e-01, -7.91419742017035133e-01},
		{-6.19179178753124115e-01, 5.95068272145819699e-01, +5.12358171676802088e-01},
	})
	chk.Vector(tst, "v", 1e-17, v, []float64{-1.55809924785903786e+00, 6.69537390404459476e+00, 8.62725343814443657e-01})
}

func TestJacobi04(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Jacobi04")

	A := NewMatrixSlice([][]float64{
		{1, 2, 3, 4, 5},
		{2, 3, 0, 2, 4},
		{3, 0, 2, 1, 3},
		{4, 2, 1, 1, 2},
		{5, 4, 3, 2, 1},
	})

	Q := NewMatrix(5, 5)
	v := NewVector(5)
	err := Jacobi(Q, v, A)
	if err != nil {
		chk.Panic("Jacobi failed:\n%v", err)
	}

	chk.Matrix(tst, "Q", 1e-14, Q.GetSlice(), [][]float64{
		{+4.265261184874604e-01, +5.285232769688938e-01, +1.854383137677959e-01, +2.570216184506737e-01, -6.620355997875309e-01},
		{-3.636641874245161e-01, +4.182907021187977e-01, -7.200691218899387e-01, -3.444995789572199e-01, -2.358002271092630e-01},
		{-5.222548807800880e-01, +3.413546312786976e-01, +6.672573809673910e-01, -4.053509412317634e-01, -3.442465966457679e-02},
		{-4.133525029362699e-01, +3.807798553184266e-01, -3.950209555261502e-02, +7.608554466087614e-01, +3.220015278111787e-01},
		{+4.921517823299884e-01, +5.330851261396132e-01, -1.789590676939640e-02, -2.684204380363021e-01, +6.334327718104180e-01},
	})
	chk.Vector(tst, "v", 1e-13, v, []float64{-2.485704750172629e+00, +1.244545682971212e+01, +2.694072690168129e+00, +2.073336609414627e-01, -4.861158430649138e+00})
}