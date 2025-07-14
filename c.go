/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1011 struct {
	dm_build_1012 []byte
	dm_build_1013 int
}

func Dm_build_1014(dm_build_1015 int) *Dm_build_1011 {
	return &Dm_build_1011{make([]byte, 0, dm_build_1015), 0}
}

func Dm_build_1016(dm_build_1017 []byte) *Dm_build_1011 {
	return &Dm_build_1011{dm_build_1017, 0}
}

func (dm_build_1019 *Dm_build_1011) dm_build_1018(dm_build_1020 int) *Dm_build_1011 {

	dm_build_1021 := len(dm_build_1019.dm_build_1012)
	dm_build_1022 := cap(dm_build_1019.dm_build_1012)

	if dm_build_1021+dm_build_1020 <= dm_build_1022 {
		dm_build_1019.dm_build_1012 = dm_build_1019.dm_build_1012[:dm_build_1021+dm_build_1020]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1022), float64(dm_build_1020+dm_build_1021)))

		nbuf := make([]byte, dm_build_1020+dm_build_1021, calCap)
		copy(nbuf, dm_build_1019.dm_build_1012)
		dm_build_1019.dm_build_1012 = nbuf
	}

	return dm_build_1019
}

func (dm_build_1024 *Dm_build_1011) Dm_build_1023() int {
	return len(dm_build_1024.dm_build_1012)
}

func (dm_build_1026 *Dm_build_1011) Dm_build_1025(dm_build_1027 int) *Dm_build_1011 {
	for i := dm_build_1027; i < len(dm_build_1026.dm_build_1012); i++ {
		dm_build_1026.dm_build_1012[i] = 0
	}
	dm_build_1026.dm_build_1012 = dm_build_1026.dm_build_1012[:dm_build_1027]
	return dm_build_1026
}

func (dm_build_1029 *Dm_build_1011) Dm_build_1028(dm_build_1030 int) *Dm_build_1011 {
	dm_build_1029.dm_build_1013 = dm_build_1030
	return dm_build_1029
}

func (dm_build_1032 *Dm_build_1011) Dm_build_1031() int {
	return dm_build_1032.dm_build_1013
}

func (dm_build_1034 *Dm_build_1011) Dm_build_1033(dm_build_1035 bool) int {
	return len(dm_build_1034.dm_build_1012) - dm_build_1034.dm_build_1013
}

func (dm_build_1037 *Dm_build_1011) Dm_build_1036(dm_build_1038 int, dm_build_1039 bool, dm_build_1040 bool) *Dm_build_1011 {

	if dm_build_1039 {
		if dm_build_1040 {
			dm_build_1037.dm_build_1018(dm_build_1038)
		} else {
			dm_build_1037.dm_build_1012 = dm_build_1037.dm_build_1012[:len(dm_build_1037.dm_build_1012)-dm_build_1038]
		}
	} else {
		if dm_build_1040 {
			dm_build_1037.dm_build_1013 += dm_build_1038
		} else {
			dm_build_1037.dm_build_1013 -= dm_build_1038
		}
	}

	return dm_build_1037
}

