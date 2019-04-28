package driver

type API interface {
	Abs(x float64) float64
	Acos(x float64) float64
	Acosh(x float64) float64
	Asin(x float64) float64
	Asinh(x float64) float64
	Atan(x float64) float64
	Atan2(y, x float64) float64
	Atanh(x float64) float64
	Cbrt(x float64) float64
	Ceil(x float64) float64
	Copysign(x, y float64) float64
	Cos(x float64) float64
	Cosh(x float64) float64
	Dim(x, y float64) float64
	Erf(x float64) float64
	Erfc(x float64) float64
	Erfcinv(x float64) float64
	Erfinv(x float64) float64
	Exp(x float64) float64
	Exp2(x float64) float64
	Expm1(x float64) float64
	Float32bits(f float32) uint32
	Float32frombits(b uint32) float32
	Float64bits(f float64) uint64
	Float64frombits(b uint64) float64
	Floor(x float64) float64
	Frexp(f float64) (frac float64, exp int)
	Gamma(x float64) float64
	Hypot(p, q float64) float64
	Ilogb(x float64) int
	Inf(sign int) float64
	IsInf(f float64, sign int) bool
	IsNaN(f float64) (is bool)
	J0(x float64) float64
	J1(x float64) float64
	Jn(n int, x float64) float64
	Ldexp(frac float64, exp int) float64
	Lgamma(x float64) (lgamma float64, sign int)
	Log(x float64) float64
	Log10(x float64) float64
	Log1p(x float64) float64
	Log2(x float64) float64
	Logb(x float64) float64
	Max(x, y float64) float64
	Min(x, y float64) float64
	Mod(x, y float64) float64
	Modf(f float64) (int float64, frac float64)
	NaN() float64
	Nextafter(x, y float64) (r float64)
	Nextafter32(x, y float32) (r float32)
	Pow(x, y float64) float64
	Pow10(n int) float64
	Remainder(x, y float64) float64
	Round(x float64) float64
	RoundToEven(x float64) float64
	Signbit(x float64) bool
	Sin(x float64) float64
	Sincos(x float64) (sin, cos float64)
	Sinh(x float64) float64
	Sqrt(x float64) float64
	Tan(x float64) float64
	Tanh(x float64) float64
	Trunc(x float64) float64
	Y0(x float64) float64
	Y1(x float64) float64
	Yn(n int, x float64) float64
}
