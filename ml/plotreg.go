// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ml

import (
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/plt"
	"github.com/cpmech/gosl/utl"
)

// PlotterReg defines a plotter to plot results from regression models
type PlotterReg struct {

	// input
	data   *Data      // data
	params *ParamsReg // parameters
	model  Regression // model

	// constants
	YlineNpts int // number of points for ModelY()
	MgridNpts int // nubmer of poitns for meshgrid (for contours)

	// arguments
	ArgsYdata    *plt.A         // arguments for x-y data plots
	ArgsYmodel   *plt.A         // arguments for x-y model line
	ArgsYclasses map[int]*plt.A // maps y classes [0, 1, 2, ...] to plot arguments
	ArgsCcost    *plt.A         // arguments for ContourCost
	ArgsCcostMdl *plt.A         // arguments for the model parameters in ContourCost
}

// NewPlotterReg returns a new ploter
func NewPlotterReg(data *Data, params *ParamsReg, reg Regression) (o *PlotterReg) {

	// input
	o = new(PlotterReg)
	o.data = data
	o.params = params
	o.model = reg

	// constants
	o.YlineNpts = 21
	o.MgridNpts = 21

	// arguments
	o.ArgsYdata = &plt.A{C: plt.C(2, 0), M: plt.M(0, 0), Ls: "None", NoClip: true}
	o.ArgsYmodel = &plt.A{C: plt.C(0, 0), M: "None", Ls: "-", NoClip: true}
	o.ArgsYclasses = make(map[int]*plt.A)
	nMaxClassesIni := 10
	for k := 0; k < nMaxClassesIni; k++ {
		o.ArgsYclasses[k] = &plt.A{C: plt.C(k, 0), M: plt.M(k, 0), NoClip: true}
	}
	o.ArgsCcost = &plt.A{}
	o.ArgsCcostMdl = &plt.A{C: "yellow", M: "o"}
	return
}

// DataY plots data x[iFeature] versus data y values
func (o *PlotterReg) DataY(iFeature int) {
	u := o.data.X.GetCol(iFeature)
	v := o.data.Y
	plt.Plot(u, v, o.ArgsYdata)
	plt.HideTRborders()
	plt.Gll(io.Sf("$x_{%d}$", iFeature), "$y$", nil)
}

// ModelY plots model y values
func (o *PlotterReg) ModelY(iFeature int, xmin, xmax float64) {
	x := la.NewVector(o.data.Nfeatures)
	u := utl.LinSpace(xmin, xmax, o.YlineNpts)
	v := utl.GetMapped(u, func(xi float64) float64 {
		x[iFeature] = xi
		return o.model.Predict(x)
	})
	plt.Plot(u, v, o.ArgsYmodel)
}

// DataClass plots data classes
func (o *PlotterReg) DataClass(iFeature, jFeature int) {
	for iSample := 0; iSample < o.data.Nsamples; iSample++ {
		args := o.ArgsYclasses[iSample%len(o.ArgsYclasses)]
		ui := o.data.X.Get(iSample, iFeature)
		vi := o.data.X.Get(iSample, jFeature)
		plt.PlotOne(ui, vi, args)
	}
}

// ContourCost plots a contour of Cost for many parameters values
//  iParam, jParam -- selected parameters [use -1 for bias]
func (o *PlotterReg) ContourCost(iParam, jParam int, pimin, pimax, pjmin, pjmax float64) {

	// create meshgrid
	o.params.Backup()
	U, V, W := utl.MeshGrid2dF(pimin, pimax, pjmin, pjmax, o.MgridNpts, o.MgridNpts, func(s, t float64) (w float64) {
		o.params.Restore(true)
		o.params.SetParam(iParam, s)
		o.params.SetParam(jParam, t)
		w = o.model.Cost()
		return
	})

	// plot contour
	plt.ContourF(U, V, W, o.ArgsCcost)

	// plot optimal solution
	o.params.Restore(true)
	plt.PlotOne(o.params.GetParam(iParam), o.params.GetParam(jParam), o.ArgsCcostMdl)

	// set labels
	stri := "$b$"
	strj := "$b$"
	if iParam >= 0 {
		stri = io.Sf("$\\theta_{%d}$", iParam)
	}
	if jParam >= 0 {
		strj = io.Sf("$\\theta_{%d}$", jParam)
	}
	plt.SetXlabel(stri, nil)
	plt.SetYlabel(strj, nil)
}
