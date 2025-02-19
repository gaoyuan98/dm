/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gaoyuan98/dm/security"
	"net"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	Dm_build_694 = 8192
	Dm_build_695 = 2 * time.Second
)

type dm_build_696 struct {
	dm_build_697 net.Conn
	dm_build_698 *tls.Conn
	dm_build_699 *Dm_build_360
	dm_build_700 *DmConnection
	dm_build_701 security.Cipher
	dm_build_702 bool
	dm_build_703 bool
	dm_build_704 *security.DhKey

	dm_build_705 bool
	dm_build_706 string
	dm_build_707 bool
}

func dm_build_708(dm_build_709 context.Context, dm_build_710 *DmConnection) (*dm_build_696, error) {
	var dm_build_711 net.Conn
	var dm_build_712 error

	dialsLock.RLock()
	dm_build_713, dm_build_714 := dials[dm_build_710.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_714 {
		dm_build_711, dm_build_712 = dm_build_713(dm_build_709, dm_build_710.dmConnector.host+":"+strconv.Itoa(int(dm_build_710.dmConnector.port)))
	} else {
		dm_build_711, dm_build_712 = dm_build_716(dm_build_710.dmConnector.host+":"+strconv.Itoa(int(dm_build_710.dmConnector.port)), time.Duration(dm_build_710.dmConnector.socketTimeout)*time.Second)
	}
	if dm_build_712 != nil {
		return nil, dm_build_712
	}

	dm_build_715 := dm_build_696{}
	dm_build_715.dm_build_697 = dm_build_711
	dm_build_715.dm_build_699 = Dm_build_363(Dm_build_989)
	dm_build_715.dm_build_700 = dm_build_710
	dm_build_715.dm_build_702 = false
	dm_build_715.dm_build_703 = false
	dm_build_715.dm_build_705 = false
	dm_build_715.dm_build_706 = ""
	dm_build_715.dm_build_707 = false
	dm_build_710.Access = &dm_build_715

	return &dm_build_715, nil
}

func dm_build_716(dm_build_717 string, dm_build_718 time.Duration) (net.Conn, error) {
	dm_build_719, dm_build_720 := net.DialTimeout("tcp", dm_build_717, dm_build_718)
	if dm_build_720 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_717).throw()
	}

	if tcpConn, ok := dm_build_719.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_695)
		tcpConn.SetNoDelay(true)

	}
	return dm_build_719, nil
}

func (dm_build_722 *dm_build_696) dm_build_721(dm_build_723 dm_build_1110) bool {
	var dm_build_724 = dm_build_722.dm_build_700.dmConnector.compress
	if dm_build_723.dm_build_1125() == Dm_build_1017 || dm_build_724 == Dm_build_1066 {
		return false
	}

	if dm_build_724 == Dm_build_1064 {
		return true
	} else if dm_build_724 == Dm_build_1065 {
		return !dm_build_722.dm_build_700.Local && dm_build_723.dm_build_1123() > Dm_build_1063
	}

	return false
}

func (dm_build_726 *dm_build_696) dm_build_725(dm_build_727 dm_build_1110) bool {
	var dm_build_728 = dm_build_726.dm_build_700.dmConnector.compress
	if dm_build_727.dm_build_1125() == Dm_build_1017 || dm_build_728 == Dm_build_1066 {
		return false
	}

	if dm_build_728 == Dm_build_1064 {
		return true
	} else if dm_build_728 == Dm_build_1065 {
		return dm_build_726.dm_build_699.Dm_build_627(Dm_build_1025) == 1
	}

	return false
}

