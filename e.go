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

type dm_build_1331 struct{}

var Dm_build_1332 = &dm_build_1331{}

func (Dm_build_1334 *dm_build_1331) Dm_build_1333(dm_build_1335 []byte, dm_build_1336 int, dm_build_1337 byte) int {
	dm_build_1335[dm_build_1336] = dm_build_1337
	return 1
}

func (Dm_build_1339 *dm_build_1331) Dm_build_1338(dm_build_1340 []byte, dm_build_1341 int, dm_build_1342 int8) int {
	dm_build_1340[dm_build_1341] = byte(dm_build_1342)
	return 1
}

func (Dm_build_1344 *dm_build_1331) Dm_build_1343(dm_build_1345 []byte, dm_build_1346 int, dm_build_1347 int16) int {
	dm_build_1345[dm_build_1346] = byte(dm_build_1347)
	dm_build_1346++
	dm_build_1345[dm_build_1346] = byte(dm_build_1347 >> 8)
	return 2
}

func (Dm_build_1349 *dm_build_1331) Dm_build_1348(dm_build_1350 []byte, dm_build_1351 int, dm_build_1352 int32) int {
	dm_build_1350[dm_build_1351] = byte(dm_build_1352)
	dm_build_1351++
	dm_build_1350[dm_build_1351] = byte(dm_build_1352 >> 8)
	dm_build_1351++
	dm_build_1350[dm_build_1351] = byte(dm_build_1352 >> 16)
	dm_build_1351++
	dm_build_1350[dm_build_1351] = byte(dm_build_1352 >> 24)
	dm_build_1351++
	return 4
}

func (Dm_build_1354 *dm_build_1331) Dm_build_1353(dm_build_1355 []byte, dm_build_1356 int, dm_build_1357 int64) int {
	dm_build_1355[dm_build_1356] = byte(dm_build_1357)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 8)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 16)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 24)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 32)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 40)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 48)
	dm_build_1356++
	dm_build_1355[dm_build_1356] = byte(dm_build_1357 >> 56)
	return 8
}

func (Dm_build_1359 *dm_build_1331) Dm_build_1358(dm_build_1360 []byte, dm_build_1361 int, dm_build_1362 float32) int {
	return Dm_build_1359.Dm_build_1378(dm_build_1360, dm_build_1361, math.Float32bits(dm_build_1362))
}

func (Dm_build_1364 *dm_build_1331) Dm_build_1363(dm_build_1365 []byte, dm_build_1366 int, dm_build_1367 float64) int {
	return Dm_build_1364.Dm_build_1383(dm_build_1365, dm_build_1366, math.Float64bits(dm_build_1367))
}

func (Dm_build_1369 *dm_build_1331) Dm_build_1368(dm_build_1370 []byte, dm_build_1371 int, dm_build_1372 uint8) int {
	dm_build_1370[dm_build_1371] = byte(dm_build_1372)
	return 1
}

func (Dm_build_1374 *dm_build_1331) Dm_build_1373(dm_build_1375 []byte, dm_build_1376 int, dm_build_1377 uint16) int {
	dm_build_1375[dm_build_1376] = byte(dm_build_1377)
	dm_build_1376++
	dm_build_1375[dm_build_1376] = byte(dm_build_1377 >> 8)
	return 2
}

func (Dm_build_1379 *dm_build_1331) Dm_build_1378(dm_build_1380 []byte, dm_build_1381 int, dm_build_1382 uint32) int {
	dm_build_1380[dm_build_1381] = byte(dm_build_1382)
	dm_build_1381++
	dm_build_1380[dm_build_1381] = byte(dm_build_1382 >> 8)
	dm_build_1381++
	dm_build_1380[dm_build_1381] = byte(dm_build_1382 >> 16)
	dm_build_1381++
	dm_build_1380[dm_build_1381] = byte(dm_build_1382 >> 24)
	return 3
}

