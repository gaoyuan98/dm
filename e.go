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

type dm_build_1267 struct{}

var Dm_build_1268 = &dm_build_1267{}

func (Dm_build_1270 *dm_build_1267) Dm_build_1269(dm_build_1271 []byte, dm_build_1272 int, dm_build_1273 byte) int {
	dm_build_1271[dm_build_1272] = dm_build_1273
	return 1
}

func (Dm_build_1275 *dm_build_1267) Dm_build_1274(dm_build_1276 []byte, dm_build_1277 int, dm_build_1278 int8) int {
	dm_build_1276[dm_build_1277] = byte(dm_build_1278)
	return 1
}

func (Dm_build_1280 *dm_build_1267) Dm_build_1279(dm_build_1281 []byte, dm_build_1282 int, dm_build_1283 int16) int {
	dm_build_1281[dm_build_1282] = byte(dm_build_1283)
	dm_build_1282++
	dm_build_1281[dm_build_1282] = byte(dm_build_1283 >> 8)
	return 2
}

func (Dm_build_1285 *dm_build_1267) Dm_build_1284(dm_build_1286 []byte, dm_build_1287 int, dm_build_1288 int32) int {
	dm_build_1286[dm_build_1287] = byte(dm_build_1288)
	dm_build_1287++
	dm_build_1286[dm_build_1287] = byte(dm_build_1288 >> 8)
	dm_build_1287++
	dm_build_1286[dm_build_1287] = byte(dm_build_1288 >> 16)
	dm_build_1287++
	dm_build_1286[dm_build_1287] = byte(dm_build_1288 >> 24)
	dm_build_1287++
	return 4
}

func (Dm_build_1290 *dm_build_1267) Dm_build_1289(dm_build_1291 []byte, dm_build_1292 int, dm_build_1293 int64) int {
	dm_build_1291[dm_build_1292] = byte(dm_build_1293)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 8)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 16)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 24)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 32)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 40)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 48)
	dm_build_1292++
	dm_build_1291[dm_build_1292] = byte(dm_build_1293 >> 56)
	return 8
}

func (Dm_build_1295 *dm_build_1267) Dm_build_1294(dm_build_1296 []byte, dm_build_1297 int, dm_build_1298 float32) int {
	return Dm_build_1295.Dm_build_1314(dm_build_1296, dm_build_1297, math.Float32bits(dm_build_1298))
}

func (Dm_build_1300 *dm_build_1267) Dm_build_1299(dm_build_1301 []byte, dm_build_1302 int, dm_build_1303 float64) int {
	return Dm_build_1300.Dm_build_1319(dm_build_1301, dm_build_1302, math.Float64bits(dm_build_1303))
}

func (Dm_build_1305 *dm_build_1267) Dm_build_1304(dm_build_1306 []byte, dm_build_1307 int, dm_build_1308 uint8) int {
	dm_build_1306[dm_build_1307] = byte(dm_build_1308)
	return 1
}

func (Dm_build_1310 *dm_build_1267) Dm_build_1309(dm_build_1311 []byte, dm_build_1312 int, dm_build_1313 uint16) int {
	dm_build_1311[dm_build_1312] = byte(dm_build_1313)
	dm_build_1312++
	dm_build_1311[dm_build_1312] = byte(dm_build_1313 >> 8)
	return 2
}

func (Dm_build_1315 *dm_build_1267) Dm_build_1314(dm_build_1316 []byte, dm_build_1317 int, dm_build_1318 uint32) int {
	dm_build_1316[dm_build_1317] = byte(dm_build_1318)
	dm_build_1317++
	dm_build_1316[dm_build_1317] = byte(dm_build_1318 >> 8)
	dm_build_1317++
	dm_build_1316[dm_build_1317] = byte(dm_build_1318 >> 16)
	dm_build_1317++
	dm_build_1316[dm_build_1317] = byte(dm_build_1318 >> 24)
	return 3
}

