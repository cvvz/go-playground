package deepcopy

import (
	"testing"
)

func TestDeepCopy(t *testing.T) {
	s2 := S2{
		i:  2,
		s3: &S3{"str"},
	}

	s := S{
		i:   1,
		pI:  new(int32),
		s2:  s2,
		pS2: &s2,
		str: "str",
	}

	sCopy := s

	// 浅拷贝的问题：没有新建内存，只是复制了指针指向同一片内存，所以修改sCopy的值，会影响到s的值
	sCopy.s2.i = 200
	sCopy.pS2.i = 200
	t.Logf("original: %d %d, copy: %d %d", s.s2.i, s.pS2.i, sCopy.s2.i, sCopy.pS2.i)
}
