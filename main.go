package main

import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

/*
#cgo LDFLAGS: -lpcre2-32
#cgo CFLAGS: -DPCRE2_CODE_UNIT_WIDTH=32
#include <pcre2.h>
#include <stdio.h>
#include <string.h>
long  lens = 256;
pcre2_code * compile_regex( char *pattern) {
	printf("%s\n", pattern);
   PCRE2_SIZE erroff;
   int errorcode;
   pcre2_code *re = pcre2_compile((PCRE2_SPTR)pattern, PCRE2_ZERO_TERMINATED, 0, &errorcode, &erroff, NULL);
   if (re == NULL) {
		printf("tttttttttttttttttttttttt\n");
		PCRE2_UCHAR buffer[256];
        pcre2_get_error_message(errorcode, buffer, sizeof(buffer));
        printf("PCRE2 compilation failed at offset %lu: %u\n", (unsigned long)erroff, buffer[0]);
		return NULL;
   }
	printf("rrrrrrrrrrrrrrrrrrrrrr\n");
   return re;
}

int match_regex(pcre2_code  *re,  char *subject) {
	PCRE2_SIZE *ovector;
	PCRE2_SPTR s = (PCRE2_SPTR)subject;
	PCRE2_SIZE len = strlen(subject);

	printf("%s\n", subject);
	pcre2_match_data *match_data = pcre2_match_data_create_from_pattern(re, NULL);
	int rc = pcre2_match(re, s, len, 0, 0, match_data, NULL);
	if (rc < 0){
		switch(rc){
			case PCRE2_ERROR_NOMATCH: printf("No match\n"); break;
			default: printf("Matching error %d\n", rc); break;
		}
		pcre2_match_data_free(match_data);
		pcre2_code_free(re);
		return -1;
	}
	ovector = pcre2_get_ovector_pointer(match_data);
	printf("\nMatch succeeded at offset %d\n", (int)ovector[0]);
	return rc;
}
*/
import "C"

func main() {
	// 目标文本
	target := "a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999"

	// 筛选规则
	pattern := `(?<=\d{4})[^0-9\s]{3,11}(?!\s|$)`
	//pattern := `^a`
	// 编译正则表达式
	cPattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cPattern))
	re := C.compile_regex(cPattern)
	if re == nil {
		os.Exit(1)
	}
	defer C.pcre2_code_free(re)
	fmt.Println(re)

	// 匹配目标文本
	subject := C.CString(target)
	defer C.free(unsafe.Pointer(subject))
	rc := C.match_regex(re, subject)
	if rc < 0 {
		fmt.Printf("No match\n")
		os.Exit(1)
	}
	fmt.Println(rc)
	// 提取结果字符串
	//matchData := make([]C.int, rc*2)
	//rc = C.pcre2_get_ovector_pointer(*C.pcre2_match_data)
	//if rc < 0 {
	//	fmt.Printf("Error getting ovector pointer: %d\n", rc)
	//	os.Exit(1)
	//}
	//result := target[int(matchData[2]):int(matchData[3])]
	//
	//// 检查结果字符串是否符合规则
	//validPattern := `^\D{3,11}$`
	//validRegex, err := regexp.Compile(validPattern)
	//if err != nil {
	//	fmt.Printf("Error compiling valid regex: %v\n", err)
	//	os.Exit(1)
	//}
	//if !validRegex.MatchString(result) {
	//	fmt.Printf("Result string '%s' does not match valid pattern '%s'\n", result, validPattern)
	//	os.Exit(1)
	//}
	//
	//// 发送结果字符串给 bash 脚本
	//conn, errs := net.Dial("udp", "localhost:12345")
	//if errs != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to connect to server: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close()
	//
	//_, errs = conn.Write([]byte(result))
	//if errs != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to send message: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println("Result sent successfully")
}
