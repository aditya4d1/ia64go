package mmx_ia64

/*
 #include<mmintrin.h>
*/
import "C"

func MM_Addpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_add_pi16(a, b)
}

func MM_Addpi32(a C.__m64, b C.__m64) C.__m64 {
}
