package gofs

import "testing"

func TestGoFs(t *testing.T) {
	p := PathHandler("testpath")
	p.Dir().CreateIfNotExist()
	dir := p.Dir()
	dir.Delete(true)
}
