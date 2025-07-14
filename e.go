/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_651 struct{}

var Dm_build_652 = &dm_build_651{}

func (Dm_build_654 *dm_build_651) Dm_build_653(dm_build_655 []byte, dm_build_656 int, dm_build_657 byte) int {
	dm_build_655[dm_build_656] = dm_build_657
	return 1
}

func (Dm_build_659 *dm_build_651) Dm_build_658(dm_build_660 []byte, dm_build_661 int, dm_build_662 int8) int {
	dm_build_660[dm_build_661] = byte(dm_build_662)
	return 1
}

func (Dm_build_664 *dm_build_651) Dm_build_663(dm_build_665 []byte, dm_build_666 int, dm_build_667 int16) int {
	dm_build_665[dm_build_666] = byte(dm_build_667)
	dm_build_666++
	dm_build_665[dm_build_666] = byte(dm_build_667 >> 8)
	return 2
}

func (Dm_build_669 *dm_build_651) Dm_build_668(dm_build_670 []byte, dm_build_671 int, dm_build_672 int32) int {
	dm_build_670[dm_build_671] = byte(dm_build_672)
	dm_build_671++
	dm_build_670[dm_build_671] = byte(dm_build_672 >> 8)
	dm_build_671++
	dm_build_670[dm_build_671] = byte(dm_build_672 >> 16)
	dm_build_671++
	dm_build_670[dm_build_671] = byte(dm_build_672 >> 24)
	dm_build_671++
	return 4
}

func (Dm_build_674 *dm_build_651) Dm_build_673(dm_build_675 []byte, dm_build_676 int, dm_build_677 int64) int {
	dm_build_675[dm_build_676] = byte(dm_build_677)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 8)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 16)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 24)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 32)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 40)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 48)
	dm_build_676++
	dm_build_675[dm_build_676] = byte(dm_build_677 >> 56)
	return 8
}

func (Dm_build_679 *dm_build_651) Dm_build_678(dm_build_680 []byte, dm_build_681 int, dm_build_682 float32) int {
	return Dm_build_679.Dm_build_698(dm_build_680, dm_build_681, math.Float32bits(dm_build_682))
}

func (Dm_build_684 *dm_build_651) Dm_build_683(dm_build_685 []byte, dm_build_686 int, dm_build_687 float64) int {
	return Dm_build_684.Dm_build_703(dm_build_685, dm_build_686, math.Float64bits(dm_build_687))
}

func (Dm_build_689 *dm_build_651) Dm_build_688(dm_build_690 []byte, dm_build_691 int, dm_build_692 uint8) int {
	dm_build_690[dm_build_691] = byte(dm_build_692)
	return 1
}

func (Dm_build_694 *dm_build_651) Dm_build_693(dm_build_695 []byte, dm_build_696 int, dm_build_697 uint16) int {
	dm_build_695[dm_build_696] = byte(dm_build_697)
	dm_build_696++
	dm_build_695[dm_build_696] = byte(dm_build_697 >> 8)
	return 2
}

func (Dm_build_699 *dm_build_651) Dm_build_698(dm_build_700 []byte, dm_build_701 int, dm_build_702 uint32) int {
	dm_build_700[dm_build_701] = byte(dm_build_702)
	dm_build_701++
	dm_build_700[dm_build_701] = byte(dm_build_702 >> 8)
	dm_build_701++
	dm_build_700[dm_build_701] = byte(dm_build_702 >> 16)
	dm_build_701++
	dm_build_700[dm_build_701] = byte(dm_build_702 >> 24)
	return 3
}

func (Dm_build_704 *dm_build_651) Dm_build_703(dm_build_705 []byte, dm_build_706 int, dm_build_707 uint64) int {
	dm_build_705[dm_build_706] = byte(dm_build_707)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 8)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 16)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 24)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 32)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 40)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 48)
	dm_build_706++
	dm_build_705[dm_build_706] = byte(dm_build_707 >> 56)
	return 3
}

func (Dm_build_709 *dm_build_651) Dm_build_708(dm_build_710 []byte, dm_build_711 int, dm_build_712 []byte, dm_build_713 int, dm_build_714 int) int {
	copy(dm_build_710[dm_build_711:dm_build_711+dm_build_714], dm_build_712[dm_build_713:dm_build_713+dm_build_714])
	return dm_build_714
}