func (Dm_build_1384 *dm_build_1331) Dm_build_1383(dm_build_1385 []byte, dm_build_1386 int, dm_build_1387 uint64) int {
	dm_build_1385[dm_build_1386] = byte(dm_build_1387)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 8)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 16)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 24)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 32)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 40)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 48)
	dm_build_1386++
	dm_build_1385[dm_build_1386] = byte(dm_build_1387 >> 56)
	return 3
}

func (Dm_build_1389 *dm_build_1331) Dm_build_1388(dm_build_1390 []byte, dm_build_1391 int, dm_build_1392 []byte, dm_build_1393 int, dm_build_1394 int) int {
	copy(dm_build_1390[dm_build_1391:dm_build_1391+dm_build_1394], dm_build_1392[dm_build_1393:dm_build_1393+dm_build_1394])
	return dm_build_1394
}

func (Dm_build_1396 *dm_build_1331) Dm_build_1395(dm_build_1397 []byte, dm_build_1398 int, dm_build_1399 []byte, dm_build_1400 int, dm_build_1401 int) int {
	dm_build_1398 += Dm_build_1396.Dm_build_1378(dm_build_1397, dm_build_1398, uint32(dm_build_1401))
	return 4 + Dm_build_1396.Dm_build_1388(dm_build_1397, dm_build_1398, dm_build_1399, dm_build_1400, dm_build_1401)
}

func (Dm_build_1403 *dm_build_1331) Dm_build_1402(dm_build_1404 []byte, dm_build_1405 int, dm_build_1406 []byte, dm_build_1407 int, dm_build_1408 int) int {
	dm_build_1405 += Dm_build_1403.Dm_build_1373(dm_build_1404, dm_build_1405, uint16(dm_build_1408))
	return 2 + Dm_build_1403.Dm_build_1388(dm_build_1404, dm_build_1405, dm_build_1406, dm_build_1407, dm_build_1408)
}

func (Dm_build_1410 *dm_build_1331) Dm_build_1409(dm_build_1411 []byte, dm_build_1412 int, dm_build_1413 string, dm_build_1414 string, dm_build_1415 *DmConnection) int {
	dm_build_1416 := Dm_build_1410.Dm_build_1548(dm_build_1413, dm_build_1414, dm_build_1415)
	dm_build_1412 += Dm_build_1410.Dm_build_1378(dm_build_1411, dm_build_1412, uint32(len(dm_build_1416)))
	return 4 + Dm_build_1410.Dm_build_1388(dm_build_1411, dm_build_1412, dm_build_1416, 0, len(dm_build_1416))
}

func (Dm_build_1418 *dm_build_1331) Dm_build_1417(dm_build_1419 []byte, dm_build_1420 int, dm_build_1421 string, dm_build_1422 string, dm_build_1423 *DmConnection) int {
	dm_build_1424 := Dm_build_1418.Dm_build_1548(dm_build_1421, dm_build_1422, dm_build_1423)

	dm_build_1420 += Dm_build_1418.Dm_build_1373(dm_build_1419, dm_build_1420, uint16(len(dm_build_1424)))
	return 2 + Dm_build_1418.Dm_build_1388(dm_build_1419, dm_build_1420, dm_build_1424, 0, len(dm_build_1424))
}

func (Dm_build_1426 *dm_build_1331) Dm_build_1425(dm_build_1427 []byte, dm_build_1428 int) byte {
	return dm_build_1427[dm_build_1428]
}

func (Dm_build_1430 *dm_build_1331) Dm_build_1429(dm_build_1431 []byte, dm_build_1432 int) int16 {
	var dm_build_1433 int16
	dm_build_1433 = int16(dm_build_1431[dm_build_1432] & 0xff)
	dm_build_1432++
	dm_build_1433 |= int16(dm_build_1431[dm_build_1432]&0xff) << 8
	return dm_build_1433
}

func (Dm_build_1435 *dm_build_1331) Dm_build_1434(dm_build_1436 []byte, dm_build_1437 int) int32 {
	var dm_build_1438 int32
	dm_build_1438 = int32(dm_build_1436[dm_build_1437] & 0xff)
	dm_build_1437++
	dm_build_1438 |= int32(dm_build_1436[dm_build_1437]&0xff) << 8
	dm_build_1437++
	dm_build_1438 |= int32(dm_build_1436[dm_build_1437]&0xff) << 16
	dm_build_1437++
	dm_build_1438 |= int32(dm_build_1436[dm_build_1437]&0xff) << 24
	return dm_build_1438
}