func (Dm_build_1320 *dm_build_1267) Dm_build_1319(dm_build_1321 []byte, dm_build_1322 int, dm_build_1323 uint64) int {
	dm_build_1321[dm_build_1322] = byte(dm_build_1323)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 8)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 16)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 24)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 32)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 40)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 48)
	dm_build_1322++
	dm_build_1321[dm_build_1322] = byte(dm_build_1323 >> 56)
	return 3
}

func (Dm_build_1325 *dm_build_1267) Dm_build_1324(dm_build_1326 []byte, dm_build_1327 int, dm_build_1328 []byte, dm_build_1329 int, dm_build_1330 int) int {
	copy(dm_build_1326[dm_build_1327:dm_build_1327+dm_build_1330], dm_build_1328[dm_build_1329:dm_build_1329+dm_build_1330])
	return dm_build_1330
}

func (Dm_build_1332 *dm_build_1267) Dm_build_1331(dm_build_1333 []byte, dm_build_1334 int, dm_build_1335 []byte, dm_build_1336 int, dm_build_1337 int) int {
	dm_build_1334 += Dm_build_1332.Dm_build_1314(dm_build_1333, dm_build_1334, uint32(dm_build_1337))
	return 4 + Dm_build_1332.Dm_build_1324(dm_build_1333, dm_build_1334, dm_build_1335, dm_build_1336, dm_build_1337)
}

func (Dm_build_1339 *dm_build_1267) Dm_build_1338(dm_build_1340 []byte, dm_build_1341 int, dm_build_1342 []byte, dm_build_1343 int, dm_build_1344 int) int {
	dm_build_1341 += Dm_build_1339.Dm_build_1309(dm_build_1340, dm_build_1341, uint16(dm_build_1344))
	return 2 + Dm_build_1339.Dm_build_1324(dm_build_1340, dm_build_1341, dm_build_1342, dm_build_1343, dm_build_1344)
}

func (Dm_build_1346 *dm_build_1267) Dm_build_1345(dm_build_1347 []byte, dm_build_1348 int, dm_build_1349 string, dm_build_1350 string, dm_build_1351 *DmConnection) int {
	dm_build_1352 := Dm_build_1346.Dm_build_1484(dm_build_1349, dm_build_1350, dm_build_1351)
	dm_build_1348 += Dm_build_1346.Dm_build_1314(dm_build_1347, dm_build_1348, uint32(len(dm_build_1352)))
	return 4 + Dm_build_1346.Dm_build_1324(dm_build_1347, dm_build_1348, dm_build_1352, 0, len(dm_build_1352))
}

func (Dm_build_1354 *dm_build_1267) Dm_build_1353(dm_build_1355 []byte, dm_build_1356 int, dm_build_1357 string, dm_build_1358 string, dm_build_1359 *DmConnection) int {
	dm_build_1360 := Dm_build_1354.Dm_build_1484(dm_build_1357, dm_build_1358, dm_build_1359)

	dm_build_1356 += Dm_build_1354.Dm_build_1309(dm_build_1355, dm_build_1356, uint16(len(dm_build_1360)))
	return 2 + Dm_build_1354.Dm_build_1324(dm_build_1355, dm_build_1356, dm_build_1360, 0, len(dm_build_1360))
}

func (Dm_build_1362 *dm_build_1267) Dm_build_1361(dm_build_1363 []byte, dm_build_1364 int) byte {
	return dm_build_1363[dm_build_1364]
}

func (Dm_build_1366 *dm_build_1267) Dm_build_1365(dm_build_1367 []byte, dm_build_1368 int) int16 {
	var dm_build_1369 int16
	dm_build_1369 = int16(dm_build_1367[dm_build_1368] & 0xff)
	dm_build_1368++
	dm_build_1369 |= int16(dm_build_1367[dm_build_1368]&0xff) << 8
	return dm_build_1369
}

func (Dm_build_1371 *dm_build_1267) Dm_build_1370(dm_build_1372 []byte, dm_build_1373 int) int32 {
	var dm_build_1374 int32
	dm_build_1374 = int32(dm_build_1372[dm_build_1373] & 0xff)
	dm_build_1373++
	dm_build_1374 |= int32(dm_build_1372[dm_build_1373]&0xff) << 8
	dm_build_1373++
	dm_build_1374 |= int32(dm_build_1372[dm_build_1373]&0xff) << 16
	dm_build_1373++
	dm_build_1374 |= int32(dm_build_1372[dm_build_1373]&0xff) << 24
	return dm_build_1374
}