func (Dm_build_716 *dm_build_651) Dm_build_715(dm_build_717 []byte, dm_build_718 int, dm_build_719 []byte, dm_build_720 int, dm_build_721 int) int {
	dm_build_718 += Dm_build_716.Dm_build_698(dm_build_717, dm_build_718, uint32(dm_build_721))
	return 4 + Dm_build_716.Dm_build_708(dm_build_717, dm_build_718, dm_build_719, dm_build_720, dm_build_721)
}

func (Dm_build_723 *dm_build_651) Dm_build_722(dm_build_724 []byte, dm_build_725 int, dm_build_726 []byte, dm_build_727 int, dm_build_728 int) int {
	dm_build_725 += Dm_build_723.Dm_build_693(dm_build_724, dm_build_725, uint16(dm_build_728))
	return 2 + Dm_build_723.Dm_build_708(dm_build_724, dm_build_725, dm_build_726, dm_build_727, dm_build_728)
}

func (Dm_build_730 *dm_build_651) Dm_build_729(dm_build_731 []byte, dm_build_732 int, dm_build_733 string, dm_build_734 string, dm_build_735 *DmConnection) int {
	dm_build_736 := Dm_build_730.Dm_build_868(dm_build_733, dm_build_734, dm_build_735)
	dm_build_732 += Dm_build_730.Dm_build_698(dm_build_731, dm_build_732, uint32(len(dm_build_736)))
	return 4 + Dm_build_730.Dm_build_708(dm_build_731, dm_build_732, dm_build_736, 0, len(dm_build_736))
}

func (Dm_build_738 *dm_build_651) Dm_build_737(dm_build_739 []byte, dm_build_740 int, dm_build_741 string, dm_build_742 string, dm_build_743 *DmConnection) int {
	dm_build_744 := Dm_build_738.Dm_build_868(dm_build_741, dm_build_742, dm_build_743)

	dm_build_740 += Dm_build_738.Dm_build_693(dm_build_739, dm_build_740, uint16(len(dm_build_744)))
	return 2 + Dm_build_738.Dm_build_708(dm_build_739, dm_build_740, dm_build_744, 0, len(dm_build_744))
}

func (Dm_build_746 *dm_build_651) Dm_build_745(dm_build_747 []byte, dm_build_748 int) byte {
	return dm_build_747[dm_build_748]
}

func (Dm_build_750 *dm_build_651) Dm_build_749(dm_build_751 []byte, dm_build_752 int) int16 {
	var dm_build_753 int16
	dm_build_753 = int16(dm_build_751[dm_build_752] & 0xff)
	dm_build_752++
	dm_build_753 |= int16(dm_build_751[dm_build_752]&0xff) << 8
	return dm_build_753
}

func (Dm_build_755 *dm_build_651) Dm_build_754(dm_build_756 []byte, dm_build_757 int) int32 {
	var dm_build_758 int32
	dm_build_758 = int32(dm_build_756[dm_build_757] & 0xff)
	dm_build_757++
	dm_build_758 |= int32(dm_build_756[dm_build_757]&0xff) << 8
	dm_build_757++
	dm_build_758 |= int32(dm_build_756[dm_build_757]&0xff) << 16
	dm_build_757++
	dm_build_758 |= int32(dm_build_756[dm_build_757]&0xff) << 24
	return dm_build_758
}

func (Dm_build_760 *dm_build_651) Dm_build_759(dm_build_761 []byte, dm_build_762 int) int64 {
	var dm_build_763 int64
	dm_build_763 = int64(dm_build_761[dm_build_762] & 0xff)
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 8
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 16
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 24
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 32
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 40
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 48
	dm_build_762++
	dm_build_763 |= int64(dm_build_761[dm_build_762]&0xff) << 56
	return dm_build_763
}

func (Dm_build_765 *dm_build_651) Dm_build_764(dm_build_766 []byte, dm_build_767 int) float32 {
	return math.Float32frombits(Dm_build_765.Dm_build_781(dm_build_766, dm_build_767))
}

func (Dm_build_769 *dm_build_651) Dm_build_768(dm_build_770 []byte, dm_build_771 int) float64 {
	return math.Float64frombits(Dm_build_769.Dm_build_786(dm_build_770, dm_build_771))
}

func (Dm_build_773 *dm_build_651) Dm_build_772(dm_build_774 []byte, dm_build_775 int) uint8 {
	return uint8(dm_build_774[dm_build_775] & 0xff)
}

