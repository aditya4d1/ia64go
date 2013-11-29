package avx_ia64

/*
 #include<immintrin.h>
*/
import "C"

func M256_Addpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_add_pd(a, b)
}

func M256_Addps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_add_ps(a, b)
}

func M256_Addsubpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_addsub_pd(a, b)
}

func M256_Addsubps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_addsub_ps(a, b)
}

func M256_Andpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_and_pd(a, b)
}

func M256_Andps(a C.__m256, b C.__m256) C.__m256) C.__m256 {
	return C._mm256_and_ps(a, b)
}

func M256_Andnotpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_andnot_pd(a, b)
}

func M256_Andnotps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_andnot_ps(a, b)
}

func M256_Blendpd(a C.__m256d, b C.__m256d, c int) C.__m256d {
	return C._mm256_blend_pd(a, b, (C.int)(c))
}

func M256_Blendps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C._mm256_blend_ps(a, b, (C.int)(c))
}

func M256_Blendvpd(a C.__m256d, b C.__m256d, c C.__m256d) C.__m256d {
	return C._mm256_blendv_pd(a, b, c)
}

func M256_Blendvps(a C.__m256, b C.__m256, c C.__m256) C.__m256 {
	return C._mm256_blendv_ps(a, b, c)
}

func M256_Broadcastpd(a *C.__m128d) C.__m256d {
	return C._mm256_broadcast_pd(a)
}

func M256_Broadcastps(a *C.__m128) C.__m256 {
	return C._mm256_broadcast_ps(a)
}

func M256_Broadcastsd(a *float64) C.__m256d {
	return C._mm256_broadcast_sd((*C.double)(a))
}

func MM_Broadcastss(a *float32) C.__m128 {
	return C._mm_broadcast_ss((*C.float)(a))
}

func M256_Broadcastss(a *float32) C.__m256 {
	return C._mm256_broadcast_ss((*C.float)(a))
}

func M256_Castpd128pd256(a C.__m128d) C.__m256d {
	return C._mm256_castpd128_pd256(a)
}

func M256_Castpd256pd128(a C.__m256d) C.__m128d {
	return C._mm256_castpd256_pd128(a)
}

func M256_Castpdps(a C.__m256d) C.__m256 {
	return C._mm256_castpd_ps(a)
}

func M256_Castpdsi256(a C.__m256d) C.__m256i {
	return C._mm256_castpd_si256(a)
}

func M256_Castps128ps256(a C.__m128) C.__m256 {
	return C._mm256_castps128_ps256(a)
}

func M256_Castps256ps128(a C.__m256) C.__m128 {
	return C._mm256_castps256_ps128(a)
}

func M256_Castpspd(a C.__m256) C.__m256d {
	return C._mm256_castps_pd(a)
}

func M256_Castpssi256(a C.__m256) C.__m256i {
	return C._mm256_castps_si256(a)
}

func M256_Castsi128si256(a C.__m128) C.__m256i {
	return C._mm256_castsi128_si256(a)
}

func M256_Castsi256pd(a C.__m256i) C.__m256d {
	return C._mm256_castsi256_pd(a)
}

func M256_Castsi256ps(a C.__m256i) C.__m256 {
	return C._mm256_castsi256_ps(a)
}

func M256_Castsi256si128(a C.__m256i) C.__mm128i {
	return C._mm256_castsi256_si128(a)
}

func M256_Ceilpd(a C.__m256d) C.__m256d {
	return C._mm256_ceil_pd(a)
}

func M256_Ceilps(a C.__m256) C.__m256 {
	return C._mm256_ceil_ps(a)
}

func MM_Cmpps(a C.__m128d, b C.__m128d, c int) C.__m128d {
	return C._mm_cmp_pd(a, b, c)
}

func M256_Cmppd(a C.__m256d, b C.__m256d, c int) C.__m256d {
	return C._m256_cmp_pd(a, b, (C.int)(c))
}

func MM_Cmpps(a C.__m128, b C.__m128, c int) C.__m128 {
	return C._mm_cmp_ps(a, b, (C.int)(c))
}

func M256_Cmpps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C._m256_cmp_ps(a, b, (C.int)(c))
}

func MM_Cmpsd(a C.__m128d, b C.__m128d, c int) C.__m128d {
	return C._mm_cmp_sd(a, b, (C.int)(c))
}

func MM_Cmpss(a C.__m128, b C.__m128, c int) C.__m128 {
	return C._mm_cmp_ss(a, b, (C.int)(c))
}

func M256_Cvtepi32pd(a C.__m128i) C.__m256d {
	return C._mm256_cvtepi32_pd(a)
}

func M256_Cvtepi32ps(a C.__m256i) C.__m256 {
	return C._mm256_cvtepi32_ps(a)
}

func M256_Cvtpdepi32(a C.__m256d) C.__m128i {
	return C._mm256_cvtpd_epi32(a)
}

func M256_Cvtpdps(a C.__m256d) C.__m128 {
	return C._mm256_cvtpd_ps(a)
}

func M256_Cvtpsepi32(a C.__m256) C.__m256i {
	return C._mm256_cvtps_epi32(a)
}

func M256_Cvtpspd(a C.__m256d) C.__m256d {
	return C._mm256_cvtps_pd(a)
}

func M256_Cvttpdepi32(a C.__m256d) C.__m128i {
	return C._mm256_cvttpd_epi32(a)
}

func M256_Cvttpsepi32(a C.__m256) C.__m256i {
	return C._mm256_cvttps_epi32(a)
}

func M256_Divpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_div_pd(a, b)
}

func M256_Divps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_div_ps(a, b)
}

func M256_Dpps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C._mm256_dp_ps(a, b, (C.int)(c))
}

func M256_Extractf128pd(a C.__m256d, b int) C.__m128d {
	return C._mm256_extractf128_pd(a, (C.int)(b))
}

func M256_Extractf128ps(a C.__m256, b int) C.__m128 {
	return C._mm256_extract128_ps(a, (C.int)(b))
}

func M256_Extractf128si(a C.__m256i, b int) C.__m128i {
	return C._mm256_extract128_si(a, (C.int)(b))
}

func M256_Floorpd(a C.__m256d) C.__m256d {
	return C._mm256_floor_pd(a)
}

func M256_Floorps(a C.__m256) C.__m256 {
	return C._mm256_floor_ps(a)
}

func M256_Haddpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_hadd_ps(a, b)
}

func M256_Haddps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_hadd_ps(a, b)
}

func M256_Hsubpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_hsub_pd(a, b)
}

func M256_Hsubps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_hsub_ps(a, b)
}

func M256_Insertf128pd(a C.__m256d, b C.__m128d, c int) C.__m256d {
	return C._mm256_insertf128_pd(a, b, (C.int)(c))
}

func M256_Insertf128ps(a C.__m256, b C.__m128, c int) C.__m256 {
	return C._mm256_insertf128_ps(a, b, (C.int)(c))
}

func M256_Insertf128si256(a C.__m256i, b C.__m128i, c int) C.__m256i {
	return C._mm256_insertf128_si256(a, b, (C.int)(c))
}