func (Dm_build_1376 *dm_build_1267) Dm_build_1375(dm_build_1377 []byte, dm_build_1378 int) int64 {
	var dm_build_1379 int64
	dm_build_1379 = int64(dm_build_1377[dm_build_1378] & 0xff)
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 8
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 16
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 24
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 32
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 40
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 48
	dm_build_1378++
	dm_build_1379 |= int64(dm_build_1377[dm_build_1378]&0xff) << 56
	return dm_build_1379
}

func (Dm_build_1381 *dm_build_1267) Dm_build_1380(dm_build_1382 []byte, dm_build_1383 int) float32 {
	return math.Float32frombits(Dm_build_1381.Dm_build_1397(dm_build_1382, dm_build_1383))
}

func (Dm_build_1385 *dm_build_1267) Dm_build_1384(dm_build_1386 []byte, dm_build_1387 int) float64 {
	return math.Float64frombits(Dm_build_1385.Dm_build_1402(dm_build_1386, dm_build_1387))
}

func (Dm_build_1389 *dm_build_1267) Dm_build_1388(dm_build_1390 []byte, dm_build_1391 int) uint8 {
	return uint8(dm_build_1390[dm_build_1391] & 0xff)
}

func (Dm_build_1393 *dm_build_1267) Dm_build_1392(dm_build_1394 []byte, dm_build_1395 int) uint16 {
	var dm_build_1396 uint16
	dm_build_1396 = uint16(dm_build_1394[dm_build_1395] & 0xff)
	dm_build_1395++
	dm_build_1396 |= uint16(dm_build_1394[dm_build_1395]&0xff) << 8
	return dm_build_1396
}

func (Dm_build_1398 *dm_build_1267) Dm_build_1397(dm_build_1399 []byte, dm_build_1400 int) uint32 {
	var dm_build_1401 uint32
	dm_build_1401 = uint32(dm_build_1399[dm_build_1400] & 0xff)
	dm_build_1400++
	dm_build_1401 |= uint32(dm_build_1399[dm_build_1400]&0xff) << 8
	dm_build_1400++
	dm_build_1401 |= uint32(dm_build_1399[dm_build_1400]&0xff) << 16
	dm_build_1400++
	dm_build_1401 |= uint32(dm_build_1399[dm_build_1400]&0xff) << 24
	return dm_build_1401
}

func (Dm_build_1403 *dm_build_1267) Dm_build_1402(dm_build_1404 []byte, dm_build_1405 int) uint64 {
	var dm_build_1406 uint64
	dm_build_1406 = uint64(dm_build_1404[dm_build_1405] & 0xff)
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 8
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 16
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 24
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 32
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 40
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 48
	dm_build_1405++
	dm_build_1406 |= uint64(dm_build_1404[dm_build_1405]&0xff) << 56
	return dm_build_1406
}

func (Dm_build_1408 *dm_build_1267) Dm_build_1407(dm_build_1409 []byte, dm_build_1410 int) []byte {
	dm_build_1411 := Dm_build_1408.Dm_build_1397(dm_build_1409, dm_build_1410)

	dm_build_1412 := make([]byte, dm_build_1411)
	copy(dm_build_1412[:int(dm_build_1411)], dm_build_1409[dm_build_1410+4:dm_build_1410+4+int(dm_build_1411)])
	return dm_build_1412
}

func (Dm_build_1414 *dm_build_1267) Dm_build_1413(dm_build_1415 []byte, dm_build_1416 int) []byte {
	dm_build_1417 := Dm_build_1414.Dm_build_1392(dm_build_1415, dm_build_1416)

	dm_build_1418 := make([]byte, dm_build_1417)
	copy(dm_build_1418[:int(dm_build_1417)], dm_build_1415[dm_build_1416+2:dm_build_1416+2+int(dm_build_1417)])
	return dm_build_1418
}

