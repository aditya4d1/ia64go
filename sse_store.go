

package sse_store

/*
 #include<xmmintrin.h>
*/
import "C"

func MM_Store1ps(x *float32, a C.__m128) {
	C._mm_store1_ps((*C.float)(x), (C.__m128)(a))
}

func MM_Storeps(x *float32, a C.__m128) {
	C._mm_store_ps((*C.float)(x), (C.__m128)(a))
}

func MM_Storeps1(x *float32, a C.__m128) {
	C._mm_store_ps1((*C.float)(x), (C.__m128)(a))
}

func MM_Storess(x *float32, a C.__m128) {
	C._mm_store_ss((*C.float)(x), (C.__m128)(a))
}

func MM_Storehpi(x *C.__m64, a C.__m128) {
	C._mm_storeh_pi((*C.__m64)(x), (C.__m128)(a))
}

func MM_Storelpi(x *C.__m64, a C.__m128) {
	C._mm_storel_pi((*C.__m64)(x), (C.__m128)(a))
}

func MM_Storerps(x *float32, a C.__m128) {
	C._mm_storer_ps((*C.float)(x), (C.__m128)(a))
}

func MM_Storeups(x *float32, a C.__m128) {
	C._mm_storeu_ps((*C.float)(x), (C.__m128)(a))
}
