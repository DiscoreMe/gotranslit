package gotranslit

import (
	"testing"
)

func BenchmarkTest(b *testing.B) {
	s := "АаБбВвГгДдЕе ЁёЖжЗзИиЙйКкЛ лМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщ"
	for i := 0; i < b.N; i++ {
		ToCyrillic(s)
	}
}

func TestTranslit(t *testing.T) {
	s := "АаБбВвГгДдЕе ЁёЖжЗзИиЙйКкЛ лМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯя"
	out := ToCyrillic(s)
	if out != "AaBbVvGgDdEe_YOyoZHzhZzIiJjKkL_lMmNnOoPpRrSsTtUuFfHhCcCHchSHshSCHschYyEeJUjuJAja" {
		t.Errorf("%s != AaBbVvGgDdEe_YOyoZHzhZzIiJjKkL_lMmNnOoPpRrSsTtUuFfHhCcCHchSHshSCHschYyEeJUjuJAja", out)
	}
}