func (Dm_build_1440 *dm_build_1331) Dm_build_1439(dm_build_1441 []byte, dm_build_1442 int) int64 {
	var dm_build_1443 int64
	dm_build_1443 = int64(dm_build_1441[dm_build_1442] & 0xff)
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 8
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 16
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 24
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 32
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 40
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 48
	dm_build_1442++
	dm_build_1443 |= int64(dm_build_1441[dm_build_1442]&0xff) << 56
	return dm_build_1443
}

func (Dm_build_1445 *dm_build_1331) Dm_build_1444(dm_build_1446 []byte, dm_build_1447 int) float32 {
	return math.Float32frombits(Dm_build_1445.Dm_build_1461(dm_build_1446, dm_build_1447))
}

func (Dm_build_1449 *dm_build_1331) Dm_build_1448(dm_build_1450 []byte, dm_build_1451 int) float64 {
	return math.Float64frombits(Dm_build_1449.Dm_build_1466(dm_build_1450, dm_build_1451))
}

func (Dm_build_1453 *dm_build_1331) Dm_build_1452(dm_build_1454 []byte, dm_build_1455 int) uint8 {
	return uint8(dm_build_1454[dm_build_1455] & 0xff)
}

func (Dm_build_1457 *dm_build_1331) Dm_build_1456(dm_build_1458 []byte, dm_build_1459 int) uint16 {
	var dm_build_1460 uint16
	dm_build_1460 = uint16(dm_build_1458[dm_build_1459] & 0xff)
	dm_build_1459++
	dm_build_1460 |= uint16(dm_build_1458[dm_build_1459]&0xff) << 8
	return dm_build_1460
}

func (Dm_build_1462 *dm_build_1331) Dm_build_1461(dm_build_1463 []byte, dm_build_1464 int) uint32 {
	var dm_build_1465 uint32
	dm_build_1465 = uint32(dm_build_1463[dm_build_1464] & 0xff)
	dm_build_1464++
	dm_build_1465 |= uint32(dm_build_1463[dm_build_1464]&0xff) << 8
	dm_build_1464++
	dm_build_1465 |= uint32(dm_build_1463[dm_build_1464]&0xff) << 16
	dm_build_1464++
	dm_build_1465 |= uint32(dm_build_1463[dm_build_1464]&0xff) << 24
	return dm_build_1465
}

func (Dm_build_1467 *dm_build_1331) Dm_build_1466(dm_build_1468 []byte, dm_build_1469 int) uint64 {
	var dm_build_1470 uint64
	dm_build_1470 = uint64(dm_build_1468[dm_build_1469] & 0xff)
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 8
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 16
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 24
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 32
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 40
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 48
	dm_build_1469++
	dm_build_1470 |= uint64(dm_build_1468[dm_build_1469]&0xff) << 56
	return dm_build_1470
}

func (Dm_build_1472 *dm_build_1331) Dm_build_1471(dm_build_1473 []byte, dm_build_1474 int) []byte {
	dm_build_1475 := Dm_build_1472.Dm_build_1461(dm_build_1473, dm_build_1474)

	dm_build_1476 := make([]byte, dm_build_1475)
	copy(dm_build_1476[:int(dm_build_1475)], dm_build_1473[dm_build_1474+4:dm_build_1474+4+int(dm_build_1475)])
	return dm_build_1476
}

func (Dm_build_1478 *dm_build_1331) Dm_build_1477(dm_build_1479 []byte, dm_build_1480 int) []byte {
	dm_build_1481 := Dm_build_1478.Dm_build_1456(dm_build_1479, dm_build_1480)

	dm_build_1482 := make([]byte, dm_build_1481)
	copy(dm_build_1482[:int(dm_build_1481)], dm_build_1479[dm_build_1480+2:dm_build_1480+2+int(dm_build_1481)])
	return dm_build_1482
}