func (Dm_build_777 *dm_build_651) Dm_build_776(dm_build_778 []byte, dm_build_779 int) uint16 {
	var dm_build_780 uint16
	dm_build_780 = uint16(dm_build_778[dm_build_779] & 0xff)
	dm_build_779++
	dm_build_780 |= uint16(dm_build_778[dm_build_779]&0xff) << 8
	return dm_build_780
}

func (Dm_build_782 *dm_build_651) Dm_build_781(dm_build_783 []byte, dm_build_784 int) uint32 {
	var dm_build_785 uint32
	dm_build_785 = uint32(dm_build_783[dm_build_784] & 0xff)
	dm_build_784++
	dm_build_785 |= uint32(dm_build_783[dm_build_784]&0xff) << 8
	dm_build_784++
	dm_build_785 |= uint32(dm_build_783[dm_build_784]&0xff) << 16
	dm_build_784++
	dm_build_785 |= uint32(dm_build_783[dm_build_784]&0xff) << 24
	return dm_build_785
}

func (Dm_build_787 *dm_build_651) Dm_build_786(dm_build_788 []byte, dm_build_789 int) uint64 {
	var dm_build_790 uint64
	dm_build_790 = uint64(dm_build_788[dm_build_789] & 0xff)
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 8
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 16
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 24
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 32
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 40
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 48
	dm_build_789++
	dm_build_790 |= uint64(dm_build_788[dm_build_789]&0xff) << 56
	return dm_build_790
}

func (Dm_build_792 *dm_build_651) Dm_build_791(dm_build_793 []byte, dm_build_794 int) []byte {
	dm_build_795 := Dm_build_792.Dm_build_781(dm_build_793, dm_build_794)

	dm_build_796 := make([]byte, dm_build_795)
	copy(dm_build_796[:int(dm_build_795)], dm_build_793[dm_build_794+4:dm_build_794+4+int(dm_build_795)])
	return dm_build_796
}

func (Dm_build_798 *dm_build_651) Dm_build_797(dm_build_799 []byte, dm_build_800 int) []byte {
	dm_build_801 := Dm_build_798.Dm_build_776(dm_build_799, dm_build_800)

	dm_build_802 := make([]byte, dm_build_801)
	copy(dm_build_802[:int(dm_build_801)], dm_build_799[dm_build_800+2:dm_build_800+2+int(dm_build_801)])
	return dm_build_802
}

func (Dm_build_804 *dm_build_651) Dm_build_803(dm_build_805 []byte, dm_build_806 int, dm_build_807 int) []byte {

	dm_build_808 := make([]byte, dm_build_807)
	copy(dm_build_808[:dm_build_807], dm_build_805[dm_build_806:dm_build_806+dm_build_807])
	return dm_build_808
}

func (Dm_build_810 *dm_build_651) Dm_build_809(dm_build_811 []byte, dm_build_812 int, dm_build_813 int, dm_build_814 string, dm_build_815 *DmConnection) string {
	return Dm_build_810.Dm_build_904(dm_build_811[dm_build_812:dm_build_812+dm_build_813], dm_build_814, dm_build_815)
}

func (Dm_build_817 *dm_build_651) Dm_build_816(dm_build_818 []byte, dm_build_819 int, dm_build_820 string, dm_build_821 *DmConnection) string {
	dm_build_822 := Dm_build_817.Dm_build_781(dm_build_818, dm_build_819)
	dm_build_819 += 4
	return Dm_build_817.Dm_build_809(dm_build_818, dm_build_819, int(dm_build_822), dm_build_820, dm_build_821)
}

func (Dm_build_824 *dm_build_651) Dm_build_823(dm_build_825 []byte, dm_build_826 int, dm_build_827 string, dm_build_828 *DmConnection) string {
	dm_build_829 := Dm_build_824.Dm_build_776(dm_build_825, dm_build_826)
	dm_build_826 += 2
	return Dm_build_824.Dm_build_809(dm_build_825, dm_build_826, int(dm_build_829), dm_build_827, dm_build_828)
}

func (Dm_build_831 *dm_build_651) Dm_build_830(dm_build_832 byte) []byte {
	return []byte{dm_build_832}
}

func (Dm_build_834 *dm_build_651) Dm_build_833(dm_build_835 int8) []byte {
	return []byte{byte(dm_build_835)}
}

func (Dm_build_837 *dm_build_651) Dm_build_836(dm_build_838 int16) []byte {
	return []byte{byte(dm_build_838), byte(dm_build_838 >> 8)}
}

