
package mmx_ia64

/*
 #include<mmintrin.h>
*/

func MM_Addpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_add_pi16(a, b)
}

func MM_Addpi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_add_pi32(a, b)
}

func MM_Addpi8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_add_pi8(a, b)
}

func MM_Addspi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_adds_pi16(a, b)
}

func MM_Addspi8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_adds_pi8(a, b)
}

func MM_Addspu16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_adds_pu16(a, b)
}

func MM_Addspu8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_adds_pu8(a, b)
}

func MM_Andsi64(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_and_si64(a, b)
}

func MM_Andnotsi64(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_andnot_si64(a, b)
}

func MM_Cmpeqpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_cmpeq_pi16(a, b)
}

func MM_Cmpeqpi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_cmpeq_pi8(a, b)
}

func MM_Cmpeqpi8 (a C.__m64,  b C.__m64) C.__m64 {
	return C._mm_cmpeq_pi8(a, b)
}

func MM_Cmpgtpi16 (a C.__m64, b C.__m64) C.__m64 {
	return C._mm_cmpgt_pi16(a, b)
}

func MM_Cmpgtpi32 (a C.__m64, b C.__m64) C.__m64 {
	return C._mm_cmpgt_pi32(a, b)
}

func MM_Cmpgetpi8 (a C.__m64, b C.__m64) C.__m64 {
	return C._mm_cmpgt_pi8(a, b);
}

func MM_Cvtm64si64(a C.__m64, b C.__m64) int64 {
	return int64(C._mm_cvtm64_si64(a , b ))
}

func MM_Cvtsi32_si64 (a int) C.__m64 {
	return C._mm_cvtsi32_si64((C.int)(a))
}

func MM_Cvtsi64m64(a int64) C.__m64 {
	return C._mm_cvtsi64_m64((C.longlong)(a))
}

func MM_Cvtsi64si32(a C.__m64) int {
	return int(C._mm_cvtsi64_si32(a))
}

func MM_Empty() {
	C._mm_empty()
}
func M_Empty() {
	C._m_empty()
}

func M_Fromint(a int) C.__m64 {
	return C._m_from_int((C.int)(a))
}

func M_Fromint64(a int64) C.__m64 {
	return C._m_from_int64((C.longlong)(a))
}

func M_Packssdw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_packssdw(a, b)
}

func M_Packsswb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_packsswb
}

func M_Packuswb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_packuswb(a, b)
}

func M_Paddb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddb(a, b)
}

func M_Paddd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddd(a, b)
}

func M_Paddsb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddsb(a, b)
}

func M_Paddsw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddsw(a, b)
}

func M_Paddusb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddusb(a, b)
}

func M_Paddusw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddusw(a, b)
}

func M_Paddw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_paddw(a, b)
}

func M_Pand(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pand(a, b)
}

func M_Pandn(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pandn(a, b)
}

func M_Pcmpeqb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pcmpeqb(a, b)
}

func M_Pcmpeqd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pcmpeqd(a, b)
}

func M_Pcmpeqw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pcmpeqw(a, b)
}

func M_Pcmpgtb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pcmpgtb(a, b)
}

func M_Pcmpgtd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pcmpgtd(a, b)
}

func M_Pcmpgtw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pcmpgtw(a, b)
}

func M_Pmaddwd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pmaddwd(a, b)
}

func M_Pmulhw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pmulhw(a, b)
}

func M_Pmullw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pmullw(a, b)
}

func M_Por(a C.__m64, b C.__m64) C.__m64 {
	return C._m_por(a, b)
}

func M_Pslld(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pslld(a, b)
}

func M_Pslldi(a C.__m64, b int) C.__m64 {
	return C._m_pslldi(a, (C.int)(b))
}

func M_Psllq(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psllq(a, b)
}

func M_Psllqi(a C.__m64, b int) C.__m64 {
	return C._m_psllqi(a, (C.int)(b))
}

func M_Psllw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psllw(a, b)
}

func M_Psllwi(a C.__m64, b int) C.__m64 {
	return C._m_psllwi(a, (C.int)(b))
}

func M_Psrad(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psrad(a, b)
}

func M_Psradi(a C.__m64, b int) C.__m64 {
	return C._m_psradi(a, (C.int)(b))
}

func M_Psraw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psraw(a, b)
}

func M_Psrawi(a C.__m64, b int) C.__m64 {
	return C._m_psrawi(a, (C.int)(b))
}

func M_Psrld(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psrld(a, b)
}

func M_Psrldi(a C.__m64, b int) C.__m64 {
	return C._m_psrldi(a, (C.int)(b))
}

func M_Psrlq(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psrlq(a, b)
}

func M_Psrlqi(a C.__m64, b int) C.__m64 {
	return C._m_psrlqi(a, (C.int)(b))
}

func M_Psrlw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psrlw(a, b)
}

func M_Psrlwi(a C.__m64, b int) C.__m64 {
	return C._m_psrlwi(a, (C.int)(b))
}

func M_Psubb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psubb(a, b)
}

func M_Psubd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psubd(a, b)
}

func M_Psubsb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psubsb(a, b)
}

func M_Psubusb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psubusb(a, b)
}

func M_Psubusw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psubusw(a, b)
}

func M_Psubw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psubw(a, b)
}