func (Dm_build_1484 *dm_build_1331) Dm_build_1483(dm_build_1485 []byte, dm_build_1486 int, dm_build_1487 int) []byte {

	dm_build_1488 := make([]byte, dm_build_1487)
	copy(dm_build_1488[:dm_build_1487], dm_build_1485[dm_build_1486:dm_build_1486+dm_build_1487])
	return dm_build_1488
}

func (Dm_build_1490 *dm_build_1331) Dm_build_1489(dm_build_1491 []byte, dm_build_1492 int, dm_build_1493 int, dm_build_1494 string, dm_build_1495 *DmConnection) string {
	return Dm_build_1490.Dm_build_1584(dm_build_1491[dm_build_1492:dm_build_1492+dm_build_1493], dm_build_1494, dm_build_1495)
}

func (Dm_build_1497 *dm_build_1331) Dm_build_1496(dm_build_1498 []byte, dm_build_1499 int, dm_build_1500 string, dm_build_1501 *DmConnection) string {
	dm_build_1502 := Dm_build_1497.Dm_build_1461(dm_build_1498, dm_build_1499)
	dm_build_1499 += 4
	return Dm_build_1497.Dm_build_1489(dm_build_1498, dm_build_1499, int(dm_build_1502), dm_build_1500, dm_build_1501)
}

func (Dm_build_1504 *dm_build_1331) Dm_build_1503(dm_build_1505 []byte, dm_build_1506 int, dm_build_1507 string, dm_build_1508 *DmConnection) string {
	dm_build_1509 := Dm_build_1504.Dm_build_1456(dm_build_1505, dm_build_1506)
	dm_build_1506 += 2
	return Dm_build_1504.Dm_build_1489(dm_build_1505, dm_build_1506, int(dm_build_1509), dm_build_1507, dm_build_1508)
}

func (Dm_build_1511 *dm_build_1331) Dm_build_1510(dm_build_1512 byte) []byte {
	return []byte{dm_build_1512}
}

func (Dm_build_1514 *dm_build_1331) Dm_build_1513(dm_build_1515 int8) []byte {
	return []byte{byte(dm_build_1515)}
}

func (Dm_build_1517 *dm_build_1331) Dm_build_1516(dm_build_1518 int16) []byte {
	return []byte{byte(dm_build_1518), byte(dm_build_1518 >> 8)}
}

func (Dm_build_1520 *dm_build_1331) Dm_build_1519(dm_build_1521 int32) []byte {
	return []byte{byte(dm_build_1521), byte(dm_build_1521 >> 8), byte(dm_build_1521 >> 16), byte(dm_build_1521 >> 24)}
}

func (Dm_build_1523 *dm_build_1331) Dm_build_1522(dm_build_1524 int64) []byte {
	return []byte{byte(dm_build_1524), byte(dm_build_1524 >> 8), byte(dm_build_1524 >> 16), byte(dm_build_1524 >> 24), byte(dm_build_1524 >> 32),
		byte(dm_build_1524 >> 40), byte(dm_build_1524 >> 48), byte(dm_build_1524 >> 56)}
}

func (Dm_build_1526 *dm_build_1331) Dm_build_1525(dm_build_1527 float32) []byte {
	return Dm_build_1526.Dm_build_1537(math.Float32bits(dm_build_1527))
}

func (Dm_build_1529 *dm_build_1331) Dm_build_1528(dm_build_1530 float64) []byte {
	return Dm_build_1529.Dm_build_1540(math.Float64bits(dm_build_1530))
}

func (Dm_build_1532 *dm_build_1331) Dm_build_1531(dm_build_1533 uint8) []byte {
	return []byte{byte(dm_build_1533)}
}

func (Dm_build_1535 *dm_build_1331) Dm_build_1534(dm_build_1536 uint16) []byte {
	return []byte{byte(dm_build_1536), byte(dm_build_1536 >> 8)}
}

