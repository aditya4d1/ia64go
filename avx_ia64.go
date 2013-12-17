package avx_ia64

/*
 #cgo CFLAGS: -mavx
 #include<immintrin.h>
 __m256d MM256PdCmp(__m256d a, __m256d b, int c){
	return (__m256d)_mm256_cmp_pd(a, b, c);
 }
 __m256 MM256PsCmp(__m256 a, __m256 b, int c){
	return (__m256)_mm256_cmp_ps(a, b, c);
 }
 __m128d MMPdCmp(__m128d a, __m128d b, int c){
	return (__m128d)_mm_cmp_pd(a, b, c);
 }
 __m128d MMSdCmp(__m128d a, __m128d b, int c){
	return (__m128d)_mm_cmp_sd(a, b, c);
 }
 __m128 MMSsCmp(__m128 a, __m128 b, int c){
	return (__m128)_mm_cmp_ss(a, b, c);
 }
 __m128 MMPsCmp(__m128 a, __m128 b, int c){
	return (__m128)_mm_cmp_ps(a, b, c);
 }
 __m256d MM256PdInsertf128(__m256d a, __m128d b, int c){
	return (__m256d)_mm256_insertf128_pd(a, b, c);
 }
 __m256 MM256PsInsertf128(__m256 a, __m128 b, int c){
	return (__m256)_mm256_insertf128_ps(a, b, c);
 }
 __m256d MM256PdFloor(__m256d a){
	return (__m256d)_mm256_floor_pd(a);
 }
 __m256 MM256PsFloor(__m256 a){
	return (__m256)_mm256_floor_ps(a);
 }
 __m256d MM256PdCeil(__m256d a){
	return (__m256d)_mm256_ceil_pd(a);
 }
 __m256 MM256PsCeil(__m256 a){
	return (__m256)_mm256_ceil_ps(a);
 }
 __m256d MM256PdPermute(__m256d a, int b){
	return (__m256d)_mm256_permute_pd(a, b);
 }
 __m256 MM256PsPermute(__m256 a, int b){
	return (__m256)_mm256_permute_ps(a, b);
 }
 __m128d MMPdPermute(__m128d a, int b){
	return (__m128d)_mm_permute_pd(a, b);
 }
 __m128 MMPsPermute(__m128 a, int b){
	return (__m128)_mm_permute_ps(a, b);
 }
 __m256d MM256PdBlend(__m256d a, __m256d b, int c){
	return (__m256d)_mm256_blend_pd(a, b, c);
 }
 __m256 MM256PsBlend(__m256 a, __m256 b, int c){
	return (__m256)_mm256_blend_ps(a, b, c);
 }
 __m256d MM256PdRound(__m256d a, int b){
	return (__m256d)_mm256_round_pd(a, b);
 }
 __m256 MM256PsRound(__m256 a, int b){
	return (__m256)_mm256_round_ps(a, b);
 }
/*
 Not defined in all GCC
 __m256 MM256M128Set(__m128 a, __m128 b){
	return (__m256)_mm256_set_m128(a, b);
 }
 __m256d MM256M128dSet(__m128d a, __m128d b){
	return (__m256d)_mm256_set_m128d(a, b);
 }
 __m256i MM256M128iSet(__m128i a, __m128i b){
	return (__m256i)_mm256_set_m128i(a, b);
 }
 __m256 MM256M128Setr(__m128 a, __m128 b){
	return (__m256)_mm256_setr_m128(a, b);
 }
 __m256i MM256M128iSetr(__m128i a, __m128i b){
	return (__m256i)_mm256_setr_m128i(a, b);
 }
 __m256d MM256M128dSetr(__m128d a, __m128d b){
	return (__m256d)_mm256_setr_m128d(a, b);
 }
 __m128d MM256PdExtractf128(__m256d a, int b){
	return (__m128d)_mm256_extractf128_pd(a, b);
 }
 __m128 MM256PsExtractf128(__m256 a, int b){
	return (__m128)_mm256_extractf128(a, b);
 }
 __m128i MM256Si256Extractf128(__m256i a, int b){
	return (__m128i)_mm256_extractf128_si256(a, b);
 }
 __m256d MM256PdPermute2f128(__m256d a, __m256d b, int c){
	return (__m256d)_mm256_permute2f128_pd(a, b, c);
 }
 __m256 MM256PsPermute2f128(__m256 a, __m256 b, int c){
	return (__m256)_mm256_permute2f128_ps(a, b, c);
 }
 __m256i MM256Si256Permute2f128(__m256i a, __m256i b, int c){
	return (__m256i)_mm256_permute2f128_si256(a, b, c);
 }
 __m256i MM256Si256Insertf128(__m256i a, __m128i b, int c){
	return (__m256i)_mm256_insertf128_si256(a, b, c);
 }
 __m256 MM256PsShuffle(__m256 a, __m256 b, int c){
	return (__m256)_mm256_shuffle_ps(a, b, c);
 }
 __m256d MM256PdShuffle(__m256d a, __m256d b, int c){
	return (__m256d)_mm256_shuffle_pd(a, b, c);
 }
 __m256 MM256PsDp(__m256 a, __m256 b, int c){
	return (__m256)_mm256_dp_ps(a, b, c);
 }
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

func M256_Andps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_and_ps(a, b)
}

func M256_Andnotpd(a C.__m256d, b C.__m256d) C.__m256d {
	return C._mm256_andnot_pd(a, b)
}

func M256_Andnotps(a C.__m256, b C.__m256) C.__m256 {
	return C._mm256_andnot_ps(a, b)
}

func M256_Blendpd(a C.__m256d, b C.__m256d, c int) C.__m256d {
	return C.MM256PdBlend(a, b, (C.int)(c))
}

func M256_Blendps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C.MM256PsBlend(a, b, (C.int)(c))
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

func M256_Castsi256si128(a C.__m256i) C.__m128i {
	return C._mm256_castsi256_si128(a)
}

func M256_Ceilpd(a C.__m256d) C.__m256d {
	return C.MM256PdCeil(a)
}

func M256_Ceilps(a C.__m256) C.__m256 {
	return C.MM256PsCeil(a)
}

func MM_Cmpps(a C.__m128d, b C.__m128d, c int) C.__m128d {
	return C.MMPsCmp(a, b, c)
}

func M256_Cmppd(a C.__m256d, b C.__m256d, c int) C.__m256d {
	return C.MM256PdCmp(a, b, (C.int)(c))
}

func MM_Cmpps(a C.__m128, b C.__m128, c int) C.__m128 {
	return C.MMPsCmp(a, b, (C.int)(c))
}

func M256_Cmpps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C.MM256PsCmp(a, b, (C.int)(c))
}

func MM_Cmpsd(a C.__m128d, b C.__m128d, c int) C.__m128d {
	return C.MMSdCmp(a, b, (C.int)(c))
}

func MM_Cmpss(a C.__m128, b C.__m128, c int) C.__m128 {
	return C.MMSsCmp(a, b, (C.int)(c))
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
	return C.MM256PsDp(a, b, (C.int)(c))
}

func M256_Extractf128pd(a C.__m256d, b int) C.__m128d {
	return C.MM256PdExtractf128(a, (C.int)(b))
}

func M256_Extractf128ps(a C.__m256, b int) C.__m128 {
	return C.MM256PsExtractf128(a, (C.int)(b))
}

func M256_Extractf128si256(a C.__m256i, b int) C.__m128i {
	return C.MM256Si256Extractf128(a, (C.int)(b))
}

func M256_Floorpd(a C.__m256d) C.__m256d {
	return C.MM256PdFloor(a)
}

func M256_Floorps(a C.__m256) C.__m256 {
	return C.MM256PsFloor(a)
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
	return C.MM256PdInsertf128(a, b, (C.int)(c))
}

func M256_Insertf128ps(a C.__m256, b C.__m128, c int) C.__m256 {
	return C.MM256PsInsertf128(a, b, (C.int)(c))
}

func M256_Insertf128si256(a C.__m256i, b C.__m128i, c int) C.__m256i {
	return C.MM256Si256Insertf128(a, b, (C.int)(c))
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
	return C._mm256_movehdup_ps(a)
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
	return C.MM256PdPermute2f128(a, b, (C.int)(c))
}

func M256_Permute2f128ps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C.MM256PsPermute2f128(a, b, (C.int)(c))
}

func M256_Permute2f128si256(a C.__m256i, b C.__m256i, c int) C.__m256i {
	return C.MM256Si256Permute2f128(a, b, (C.int)(c))
}

func MM_Permutepd(a C.__m128d, b int) C.__m128d {
	return C.MMPdPermute(a, (C.int)(b))
}

func M256_Permutepd(a C.__m256d, b int) C.__m256d {
	return C.MM256PdPermute(a, (C.int)(b))
}

func MM_Permuteps(a C.__m128, b int) C.__m128 {
	return C.MMPsPermute(a, (C.int)(b))
}

func M256_Permuteps(a C.__m256, b int) C.__m256 {
	return C.MM256PsPermute(a, (C.int)(b))
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
	return C.MM256PdRound(a, (C.int)(b))
}

func M256_Roundps(a C.__m256, b int) C.__m256 {
	return C.MM256PsRound(a, (C.int)(b))
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
	return C.MM256M128Set(a, b)
}

func M256_Setm128d(a C.__m128d, b C.__m128d) C.__m256d {
	return C.MM256M128dSet(a, b)
}

func M256_Setm128i(a C.__m128i, b C.__m128i) C.__m256i {
	return C.MM256M128iSet(a, b)
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
	return C.MM256M128Setr(a, b)
}

func M256_Setrm128d(a C.__m128d, b C.__m128d) C.__m256d {
	return C.MM256M128dSetr(a, b)
}

func M256_Setrm128i(a C.__m128i, b C.__m128i) C.__m256i {
	return C.MM256M128iSetr(a, b)
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
	return C.MM256PdShuffle(a, b, (C.int)(c))
}

func M256_Shuffleps(a C.__m256, b C.__m256, c int) C.__m256 {
	return C.MM256PsShuffle(a, b, (C.int)(c))
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
