package point

import "testing"

func TestPrintPoint(t *testing.T) {
	p := Point{X : 3, Y: 2}
	chaineCalcul := PrintPoint(p)
	chaineResultat := "(3,2)"
    if chaineCalcul != chaineResultat {
       t.Errorf("PrintPoint() incorrect : calcule: \"%s\", et veut: \"%s\".", chaineCalcul, chaineResultat)
    }
}

func TestPrintPointReverse(t *testing.T) {
	p := Point{X : 3, Y: 2}
	chaineCalcul := PrintPointReverse(p)
	chaineResultat := "(2,3)"
    if chaineCalcul != chaineResultat {
       t.Errorf("PrintPoint() incorrect : calcule: \"%s\", et veut: \"%s\".", chaineCalcul, chaineResultat)
    }
}