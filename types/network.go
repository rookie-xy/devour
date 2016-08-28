
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


const (
    DVR_NETWORK_MODULE = 0x544E5645;
    DvrNetworkConf     = 0x02000000;
)


type DvrNetworkConfType struct {
    Connections       uint;
    Use               uint;

    MultiAccept       bool;
    AcceptMutex       bool;

    AcceptMutexDelay  int;

    Name              string;
    DebugConnection   DvrArrayType;
};
