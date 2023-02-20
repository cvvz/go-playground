package deepcopy

type S struct {
	i   int32
	pI  *int32
	s2  S2
	pS2 *S2
	str string
}

type S2 struct {
	i  int
	s3 *S3
}

type S3 struct {
	str string
}
