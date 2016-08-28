package main


import (
	"fmt"
    "unsafe"
//	"reflect"
)


type ab struct {
	a bool;
	b int16;
	c string;
};

type Void unsafe.Pointer;

var aa = ab{false, 32, "mengshi"};

func haha(v uintptr) {
    pb := (*int16)(Void(v + unsafe.Offsetof(aa.b)))
    *pb = 999
    //fmt.Println(aa.b)
//	pb := (*int16)(Pointer(uintptr(Pointer(v)) + Offsetof(aa.b)));
//	fmt.Println(*pb);
}

func main() {
/*
    cc = aa;

    //r := cc.(ab);
v := reflect.ValueOf(cc)
//k := v.Type();
//    u := reflect.TypeOf(cc)
//uu := u.Type();
//    fmt.Println(u.Kind());
    //r := cc.(ab);

    
//Offsetof(aa.b)
    pb := (*int16)(Pointer(v.Pointer() + 8))

    fmt.Println(*pb);
*/

//d := Sizeof(&aa);


    haha(uintptr(Void(&aa)));
    fmt.Println(aa.a, aa.b, aa.c)
 //   fmt.Println(aa.b)
 /*
    fmt.Println(aa.a, aa.b, aa.c)
    pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&aa)) + unsafe.Offsetof(aa.b)))
    *pb = 42
    fmt.Println(aa.a, aa.b, aa.c)
    pc := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&aa)) + unsafe.Offsetof(aa.c)))
    *pc = "zhangyue"
    fmt.Println(aa.a, aa.b, aa.c)

    pa := (*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&aa)) + unsafe.Offsetof(aa.a)))
    *pa = true
    fmt.Println(aa.a, aa.b, aa.c)
    */


//    fmt.Println(unsafe.Pointer(&aa));

var a string = "123";
var b = []byte{'1', '2', '3'};
c := string(b)
if a == c {
	fmt.Printf("kkkkkkkkkkkkkkk:%s\n", c);
}

    return;
}
