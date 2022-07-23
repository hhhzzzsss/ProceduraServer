package region

import "github.com/hhhzzzsss/procedura-generator/util"

type RegionCache[T any] struct {
	xdim, ydim, zdim int
	contents         []T
}

func MakeRegionCache[T any](r *Region) RegionCache[T] {
	return RegionCache[T]{
		r.xdim, r.ydim, r.zdim,
		make([]T, r.xdim*r.ydim*r.zdim),
	}
}

func (r *RegionCache[T]) Set(x, y, z int, val T) {
	if !r.IsInRange(x, y, z) {
		return
	}
	r.contents[y*r.zdim*r.xdim+z*r.xdim+x] = val
}

func (r *RegionCache[T]) Get(x, y, z int) T {
	if !r.IsInRange(x, y, z) {
		return *new(T)
	}
	return r.contents[y*r.zdim*r.xdim+z*r.xdim+x]
}

func (r *RegionCache[T]) ForEach(generator func(x, y, z int) T) {
	var bar util.ProgressBar
	bar.Initialize(r.ydim)
	for y := 0; y < r.ydim; y++ {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				r.Set(x, y, z, generator(x, y, z))
			}
		}
		bar.Play(y + 1)
	}
	bar.Finish()
}

func (r *RegionCache[T]) ForEachNormalized(generator func(x, y, z float64) T) {
	var bar util.ProgressBar
	bar.Initialize(r.ydim)
	minDim := r.xdim
	if r.ydim < minDim {
		minDim = r.ydim
	}
	if r.zdim < minDim {
		minDim = r.zdim
	}
	for y := 0; y < r.ydim; y++ {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				xNorm := 2.0*float64(x)/float64(r.xdim) - 1.0
				yNorm := 2.0*float64(y)/float64(r.ydim) - 1.0
				zNorm := 2.0*float64(z)/float64(r.zdim) - 1.0
				xNorm *= float64(r.xdim) / float64(minDim)
				yNorm *= float64(r.ydim) / float64(minDim)
				zNorm *= float64(r.zdim) / float64(minDim)
				r.Set(x, y, z, generator(xNorm, yNorm, zNorm))
			}
		}
		bar.Play(y + 1)
	}
	bar.Finish()
}

func (r *RegionCache[T]) ForEachParallel(threads int, generator func(x, y, z int) T) {
	var bar util.ProgressBar
	bar.Initialize(r.ydim)
	cLayer := make(chan int, r.ydim)
	cDone := make(chan struct{}, r.ydim)
	for i := 0; i < threads; i++ {
		go r.forEachWorker(cLayer, cDone, generator)
	}
	for y := 0; y < r.ydim; y++ {
		cLayer <- y
	}
	for y := 0; y < r.ydim; y++ {
		<-cDone
		bar.Play(y + 1)
	}
	bar.Finish()
}

func (r *RegionCache[T]) forEachWorker(cLayer chan int, cDone chan struct{}, generator func(x, y, z int) T) {
	for y := range cLayer {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				r.Set(x, y, z, generator(x, y, z))
			}
		}
		cDone <- struct{}{}
	}
}

func (r *RegionCache[T]) ForEachNormalizedParallel(threads int, generator func(x, y, z float64) T) {
	var bar util.ProgressBar
	bar.Initialize(r.ydim)
	cLayer := make(chan int, r.ydim)
	cDone := make(chan struct{}, r.ydim)
	for i := 0; i < threads; i++ {
		go r.forEachNormalizedWorker(cLayer, cDone, generator)
	}
	for y := 0; y < r.ydim; y++ {
		cLayer <- y
	}
	for y := 0; y < r.ydim; y++ {
		<-cDone
		bar.Play(y + 1)
	}
	bar.Finish()
}

func (r *RegionCache[T]) forEachNormalizedWorker(cLayer chan int, cDone chan struct{}, generator func(x, y, z float64) T) {
	minDim := r.xdim
	if r.ydim < minDim {
		minDim = r.ydim
	}
	if r.zdim < minDim {
		minDim = r.zdim
	}
	for y := range cLayer {
		for z := 0; z < r.zdim; z++ {
			for x := 0; x < r.xdim; x++ {
				xNorm := 2.0*float64(x)/float64(r.xdim) - 1.0
				yNorm := 2.0*float64(y)/float64(r.ydim) - 1.0
				zNorm := 2.0*float64(z)/float64(r.zdim) - 1.0
				xNorm *= float64(r.xdim) / float64(minDim)
				yNorm *= float64(r.ydim) / float64(minDim)
				zNorm *= float64(r.zdim) / float64(minDim)
				r.Set(x, y, z, generator(xNorm, yNorm, zNorm))
			}
		}
		cDone <- struct{}{}
	}
}

func (r *RegionCache[T]) IsInRange(x, y, z int) bool {
	return x >= 0 && x < r.xdim && y >= 0 && y < r.ydim && z >= 0 && z < r.zdim
}
