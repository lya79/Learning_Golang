package TestNdkBuildStatic

/*
#cgo CPPFLAGS: -I./
#cgo LDFLAGS: -L./ -lHelloTestFromNDKBuild
#include "HelloTestFromNDKBuild.h"
*/
import "C"
import (
	"fmt"
	"log"
)

func HelloFromNdkBuildStatic() {
	iSing99 := C.hi()
	fmt.Println("success from dynamic lib!", iSing99)
	log.Println("success from dynamic lib!", iSing99)
}
