
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


type DvrRecvHandleFunc func(cycle *DvrCycleType) int;
type DvrRecvChainHandleFunc func(cycle *DvrCycleType) int;
type DvrSendHandleFunc func(cycle *DvrCycleType) int;
type DvrSendChainHandleFunc func(cycle *DvrCycleType) int;


type DvrOsIoType struct {
    Recv       DvrRecvHandleFunc;
    RecvChain  DvrRecvChainHandleFunc;
    Send       DvrSendHandleFunc;
    SendChain  DvrSendChainHandleFunc;
    Flags      uint;
};