func (Dm_build_840 *dm_build_651) Dm_build_839(dm_build_841 int32) []byte {
	return []byte{byte(dm_build_841), byte(dm_build_841 >> 8), byte(dm_build_841 >> 16), byte(dm_build_841 >> 24)}
}

func (Dm_build_843 *dm_build_651) Dm_build_842(dm_build_844 int64) []byte {
	return []byte{byte(dm_build_844), byte(dm_build_844 >> 8), byte(dm_build_844 >> 16), byte(dm_build_844 >> 24), byte(dm_build_844 >> 32),
		byte(dm_build_844 >> 40), byte(dm_build_844 >> 48), byte(dm_build_844 >> 56)}
}

func (Dm_build_846 *dm_build_651) Dm_build_845(dm_build_847 float32) []byte {
	return Dm_build_846.Dm_build_857(math.Float32bits(dm_build_847))
}

func (Dm_build_849 *dm_build_651) Dm_build_848(dm_build_850 float64) []byte {
	return Dm_build_849.Dm_build_860(math.Float64bits(dm_build_850))
}

func (Dm_build_852 *dm_build_651) Dm_build_851(dm_build_853 uint8) []byte {
	return []byte{byte(dm_build_853)}
}

func (Dm_build_855 *dm_build_651) Dm_build_854(dm_build_856 uint16) []byte {
	return []byte{byte(dm_build_856), byte(dm_build_856 >> 8)}
}

func (Dm_build_858 *dm_build_651) Dm_build_857(dm_build_859 uint32) []byte {
	return []byte{byte(dm_build_859), byte(dm_build_859 >> 8), byte(dm_build_859 >> 16), byte(dm_build_859 >> 24)}
}

func (Dm_build_861 *dm_build_651) Dm_build_860(dm_build_862 uint64) []byte {
	return []byte{byte(dm_build_862), byte(dm_build_862 >> 8), byte(dm_build_862 >> 16), byte(dm_build_862 >> 24), byte(dm_build_862 >> 32), byte(dm_build_862 >> 40), byte(dm_build_862 >> 48), byte(dm_build_862 >> 56)}
}

