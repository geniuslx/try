// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"math"

	"github.com/cpmech/goga"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
	"github.com/cpmech/gosl/utl"
)

// case A: finding the minimum of 2.0 + (1+x)²
func fcnA(f, g, h, x []float64, y []int, cpu int) {
	if x[0] == 0 {
		f[0] = 0.0
		return
	}
	x4 := math.Pow(x[0], 4.0)
	f[0] = x4*math.Cos(x[0]) + 2.0*x4
}

// main function
func main() {

	// problem definition
	nf := 1 // number of objective functions
	ng := 0 // number of inequality constraints
	nh := 0 // number of equality constraints

	// the solver (optimiser)
	var opt goga.Optimiser
	opt.Default()              // must call this to set default constants
	opt.FltMin = []float64{-2} // must set minimum
	opt.FltMax = []float64{2}  // must set maximum

	// initialise the solver
	opt.Init(goga.GenTrialSolutions, nil, fcnA, nf, ng, nh)

	// solve problem
	opt.Solve()

	// auxiliary
	fvec := []float64{0} // temporary vector to use with fcnA
	xvec := []float64{0} // temporary vector to use with fcnA

	// print results
	xvec[0] = opt.Solutions[0].Flt[0]  // best x value
	fcnA(fvec, nil, nil, xvec, nil, 0) // fvec[0] is the best f value
	io.Pf("xBest    = %v\n", xvec[0])
	io.Pf("f(xBest) = %v\n", fvec[0])

	// plotting
	npts := 101
	X := utl.LinSpace(-2, 2, npts)
	F := make([]float64, npts)
	for i := 0; i < npts; i++ {
		xvec[0] = X[i]
		fcnA(fvec, nil, nil, xvec, nil, 0)
		F[i] = fvec[0]
	}
	plt.Reset(true, nil)
	plt.PlotOne(xvec[0], fvec[0], &plt.A{C: "r", M: "o", Ms: 20, NoClip: true})
	plt.Plot(X, F, nil)
	plt.Gll("$x$", "$f$", nil)
	plt.Save("/tmp", "tutorial2")
}
