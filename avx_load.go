package avx_load

/*
 #cgo CFLAGS: -mavx
 #include<immintrin.h>
*/
import "C"

func MM256_Lddqusi256(a *C.__m256i) C.__m256i {
	return C._mm256_lddqu_si256(a)
}

func MM256_Loadpd(a *float64) C.__m256d {
	return C._mm256_load_pd((*C.double)(a))
}

func MM256_Loadps(a *float32) C.__m256 {
	return C._mm256_load_ps((*C.float)(a))
}

func MM256_Loadsi256(a *C.__m256i) C.__m256i {
	return C._mm256_load_si256(a)
}
/*
func MM256_Loadu2m128(a *float32, b *float32) C.__m256 {
	return C._mm256_loadu2_m128((*C.float)(a),(*C.float)(b))
}

func MM256_Loadu2m128d(a *float64, b *float64) C.__m256d {
	return C._mm256_loadu2_m128d((*C.double)(a), (*C.double)(b))
}

func MM256_Loadu2m128i(a *C.__m128i, b *C.__m128i) C.__m256i {
	return C._mm256_loadu2_m128i(a, b)
}
*/
func MM256_Loadupd(a *float64) C.__m256d {
	return C._mm256_loadu_pd((*C.double)(a))
}

func MM256_Loadups(a *float32) C.__m256 {
	return C._mm256_loadu_ps((*C.float)(a))
}

func MM256_Loadusi256(a *C.__m256i) C.__m256i {
	return C._mm256_loadu_si256(a)
}

func MM_Maskloadpd(a *float64, b C.__m128i) C.__m128d {
	return C._mm_maskload_pd((*C.double)(a), b)
}

func MM256_Maskloadpd(a *float64, b C.__m256i) C.__m256d {
	return C._mm256_maskload_pd((*C.double)(a), b)
}

func MM_Maskloadps(a *float32, b C.__m128i) C.__m128 {
	return C._mm_maskload_ps((*C.float)(a), b)
}

func MM256_Maskloadps(a *float32, b C.__m256i) C.__m256 {
	return C._mm256_maskload_ps((*C.float)(a), b)
}
