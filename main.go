package main

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
#include <errno.h>
long  lens = 256;
char * compile_regex_and_match( char *pattern,  char *subject) {
	pcre2_code *re;
	PCRE2_SIZE erroroffset;
	PCRE2_SIZE *ovector;
	int errornumber;
	int rc;

	//PCRE2_SPTR p = (PCRE2_SPTR)pattern;
	//PCRE2_SPTR s = (PCRE2_SPTR)subject;
	PCRE2_SIZE len = strlen((char *)subject);


	size_t subject_length;
	pcre2_match_data *match_data;
	printf("tttttttttt1\n");
   	re = pcre2_compile(
	  	(PCRE2_SPTR)pattern,
		PCRE2_ZERO_TERMINATED,
		0,
		&errornumber,
		&erroroffset,
		NULL);

	printf("tttttttttt2\n");
	if (re == NULL)
	{
		//PCRE2_UCHAR buffer[256];
		//pcre2_get_error_message(errornumber, buffer, sizeof(buffer));
		printf("PCRE2 compilation failed at offset %d\n", (int)erroroffset);
		//struct Result r = { -1, "PCRE2 compilation failed" };
		//return &r;
		errno = EINVAL;
        return NULL;
	}

	printf("tttttttttt3\n");
	match_data = pcre2_match_data_create_from_pattern(re, NULL);
	printf("tttttttttt4\n");
	rc = pcre2_match(
		re,
		(PCRE2_SPTR)subject,
		subject_length,
		0,
		0,
		match_data,
		NULL);

	printf("tttttttttt5\n");
	if (rc < 0)
	  {
	  	switch(rc)
			{
			case PCRE2_ERROR_NOMATCH: printf("No match\n"); break;
			default: printf("Matching error %d\n", rc); break;
			}
		pcre2_match_data_free(match_data);
		pcre2_code_free(re);
		//struct Result r = { -1, "PCRE2 match failed" };
		//return &r;

		errno = EINVAL;
        return NULL;
	}
	printf("tttttttttt6\n");
	ovector = pcre2_get_ovector_pointer(match_data);
	printf("\nMatch succeeded at offset %d\n", (int)ovector[0]);
	if (rc == 0)
  		printf("ovector was not big enough for all the captured substrings\n");
	pcre2_match_data_free(match_data);
	pcre2_code_free(re);
	PCRE2_SPTR substring_start = (PCRE2_SPTR)subject + ovector[0];
	size_t substring_length = ovector[1] - ovector[0];

	printf("---%2d: %.*s\n", 0, (int)substring_length, (char *)substring_start);

	//struct Result r = { 0, (char *)substring_start };
	//return &r;
	return (char *)substring_start;
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

	// 匹配目标文本
	subject := C.CString(target)
	defer C.free(unsafe.Pointer(subject))

	reC, cErr := C.compile_regex_and_match(cPattern, subject)
	if cErr != nil {
		os.Exit(1)
	}
	fmt.Println(reC)
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
