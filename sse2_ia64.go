package sse2_ia64 

/* 
#include<> 
*/ 
import "C" 

func MM_Addepi16(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_add_epi16(a, b) 
} 

func MM_Addepi32(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_add_epi32(a, b) 
} 

func MM_Addepi64(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_add_epi64(a, b) 
} 

func MM_Addepi8(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_add_epi8(a, b) 
} 

func MM_Addepd(a C.__m128d, b C.__m128d) C.__m128d { 
        return C._mm_add_pd(a, b) 
} 

func MM_Addepsd(a C.__m128d, b C.__m128d) C.__m128d { 
        return C._mm_add_ps(a, b) 
} 

func MM_Addsi64(a C.__m64, b C.__m64) C.__m64 { 
        return C._mm_add_si64(a, b) 
} 

func MM_Addsepi16(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_adds_epi16(a, b) 
} 

func MM_Addsepi8(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_adds_epi8(a, b) 
} 

func MM_Addsepu16(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_adds_epu16(a, b) 
} 

func MM_Addsepu8(a C.__m128i, b C.__m128i) C.__m128i { 
        return C._mm_adds_epu8(a, b) 
} 

