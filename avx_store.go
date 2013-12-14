
package avx_store

/*
 #cgo CFLAGS: -mavx
 #include<immintrin.h>
*/
import "C"

func MM_Maskstorepd(a *float64, b C.__m128i, c C.__m128d) {
	C._mm_maskstore_pd((*C.double)(a), b, c)
}

func M256_Maskstorepd(a *float64, b C.__m256i, c C.__m256d) {
	C._mm256_maskstore_pd((*C.double)(a), b, c)
}

func MM_Maskstoreps(a *float32, b C.__m128i, c C.__m128) {
	C._mm_maskstore_ps((*C.float)(a), b, c)
}

func M256_Maskstoreps(a *float32, b C.__m256i, c C.__m256) {
	C._mm256_maskstore_ps((*C.float)(a), b, c)
}

func M256_Storepd(a *float64, b C.__m256d) {
	C._mm256_store_pd((*C.double)(a), b)
}

func M256_Storeps(a *float32, b C.__m256) {
	C._mm256_store_ps((*C.float)(a), b)
}

func M256_Storesi256(a *C.__m256i, b C.__m256i) {
	C._mm256_store_si256(a, b)
}
/*
func M256_Storeu2m128(a *float32, b *float32, c C.__m256) {
	C._mm256_storeu2_m128((*C.float)(a), (*C.float)(b), c)
}

func M256_Storeu2m128d(a *float64, b *float64, c C.__m256d) {
	C._mm256_storeu2_m128d((*C.longlong)(a), (*C.longlong)(b), c)
}

func M256_Storeu2m128i(a *C.__m128i, b *C.__m128i, c C.__m256i) {
	C._mm256_storeu2_m128i(a, b, c)
}
*/
func M256_Storeupd(a *float64, b C.__m256d) {
	C._mm256_storeu_pd((*C.double)(a), b)
}

func M256_Storeups(a *float32, b C.__m256) {
	C._mm256_storeu_ps((*C.float)(a), b)
}

func M256_Storeusi256(a *C.__m256i, b C.__m256i) {
	C._mm256_storeu_si256(a, b)
}
