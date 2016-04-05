package mathbeta

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// util func for convert other types to string
func conv(i interface{}) string {
	switch i.(type) {
	case int8:
		return strconv.FormatInt(int64(i.(int8)), 10)
	case byte:
		return strconv.FormatUint(uint64(i.(byte)), 10)
	case int16:
		return strconv.FormatInt(int64(i.(int16)), 10)
	case uint16:
		return strconv.FormatUint(uint64(i.(uint16)), 10)
	case int32:
		return strconv.FormatInt(int64(i.(int32)), 10)
	case uint32:
		return strconv.FormatUint(uint64(i.(uint32)), 10)
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case uint64:
		return strconv.FormatUint(i.(uint64), 10)
	case int:
		return strconv.Itoa(i.(int))
	case float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	default:
		return ""
	}
}

type Matrix struct {
	Row, Column int
	Data        [][]float64
}

// create a new matrix of row * column dimensions, filling element row-oriented with data
func NewMatrix(row, column int, data ...float64) *Matrix {
	d := make([][]float64, row)
	for i := 0; i < row; i++ {
		d[i] = make([]float64, column)
	}
	dl := len(data)
	for i := 0; i < dl; i++ {
		d[i/column][i%column] = data[i]
	}
	return &Matrix{Data: d, Row: row, Column: column}
}

// create a random matrix of row * column dimensions
func RandMatrix(row, column int) *Matrix {
	d := make([][]float64, row)
	for i := 0; i < row; i++ {
		d[i] = make([]float64, column)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			d[i][j] = r.Float64()
		}
	}
	return &Matrix{Data: d, Row: row, Column: column}
}

func Eye(n int) *Matrix {
	d := make([][]float64, n)
	for i := 0; i < n; i++ {
		d[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		d[i][i] = 1
	}
	return &Matrix{Data: d, Row: n, Column: n}
}

func Ones(row, column int) *Matrix {
	d := make([][]float64, row)
	for i := 0; i < row; i++ {
		d[i] = make([]float64, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			d[i][j] = 1
		}
	}
	return &Matrix{Data: d, Row: row, Column: column}
}

func Zeros(row, column int) *Matrix {
	d := make([][]float64, row)
	for i := 0; i < row; i++ {
		d[i] = make([]float64, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			d[i][j] = 0
		}
	}
	return &Matrix{Data: d, Row: row, Column: column}
}

func (m *Matrix) Print() {
	fmt.Print(m)
}

// return a string representation of the matrix
func (m *Matrix) String() string {
	s := ""
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column-1; j++ {
			s += conv(m.Data[i][j]) + " "
		}
		s += conv(m.Data[i][m.Column-1]) + "\n"
	}
	return s
}

func (m *Matrix) Add(a *Matrix) *Matrix {
	if m.Row != a.Row || m.Column != a.Column {
		panic("dimensions not consistent!")
	}
	d := make([][]float64, m.Row)
	for i := 0; i < m.Row; i++ {
		d[i] = make([]float64, m.Column)
	}
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			d[i][j] = a.Data[i][j] + m.Data[i][j]
		}
	}
	return &Matrix{Row: m.Row, Column: m.Column, Data: d}
}

func (m *Matrix) Subtract(a *Matrix) *Matrix {
	if m.Row != a.Row || m.Column != a.Column {
		panic("dimensions not consistent!")
	}
	d := make([][]float64, m.Row)
	for i := 0; i < m.Row; i++ {
		d[i] = make([]float64, m.Column)
	}
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			d[i][j] = m.Data[i][j] - a.Data[i][j]
		}
	}
	return &Matrix{Row: m.Row, Column: m.Column, Data: d}
}

func (m *Matrix) Multiply(a *Matrix) *Matrix {
	if m.Column != a.Row {
		panic("dimensions not consistent!")
	}
	d := make([][]float64, m.Row)
	for i := 0; i < m.Row; i++ {
		d[i] = make([]float64, a.Column)
	}
	for i := 0; i < m.Row; i++ {
		for j := 0; j < a.Column; j++ {
			d[i][j] = 0
			for h := 0; h < m.Column; h++ {
				d[i][j] += m.Data[i][h] * a.Data[h][j]
			}
		}
	}
	return &Matrix{Row: m.Row, Column: a.Column, Data: d}
}

// get the transpose of the current matrix
func (m *Matrix) Transpose() *Matrix {
	d := make([][]float64, m.Column)
	for i := 0; i < m.Column; i++ {
		d[i] = make([]float64, m.Row)
	}
	for i := 0; i < m.Column; i++ {
		for j := 0; j < m.Row; j++ {
			d[i][j] = m.Data[j][i]
		}
	}

	return &Matrix{Row: m.Column, Column: m.Row, Data: d}
}

// copy the matrix to a new one
func (m *Matrix) Copy() *Matrix {
	d := make([][]float64, m.Row)
	for i := 0; i < m.Row; i++ {
		d[i] = make([]float64, m.Column)
	}
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			d[i][j] = m.Data[i][j]
		}
	}

	return &Matrix{Row: m.Row, Column: m.Column, Data: d}
}

