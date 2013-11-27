package sse_load

/*
 #include<xmmintrin.h>
*/
import "C"

func MM_Load1ps(x *float32) C.__m128 {
	return C._mm_load1_ps((*C.float)(x))
}

func MM_Loadps(x *float32) C.__m128 {
	return C._mm_load_ps((*C.float)(x))
}

func MM_Loadps1(x *float32) C.__m128 {
	return C._mm_load_ps1((*C.float)(x))
}

func MM_Loadss(x *float32) C.__m128 {
	return C._mm_load_ss((*C.float)(x))
}

func MM_Loadhpi(a C.__m128, x *C.__m64) C.__m128 {
	return C._mm_loadh_pi((C.__m128)(a), (*C.__m64)(x))
}

func MM_Loadlpi(a C.__m128, x *C.__m64) C.__m128 {
	return C._mm_loadl_pi((C.__m128)(a), (*C.__m64)(x))
}

func MM_Loadrps(x *float32) C.__m128 {
	return C._mm_loadr_ps((*C.float)(x))
}

func MM_Loadups(x *float32) C.__m128 {
	return C._mm_loadu_ps((*C.float)(x))
}