func (Dm_build_1420 *dm_build_1267) Dm_build_1419(dm_build_1421 []byte, dm_build_1422 int, dm_build_1423 int) []byte {

	dm_build_1424 := make([]byte, dm_build_1423)
	copy(dm_build_1424[:dm_build_1423], dm_build_1421[dm_build_1422:dm_build_1422+dm_build_1423])
	return dm_build_1424
}

func (Dm_build_1426 *dm_build_1267) Dm_build_1425(dm_build_1427 []byte, dm_build_1428 int, dm_build_1429 int, dm_build_1430 string, dm_build_1431 *DmConnection) string {
	return Dm_build_1426.Dm_build_1520(dm_build_1427[dm_build_1428:dm_build_1428+dm_build_1429], dm_build_1430, dm_build_1431)
}

func (Dm_build_1433 *dm_build_1267) Dm_build_1432(dm_build_1434 []byte, dm_build_1435 int, dm_build_1436 string, dm_build_1437 *DmConnection) string {
	dm_build_1438 := Dm_build_1433.Dm_build_1397(dm_build_1434, dm_build_1435)
	dm_build_1435 += 4
	return Dm_build_1433.Dm_build_1425(dm_build_1434, dm_build_1435, int(dm_build_1438), dm_build_1436, dm_build_1437)
}

func (Dm_build_1440 *dm_build_1267) Dm_build_1439(dm_build_1441 []byte, dm_build_1442 int, dm_build_1443 string, dm_build_1444 *DmConnection) string {
	dm_build_1445 := Dm_build_1440.Dm_build_1392(dm_build_1441, dm_build_1442)
	dm_build_1442 += 2
	return Dm_build_1440.Dm_build_1425(dm_build_1441, dm_build_1442, int(dm_build_1445), dm_build_1443, dm_build_1444)
}

func (Dm_build_1447 *dm_build_1267) Dm_build_1446(dm_build_1448 byte) []byte {
	return []byte{dm_build_1448}
}

func (Dm_build_1450 *dm_build_1267) Dm_build_1449(dm_build_1451 int8) []byte {
	return []byte{byte(dm_build_1451)}
}

func (Dm_build_1453 *dm_build_1267) Dm_build_1452(dm_build_1454 int16) []byte {
	return []byte{byte(dm_build_1454), byte(dm_build_1454 >> 8)}
}

func (Dm_build_1456 *dm_build_1267) Dm_build_1455(dm_build_1457 int32) []byte {
	return []byte{byte(dm_build_1457), byte(dm_build_1457 >> 8), byte(dm_build_1457 >> 16), byte(dm_build_1457 >> 24)}
}

func (Dm_build_1459 *dm_build_1267) Dm_build_1458(dm_build_1460 int64) []byte {
	return []byte{byte(dm_build_1460), byte(dm_build_1460 >> 8), byte(dm_build_1460 >> 16), byte(dm_build_1460 >> 24), byte(dm_build_1460 >> 32),
		byte(dm_build_1460 >> 40), byte(dm_build_1460 >> 48), byte(dm_build_1460 >> 56)}
}

func (Dm_build_1462 *dm_build_1267) Dm_build_1461(dm_build_1463 float32) []byte {
	return Dm_build_1462.Dm_build_1473(math.Float32bits(dm_build_1463))
}

func (Dm_build_1465 *dm_build_1267) Dm_build_1464(dm_build_1466 float64) []byte {
	return Dm_build_1465.Dm_build_1476(math.Float64bits(dm_build_1466))
}

func (Dm_build_1468 *dm_build_1267) Dm_build_1467(dm_build_1469 uint8) []byte {
	return []byte{byte(dm_build_1469)}
}

func (Dm_build_1471 *dm_build_1267) Dm_build_1470(dm_build_1472 uint16) []byte {
	return []byte{byte(dm_build_1472), byte(dm_build_1472 >> 8)}
}