// get the rank of the matrix, using row transformation method
func (m *Matrix) Rank() int {
	if m.Row <= m.Column {
		// copy the original matrix
		mm := m.Copy()
		// mark the current processing column
		k := 0
		for i := 0; i < mm.Row && k < mm.Column; {
			pivot := mm.Data[i]
			// the head of the current row is 0,
			// we should switch the row with other row whose head is non-zero
			if pivot[k] == 0 {
				j := i + 1
				for ; j < mm.Row; j++ {
					if mm.Data[j][k] != 0 {
						for h := i; h < mm.Column; h++ {
							t := pivot[h]
							pivot[h] = mm.Data[j][h]
							mm.Data[j][h] = t
						}
						break
					}
				}
				// the current column and rows after that are all with zero head
				// go on to the next column
				if j >= mm.Row {
					k++
					continue
				}
			}
			p := pivot[k]
			// set the head of the current row to 1(stardardize the row)
			for j := i; j < mm.Column; j++ {
				pivot[j] = pivot[j] / p
			}
			// eliminate the rows after the current one
			for j := i + 1; j < mm.Row; j++ {
				f := -mm.Data[j][k]
				mm.Data[j][k] = 0
				for h := k + 1; h < mm.Column; h++ {
					mm.Data[j][h] = mm.Data[j][h] + f*mm.Data[i][h]
				}
			}
			i++
		}
		r := 0
		// count all the non-zero rows as the rank
		for i := 0; i < mm.Row; i++ {
			zeros := true
			for j := 0; j < mm.Column; j++ {
				if mm.Data[i][j] != 0 {
					zeros = false
					break
				}
			}
			if !zeros {
				r++
			}
		}
		return r
	}

	return m.Transpose().Rank()
}

// get the determinant of the matrix, using row transformation method
func (m *Matrix) Determinant() float64 {
	if m.Column != m.Row {
		panic("matrix with different row count and column count has no determinant")
	}
	if m.Row == 1 {
		return m.Data[0][0]
	}
	mm := m.Copy()
	for i := 0; i < mm.Row; {
		pivot := mm.Data[i]
		// the head of the current row is 0,
		// we should switch the row with other row whose head is non-zero
		if pivot[i] == 0 {
			j := i + 1
			for ; j < mm.Row; j++ {
				if mm.Data[j][i] != 0 {
					for h := i; h < mm.Column; h++ {
						t := pivot[h]
						pivot[h] = mm.Data[j][h]
						mm.Data[j][h] = t
					}
					break
				}
			}
			// the current column and rows after that are all with zero head
			// go on to the next column
			if j >= mm.Row {
				return 0
			}
		}
		p := pivot[i]
		// eliminate the rows after the current one
		for j := i + 1; j < mm.Row; j++ {
			f := -mm.Data[j][i] / p
			mm.Data[j][i] = 0
			for h := i + 1; h < mm.Column; h++ {
				mm.Data[j][h] = mm.Data[j][h] + f*mm.Data[i][h]
			}
		}
		i++
	}
	d := 1.0
	for i := 0; i < mm.Row; i++ {
		d *= mm.Data[i][i]
	}
	return d
}

// inverse of the matrix, using the row transformation method
func (m *Matrix) Inverse() *Matrix {
	if m.Row != m.Column {
		panic("matrix with different row count and column count has no inversion")
	}
	eye := Eye(m.Row)
	mm := m.Copy()
	// zero the bottom left part
	for i := 0; i < m.Row; i++ {
		if mm.Data[i][i] == 0 {
			j := i + 1
			for ; j < mm.Row; j++ {
				if mm.Data[j][i] != 0 {
					for k := i; k < mm.Column; k++ {
						mm.Data[i][k] += mm.Data[j][k]
					}
					// all the operation to the eye must begin from first column
					for k := 0; k < eye.Column; k++ {
						eye.Data[i][k] += eye.Data[j][k]
					}
					break
				}
			}
			// matrix determinant is 0, so has no inversion
			if j >= mm.Row {
				panic("matrix with with determinant 0 has no inversion")
			}
		}
		p := mm.Data[i][i]
		for j := i; j < mm.Column; j++ {
			mm.Data[i][j] /= p
		}
		// all the operation to the eye must begin from first column
		for j := 0; j < eye.Column; j++ {
			eye.Data[i][j] /= p
		}

		for j := i + 1; j < mm.Row; j++ {
			p := -mm.Data[j][i]
			for k := i; k < mm.Column; k++ {
				mm.Data[j][k] = mm.Data[j][k] + p*mm.Data[i][k]
			}
			// all the operation to the eye must begin from first column
			for k := 0; k < mm.Column; k++ {
				eye.Data[j][k] = eye.Data[j][k] + p*eye.Data[i][k]
			}
		}
	}
	// zero the top right part
	for j := mm.Column - 1; j > 0; j-- {
		for i := j - 1; i > -1; i-- {
			p := -mm.Data[i][j]
			mm.Data[i][j] = 0
			// all the operation to the eye must begin from first column
			for k := 0; k < eye.Column; k++ {
				eye.Data[i][k] = eye.Data[i][k] + p*eye.Data[j][k]
			}
		}
	}

	return eye
}