func (dm_build_1042 *Dm_build_1011) Dm_build_1041(dm_build_1043 io.Reader, dm_build_1044 int) (int, error) {
	dm_build_1045 := len(dm_build_1042.dm_build_1012)
	dm_build_1042.dm_build_1018(dm_build_1044)
	dm_build_1046 := 0
	for dm_build_1044 > 0 {
		n, err := dm_build_1043.Read(dm_build_1042.dm_build_1012[dm_build_1045+dm_build_1046:])
		if n > 0 && err == io.EOF {
			dm_build_1046 += n
			dm_build_1042.dm_build_1012 = dm_build_1042.dm_build_1012[:dm_build_1045+dm_build_1046]
			return dm_build_1046, nil
		} else if n > 0 && err == nil {
			dm_build_1044 -= n
			dm_build_1046 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1046, nil
}

func (dm_build_1048 *Dm_build_1011) Dm_build_1047(dm_build_1049 io.Writer) (*Dm_build_1011, error) {
	if _, err := dm_build_1049.Write(dm_build_1048.dm_build_1012); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1048, nil
}

func (dm_build_1051 *Dm_build_1011) Dm_build_1050(dm_build_1052 bool) int {
	dm_build_1053 := len(dm_build_1051.dm_build_1012)
	dm_build_1051.dm_build_1018(1)

	if dm_build_1052 {
		return copy(dm_build_1051.dm_build_1012[dm_build_1053:], []byte{1})
	} else {
		return copy(dm_build_1051.dm_build_1012[dm_build_1053:], []byte{0})
	}
}

func (dm_build_1055 *Dm_build_1011) Dm_build_1054(dm_build_1056 byte) int {
	dm_build_1057 := len(dm_build_1055.dm_build_1012)
	dm_build_1055.dm_build_1018(1)

	return copy(dm_build_1055.dm_build_1012[dm_build_1057:], Dm_build_652.Dm_build_830(dm_build_1056))
}

func (dm_build_1059 *Dm_build_1011) Dm_build_1058(dm_build_1060 int8) int {
	dm_build_1061 := len(dm_build_1059.dm_build_1012)
	dm_build_1059.dm_build_1018(1)

	return copy(dm_build_1059.dm_build_1012[dm_build_1061:], Dm_build_652.Dm_build_833(dm_build_1060))
}

func (dm_build_1063 *Dm_build_1011) Dm_build_1062(dm_build_1064 int16) int {
	dm_build_1065 := len(dm_build_1063.dm_build_1012)
	dm_build_1063.dm_build_1018(2)

	return copy(dm_build_1063.dm_build_1012[dm_build_1065:], Dm_build_652.Dm_build_836(dm_build_1064))
}

func (dm_build_1067 *Dm_build_1011) Dm_build_1066(dm_build_1068 int32) int {
	dm_build_1069 := len(dm_build_1067.dm_build_1012)
	dm_build_1067.dm_build_1018(4)

	return copy(dm_build_1067.dm_build_1012[dm_build_1069:], Dm_build_652.Dm_build_839(dm_build_1068))
}

func (dm_build_1071 *Dm_build_1011) Dm_build_1070(dm_build_1072 uint8) int {
	dm_build_1073 := len(dm_build_1071.dm_build_1012)
	dm_build_1071.dm_build_1018(1)

	return copy(dm_build_1071.dm_build_1012[dm_build_1073:], Dm_build_652.Dm_build_851(dm_build_1072))
}

func (dm_build_1075 *Dm_build_1011) Dm_build_1074(dm_build_1076 uint16) int {
	dm_build_1077 := len(dm_build_1075.dm_build_1012)
	dm_build_1075.dm_build_1018(2)

	return copy(dm_build_1075.dm_build_1012[dm_build_1077:], Dm_build_652.Dm_build_854(dm_build_1076))
}

func (dm_build_1079 *Dm_build_1011) Dm_build_1078(dm_build_1080 uint32) int {
	dm_build_1081 := len(dm_build_1079.dm_build_1012)
	dm_build_1079.dm_build_1018(4)

	return copy(dm_build_1079.dm_build_1012[dm_build_1081:], Dm_build_652.Dm_build_857(dm_build_1080))
}

func (dm_build_1083 *Dm_build_1011) Dm_build_1082(dm_build_1084 uint64) int {
	dm_build_1085 := len(dm_build_1083.dm_build_1012)
	dm_build_1083.dm_build_1018(8)

	return copy(dm_build_1083.dm_build_1012[dm_build_1085:], Dm_build_652.Dm_build_860(dm_build_1084))
}

func (dm_build_1087 *Dm_build_1011) Dm_build_1086(dm_build_1088 float32) int {
	dm_build_1089 := len(dm_build_1087.dm_build_1012)
	dm_build_1087.dm_build_1018(4)

	return copy(dm_build_1087.dm_build_1012[dm_build_1089:], Dm_build_652.Dm_build_857(math.Float32bits(dm_build_1088)))
}

func (dm_build_1091 *Dm_build_1011) Dm_build_1090(dm_build_1092 float64) int {
	dm_build_1093 := len(dm_build_1091.dm_build_1012)
	dm_build_1091.dm_build_1018(8)

	return copy(dm_build_1091.dm_build_1012[dm_build_1093:], Dm_build_652.Dm_build_860(math.Float64bits(dm_build_1092)))
}

func (dm_build_1095 *Dm_build_1011) Dm_build_1094(dm_build_1096 []byte) int {
	dm_build_1097 := len(dm_build_1095.dm_build_1012)
	dm_build_1095.dm_build_1018(len(dm_build_1096))
	return copy(dm_build_1095.dm_build_1012[dm_build_1097:], dm_build_1096)
}

func (dm_build_1099 *Dm_build_1011) Dm_build_1098(dm_build_1100 []byte) int {
	return dm_build_1099.Dm_build_1066(int32(len(dm_build_1100))) + dm_build_1099.Dm_build_1094(dm_build_1100)
}

func (dm_build_1102 *Dm_build_1011) Dm_build_1101(dm_build_1103 []byte) int {
	return dm_build_1102.Dm_build_1070(uint8(len(dm_build_1103))) + dm_build_1102.Dm_build_1094(dm_build_1103)
}

func (dm_build_1105 *Dm_build_1011) Dm_build_1104(dm_build_1106 []byte) int {
	return dm_build_1105.Dm_build_1074(uint16(len(dm_build_1106))) + dm_build_1105.Dm_build_1094(dm_build_1106)
}

func (dm_build_1108 *Dm_build_1011) Dm_build_1107(dm_build_1109 []byte) int {
	return dm_build_1108.Dm_build_1094(dm_build_1109) + dm_build_1108.Dm_build_1054(0)
}

func (dm_build_1111 *Dm_build_1011) Dm_build_1110(dm_build_1112 string, dm_build_1113 string, dm_build_1114 *DmConnection) int {
	dm_build_1115 := Dm_build_652.Dm_build_868(dm_build_1112, dm_build_1113, dm_build_1114)
	return dm_build_1111.Dm_build_1098(dm_build_1115)
}

func (dm_build_1117 *Dm_build_1011) Dm_build_1116(dm_build_1118 string, dm_build_1119 string, dm_build_1120 *DmConnection) int {
	dm_build_1121 := Dm_build_652.Dm_build_868(dm_build_1118, dm_build_1119, dm_build_1120)
	return dm_build_1117.Dm_build_1101(dm_build_1121)
}

func (dm_build_1123 *Dm_build_1011) Dm_build_1122(dm_build_1124 string, dm_build_1125 string, dm_build_1126 *DmConnection) int {
	dm_build_1127 := Dm_build_652.Dm_build_868(dm_build_1124, dm_build_1125, dm_build_1126)
	return dm_build_1123.Dm_build_1104(dm_build_1127)
}

func (dm_build_1129 *Dm_build_1011) Dm_build_1128(dm_build_1130 string, dm_build_1131 string, dm_build_1132 *DmConnection) int {
	dm_build_1133 := Dm_build_652.Dm_build_868(dm_build_1130, dm_build_1131, dm_build_1132)
	return dm_build_1129.Dm_build_1107(dm_build_1133)
}

func (dm_build_1135 *Dm_build_1011) Dm_build_1134() byte {
	dm_build_1136 := Dm_build_652.Dm_build_745(dm_build_1135.dm_build_1012, dm_build_1135.dm_build_1013)
	dm_build_1135.dm_build_1013++
	return dm_build_1136
}

func (dm_build_1138 *Dm_build_1011) Dm_build_1137() int16 {
	dm_build_1139 := Dm_build_652.Dm_build_749(dm_build_1138.dm_build_1012, dm_build_1138.dm_build_1013)
	dm_build_1138.dm_build_1013 += 2
	return dm_build_1139
}

func (dm_build_1141 *Dm_build_1011) Dm_build_1140() int32 {
	dm_build_1142 := Dm_build_652.Dm_build_754(dm_build_1141.dm_build_1012, dm_build_1141.dm_build_1013)
	dm_build_1141.dm_build_1013 += 4
	return dm_build_1142
}

func (dm_build_1144 *Dm_build_1011) Dm_build_1143() int64 {
	dm_build_1145 := Dm_build_652.Dm_build_759(dm_build_1144.dm_build_1012, dm_build_1144.dm_build_1013)
	dm_build_1144.dm_build_1013 += 8
	return dm_build_1145
}

func (dm_build_1147 *Dm_build_1011) Dm_build_1146() float32 {
	dm_build_1148 := Dm_build_652.Dm_build_764(dm_build_1147.dm_build_1012, dm_build_1147.dm_build_1013)
	dm_build_1147.dm_build_1013 += 4
	return dm_build_1148
}

func (dm_build_1150 *Dm_build_1011) Dm_build_1149() float64 {
	dm_build_1151 := Dm_build_652.Dm_build_768(dm_build_1150.dm_build_1012, dm_build_1150.dm_build_1013)
	dm_build_1150.dm_build_1013 += 8
	return dm_build_1151
}

func (dm_build_1153 *Dm_build_1011) Dm_build_1152() uint8 {
	dm_build_1154 := Dm_build_652.Dm_build_772(dm_build_1153.dm_build_1012, dm_build_1153.dm_build_1013)
	dm_build_1153.dm_build_1013 += 1
	return dm_build_1154
}

func (dm_build_1156 *Dm_build_1011) Dm_build_1155() uint16 {
	dm_build_1157 := Dm_build_652.Dm_build_776(dm_build_1156.dm_build_1012, dm_build_1156.dm_build_1013)
	dm_build_1156.dm_build_1013 += 2
	return dm_build_1157
}

func (dm_build_1159 *Dm_build_1011) Dm_build_1158() uint32 {
	dm_build_1160 := Dm_build_652.Dm_build_781(dm_build_1159.dm_build_1012, dm_build_1159.dm_build_1013)
	dm_build_1159.dm_build_1013 += 4
	return dm_build_1160
}

func (dm_build_1162 *Dm_build_1011) Dm_build_1161(dm_build_1163 int) []byte {
	dm_build_1164 := Dm_build_652.Dm_build_803(dm_build_1162.dm_build_1012, dm_build_1162.dm_build_1013, dm_build_1163)
	dm_build_1162.dm_build_1013 += dm_build_1163
	return dm_build_1164
}

func (dm_build_1166 *Dm_build_1011) Dm_build_1165() []byte {
	return dm_build_1166.Dm_build_1161(int(dm_build_1166.Dm_build_1140()))
}

func (dm_build_1168 *Dm_build_1011) Dm_build_1167() []byte {
	return dm_build_1168.Dm_build_1161(int(dm_build_1168.Dm_build_1134()))
}

func (dm_build_1170 *Dm_build_1011) Dm_build_1169() []byte {
	return dm_build_1170.Dm_build_1161(int(dm_build_1170.Dm_build_1137()))
}

func (dm_build_1172 *Dm_build_1011) Dm_build_1171(dm_build_1173 int) []byte {
	return dm_build_1172.Dm_build_1161(dm_build_1173)
}

func (dm_build_1175 *Dm_build_1011) Dm_build_1174() []byte {
	dm_build_1176 := 0
	for dm_build_1175.Dm_build_1134() != 0 {
		dm_build_1176++
	}
	dm_build_1175.Dm_build_1036(dm_build_1176, false, false)
	return dm_build_1175.Dm_build_1161(dm_build_1176)
}

func (dm_build_1178 *Dm_build_1011) Dm_build_1177(dm_build_1179 int, dm_build_1180 string, dm_build_1181 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1178.Dm_build_1161(dm_build_1179), dm_build_1180, dm_build_1181)
}

func (dm_build_1183 *Dm_build_1011) Dm_build_1182(dm_build_1184 string, dm_build_1185 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1183.Dm_build_1165(), dm_build_1184, dm_build_1185)
}

