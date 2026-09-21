/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/gaoyuan98/dm/security"
)

const (
	Dm_build_334 = 8192
	Dm_build_335 = 2 * time.Second
)

type dm_build_336 struct {
	dm_build_337 Dm_build_602
	dm_build_338 *Dm_build_0
	dm_build_339 *DmConnection
	dm_build_340 security.Cipher
	dm_build_341 bool
	dm_build_342 bool
	dm_build_343 *security.DhKey

	dm_build_344 bool
	dm_build_345 string
	dm_build_346 bool
}

func dm_build_347(dm_build_348 context.Context, dm_build_349 *DmConnection) (*dm_build_336, error) {

	dm_build_350 := &Dm_build_604{}
	if err := dm_build_350.Dm_build_603(dm_build_348, dm_build_349); err != nil {
		return nil, err
	}

	dm_build_351 := &dm_build_336{}
	dm_build_351.dm_build_337 = dm_build_350
	dm_build_351.dm_build_338 = Dm_build_3(Dm_build_686)
	dm_build_351.dm_build_339 = dm_build_349
	dm_build_351.dm_build_341 = false
	dm_build_351.dm_build_342 = false
	dm_build_351.dm_build_344 = false
	dm_build_351.dm_build_345 = ""
	dm_build_351.dm_build_346 = false
	dm_build_349.Access = dm_build_351

	return dm_build_351, nil
}

func (dm_build_353 *dm_build_336) dm_build_352(dm_build_354 dm_build_808) bool {
	var dm_build_355 = dm_build_353.dm_build_339.dmConnector.compress
	if dm_build_354.dm_build_823() == Dm_build_714 || dm_build_355 == Dm_build_763 {
		return false
	}

	if dm_build_355 == Dm_build_761 {
		return true
	} else if dm_build_355 == Dm_build_762 {
		return !dm_build_353.dm_build_339.Local && dm_build_354.dm_build_821() > Dm_build_760
	}

	return false
}

func (dm_build_357 *dm_build_336) dm_build_356(dm_build_358 dm_build_808) bool {
	var dm_build_359 = dm_build_357.dm_build_339.dmConnector.compress
	if dm_build_358.dm_build_823() == Dm_build_714 || dm_build_359 == Dm_build_763 {
		return false
	}

	if dm_build_359 == Dm_build_761 {
		return true
	} else if dm_build_359 == Dm_build_762 {
		return dm_build_357.dm_build_338.Dm_build_267(Dm_build_722) == 1
	}

	return false
}

