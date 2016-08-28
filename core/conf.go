
/*
 * Copyright (C) 2016 Meng Shi
 */


package core


import (
	"fmt"
      "os"
      "strings"

    . "devour/types"
    . "devour/autoconf"
)


var DvrArgumentNumber = []uint{
    DvrConfNoargs,
    DvrConfTake1,
    DvrConfTake2,
    DvrConfTake3,
    DvrConfTake4,
    DvrConfTake5,
    DvrConfTake6,
    DvrConfTake7,
};


func DvrConfFullName(cycle *DvrCycleType, name *DvrStrType) int {
    if name.Data.(string)[0:1] == "/" {
        return DvrOk;
    }

    return DvrOk;
}

//DvrCoreConfType 
func DvrGetConf(confCtx []*DvrVoidType, module DvrMoudleType) *DvrVoidType {
	fmt.Printf("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx:%d\n", module.Index);
    return confCtx[module.Index];
    //return nil;
}


func DvrConfReadToken(cf *DvrConfType) int {
    var src, dst                                    []byte;
    var ch                                          byte;
    var l, start                                    int;
    var found, needSpace, lastSpace, sharpComment   bool;
    var quoted, Squoted, Dquoted                    bool;
    var n                                           int;
    var word                                       *DvrStrType;
    var b                                          *DvrBufType;

    found = false;

    needSpace = false;
    lastSpace = true;
    sharpComment = false;

    quoted, Squoted, Dquoted = false, false, false;

    src, dst = nil, nil;
    cf.Args.Nelts = 0;
    b = cf.ConfFile.Buffer;
    start = b.Pos;

    for ;; {
        if b.Pos >= b.Last {
	    if cf.ConfFile.File.Offset >= DvrFileSize(cf.ConfFile.File.Stat) {
                return DvrConfFileDone;
	    }

	    n = DvrReadFile(&cf.ConfFile.File, b.Buf[b.Start + (b.Pos - start):],
	                                      b.End - (b.Start + (b.Pos - start)),
					      cf.ConfFile.File.Offset);

	    if n == DvrError {
	        return DvrError;
	    }

	    b.Buf = b.Buf[b.Start + (b.Pos - start):];
	    start = b.Start;
	    b.Last = b.Pos + n;
	}

	ch = b.Buf[b.Pos]; b.Pos++;

	if ch == LF {
	    cf.ConfFile.Line++;

	    if sharpComment {
                sharpComment = false;
	    }
	}

        if sharpComment {
	    continue;
        }

	if quoted {
            quoted = false;
	    continue;
	}

	if needSpace {
            if ch == ' ' || ch == '\t' || ch == CR || ch == LF {
                lastSpace = true;
                needSpace = false;
                continue;
            }

            if ch == ';' || ch == '{' {
                return DvrOk;
            }

            DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                           "unexpected '%c' in %s:%d",
                           ch, cf.ConfFile.File.Name.Data,
                           cf.ConfFile.Line);

            return DvrError;
        }

	if lastSpace {
            if ch == ' ' || ch == '\t' || ch == CR || ch == LF {
                continue;
            }

            start = b.Pos - 1;

            switch ch {

            case ';':
            case '{':
                if cf.Args.Nelts == 0 {
                    DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                                   "unexpected '%c' in %s:%d",
                                   ch, cf.ConfFile.File.Name.Data,
                                   cf.ConfFile.Line);
                    return DvrError;
                }

                return DvrOk;

            case '}':
                if cf.Args.Nelts > 0 {
                    DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                                   "unexpected '}' in %s:%d",
                                   ch, cf.ConfFile.File.Name.Data,
                                   cf.ConfFile.Line);
                    return DvrError;
                }

                return DvrConfBlockDone;

            case '#':
                sharpComment = true;
                continue;

            case '\\':
                quoted = true;
                lastSpace = false;
                continue;

            case '"':
                start++;
                Dquoted = true;
                lastSpace = false;
                continue;

            case '\'':
                start++;
                Squoted = true;
                lastSpace = false;
                continue;

            default:
                lastSpace = false;
            }

	} else {
            if ch == '\\' {
                quoted = true;
                continue;
            }

            if Dquoted {
                if ch == '"' {
                    Dquoted = false;
                    needSpace = true;
                    found = true;
                }

            } else if Squoted {
                if ch == '\'' {
                    Dquoted = false;
                    needSpace = true;
                    found = true;
                }

            } else if ch == ' ' || ch == '\t' || ch == CR || ch == LF || ch == ';' || ch == '{' {
                lastSpace = true;
                found = true;
            }

            if found {
                word = &DvrStrType{};
		dst = make([]byte, b.Pos);
		src = b.Buf;
                i := start;
                j := 0;

		for l = 0; i < b.Pos - 1; l++ {
                    if src[i] == '\\' {
                        switch src[1] {
                        case '"':
                        case '\'':
                        case '\\':
                            i++;
                            break;

                        case 't':
			    j++;
                            dst[j] = '\t';
                            i += 2
                            continue;

                        case 'r':
                            j++;
                            dst[j] = '\r';
			    i += 2;
                            continue;

                        case 'n':
                            j++;
                            dst[j] = '\n';
			    i += 2;
                            continue;
                        }
                    }

                    dst[j] = src[i];
                    i++; j++;
                }

		word.Data = dst;
                word.Len = l;

                cf.Args.Elts[cf.Args.Nelts] = word;
		cf.Args.Nelts = cf.Args.Nelts + 1;

                if ch == ';' || ch == '{' {
                    return DvrOk;
                }

                found = false;
            }
        }
    }
}