func (dm_build_1187 *Dm_build_1011) Dm_build_1186(dm_build_1188 string, dm_build_1189 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1187.Dm_build_1167(), dm_build_1188, dm_build_1189)
}

func (dm_build_1191 *Dm_build_1011) Dm_build_1190(dm_build_1192 string, dm_build_1193 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1191.Dm_build_1169(), dm_build_1192, dm_build_1193)
}

func (dm_build_1195 *Dm_build_1011) Dm_build_1194(dm_build_1196 string, dm_build_1197 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1195.Dm_build_1174(), dm_build_1196, dm_build_1197)
}

func (dm_build_1199 *Dm_build_1011) Dm_build_1198(dm_build_1200 int, dm_build_1201 byte) int {
	return dm_build_1199.Dm_build_1234(dm_build_1200, Dm_build_652.Dm_build_830(dm_build_1201))
}

func (dm_build_1203 *Dm_build_1011) Dm_build_1202(dm_build_1204 int, dm_build_1205 int16) int {
	return dm_build_1203.Dm_build_1234(dm_build_1204, Dm_build_652.Dm_build_836(dm_build_1205))
}

func (dm_build_1207 *Dm_build_1011) Dm_build_1206(dm_build_1208 int, dm_build_1209 int32) int {
	return dm_build_1207.Dm_build_1234(dm_build_1208, Dm_build_652.Dm_build_839(dm_build_1209))
}