func (dm_build_361 *dm_build_336) dm_build_360(dm_build_362 dm_build_808) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_364 := dm_build_362.dm_build_821()

	if dm_build_364 > 0 {

		if dm_build_361.dm_build_352(dm_build_362) {
			var retBytes, err = Compress(dm_build_361.dm_build_338, Dm_build_715, int(dm_build_364), int(dm_build_361.dm_build_339.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_361.dm_build_338.Dm_build_14(Dm_build_715)

			dm_build_361.dm_build_338.Dm_build_55(dm_build_364)

			dm_build_361.dm_build_338.Dm_build_83(retBytes)

			dm_build_362.dm_build_822(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_361.dm_build_338.Dm_build_187(Dm_build_722, 1)
		}

		if dm_build_361.dm_build_342 {
			dm_build_364 = dm_build_362.dm_build_821()
			var retBytes = dm_build_361.dm_build_340.Encrypt(dm_build_361.dm_build_338.Dm_build_294(Dm_build_715, int(dm_build_364)), true)

			dm_build_361.dm_build_338.Dm_build_14(Dm_build_715)

			dm_build_361.dm_build_338.Dm_build_83(retBytes)

			dm_build_362.dm_build_822(int32(len(retBytes)))
		}
	}

	if dm_build_361.dm_build_338.Dm_build_12() > Dm_build_687 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_362.dm_build_817()

	dm_build_361.dm_build_338.Dm_build_17(0)
	if _, err := dm_build_361.dm_build_338.Dm_build_36(dm_build_361.dm_build_337); err != nil {
		return err
	}

	return nil
}

func (dm_build_366 *dm_build_336) dm_build_365(dm_build_367 dm_build_808) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_369 := int32(0)

	dm_build_366.dm_build_338.Dm_build_14(0)

	if _, err := dm_build_366.dm_build_338.Dm_build_30(dm_build_366.dm_build_337, Dm_build_715); err != nil {
		return err
	}
	dm_build_369 = dm_build_367.dm_build_821()

	if dm_build_369 > 0 {
		if _, err := dm_build_366.dm_build_338.Dm_build_30(dm_build_366.dm_build_337, int(dm_build_369)); err != nil {
			return err
		}
	}

	dm_build_367.dm_build_818()
	dm_build_369 = dm_build_367.dm_build_821()
	if dm_build_369 <= 0 {
		return nil
	}

	if dm_build_366.dm_build_342 {
		ebytes := dm_build_366.dm_build_338.Dm_build_294(Dm_build_715, int(dm_build_369))
		bytes, err := dm_build_366.dm_build_340.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_366.dm_build_338.Dm_build_14(Dm_build_715)
		dm_build_366.dm_build_338.Dm_build_83(bytes)
		dm_build_367.dm_build_822(int32(len(bytes)))
	}

	if dm_build_366.dm_build_356(dm_build_367) {

		dm_build_369 = dm_build_367.dm_build_821()
		cbytes := dm_build_366.dm_build_338.Dm_build_294(Dm_build_715+ULINT_SIZE, int(dm_build_369-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_366.dm_build_339.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_366.dm_build_338.Dm_build_14(Dm_build_715)
		dm_build_366.dm_build_338.Dm_build_83(bytes)
		dm_build_367.dm_build_822(int32(len(bytes)))
	}
	return nil
}

func (dm_build_371 *dm_build_336) dm_build_370(dm_build_372 dm_build_808) (dm_build_373 interface{}, dm_build_374 error) {
	if dm_build_371.dm_build_346 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_375 := dm_build_371.dm_build_339
	dm_build_375.mu.Lock()
	defer dm_build_375.mu.Unlock()
	dm_build_374 = dm_build_372.dm_build_812(dm_build_372)
	if dm_build_374 != nil {
		return nil, dm_build_374
	}

	dm_build_374 = dm_build_371.dm_build_360(dm_build_372)
	if dm_build_374 != nil {
		return nil, dm_build_374
	}

	dm_build_374 = dm_build_371.dm_build_365(dm_build_372)
	if dm_build_374 != nil {
		return nil, dm_build_374
	}

	return dm_build_372.dm_build_816(dm_build_372)
}

func (dm_build_377 *dm_build_336) dm_build_376() (*dm_build_1270, error) {

	Dm_build_378 := dm_build_1276(dm_build_377)
	_, dm_build_379 := dm_build_377.dm_build_370(Dm_build_378)
	if dm_build_379 != nil {
		return nil, dm_build_379
	}

	return Dm_build_378, nil
}

func (dm_build_381 *dm_build_336) dm_build_380() error {

	dm_build_382 := dm_build_1134(dm_build_381)
	_, dm_build_383 := dm_build_381.dm_build_370(dm_build_382)
	if dm_build_383 != nil {
		return dm_build_383
	}

	return nil
}

func (dm_build_385 *dm_build_336) dm_build_384() error {

	var dm_build_386 *dm_build_1270
	var err error
	if dm_build_386, err = dm_build_385.dm_build_376(); err != nil {
		return err
	}

	if err = dm_build_385.dm_build_599(); err != nil {
		return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
	}

	if dm_build_385.dm_build_342 || dm_build_385.dm_build_341 {
		k, err := dm_build_385.dm_build_589()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_386.Dm_build_1274)
		encryptType := dm_build_386.dm_build_1272
		hashType := int(dm_build_386.Dm_build_1273)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_385.dm_build_592(encryptType, sessionKey, dm_build_385.dm_build_339.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_385.dm_build_380(); err != nil {
		return err
	}
	return nil
}

func (dm_build_389 *dm_build_336) Dm_build_388(dm_build_390 *DmStatement) error {
	dm_build_391 := dm_build_1303(dm_build_389, dm_build_390)
	_, dm_build_392 := dm_build_389.dm_build_370(dm_build_391)
	if dm_build_392 != nil {
		return dm_build_392
	}

	return nil
}

func (dm_build_394 *dm_build_336) Dm_build_393(dm_build_395 int32) error {
	dm_build_396 := dm_build_1313(dm_build_394, dm_build_395)
	_, dm_build_397 := dm_build_394.dm_build_370(dm_build_396)
	if dm_build_397 != nil {
		return dm_build_397
	}

	return nil
}

func (dm_build_399 *dm_build_336) Dm_build_398(dm_build_400 *DmStatement, dm_build_401 bool, dm_build_402 int16) (*execRetInfo, error) {
	dm_build_403 := dm_build_1176(dm_build_399, dm_build_400, dm_build_401, dm_build_402)
	dm_build_404, dm_build_405 := dm_build_399.dm_build_370(dm_build_403)
	if dm_build_405 != nil {
		return nil, dm_build_405
	}
	return dm_build_404.(*execRetInfo), nil
}

func (dm_build_407 *dm_build_336) Dm_build_406(dm_build_408 *DmStatement, dm_build_409 int16) (*execRetInfo, error) {
	return dm_build_407.Dm_build_398(dm_build_408, false, Dm_build_767)
}

func (dm_build_411 *dm_build_336) Dm_build_410(dm_build_412 *DmStatement, dm_build_413 []OptParameter) (*execRetInfo, error) {
	dm_build_414, dm_build_415 := dm_build_411.dm_build_370(dm_build_911(dm_build_411, dm_build_412, dm_build_413))
	if dm_build_415 != nil {
		return nil, dm_build_415
	}

	return dm_build_414.(*execRetInfo), nil
}

func (dm_build_417 *dm_build_336) Dm_build_416(dm_build_418 *DmStatement, dm_build_419 int16) (*execRetInfo, error) {
	return dm_build_417.Dm_build_398(dm_build_418, true, dm_build_419)
}

func (dm_build_421 *dm_build_336) Dm_build_420(dm_build_422 *DmStatement, dm_build_423 [][]interface{}) (*execRetInfo, error) {
	dm_build_424 := dm_build_943(dm_build_421, dm_build_422, dm_build_423)
	dm_build_425, dm_build_426 := dm_build_421.dm_build_370(dm_build_424)
	if dm_build_426 != nil {
		return nil, dm_build_426
	}
	return dm_build_425.(*execRetInfo), nil
}

func (dm_build_428 *dm_build_336) Dm_build_427(dm_build_429 *DmStatement, dm_build_430 [][]interface{}, dm_build_431 bool) (*execRetInfo, error) {
	var dm_build_432, dm_build_433 = 0, 0
	var dm_build_434 = len(dm_build_430)
	var dm_build_435 [][]interface{}
	var dm_build_436 = NewExceInfo()
	dm_build_436.updateCounts = make([]int64, dm_build_434)
	var dm_build_437 = false
	for dm_build_432 < dm_build_434 {
		for dm_build_433 = dm_build_432; dm_build_433 < dm_build_434; dm_build_433++ {
			paramData := dm_build_430[dm_build_433]
			bindData := make([]interface{}, dm_build_429.paramCount)
			dm_build_437 = false
			for icol := 0; icol < int(dm_build_429.paramCount); icol++ {
				if dm_build_429.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_428.dm_build_572(bindData, paramData, icol) {
					dm_build_437 = true
					break
				}
			}

			if dm_build_437 {
				break
			}
			dm_build_435 = append(dm_build_435, bindData)
		}

		if dm_build_433 != dm_build_432 {
			tmpExecInfo, err := dm_build_428.Dm_build_420(dm_build_429, dm_build_435)
			if err != nil {
				return nil, err
			}
			dm_build_435 = dm_build_435[0:0]
			dm_build_436.union(tmpExecInfo, dm_build_432, dm_build_433-dm_build_432)
		}

		if dm_build_433 < dm_build_434 {
			tmpExecInfo, err := dm_build_428.Dm_build_446(dm_build_429, dm_build_430[dm_build_433], dm_build_431)
			if err != nil {
				return nil, err
			}

			dm_build_431 = true
			dm_build_436.union(tmpExecInfo, dm_build_433, 1)
		}

		dm_build_432 = dm_build_433 + 1
	}
	for _, i := range dm_build_436.updateCounts {
		if i > 0 {
			dm_build_436.updateCount += i
		}
	}
	return dm_build_436, nil
}

func (dm_build_439 *dm_build_336) dm_build_438(dm_build_440 *DmStatement, dm_build_441 []parameter) error {
	if !dm_build_440.prepared {
		retInfo, err := dm_build_439.Dm_build_398(dm_build_440, false, Dm_build_767)
		if err != nil {
			return nil
		}
		dm_build_440.serverParams = retInfo.serverParams
		dm_build_440.paramCount = int32(len(dm_build_440.serverParams))
		dm_build_440.prepared = true
	}

	dm_build_442 := dm_build_1165(dm_build_439, dm_build_440, dm_build_440.bindParams)
	dm_build_443, err := dm_build_439.dm_build_370(dm_build_442)
	if err != nil {
		return nil
	}
	retInfo := dm_build_443.(*execRetInfo)
	if retInfo.serverParams != nil && len(retInfo.serverParams) > 0 {
		dm_build_440.serverParams = retInfo.serverParams
		dm_build_440.paramCount = int32(len(dm_build_440.serverParams))
	}
	dm_build_440.preExec = true
	return nil
}

func (dm_build_447 *dm_build_336) Dm_build_446(dm_build_448 *DmStatement, dm_build_449 []interface{}, dm_build_450 bool) (*execRetInfo, error) {

	var dm_build_451 = make([]interface{}, dm_build_448.paramCount)
	for icol := 0; icol < int(dm_build_448.paramCount); icol++ {
		if dm_build_448.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_447.dm_build_572(dm_build_451, dm_build_449, icol) {

			if !dm_build_450 {
				dm_build_447.dm_build_438(dm_build_448, dm_build_448.bindParams)

				dm_build_450 = true
			}

			dm_build_447.dm_build_578(dm_build_448, dm_build_448.bindParams[icol], icol, dm_build_449[icol].(iOffRowBinder))
			dm_build_451[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_452 = make([][]interface{}, 1, 1)
	dm_build_452[0] = dm_build_451

	dm_build_453 := dm_build_943(dm_build_447, dm_build_448, dm_build_452)
	dm_build_454, dm_build_455 := dm_build_447.dm_build_370(dm_build_453)
	if dm_build_455 != nil {
		return nil, dm_build_455
	}
	return dm_build_454.(*execRetInfo), nil
}

func (dm_build_457 *dm_build_336) Dm_build_456(dm_build_458 *DmStatement, dm_build_459 int16) (*execRetInfo, error) {
	dm_build_460 := dm_build_1152(dm_build_457, dm_build_458, dm_build_459)

	dm_build_461, dm_build_462 := dm_build_457.dm_build_370(dm_build_460)
	if dm_build_462 != nil {
		return nil, dm_build_462
	}
	return dm_build_461.(*execRetInfo), nil
}

func (dm_build_464 *dm_build_336) Dm_build_463(dm_build_465 *innerRows, dm_build_466 int64) (*execRetInfo, error) {
	dm_build_467 := dm_build_1051(dm_build_464, dm_build_465, dm_build_466, INT64_MAX)
	dm_build_468, dm_build_469 := dm_build_464.dm_build_370(dm_build_467)
	if dm_build_469 != nil {
		return nil, dm_build_469
	}
	return dm_build_468.(*execRetInfo), nil
}

func (dm_build_471 *dm_build_336) Commit() error {
	dm_build_472 := dm_build_896(dm_build_471)
	_, dm_build_473 := dm_build_471.dm_build_370(dm_build_472)
	if dm_build_473 != nil {
		return dm_build_473
	}

	return nil
}

func (dm_build_475 *dm_build_336) Rollback() error {
	dm_build_476 := dm_build_1214(dm_build_475)
	_, dm_build_477 := dm_build_475.dm_build_370(dm_build_476)
	if dm_build_477 != nil {
		return dm_build_477
	}

	return nil
}

func (dm_build_479 *dm_build_336) Dm_build_478(dm_build_480 *DmConnection) error {
	dm_build_481 := dm_build_1219(dm_build_479, dm_build_480.IsoLevel)
	_, dm_build_482 := dm_build_479.dm_build_370(dm_build_481)
	if dm_build_482 != nil {
		return dm_build_482
	}

	return nil
}

func (dm_build_484 *dm_build_336) Dm_build_483(dm_build_485 *DmStatement, dm_build_486 string) error {
	dm_build_487 := dm_build_901(dm_build_484, dm_build_485, dm_build_486)
	_, dm_build_488 := dm_build_484.dm_build_370(dm_build_487)
	if dm_build_488 != nil {
		return dm_build_488
	}

	return nil
}

func (dm_build_490 *dm_build_336) Dm_build_489(dm_build_491 []uint32) ([]int64, error) {
	dm_build_492 := dm_build_1321(dm_build_490, dm_build_491)
	dm_build_493, dm_build_494 := dm_build_490.dm_build_370(dm_build_492)
	if dm_build_494 != nil {
		return nil, dm_build_494
	}
	return dm_build_493.([]int64), nil
}

func (dm_build_496 *dm_build_336) Close() error {
	if dm_build_496.dm_build_346 {
		return nil
	}

	dm_build_497 := dm_build_496.dm_build_337.Close()
	if dm_build_497 != nil {
		return dm_build_497
	}

	dm_build_496.dm_build_346 = true
	return nil
}

func (dm_build_499 *dm_build_336) dm_build_498(dm_build_500 *lob) (int64, error) {
	dm_build_501 := dm_build_1084(dm_build_499, dm_build_500)
	dm_build_502, dm_build_503 := dm_build_499.dm_build_370(dm_build_501)
	if dm_build_503 != nil {
		return 0, dm_build_503
	}
	return dm_build_502.(int64), nil
}

func (dm_build_505 *dm_build_336) dm_build_504(dm_build_506 *lob, dm_build_507 int32, dm_build_508 int32) (*lobRetInfo, error) {
	dm_build_509 := dm_build_1069(dm_build_505, dm_build_506, int(dm_build_507), int(dm_build_508))
	dm_build_510, dm_build_511 := dm_build_505.dm_build_370(dm_build_509)
	if dm_build_511 != nil {
		return nil, dm_build_511
	}
	return dm_build_510.(*lobRetInfo), nil
}

func (dm_build_513 *dm_build_336) dm_build_512(dm_build_514 *DmBlob, dm_build_515 int32, dm_build_516 int32) ([]byte, error) {
	var dm_build_517 = make([]byte, dm_build_516)
	var dm_build_518 int32 = 0
	var dm_build_519 int32 = 0
	var dm_build_520 *lobRetInfo
	var dm_build_521 []byte
	var dm_build_522 error
	for dm_build_518 < dm_build_516 {
		dm_build_519 = dm_build_516 - dm_build_518
		if dm_build_519 > Dm_build_801 {
			dm_build_519 = Dm_build_801
		}
		dm_build_520, dm_build_522 = dm_build_513.dm_build_504(&dm_build_514.lob, dm_build_515+dm_build_518, dm_build_519)
		if dm_build_522 != nil {
			return nil, dm_build_522
		}
		dm_build_521 = dm_build_520.data
		if dm_build_521 == nil || len(dm_build_521) == 0 {
			break
		}
		Dm_build_1332.Dm_build_1388(dm_build_517, int(dm_build_518), dm_build_521, 0, len(dm_build_521))
		dm_build_518 += int32(len(dm_build_521))
		if dm_build_514.readOver {
			break
		}
	}
	return dm_build_517, nil
}

func (dm_build_524 *dm_build_336) dm_build_523(dm_build_525 *DmClob, dm_build_526 int32, dm_build_527 int32) (string, error) {
	var dm_build_528 bytes.Buffer
	var dm_build_529 int32 = 0
	var dm_build_530 int32 = 0
	var dm_build_531 *lobRetInfo
	var dm_build_532 []byte
	var dm_build_533 string
	var dm_build_534 error
	for dm_build_529 < dm_build_527 {
		dm_build_530 = dm_build_527 - dm_build_529
		if dm_build_530 > Dm_build_801/2 {
			dm_build_530 = Dm_build_801 / 2
		}
		dm_build_531, dm_build_534 = dm_build_524.dm_build_504(&dm_build_525.lob, dm_build_526+dm_build_529, dm_build_530)
		if dm_build_534 != nil {
			return "", dm_build_534
		}
		dm_build_532 = dm_build_531.data
		if dm_build_532 == nil || len(dm_build_532) == 0 {
			break
		}
		dm_build_533 = Dm_build_1332.Dm_build_1489(dm_build_532, 0, len(dm_build_532), dm_build_525.serverEncoding, dm_build_524.dm_build_339)

		dm_build_528.WriteString(dm_build_533)
		var strLen = dm_build_531.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_533))
		}
		dm_build_529 += int32(strLen)
		if dm_build_525.readOver {
			break
		}
	}
	return dm_build_528.String(), nil
}

func (dm_build_536 *dm_build_336) dm_build_535(dm_build_537 *DmClob, dm_build_538 int, dm_build_539 string, dm_build_540 string) (int, error) {
	var dm_build_541 = Dm_build_1332.Dm_build_1548(dm_build_539, dm_build_540, dm_build_536.dm_build_339)
	var dm_build_542 = 0
	var dm_build_543 = len(dm_build_541)
	var dm_build_544 = 0
	var dm_build_545 = 0
	var dm_build_546 = 0
	var dm_build_547 = dm_build_543/Dm_build_800 + 1
	var dm_build_548 byte = 0
	var dm_build_549 byte = 0x01
	var dm_build_550 byte = 0x02
	for i := 0; i < dm_build_547; i++ {
		dm_build_548 = 0
		if i == 0 {
			dm_build_548 |= dm_build_549
		}
		if i == dm_build_547-1 {
			dm_build_548 |= dm_build_550
		}
		dm_build_546 = dm_build_543 - dm_build_545
		if dm_build_546 > Dm_build_800 {
			dm_build_546 = Dm_build_800
		}

		setLobData := dm_build_1233(dm_build_536, &dm_build_537.lob, dm_build_548, dm_build_538, dm_build_541, dm_build_542, dm_build_546)
		ret, err := dm_build_536.dm_build_370(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_544, nil
		} else {
			dm_build_538 += int(tmp)
			dm_build_544 += int(tmp)
			dm_build_545 += dm_build_546
			dm_build_542 += dm_build_546
		}
	}
	return dm_build_544, nil
}

func (dm_build_552 *dm_build_336) dm_build_551(dm_build_553 *DmBlob, dm_build_554 int, dm_build_555 []byte) (int, error) {
	var dm_build_556 = 0
	var dm_build_557 = len(dm_build_555)
	var dm_build_558 = 0
	var dm_build_559 = 0
	var dm_build_560 = 0
	var dm_build_561 = dm_build_557/Dm_build_800 + 1
	var dm_build_562 byte = 0
	var dm_build_563 byte = 0x01
	var dm_build_564 byte = 0x02
	for i := 0; i < dm_build_561; i++ {
		dm_build_562 = 0
		if i == 0 {
			dm_build_562 |= dm_build_563
		}
		if i == dm_build_561-1 {
			dm_build_562 |= dm_build_564
		}
		dm_build_560 = dm_build_557 - dm_build_559
		if dm_build_560 > Dm_build_800 {
			dm_build_560 = Dm_build_800
		}

		setLobData := dm_build_1233(dm_build_552, &dm_build_553.lob, dm_build_562, dm_build_554, dm_build_555, dm_build_556, dm_build_560)
		ret, err := dm_build_552.dm_build_370(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_558, nil
		} else {
			dm_build_554 += int(tmp)
			dm_build_558 += int(tmp)
			dm_build_559 += dm_build_560
			dm_build_556 += dm_build_560
		}
	}
	return dm_build_558, nil
}

func (dm_build_566 *dm_build_336) dm_build_565(dm_build_567 *lob, dm_build_568 int) (int64, error) {
	dm_build_569 := dm_build_1095(dm_build_566, dm_build_567, dm_build_568)
	dm_build_570, dm_build_571 := dm_build_566.dm_build_370(dm_build_569)
	if dm_build_571 != nil {
		return dm_build_567.length, dm_build_571
	}
	return dm_build_570.(int64), nil
}

func (dm_build_573 *dm_build_336) dm_build_572(dm_build_574 []interface{}, dm_build_575 []interface{}, dm_build_576 int) bool {
	var dm_build_577 = false
	dm_build_574[dm_build_576] = dm_build_575[dm_build_576]

	if binder, ok := dm_build_575[dm_build_576].(iOffRowBinder); ok {
		dm_build_577 = true
		dm_build_574[dm_build_576] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_573.dm_build_339) {
			dm_build_574[dm_build_576] = &lobCtl{lob.buildCtlData()}
			dm_build_577 = false
		}
	} else {
		dm_build_574[dm_build_576] = dm_build_575[dm_build_576]
	}
	return dm_build_577
}

func (dm_build_579 *dm_build_336) dm_build_578(dm_build_580 *DmStatement, dm_build_581 parameter, dm_build_582 int, dm_build_583 iOffRowBinder) error {
	var dm_build_584 = Dm_build_1617()
	dm_build_583.read(dm_build_584)
	var dm_build_585 = 0
	for !dm_build_583.isReadOver() || dm_build_584.Dm_build_1618() > 0 {
		if !dm_build_583.isReadOver() && dm_build_584.Dm_build_1618() < Dm_build_800 {
			dm_build_583.read(dm_build_584)
		}
		if dm_build_584.Dm_build_1618() > Dm_build_800 {
			dm_build_585 = Dm_build_800
		} else {
			dm_build_585 = dm_build_584.Dm_build_1618()
		}

		putData := dm_build_1204(dm_build_579, dm_build_580, int16(dm_build_582), dm_build_584, int32(dm_build_585))
		_, err := dm_build_579.dm_build_370(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_587 *dm_build_336) dm_build_586() ([]byte, error) {
	var dm_build_588 error
	if dm_build_587.dm_build_343 == nil {
		if dm_build_587.dm_build_343, dm_build_588 = security.NewClientKeyPair(); dm_build_588 != nil {
			return nil, dm_build_588
		}
	}
	return security.Bn2Bytes(dm_build_587.dm_build_343.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_590 *dm_build_336) dm_build_589() (*security.DhKey, error) {
	var dm_build_591 error
	if dm_build_590.dm_build_343 == nil {
		if dm_build_590.dm_build_343, dm_build_591 = security.NewClientKeyPair(); dm_build_591 != nil {
			return nil, dm_build_591
		}
	}
	return dm_build_590.dm_build_343, nil
}

func (dm_build_593 *dm_build_336) dm_build_592(dm_build_594 int, dm_build_595 []byte, dm_build_596 string, dm_build_597 int) (dm_build_598 error) {
	if dm_build_594 > 0 && dm_build_594 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_595 != nil {
		dm_build_593.dm_build_340, dm_build_598 = security.NewSymmCipher(dm_build_594, dm_build_595)
	} else if dm_build_594 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_593.dm_build_340, dm_build_598 = security.NewThirdPartCipher(dm_build_594, dm_build_595, dm_build_596, dm_build_597); dm_build_598 != nil {
			dm_build_598 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_598.Error()).throw()
		}
	}
	return
}

func (dm_build_600 *dm_build_336) dm_build_599() (dm_build_601 error) {
	if dm_build_600.dm_build_339.sslEncrypt == 2 {

		socket := &Dm_build_631{
			dm_build_632: dm_build_600.dm_build_337,
			dm_build_634: true,
		}
		if dm_build_601 = socket.Dm_build_603(context.Background(), dm_build_600.dm_build_339); dm_build_601 != nil {
			return dm_build_601
		}
		dm_build_600.dm_build_337 = socket
	} else if dm_build_600.dm_build_339.sslEncrypt == 1 || dm_build_600.dm_build_339.sslEncrypt == 4 {
		socket := &Dm_build_631{
			dm_build_632: dm_build_600.dm_build_337,
			dm_build_634: false,
		}
		if dm_build_601 = socket.Dm_build_603(context.Background(), dm_build_600.dm_build_339); dm_build_601 != nil {
			return dm_build_601
		}
		dm_build_600.dm_build_337 = socket
	}

	return
}

type Dm_build_602 interface {
	io.ReadWriteCloser

	Dm_build_603(ctx context.Context, dmConn *DmConnection) error
}

type Dm_build_604 struct {
	dm_build_605 net.Conn
}

func (dm_build_607 *Dm_build_604) Dm_build_603(dm_build_608 context.Context, dm_build_609 *DmConnection) (dm_build_610 error) {
	dialsLock.RLock()
	dm_build_611, dm_build_612 := dials[dm_build_609.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_612 {
		dm_build_607.dm_build_605, dm_build_610 = dm_build_611(dm_build_608, dm_build_609.dmConnector.host+":"+strconv.Itoa(int(dm_build_609.dmConnector.port)))
	} else {
		dm_build_607.dm_build_605, dm_build_610 = dm_build_607.dm_build_625(dm_build_609.dmConnector.host+":"+strconv.Itoa(int(dm_build_609.dmConnector.port)), time.Duration(dm_build_609.dmConnector.socketTimeout)*time.Second)
	}
	return dm_build_610
}

func (dm_build_614 *Dm_build_604) Close() error {
	return dm_build_614.dm_build_605.Close()
}

func (dm_build_616 *Dm_build_604) Write(dm_build_617 []byte) (dm_build_618 int, dm_build_619 error) {
	return dm_build_616.dm_build_605.Write(dm_build_617)
}

func (dm_build_621 *Dm_build_604) Read(dm_build_622 []byte) (dm_build_623 int, dm_build_624 error) {
	return dm_build_621.dm_build_605.Read(dm_build_622)
}

func (dm_build_626 *Dm_build_604) dm_build_625(dm_build_627 string, dm_build_628 time.Duration) (net.Conn, error) {
	dm_build_629, dm_build_630 := net.DialTimeout("tcp", dm_build_627, dm_build_628)
	if dm_build_630 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_627).throw()
	}

	if tcpConn, ok := dm_build_629.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_335)
		tcpConn.SetNoDelay(true)

	}
	return dm_build_629, nil
}

type Dm_build_631 struct {
	dm_build_632 Dm_build_602

	dm_build_633 tls.Conn

	dm_build_634 bool
}

func (dm_build_636 *Dm_build_631) Dm_build_603(dm_build_637 context.Context, dm_build_638 *DmConnection) (dm_build_639 error) {
	var dm_build_640 *tls.Conn
	dm_build_640, dm_build_639 = dm_build_636.dm_build_653((dm_build_636.dm_build_632).(*Dm_build_604).dm_build_605, dm_build_638)
	if dm_build_639 != nil {
		return
	}
	dm_build_636.dm_build_633 = *dm_build_640
	return nil
}

func (dm_build_642 *Dm_build_631) Close() error {
	if err := dm_build_642.dm_build_633.Close(); err != nil {
		return err
	}
	return dm_build_642.dm_build_632.Close()
}

func (dm_build_644 *Dm_build_631) Write(dm_build_645 []byte) (dm_build_646 int, dm_build_647 error) {
	if dm_build_644.dm_build_634 {
		return dm_build_644.dm_build_632.Write(dm_build_645)
	} else {
		return dm_build_644.dm_build_633.Write(dm_build_645)
	}
}

func (dm_build_649 *Dm_build_631) Read(dm_build_650 []byte) (dm_build_651 int, dm_build_652 error) {
	if dm_build_649.dm_build_634 {
		return dm_build_649.dm_build_632.Read(dm_build_650)
	} else {
		return dm_build_649.dm_build_633.Read(dm_build_650)
	}
}

func (dm_build_654 *Dm_build_631) dm_build_653(dm_build_655 net.Conn, dm_build_656 *DmConnection) (*tls.Conn, error) {
	var dm_build_657 = &tls.Config{
		InsecureSkipVerify: true,
	}

	if dm_build_656.sslEncrypt == 1 || dm_build_656.sslEncrypt == 2 {

		if customCA, err := ioutil.ReadFile(dm_build_656.dmConnector.sslCaPath); err == nil {
			dm_build_657.RootCAs, err = x509.SystemCertPool()
			if err != nil {

				dm_build_657.RootCAs = x509.NewCertPool()
			}
			dm_build_657.RootCAs.AppendCertsFromPEM(customCA)
		}

		if dm_build_656.dmConnector.sslCertPath == "" || dm_build_656.dmConnector.sslKeyPath == "" {

			return nil, errors.New("sslPath can not be empty!")
		}

		cer, err := dm_build_654.dm_build_659(dm_build_656.dmConnector.sslCertPath, dm_build_656.dmConnector.sslKeyPath, dm_build_656.dmConnector.sslPwd)
		if err != nil {
			return nil, err
		}

		dm_build_657.Certificates = []tls.Certificate{cer}

		dm_build_657.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {

			certs := make([]*x509.Certificate, len(rawCerts))
			for i, asn1Data := range rawCerts {
				cert, err := x509.ParseCertificate(asn1Data)
				if err != nil {
					return fmt.Errorf("tls: failed to parse certificate from server: %w", err)
				}
				certs[i] = cert
			}

			opts := x509.VerifyOptions{
				Roots:       dm_build_657.RootCAs,
				CurrentTime: time.Now(),

				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range certs[1:] {
				opts.Intermediates.AddCert(cert)
			}
			if _, err := certs[0].Verify(opts); err != nil {
				return err
			}
			return nil
		}
	}

	dm_build_658 := tls.Client(dm_build_655, dm_build_657)
	if err := dm_build_658.Handshake(); err != nil {
		return nil, err
	}
	return dm_build_658, nil

}

func (dm_build_660 *Dm_build_631) dm_build_659(dm_build_661, dm_build_662, dm_build_663 string) (tls.Certificate, error) {

	dm_build_664, err := ioutil.ReadFile(dm_build_661)
	if err != nil {
		return tls.Certificate{}, err
	}
	dm_build_666, err := ioutil.ReadFile(dm_build_662)
	if err != nil {
		return tls.Certificate{}, err
	}

	dm_build_668, _ := pem.Decode(dm_build_666)
	if dm_build_668 == nil {
		return tls.Certificate{}, fmt.Errorf("tls: failed to find any PEM data in key input")
	}

	if x509.IsEncryptedPEMBlock(dm_build_668) {

		dm_build_666, err = x509.DecryptPEMBlock(dm_build_668, []byte(dm_build_663))
		if err != nil {
			return tls.Certificate{}, err
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(dm_build_666)
		if err != nil {

			key, err2 := x509.ParsePKCS8PrivateKey(dm_build_666)
			if err2 != nil {
				return tls.Certificate{}, err2
			}
			dm_build_666, _ = x509.MarshalPKCS8PrivateKey(key)
		} else {
			dm_build_666 = x509.MarshalPKCS1PrivateKey(privateKey)
		}

		dm_build_666 = pem.EncodeToMemory(&pem.Block{
			Type:  dm_build_668.Type,
			Bytes: dm_build_666,
		})
	} else if dm_build_668.Type == "ENCRYPTED PRIVATE KEY" {

		return tls.Certificate{}, fmt.Errorf("PKCS#8 didn't supported yet")
	}

	dm_build_669, err := tls.X509KeyPair(dm_build_664, dm_build_666)
	if err != nil {
		return tls.Certificate{}, err
	}

	return dm_build_669, nil
}