func (Dm_build_1474 *dm_build_1267) Dm_build_1473(dm_build_1475 uint32) []byte {
	return []byte{byte(dm_build_1475), byte(dm_build_1475 >> 8), byte(dm_build_1475 >> 16), byte(dm_build_1475 >> 24)}
}

func (Dm_build_1477 *dm_build_1267) Dm_build_1476(dm_build_1478 uint64) []byte {
	return []byte{byte(dm_build_1478), byte(dm_build_1478 >> 8), byte(dm_build_1478 >> 16), byte(dm_build_1478 >> 24), byte(dm_build_1478 >> 32), byte(dm_build_1478 >> 40), byte(dm_build_1478 >> 48), byte(dm_build_1478 >> 56)}
}

func (Dm_build_1480 *dm_build_1267) Dm_build_1479(dm_build_1481 []byte, dm_build_1482 string, dm_build_1483 *DmConnection) []byte {
	if dm_build_1482 == "UTF-8" {
		return dm_build_1481
	}

	if dm_build_1483 == nil {
		if e := dm_build_1525(dm_build_1482); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1481), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1483.encodeBuffer == nil {
		dm_build_1483.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1483.encode = dm_build_1525(dm_build_1483.getServerEncoding())
		dm_build_1483.transformReaderDst = make([]byte, 4096)
		dm_build_1483.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1483.encode; e != nil {

		dm_build_1483.encodeBuffer.Reset()

		n, err := dm_build_1483.encodeBuffer.ReadFrom(
			Dm_build_1539(bytes.NewReader(dm_build_1481), e.NewEncoder(), dm_build_1483.transformReaderDst, dm_build_1483.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1483.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1485 *dm_build_1267) Dm_build_1484(dm_build_1486 string, dm_build_1487 string, dm_build_1488 *DmConnection) []byte {
	return Dm_build_1485.Dm_build_1479([]byte(dm_build_1486), dm_build_1487, dm_build_1488)
}

func (Dm_build_1490 *dm_build_1267) Dm_build_1489(dm_build_1491 []byte) byte {
	return Dm_build_1490.Dm_build_1361(dm_build_1491, 0)
}

func (Dm_build_1493 *dm_build_1267) Dm_build_1492(dm_build_1494 []byte) int16 {
	return Dm_build_1493.Dm_build_1365(dm_build_1494, 0)
}

func (Dm_build_1496 *dm_build_1267) Dm_build_1495(dm_build_1497 []byte) int32 {
	return Dm_build_1496.Dm_build_1370(dm_build_1497, 0)
}

func (Dm_build_1499 *dm_build_1267) Dm_build_1498(dm_build_1500 []byte) int64 {
	return Dm_build_1499.Dm_build_1375(dm_build_1500, 0)
}

func (Dm_build_1502 *dm_build_1267) Dm_build_1501(dm_build_1503 []byte) float32 {
	return Dm_build_1502.Dm_build_1380(dm_build_1503, 0)
}

func (Dm_build_1505 *dm_build_1267) Dm_build_1504(dm_build_1506 []byte) float64 {
	return Dm_build_1505.Dm_build_1384(dm_build_1506, 0)
}

func (Dm_build_1508 *dm_build_1267) Dm_build_1507(dm_build_1509 []byte) uint8 {
	return Dm_build_1508.Dm_build_1388(dm_build_1509, 0)
}

func (Dm_build_1511 *dm_build_1267) Dm_build_1510(dm_build_1512 []byte) uint16 {
	return Dm_build_1511.Dm_build_1392(dm_build_1512, 0)
}

func (Dm_build_1514 *dm_build_1267) Dm_build_1513(dm_build_1515 []byte) uint32 {
	return Dm_build_1514.Dm_build_1397(dm_build_1515, 0)
}

func (Dm_build_1517 *dm_build_1267) Dm_build_1516(dm_build_1518 []byte, dm_build_1519 string) []byte {
	if dm_build_1519 == "UTF-8" {
		return dm_build_1518
	}

	if e := dm_build_1525(dm_build_1519); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1518), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1521 *dm_build_1267) Dm_build_1520(dm_build_1522 []byte, dm_build_1523 string, dm_build_1524 *DmConnection) string {
	return string(Dm_build_1521.Dm_build_1516(dm_build_1522, dm_build_1523))
}

func dm_build_1525(dm_build_1526 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1526); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1527 struct {
	dm_build_1528 io.Reader
	dm_build_1529 transform.Transformer
	dm_build_1530 error

	dm_build_1531                []byte
	dm_build_1532, dm_build_1533 int

	dm_build_1534                []byte
	dm_build_1535, dm_build_1536 int

	dm_build_1537 bool
}

const dm_build_1538 = 4096

func Dm_build_1539(dm_build_1540 io.Reader, dm_build_1541 transform.Transformer, dm_build_1542 []byte, dm_build_1543 []byte) *Dm_build_1527 {
	dm_build_1541.Reset()
	return &Dm_build_1527{
		dm_build_1528: dm_build_1540,
		dm_build_1529: dm_build_1541,
		dm_build_1531: dm_build_1542,
		dm_build_1534: dm_build_1543,
	}
}

func (dm_build_1545 *Dm_build_1527) Read(dm_build_1546 []byte) (int, error) {
	dm_build_1547, dm_build_1548 := 0, error(nil)
	for {

		if dm_build_1545.dm_build_1532 != dm_build_1545.dm_build_1533 {
			dm_build_1547 = copy(dm_build_1546, dm_build_1545.dm_build_1531[dm_build_1545.dm_build_1532:dm_build_1545.dm_build_1533])
			dm_build_1545.dm_build_1532 += dm_build_1547
			if dm_build_1545.dm_build_1532 == dm_build_1545.dm_build_1533 && dm_build_1545.dm_build_1537 {
				return dm_build_1547, dm_build_1545.dm_build_1530
			}
			return dm_build_1547, nil
		} else if dm_build_1545.dm_build_1537 {
			return 0, dm_build_1545.dm_build_1530
		}

		if dm_build_1545.dm_build_1535 != dm_build_1545.dm_build_1536 || dm_build_1545.dm_build_1530 != nil {
			dm_build_1545.dm_build_1532 = 0
			dm_build_1545.dm_build_1533, dm_build_1547, dm_build_1548 = dm_build_1545.dm_build_1529.Transform(dm_build_1545.dm_build_1531, dm_build_1545.dm_build_1534[dm_build_1545.dm_build_1535:dm_build_1545.dm_build_1536], dm_build_1545.dm_build_1530 == io.EOF)
			dm_build_1545.dm_build_1535 += dm_build_1547

			switch {
			case dm_build_1548 == nil:
				if dm_build_1545.dm_build_1535 != dm_build_1545.dm_build_1536 {
					dm_build_1545.dm_build_1530 = nil
				}

				dm_build_1545.dm_build_1537 = dm_build_1545.dm_build_1530 != nil
				continue
			case dm_build_1548 == transform.ErrShortDst && (dm_build_1545.dm_build_1533 != 0 || dm_build_1547 != 0):

				continue
			case dm_build_1548 == transform.ErrShortSrc && dm_build_1545.dm_build_1536-dm_build_1545.dm_build_1535 != len(dm_build_1545.dm_build_1534) && dm_build_1545.dm_build_1530 == nil:

			default:
				dm_build_1545.dm_build_1537 = true

				if dm_build_1545.dm_build_1530 == nil || dm_build_1545.dm_build_1530 == io.EOF {
					dm_build_1545.dm_build_1530 = dm_build_1548
				}
				continue
			}
		}

		if dm_build_1545.dm_build_1535 != 0 {
			dm_build_1545.dm_build_1535, dm_build_1545.dm_build_1536 = 0, copy(dm_build_1545.dm_build_1534, dm_build_1545.dm_build_1534[dm_build_1545.dm_build_1535:dm_build_1545.dm_build_1536])
		}
		dm_build_1547, dm_build_1545.dm_build_1530 = dm_build_1545.dm_build_1528.Read(dm_build_1545.dm_build_1534[dm_build_1545.dm_build_1536:])
		dm_build_1545.dm_build_1536 += dm_build_1547
	}
}