func (dm_build_1211 *Dm_build_1011) Dm_build_1210(dm_build_1212 int, dm_build_1213 int64) int {
	return dm_build_1211.Dm_build_1234(dm_build_1212, Dm_build_652.Dm_build_842(dm_build_1213))
}

func (dm_build_1215 *Dm_build_1011) Dm_build_1214(dm_build_1216 int, dm_build_1217 float32) int {
	return dm_build_1215.Dm_build_1234(dm_build_1216, Dm_build_652.Dm_build_845(dm_build_1217))
}

func (dm_build_1219 *Dm_build_1011) Dm_build_1218(dm_build_1220 int, dm_build_1221 float64) int {
	return dm_build_1219.Dm_build_1234(dm_build_1220, Dm_build_652.Dm_build_848(dm_build_1221))
}

func (dm_build_1223 *Dm_build_1011) Dm_build_1222(dm_build_1224 int, dm_build_1225 uint8) int {
	return dm_build_1223.Dm_build_1234(dm_build_1224, Dm_build_652.Dm_build_851(dm_build_1225))
}

func (dm_build_1227 *Dm_build_1011) Dm_build_1226(dm_build_1228 int, dm_build_1229 uint16) int {
	return dm_build_1227.Dm_build_1234(dm_build_1228, Dm_build_652.Dm_build_854(dm_build_1229))
}