func M_Punpckhbw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_punpckhbw(a, b)
}

func M_Punpckhdq(a C.__m64, b C.__m64) C.__m64 {
	return C._m_punpckhdq(a, b)
}

func M_Punpckhwd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_punpckhwd(a, b)
}

func M_Punpcklbw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_punpcklbw(a, b)
}

func M_Punpckldq(a C.__m64, b C.__m64) C.__m64 {
	return C._m_punpckldq(a, b)
}

func M_Punpcklwd(a C.__m64, b C.__m64) C.__m64 {
	return C._m_punpcklwd(a, b)
}

func M_Pxor(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pxor(a, b)
}

func M_Toint(a C.__m64) int {
	return (int)(C._m_to_int(a))
}

func M_Toint64(a C.__m64){
	return (int64)(C._m_to_int64(a))
}

func MM_Maddpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_madd_pi16(a, b)
}

func MM_Mulhipi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_mulhi_pi16(a, b)
}

func MM_Mullopi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_mullo_pi16(a, b)
}

func MM_Orsi64(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_or_si64(a, b)
}

func MM_Packspi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_packs_pi16(a, b)
}

func MM_Packspi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_packs_pi32(a, b)
}

func MM_Packspu16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_packs_pu16(a, b)
}

func MM_Set1pi16(a int16, b int16, c int16, d int16) C.__m64 {
	return C._mm_set1_pi16((C.short)(a), (C. short)(b), (C.short)(c), (C.short)(d))
}

func MM_Set1pi32(a int, b int) C.__m64 {
	return C._mm_set1_pi32((C.int)(a),(C.int)(b))
}

func MM_Set1pi8(a byte, b byte, c byte, d byte, e byte, f byte, g byte, h byte) C.__m64 {
	return C._mm_set1_pi8((C.char)(a), (C.char)(b), (C.char)(c), (C.char)(d), (C.char)(e), (C.char)(f), (C.char)(g), (C.char)(h))
}

func MM_Setrpi16(a int16, b int16, c int16, d int16) C.__m64 {
	return C._mm_setr_pi16((C.short)(a), (C.short)(b), (C.short)(c), (C.short)(d))
}

func MM_Setrpi32(a int, b int) C.__m64 {
	return C._mm_setr_pi32((C.int)(a),(C.int)(b))
}

func MM_Setrpi8(a byte, b byte, c byte, d byte, e byte, f byte, g byte, h byte) C.__m64 {
	return C._mm_setr_pi8((C.char)(a), (C.char)(b), (C.char)(c), (C.char)(d), (C.char)(e), (C.char)(f), (C.char)(g), (C.char)(h))
}

func MM_Setzerosi64() C.__m64 {
	return C._mm_setzero_si64()
}

func MM_Sllpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sll_pi16(a, b)
}

func MM_Sllpi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sll_pi32(a, b)
}

func MM_Sllsi64(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sll_si64(a, b)
}

func MM_Sllipi16(a C.__m64, b int) C.__m64 {
	return C._mm_slli_pi16(a, (C.int)(b))
}

func MM_Sllipi32(a C.__m64, b int) C.__m64 {
	return C._mm_slli_pi32(a, (C.int)(b))
}

func MM_Sllipi64(a C.__m64, b int) C.__m64 {
	return C._mm_slli_pi64(a, (C.int)(b))
}

func MM_Srapi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sra_pi16(a, b)
}

func MM_Srapi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sra_pi32(a, b)
}

func MM_Sraipi16(a C.__m64, b int) C.__m64 {
	return C._mm_srai_pi16(a, (C.int)(b))
}

func MM_Sraipi32(a C.__m64, b int) C.__m64 {
	return C._mm_srai_pi32(a, (C.int)(b))
}

func MM_Srlpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_srl_pi16(a, b)
}

func MM_Srlpi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_srl_pi32(a, b)
}

func MM_Srlsi64(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_srl_si64(a, b)
}

func MM_Srlipi16(a C.__m64, b int) C.__m64 {
	return C._mm_srli_pi16(a, (C.int)(b))
}

func MM_Srlipi32(a C.__m64, b int) C.__m64 {
	return C._mm_srli_pi32(a, (C.int)(b))
}

func MM_Srlisi64(a C.__m64, b int) C.__m64 {
	return C._mm_srli_si64(a, (C.int)(b))
}

func MM_Subpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sub_pi16(a, b)
}

func MM_Subpi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sub_pi32(a, b)
}

func MM_Subpi8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_sub_pi8(a, b)
}

func MM_Subspi8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_subs_pi8(a, b)
}

func MM_Subspi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_subs_pi16(a, b)
}

func MM_Subspu16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_subs_pu16(a, b)
}

func MM_Subspu8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_subs_pu8(a, b)
}

func MM_Unpackhipi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_unpackhi_pi16(a, b)
}

func MM_Unpackhipi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_unpackhi_pi32(a, b)
}

func MM_Unpackhipi8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_unpackhi_pi8(a, b)
}

func MM_Unpacklopi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_unpacklo_pi16(a, b)
}

func MM_Unpacklopi32(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_unpacklo_pi32(a, b)
}

func MM_Unpacklopi8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_unpacklo_pi8(a, b)
}

func MM_Xorsi64(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_xor_si64(a, b)
}
