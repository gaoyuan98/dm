/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_933 struct {
	dm_build_934 *list.List
	dm_build_935 *dm_build_987
	dm_build_936 int
}

func Dm_build_937() *Dm_build_933 {
	return &Dm_build_933{
		dm_build_934: list.New(),
		dm_build_936: 0,
	}
}

func (dm_build_939 *Dm_build_933) Dm_build_938() int {
	return dm_build_939.dm_build_936
}

func (dm_build_941 *Dm_build_933) Dm_build_940(dm_build_942 *Dm_build_1011, dm_build_943 int) int {
	var dm_build_944 = 0
	var dm_build_945 = 0
	for dm_build_944 < dm_build_943 && dm_build_941.dm_build_935 != nil {
		dm_build_945 = dm_build_941.dm_build_935.dm_build_995(dm_build_942, dm_build_943-dm_build_944)
		if dm_build_941.dm_build_935.dm_build_990 == 0 {
			dm_build_941.dm_build_977()
		}
		dm_build_944 += dm_build_945
		dm_build_941.dm_build_936 -= dm_build_945
	}
	return dm_build_944
}

func (dm_build_947 *Dm_build_933) Dm_build_946(dm_build_948 []byte, dm_build_949 int, dm_build_950 int) int {
	var dm_build_951 = 0
	var dm_build_952 = 0
	for dm_build_951 < dm_build_950 && dm_build_947.dm_build_935 != nil {
		dm_build_952 = dm_build_947.dm_build_935.dm_build_999(dm_build_948, dm_build_949, dm_build_950-dm_build_951)
		if dm_build_947.dm_build_935.dm_build_990 == 0 {
			dm_build_947.dm_build_977()
		}
		dm_build_951 += dm_build_952
		dm_build_947.dm_build_936 -= dm_build_952
		dm_build_949 += dm_build_952
	}
	return dm_build_951
}

func (dm_build_954 *Dm_build_933) Dm_build_953(dm_build_955 io.Writer, dm_build_956 int) int {
	var dm_build_957 = 0
	var dm_build_958 = 0
	for dm_build_957 < dm_build_956 && dm_build_954.dm_build_935 != nil {
		dm_build_958 = dm_build_954.dm_build_935.dm_build_1004(dm_build_955, dm_build_956-dm_build_957)
		if dm_build_954.dm_build_935.dm_build_990 == 0 {
			dm_build_954.dm_build_977()
		}
		dm_build_957 += dm_build_958
		dm_build_954.dm_build_936 -= dm_build_958
	}
	return dm_build_957
}

func (dm_build_960 *Dm_build_933) Dm_build_959(dm_build_961 []byte, dm_build_962 int, dm_build_963 int) {
	if dm_build_963 == 0 {
		return
	}
	var dm_build_964 = dm_build_991(dm_build_961, dm_build_962, dm_build_963)
	if dm_build_960.dm_build_935 == nil {
		dm_build_960.dm_build_935 = dm_build_964
	} else {
		dm_build_960.dm_build_934.PushBack(dm_build_964)
	}
	dm_build_960.dm_build_936 += dm_build_963
}

func (dm_build_966 *Dm_build_933) dm_build_965(dm_build_967 int) byte {
	var dm_build_968 = dm_build_967
	var dm_build_969 = dm_build_966.dm_build_935
	for dm_build_968 > 0 && dm_build_969 != nil {
		if dm_build_969.dm_build_990 == 0 {
			continue
		}
		if dm_build_968 > dm_build_969.dm_build_990-1 {
			dm_build_968 -= dm_build_969.dm_build_990
			dm_build_969 = dm_build_966.dm_build_934.Front().Value.(*dm_build_987)
		} else {
			break
		}
	}
	return dm_build_969.dm_build_1008(dm_build_968)
}
func (dm_build_971 *Dm_build_933) Dm_build_970(dm_build_972 *Dm_build_933) {
	if dm_build_972.dm_build_936 == 0 {
		return
	}
	var dm_build_973 = dm_build_972.dm_build_935
	for dm_build_973 != nil {
		dm_build_971.dm_build_974(dm_build_973)
		dm_build_972.dm_build_977()
		dm_build_973 = dm_build_972.dm_build_935
	}
	dm_build_972.dm_build_936 = 0
}
func (dm_build_975 *Dm_build_933) dm_build_974(dm_build_976 *dm_build_987) {
	if dm_build_976.dm_build_990 == 0 {
		return
	}
	if dm_build_975.dm_build_935 == nil {
		dm_build_975.dm_build_935 = dm_build_976
	} else {
		dm_build_975.dm_build_934.PushBack(dm_build_976)
	}
	dm_build_975.dm_build_936 += dm_build_976.dm_build_990
}