func (dm_build_1231 *Dm_build_1011) Dm_build_1230(dm_build_1232 int, dm_build_1233 uint32) int {
	return dm_build_1231.Dm_build_1234(dm_build_1232, Dm_build_652.Dm_build_857(dm_build_1233))
}

func (dm_build_1235 *Dm_build_1011) Dm_build_1234(dm_build_1236 int, dm_build_1237 []byte) int {
	return copy(dm_build_1235.dm_build_1012[dm_build_1236:], dm_build_1237)
}

func (dm_build_1239 *Dm_build_1011) Dm_build_1238(dm_build_1240 int, dm_build_1241 []byte) int {
	return dm_build_1239.Dm_build_1206(dm_build_1240, int32(len(dm_build_1241))) + dm_build_1239.Dm_build_1234(dm_build_1240+4, dm_build_1241)
}

func (dm_build_1243 *Dm_build_1011) Dm_build_1242(dm_build_1244 int, dm_build_1245 []byte) int {
	return dm_build_1243.Dm_build_1198(dm_build_1244, byte(len(dm_build_1245))) + dm_build_1243.Dm_build_1234(dm_build_1244+1, dm_build_1245)
}

func (dm_build_1247 *Dm_build_1011) Dm_build_1246(dm_build_1248 int, dm_build_1249 []byte) int {
	return dm_build_1247.Dm_build_1202(dm_build_1248, int16(len(dm_build_1249))) + dm_build_1247.Dm_build_1234(dm_build_1248+2, dm_build_1249)
}

func (dm_build_1251 *Dm_build_1011) Dm_build_1250(dm_build_1252 int, dm_build_1253 []byte) int {
	return dm_build_1251.Dm_build_1234(dm_build_1252, dm_build_1253) + dm_build_1251.Dm_build_1198(dm_build_1252+len(dm_build_1253), 0)
}

func (dm_build_1255 *Dm_build_1011) Dm_build_1254(dm_build_1256 int, dm_build_1257 string, dm_build_1258 string, dm_build_1259 *DmConnection) int {
	return dm_build_1255.Dm_build_1238(dm_build_1256, Dm_build_652.Dm_build_868(dm_build_1257, dm_build_1258, dm_build_1259))
}

func (dm_build_1261 *Dm_build_1011) Dm_build_1260(dm_build_1262 int, dm_build_1263 string, dm_build_1264 string, dm_build_1265 *DmConnection) int {
	return dm_build_1261.Dm_build_1242(dm_build_1262, Dm_build_652.Dm_build_868(dm_build_1263, dm_build_1264, dm_build_1265))
}

func (dm_build_1267 *Dm_build_1011) Dm_build_1266(dm_build_1268 int, dm_build_1269 string, dm_build_1270 string, dm_build_1271 *DmConnection) int {
	return dm_build_1267.Dm_build_1246(dm_build_1268, Dm_build_652.Dm_build_868(dm_build_1269, dm_build_1270, dm_build_1271))
}

func (dm_build_1273 *Dm_build_1011) Dm_build_1272(dm_build_1274 int, dm_build_1275 string, dm_build_1276 string, dm_build_1277 *DmConnection) int {
	return dm_build_1273.Dm_build_1250(dm_build_1274, Dm_build_652.Dm_build_868(dm_build_1275, dm_build_1276, dm_build_1277))
}

func (dm_build_1279 *Dm_build_1011) Dm_build_1278(dm_build_1280 int) byte {
	return Dm_build_652.Dm_build_873(dm_build_1279.Dm_build_1305(dm_build_1280, 1))
}

