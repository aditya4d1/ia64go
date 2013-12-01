package avx_ia64

/*
 #include<immintrin.h>
 #include<avxintrin.h>
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

func M256_Maxpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_max_pd(a,b)
}

func M256_Maxps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_max_ps(a, b)
}

func M256_Minpd(a C.__m256d, c C.__m256d) C.__m256d {
	return C._mm256_min_pd(a, b)
}

func M256_Minps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_min_ps(a, b)
}

func M256_Moveduppd(a C.__m256d) C.__m256d {
	return C._mm256_movedup_pd(a)
}

func M256_Movehdupps(a C.__m256) C.__m256 {
	return C._mm256_movehdupps(a)
}

func M256_Moveldupps(a C.__m256) C.__m256 {
	return C._mm256_moveldup_ps(a)
}

func M256_Movemaskpd(a C.__m256d) int {
	return (int)(C._mm256_movemask_pd(a))
}

func M256_Movemaskps(a C.__m256) int {
	return (int)(C._mm256_movemask_ps(a))
}

func M256_Mulpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_mul_pd(a, b)
}

func M256_Mulps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_mul_ps(a, b)
}

func M256_Orpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_or_pd(a, b)
}

func M256_Orps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_or_ps(a, b)
}

func M256_Permute2f128pd(a C.__m256d, b C.__m256d, c int) C.__m256d {
	return C._mm256_permute2f128_pd(a, b, (C.int)(c))
}

func M256_Permute2f128ps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C._mm256_permute2f128_ps(a, b, (C.int)(c))
}

func M256_Permute2f128si256(a C.__m256i, b C.__m256i, c int) C.__m256i {
	return C._mm256_permute2f128_si256(a, b, (C.int)(c))
}

func MM_Permutepd(a C.__m128d, b int) C.__m128d {
	return C._mm_permute_pd(a, (C.int)(b))
}

func M256_Permutepd(a C.__m256d, b int) C.__m256d {
	return C._mm256_permute_pd(a, (C.int)(b))
}

func MM_Permuteps(a C.__m128, b int) C.__m128 {
	return C._mm_permute_ps(a, (C.int)(b))
}

func M256_Permuteps(a C.__m256, b int) C.__m256 {
	return C._mm256_permute_ps(a, (C.int)(b))
}

func MM_Permutevarpd(a C.__m128d, b C.__m128i) C.__m128d {
	return C._mm_permutevar_pd(a, b)
}

func M256_Permutevarpd(a C.__m256d, b C.__m256i) C.__m256d {
	return C._mm256_permutevar_pd(a, b)
}

func MM_Permutevarps(a C.__m128, b C.__m128i) C.__m128 {
	return C._mm_permutevar_ps(a, b)
}

func M256_Permutevarps(a C.__m256, b C.__m256i) C.__m256 {
	return C._mm256_permutevar_ps(a, b)
}

func M256_Rcpps(a C.__m256) C.__m256 {
	return C._mm256_rcp_ps(a)
}

func M256_Roundpd(a C.__m256d, b int) C.__m256d {
	return C._mm256_round_pd(a, (C.int)(b))
}

func M256_Roundps(a C.__m256, b int) C.__m256 {
	return C._mm256_round_ps(a, (C.int)(b))
}

func M256_Rsqrtps(a C.__m256) C.__m256 {
	return C._mm256_rsqrt_ps(a)
}
//Check here! int16 to short
func M256_Set1epi16(a int16) C.__m256i {
	return C._mm256_set1_epi16((C.short)(a))
}

func M256_Set1epi32(a int) C.__m256i {
	return C._mm256_set1_epi32((C.int)(a))
}

func M256_Set1epi64x(a int64) C.__m256i {
	return C._mm256_set1_epi64x((C.longlong)(a))
}

func M256_Set1epi8(a byte) C.__m256i {
	return C._mm256_set1_epi8((C.char)(a))
}

func M256_Set1pd(a float64) C.__m256d {
	return C._mm256_set1_pd((C.double)(a))
}

func M256_Set1ps(a float32) C.__m256 {
	return C._mm256_set1_ps((C.float)(a))
}

func M256_Setepi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) C.__m256i {
	return C._mm256_set_epi16((C.short)(e15), (C.short)(e14), (C.short)(e13), (C.short)(e12), (C.short)(e11), (C.short)(e10), (C.short)(e9), (C.short)(e8), (C.short)(e7), (C.short)(e6), (C.short)(e5), (C.short)(e4), (C.short)(e3), (C.short)(e2), (C.short)(e1), (C.short)(e0))
}

func M256_Setepi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) C.__m256i {
	return C._mm256_set_epi32((C.int)(e7), (C.int)(e6), (C.int)(e5), (C.int)(e4), (C.int)(e3), (C.int)(e2), (C.int)(e1), (C.int)(e0))
}

func M256_Setepi64x(a int64, b int64, c int64, d int64) C.__m256i {
	return C._mm256_set_epi64x((C.longlong)(a), (C.longlong)(b), (C.longlong)(c), (C.longlong)(d))
}

func M256_Setepi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) C.__m256i {
	return C._mm256_set_epi8((C.char)(e31), (C.char)(e30), (C.char)(e29), (C.char)(e28), (C.char)(e27), (C.char)(e26), (C.char)(e25), (C.char)(e24), (C.char)(e23), (C.char)(e22), (C.char)(e21), (C.char)(e20), (C.char)(e19), (C.char)(e18), (C.char)(e17), (C.char)(e16), (C.char)(e15), (C.char)(e14), (C.char)(e13), (C.char)(e12), (C.char)(e11), (C.char)(e10), (C.char)(e9), (C.char)(e8), (C.char)(e7), (C.char)(e6), (C.char)(e5), (C.char)(e4), (C.char)(e3), (C.char)(e2), (C.char)(e1), (C.char)(e0))
}

func M256_Setm128(a C.__m128, b C.__m128) C.__m256 {
	return C._mm256_set_m128(a, b)
}

func M256_Setm128d(a C.__m128d, b C.__m128d) C.__m256d {
	return C._mm256_set_m128d(a, b)
}

func M256_Setm128i(a C.__m128i, b C.__m128i) C.__m256i {
	return C._mm256_set_m128i(a, b)
}

func M256_Setpd(a float64, b float64, c float64, d float64) C.__m256d {
	return C._mm256_set_pd((C.double)(a), (C.double)(b), (C.double)(c), (C.double)(d))
}

func M256_Setps(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) C.__m256 {
	return C._mm256_set_ps((C.float)(e7), (C.float)(e6), (C.float)(e5), (C.float)(e4), (C.float)(e3), (C.float)(e2), (C.float)(e1), (C.float)(e0))
}

func M256_Setrepi16(e15 int16, e14 int16, e13 int16, e12 int16, e11 int16, e10 int16, e9 int16, e8 int16, e7 int16, e6 int16, e5 int16, e4 int16, e3 int16, e2 int16, e1 int16, e0 int16) C.__m256i {
	return C._mm256_setr_epi16((C.short)(e15), (C.short)(e14), (C.short)(e13), (C.short)(e12), (C.short)(e11), (C.short)(e10), (C.short)(e9), (C.short)(e8), (C.short)(e7), (C.short)(e6), (C.short)(e5), (C.short)(e4), (C.short)(e3), (C.short)(e2), (C.short)(e1), (C.short)(e0))
}

func M256_Setrepi32(e7 int, e6 int, e5 int, e4 int, e3 int, e2 int, e1 int, e0 int) C.__m256i {
	return C._mm256_setr_epi32((C.int)(e7), (C.int)(e6), (C.int)(e5), (C.int)(e4), (C.int)(e3), (C.int)(e2), (C.int)(e1), (C.int)(e0))
}

func M256_Setrepi64x(a int64, b int64, c int64, d int64) C.__m256i {
	return C._mm256_setr_epi64x((C.longlong)(a), (C.longlong)(b), (C.longlong)(c), (C.longlong)(d))
}

func M256_Setrepi8(e31 byte, e30 byte, e29 byte, e28 byte, e27 byte, e26 byte, e25 byte, e24 byte, e23 byte, e22 byte, e21 byte, e20 byte, e19 byte, e18 byte, e17 byte, e16 byte, e15 byte, e14 byte, e13 byte, e12 byte, e11 byte, e10 byte, e9 byte, e8 byte, e7 byte, e6 byte, e5 byte, e4 byte, e3 byte, e2 byte, e1 byte, e0 byte) C.__m256i {
	return C._mm256_setr_epi8((C.char)(e31), (C.char)(e30), (C.char)(e29), (C.char)(e28), (C.char)(e27), (C.char)(e26), (C.char)(e25), (C.char)(e24), (C.char)(e23), (C.char)(e22), (C.char)(e21), (C.char)(e20), (C.char)(e19), (C.char)(e18), (C.char)(e17), (C.char)(e16), (C.char)(e15), (C.char)(e14), (C.char)(e13), (C.char)(e12), (C.char)(e11), (C.char)(e10), (C.char)(e9), (C.char)(e8), (C.char)(e7), (C.char)(e6), (C.char)(e5), (C.char)(e4), (C.char)(e3), (C.char)(e2), (C.char)(e1), (C.char)(e0))
}

func M256_Setrm128(a C.__m128, b C.__m128) C.__m256 {
	return C._mm256_setr_m128(a, b)
}

func M256_Setrm128d(a C.__m128d, b C.__m128d) C.__m256d {
	return C._mm256_setr_m128d(a, b)
}

func M256_Setrm128i(a C.__m128i, b C.__m128i) C.__m256i {
	return C._mm256_setr_m128i(a, b)
}

func M256_Setrpd(a float64, b float64, c float64, d float64) C.__m256d {
	return C._mm256_setr_pd((C.double)(a), (C.double)(b), (C.double)(c), (C.double)(d))
}

func M256_Setrps(e7 float32, e6 float32, e5 float32, e4 float32, e3 float32, e2 float32, e1 float32, e0 float32) C.__m256 {
	return C._mm256_setr_ps((C.float)(e7), (C.float)(e6), (C.float)(e5), (C.float)(e4), (C.float)(e3), (C.float)(e2), (C.float)(e1), (C.float)(e0))
}

func M256_Setzeropd() C.__m256d {
	return C._mm256_setzero_pd()
}

func M256_Setzerops() C.__m256 {
	return C._mm256_setzero_ps()
}

func M256_Setzerosi256() C.__m256i {
	return C._mm256_setzero_si256()
}

func M256_Shufflepd(a C.__m256d, b C.__m256d, c int) C.__m256d {
	return C._mm256_shuffle_pd(a, b, (C.int)(c))
}

func M256_Shuffleps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C._mm256_shuffle_ps(a, b, (C.int)(c))
}

func M256_Sqrtpd(a C.__m256d) C.__m256d {
	return C._mm256_sqrt_pd(a)
}

func M256_Sqrtps(a C.__m256) C.__m256 {
	return C._mm256_sqrt_ps(a)
}

func M256_Streampd(a *float64) {
	C._mm256_stream_pd((*C.double)(a))
}

func M256_Streamps(a *float32) {
	C._mm256_stream_ps((*C.float)(a))
}

func M256_Streamsi256(a *C.__m256i, v C.__m256i) {
	C._mm256_stream_si256(a, b)
}

func M256_Subpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_sub_pd(a, b)
}

func M256_Subps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_sub_ps(a, b)
}