func (Dm_build_1538 *dm_build_1331) Dm_build_1537(dm_build_1539 uint32) []byte {
	return []byte{byte(dm_build_1539), byte(dm_build_1539 >> 8), byte(dm_build_1539 >> 16), byte(dm_build_1539 >> 24)}
}

func (Dm_build_1541 *dm_build_1331) Dm_build_1540(dm_build_1542 uint64) []byte {
	return []byte{byte(dm_build_1542), byte(dm_build_1542 >> 8), byte(dm_build_1542 >> 16), byte(dm_build_1542 >> 24), byte(dm_build_1542 >> 32), byte(dm_build_1542 >> 40), byte(dm_build_1542 >> 48), byte(dm_build_1542 >> 56)}
}

func (Dm_build_1544 *dm_build_1331) Dm_build_1543(dm_build_1545 []byte, dm_build_1546 string, dm_build_1547 *DmConnection) []byte {
	if dm_build_1546 == "UTF-8" {
		return dm_build_1545
	}

	if dm_build_1547 == nil {
		if e := dm_build_1589(dm_build_1546); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1545), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1547.encodeBuffer == nil {
		dm_build_1547.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1547.encode = dm_build_1589(dm_build_1547.getServerEncoding())
		dm_build_1547.transformReaderDst = make([]byte, 4096)
		dm_build_1547.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1547.encode; e != nil {

		dm_build_1547.encodeBuffer.Reset()

		n, err := dm_build_1547.encodeBuffer.ReadFrom(
			Dm_build_1603(bytes.NewReader(dm_build_1545), e.NewEncoder(), dm_build_1547.transformReaderDst, dm_build_1547.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1547.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1549 *dm_build_1331) Dm_build_1548(dm_build_1550 string, dm_build_1551 string, dm_build_1552 *DmConnection) []byte {
	return Dm_build_1549.Dm_build_1543([]byte(dm_build_1550), dm_build_1551, dm_build_1552)
}

func (Dm_build_1554 *dm_build_1331) Dm_build_1553(dm_build_1555 []byte) byte {
	return Dm_build_1554.Dm_build_1425(dm_build_1555, 0)
}

func (Dm_build_1557 *dm_build_1331) Dm_build_1556(dm_build_1558 []byte) int16 {
	return Dm_build_1557.Dm_build_1429(dm_build_1558, 0)
}

func (Dm_build_1560 *dm_build_1331) Dm_build_1559(dm_build_1561 []byte) int32 {
	return Dm_build_1560.Dm_build_1434(dm_build_1561, 0)
}

func (Dm_build_1563 *dm_build_1331) Dm_build_1562(dm_build_1564 []byte) int64 {
	return Dm_build_1563.Dm_build_1439(dm_build_1564, 0)
}

func (Dm_build_1566 *dm_build_1331) Dm_build_1565(dm_build_1567 []byte) float32 {
	return Dm_build_1566.Dm_build_1444(dm_build_1567, 0)
}

func (Dm_build_1569 *dm_build_1331) Dm_build_1568(dm_build_1570 []byte) float64 {
	return Dm_build_1569.Dm_build_1448(dm_build_1570, 0)
}

func (Dm_build_1572 *dm_build_1331) Dm_build_1571(dm_build_1573 []byte) uint8 {
	return Dm_build_1572.Dm_build_1452(dm_build_1573, 0)
}

func (Dm_build_1575 *dm_build_1331) Dm_build_1574(dm_build_1576 []byte) uint16 {
	return Dm_build_1575.Dm_build_1456(dm_build_1576, 0)
}

func (Dm_build_1578 *dm_build_1331) Dm_build_1577(dm_build_1579 []byte) uint32 {
	return Dm_build_1578.Dm_build_1461(dm_build_1579, 0)
}

func (Dm_build_1581 *dm_build_1331) Dm_build_1580(dm_build_1582 []byte, dm_build_1583 string) []byte {
	if dm_build_1583 == "UTF-8" {
		return dm_build_1582
	}

	if e := dm_build_1589(dm_build_1583); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1582), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1585 *dm_build_1331) Dm_build_1584(dm_build_1586 []byte, dm_build_1587 string, dm_build_1588 *DmConnection) string {
	return string(Dm_build_1585.Dm_build_1580(dm_build_1586, dm_build_1587))
}

func dm_build_1589(dm_build_1590 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1590); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1591 struct {
	dm_build_1592 io.Reader
	dm_build_1593 transform.Transformer
	dm_build_1594 error

	dm_build_1595                []byte
	dm_build_1596, dm_build_1597 int

	dm_build_1598                []byte
	dm_build_1599, dm_build_1600 int

	dm_build_1601 bool
}

const dm_build_1602 = 4096

func Dm_build_1603(dm_build_1604 io.Reader, dm_build_1605 transform.Transformer, dm_build_1606 []byte, dm_build_1607 []byte) *Dm_build_1591 {
	dm_build_1605.Reset()
	return &Dm_build_1591{
		dm_build_1592: dm_build_1604,
		dm_build_1593: dm_build_1605,
		dm_build_1595: dm_build_1606,
		dm_build_1598: dm_build_1607,
	}
}

func (dm_build_1609 *Dm_build_1591) Read(dm_build_1610 []byte) (int, error) {
	dm_build_1611, dm_build_1612 := 0, error(nil)
	for {

		if dm_build_1609.dm_build_1596 != dm_build_1609.dm_build_1597 {
			dm_build_1611 = copy(dm_build_1610, dm_build_1609.dm_build_1595[dm_build_1609.dm_build_1596:dm_build_1609.dm_build_1597])
			dm_build_1609.dm_build_1596 += dm_build_1611
			if dm_build_1609.dm_build_1596 == dm_build_1609.dm_build_1597 && dm_build_1609.dm_build_1601 {
				return dm_build_1611, dm_build_1609.dm_build_1594
			}
			return dm_build_1611, nil
		} else if dm_build_1609.dm_build_1601 {
			return 0, dm_build_1609.dm_build_1594
		}

		if dm_build_1609.dm_build_1599 != dm_build_1609.dm_build_1600 || dm_build_1609.dm_build_1594 != nil {
			dm_build_1609.dm_build_1596 = 0
			dm_build_1609.dm_build_1597, dm_build_1611, dm_build_1612 = dm_build_1609.dm_build_1593.Transform(dm_build_1609.dm_build_1595, dm_build_1609.dm_build_1598[dm_build_1609.dm_build_1599:dm_build_1609.dm_build_1600], dm_build_1609.dm_build_1594 == io.EOF)
			dm_build_1609.dm_build_1599 += dm_build_1611

			switch {
			case dm_build_1612 == nil:
				if dm_build_1609.dm_build_1599 != dm_build_1609.dm_build_1600 {
					dm_build_1609.dm_build_1594 = nil
				}

				dm_build_1609.dm_build_1601 = dm_build_1609.dm_build_1594 != nil
				continue
			case dm_build_1612 == transform.ErrShortDst && (dm_build_1609.dm_build_1597 != 0 || dm_build_1611 != 0):

				continue
			case dm_build_1612 == transform.ErrShortSrc && dm_build_1609.dm_build_1600-dm_build_1609.dm_build_1599 != len(dm_build_1609.dm_build_1598) && dm_build_1609.dm_build_1594 == nil:

			default:
				dm_build_1609.dm_build_1601 = true

				if dm_build_1609.dm_build_1594 == nil || dm_build_1609.dm_build_1594 == io.EOF {
					dm_build_1609.dm_build_1594 = dm_build_1612
				}
				continue
			}
		}

		if dm_build_1609.dm_build_1599 != 0 {
			dm_build_1609.dm_build_1599, dm_build_1609.dm_build_1600 = 0, copy(dm_build_1609.dm_build_1598, dm_build_1609.dm_build_1598[dm_build_1609.dm_build_1599:dm_build_1609.dm_build_1600])
		}
		dm_build_1611, dm_build_1609.dm_build_1594 = dm_build_1609.dm_build_1592.Read(dm_build_1609.dm_build_1598[dm_build_1609.dm_build_1600:])
		dm_build_1609.dm_build_1600 += dm_build_1611
	}
}