func (dm_build_1282 *Dm_build_1011) Dm_build_1281(dm_build_1283 int) int16 {
	return Dm_build_652.Dm_build_876(dm_build_1282.Dm_build_1305(dm_build_1283, 2))
}

func (dm_build_1285 *Dm_build_1011) Dm_build_1284(dm_build_1286 int) int32 {
	return Dm_build_652.Dm_build_879(dm_build_1285.Dm_build_1305(dm_build_1286, 4))
}

func (dm_build_1288 *Dm_build_1011) Dm_build_1287(dm_build_1289 int) int64 {
	return Dm_build_652.Dm_build_882(dm_build_1288.Dm_build_1305(dm_build_1289, 8))
}

func (dm_build_1291 *Dm_build_1011) Dm_build_1290(dm_build_1292 int) float32 {
	return Dm_build_652.Dm_build_885(dm_build_1291.Dm_build_1305(dm_build_1292, 4))
}

func (dm_build_1294 *Dm_build_1011) Dm_build_1293(dm_build_1295 int) float64 {
	return Dm_build_652.Dm_build_888(dm_build_1294.Dm_build_1305(dm_build_1295, 8))
}

func (dm_build_1297 *Dm_build_1011) Dm_build_1296(dm_build_1298 int) uint8 {
	return Dm_build_652.Dm_build_891(dm_build_1297.Dm_build_1305(dm_build_1298, 1))
}

func (dm_build_1300 *Dm_build_1011) Dm_build_1299(dm_build_1301 int) uint16 {
	return Dm_build_652.Dm_build_894(dm_build_1300.Dm_build_1305(dm_build_1301, 2))
}

func (dm_build_1303 *Dm_build_1011) Dm_build_1302(dm_build_1304 int) uint32 {
	return Dm_build_652.Dm_build_897(dm_build_1303.Dm_build_1305(dm_build_1304, 4))
}

func (dm_build_1306 *Dm_build_1011) Dm_build_1305(dm_build_1307 int, dm_build_1308 int) []byte {
	return dm_build_1306.dm_build_1012[dm_build_1307 : dm_build_1307+dm_build_1308]
}

func (dm_build_1310 *Dm_build_1011) Dm_build_1309(dm_build_1311 int) []byte {
	dm_build_1312 := dm_build_1310.Dm_build_1284(dm_build_1311)
	return dm_build_1310.Dm_build_1305(dm_build_1311+4, int(dm_build_1312))
}

func (dm_build_1314 *Dm_build_1011) Dm_build_1313(dm_build_1315 int) []byte {
	dm_build_1316 := dm_build_1314.Dm_build_1278(dm_build_1315)
	return dm_build_1314.Dm_build_1305(dm_build_1315+1, int(dm_build_1316))
}

func (dm_build_1318 *Dm_build_1011) Dm_build_1317(dm_build_1319 int) []byte {
	dm_build_1320 := dm_build_1318.Dm_build_1281(dm_build_1319)
	return dm_build_1318.Dm_build_1305(dm_build_1319+2, int(dm_build_1320))
}

func (dm_build_1322 *Dm_build_1011) Dm_build_1321(dm_build_1323 int) []byte {
	dm_build_1324 := 0
	for dm_build_1322.Dm_build_1278(dm_build_1323) != 0 {
		dm_build_1323++
		dm_build_1324++
	}

	return dm_build_1322.Dm_build_1305(dm_build_1323-dm_build_1324, int(dm_build_1324))
}

func (dm_build_1326 *Dm_build_1011) Dm_build_1325(dm_build_1327 int, dm_build_1328 string, dm_build_1329 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1326.Dm_build_1309(dm_build_1327), dm_build_1328, dm_build_1329)
}

func (dm_build_1331 *Dm_build_1011) Dm_build_1330(dm_build_1332 int, dm_build_1333 string, dm_build_1334 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1331.Dm_build_1313(dm_build_1332), dm_build_1333, dm_build_1334)
}

func (dm_build_1336 *Dm_build_1011) Dm_build_1335(dm_build_1337 int, dm_build_1338 string, dm_build_1339 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1336.Dm_build_1317(dm_build_1337), dm_build_1338, dm_build_1339)
}

func (dm_build_1341 *Dm_build_1011) Dm_build_1340(dm_build_1342 int, dm_build_1343 string, dm_build_1344 *DmConnection) string {
	return Dm_build_652.Dm_build_904(dm_build_1341.Dm_build_1321(dm_build_1342), dm_build_1343, dm_build_1344)
}
