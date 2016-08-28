
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


import (
    . "os"
    . "unsafe"
    "fmt"
    "strconv"
)


const (
    DvrConfNoargs     =  0x00000001
    DvrConfTake1      =  0x00000002
    DvrConfTake2      =  0x00000004
    DvrConfTake3      =  0x00000008
    DvrConfTake4      =  0x00000010
    DvrConfTake5      =  0x00000020
    DvrConfTake6      =  0x00000040
    DvrConfTake7      =  0x00000080
    DvrConfTake12     =  (DvrConfTake1|DvrConfTake2)

    DvrConfMore2      =  0x00001000
    DvrConfMore1      =  0x00000800
    DvrConfAny        =  0x00000400

    DvrConfBlock      =  0x00000100
    DvrConfFlag       =  0x00000200

    DvrMainConf       =  0x01000000
    DvrAnyConf        =  0x0F000000

    DvrDirectConf     =  0x00010000
)


const (
    DVR_CORE_MODULE   = 0x45524F43
    DVR_CONF_MODULE   = 0x464E4F43
)


const (
    DvrConfBlockDone = 1
    DvrConfFileDone = 2
)


var (
    DvrNilString  = DvrStrType{0, nil};
    DvrNilCommand = DvrCommandType{ DvrNilString, 0, nil, 0, 0, nil };
)


var (
    DvrCcf DvrCoreConfType;
    DvrNcf = DvrNetworkConfType{};
)


const (
    DvrConfUnsetUint = 0;
    DvrConfUnsetInt = -1;
    DvrConfUnsetBool = false;
)


const (
    DvrDate  = 0;
    DvrGmt   = 1;
    DvrNginx = 2;
)


var  DvrConfOk string = "0";
var  DvrConfError string = "-1";


type DvrConfHandleFunc func(cf *DvrConfType, dummy *DvrCommandType, conf DvrVoidType) string;
type DvrConfParseHandleFunc func(cf *DvrConfType, filename *DvrStrType) string;


type DvrConfType struct {
    Name          string;
    Args         *DvrArrayType;
    Cycle        *DvrCycleType;

    ConfFile     *DvrConfFileType;
    Log          *DvrErrlogType;

    Ctx          []*DvrVoidType;

    ModuleType    uint;
    CommandType   uint;

    Modules       DvrInterfType;

    Handler       DvrConfHandleFunc;
    ParseHandler  DvrConfParseHandleFunc;

    HandlerConf   DvrVoidType;
};


type DvrConfPostHandleFunc func(cf *DvrConfType, data DvrInterfType, conf DvrVoidType) DvrInterfType;

type DvrConfPostType struct {
    Handler DvrConfPostHandleFunc;
};


type DvrConfFileType struct {
    File     DvrFileType;
    Buffer  *DvrBufType;
    Line     uint;
};


type DvrOpenFileType struct {
    *File;
};


func DvrConfSetFlagSlot(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    p := conf;

    fp := (*bool)(Pointer(uintptr(*p) + cmd.Offset));
    /*
    if *fp != DvrConfUnset {
	return "is duplicate";
    }
    */

    value := cf.Args.Elts[1];

    flag := string(value.Data.([]byte))[0:value.Len];

    if flag == "on" {
	*fp = true;

    } else if flag == "off" {
	*fp = false;

    } else {
	//DvrConfLogPrint
	/*
        DvrErrlogPrint(DvrLogEmerg, cf, 0,
                       "invalid value \"%s\" in \"%s\" directive, ",
                       "it must be \"on\" or \"off\"",
                       flag, cmd.Name.Data);
		       */
        return DvrConfError;
    }

    if cmd.Post != nil {
	post := cmd.Post.(DvrConfPostType);
	post.Handler(cf, post, *p);
    }

    return DvrConfOk;
}

func DvrConfSetNumSlot(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    var err error;

    p := conf;

    fp := (*int)(Pointer(uintptr(*p) + cmd.Offset));

    if *fp != DvrConfUnsetInt {
	return "is duplicate";
    }

    value := cf.Args.Elts[1];
    flag := string(value.Data.([]byte))[0:value.Len];

    *fp, err = strconv.Atoi(flag);
    if err != nil {
	return err.Error();
    }

    if cmd.Post != nil {
	post := cmd.Post.(DvrConfPostType);
	post.Handler(cf, post, *p);
    }

    return DvrConfOk;
}


func DvrConfSetStrSlot(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    var field *DvrStrType;

    p := conf;

    field = (*DvrStrType)(Pointer(uintptr(*p) + cmd.Offset));

    if field.Data != nil {
	return "is duplicate";
    }

    value := cf.Args.Elts[1];
    fmt.Println(value.Len, string(value.Data.([]byte)));
    field.Len = value.Len;
    field.Data = value.Data;


    //fmt.Println(field.Len, string(field.Data.([]byte)));

    if cmd.Post != nil {
	post := cmd.Post.(DvrConfPostType);
	post.Handler(cf, post, *p);
    }

    return DvrConfOk;
}


func DvrConfSetMsecSlot(cf *DvrConfType, cmd *DvrCommandType, conf *DvrVoidType) string {
    return DvrConfOk;
}

func DvrSizeof(v string) int {
    return len(v) + 1;
}