func (dm_build_730 *dm_build_696) dm_build_729(dm_build_731 dm_build_1110) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_733 := dm_build_731.dm_build_1123()

	if dm_build_733 > 0 {

		if dm_build_730.dm_build_721(dm_build_731) {
			var retBytes, err = Compress(dm_build_730.dm_build_699, Dm_build_1018, int(dm_build_733), int(dm_build_730.dm_build_700.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_730.dm_build_699.Dm_build_374(Dm_build_1018)

			dm_build_730.dm_build_699.Dm_build_415(dm_build_733)

			dm_build_730.dm_build_699.Dm_build_443(retBytes)

			dm_build_731.dm_build_1124(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_730.dm_build_699.Dm_build_547(Dm_build_1025, 1)
		}

		if dm_build_730.dm_build_703 {
			dm_build_733 = dm_build_731.dm_build_1123()
			var retBytes = dm_build_730.dm_build_701.Encrypt(dm_build_730.dm_build_699.Dm_build_654(Dm_build_1018, int(dm_build_733)), true)

			dm_build_730.dm_build_699.Dm_build_374(Dm_build_1018)

			dm_build_730.dm_build_699.Dm_build_443(retBytes)

			dm_build_731.dm_build_1124(int32(len(retBytes)))
		}
	}

	if dm_build_730.dm_build_699.Dm_build_372() > Dm_build_990 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_731.dm_build_1119()
	if dm_build_730.dm_build_972(dm_build_731) {
		if dm_build_730.dm_build_698 != nil {
			dm_build_730.dm_build_699.Dm_build_377(0)
			if _, err := dm_build_730.dm_build_699.Dm_build_396(dm_build_730.dm_build_698); err != nil {
				return err
			}
		}
	} else {
		dm_build_730.dm_build_699.Dm_build_377(0)
		if _, err := dm_build_730.dm_build_699.Dm_build_396(dm_build_730.dm_build_697); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_735 *dm_build_696) dm_build_734(dm_build_736 dm_build_1110) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_738 := int32(0)
	if dm_build_735.dm_build_972(dm_build_736) {
		if dm_build_735.dm_build_698 != nil {
			dm_build_735.dm_build_699.Dm_build_374(0)
			if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_698, Dm_build_1018); err != nil {
				return err
			}

			dm_build_738 = dm_build_736.dm_build_1123()
			if dm_build_738 > 0 {
				if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_698, int(dm_build_738)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_735.dm_build_699.Dm_build_374(0)
		if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_697, Dm_build_1018); err != nil {
			return err
		}
		dm_build_738 = dm_build_736.dm_build_1123()

		if dm_build_738 > 0 {
			if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_697, int(dm_build_738)); err != nil {
				return err
			}
		}
	}

	dm_build_736.dm_build_1120()

	dm_build_738 = dm_build_736.dm_build_1123()
	if dm_build_738 <= 0 {
		return nil
	}

	if dm_build_735.dm_build_703 {
		ebytes := dm_build_735.dm_build_699.Dm_build_654(Dm_build_1018, int(dm_build_738))
		bytes, err := dm_build_735.dm_build_701.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_735.dm_build_699.Dm_build_374(Dm_build_1018)
		dm_build_735.dm_build_699.Dm_build_443(bytes)
		dm_build_736.dm_build_1124(int32(len(bytes)))
	}

	if dm_build_735.dm_build_725(dm_build_736) {

		dm_build_738 = dm_build_736.dm_build_1123()
		cbytes := dm_build_735.dm_build_699.Dm_build_654(Dm_build_1018+ULINT_SIZE, int(dm_build_738-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_735.dm_build_700.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_735.dm_build_699.Dm_build_374(Dm_build_1018)
		dm_build_735.dm_build_699.Dm_build_443(bytes)
		dm_build_736.dm_build_1124(int32(len(bytes)))
	}
	return nil
}

func (dm_build_740 *dm_build_696) dm_build_739(dm_build_741 dm_build_1110) (dm_build_742 interface{}, dm_build_743 error) {
	if dm_build_740.dm_build_707 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_744 := dm_build_740.dm_build_700
	dm_build_744.mu.Lock()
	defer dm_build_744.mu.Unlock()
	dm_build_743 = dm_build_741.dm_build_1114(dm_build_741)
	if dm_build_743 != nil {
		return nil, dm_build_743
	}

	dm_build_743 = dm_build_740.dm_build_729(dm_build_741)
	if dm_build_743 != nil {
		return nil, dm_build_743
	}

	dm_build_743 = dm_build_740.dm_build_734(dm_build_741)
	if dm_build_743 != nil {
		return nil, dm_build_743
	}

	return dm_build_741.dm_build_1118(dm_build_741)
}

func (dm_build_746 *dm_build_696) dm_build_745() (*dm_build_1567, error) {

	Dm_build_747 := dm_build_1573(dm_build_746)
	_, dm_build_748 := dm_build_746.dm_build_739(Dm_build_747)
	if dm_build_748 != nil {
		return nil, dm_build_748
	}

	return Dm_build_747, nil
}

func (dm_build_750 *dm_build_696) dm_build_749() error {

	dm_build_751 := dm_build_1434(dm_build_750)
	_, dm_build_752 := dm_build_750.dm_build_739(dm_build_751)
	if dm_build_752 != nil {
		return dm_build_752
	}

	return nil
}

func (dm_build_754 *dm_build_696) dm_build_753() error {

	var dm_build_755 *dm_build_1567
	var err error
	if dm_build_755, err = dm_build_754.dm_build_745(); err != nil {
		return err
	}

	if dm_build_754.dm_build_700.sslEncrypt == 2 {
		if err = dm_build_754.dm_build_968(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_754.dm_build_700.sslEncrypt == 1 {
		if err = dm_build_754.dm_build_968(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_754.dm_build_703 || dm_build_754.dm_build_702 {
		k, err := dm_build_754.dm_build_958()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_755.Dm_build_1571)
		encryptType := dm_build_755.dm_build_1569
		hashType := int(dm_build_755.Dm_build_1570)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_754.dm_build_961(encryptType, sessionKey, dm_build_754.dm_build_700.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_754.dm_build_749(); err != nil {
		return err
	}
	return nil
}

func (dm_build_758 *dm_build_696) Dm_build_757(dm_build_759 *DmStatement) error {
	dm_build_760 := dm_build_1596(dm_build_758, dm_build_759)
	_, dm_build_761 := dm_build_758.dm_build_739(dm_build_760)
	if dm_build_761 != nil {
		return dm_build_761
	}

	return nil
}

func (dm_build_763 *dm_build_696) Dm_build_762(dm_build_764 int32) error {
	dm_build_765 := dm_build_1606(dm_build_763, dm_build_764)
	_, dm_build_766 := dm_build_763.dm_build_739(dm_build_765)
	if dm_build_766 != nil {
		return dm_build_766
	}

	return nil
}

func (dm_build_768 *dm_build_696) Dm_build_767(dm_build_769 *DmStatement, dm_build_770 bool, dm_build_771 int16) (*execRetInfo, error) {
	dm_build_772 := dm_build_1473(dm_build_768, dm_build_769, dm_build_770, dm_build_771)
	dm_build_773, dm_build_774 := dm_build_768.dm_build_739(dm_build_772)
	if dm_build_774 != nil {
		return nil, dm_build_774
	}
	return dm_build_773.(*execRetInfo), nil
}

func (dm_build_776 *dm_build_696) Dm_build_775(dm_build_777 *DmStatement, dm_build_778 int16) (*execRetInfo, error) {
	return dm_build_776.Dm_build_767(dm_build_777, false, Dm_build_1070)
}

func (dm_build_780 *dm_build_696) Dm_build_779(dm_build_781 *DmStatement, dm_build_782 []OptParameter) (*execRetInfo, error) {
	dm_build_783, dm_build_784 := dm_build_780.dm_build_739(dm_build_1213(dm_build_780, dm_build_781, dm_build_782))
	if dm_build_784 != nil {
		return nil, dm_build_784
	}

	return dm_build_783.(*execRetInfo), nil
}

func (dm_build_786 *dm_build_696) Dm_build_785(dm_build_787 *DmStatement, dm_build_788 int16) (*execRetInfo, error) {
	return dm_build_786.Dm_build_767(dm_build_787, true, dm_build_788)
}

func (dm_build_790 *dm_build_696) Dm_build_789(dm_build_791 *DmStatement, dm_build_792 [][]interface{}) (*execRetInfo, error) {
	dm_build_793 := dm_build_1245(dm_build_790, dm_build_791, dm_build_792)
	dm_build_794, dm_build_795 := dm_build_790.dm_build_739(dm_build_793)
	if dm_build_795 != nil {
		return nil, dm_build_795
	}
	return dm_build_794.(*execRetInfo), nil
}

func (dm_build_797 *dm_build_696) Dm_build_796(dm_build_798 *DmStatement, dm_build_799 [][]interface{}, dm_build_800 bool) (*execRetInfo, error) {
	var dm_build_801, dm_build_802 = 0, 0
	var dm_build_803 = len(dm_build_799)
	var dm_build_804 [][]interface{}
	var dm_build_805 = NewExceInfo()
	dm_build_805.updateCounts = make([]int64, dm_build_803)
	var dm_build_806 = false
	for dm_build_801 < dm_build_803 {
		for dm_build_802 = dm_build_801; dm_build_802 < dm_build_803; dm_build_802++ {
			paramData := dm_build_799[dm_build_802]
			bindData := make([]interface{}, dm_build_798.paramCount)
			dm_build_806 = false
			for icol := 0; icol < int(dm_build_798.paramCount); icol++ {
				if dm_build_798.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_797.dm_build_941(bindData, paramData, icol) {
					dm_build_806 = true
					break
				}
			}

			if dm_build_806 {
				break
			}
			dm_build_804 = append(dm_build_804, bindData)
		}

		if dm_build_802 != dm_build_801 {
			tmpExecInfo, err := dm_build_797.Dm_build_789(dm_build_798, dm_build_804)
			if err != nil {
				return nil, err
			}
			dm_build_804 = dm_build_804[0:0]
			dm_build_805.union(tmpExecInfo, dm_build_801, dm_build_802-dm_build_801)
		}

		if dm_build_802 < dm_build_803 {
			tmpExecInfo, err := dm_build_797.Dm_build_815(dm_build_798, dm_build_799[dm_build_802], dm_build_800)
			if err != nil {
				return nil, err
			}

			dm_build_800 = true
			dm_build_805.union(tmpExecInfo, dm_build_802, 1)
		}

		dm_build_801 = dm_build_802 + 1
	}
	for _, i := range dm_build_805.updateCounts {
		if i > 0 {
			dm_build_805.updateCount += i
		}
	}
	return dm_build_805, nil
}

func (dm_build_808 *dm_build_696) dm_build_807(dm_build_809 *DmStatement, dm_build_810 []parameter) error {
	if !dm_build_809.prepared {
		retInfo, err := dm_build_808.Dm_build_767(dm_build_809, false, Dm_build_1070)
		if err != nil {
			return nil
		}
		dm_build_809.serverParams = retInfo.serverParams
		dm_build_809.paramCount = int32(len(dm_build_809.serverParams))
		dm_build_809.prepared = true
	}

	dm_build_811 := dm_build_1462(dm_build_808, dm_build_809, dm_build_809.bindParams)
	dm_build_812, err := dm_build_808.dm_build_739(dm_build_811)
	if err != nil {
		return nil
	}
	retInfo := dm_build_812.(*execRetInfo)
	if retInfo.serverParams != nil && len(retInfo.serverParams) > 0 {
		dm_build_809.serverParams = retInfo.serverParams
		dm_build_809.paramCount = int32(len(dm_build_809.serverParams))
	}
	dm_build_809.preExec = true
	return nil
}

func (dm_build_816 *dm_build_696) Dm_build_815(dm_build_817 *DmStatement, dm_build_818 []interface{}, dm_build_819 bool) (*execRetInfo, error) {

	var dm_build_820 = make([]interface{}, dm_build_817.paramCount)
	for icol := 0; icol < int(dm_build_817.paramCount); icol++ {
		if dm_build_817.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_816.dm_build_941(dm_build_820, dm_build_818, icol) {

			if !dm_build_819 {
				dm_build_816.dm_build_807(dm_build_817, dm_build_817.bindParams)

				dm_build_819 = true
			}

			dm_build_816.dm_build_947(dm_build_817, dm_build_817.bindParams[icol], icol, dm_build_818[icol].(iOffRowBinder))
			dm_build_820[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_821 = make([][]interface{}, 1, 1)
	dm_build_821[0] = dm_build_820

	dm_build_822 := dm_build_1245(dm_build_816, dm_build_817, dm_build_821)
	dm_build_823, dm_build_824 := dm_build_816.dm_build_739(dm_build_822)
	if dm_build_824 != nil {
		return nil, dm_build_824
	}
	return dm_build_823.(*execRetInfo), nil
}

func (dm_build_826 *dm_build_696) Dm_build_825(dm_build_827 *DmStatement, dm_build_828 int16) (*execRetInfo, error) {
	dm_build_829 := dm_build_1449(dm_build_826, dm_build_827, dm_build_828)

	dm_build_830, dm_build_831 := dm_build_826.dm_build_739(dm_build_829)
	if dm_build_831 != nil {
		return nil, dm_build_831
	}
	return dm_build_830.(*execRetInfo), nil
}

func (dm_build_833 *dm_build_696) Dm_build_832(dm_build_834 *innerRows, dm_build_835 int64) (*execRetInfo, error) {
	dm_build_836 := dm_build_1352(dm_build_833, dm_build_834, dm_build_835, INT64_MAX)
	dm_build_837, dm_build_838 := dm_build_833.dm_build_739(dm_build_836)
	if dm_build_838 != nil {
		return nil, dm_build_838
	}
	return dm_build_837.(*execRetInfo), nil
}

func (dm_build_840 *dm_build_696) Commit() error {
	dm_build_841 := dm_build_1198(dm_build_840)
	_, dm_build_842 := dm_build_840.dm_build_739(dm_build_841)
	if dm_build_842 != nil {
		return dm_build_842
	}

	return nil
}

func (dm_build_844 *dm_build_696) Rollback() error {
	dm_build_845 := dm_build_1511(dm_build_844)
	_, dm_build_846 := dm_build_844.dm_build_739(dm_build_845)
	if dm_build_846 != nil {
		return dm_build_846
	}

	return nil
}

func (dm_build_848 *dm_build_696) Dm_build_847(dm_build_849 *DmConnection) error {
	dm_build_850 := dm_build_1516(dm_build_848, dm_build_849.IsoLevel)
	_, dm_build_851 := dm_build_848.dm_build_739(dm_build_850)
	if dm_build_851 != nil {
		return dm_build_851
	}

	return nil
}

func (dm_build_853 *dm_build_696) Dm_build_852(dm_build_854 *DmStatement, dm_build_855 string) error {
	dm_build_856 := dm_build_1203(dm_build_853, dm_build_854, dm_build_855)
	_, dm_build_857 := dm_build_853.dm_build_739(dm_build_856)
	if dm_build_857 != nil {
		return dm_build_857
	}

	return nil
}

func (dm_build_859 *dm_build_696) Dm_build_858(dm_build_860 []uint32) ([]int64, error) {
	dm_build_861 := dm_build_1614(dm_build_859, dm_build_860)
	dm_build_862, dm_build_863 := dm_build_859.dm_build_739(dm_build_861)
	if dm_build_863 != nil {
		return nil, dm_build_863
	}
	return dm_build_862.([]int64), nil
}

func (dm_build_865 *dm_build_696) Close() error {
	if dm_build_865.dm_build_707 {
		return nil
	}

	dm_build_866 := dm_build_865.dm_build_697.Close()
	if dm_build_866 != nil {
		return dm_build_866
	}

	dm_build_865.dm_build_700 = nil
	dm_build_865.dm_build_707 = true
	return nil
}

func (dm_build_868 *dm_build_696) dm_build_867(dm_build_869 *lob) (int64, error) {
	dm_build_870 := dm_build_1385(dm_build_868, dm_build_869)
	dm_build_871, dm_build_872 := dm_build_868.dm_build_739(dm_build_870)
	if dm_build_872 != nil {
		return 0, dm_build_872
	}
	return dm_build_871.(int64), nil
}

func (dm_build_874 *dm_build_696) dm_build_873(dm_build_875 *lob, dm_build_876 int32, dm_build_877 int32) (*lobRetInfo, error) {
	dm_build_878 := dm_build_1370(dm_build_874, dm_build_875, int(dm_build_876), int(dm_build_877))
	dm_build_879, dm_build_880 := dm_build_874.dm_build_739(dm_build_878)
	if dm_build_880 != nil {
		return nil, dm_build_880
	}
	return dm_build_879.(*lobRetInfo), nil
}

func (dm_build_882 *dm_build_696) dm_build_881(dm_build_883 *DmBlob, dm_build_884 int32, dm_build_885 int32) ([]byte, error) {
	var dm_build_886 = make([]byte, dm_build_885)
	var dm_build_887 int32 = 0
	var dm_build_888 int32 = 0
	var dm_build_889 *lobRetInfo
	var dm_build_890 []byte
	var dm_build_891 error
	for dm_build_887 < dm_build_885 {
		dm_build_888 = dm_build_885 - dm_build_887
		if dm_build_888 > Dm_build_1103 {
			dm_build_888 = Dm_build_1103
		}
		dm_build_889, dm_build_891 = dm_build_882.dm_build_873(&dm_build_883.lob, dm_build_884+dm_build_887, dm_build_888)
		if dm_build_891 != nil {
			return nil, dm_build_891
		}
		dm_build_890 = dm_build_889.data
		if dm_build_890 == nil || len(dm_build_890) == 0 {
			break
		}
		Dm_build_1.Dm_build_57(dm_build_886, int(dm_build_887), dm_build_890, 0, len(dm_build_890))
		dm_build_887 += int32(len(dm_build_890))
		if dm_build_883.readOver {
			break
		}
	}
	return dm_build_886, nil
}

func (dm_build_893 *dm_build_696) dm_build_892(dm_build_894 *DmClob, dm_build_895 int32, dm_build_896 int32) (string, error) {
	var dm_build_897 bytes.Buffer
	var dm_build_898 int32 = 0
	var dm_build_899 int32 = 0
	var dm_build_900 *lobRetInfo
	var dm_build_901 []byte
	var dm_build_902 string
	var dm_build_903 error
	for dm_build_898 < dm_build_896 {
		dm_build_899 = dm_build_896 - dm_build_898
		if dm_build_899 > Dm_build_1103/2 {
			dm_build_899 = Dm_build_1103 / 2
		}
		dm_build_900, dm_build_903 = dm_build_893.dm_build_873(&dm_build_894.lob, dm_build_895+dm_build_898, dm_build_899)
		if dm_build_903 != nil {
			return "", dm_build_903
		}
		dm_build_901 = dm_build_900.data
		if dm_build_901 == nil || len(dm_build_901) == 0 {
			break
		}
		dm_build_902 = Dm_build_1.Dm_build_158(dm_build_901, 0, len(dm_build_901), dm_build_894.serverEncoding, dm_build_893.dm_build_700)

		dm_build_897.WriteString(dm_build_902)
		var strLen = dm_build_900.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_902))
		}
		dm_build_898 += int32(strLen)
		if dm_build_894.readOver {
			break
		}
	}
	return dm_build_897.String(), nil
}

func (dm_build_905 *dm_build_696) dm_build_904(dm_build_906 *DmClob, dm_build_907 int, dm_build_908 string, dm_build_909 string) (int, error) {
	var dm_build_910 = Dm_build_1.Dm_build_217(dm_build_908, dm_build_909, dm_build_905.dm_build_700)
	var dm_build_911 = 0
	var dm_build_912 = len(dm_build_910)
	var dm_build_913 = 0
	var dm_build_914 = 0
	var dm_build_915 = 0
	var dm_build_916 = dm_build_912/Dm_build_1102 + 1
	var dm_build_917 byte = 0
	var dm_build_918 byte = 0x01
	var dm_build_919 byte = 0x02
	for i := 0; i < dm_build_916; i++ {
		dm_build_917 = 0
		if i == 0 {
			dm_build_917 |= dm_build_918
		}
		if i == dm_build_916-1 {
			dm_build_917 |= dm_build_919
		}
		dm_build_915 = dm_build_912 - dm_build_914
		if dm_build_915 > Dm_build_1102 {
			dm_build_915 = Dm_build_1102
		}

		setLobData := dm_build_1530(dm_build_905, &dm_build_906.lob, dm_build_917, dm_build_907, dm_build_910, dm_build_911, dm_build_915)
		ret, err := dm_build_905.dm_build_739(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_913, nil
		} else {
			dm_build_907 += int(tmp)
			dm_build_913 += int(tmp)
			dm_build_914 += dm_build_915
			dm_build_911 += dm_build_915
		}
	}
	return dm_build_913, nil
}

func (dm_build_921 *dm_build_696) dm_build_920(dm_build_922 *DmBlob, dm_build_923 int, dm_build_924 []byte) (int, error) {
	var dm_build_925 = 0
	var dm_build_926 = len(dm_build_924)
	var dm_build_927 = 0
	var dm_build_928 = 0
	var dm_build_929 = 0
	var dm_build_930 = dm_build_926/Dm_build_1102 + 1
	var dm_build_931 byte = 0
	var dm_build_932 byte = 0x01
	var dm_build_933 byte = 0x02
	for i := 0; i < dm_build_930; i++ {
		dm_build_931 = 0
		if i == 0 {
			dm_build_931 |= dm_build_932
		}
		if i == dm_build_930-1 {
			dm_build_931 |= dm_build_933
		}
		dm_build_929 = dm_build_926 - dm_build_928
		if dm_build_929 > Dm_build_1102 {
			dm_build_929 = Dm_build_1102
		}

		setLobData := dm_build_1530(dm_build_921, &dm_build_922.lob, dm_build_931, dm_build_923, dm_build_924, dm_build_925, dm_build_929)
		ret, err := dm_build_921.dm_build_739(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_927, nil
		} else {
			dm_build_923 += int(tmp)
			dm_build_927 += int(tmp)
			dm_build_928 += dm_build_929
			dm_build_925 += dm_build_929
		}
	}
	return dm_build_927, nil
}

func (dm_build_935 *dm_build_696) dm_build_934(dm_build_936 *lob, dm_build_937 int) (int64, error) {
	dm_build_938 := dm_build_1396(dm_build_935, dm_build_936, dm_build_937)
	dm_build_939, dm_build_940 := dm_build_935.dm_build_739(dm_build_938)
	if dm_build_940 != nil {
		return dm_build_936.length, dm_build_940
	}
	return dm_build_939.(int64), nil
}

func (dm_build_942 *dm_build_696) dm_build_941(dm_build_943 []interface{}, dm_build_944 []interface{}, dm_build_945 int) bool {
	var dm_build_946 = false
	dm_build_943[dm_build_945] = dm_build_944[dm_build_945]

	if binder, ok := dm_build_944[dm_build_945].(iOffRowBinder); ok {
		dm_build_946 = true
		dm_build_943[dm_build_945] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_942.dm_build_700) {
			dm_build_943[dm_build_945] = &lobCtl{lob.buildCtlData()}
			dm_build_946 = false
		}
	} else {
		dm_build_943[dm_build_945] = dm_build_944[dm_build_945]
	}
	return dm_build_946
}

func (dm_build_948 *dm_build_696) dm_build_947(dm_build_949 *DmStatement, dm_build_950 parameter, dm_build_951 int, dm_build_952 iOffRowBinder) error {
	var dm_build_953 = Dm_build_286()
	dm_build_952.read(dm_build_953)
	var dm_build_954 = 0
	for !dm_build_952.isReadOver() || dm_build_953.Dm_build_287() > 0 {
		if !dm_build_952.isReadOver() && dm_build_953.Dm_build_287() < Dm_build_1102 {
			dm_build_952.read(dm_build_953)
		}
		if dm_build_953.Dm_build_287() > Dm_build_1102 {
			dm_build_954 = Dm_build_1102
		} else {
			dm_build_954 = dm_build_953.Dm_build_287()
		}

		putData := dm_build_1501(dm_build_948, dm_build_949, int16(dm_build_951), dm_build_953, int32(dm_build_954))
		_, err := dm_build_948.dm_build_739(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_956 *dm_build_696) dm_build_955() ([]byte, error) {
	var dm_build_957 error
	if dm_build_956.dm_build_704 == nil {
		if dm_build_956.dm_build_704, dm_build_957 = security.NewClientKeyPair(); dm_build_957 != nil {
			return nil, dm_build_957
		}
	}
	return security.Bn2Bytes(dm_build_956.dm_build_704.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_959 *dm_build_696) dm_build_958() (*security.DhKey, error) {
	var dm_build_960 error
	if dm_build_959.dm_build_704 == nil {
		if dm_build_959.dm_build_704, dm_build_960 = security.NewClientKeyPair(); dm_build_960 != nil {
			return nil, dm_build_960
		}
	}
	return dm_build_959.dm_build_704, nil
}

func (dm_build_962 *dm_build_696) dm_build_961(dm_build_963 int, dm_build_964 []byte, dm_build_965 string, dm_build_966 int) (dm_build_967 error) {
	if dm_build_963 > 0 && dm_build_963 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_964 != nil {
		dm_build_962.dm_build_701, dm_build_967 = security.NewSymmCipher(dm_build_963, dm_build_964)
	} else if dm_build_963 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_962.dm_build_701, dm_build_967 = security.NewThirdPartCipher(dm_build_963, dm_build_964, dm_build_965, dm_build_966); dm_build_967 != nil {
			dm_build_967 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_967.Error()).throw()
		}
	}
	return
}

func (dm_build_969 *dm_build_696) dm_build_968(dm_build_970 bool) (dm_build_971 error) {
	if dm_build_969.dm_build_698, dm_build_971 = security.NewTLSFromTCP(dm_build_969.dm_build_697, dm_build_969.dm_build_700.dmConnector.sslCertPath, dm_build_969.dm_build_700.dmConnector.sslKeyPath, dm_build_969.dm_build_700.dmConnector.user); dm_build_971 != nil {
		return
	}
	if !dm_build_970 {
		dm_build_969.dm_build_698 = nil
	}
	return
}

func (dm_build_973 *dm_build_696) dm_build_972(dm_build_974 dm_build_1110) bool {
	return dm_build_974.dm_build_1125() != Dm_build_1017 && dm_build_973.dm_build_700.sslEncrypt == 1
}