func (dm_build_978 *Dm_build_933) dm_build_977() {
	var dm_build_979 = dm_build_978.dm_build_934.Front()
	if dm_build_979 == nil {
		dm_build_978.dm_build_935 = nil
	} else {
		dm_build_978.dm_build_935 = dm_build_979.Value.(*dm_build_987)
		dm_build_978.dm_build_934.Remove(dm_build_979)
	}
}

func (dm_build_981 *Dm_build_933) Dm_build_980() []byte {
	var dm_build_982 = make([]byte, dm_build_981.dm_build_936)
	var dm_build_983 = dm_build_981.dm_build_935
	var dm_build_984 = 0
	var dm_build_985 = len(dm_build_982)
	var dm_build_986 = 0
	for dm_build_983 != nil {
		if dm_build_983.dm_build_990 > 0 {
			if dm_build_985 > dm_build_983.dm_build_990 {
				dm_build_986 = dm_build_983.dm_build_990
			} else {
				dm_build_986 = dm_build_985
			}
			copy(dm_build_982[dm_build_984:dm_build_984+dm_build_986], dm_build_983.dm_build_988[dm_build_983.dm_build_989:dm_build_983.dm_build_989+dm_build_986])
			dm_build_984 += dm_build_986
			dm_build_985 -= dm_build_986
		}
		if dm_build_981.dm_build_934.Front() == nil {
			dm_build_983 = nil
		} else {
			dm_build_983 = dm_build_981.dm_build_934.Front().Value.(*dm_build_987)
		}
	}
	return dm_build_982
}

type dm_build_987 struct {
	dm_build_988 []byte
	dm_build_989 int
	dm_build_990 int
}

func dm_build_991(dm_build_992 []byte, dm_build_993 int, dm_build_994 int) *dm_build_987 {
	return &dm_build_987{
		dm_build_992,
		dm_build_993,
		dm_build_994,
	}
}

func (dm_build_996 *dm_build_987) dm_build_995(dm_build_997 *Dm_build_1011, dm_build_998 int) int {
	if dm_build_996.dm_build_990 <= dm_build_998 {
		dm_build_998 = dm_build_996.dm_build_990
	}
	dm_build_997.Dm_build_1094(dm_build_996.dm_build_988[dm_build_996.dm_build_989 : dm_build_996.dm_build_989+dm_build_998])
	dm_build_996.dm_build_989 += dm_build_998
	dm_build_996.dm_build_990 -= dm_build_998
	return dm_build_998
}

func (dm_build_1000 *dm_build_987) dm_build_999(dm_build_1001 []byte, dm_build_1002 int, dm_build_1003 int) int {
	if dm_build_1000.dm_build_990 <= dm_build_1003 {
		dm_build_1003 = dm_build_1000.dm_build_990
	}
	copy(dm_build_1001[dm_build_1002:dm_build_1002+dm_build_1003], dm_build_1000.dm_build_988[dm_build_1000.dm_build_989:dm_build_1000.dm_build_989+dm_build_1003])
	dm_build_1000.dm_build_989 += dm_build_1003
	dm_build_1000.dm_build_990 -= dm_build_1003
	return dm_build_1003
}

func (dm_build_1005 *dm_build_987) dm_build_1004(dm_build_1006 io.Writer, dm_build_1007 int) int {
	if dm_build_1005.dm_build_990 <= dm_build_1007 {
		dm_build_1007 = dm_build_1005.dm_build_990
	}
	dm_build_1006.Write(dm_build_1005.dm_build_988[dm_build_1005.dm_build_989 : dm_build_1005.dm_build_989+dm_build_1007])
	dm_build_1005.dm_build_989 += dm_build_1007
	dm_build_1005.dm_build_990 -= dm_build_1007
	return dm_build_1007
}
func (dm_build_1009 *dm_build_987) dm_build_1008(dm_build_1010 int) byte {
	return dm_build_1009.dm_build_988[dm_build_1009.dm_build_989+dm_build_1010]
}