func DvrConfParse(cf *DvrConfType, filename *DvrStrType) string {
    var rc, m, i       int;
    var found, valid   bool;
    //var rv             DvrVoidType;
//    var cmd            DvrCommandType;
    var conf          *DvrVoidType;
    var name          *DvrStrType;
    var prev          *DvrConfFileType;

    if filename != nil {
	fd, err := os.OpenFile(filename.Data.(string), DvrFileFlag, DvrFileOpenMode)
	if err != nil {
            DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                           DvrOpenFileName + " %s failed", filename.Data);
            return DvrConfError;
	}

	prev = cf.ConfFile;

	cf.ConfFile = &DvrConfFileType{};
	cf.ConfFile.File.Stat = &DvrFileStatType{};
	err = DvrFdStat(fd, cf.ConfFile.File.Stat)
	if err != nil {
            DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                           DvrFdStatName + " %s failed", filename.Data);
	    return DvrConfError;
	}

        cf.ConfFile.Buffer = DvrCreateBuffer(1024);
	if cf.ConfFile.Buffer == nil {
            return DvrConfError;
	}

	cf.ConfFile.File.Fd = fd;
        cf.ConfFile.File.Name = &DvrNilString;
	cf.ConfFile.File.Name.Len = filename.Len;
        cf.ConfFile.File.Name.Data = filename.Data;
	cf.ConfFile.File.Offset = 0;
        cf.ConfFile.File.Log = cf.Log;
        cf.ConfFile.Line = 1;
    }

    for ;; {
	rc = DvrConfReadToken(cf);

	if rc == DvrError {
	    break;
	}

	if rc != DvrOk {
	    break;
	}

	if cf.Handler != nil {
	    rv := cf.Handler(cf, nil, cf.HandlerConf);

	    if (rv == DvrConfOk) {
                continue;

            } else if (rv == DvrConfError) {
                 rc = DvrError;
                 break;

            } else {
                 DvrErrlogPrint(DvrLogEmerg, cf.Log, 0, "%s in %s:%d", rv,
                                cf.ConfFile.File.Name.Data,
                                cf.ConfFile.Line);
                 rc = DvrError;
                 break;
            }
        }

	name = cf.Args.Elts[0];
	found = false;

	for m = 0; rc != DvrError && !found && DvrModules[m] != nil; m++ {
	    if DvrModules[m].Type != DVR_CONF_MODULE &&
	       DvrModules[m].Type != cf.ModuleType {
		continue;
	    }

	    commands := DvrModules[m].Commands;
            if commands == nil {
		continue;
	    }

	    for i = 0; commands[i].Name.Len != 0; i++ {
                cmd := commands[i]
		if name.Len == cmd.Name.Len &&
	           strings.Contains(string(name.Data.([]byte)), cmd.Name.Data.(string)) {

		    found = true;

		    if cmd.Type & cf.CommandType == 0 {
                        DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                                       "directive \"%s\" in %s:%d is not allowed here",
                                       name.Data,
                                       cf.ConfFile.File.Name.Data,
                                       cf.ConfFile.Line);
                        rc = DvrError;
                        break;
		    }

		    if cmd.Type & DvrConfAny != 0 {
			valid = true;


		    } else if cmd.Type & DvrConfFlag != 0 {
			if cf.Args.Nelts == 2 {
                            valid = true;
		        } else {
                            valid = false;
			}

		    } else if cmd.Type & DvrConfMore1 != 0 {
			if cf.Args.Nelts > 1 {
                            valid = true;
		        } else {
                            valid = false;
			}

		    } else if cmd.Type & DvrConfMore2 != 0 {
			if cf.Args.Nelts > 2 {
                            valid = true;
		        } else {
                            valid = false;
			}

		    } else if cf.Args.Nelts <= 10 &&
		              (cmd.Type & DvrArgumentNumber[cf.Args.Nelts - 1] != 0) {
			valid = true;

		    } else {
                        valid = false;
		    }

		    if !valid {
	                DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
			               "invalid number arguments in directive \"%s\" in %s:%d",
				       name.Data,
                                       cf.ConfFile.File.Name.Data,
                                       cf.ConfFile.Line);
                        rc = DvrError;
                        break;
		    }

		    conf = nil;

		    if cmd.Type & DvrDirectConf != 0 {
                        conf = cf.Ctx[DvrModules[m].Index];

		    } else if cmd.Type & DvrMainConf != 0 {
                        conf = cf.Ctx[DvrModules[m].Index];

		    } else if cf.Ctx != nil {
                        conf = cf.Ctx[DvrModules[m].CtxIndex];
		    }

		    rv := cmd.Set(cf, &cmd, conf);

		    if rv == DvrConfOk {
			break;

		    } else if rv == DvrConfError {
			rc = DvrError;
			break;

		    } else {
                        DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                                       "the \"%s\" directive %s in %s:%d",
                                       name.Data, rv,
                                       cf.ConfFile.File.Name.Data,
                                       cf.ConfFile.Line);
			rc = DvrError;
			break;
		    }
		}
	    }
	}

	if !found {
            DvrErrlogPrint(DvrLogEmerg, cf.Log, 0,
                           "unknown directive \"%s\" in %s:%d",
                           name.Data,
                           cf.ConfFile.File.Name.Data,
                           cf.ConfFile.Line);
	    rc = DvrError;
	    break;
	}

	if rc == DvrError {
	    break;
	}
    }

    if filename != nil {
	cf.ConfFile = prev;

        if (DvrCloseFile() == DvrFileError) {
            DvrErrlogPrint(DvrLogAlert, cf.Log, 0,
                           DvrCloseFileName + " %s failed",
                           cf.ConfFile.File.Name.Data);
            return DvrConfError;
        }
    }

    if rc == DvrError {
	return DvrConfError;
    }

    return DvrConfOk;
}