func (Dm_build_864 *dm_build_651) Dm_build_863(dm_build_865 []byte, dm_build_866 string, dm_build_867 *DmConnection) []byte {
	if dm_build_866 == "UTF-8" {
		return dm_build_865
	}

	if dm_build_867 == nil {
		if e := dm_build_909(dm_build_866); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_865), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_867.encodeBuffer == nil {
		dm_build_867.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_867.encode = dm_build_909(dm_build_867.getServerEncoding())
		dm_build_867.transformReaderDst = make([]byte, 4096)
		dm_build_867.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_867.encode; e != nil {

		dm_build_867.encodeBuffer.Reset()

		n, err := dm_build_867.encodeBuffer.ReadFrom(
			Dm_build_923(bytes.NewReader(dm_build_865), e.NewEncoder(), dm_build_867.transformReaderDst, dm_build_867.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_867.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_869 *dm_build_651) Dm_build_868(dm_build_870 string, dm_build_871 string, dm_build_872 *DmConnection) []byte {
	return Dm_build_869.Dm_build_863([]byte(dm_build_870), dm_build_871, dm_build_872)
}

func (Dm_build_874 *dm_build_651) Dm_build_873(dm_build_875 []byte) byte {
	return Dm_build_874.Dm_build_745(dm_build_875, 0)
}

func (Dm_build_877 *dm_build_651) Dm_build_876(dm_build_878 []byte) int16 {
	return Dm_build_877.Dm_build_749(dm_build_878, 0)
}

func (Dm_build_880 *dm_build_651) Dm_build_879(dm_build_881 []byte) int32 {
	return Dm_build_880.Dm_build_754(dm_build_881, 0)
}

func (Dm_build_883 *dm_build_651) Dm_build_882(dm_build_884 []byte) int64 {
	return Dm_build_883.Dm_build_759(dm_build_884, 0)
}

func (Dm_build_886 *dm_build_651) Dm_build_885(dm_build_887 []byte) float32 {
	return Dm_build_886.Dm_build_764(dm_build_887, 0)
}

func (Dm_build_889 *dm_build_651) Dm_build_888(dm_build_890 []byte) float64 {
	return Dm_build_889.Dm_build_768(dm_build_890, 0)
}

func (Dm_build_892 *dm_build_651) Dm_build_891(dm_build_893 []byte) uint8 {
	return Dm_build_892.Dm_build_772(dm_build_893, 0)
}

func (Dm_build_895 *dm_build_651) Dm_build_894(dm_build_896 []byte) uint16 {
	return Dm_build_895.Dm_build_776(dm_build_896, 0)
}

func (Dm_build_898 *dm_build_651) Dm_build_897(dm_build_899 []byte) uint32 {
	return Dm_build_898.Dm_build_781(dm_build_899, 0)
}

func (Dm_build_901 *dm_build_651) Dm_build_900(dm_build_902 []byte, dm_build_903 string) []byte {
	if dm_build_903 == "UTF-8" {
		return dm_build_902
	}

	if e := dm_build_909(dm_build_903); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_902), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_905 *dm_build_651) Dm_build_904(dm_build_906 []byte, dm_build_907 string, dm_build_908 *DmConnection) string {
	return string(Dm_build_905.Dm_build_900(dm_build_906, dm_build_907))
}

func dm_build_909(dm_build_910 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_910); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_911 struct {
	dm_build_912 io.Reader
	dm_build_913 transform.Transformer
	dm_build_914 error

	dm_build_915               []byte
	dm_build_916, dm_build_917 int

	dm_build_918               []byte
	dm_build_919, dm_build_920 int

	dm_build_921 bool
}

const dm_build_922 = 4096

func Dm_build_923(dm_build_924 io.Reader, dm_build_925 transform.Transformer, dm_build_926 []byte, dm_build_927 []byte) *Dm_build_911 {
	dm_build_925.Reset()
	return &Dm_build_911{
		dm_build_912: dm_build_924,
		dm_build_913: dm_build_925,
		dm_build_915: dm_build_926,
		dm_build_918: dm_build_927,
	}
}

func (dm_build_929 *Dm_build_911) Read(dm_build_930 []byte) (int, error) {
	dm_build_931, dm_build_932 := 0, error(nil)
	for {

		if dm_build_929.dm_build_916 != dm_build_929.dm_build_917 {
			dm_build_931 = copy(dm_build_930, dm_build_929.dm_build_915[dm_build_929.dm_build_916:dm_build_929.dm_build_917])
			dm_build_929.dm_build_916 += dm_build_931
			if dm_build_929.dm_build_916 == dm_build_929.dm_build_917 && dm_build_929.dm_build_921 {
				return dm_build_931, dm_build_929.dm_build_914
			}
			return dm_build_931, nil
		} else if dm_build_929.dm_build_921 {
			return 0, dm_build_929.dm_build_914
		}

		if dm_build_929.dm_build_919 != dm_build_929.dm_build_920 || dm_build_929.dm_build_914 != nil {
			dm_build_929.dm_build_916 = 0
			dm_build_929.dm_build_917, dm_build_931, dm_build_932 = dm_build_929.dm_build_913.Transform(dm_build_929.dm_build_915, dm_build_929.dm_build_918[dm_build_929.dm_build_919:dm_build_929.dm_build_920], dm_build_929.dm_build_914 == io.EOF)
			dm_build_929.dm_build_919 += dm_build_931

			switch {
			case dm_build_932 == nil:
				if dm_build_929.dm_build_919 != dm_build_929.dm_build_920 {
					dm_build_929.dm_build_914 = nil
				}

				dm_build_929.dm_build_921 = dm_build_929.dm_build_914 != nil
				continue
			case dm_build_932 == transform.ErrShortDst && (dm_build_929.dm_build_917 != 0 || dm_build_931 != 0):

				continue
			case dm_build_932 == transform.ErrShortSrc && dm_build_929.dm_build_920-dm_build_929.dm_build_919 != len(dm_build_929.dm_build_918) && dm_build_929.dm_build_914 == nil:

			default:
				dm_build_929.dm_build_921 = true

				if dm_build_929.dm_build_914 == nil || dm_build_929.dm_build_914 == io.EOF {
					dm_build_929.dm_build_914 = dm_build_932
				}
				continue
			}
		}

		if dm_build_929.dm_build_919 != 0 {
			dm_build_929.dm_build_919, dm_build_929.dm_build_920 = 0, copy(dm_build_929.dm_build_918, dm_build_929.dm_build_918[dm_build_929.dm_build_919:dm_build_929.dm_build_920])
		}
		dm_build_931, dm_build_929.dm_build_914 = dm_build_929.dm_build_912.Read(dm_build_929.dm_build_918[dm_build_929.dm_build_920:])
		dm_build_929.dm_build_920 += dm_build_931
	}
}
